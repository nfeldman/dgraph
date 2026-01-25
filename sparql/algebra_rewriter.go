package sparql

import (
	"fmt"
	"sort"
	"strings"
)

// AlgebraRewriteRule represents a single algebra rewrite rule.
type AlgebraRewriteRule interface {
	Name() string
	Apply(expr AlgebraExpr) (AlgebraExpr, bool)
}

// AlgebraOptimizer runs a sequence of rewrite rules to a fixed point.
type AlgebraOptimizer struct {
	rules         []AlgebraRewriteRule
	maxIterations int
}

// NewAlgebraOptimizer creates an optimizer with the provided rules.
// maxIterations defaults to 10 to avoid infinite rewrite loops.
func NewAlgebraOptimizer(rules ...AlgebraRewriteRule) *AlgebraOptimizer {
	return &AlgebraOptimizer{
		rules:         rules,
		maxIterations: 10,
	}
}

// Optimize applies rewrite rules repeatedly until no rule makes further changes
// or until maxIterations is reached.
func (o *AlgebraOptimizer) Optimize(expr AlgebraExpr) AlgebraExpr {
	if o == nil {
		return expr
	}
	current := expr
	for iter := 0; iter < o.maxIterations; iter++ {
		changed := false
		for _, rule := range o.rules {
			next, ok := rule.Apply(current)
			if ok {
				current = next
				changed = true
			}
		}
		if !changed {
			break
		}
	}
	return current
}

// -----------------------------------------------------------------------------
// Filter Pushdown Rule
// -----------------------------------------------------------------------------

type filterPushdownRule struct{}

func (filterPushdownRule) Name() string { return "filter-pushdown" }

func (filterPushdownRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return pushDownFilters(expr)
}

func pushDownFilters(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraFilter:
		// Try to push inside before handling children
		pushed, ok := pushFilterIntoChild(n)
		if ok {
			return pushed, true
		}
		// Otherwise, recurse into child
		child, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraJoin:
		left, lc := pushDownFilters(n.Left)
		right, rc := pushDownFilters(n.Right)
		if lc {
			n.Left = left
		}
		if rc {
			n.Right = right
		}
		return n, lc || rc
	case *AlgebraLeftJoin:
		input, ic := pushDownFilters(n.Input)
		if ic {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := pushDownFilters(alt)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	case *AlgebraProject:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraAgg:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraBind:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraDistinct:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraOrderBy:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraLimit:
		input, changed := pushDownFilters(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraValues:
		return n, false
	case *AlgebraEmpty:
		return n, false
	default:
		return expr, false
	}
}

func pushFilterIntoChild(f *AlgebraFilter) (AlgebraExpr, bool) {
	refs := extractVariableReferences(f.Expr)
	switch c := f.Input.(type) {
	case *AlgebraJoin:
		leftVars := toSet(c.Left.Variables())
		rightVars := toSet(c.Right.Variables())
		refsSet := toSet(refs)
		if subsetOf(refsSet, leftVars) {
			return &AlgebraJoin{Left: &AlgebraFilter{Expr: f.Expr, Input: c.Left}, Right: c.Right}, true
		}
		if subsetOf(refsSet, rightVars) {
			return &AlgebraJoin{Left: c.Left, Right: &AlgebraFilter{Expr: f.Expr, Input: c.Right}}, true
		}
	case *AlgebraUnion:
		alts := make([]AlgebraExpr, len(c.Alternatives))
		for i, alt := range c.Alternatives {
			alts[i] = &AlgebraFilter{Expr: f.Expr, Input: alt}
		}
		return &AlgebraUnion{Alternatives: alts}, true
	case *AlgebraLeftJoin:
		mainVars := toSet(c.Input.Variables())
		optVars := toSet(getTripleVars(c.Patterns))
		refsSet := toSet(refs)
		// If filter only touches main variables, push into input
		if subsetOf(refsSet, mainVars) {
			return &AlgebraLeftJoin{Input: &AlgebraFilter{Expr: f.Expr, Input: c.Input}, Patterns: c.Patterns, Filter: c.Filter}, true
		}
		// If filter only touches optional variables, merge with optional filter
		if subsetOf(refsSet, optVars) {
			merged := c.Filter
			if merged == "" {
				merged = f.Expr
			} else {
				merged = combineFilters(merged, f.Expr)
			}
			return &AlgebraLeftJoin{Input: c.Input, Patterns: c.Patterns, Filter: merged}, true
		}
	}
	return f, false
}

func subsetOf(a, b map[string]bool) bool {
	for k := range a {
		if !b[k] {
			return false
		}
	}
	return true
}

func toSet(vars []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range vars {
		m[v] = true
	}
	return m
}

func combineFilters(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	return "(" + a + ") && (" + b + ")"
}

// -----------------------------------------------------------------------------
// Join Reordering Rule (greedy by estimated cost)
// -----------------------------------------------------------------------------

type joinReorderRule struct{}

func (joinReorderRule) Name() string { return "join-reorder" }

func (joinReorderRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return reorderJoins(expr)
}

func reorderJoins(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraJoin:
		flat := flattenJoins(n)
		// Recursively reorder children first
		changedChild := false
		for i, part := range flat {
			newPart, ok := reorderJoins(part)
			if ok {
				flat[i] = newPart
				changedChild = true
			}
		}
		sort.SliceStable(flat, func(i, j int) bool {
			return estimateCost(flat[i]) < estimateCost(flat[j])
		})
		rebuilt := buildJoinChain(flat)
		if !sameJoinStructure(n, rebuilt) {
			return rebuilt, true
		}
		return rebuilt, changedChild
	case *AlgebraFilter:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraProject:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraDistinct:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraOrderBy:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraLimit:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraAgg:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraBind:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraLeftJoin:
		input, changed := reorderJoins(n.Input)
		if changed {
			n.Input = input
			return n, true
		}
		return n, false
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := reorderJoins(alt)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	case *AlgebraValues:
		return n, false
	case *AlgebraEmpty:
		return n, false
	default:
		return expr, false
	}
}

func flattenJoins(expr AlgebraExpr) []AlgebraExpr {
	switch j := expr.(type) {
	case *AlgebraJoin:
		left := flattenJoins(j.Left)
		right := flattenJoins(j.Right)
		return append(left, right...)
	default:
		return []AlgebraExpr{expr}
	}
}

func buildJoinChain(parts []AlgebraExpr) AlgebraExpr {
	if len(parts) == 0 {
		return nil
	}
	if len(parts) == 1 {
		return parts[0]
	}
	result := parts[0]
	for i := 1; i < len(parts); i++ {
		result = &AlgebraJoin{Left: result, Right: parts[i]}
	}
	return result
}

func estimateCost(expr AlgebraExpr) int {
	switch n := expr.(type) {
	case *AlgebraBGP:
		if len(n.Triples) == 0 {
			return 1
		}
		return len(n.Triples)
	case *AlgebraValues:
		return 1 + len(n.Rows)
	case *AlgebraEmpty:
		return 1
	case *AlgebraFilter:
		return estimateCost(n.Input)
	case *AlgebraProject:
		return estimateCost(n.Input)
	case *AlgebraDistinct:
		return estimateCost(n.Input) + 1
	case *AlgebraOrderBy:
		return estimateCost(n.Input) + 2
	case *AlgebraLimit:
		return estimateCost(n.Input)
	case *AlgebraAgg:
		return estimateCost(n.Input) + 3
	case *AlgebraBind:
		return estimateCost(n.Input)
	case *AlgebraLeftJoin:
		return estimateCost(n.Input) + len(n.Patterns)
	case *AlgebraUnion:
		total := 0
		for _, alt := range n.Alternatives {
			total += estimateCost(alt)
		}
		if total == 0 {
			return 10
		}
		return total
	default:
		return 10
	}
}

func sameJoinStructure(a, b AlgebraExpr) bool {
	aj, ok1 := a.(*AlgebraJoin)
	bj, ok2 := b.(*AlgebraJoin)
	if !ok1 || !ok2 {
		return a == b
	}
	return sameJoinStructure(aj.Left, bj.Left) && sameJoinStructure(aj.Right, bj.Right)
}

// -----------------------------------------------------------------------------
// OPTIONAL Simplification & Dead Variable Elimination Rule
// -----------------------------------------------------------------------------

type deadVariableEliminationRule struct{}

func (deadVariableEliminationRule) Name() string { return "dead-var-elimination" }

func (deadVariableEliminationRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	required := inferRequiredVars(expr)
	return eliminateDead(expr, required)
}

func inferRequiredVars(expr AlgebraExpr) map[string]bool {
	switch n := expr.(type) {
	case *AlgebraProject:
		return toSet(n.Vars)
	default:
		return toSet(expr.Variables())
	}
}

func eliminateDead(expr AlgebraExpr, required map[string]bool) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraProject:
		childRequired := toSet(n.Vars)
		child, changed := eliminateDead(n.Input, childRequired)
		if changed {
			n.Input = child
		}
		// Drop projected variables not defined (defensive cleanup)
		filtered := make([]string, 0, len(n.Vars))
		defined := toSet(child.DefinedVars())
		for _, v := range n.Vars {
			if defined[v] {
				filtered = append(filtered, v)
			}
		}
		if len(filtered) != len(n.Vars) {
			n.Vars = filtered
			changed = true
		}
		return n, changed
	case *AlgebraFilter:
		refs := toSet(extractVariableReferences(n.Expr))
		childRequired := unionSets(required, refs)
		child, changed := eliminateDead(n.Input, childRequired)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraOrderBy:
		refs := toSet(extractVariableReferences(strings.Join(n.Exprs, " ")))
		childRequired := unionSets(required, refs)
		child, changed := eliminateDead(n.Input, childRequired)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraDistinct:
		child, changed := eliminateDead(n.Input, required)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraLimit:
		child, changed := eliminateDead(n.Input, required)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraAgg:
		req := unionSets(required, toSet(n.Group))
		if n.Var != "" {
			req[n.Var] = true
		}
		child, changed := eliminateDead(n.Input, req)
		if changed {
			n.Input = child
		}
		// After aggregation, only group + agg var remain
		n.DefinedVars()
		return n, changed
	case *AlgebraBind:
		// If bound variable is not required, drop the bind
		if !required[n.Var] {
			child, changed := eliminateDead(n.Input, required)
			if changed {
				return child, true
			}
			return child, true
		}
		child, changed := eliminateDead(n.Input, required)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraLeftJoin:
		// Simplify only if optional contributes nothing (no patterns and no filter)
		if len(n.Patterns) == 0 && n.Filter == "" {
			return eliminateDead(n.Input, required)
		}
		child, changed := eliminateDead(n.Input, required)
		if changed {
			n.Input = child
			return n, true
		}
		return n, false
	case *AlgebraJoin:
		left, lch := eliminateDead(n.Left, required)
		right, rch := eliminateDead(n.Right, required)
		if lch {
			n.Left = left
		}
		if rch {
			n.Right = right
		}
		return n, lch || rch
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := eliminateDead(alt, required)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	default:
		return expr, false
	}
}

func unionSets(a, b map[string]bool) map[string]bool {
	out := make(map[string]bool)
	for k := range a {
		out[k] = true
	}
	for k := range b {
		out[k] = true
	}
	return out
}

func intersectionEmpty(a, b map[string]bool) bool {
	for k := range a {
		if b[k] {
			return false
		}
	}
	return true
}

// -----------------------------------------------------------------------------
// Identity & Deduplication Simplification
// -----------------------------------------------------------------------------

type identitySimplificationRule struct{}

func (identitySimplificationRule) Name() string { return "identity-simplification" }

func (identitySimplificationRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return simplifyIdentities(expr)
}

func simplifyIdentities(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraFilter:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		norm := strings.TrimSpace(n.Expr)
		if norm == "" || strings.EqualFold(norm, "true") {
			return child, true
		}
		return n, changed
	case *AlgebraJoin:
		left, lc := simplifyIdentities(n.Left)
		right, rc := simplifyIdentities(n.Right)
		n.Left, n.Right = left, right
		if isEmptyExpr(left) || isEmptyExpr(right) {
			return &AlgebraEmpty{}, true
		}
		return n, lc || rc
	case *AlgebraLeftJoin:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return &AlgebraEmpty{}, true
		}
		if len(n.Patterns) == 0 && strings.TrimSpace(n.Filter) == "" {
			return child, true
		}
		return n, changed
	case *AlgebraUnion:
		changed := false
		alts := make([]AlgebraExpr, 0, len(n.Alternatives))
		seen := make(map[string]bool)
		for _, alt := range n.Alternatives {
			simplified, ok := simplifyIdentities(alt)
			if ok {
				changed = true
			}
			if isEmptyExpr(simplified) {
				changed = true
				continue
			}
			key := fmt.Sprintf("%T:%s", simplified, simplified.String())
			if seen[key] {
				changed = true
				continue
			}
			seen[key] = true
			alts = append(alts, simplified)
		}
		if len(alts) == 0 {
			return &AlgebraEmpty{}, true
		}
		if len(alts) == 1 {
			return alts[0], true
		}
		n.Alternatives = alts
		return n, changed
	case *AlgebraProject:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		if inner, ok := child.(*AlgebraProject); ok {
			n.Input = inner.Input
			changed = true
		}
		return n, changed
	case *AlgebraDistinct:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		if _, ok := child.(*AlgebraDistinct); ok {
			return child, true
		}
		return n, changed
	case *AlgebraOrderBy:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		return n, changed
	case *AlgebraLimit:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if n.Count == 0 {
			return &AlgebraEmpty{}, true
		}
		if isEmptyExpr(child) {
			return child, true
		}
		return n, changed
	case *AlgebraAgg:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		return n, changed
	case *AlgebraBind:
		child, changed := simplifyIdentities(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return child, true
		}
		return n, changed
	case *AlgebraValues:
		dedupValuesRows(n)
		if len(n.Rows) == 0 {
			return &AlgebraEmpty{}, true
		}
		return n, false
	case *AlgebraEmpty:
		return n, false
	case *AlgebraBGP:
		return n, false
	default:
		return expr, false
	}
}

func dedupValuesRows(v *AlgebraValues) {
	if len(v.Rows) == 0 {
		return
	}
	seen := make(map[string]bool)
	uniq := make([]map[string]string, 0, len(v.Rows))
	for _, row := range v.Rows {
		key := buildRowKey(v.Vars, row)
		if seen[key] {
			continue
		}
		seen[key] = true
		uniq = append(uniq, row)
	}
	v.Rows = uniq
}

func buildRowKey(vars []string, row map[string]string) string {
	parts := make([]string, len(vars))
	for i, v := range vars {
		parts[i] = fmt.Sprintf("%s=%s", v, row[v])
	}
	return strings.Join(parts, ";")
}

func isEmptyExpr(expr AlgebraExpr) bool {
	switch e := expr.(type) {
	case *AlgebraEmpty:
		return true
	case *AlgebraValues:
		return len(e.Rows) == 0
	default:
		return false
	}
}

// -----------------------------------------------------------------------------
// VALUES simplification & projection propagation
// -----------------------------------------------------------------------------

type valuesPropagationRule struct{}

func (valuesPropagationRule) Name() string { return "values-propagation" }

func (valuesPropagationRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return propagateValues(expr)
}

func propagateValues(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraProject:
		child, changed := propagateValues(n.Input)
		n.Input = child
		if val, ok := child.(*AlgebraValues); ok {
			projected := restrictValues(val, n.Vars)
			dedupValuesRows(projected)
			return projected, true
		}
		return n, changed
	case *AlgebraFilter:
		child, changed := propagateValues(n.Input)
		n.Input = child
		if val, ok := child.(*AlgebraValues); ok {
			norm := strings.TrimSpace(n.Expr)
			if norm == "" || strings.EqualFold(norm, "true") {
				return val, true
			}
			if strings.EqualFold(norm, "false") {
				return &AlgebraEmpty{}, true
			}
		}
		return n, changed
	case *AlgebraJoin:
		left, lc := propagateValues(n.Left)
		right, rc := propagateValues(n.Right)
		n.Left, n.Right = left, right
		if isEmptyExpr(left) || isEmptyExpr(right) {
			return &AlgebraEmpty{}, true
		}
		return n, lc || rc
	case *AlgebraLeftJoin:
		child, changed := propagateValues(n.Input)
		n.Input = child
		if isEmptyExpr(child) {
			return &AlgebraEmpty{}, true
		}
		return n, changed
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := propagateValues(alt)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	case *AlgebraDistinct:
		child, changed := propagateValues(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraOrderBy:
		child, changed := propagateValues(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraLimit:
		child, changed := propagateValues(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraAgg:
		child, changed := propagateValues(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraBind:
		child, changed := propagateValues(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraValues, *AlgebraEmpty, *AlgebraBGP:
		return n, false
	default:
		return expr, false
	}
}

func restrictValues(v *AlgebraValues, vars []string) *AlgebraValues {
	trimmed := make([]map[string]string, 0, len(v.Rows))
	for _, row := range v.Rows {
		newRow := make(map[string]string)
		for _, proj := range vars {
			if val, ok := row[proj]; ok {
				newRow[proj] = val
			}
		}
		trimmed = append(trimmed, newRow)
	}
	return &AlgebraValues{Vars: vars, Rows: trimmed}
}

// -----------------------------------------------------------------------------
// Filter normalization
// -----------------------------------------------------------------------------

type filterNormalizationRule struct{}

func (filterNormalizationRule) Name() string { return "filter-normalization" }

func (filterNormalizationRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return normalizeFilters(expr)
}

func normalizeFilters(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraFilter:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		norm := normalizeExprString(n.Expr)
		if norm != n.Expr {
			n.Expr = norm
			changed = true
		}
		return n, changed
	case *AlgebraLeftJoin:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		norm := normalizeExprString(n.Filter)
		if norm != n.Filter {
			n.Filter = norm
			changed = true
		}
		return n, changed
	case *AlgebraJoin:
		left, lc := normalizeFilters(n.Left)
		right, rc := normalizeFilters(n.Right)
		n.Left, n.Right = left, right
		return n, lc || rc
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := normalizeFilters(alt)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	case *AlgebraProject:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraDistinct:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraOrderBy:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraLimit:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraAgg:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraBind:
		child, changed := normalizeFilters(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraValues, *AlgebraEmpty, *AlgebraBGP:
		return n, false
	default:
		return expr, false
	}
}

func normalizeExprString(expr string) string {
	trimmed := strings.TrimSpace(expr)
	for len(trimmed) > 1 && strings.HasPrefix(trimmed, "(") && strings.HasSuffix(trimmed, ")") {
		inner := strings.TrimSpace(trimmed[1 : len(trimmed)-1])
		if inner == "" {
			break
		}
		if !balancedParens(inner) {
			break
		}
		trimmed = inner
	}
	return trimmed
}

func balancedParens(expr string) bool {
	depth := 0
	for _, r := range expr {
		switch r {
		case '(':
			depth++
		case ')':
			depth--
			if depth < 0 {
				return false
			}
		}
	}
	return depth == 0
}

// -----------------------------------------------------------------------------
// Property path desugaring (simple sequence p1/p2)
// -----------------------------------------------------------------------------

type pathDesugarRule struct{}

func (pathDesugarRule) Name() string { return "path-desugar" }

func (pathDesugarRule) Apply(expr AlgebraExpr) (AlgebraExpr, bool) {
	return desugarPathExpressions(expr)
}

func desugarPathExpressions(expr AlgebraExpr) (AlgebraExpr, bool) {
	switch n := expr.(type) {
	case *AlgebraBGP:
		expanded := make([]*Triple, 0, len(n.Triples))
		changed := false
		for idx, t := range n.Triples {
			triples, ok := expandPathTriple(t, idx)
			if ok {
				changed = true
			}
			expanded = append(expanded, triples...)
		}
		if !changed {
			return n, false
		}
		return &AlgebraBGP{Triples: expanded}, true
	case *AlgebraJoin:
		left, lc := desugarPathExpressions(n.Left)
		right, rc := desugarPathExpressions(n.Right)
		n.Left, n.Right = left, right
		return n, lc || rc
	case *AlgebraFilter:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraLeftJoin:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraUnion:
		changed := false
		for i, alt := range n.Alternatives {
			newAlt, ok := desugarPathExpressions(alt)
			if ok {
				n.Alternatives[i] = newAlt
				changed = true
			}
		}
		return n, changed
	case *AlgebraProject:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraAgg:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraBind:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraDistinct:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraOrderBy:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraLimit:
		child, changed := desugarPathExpressions(n.Input)
		n.Input = child
		return n, changed
	case *AlgebraValues, *AlgebraEmpty:
		return n, false
	default:
		return expr, false
	}
}

func expandPathTriple(t *Triple, idx int) ([]*Triple, bool) {
	parts, ok := splitPathPredicate(t.Predicate)
	if !ok || len(parts) <= 1 {
		return []*Triple{t}, false
	}
	triples := make([]*Triple, 0, len(parts))
	currentSubject := t.Subject
	for i, pred := range parts {
		obj := t.Object
		objIsVar := t.ObjectIsVar
		if i < len(parts)-1 {
			obj = fmt.Sprintf("?path_%d_%d", idx, i)
			objIsVar = true
		}
		triples = append(triples, &Triple{Subject: currentSubject, Predicate: pred, Object: obj, ObjectIsVar: objIsVar, Graph: t.Graph})
		currentSubject = obj
	}
	return triples, true
}

func splitPathPredicate(pred string) ([]string, bool) {
	if strings.Contains(pred, "://") {
		return nil, false
	}
	parts := strings.Split(pred, "/")
	if len(parts) <= 1 {
		return nil, false
	}
	clean := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p == "" {
			return nil, false
		}
		clean = append(clean, p)
	}
	return clean, true
}
