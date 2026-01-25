package sparql

import (
	"math"
)

// IndexSelectionOptimizer selects optimal indexes for query execution.
// It:
// - Analyzes available indexes
// - Matches predicates to indexes
// - Calculates index cost
// - Selects the best index for each predicate
type IndexSelectionOptimizer struct {
	Schema *SchemaAnalyzer
}

// NewIndexSelectionOptimizer creates a new index selection optimizer.
func NewIndexSelectionOptimizer(schema *SchemaAnalyzer) *IndexSelectionOptimizer {
	return &IndexSelectionOptimizer{
		Schema: schema,
	}
}

// IndexInfo contains information about an index.
type IndexInfo struct {
	Predicate   string
	IndexType   string
	Selectivity float64
	Cost        float64
	Recommended bool
}

// SelectIndexes analyzes an expression and selects optimal indexes.
func (iso *IndexSelectionOptimizer) SelectIndexes(expr AlgebraExpr) map[string]*IndexInfo {
	result := make(map[string]*IndexInfo)

	if expr == nil {
		return result
	}

	iso.collectPredicates(expr, result)
	iso.calculateIndexCosts(result)

	return result
}

// collectPredicates collects all predicates from an expression.
func (iso *IndexSelectionOptimizer) collectPredicates(expr AlgebraExpr, indexes map[string]*IndexInfo) {
	if expr == nil {
		return
	}

	switch e := expr.(type) {
	case *AlgebraBGP:
		if e != nil {
			for _, triple := range e.Triples {
				if triple != nil && triple.Predicate != "" {
					if _, exists := indexes[triple.Predicate]; !exists {
						indexes[triple.Predicate] = iso.getBestIndex(triple.Predicate)
					}
				}
			}
		}
	case *AlgebraJoin:
		if e.Left != nil {
			iso.collectPredicates(e.Left, indexes)
		}
		if e.Right != nil {
			iso.collectPredicates(e.Right, indexes)
		}
	case *AlgebraUnion:
		for _, alt := range e.Alternatives {
			iso.collectPredicates(alt, indexes)
		}
	case *AlgebraFilter:
		if e.Input != nil {
			iso.collectPredicates(e.Input, indexes)
		}
	}
}

// getBestIndex returns the best index for a predicate.
func (iso *IndexSelectionOptimizer) getBestIndex(predicate string) *IndexInfo {
	if iso.Schema == nil {
		return &IndexInfo{
			Predicate:   predicate,
			IndexType:   "",
			Selectivity: 0.5,
			Cost:        0.5,
			Recommended: false,
		}
	}

	predInfo := iso.Schema.GetPredicate(predicate)
	if predInfo == nil {
		return &IndexInfo{
			Predicate:   predicate,
			IndexType:   "",
			Selectivity: 0.5,
			Cost:        1.0,
			Recommended: false,
		}
	}

	info := &IndexInfo{
		Predicate:   predicate,
		IndexType:   predInfo.IndexType,
		Selectivity: iso.Schema.GetSelectivity(predicate),
		Recommended: predInfo.IndexType != "",
	}

	info.Cost = iso.calculateIndexCost(info)

	return info
}

// calculateIndexCosts calculates the cost for all indexes.
func (iso *IndexSelectionOptimizer) calculateIndexCosts(indexes map[string]*IndexInfo) {
	for _, info := range indexes {
		if info != nil {
			info.Cost = iso.calculateIndexCost(info)
		}
	}
}

// calculateIndexCost calculates the cost of using an index.
// Lower cost is better. Cost is based on:
// - Whether index exists
// - Index type (hash < range < tokenizer)
// - Selectivity
// - Cardinality
func (iso *IndexSelectionOptimizer) calculateIndexCost(info *IndexInfo) float64 {
	if info == nil {
		return math.MaxFloat64
	}

	// Base cost is inverse of selectivity
	// More selective = lower cost
	baseCost := 1.0 / info.Selectivity

	// Adjust for index type
	indexMultiplier := 1.0
	switch info.IndexType {
	case "hash":
		indexMultiplier = 0.5 // Best for exact match
	case "range":
		indexMultiplier = 0.7 // Good for comparisons
	case "":
		indexMultiplier = 2.0 // No index is expensive
	default:
		indexMultiplier = 1.0
	}

	// Get cardinality if available
	cardinalityFactor := 1.0
	if iso.Schema != nil {
		predInfo := iso.Schema.GetPredicate(info.Predicate)
		if predInfo != nil && predInfo.Cardinality > 0 {
			// Log of cardinality as factor
			cardinalityFactor = math.Log(float64(predInfo.Cardinality) + 1.0)
		}
	}

	cost := baseCost * indexMultiplier / math.Max(1.0, cardinalityFactor)
	return math.Max(cost, 0.1) // Ensure minimum cost
}

// GetIndexRecommendations returns index recommendations for optimization.
func (iso *IndexSelectionOptimizer) GetIndexRecommendations(expr AlgebraExpr) []*IndexRecommendation {
	recommendations := make([]*IndexRecommendation, 0)

	indexes := iso.SelectIndexes(expr)

	for predicate, info := range indexes {
		if info != nil && info.IndexType == "" {
			// Recommend creating an index for unindexed predicates
			recommendations = append(recommendations, &IndexRecommendation{
				Predicate:     predicate,
				RecommendType: iso.getRecommendedIndexType(predicate),
				Priority:      iso.calculateRecommendationPriority(info),
				Reason:        "Frequently used predicate without index",
			})
		}
	}

	return recommendations
}

// getRecommendedIndexType recommends an index type for a predicate.
func (iso *IndexSelectionOptimizer) getRecommendedIndexType(predicate string) string {
	// Simple heuristic: recommend hash for most cases
	// In the future, could analyze predicate usage patterns
	return "hash"
}

// calculateRecommendationPriority calculates priority for an index recommendation.
func (iso *IndexSelectionOptimizer) calculateRecommendationPriority(info *IndexInfo) float64 {
	// Higher selectivity = higher priority
	return 1.0 - info.Selectivity
}

// AnalyzeIndexUsage analyzes how well indexes are being used.
func (iso *IndexSelectionOptimizer) AnalyzeIndexUsage(expr AlgebraExpr) *IndexUsageAnalysis {
	analysis := &IndexUsageAnalysis{
		TotalPredicates:     0,
		IndexedPredicates:   0,
		UnindexedPredicates: 0,
		TotalIndexCost:      0.0,
		AverageIndexCost:    0.0,
	}

	indexes := iso.SelectIndexes(expr)

	for _, info := range indexes {
		if info != nil {
			analysis.TotalPredicates++

			if info.IndexType != "" {
				analysis.IndexedPredicates++
			} else {
				analysis.UnindexedPredicates++
			}

			analysis.TotalIndexCost += info.Cost
		}
	}

	if analysis.TotalPredicates > 0 {
		analysis.AverageIndexCost = analysis.TotalIndexCost / float64(analysis.TotalPredicates)
	}

	return analysis
}

// IndexRecommendation contains a recommendation for creating an index.
type IndexRecommendation struct {
	Predicate     string
	RecommendType string
	Priority      float64
	Reason        string
}

// IndexUsageAnalysis contains analysis of index usage in a query.
type IndexUsageAnalysis struct {
	TotalPredicates     int
	IndexedPredicates   int
	UnindexedPredicates int
	TotalIndexCost      float64
	AverageIndexCost    float64
}

// GetIndexCost returns the cost for a specific predicate's index.
func (iso *IndexSelectionOptimizer) GetIndexCost(predicate string) float64 {
	info := iso.getBestIndex(predicate)
	if info == nil {
		return 1.0
	}
	return info.Cost
}

// IsIndexedPredicate checks if a predicate has an index.
func (iso *IndexSelectionOptimizer) IsIndexedPredicate(predicate string) bool {
	info := iso.getBestIndex(predicate)
	if info == nil {
		return false
	}
	return info.IndexType != ""
}
