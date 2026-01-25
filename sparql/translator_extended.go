package sparql

import (
	"context"
	"fmt"
	"regexp"
	"strings"

	dql "github.com/dgraph-io/dgraph/v25/dql"
	"github.com/dgraph-io/dgraph/v25/protos/pb"
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

// translateGraphPattern handles different pattern types (BGP, OPTIONAL, UNION).
func translateGraphPattern(pattern GraphPattern, fromGraphs, fromNamedGraphs []string, parent *dql.GraphQuery, opts TranslateOptions) error {
	switch p := pattern.(type) {
	case *BGP:
		// Regular basic graph pattern
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
		// OPTIONAL pattern - in DQL, these are naturally optional (missing edges don't fail)
		// Add patterns with a filter to check if the property has a value
		for _, subPattern := range p.Patterns {
			if err := translateGraphPattern(subPattern, fromGraphs, fromNamedGraphs, parent, opts); err != nil {
				return err
			}
		}
		// In DQL, patterns are optional by default if not required
		return nil

	case *UnionPattern:
		// UNION - multiple alternatives
		// In DQL, we can model this as OR filters or multiple query blocks
		// For each alternative, add as separate query path
		for _, alternative := range p.Alternatives {
			for _, altPattern := range alternative {
				if err := translateGraphPattern(altPattern, fromGraphs, fromNamedGraphs, parent, opts); err != nil {
					return err
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

// applyBindExpression applies variable binding expressions (BIND ... AS ...)
func applyBindExpression(query *dql.GraphQuery, bind *BindExpression) error {
	// BIND expressions create new variables from expressions
	// Map to DQL's variable binding system
	// Example: BIND (?x + ?y AS ?sum) -> creates variable with math expression

	// For now, store as a special function node
	// Full implementation would need to parse and evaluate the expression
	if query.Args == nil {
		query.Args = make(map[string]string)
	}
	query.Args["bind_"+bind.Variable] = bind.Expression

	return nil
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
