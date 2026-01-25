package sparql

import (
	"fmt"
	"sort"
)

// JoinOrderOptimizer optimizes the order of joins in SPARQL patterns.
// It reorders joins to minimize intermediate result sizes by:
// - Analyzing join structure
// - Calculating selectivity per triple
// - Ordering joins from most to least selective
// - Handling multi-way joins
type JoinOrderOptimizer struct {
	Schema *SchemaAnalyzer
}

// NewJoinOrderOptimizer creates a new join order optimizer.
func NewJoinOrderOptimizer(schema *SchemaAnalyzer) *JoinOrderOptimizer {
	return &JoinOrderOptimizer{
		Schema: schema,
	}
}

// OptimizeJoinOrder optimizes the order of joins in an algebra expression.
func (joo *JoinOrderOptimizer) OptimizeJoinOrder(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return nil, fmt.Errorf("expression cannot be nil")
	}

	switch e := expr.(type) {
	case *AlgebraBGP:
		return joo.optimizeBGP(e)
	case *AlgebraJoin:
		return joo.optimizeJoin(e)
	case *AlgebraUnion:
		return joo.optimizeUnion(e)
	case *AlgebraFilter:
		if e.Input != nil {
			optimized, err := joo.OptimizeJoinOrder(e.Input)
			if err != nil {
				return nil, err
			}
			e.Input = optimized
		}
		return e, nil
	default:
		return expr, nil
	}
}

// optimizeBGP optimizes join order within a Basic Graph Pattern.
func (joo *JoinOrderOptimizer) optimizeBGP(bgp *AlgebraBGP) (AlgebraExpr, error) {
	if bgp == nil || len(bgp.Triples) <= 1 {
		return bgp, nil
	}

	// Calculate selectivity for each triple
	triples := bgp.Triples
	selectivities := make([]*tripleSelectivity, len(triples))

	for i, triple := range triples {
		selectivities[i] = &tripleSelectivity{
			Triple:      triple,
			Index:       i,
			Selectivity: joo.calculateSelectivity(triple),
		}
	}

	// Sort by selectivity (most selective first)
	sort.Slice(selectivities, func(i, j int) bool {
		return selectivities[i].Selectivity < selectivities[j].Selectivity
	})

	// Reorder triples
	reordered := make([]*Triple, len(triples))
	for i, ts := range selectivities {
		reordered[i] = ts.Triple
	}

	bgp.Triples = reordered
	return bgp, nil
}

// optimizeJoin optimizes nested join structures.
func (joo *JoinOrderOptimizer) optimizeJoin(join *AlgebraJoin) (AlgebraExpr, error) {
	if join == nil {
		return join, nil
	}

	// Optimize both sides
	var err error
	if join.Left != nil {
		join.Left, err = joo.OptimizeJoinOrder(join.Left)
		if err != nil {
			return nil, err
		}
	}

	if join.Right != nil {
		join.Right, err = joo.OptimizeJoinOrder(join.Right)
		if err != nil {
			return nil, err
		}
	}

	return join, nil
}

// optimizeUnion optimizes all alternatives in a union.
func (joo *JoinOrderOptimizer) optimizeUnion(union *AlgebraUnion) (AlgebraExpr, error) {
	if union == nil || len(union.Alternatives) == 0 {
		return union, nil
	}

	var err error
	for i, expr := range union.Alternatives {
		union.Alternatives[i], err = joo.OptimizeJoinOrder(expr)
		if err != nil {
			return nil, err
		}
	}

	return union, nil
}

// calculateSelectivity calculates the selectivity of a triple.
// Lower selectivity = more restrictive = should be processed first.
// Returns value between 0.0 (highly selective) and 1.0 (not selective).
func (joo *JoinOrderOptimizer) calculateSelectivity(triple *Triple) float64 {
	if triple == nil {
		return 0.5
	}

	selectivity := 1.0

	// Indexed predicates are more selective
	if joo.isIndexed(triple.Predicate) {
		selectivity *= 0.3
	} else if joo.hasReverseIndex(triple.Predicate) {
		selectivity *= 0.5
	}

	// Variables in both subject and object are less selective
	subjectIsVar := isVar(triple.Subject)
	objectIsVar := triple.ObjectIsVar && isVar(triple.Object)

	if !subjectIsVar && objectIsVar {
		// Subject is bound, object is variable - more selective
		selectivity *= 0.7
	} else if subjectIsVar && !objectIsVar {
		// Subject is variable, object is bound - more selective
		selectivity *= 0.7
	} else if !subjectIsVar && !objectIsVar {
		// Both bound - most selective
		selectivity *= 0.1
	}
	// If both are variables, multiply by 1.0 (no change)

	return selectivity
}

// isIndexed checks if a predicate is indexed.
func (joo *JoinOrderOptimizer) isIndexed(predicate string) bool {
	if joo.Schema == nil {
		return false
	}

	predInfo := joo.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return false
	}

	return predInfo.IndexType != ""
}

// hasReverseIndex checks if a predicate has a reverse index.
func (joo *JoinOrderOptimizer) hasReverseIndex(predicate string) bool {
	if joo.Schema == nil {
		return false
	}

	predInfo := joo.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return false
	}

	return predInfo.ReverseIndex
}

// GetJoinStatistics returns statistics about join optimization.
func (joo *JoinOrderOptimizer) GetJoinStatistics(expr AlgebraExpr) *JoinStatistics {
	stats := &JoinStatistics{
		TotalJoins:         0,
		IndexedJoins:       0,
		OptimizedJoins:     0,
		AverageSelectivity: 0.0,
	}

	if expr == nil {
		return stats
	}

	joo.collectStats(expr, stats)
	return stats
}

// collectStats collects join statistics from an expression.
func (joo *JoinOrderOptimizer) collectStats(expr AlgebraExpr, stats *JoinStatistics) {
	if expr == nil {
		return
	}

	switch e := expr.(type) {
	case *AlgebraBGP:
		if len(e.Triples) > 1 {
			stats.TotalJoins += len(e.Triples) - 1
			selectivities := 0.0

			for _, triple := range e.Triples {
				selectivity := joo.calculateSelectivity(triple)
				selectivities += selectivity
				if joo.isIndexed(triple.Predicate) {
					stats.IndexedJoins++
				}
			}

			if len(e.Triples) > 0 {
				stats.AverageSelectivity = selectivities / float64(len(e.Triples))
			}
			stats.OptimizedJoins += len(e.Triples)
		}
	case *AlgebraJoin:
		if e.Left != nil {
			joo.collectStats(e.Left, stats)
		}
		if e.Right != nil {
			joo.collectStats(e.Right, stats)
		}
	case *AlgebraUnion:
		for _, alt := range e.Alternatives {
			joo.collectStats(alt, stats)
		}
	case *AlgebraFilter:
		if e.Input != nil {
			joo.collectStats(e.Input, stats)
		}
	}
}

// tripleSelectivity represents a triple and its calculated selectivity.
type tripleSelectivity struct {
	Triple      *Triple
	Index       int
	Selectivity float64
}

// JoinStatistics contains statistics about joins in a query.
type JoinStatistics struct {
	TotalJoins         int
	IndexedJoins       int
	OptimizedJoins     int
	AverageSelectivity float64
}
