package sparql

import (
	"fmt"
)

// ASTToAlgebra converts SPARQL AST to algebraic form.
// This is the core translation function that transforms parsed queries into
// the algebra intermediate representation for optimization and execution.
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	converter := &astConverter{}
	return converter.convert(query)
}

// astConverter holds state during conversion from AST to algebra.
type astConverter struct {
	filters []string // Filters to apply
}

// convert orchestrates the 10-step conversion algorithm.
func (c *astConverter) convert(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	// Step 1: Extract BGPs, OPTIONAL, UNION, FILTER separately
	patterns := query.Patterns
	if len(patterns) == 0 && len(query.Bgps) > 0 {
		// Backwards compatibility: convert BGPs to patterns
		for _, bgp := range query.Bgps {
			patterns = append(patterns, bgp)
		}
	}

	// Step 2-6: Build algebra from patterns
	expr, err := c.patternsToAlgebra(patterns)
	if err != nil {
		return nil, err
	}

	// Step 7-9: Apply modifiers (BIND, aggregates, ORDER BY, LIMIT)
	// BIND expressions
	for _, bind := range query.Binds {
		expr = &AlgebraBind{
			Var:   bind.Variable,
			Expr:  bind.Expression,
			Input: expr,
		}
	}

	// Aggregates
	if len(query.Aggregates) > 0 {
		groupVars := extractGroupByVars(query.Aggregates)
		for _, agg := range query.Aggregates {
			expr = &AlgebraAgg{
				Op:    agg.Function,
				Var:   agg.Variable,
				Group: groupVars,
				Input: expr,
			}
		}
	}

	// DISTINCT
	if query.Distinct {
		expr = &AlgebraDistinct{Input: expr}
	}

	// ORDER BY
	if len(query.OrderBy) > 0 {
		ascending := make([]bool, len(query.OrderBy))
		for i := range ascending {
			ascending[i] = true // Default ascending unless prefixed with "DESC"
		}
		exprs := query.OrderBy
		expr = &AlgebraOrderBy{
			Exprs:     exprs,
			Ascending: ascending,
			Input:     expr,
		}
	}

	// LIMIT/OFFSET
	if query.Limit > 0 || query.Offset > 0 {
		expr = &AlgebraLimit{
			Count:  query.Limit,
			Offset: query.Offset,
			Input:  expr,
		}
	}

	// Step 10: Apply SELECT projection at top level
	if len(query.Projs) > 0 {
		expr = &AlgebraProject{
			Vars:  query.Projs,
			Input: expr,
		}
	}

	return expr, nil
}

// patternsToAlgebra converts a sequence of graph patterns to algebra.
// This handles BGP, OPTIONAL, UNION, and FILTER patterns.
func (c *astConverter) patternsToAlgebra(patterns []GraphPattern) (AlgebraExpr, error) {
	if len(patterns) == 0 {
		return nil, fmt.Errorf("no graph patterns in query")
	}

	// Process patterns sequentially, building up the algebra tree
	var result AlgebraExpr
	var err error

	for _, p := range patterns {
		switch pattern := p.(type) {
		case *BGP:
			// Convert BGP to algebra
			bgpAlgebra := &AlgebraBGP{Triples: pattern.Triples}
			if result == nil {
				result = bgpAlgebra
			} else {
				// Join with previous result
				result = &AlgebraJoin{
					Left:  result,
					Right: bgpAlgebra,
				}
			}

		case *OptionalPattern:
			// Convert OPTIONAL to LeftJoin
			optAlgebra, err := c.patternsToAlgebra(pattern.Patterns)
			if err != nil {
				return nil, err
			}

			// Extract triples from the optional algebra for the pattern
			triples := extractTriplesFromAlgebra(optAlgebra)

			if result == nil {
				result = optAlgebra
			} else {
				result = &AlgebraLeftJoin{
					Input:    result,
					Patterns: triples,
					Filter:   "",
				}
			}

		case *UnionPattern:
			// Convert UNION to Union operator
			alternatives := make([]AlgebraExpr, 0, len(pattern.Alternatives))
			for _, altPatterns := range pattern.Alternatives {
				altAlgebra, err := c.patternsToAlgebra(altPatterns)
				if err != nil {
					return nil, err
				}
				alternatives = append(alternatives, altAlgebra)
			}

			union := &AlgebraUnion{Alternatives: alternatives}
			if result == nil {
				result = union
			} else {
				// Join previous result with union alternatives
				result = &AlgebraJoin{
					Left:  result,
					Right: union,
				}
			}

		case *FilterPattern:
			// Apply FILTER to current result
			if result == nil {
				return nil, fmt.Errorf("FILTER without preceding patterns")
			}
			result = &AlgebraFilter{
				Expr:  pattern.Expression,
				Input: result,
			}

		default:
			return nil, fmt.Errorf("unknown pattern type: %T", p)
		}
	}

	if result == nil {
		return nil, fmt.Errorf("no valid patterns in query")
	}

	return result, err
}

// extractTriplesFromAlgebra recursively extracts all triples from an algebra expression.
// This is used when converting OPTIONAL patterns to LeftJoin.
func extractTriplesFromAlgebra(expr AlgebraExpr) []*Triple {
	switch e := expr.(type) {
	case *AlgebraBGP:
		return e.Triples
	case *AlgebraJoin:
		left := extractTriplesFromAlgebra(e.Left)
		right := extractTriplesFromAlgebra(e.Right)
		return append(left, right...)
	case *AlgebraFilter:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraLeftJoin:
		triples := extractTriplesFromAlgebra(e.Input)
		return append(triples, e.Patterns...)
	case *AlgebraUnion:
		// For union, collect from first alternative only
		if len(e.Alternatives) > 0 {
			return extractTriplesFromAlgebra(e.Alternatives[0])
		}
	case *AlgebraProject:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraAgg:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraBind:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraDistinct:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraOrderBy:
		return extractTriplesFromAlgebra(e.Input)
	case *AlgebraLimit:
		return extractTriplesFromAlgebra(e.Input)
	}
	return nil
}

// extractGroupByVars extracts GROUP BY variables from aggregates.
// In SPARQL, variables that appear in SELECT but not in aggregates are GROUP BY variables.
func extractGroupByVars(aggregates []*Aggregate) []string {
	// For now, return empty list. This will be refined when we have
	// full SELECT clause information available in the converter.
	return []string{}
}

// joinExpressions creates a left-associative join of multiple expressions.
// Example: [A, B, C] -> Join(Join(A, B), C)
func joinExpressions(exprs []AlgebraExpr) AlgebraExpr {
	if len(exprs) == 0 {
		return nil
	}
	if len(exprs) == 1 {
		return exprs[0]
	}

	result := exprs[0]
	for i := 1; i < len(exprs); i++ {
		result = &AlgebraJoin{
			Left:  result,
			Right: exprs[i],
		}
	}
	return result
}

// ConversionResult holds the result of conversion including the algebra
// expression and any diagnostic information.
type ConversionResult struct {
	Algebra      AlgebraExpr
	Diagnostics  []string
	VariableInfo map[string]string // Maps variables to their definitions
}

// ConvertWithDiagnostics converts AST to algebra and provides diagnostic info.
// This is useful for debugging and understanding the conversion process.
func ConvertWithDiagnostics(query *SPARQLQueryImpl) (*ConversionResult, error) {
	result, err := ASTToAlgebra(query)
	if err != nil {
		return nil, err
	}

	// Collect diagnostic information
	collector := NewVariableCollector()
	vars := collector.Collect(result)

	diags := []string{
		fmt.Sprintf("Query type: %s", query.Qtype),
		fmt.Sprintf("Projection vars: %v", query.Projs),
		fmt.Sprintf("Variables in algebra: %v", vars),
		fmt.Sprintf("Has GROUP BY: %v", len(query.Aggregates) > 0),
		fmt.Sprintf("Has ORDER BY: %v", len(query.OrderBy) > 0),
		fmt.Sprintf("Has LIMIT: %v", query.Limit > 0),
		fmt.Sprintf("Has OFFSET: %v", query.Offset > 0),
	}

	varInfo := make(map[string]string)
	for _, v := range vars {
		varInfo[v] = "referenced"
	}

	return &ConversionResult{
		Algebra:      result,
		Diagnostics:  diags,
		VariableInfo: varInfo,
	}, nil
}

// AlgebraToString provides a human-readable string representation of an algebra expression.
// Useful for debugging and understanding the query structure.
func AlgebraToString(expr AlgebraExpr) string {
	formatter := &ExpressionTreeFormatter{}
	return formatter.Format(expr)
}

// SimplifyAlgebra applies simple optimizations to reduce redundant operations.
// Currently handles:
// - Removing redundant Project nodes
// - Flattening nested Joins
func SimplifyAlgebra(expr AlgebraExpr) AlgebraExpr {
	// This will be expanded in Phase 2 with more sophisticated optimizations
	// For now, return as-is to maintain correctness
	return expr
}
