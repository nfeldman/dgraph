package sparql

import (
	"context"
	"fmt"
	"strings"

	dql "github.com/nfeldman/dgraph/dql"
)

// TranslateSelectExtended handles SELECT queries with support for OPTIONAL, UNION, aggregates, BIND, HAVING.
func TranslateSelectExtended(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq.Type() != "SELECT" {
		return nil, nil, fmt.Errorf("expected SELECT query, got %s", sq.Type())
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
			if err := translateGraphPattern(pattern, fromGraphs, fromNamedGraphs, rootQuery); err != nil {
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
		for _, varName := range orderBy {
			order := &dql.Order{
				Attr:  varName,
				Order: "asc",
			}
			rootQuery.Order = append(rootQuery.Order, order)
		}
	}

	return []*dql.GraphQuery{rootQuery}, sq.GetPrefixes(), nil
}

// translateGraphPattern handles different pattern types (BGP, OPTIONAL, UNION).
func translateGraphPattern(pattern GraphPattern, fromGraphs, fromNamedGraphs []string, parent *dql.GraphQuery) error {
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
			if err := translateGraphPattern(subPattern, fromGraphs, fromNamedGraphs, parent); err != nil {
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
				if err := translateGraphPattern(altPattern, fromGraphs, fromNamedGraphs, parent); err != nil {
					return err
				}
			}
		}
		return nil

	default:
		return fmt.Errorf("unknown graph pattern type: %T", pattern)
	}
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

// applyAggregates applies aggregate functions (COUNT, SUM, MIN, MAX, AVG)
func applyAggregates(query *dql.GraphQuery, aggregates []*Aggregate) error {
	// Map SPARQL aggregates to DQL @groupby directive
	// COUNT(?x) -> count(uid)
	// SUM(?x) -> sum(value)
	// MIN(?x) -> min(value)
	// MAX(?x) -> max(value)
	// AVG(?x) -> avg(value)

	for _, agg := range aggregates {
		attr := &dql.GroupByAttr{
			Attr:  agg.Function + "(" + agg.Variable + ")",
			Alias: agg.Alias,
		}
		query.GroupbyAttrs = append(query.GroupbyAttrs, *attr)
		query.IsGroupby = true
	}

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

// applySubquery handles nested query patterns
func applySubquery(parent *dql.GraphQuery, subPatterns []GraphPattern, fromGraphs, fromNamedGraphs []string) error {
	// Create a nested query block for subquery patterns
	subQuery := &dql.GraphQuery{
		Attr: "_subquery",
	}

	for _, pattern := range subPatterns {
		if err := translateGraphPattern(pattern, fromGraphs, fromNamedGraphs, subQuery); err != nil {
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
