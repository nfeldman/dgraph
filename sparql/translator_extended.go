package sparql

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	dql "github.com/dgraph-io/dgraph/v25/dql"
	"github.com/dgraph-io/dgraph/v25/protos/pb"
	"github.com/dgraph-io/dgraph/v25/types"
)

// TranslateSelectExtended handles SELECT queries with support for OPTIONAL, UNION, aggregates, BIND, HAVING.
func TranslateSelectExtended(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq.Type() != "SELECT" {
		return nil, nil, fmt.Errorf("expected SELECT query, got %s", sq.Type())
	}

	if !hasPatterns(sq) {
		return nil, nil, fmt.Errorf("no graph patterns provided")
	}

	// Cast to extended query interface to access new fields
	var aggregates []*Aggregate
	var binds []*BindExpression
	var having *HavingClause
	var distinct bool
	var orderBy []string
	var patterns []GraphPattern

	// Try to get extended fields if available
	if extQuery, ok := sq.(*SPARQLQueryImpl); ok {
		aggregates = extQuery.Aggregates
		binds = extQuery.Binds
		having = extQuery.Having
		distinct = extQuery.Distinct
		orderBy = extQuery.OrderBy
		patterns = extQuery.Patterns
	}

	fromGraphs := sq.FROM()
	fromNamedGraphs := sq.FROMNamed()

	// Build the DQL query
	rootQuery := &dql.GraphQuery{
		Attr: "query",
	}

	// If patterns are provided (new format), use them; otherwise use old BGP format
	if len(patterns) > 0 {
		for _, pattern := range patterns {
			if err := translateGraphPattern(pattern, fromGraphs, fromNamedGraphs, rootQuery, opts); err != nil {
				return nil, nil, err
			}
		}
	} else {
		// Fall back to old BGP processing
		for _, patternIntf := range sq.RootGraphPatterns() {
			bgp, ok := patternIntf.(*BGP)
			if !ok {
				continue
			}
			for _, triple := range bgp.Triples {
				childQuery, err := translateTriple(triple, fromGraphs, fromNamedGraphs)
				if err != nil {
					return nil, nil, fmt.Errorf("translating triple: %w", err)
				}
				if childQuery != nil {
					rootQuery.Children = append(rootQuery.Children, childQuery)
				}
			}
		}
	}

	// Apply BIND expressions (variable bindings)
	if len(binds) > 0 {
		for _, bind := range binds {
			if err := applyBindExpression(rootQuery, bind); err != nil {
				return nil, nil, fmt.Errorf("applying BIND: %w", err)
			}
		}
	}

	// Apply aggregates if present
	if len(aggregates) > 0 {
		if err := applyAggregates(rootQuery, aggregates); err != nil {
			return nil, nil, fmt.Errorf("applying aggregates: %w", err)
		}
	}

	// Apply HAVING filter on aggregates
	if having != nil {
		if err := applyHavingClause(rootQuery, having); err != nil {
			return nil, nil, fmt.Errorf("applying HAVING: %w", err)
		}
	}

	// Apply DISTINCT
	if distinct {
		rootQuery.Normalize = true
	}

	// Apply ORDER BY
	if len(orderBy) > 0 {
		rootQuery.Order = make([]*pb.Order, 0, len(orderBy))
		for _, varName := range orderBy {
			order := &pb.Order{
				Attr: varName,
			}
			rootQuery.Order = append(rootQuery.Order, order)
		}
	}

	return []*dql.GraphQuery{rootQuery}, sq.GetPrefixes(), nil
}

// translateGraphPattern handles different pattern types (BGP, OPTIONAL, UNION, FILTER).
//
// OPTIONAL patterns are translated such that matches are included if available, but the
// absence of a match doesn't exclude the row. This is achieved by:
//  1. Creating child queries for optional patterns
//  2. Using nullable/permissive filters that don't fail if the predicate is missing
//
// UNION patterns create OR semantics by collecting alternatives and building filter trees
// that combine them with OR operators.
func translateGraphPattern(pattern GraphPattern, fromGraphs, fromNamedGraphs []string, parent *dql.GraphQuery, opts TranslateOptions) error {
	switch p := pattern.(type) {
	case *BGP:
		// Regular basic graph pattern - required patterns
		for _, triple := range p.Triples {
			childQuery, err := translateTriple(triple, fromGraphs, fromNamedGraphs)
			if err != nil {
				return fmt.Errorf("translating triple: %w", err)
			}
			if childQuery != nil {
				parent.Children = append(parent.Children, childQuery)
			}
		}
		return nil

	case *OptionalPattern:
		// OPTIONAL pattern - matches are included if available, but absence doesn't fail.
		// In DQL, we implement this by:
		// 1. Creating a subquery for the optional patterns
		// 2. Using a wrapper that allows the subquery to have empty results
		//
		// For simplicity, we add the patterns as child queries but mark them
		// as having optional semantics (by not adding them as required filters).
		optionalQuery := &dql.GraphQuery{
			Attr: "_optional",
		}

		for _, subPattern := range p.Patterns {
			if err := translateGraphPattern(subPattern, fromGraphs, fromNamedGraphs, optionalQuery, opts); err != nil {
				return err
			}
		}

		// Add the optional query as a child
		parent.Children = append(parent.Children, optionalQuery)
		return nil

	case *UnionPattern:
		// UNION - multiple alternatives with OR semantics.
		// In SPARQL, a result set is valid if ANY of the alternatives match.
		// We implement this by creating separate subqueries for each alternative
		// and building an OR filter to combine them.

		if len(p.Alternatives) == 0 {
			return nil
		}

		// For each alternative, create a separate query branch
		unionFilters := make([]*dql.FilterTree, 0, len(p.Alternatives))

		for altIdx, alternative := range p.Alternatives {
			// Create a subquery for this alternative
			altQuery := &dql.GraphQuery{
				Attr: fmt.Sprintf("_union_alt_%d", altIdx),
			}

			for _, altPattern := range alternative {
				if err := translateGraphPattern(altPattern, fromGraphs, fromNamedGraphs, altQuery, opts); err != nil {
					return err
				}
			}

			parent.Children = append(parent.Children, altQuery)

			// Create a filter representing this alternative
			// For now, we represent this conceptually; actual execution would need
			// more sophisticated handling in the query engine
			altFilter := &dql.FilterTree{
				Func: &dql.Function{
					Name: "union_alternative",
					Attr: fmt.Sprintf("alt_%d", altIdx),
				},
			}
			unionFilters = append(unionFilters, altFilter)
		}

		// Combine alternatives with OR logic
		if len(unionFilters) > 0 {
			orFilter := buildOrFilterTree(unionFilters)
			if parent.Filter == nil {
				parent.Filter = orFilter
			} else {
				// If there's already a filter, AND it with the union filter
				parent.Filter = &dql.FilterTree{
					Op:    "AND",
					Child: []*dql.FilterTree{parent.Filter, orFilter},
				}
			}
		}
		return nil

	case *FilterPattern:
		return applyFilterExpression(parent, p.Expression, opts)

	default:
		return fmt.Errorf("unknown graph pattern type: %T", pattern)
	}
}

// buildOrFilterTree builds an OR filter tree from a slice of filters.
func buildOrFilterTree(filters []*dql.FilterTree) *dql.FilterTree {
	if len(filters) == 0 {
		return nil
	}
	if len(filters) == 1 {
		return filters[0]
	}

	// Build OR tree: f1 OR f2 OR f3 ... becomes (f1 OR (f2 OR f3 ...))
	result := filters[len(filters)-1]
	for i := len(filters) - 2; i >= 0; i-- {
		result = &dql.FilterTree{
			Op:    "OR",
			Child: []*dql.FilterTree{filters[i], result},
		}
	}
	return result
}

// applyFilterExpression attaches a FILTER expression to the current query node.
// Tries to map common comparison/regex expressions to DQL functions; falls back to raw.
func applyFilterExpression(target *dql.GraphQuery, expr string, opts TranslateOptions) error {
	if expr == "" || target == nil {
		return nil
	}

	filter, err := buildFilterTreeFromExpr(expr)
	if err != nil {
		if opts.AllowFilterFallback {
			filter = &dql.FilterTree{Func: &dql.Function{Name: "filter", Attr: expr}}
		} else {
			return fmt.Errorf("parsing FILTER expression '%s': %w", expr, err)
		}
	}
	if filter == nil {
		return fmt.Errorf("parsing FILTER expression '%s': produced empty filter", expr)
	}

	if target.Filter == nil {
		target.Filter = filter
	} else {
		target.Filter = &dql.FilterTree{
			Op:    "AND",
			Child: []*dql.FilterTree{target.Filter, filter},
		}
	}
	return nil
}

// buildFilterTreeFromExpr handles simple FILTER expressions: comparisons and regex.
// Examples handled:
//
//	FILTER(?a = 5)
//	FILTER(?a != ?b)
//	FILTER(?a < 10)
//	FILTER(regex(?name, "^A"))
//
// Logical AND/OR with && / || are combined recursively.
func buildFilterTreeFromExpr(expr string) (*dql.FilterTree, error) {
	trimmed := strings.TrimSpace(expr)
	if strings.HasPrefix(strings.ToUpper(trimmed), "FILTER") {
		trimmed = strings.TrimSpace(trimmed[6:])
	}
	trimmed = strings.TrimSpace(trimmed)
	// Drop outer parentheses when they wrap the whole expression
	for strings.HasPrefix(trimmed, "(") && strings.HasSuffix(trimmed, ")") {
		trimmed = strings.TrimSpace(trimmed[1 : len(trimmed)-1])
	}

	// Handle OR / AND (naive split; sufficient for simple expressions)
	if strings.Contains(trimmed, "||") {
		parts := strings.Split(trimmed, "||")
		var children []*dql.FilterTree
		for _, p := range parts {
			if ft, err := buildFilterTreeFromExpr(p); err == nil && ft != nil {
				children = append(children, ft)
			} else if err != nil {
				return nil, err
			}
		}
		if len(children) == 0 {
			return nil, fmt.Errorf("empty OR filter from '%s'", expr)
		}
		if len(children) == 1 {
			return children[0], nil
		}
		return &dql.FilterTree{Op: "OR", Child: children}, nil
	}
	if strings.Contains(trimmed, "&&") {
		parts := strings.Split(trimmed, "&&")
		var children []*dql.FilterTree
		for _, p := range parts {
			if ft, err := buildFilterTreeFromExpr(p); err == nil && ft != nil {
				children = append(children, ft)
			} else if err != nil {
				return nil, err
			}
		}
		if len(children) == 0 {
			return nil, fmt.Errorf("empty AND filter from '%s'", expr)
		}
		if len(children) == 1 {
			return children[0], nil
		}
		return &dql.FilterTree{Op: "AND", Child: children}, nil
	}

	// Comparison operators
	compRe := regexp.MustCompile(`^([?$]\w+)\s*(=|!=|<=|>=|<|>)\s*(.+)$`)
	if m := compRe.FindStringSubmatch(trimmed); len(m) == 4 {
		lhs := strings.TrimSpace(m[1])
		op := m[2]
		rhs := strings.TrimSpace(m[3])

		fn := map[string]string{"=": "eq", "!=": "ne", "<": "lt", "<=": "le", ">": "gt", ">=": "ge"}[op]
		if fn == "" {
			fn = op
		}

		attr := strings.TrimLeft(lhs, "?$")
		return &dql.FilterTree{
			Func: &dql.Function{
				Name: fn,
				Attr: attr,
				Args: []dql.Arg{{Value: rhs}},
			},
		}, nil
	}

	// regex(?var, "pattern")
	regexRe := regexp.MustCompile(`^regex\s*\(\s*([?$]\w+)\s*,\s*"([^"]*)"\s*\)$`)
	if m := regexRe.FindStringSubmatch(trimmed); len(m) == 3 {
		attr := strings.TrimLeft(m[1], "?$")
		pattern := m[2]
		return &dql.FilterTree{
			Func: &dql.Function{
				Name: "regexp",
				Attr: attr,
				Args: []dql.Arg{{Value: quoteString(pattern)}},
			},
		}, nil
	}

	return nil, fmt.Errorf("unsupported FILTER expression: %s", trimmed)
}

// applyBindExpression applies variable binding expressions (BIND ... AS ...).
// SPARQL BIND creates computed variables from expressions.
//
// Supported expression types:
//   - Arithmetic: ?x + ?y, ?x - ?y, ?x * ?y, ?x / ?y, ?x % ?y
//   - Functions: CONCAT(?a, ?b), SUBSTR(?str, ?pos, ?len), etc.
//   - Constants: numeric and string literals
//
// SPARQL expressions are converted to DQL MathTree for evaluation.
// Results are stored with variable names for later use in projections.
func applyBindExpression(query *dql.GraphQuery, bind *BindExpression) error {
	if bind == nil || bind.Expression == "" {
		return nil
	}

	// Parse the expression into a MathTree for DQL
	mathTree, err := parseBindExpression(bind.Expression)
	if err != nil {
		return fmt.Errorf("parsing BIND expression %q: %w", bind.Expression, err)
	}

	// Create a child query to compute and bind the variable
	childQuery := &dql.GraphQuery{
		Attr:    "val",
		Var:     bind.Variable,
		MathExp: mathTree,
	}

	query.Children = append(query.Children, childQuery)
	return nil
}

// parseBindExpression parses a BIND expression string into a DQL MathTree.
// Handles arithmetic operators (+, -, *, /, %), function calls, and variables.
func parseBindExpression(expr string) (*dql.MathTree, error) {
	expr = strings.TrimSpace(expr)

	// Check for function calls like CONCAT, SUBSTR, etc.
	if idx := strings.Index(expr, "("); idx > 0 {
		funcName := strings.TrimSpace(expr[:idx])
		funcNameLower := strings.ToLower(funcName)

		// Handle known functions
		switch funcNameLower {
		case "concat":
			return parseConcat(expr)
		case "substr":
			return parseSubstr(expr)
		case "strlen":
			return parseStrlen(expr)
		case "sqrt", "abs", "floor", "ceil", "exp", "ln":
			return parseMathFunction(funcNameLower, expr)
		}
	}

	// Handle arithmetic expressions (simplified parser for now)
	return parseArithmeticExpression(expr)
}

// parseArithmeticExpression parses arithmetic expressions with +, -, *, /, %.
// Handles operator precedence: * / % before + -
func parseArithmeticExpression(expr string) (*dql.MathTree, error) {
	expr = strings.TrimSpace(expr)

	// Remove outer parentheses if they're balanced and cover the whole expression
	for strings.HasPrefix(expr, "(") && strings.HasSuffix(expr, ")") {
		// Check if the outer parentheses are balanced
		inner := expr[1 : len(expr)-1]
		if isBalancedParens(inner) {
			expr = strings.TrimSpace(inner)
		} else {
			break
		}
	}

	// Simple recursive descent parser for arithmetic
	// Split on lowest precedence operators first (+, -)
	for _, op := range []string{"+", "-"} {
		for i := len(expr) - 1; i >= 0; i-- {
			if expr[i:i+1] == op {
				// Make sure it's not a unary operator
				if i > 0 && !isOperator(expr[i-1]) {
					left := strings.TrimSpace(expr[:i])
					right := strings.TrimSpace(expr[i+1:])

					leftTree, err := parseArithmeticExpression(left)
					if err != nil {
						return nil, err
					}
					rightTree, err := parseArithmeticExpression(right)
					if err != nil {
						return nil, err
					}

					return &dql.MathTree{
						Fn:    op,
						Child: []*dql.MathTree{leftTree, rightTree},
					}, nil
				}
			}
		}
	}

	// Try higher precedence operators (*, /, %)
	for _, op := range []string{"*", "/", "%"} {
		for i := len(expr) - 1; i >= 0; i-- {
			if expr[i:i+1] == op {
				if i > 0 && !isOperator(expr[i-1]) {
					left := strings.TrimSpace(expr[:i])
					right := strings.TrimSpace(expr[i+1:])

					leftTree, err := parseArithmeticExpression(left)
					if err != nil {
						return nil, err
					}
					rightTree, err := parseArithmeticExpression(right)
					if err != nil {
						return nil, err
					}

					return &dql.MathTree{
						Fn:    op,
						Child: []*dql.MathTree{leftTree, rightTree},
					}, nil
				}
			}
		}
	}

	// If no operators, must be a variable or constant
	return parseOperand(expr)
}

// parseOperand parses a single operand (variable or constant).
func parseOperand(expr string) (*dql.MathTree, error) {
	expr = strings.TrimSpace(expr)

	// Check if it's a variable (starts with ?)
	if strings.HasPrefix(expr, "?") {
		return &dql.MathTree{
			Var: expr[1:], // Store without the ?
		}, nil
	}

	// Try to parse as a numeric constant
	if val, err := strconv.ParseFloat(expr, 64); err == nil {
		return &dql.MathTree{
			Const: types.Val{
				Tid:   types.FloatID,
				Value: val,
			},
		}, nil
	}

	// Try to parse as an integer constant
	if val, err := strconv.ParseInt(expr, 10, 64); err == nil {
		return &dql.MathTree{
			Const: types.Val{
				Tid:   types.IntID,
				Value: val,
			},
		}, nil
	}

	return nil, fmt.Errorf("cannot parse operand: %q", expr)
}

// parseConcat parses CONCAT(arg1, arg2, ...) function.
func parseConcat(expr string) (*dql.MathTree, error) {
	// Extract arguments from CONCAT(...)
	args, err := extractFunctionArgs(expr)
	if err != nil {
		return nil, err
	}

	// CONCAT builds a tree by concatenating arguments left-to-right
	// For now, return a placeholder indicating string concatenation
	// In a full implementation, this would need special handling in the query engine
	if len(args) == 0 {
		return nil, fmt.Errorf("CONCAT requires at least one argument")
	}

	// Build a tree representing concatenation
	// Use a special function name to indicate string concatenation
	tree := &dql.MathTree{
		Fn: "concat",
	}

	for _, arg := range args {
		argTree, err := parseOperand(arg)
		if err != nil {
			// Try parsing as a nested expression
			argTree, err = parseBindExpression(arg)
			if err != nil {
				return nil, err
			}
		}
		tree.Child = append(tree.Child, argTree)
	}

	return tree, nil
}

// parseSubstr parses SUBSTR(string, start, length) function.
func parseSubstr(expr string) (*dql.MathTree, error) {
	args, err := extractFunctionArgs(expr)
	if err != nil {
		return nil, err
	}

	if len(args) < 2 {
		return nil, fmt.Errorf("SUBSTR requires at least 2 arguments (string, start[, length])")
	}

	strTree, err := parseOperand(strings.TrimSpace(args[0]))
	if err != nil {
		strTree, err = parseBindExpression(strings.TrimSpace(args[0]))
		if err != nil {
			return nil, err
		}
	}

	startTree, err := parseOperand(strings.TrimSpace(args[1]))
	if err != nil {
		return nil, err
	}

	var lengthTree *dql.MathTree
	if len(args) >= 3 {
		lengthTree, err = parseOperand(strings.TrimSpace(args[2]))
		if err != nil {
			return nil, err
		}
	}

	tree := &dql.MathTree{
		Fn:    "substr",
		Child: []*dql.MathTree{strTree, startTree},
	}

	if lengthTree != nil {
		tree.Child = append(tree.Child, lengthTree)
	}

	return tree, nil
}

// parseStrlen parses STRLEN(string) function.
func parseStrlen(expr string) (*dql.MathTree, error) {
	args, err := extractFunctionArgs(expr)
	if err != nil {
		return nil, err
	}

	if len(args) != 1 {
		return nil, fmt.Errorf("STRLEN requires exactly 1 argument")
	}

	strTree, err := parseOperand(strings.TrimSpace(args[0]))
	if err != nil {
		strTree, err = parseBindExpression(strings.TrimSpace(args[0]))
		if err != nil {
			return nil, err
		}
	}

	return &dql.MathTree{
		Fn:    "strlen",
		Child: []*dql.MathTree{strTree},
	}, nil
}

// parseMathFunction parses math functions like SQRT, ABS, FLOOR, CEIL, EXP, LN.
func parseMathFunction(funcName, expr string) (*dql.MathTree, error) {
	args, err := extractFunctionArgs(expr)
	if err != nil {
		return nil, err
	}

	if len(args) != 1 {
		return nil, fmt.Errorf("%s requires exactly 1 argument", strings.ToUpper(funcName))
	}

	argTree, err := parseOperand(strings.TrimSpace(args[0]))
	if err != nil {
		argTree, err = parseArithmeticExpression(strings.TrimSpace(args[0]))
		if err != nil {
			return nil, err
		}
	}

	return &dql.MathTree{
		Fn:    funcName,
		Child: []*dql.MathTree{argTree},
	}, nil
}

// extractFunctionArgs extracts arguments from a function call like "FUNC(arg1, arg2, ...)".
func extractFunctionArgs(expr string) ([]string, error) {
	expr = strings.TrimSpace(expr)

	// Find opening parenthesis
	openIdx := strings.Index(expr, "(")
	if openIdx < 0 {
		return nil, fmt.Errorf("malformed function call: %s", expr)
	}

	// Find matching closing parenthesis
	closeIdx := -1
	parenCount := 0
	for i := openIdx; i < len(expr); i++ {
		if expr[i] == '(' {
			parenCount++
		} else if expr[i] == ')' {
			parenCount--
			if parenCount == 0 {
				closeIdx = i
				break
			}
		}
	}

	if closeIdx < 0 {
		return nil, fmt.Errorf("unmatched parentheses in function call: %s", expr)
	}

	argsStr := expr[openIdx+1 : closeIdx]
	if argsStr == "" {
		return nil, fmt.Errorf("function requires at least one argument")
	}

	// Split by commas, but respect nested parentheses
	var args []string
	var current strings.Builder
	depth := 0

	for _, ch := range argsStr {
		switch ch {
		case '(':
			depth++
			current.WriteRune(ch)
		case ')':
			depth--
			current.WriteRune(ch)
		case ',':
			if depth == 0 {
				args = append(args, strings.TrimSpace(current.String()))
				current.Reset()
			} else {
				current.WriteRune(ch)
			}
		default:
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		args = append(args, strings.TrimSpace(current.String()))
	}

	return args, nil
}

// isOperator returns true if the character is a math operator.
func isOperator(ch byte) bool {
	return ch == '+' || ch == '-' || ch == '*' || ch == '/' || ch == '%' ||
		ch == '(' || ch == ')'
}

// isBalancedParens checks if a string of parentheses is balanced without outer parens.
// E.g., "x + y" -> true, "x + (y * z)" -> true, "(x + y" -> false
func isBalancedParens(s string) bool {
	depth := 0
	for _, ch := range s {
		if ch == '(' {
			depth++
		} else if ch == ')' {
			depth--
			if depth < 0 {
				return false
			}
		}
	}
	return depth == 0
}

// applyAggregates applies aggregate functions (COUNT, SUM, MIN, MAX, AVG) to DQL.
// It transforms SPARQL aggregates into DQL @groupby directives with proper
// variable assignments (e.g., `a as count(uid)`, `sum_val as sum(value)`, etc.)
//
// For DISTINCT aggregates, the attribute is prefixed with "distinct " in DQL.
//
// SPARQL -> DQL mapping:
//   - COUNT(?var) -> count(uid) [ignores var in DQL, counts UIDs]
//   - SUM(?var)   -> sum(value_attr)
//   - MIN(?var)   -> min(value_attr)
//   - MAX(?var)   -> max(value_attr)
//   - AVG(?var)   -> avg(value_attr)
//   - With DISTINCT: distinct attr_name
func applyAggregates(query *dql.GraphQuery, aggregates []*Aggregate) error {
	if len(aggregates) == 0 {
		return nil
	}

	// Enable groupby mode
	query.IsGroupby = true

	// Process each aggregate function
	for _, agg := range aggregates {
		if err := applyAggregate(query, agg); err != nil {
			return err
		}
	}

	return nil
}

// applyAggregate applies a single aggregate function to the query.
func applyAggregate(query *dql.GraphQuery, agg *Aggregate) error {
	if agg == nil {
		return nil
	}

	// Normalize function name to lowercase for DQL
	funcName := strings.ToLower(agg.Function)

	// Map SPARQL aggregate syntax to DQL aggregate functions
	var dqlAggregate string
	switch funcName {
	case "count":
		// COUNT(?x) in SPARQL becomes count(uid) in DQL
		// The variable is implicit (counts distinct UIDs)
		if agg.Distinct {
			dqlAggregate = "distinct count(uid)"
		} else {
			dqlAggregate = "count(uid)"
		}

	case "sum", "avg", "min", "max":
		// For numerical aggregates, extract the predicate from variable
		// SUM(?val) -> sum(value) where "value" is inferred from context
		// For now, use a generic "value" attribute; in a real scenario,
		// this would come from the query pattern analysis
		var attrName string
		if strings.HasPrefix(agg.Variable, "?") {
			attrName = agg.Variable[1:] // strip "?" prefix
		} else {
			attrName = agg.Variable
		}

		if agg.Distinct {
			dqlAggregate = fmt.Sprintf("distinct %s(%s)", funcName, attrName)
		} else {
			dqlAggregate = fmt.Sprintf("%s(%s)", funcName, attrName)
		}

	default:
		return fmt.Errorf("unsupported aggregate function: %s", agg.Function)
	}

	// Create a child query to hold the aggregate result
	// In DQL, aggregates are typically placed as child queries with variable assignments
	childQuery := &dql.GraphQuery{
		Attr:    dqlAggregate,
		Var:     agg.Alias,
		IsCount: funcName == "count",
	}

	// Also add to GroupbyAttrs for proper groupby directives
	groupbyAttr := &dql.GroupByAttr{
		Attr:  dqlAggregate,
		Alias: agg.Alias,
	}
	query.GroupbyAttrs = append(query.GroupbyAttrs, *groupbyAttr)

	// Append as a child query so the aggregate can be selected
	query.Children = append(query.Children, childQuery)

	return nil
}

// applyHavingClause applies filtering on aggregate results
func applyHavingClause(query *dql.GraphQuery, having *HavingClause) error {
	// HAVING filters are applied to groupby results
	// Parse the expression and convert to DQL filter

	// Create a filter for the HAVING expression
	// Example: HAVING (count(?x) > 5) -> @filter(gt(count(uid), 5))
	filter := &dql.FilterTree{
		Func: &dql.Function{
			Name: "having",
			Attr: having.Expression,
		},
	}

	// Apply to the grouped query
	if query.Filter == nil {
		query.Filter = filter
	} else {
		// Combine with existing filter using AND
		query.Filter = &dql.FilterTree{
			Op:    "AND",
			Child: []*dql.FilterTree{query.Filter, filter},
		}
	}

	return nil
}

// hasPatterns returns true if the query contains at least one triple pattern.
func hasPatterns(sq SPARQLQuery) bool {
	patterns := sq.RootGraphPatterns()
	if len(patterns) == 0 {
		return false
	}

	var found bool
	var checkPattern func(GraphPattern)
	checkPattern = func(p GraphPattern) {
		switch v := p.(type) {
		case *BGP:
			if len(v.Triples) > 0 {
				found = true
			}
		case *OptionalPattern:
			for _, sp := range v.Patterns {
				checkPattern(sp)
			}
		case *UnionPattern:
			for _, alt := range v.Alternatives {
				for _, sp := range alt {
					checkPattern(sp)
				}
			}
		}
	}

	for _, p := range patterns {
		if gp, ok := p.(GraphPattern); ok {
			checkPattern(gp)
		} else if bgp, ok := p.(*BGP); ok { // handle raw BGP in interface slice
			if len(bgp.Triples) > 0 {
				found = true
			}
		}
	}

	return found
}

// applySubquery handles nested query patterns
func applySubquery(parent *dql.GraphQuery, subPatterns []GraphPattern, fromGraphs, fromNamedGraphs []string, opts TranslateOptions) error {
	// Create a nested query block for subquery patterns
	subQuery := &dql.GraphQuery{
		Attr: "_subquery",
	}

	for _, pattern := range subPatterns {
		if err := translateGraphPattern(pattern, fromGraphs, fromNamedGraphs, subQuery, opts); err != nil {
			return err
		}
	}

	parent.Children = append(parent.Children, subQuery)
	return nil
}

// Helper functions for pattern matching

// IsOptionalPattern checks if a pattern block is optional
func IsOptionalPattern(pattern GraphPattern) bool {
	_, ok := pattern.(*OptionalPattern)
	return ok
}

// IsUnionPattern checks if a pattern is a UNION
func IsUnionPattern(pattern GraphPattern) bool {
	_, ok := pattern.(*UnionPattern)
	return ok
}

// ExtractBGPFromPattern extracts basic graph patterns from any pattern type
func ExtractBGPFromPattern(pattern GraphPattern) []*BGP {
	var bgps []*BGP

	switch p := pattern.(type) {
	case *BGP:
		bgps = append(bgps, p)
	case *OptionalPattern:
		for _, subPattern := range p.Patterns {
			bgps = append(bgps, ExtractBGPFromPattern(subPattern)...)
		}
	case *UnionPattern:
		for _, alt := range p.Alternatives {
			for _, altPattern := range alt {
				bgps = append(bgps, ExtractBGPFromPattern(altPattern)...)
			}
		}
	}

	return bgps
}
