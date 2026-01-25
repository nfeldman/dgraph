package sparql

import (
	"fmt"
)

// PredicateOptimizer optimizes SPARQL patterns based on predicate information.
// It leverages schema information to:
// - Identify indexed predicates
// - Push indexed predicates forward in queries
// - Optimize pattern order for better performance
// - Handle inverse and reverse predicates
type PredicateOptimizer struct {
	Schema *SchemaAnalyzer
}

// NewPredicateOptimizer creates a new predicate optimizer with schema information.
func NewPredicateOptimizer(schema *SchemaAnalyzer) *PredicateOptimizer {
	return &PredicateOptimizer{
		Schema: schema,
	}
}

// Optimize optimizes algebra expressions based on predicate information.
// It reorders patterns to push indexed predicates forward.
func (po *PredicateOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return nil, fmt.Errorf("expression cannot be nil")
	}

	switch e := expr.(type) {
	case *AlgebraBGP:
		return po.optimizeBGP(e)
	case *AlgebraJoin:
		return po.optimizeJoin(e)
	case *AlgebraUnion:
		return po.optimizeUnion(e)
	case *AlgebraFilter:
		// Push filter optimization to child, but don't reorder filter itself
		optimized, err := po.Optimize(e.Input)
		if err != nil {
			return nil, err
		}
		e.Input = optimized
		return e, nil
	default:
		return expr, nil
	}
}

// optimizeBGP optimizes a Basic Graph Pattern by reordering triples.
// Moves indexed predicates to the front for better performance.
func (po *PredicateOptimizer) optimizeBGP(bgp *AlgebraBGP) (AlgebraExpr, error) {
	if bgp == nil || len(bgp.Triples) == 0 {
		return bgp, nil
	}

	// Classify triples by their indexing characteristics
	indexed := []*Triple{}
	reverse := []*Triple{}
	unindexed := []*Triple{}

	for _, triple := range bgp.Triples {
		if po.isIndexed(triple.Predicate) {
			indexed = append(indexed, triple)
		} else if po.hasReverseIndex(triple.Predicate) {
			reverse = append(reverse, triple)
		} else {
			unindexed = append(unindexed, triple)
		}
	}

	// Reorder: indexed first, then reverse-indexed, then unindexed
	reordered := append(indexed, reverse...)
	reordered = append(reordered, unindexed...)

	bgp.Triples = reordered
	return bgp, nil
}

// optimizeJoin optimizes a join expression by optimizing its children.
func (po *PredicateOptimizer) optimizeJoin(join *AlgebraJoin) (AlgebraExpr, error) {
	var err error

	// Optimize left child
	if join.Left != nil {
		join.Left, err = po.Optimize(join.Left)
		if err != nil {
			return nil, err
		}
	}

	// Optimize right child
	if join.Right != nil {
		join.Right, err = po.Optimize(join.Right)
		if err != nil {
			return nil, err
		}
	}

	return join, nil
}

// optimizeUnion optimizes a union expression by optimizing its children.
func (po *PredicateOptimizer) optimizeUnion(union *AlgebraUnion) (AlgebraExpr, error) {
	if union == nil || len(union.Alternatives) == 0 {
		return union, nil
	}

	var err error
	for i, expr := range union.Alternatives {
		union.Alternatives[i], err = po.Optimize(expr)
		if err != nil {
			return nil, err
		}
	}

	return union, nil
}

// isIndexed checks if a predicate is indexed in the schema.
func (po *PredicateOptimizer) isIndexed(predicate string) bool {
	if po.Schema == nil {
		return false
	}

	// Handle built-in predicates
	if isBuitinPredicate(predicate) {
		return false
	}

	predInfo := po.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return false
	}

	return predInfo.IndexType != ""
}

// hasReverseIndex checks if a predicate has a reverse index.
func (po *PredicateOptimizer) hasReverseIndex(predicate string) bool {
	if po.Schema == nil {
		return false
	}

	// Handle built-in predicates
	if isBuitinPredicate(predicate) {
		return false
	}

	predInfo := po.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return false
	}

	return predInfo.ReverseIndex
}

// isBuitinPredicate checks if a predicate is a built-in SPARQL predicate.
func isBuitinPredicate(predicate string) bool {
	switch predicate {
	case "rdf:type", "a", "rdfs:label", "rdfs:comment":
		return true
	default:
		return false
	}
}

// GetOptimizationInfo returns optimization information for a predicate.
// This provides details about what optimizations apply.
func (po *PredicateOptimizer) GetOptimizationInfo(predicate string) *PredicateOptimizationInfo {
	if po.Schema == nil {
		return &PredicateOptimizationInfo{
			Predicate:   predicate,
			IsIndexed:   false,
			HasReverse:  false,
			IndexType:   "",
			Selectivity: 0.5,
		}
	}

	predInfo := po.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return &PredicateOptimizationInfo{
			Predicate:   predicate,
			IsIndexed:   false,
			HasReverse:  false,
			IndexType:   "",
			Selectivity: 0.5,
		}
	}

	return &PredicateOptimizationInfo{
		Predicate:   predicate,
		IsIndexed:   predInfo.IndexType != "",
		HasReverse:  predInfo.ReverseIndex,
		IndexType:   predInfo.IndexType,
		Selectivity: po.Schema.GetSelectivity(predicate),
	}
}

// PredicateOptimizationInfo contains optimization information for a predicate.
type PredicateOptimizationInfo struct {
	Predicate   string
	IsIndexed   bool
	HasReverse  bool
	IndexType   string
	Selectivity float64
}
