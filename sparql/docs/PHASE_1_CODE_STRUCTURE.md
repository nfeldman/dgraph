# Phase 1 Code Structure Guide

**Purpose**: Concrete code structure examples and patterns for Phase 1 implementation.

---

## 1. Algebra Types File Structure (`sparql/algebra.go`)

### File Organization

```go
package sparql

import (
	"fmt"
	"regexp"
	"strings"
	// standard imports
)

// ============================================================================
// SECTION 1: AlgebraExpr Interface
// ============================================================================

// AlgebraExpr is the base interface for all algebra operations.
// Each concrete operator implements this interface to support the visitor pattern.
type AlgebraExpr interface {
	// Accept implements the Visitor pattern for traversal and transformation.
	Accept(visitor AlgebraVisitor) interface{}

	// String returns a human-readable string representation for debugging.
	// Format should be recursive: Operator(child1, child2, ...)
	String() string
}

// ============================================================================
// SECTION 2: Concrete Algebra Operators
// ============================================================================

// BGP (Basic Graph Pattern) represents a set of triple patterns.
// All triples in a BGP must be matched together.
type BGP struct {
	Triples []*Triple
}

func (bgp *BGP) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitBGP(bgp)
}

func (bgp *BGP) String() string {
	if len(bgp.Triples) == 0 {
		return "BGP()"
	}
	parts := make([]string, len(bgp.Triples))
	for i, t := range bgp.Triples {
		// Format: "?subj:pred:?obj"
		obj := t.Object
		if !t.ObjectIsVar {
			obj = fmt.Sprintf("<%s>", t.Object) // Wrap IRIs
		}
		parts[i] = fmt.Sprintf("%s:%s:%s", t.Subject, t.Predicate, obj)
	}
	return fmt.Sprintf("BGP([%s])", strings.Join(parts, ", "))
}

// Join represents the combination of two patterns.
// Results from Left and Right are joined on shared variables.
type Join struct {
	Left  AlgebraExpr
	Right AlgebraExpr
}

func (j *Join) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitJoin(j)
}

func (j *Join) String() string {
	return fmt.Sprintf("Join(%s, %s)", j.Left.String(), j.Right.String())
}

// Filter represents a WHERE constraint applied to results.
// Only solutions where Expr evaluates to true pass through.
type Filter struct {
	Expr  string       // SPARQL filter expression (e.g., "?name = 'Alice'")
	Input AlgebraExpr  // Patterns being filtered
}

func (f *Filter) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitFilter(f)
}

func (f *Filter) String() string {
	return fmt.Sprintf("Filter(%s, %s)", f.Expr, f.Input.String())
}

// LeftJoin represents OPTIONAL patterns.
// Solutions from Input are kept; optional Patterns are added if they match.
type LeftJoin struct {
	Input    AlgebraExpr
	Patterns []*Triple
	Filter   string
}

func (lj *LeftJoin) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitLeftJoin(lj)
}

func (lj *LeftJoin) String() string {
	pats := make([]string, len(lj.Patterns))
	for i, t := range lj.Patterns {
		obj := t.Object
		if !t.ObjectIsVar {
			obj = fmt.Sprintf("<%s>", t.Object)
		}
		pats[i] = fmt.Sprintf("%s:%s:%s", t.Subject, t.Predicate, obj)
	}

	filterStr := ""
	if lj.Filter != "" {
		filterStr = fmt.Sprintf(", filter=%s", lj.Filter)
	}

	return fmt.Sprintf("LeftJoin(%s, [%s]%s)",
		lj.Input.String(), strings.Join(pats, ", "), filterStr)
}

// Union represents UNION alternatives.
// Results are the union of solutions from all alternatives.
type Union struct {
	Alternatives []AlgebraExpr
}

func (u *Union) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitUnion(u)
}

func (u *Union) String() string {
	parts := make([]string, len(u.Alternatives))
	for i, alt := range u.Alternatives {
		parts[i] = alt.String()
	}
	return fmt.Sprintf("Union(%s)", strings.Join(parts, " | "))
}

// Project represents SELECT variable projection.
// Only specified variables are returned to the user.
type Project struct {
	Vars  []string
	Input AlgebraExpr
}

func (p *Project) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitProject(p)
}

func (p *Project) String() string {
	return fmt.Sprintf("Project([%s], %s)",
		strings.Join(p.Vars, ", "), p.Input.String())
}

// Aggregate represents aggregation functions (COUNT, SUM, MIN, MAX, AVG).
// Results are grouped by Group variables; others are aggregated.
type Aggregate struct {
	Op    string       // "COUNT", "SUM", "MIN", "MAX", "AVG"
	Expr  string       // Expression to aggregate (e.g., "*" for COUNT(*))
	Group []string     // GROUP BY variables
	Input AlgebraExpr
}

func (a *Aggregate) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAggregate(a)
}

func (a *Aggregate) String() string {
	if len(a.Group) == 0 {
		return fmt.Sprintf("Aggregate(%s(%s), %s)",
			a.Op, a.Expr, a.Input.String())
	}
	return fmt.Sprintf("Aggregate(%s(%s), GROUP_BY[%s], %s)",
		a.Op, a.Expr, strings.Join(a.Group, ", "), a.Input.String())
}

// Bind represents BIND expressions that compute new variables.
// Example: BIND (?age - 2024 AS ?yearBorn)
type Bind struct {
	Var   string       // Output variable name (e.g., "?age")
	Expr  string       // Expression to evaluate
	Input AlgebraExpr
}

func (b *Bind) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitBind(b)
}

func (b *Bind) String() string {
	return fmt.Sprintf("Bind(%s AS %s, %s)",
		b.Expr, b.Var, b.Input.String())
}

// Distinct represents DISTINCT modifier.
// Removes duplicate solution bindings.
type Distinct struct {
	Input AlgebraExpr
}

func (d *Distinct) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitDistinct(d)
}

func (d *Distinct) String() string {
	return fmt.Sprintf("Distinct(%s)", d.Input.String())
}

// OrderBy represents ORDER BY clause.
// Solutions are sorted by specified expressions and directions.
type OrderBy struct {
	Expressions []string  // Variable names or expressions
	Ascending   []bool    // Sort direction for each expression
	Input       AlgebraExpr
}

func (ob *OrderBy) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitOrderBy(ob)
}

func (ob *OrderBy) String() string {
	parts := make([]string, len(ob.Expressions))
	for i, expr := range ob.Expressions {
		dir := "ASC"
		if !ob.Ascending[i] {
			dir = "DESC"
		}
		parts[i] = fmt.Sprintf("%s %s", expr, dir)
	}
	return fmt.Sprintf("OrderBy([%s], %s)",
		strings.Join(parts, ", "), ob.Input.String())
}

// Limit represents LIMIT and OFFSET clauses.
// Returns at most Count solutions, skipping first Offset.
type Limit struct {
	Count  int
	Offset int
	Input  AlgebraExpr
}

func (l *Limit) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitLimit(l)
}

func (l *Limit) String() string {
	if l.Offset == 0 {
		return fmt.Sprintf("Limit(count=%d, %s)",
			l.Count, l.Input.String())
	}
	return fmt.Sprintf("Limit(count=%d, offset=%d, %s)",
		l.Count, l.Offset, l.Input.String())
}

// ============================================================================
// SECTION 3: AST to Algebra Conversion
// ============================================================================

// ASTToAlgebra converts a SPARQL AST query to an algebra expression.
// This is the main entry point for converting parsed queries.
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	if query.Qtype != "SELECT" && query.Qtype != "ASK" {
		return nil, fmt.Errorf("unsupported query type: %s (only SELECT/ASK)", query.Qtype)
	}

	// Step 1: Extract patterns
	patterns := query.Patterns
	if len(patterns) == 0 && len(query.Bgps) > 0 {
		// Backwards compatibility: convert old BGP format
		for _, bgp := range query.Bgps {
			patterns = append(patterns, bgp)
		}
	}

	if len(patterns) == 0 {
		return nil, fmt.Errorf("no patterns in query")
	}

	// Step 2: Separate filter patterns from others
	filterPatterns, otherPatterns := separateFilterPatterns(patterns)

	// Step 3: Build algebra from non-filter patterns
	algebra, err := convertPatterns(otherPatterns)
	if err != nil {
		return nil, fmt.Errorf("converting patterns: %w", err)
	}

	// Step 4: Apply filters
	algebra = applyFilters(algebra, filterPatterns)

	// Step 5: Apply BIND expressions
	for _, bind := range query.Binds {
		algebra = &Bind{
			Var:   bind.Variable,
			Expr:  bind.Expression,
			Input: algebra,
		}
	}

	// Step 6: Apply aggregates (if present)
	if len(query.Aggregates) > 0 {
		// Extract GROUP BY variables (implicit: non-aggregated variables)
		groupVars := extractGroupByVars(query.Aggregates, query.Projs)

		// Create Aggregate node (use first aggregate; extend if multiple needed)
		agg := query.Aggregates[0]
		algebra = &Aggregate{
			Op:    strings.ToUpper(agg.Function),
			Expr:  agg.Variable,
			Group: groupVars,
			Input: algebra,
		}
	}

	// Step 7: Apply DISTINCT
	if query.Distinct {
		algebra = &Distinct{Input: algebra}
	}

	// Step 8: Apply ORDER BY
	if len(query.OrderBy) > 0 {
		algebra = applyOrderBy(algebra, query.OrderBy)
	}

	// Step 9: Apply LIMIT/OFFSET
	if query.Limit > 0 || query.Offset > 0 {
		algebra = &Limit{
			Count:  query.Limit,
			Offset: query.Offset,
			Input:  algebra,
		}
	}

	// Step 10: Apply projection (outermost)
	if len(query.Projs) > 0 {
		algebra = &Project{
			Vars:  query.Projs,
			Input: algebra,
		}
	}

	return algebra, nil
}

// convertPatterns recursively converts GraphPattern list to algebra.
// Multiple patterns are combined with Join.
func convertPatterns(patterns []GraphPattern) (AlgebraExpr, error) {
	if len(patterns) == 0 {
		return nil, fmt.Errorf("no patterns to convert")
	}

	if len(patterns) == 1 {
		return convertSinglePattern(patterns[0])
	}

	// Multiple patterns: Join them
	left, err := convertPatterns(patterns[:len(patterns)-1])
	if err != nil {
		return nil, err
	}

	right, err := convertSinglePattern(patterns[len(patterns)-1])
	if err != nil {
		return nil, err
	}

	return &Join{Left: left, Right: right}, nil
}

// convertSinglePattern converts a single GraphPattern to algebra.
func convertSinglePattern(p GraphPattern) (AlgebraExpr, error) {
	switch pattern := p.(type) {
	case *BGP:
		return &BGP{Triples: pattern.Triples}, nil

	case *OptionalPattern:
		// Extract triples from optional patterns
		triples := extractTriplesFromPatterns(pattern.Patterns)
		return &LeftJoin{
			Patterns: triples,
			Filter:   "",
		}, nil

	case *UnionPattern:
		// Convert each alternative
		alternatives := make([]AlgebraExpr, len(pattern.Alternatives))
		for i, alt := range pattern.Alternatives {
			algebra, err := convertPatterns(alt)
			if err != nil {
				return nil, fmt.Errorf("converting union alternative %d: %w", i, err)
			}
			alternatives[i] = algebra
		}
		return &Union{Alternatives: alternatives}, nil

	case *FilterPattern:
		// FilterPattern is handled separately, not here
		return nil, fmt.Errorf("FilterPattern should not appear as standalone pattern")

	default:
		return nil, fmt.Errorf("unknown pattern type: %T", p)
	}
}

// separateFilterPatterns splits patterns into filters and others.
func separateFilterPatterns(patterns []GraphPattern) ([]FilterPattern, []GraphPattern) {
	var filters []FilterPattern
	var others []GraphPattern

	for _, p := range patterns {
		if fp, ok := p.(*FilterPattern); ok {
			filters = append(filters, *fp)
		} else {
			others = append(others, p)
		}
	}

	return filters, others
}

// applyFilters wraps algebra with Filter nodes for each filter expression.
func applyFilters(algebra AlgebraExpr, filters []FilterPattern) AlgebraExpr {
	// Apply filters in reverse order so they nest correctly
	for i := len(filters) - 1; i >= 0; i-- {
		algebra = &Filter{
			Expr:  filters[i].Expression,
			Input: algebra,
		}
	}
	return algebra
}

// extractTriplesFromPatterns flattens patterns into a triple list.
func extractTriplesFromPatterns(patterns []GraphPattern) []*Triple {
	var triples []*Triple
	for _, p := range patterns {
		if bgp, ok := p.(*BGP); ok {
			triples = append(triples, bgp.Triples...)
		}
	}
	return triples
}

// applyOrderBy converts OrderBy variable list to OrderBy algebra.
func applyOrderBy(algebra AlgebraExpr, orderVars []string) AlgebraExpr {
	expressions := make([]string, 0)
	ascending := make([]bool, 0)

	for _, orderVar := range orderVars {
		// Parse "?var" or "DESC(?var)" or "ASC(?var)"
		if strings.HasPrefix(strings.ToUpper(orderVar), "DESC(") {
			expr := orderVar[5 : len(orderVar)-1]
			expressions = append(expressions, expr)
			ascending = append(ascending, false)
		} else if strings.HasPrefix(strings.ToUpper(orderVar), "ASC(") {
			expr := orderVar[4 : len(orderVar)-1]
			expressions = append(expressions, expr)
			ascending = append(ascending, true)
		} else {
			// Assume ASC by default
			expressions = append(expressions, orderVar)
			ascending = append(ascending, true)
		}
	}

	return &OrderBy{
		Expressions: expressions,
		Ascending:   ascending,
		Input:       algebra,
	}
}

// extractGroupByVars extracts implicit GROUP BY variables.
// In SPARQL: non-aggregated variables in SELECT are GROUP BY.
func extractGroupByVars(aggregates []*Aggregate, projVars []string) []string {
	if len(projVars) == 0 {
		return []string{}
	}

	// GROUP BY = all projected variables except aggregates
	var groupVars []string
	aggVarMap := make(map[string]bool)
	for _, agg := range aggregates {
		aggVarMap[agg.Variable] = true
	}

	for _, proj := range projVars {
		if !aggVarMap[proj] && strings.HasPrefix(proj, "?") {
			groupVars = append(groupVars, proj)
		}
	}

	return groupVars
}

// ============================================================================
// SECTION 4: Algebra Validator
// ============================================================================

// AlgebraValidator validates algebra expressions for semantic correctness.
type AlgebraValidator struct {
	definedVars map[string]bool
	errors      []string
}

// Validate checks an algebra expression tree for semantic correctness.
// Returns error if any validation issues found.
func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
	if expr == nil {
		return fmt.Errorf("algebra expression is nil")
	}

	v.definedVars = make(map[string]bool)
	v.errors = make([]string, 0)

	// Traverse and collect errors
	v.validateExpr(expr)

	if len(v.errors) > 0 {
		return fmt.Errorf("algebra validation errors:\n  %s",
			strings.Join(v.errors, "\n  "))
	}

	return nil
}

// validateExpr dispatches to specific validator based on expression type.
func (v *AlgebraValidator) validateExpr(expr AlgebraExpr) {
	if expr == nil {
		return
	}

	switch e := expr.(type) {
	case *BGP:
		v.validateBGP(e)
	case *Join:
		v.validateJoin(e)
	case *Filter:
		v.validateFilter(e)
	case *LeftJoin:
		v.validateLeftJoin(e)
	case *Union:
		v.validateUnion(e)
	case *Project:
		v.validateProject(e)
	case *Aggregate:
		v.validateAggregate(e)
	case *Bind:
		v.validateBind(e)
	case *Distinct:
		v.validateDistinct(e)
	case *OrderBy:
		v.validateOrderBy(e)
	case *Limit:
		v.validateLimit(e)
	}
}

func (v *AlgebraValidator) validateBGP(bgp *BGP) {
	// BGP defines all variables in its triples
	for _, triple := range bgp.Triples {
		if strings.HasPrefix(triple.Subject, "?") {
			v.definedVars[triple.Subject] = true
		}
		if triple.ObjectIsVar {
			v.definedVars[triple.Object] = true
		}
	}
}

func (v *AlgebraValidator) validateJoin(j *Join) {
	// Validate left side first
	v.validateExpr(j.Left)

	// Validate right side (with left's variables in scope)
	v.validateExpr(j.Right)
}

func (v *AlgebraValidator) validateFilter(f *Filter) {
	// Validate input first (defines variables)
	v.validateExpr(f.Input)

	// Check filter only references defined variables
	refVars := extractReferencedVariables(f.Expr)
	for _, varRef := range refVars {
		if !v.definedVars[varRef] {
			v.errors = append(v.errors,
				fmt.Sprintf("Filter references undefined variable %s (defined: %v)",
					varRef, v.definedVars))
		}
	}
}

func (v *AlgebraValidator) validateProject(p *Project) {
	// Validate input first
	v.validateExpr(p.Input)

	// Check all projected variables are defined
	for _, projVar := range p.Vars {
		if !v.definedVars[projVar] {
			v.errors = append(v.errors,
				fmt.Sprintf("Project references undefined variable %s (defined: %v)",
					projVar, v.definedVars))
		}
	}
}

// Validators for other operators follow similar pattern...
// (LeftJoin, Union, Aggregate, Bind, Distinct, OrderBy, Limit)

func (v *AlgebraValidator) validateLeftJoin(lj *LeftJoin) {
	v.validateExpr(lj.Input)
	// Optional patterns are allowed to reference input variables
	// Filter is optional
}

func (v *AlgebraValidator) validateUnion(u *Union) {
	// Each alternative is validated independently
	for _, alt := range u.Alternatives {
		v.validateExpr(alt)
	}
}

func (v *AlgebraValidator) validateAggregate(a *Aggregate) {
	v.validateExpr(a.Input)
	// Aggregate can reference any variable from input
}

func (v *AlgebraValidator) validateBind(b *Bind) {
	v.validateExpr(b.Input)
	// After BIND, Var is defined
	v.definedVars[b.Var] = true
}

func (v *AlgebraValidator) validateDistinct(d *Distinct) {
	v.validateExpr(d.Input)
}

func (v *AlgebraValidator) validateOrderBy(ob *OrderBy) {
	v.validateExpr(ob.Input)
	// OrderBy expressions can reference input variables
	for _, expr := range ob.Expressions {
		refVars := extractReferencedVariables(expr)
		for _, varRef := range refVars {
			if !v.definedVars[varRef] {
				v.errors = append(v.errors,
					fmt.Sprintf("OrderBy references undefined variable %s", varRef))
			}
		}
	}
}

func (v *AlgebraValidator) validateLimit(l *Limit) {
	v.validateExpr(l.Input)
}

// extractReferencedVariables finds all ?variables in an expression string.
func extractReferencedVariables(expr string) []string {
	re := regexp.MustCompile(`\?(\w+)`)
	matches := re.FindAllStringSubmatch(expr, -1)

	varMap := make(map[string]bool)
	for _, match := range matches {
		if len(match) > 1 {
			varMap["?"+match[1]] = true
		}
	}

	vars := make([]string, 0, len(varMap))
	for v := range varMap {
		vars = append(vars, v)
	}
	return vars
}
```

---

## 2. Visitor Pattern File (`sparql/algebra_visitor.go`)

### File Structure

```go
package sparql

import (
	"bytes"
	"fmt"
	"strings"
)

// ============================================================================
// SECTION 1: AlgebraVisitor Interface
// ============================================================================

// AlgebraVisitor defines operations on algebra expressions.
// Implementations provide specific behavior (printing, validation, optimization, etc.)
type AlgebraVisitor interface {
	VisitBGP(*BGP) interface{}
	VisitJoin(*Join) interface{}
	VisitFilter(*Filter) interface{}
	VisitLeftJoin(*LeftJoin) interface{}
	VisitUnion(*Union) interface{}
	VisitProject(*Project) interface{}
	VisitAggregate(*Aggregate) interface{}
	VisitBind(*Bind) interface{}
	VisitDistinct(*Distinct) interface{}
	VisitOrderBy(*OrderBy) interface{}
	VisitLimit(*Limit) interface{}
}

// ============================================================================
// SECTION 2: AlgebraPrinter Implementation
// ============================================================================

// AlgebraPrinter prints algebra trees in human-readable format with indentation.
type AlgebraPrinter struct {
	indent int
	output bytes.Buffer
}

// NewAlgebraPrinter creates a new printer.
func NewAlgebraPrinter() *AlgebraPrinter {
	return &AlgebraPrinter{indent: 0}
}

// Print returns formatted algebra tree as string.
func (p *AlgebraPrinter) Print(expr AlgebraExpr) string {
	p.output.Reset()
	p.indent = 0
	expr.Accept(p)
	return p.output.String()
}

// writeLine writes a line with proper indentation.
func (p *AlgebraPrinter) writeLine(text string) {
	for i := 0; i < p.indent; i++ {
		p.output.WriteString("  ")
	}
	p.output.WriteString(text)
	p.output.WriteString("\n")
}

// Visitor methods for each operator type:

func (p *AlgebraPrinter) VisitBGP(bgp *BGP) interface{} {
	p.writeLine("BGP")
	p.indent++
	for _, triple := range bgp.Triples {
		obj := triple.Object
		if !triple.ObjectIsVar {
			obj = fmt.Sprintf("<%s>", obj)
		}
		p.writeLine(fmt.Sprintf("%s %s %s", triple.Subject, triple.Predicate, obj))
	}
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitJoin(j *Join) interface{} {
	p.writeLine("Join")
	p.indent++
	p.writeLine("Left:")
	p.indent++
	j.Left.Accept(p)
	p.indent--
	p.writeLine("Right:")
	p.indent++
	j.Right.Accept(p)
	p.indent--
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitFilter(f *Filter) interface{} {
	p.writeLine(fmt.Sprintf("Filter: %s", f.Expr))
	p.indent++
	f.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitLeftJoin(lj *LeftJoin) interface{} {
	p.writeLine("LeftJoin (OPTIONAL)")
	p.indent++
	p.writeLine("Input:")
	p.indent++
	lj.Input.Accept(p)
	p.indent--
	p.writeLine("Patterns:")
	p.indent++
	for _, triple := range lj.Patterns {
		obj := triple.Object
		if !triple.ObjectIsVar {
			obj = fmt.Sprintf("<%s>", obj)
		}
		p.writeLine(fmt.Sprintf("%s %s %s", triple.Subject, triple.Predicate, obj))
	}
	if lj.Filter != "" {
		p.writeLine(fmt.Sprintf("Filter: %s", lj.Filter))
	}
	p.indent--
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitUnion(u *Union) interface{} {
	p.writeLine("Union")
	p.indent++
	for i, alt := range u.Alternatives {
		p.writeLine(fmt.Sprintf("Alternative %d:", i+1))
		p.indent++
		alt.Accept(p)
		p.indent--
	}
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitProject(proj *Project) interface{} {
	p.writeLine(fmt.Sprintf("Project: %s", strings.Join(proj.Vars, ", ")))
	p.indent++
	proj.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitAggregate(agg *Aggregate) interface{} {
	groupStr := ""
	if len(agg.Group) > 0 {
		groupStr = fmt.Sprintf(" GROUP BY %s", strings.Join(agg.Group, ", "))
	}
	p.writeLine(fmt.Sprintf("Aggregate: %s(%s)%s", agg.Op, agg.Expr, groupStr))
	p.indent++
	agg.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitBind(b *Bind) interface{} {
	p.writeLine(fmt.Sprintf("Bind: %s AS %s", b.Var, b.Expr))
	p.indent++
	b.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitDistinct(d *Distinct) interface{} {
	p.writeLine("Distinct")
	p.indent++
	d.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitOrderBy(ob *OrderBy) interface{} {
	orderStr := ""
	for i, expr := range ob.Expressions {
		dir := "ASC"
		if !ob.Ascending[i] {
			dir = "DESC"
		}
		if i > 0 {
			orderStr += ", "
		}
		orderStr += fmt.Sprintf("%s %s", expr, dir)
	}
	p.writeLine(fmt.Sprintf("OrderBy: %s", orderStr))
	p.indent++
	ob.Input.Accept(p)
	p.indent--
	return nil
}

func (p *AlgebraPrinter) VisitLimit(l *Limit) interface{} {
	limitStr := fmt.Sprintf("Limit: count=%d", l.Count)
	if l.Offset > 0 {
		limitStr += fmt.Sprintf(", offset=%d", l.Offset)
	}
	p.writeLine(limitStr)
	p.indent++
	l.Input.Accept(p)
	p.indent--
	return nil
}
```

---

## 3. Execution Context File (`sparql/context.go`)

### File Structure

```go
package sparql

import (
	"context"
	"github.com/dgraph-io/dgraph/v25/graphql/schema"
)

// SPARQLExecutionContext carries context through the SPARQL execution pipeline.
// It aggregates query context, schema, auth information, and options.
type SPARQLExecutionContext struct {
	// Ctx is the request context (for cancellation and timeouts)
	Ctx context.Context

	// Schema is the GraphQL schema (used in Phase 3 for type validation)
	Schema *schema.Schema

	// Authentication context (used in Phase 4 for auth rule application)
	UserID   string   // Current authenticated user
	GroupIDs []string // User's group memberships

	// Multi-tenancy support
	Namespace uint64

	// Translation options
	Options TranslateOptions

	// Query metadata
	QueryString string
	QueryHash   string
}

// NewSPARQLExecutionContext creates a new execution context with defaults.
func NewSPARQLExecutionContext(ctx context.Context) *SPARQLExecutionContext {
	if ctx == nil {
		ctx = context.Background()
	}

	return &SPARQLExecutionContext{
		Ctx:      ctx,
		GroupIDs: make([]string, 0),
		Options:  TranslateOptions{},
	}
}

// WithSchema sets the schema for this context (fluent API).
func (c *SPARQLExecutionContext) WithSchema(s *schema.Schema) *SPARQLExecutionContext {
	c.Schema = s
	return c
}

// WithAuth sets authentication information (fluent API).
func (c *SPARQLExecutionContext) WithAuth(userID string, groupIDs []string) *SPARQLExecutionContext {
	c.UserID = userID
	c.GroupIDs = groupIDs
	return c
}

// WithNamespace sets the namespace/tenant (fluent API).
func (c *SPARQLExecutionContext) WithNamespace(ns uint64) *SPARQLExecutionContext {
	c.Namespace = ns
	return c
}

// WithOptions sets translation options (fluent API).
func (c *SPARQLExecutionContext) WithOptions(opts TranslateOptions) *SPARQLExecutionContext {
	c.Options = opts
	return c
}

// WithQueryString sets the original query string (fluent API).
func (c *SPARQLExecutionContext) WithQueryString(qs string) *SPARQLExecutionContext {
	c.QueryString = qs
	return c
}

// IsAuthenticated returns true if user is authenticated.
func (c *SPARQLExecutionContext) IsAuthenticated() bool {
	return c.UserID != ""
}

// IsMember checks if user is in a specific group.
func (c *SPARQLExecutionContext) IsMember(group string) bool {
	for _, g := range c.GroupIDs {
		if g == group {
			return true
		}
	}
	return false
}
```

---

## 4. Test Structure Template (`sparql/algebra_test.go`)

### File Organization Pattern

```go
package sparql

import (
	"testing"
	"github.com/stretchr/testify/require"
)

// ============================================================================
// SECTION 1: Algebra Type Construction Tests
// ============================================================================

func TestBGPConstruction(t *testing.T) {
	bgp := &BGP{
		Triples: []*Triple{
			{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
		},
	}

	require.NotNil(t, bgp)
	require.Equal(t, 1, len(bgp.Triples))
}

func TestBGPString(t *testing.T) {
	bgp := &BGP{
		Triples: []*Triple{
			{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
			{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	result := bgp.String()
	require.Contains(t, result, "BGP")
	require.Contains(t, result, "?s")
	require.Contains(t, result, "foaf:name")
	require.Contains(t, result, "?name")
}

// Similar test functions for Join, Filter, Project, etc.
// ~15 tests total in this section

// ============================================================================
// SECTION 2: Visitor Pattern Tests
// ============================================================================

func TestAlgebraPrinterSimpleBGP(t *testing.T) {
	bgp := &BGP{
		Triples: []*Triple{
			{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
		},
	}

	printer := NewAlgebraPrinter()
	result := printer.Print(bgp)

	require.Contains(t, result, "BGP")
	require.Contains(t, result, "?s")
	require.Contains(t, result, "rdf:type")
}

func TestAlgebraPrinterNestedStructure(t *testing.T) {
	algebra := &Project{
		Vars: []string{"?s", "?name"},
		Input: &Filter{
			Expr: "?name = 'Alice'",
			Input: &Join{
				Left: &BGP{
					Triples: []*Triple{
						{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
					},
				},
				Right: &BGP{
					Triples: []*Triple{
						{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
					},
				},
			},
		},
	}

	printer := NewAlgebraPrinter()
	result := printer.Print(algebra)

	require.Contains(t, result, "Project")
	require.Contains(t, result, "Filter")
	require.Contains(t, result, "Join")
	require.Contains(t, result, "BGP")
}

// ~10 tests total in this section

// ============================================================================
// SECTION 3: AST to Algebra Conversion Tests
// ============================================================================

func TestASTToAlgebraSimpleSelect(t *testing.T) {
	ast := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?person", "?name"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?person", Predicate: "rdf:type", Object: "Person"},
					{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
				},
			},
		},
	}

	algebra, err := ASTToAlgebra(ast)
	require.NoError(t, err)
	require.NotNil(t, algebra)

	// Verify structure
	project, ok := algebra.(*Project)
	require.True(t, ok)
	require.Equal(t, []string{"?person", "?name"}, project.Vars)
}

// Similar test functions for FILTER, OPTIONAL, UNION, aggregates, etc.
// ~50+ tests total in this section

// ============================================================================
// SECTION 4: Algebra Validator Tests
// ============================================================================

func TestValidatorValidExpression(t *testing.T) {
	algebra := &Filter{
		Expr: "?name = 'Alice'",
		Input: &BGP{
			Triples: []*Triple{
				{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			},
		},
	}

	validator := &AlgebraValidator{}
	err := validator.Validate(algebra)
	require.NoError(t, err)
}

func TestValidatorDetectsUndefinedVariable(t *testing.T) {
	algebra := &Filter{
		Expr: "?age > 18",
		Input: &BGP{
			Triples: []*Triple{
				{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			},
		},
	}

	validator := &AlgebraValidator{}
	err := validator.Validate(algebra)
	require.Error(t, err)
	require.Contains(t, err.Error(), "?age")
	require.Contains(t, err.Error(), "undefined")
}

// Similar test functions for other validation scenarios
// ~15+ tests total in this section

// ============================================================================
// SECTION 5: Execution Context Tests
// ============================================================================

func TestContextCreation(t *testing.T) {
	ctx := NewSPARQLExecutionContext(nil)
	require.NotNil(t, ctx)
	require.NotNil(t, ctx.Ctx)
	require.Equal(t, "", ctx.UserID)
}

func TestContextBuilderPattern(t *testing.T) {
	ctx := NewSPARQLExecutionContext(nil).
		WithAuth("user1", []string{"admin"}).
		WithNamespace(1)

	require.Equal(t, "user1", ctx.UserID)
	require.True(t, ctx.IsAuthenticated())
	require.True(t, ctx.IsMember("admin"))
}

// ~5 tests total in this section

// ============================================================================
// SECTION 6: Helper Test Functions
// ============================================================================

// Helper to create common test algebra structures
func createSimpleBGP() AlgebraExpr {
	return &BGP{
		Triples: []*Triple{
			{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
		},
	}
}

func createFilteredBGP(filterExpr string) AlgebraExpr {
	return &Filter{
		Expr:  filterExpr,
		Input: createSimpleBGP(),
	}
}

// Add more helpers as needed
```

---

## 5. File Size Estimates

| File               | LOC           | Components                                                |
| ------------------ | ------------- | --------------------------------------------------------- |
| algebra.go         | 500-600       | Interface, 11 operators, ASTToAlgebra, helpers, Validator |
| algebra_visitor.go | 100-150       | AlgebraVisitor interface, AlgebraPrinter implementation   |
| context.go         | 100           | SPARQLExecutionContext, factory, fluent methods           |
| algebra_test.go    | 400-500       | 100+ test cases across 5 categories                       |
| **TOTAL**          | **1100-1350** | **Complete Phase 1**                                      |

---

## 6. Implementation Sequence

**Recommended order for file creation**:

1. **algebra.go** (FIRST)
   - Define AlgebraExpr interface
   - Implement all 11 operator types
   - Implement Accept() methods
   - Implement String() methods
   - Add ASTToAlgebra() skeleton with error handling
   - Add helper functions
   - Add AlgebraValidator

2. **algebra_visitor.go** (SECOND)
   - Define AlgebraVisitor interface
   - Implement AlgebraPrinter
   - Test with algebra_test.go during development

3. **context.go** (THIRD)
   - Quick implementation
   - No dependencies on other files

4. **algebra_test.go** (THROUGHOUT)
   - Start writing tests immediately (TDD approach)
   - Tests guide implementation
   - Run frequently with `go test ./sparql -v`

---

## 7. Common Patterns & Conventions

### Pattern 1: Accept() Implementation

All algebra operators use the same Accept pattern:

```go
func (x *OperatorType) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitOperatorType(x)
}
```

### Pattern 2: String() Implementation

Format: `OperatorName(details, child1_string, child2_string)`

```go
func (x *OperatorType) String() string {
	return fmt.Sprintf("OperatorName(%s)", details...)
}
```

### Pattern 3: Recursive Conversion

Helper functions follow recursive pattern:

```go
func convertXxx(input) (AlgebraExpr, error) {
	if len(input) == 0 {
		return nil, error
	}
	if len(input) == 1 {
		return convertSingle(input[0])
	}
	// Recursively split and combine
}
```

### Pattern 4: Variable Scope Tracking

Validators track scope as they traverse:

```go
func (v *Validator) validateXxx(x *Operator) {
	v.validateExpr(x.Input)  // Define variables
	v.checkReferences(x)     // Verify references
}
```

---

**End of Code Structure Guide**
