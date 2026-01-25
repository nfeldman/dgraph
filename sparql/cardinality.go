package sparql

import (
	"math"
	"strings"
)

// CardinalityEstimator estimates the number of rows produced by algebra expressions.
// This is used for cost-based query optimization to select better execution plans.
type CardinalityEstimator interface {
	// Estimate returns the estimated cardinality (number of solution tuples) for an expression.
	Estimate(expr AlgebraExpr) float64

	// EstimateSelectivity returns a value between 0 and 1 indicating the fraction of
	// input rows that satisfy a filter condition. 1.0 means all rows pass, 0.0 means none.
	EstimateSelectivity(filter *AlgebraFilter) float64
}

// DefaultCardinalityEstimator provides cardinality estimates using heuristics.
// It doesn't have access to schema statistics, so it uses pattern-based estimation.
type DefaultCardinalityEstimator struct {
	// MaxResultSize is a hint about the maximum expected result size.
	// Used to bound estimates for unknown patterns.
	MaxResultSize float64

	// BaseTripleCardinality estimates rows returned by a single triple pattern.
	// Used as a starting point for BGP cardinality.
	BaseTripleCardinality float64

	// FilterSelectivity is the default selectivity for filters without analysis.
	FilterSelectivity float64
}

// NewDefaultCardinalityEstimator creates a new estimator with reasonable defaults.
func NewDefaultCardinalityEstimator() *DefaultCardinalityEstimator {
	return &DefaultCardinalityEstimator{
		MaxResultSize:         1000000, // 1M
		BaseTripleCardinality: 1000,    // 1K rows per triple
		FilterSelectivity:     0.3,     // 30% pass rate
	}
}

// Estimate implements CardinalityEstimator.
func (e *DefaultCardinalityEstimator) Estimate(expr AlgebraExpr) float64 {
	if expr == nil {
		return 0
	}

	switch node := expr.(type) {
	case *AlgebraBGP:
		return e.estimateBGP(node)
	case *AlgebraJoin:
		return e.estimateJoin(node)
	case *AlgebraFilter:
		return e.estimateFilter(node)
	case *AlgebraLeftJoin:
		return e.estimateLeftJoin(node)
	case *AlgebraUnion:
		return e.estimateUnion(node)
	case *AlgebraProject:
		return e.estimateProject(node)
	case *AlgebraAgg:
		return e.estimateAgg(node)
	case *AlgebraBind:
		return e.estimateBind(node)
	case *AlgebraDistinct:
		return e.estimateDistinct(node)
	case *AlgebraOrderBy:
		return e.estimateOrderBy(node)
	case *AlgebraLimit:
		return e.estimateLimit(node)
	case *AlgebraValues:
		return e.estimateValues(node)
	case *AlgebraEmpty:
		return 0
	default:
		return e.MaxResultSize
	}
}

// estimateBGP estimates cardinality of a basic graph pattern.
// Heuristic: Each triple adds specificity, reducing the result set.
func (e *DefaultCardinalityEstimator) estimateBGP(bgp *AlgebraBGP) float64 {
	if len(bgp.Triples) == 0 {
		return 1 // Empty BGP returns one empty solution
	}

	// Start with base cardinality for first triple
	cardinality := e.BaseTripleCardinality

	// Each additional triple is a join with selectivity ~0.01 (1% overlap)
	// This models the assumption that join predicates are restrictive
	for i := 1; i < len(bgp.Triples); i++ {
		cardinality *= 0.01 // 1% overlap between joins
	}

	// Bound by maximum
	if cardinality > e.MaxResultSize {
		cardinality = e.MaxResultSize
	}

	return cardinality
}

// estimateJoin estimates cardinality of a join operation.
// Heuristic: Join output is product of inputs times selectivity.
func (e *DefaultCardinalityEstimator) estimateJoin(join *AlgebraJoin) float64 {
	leftCard := e.Estimate(join.Left)
	rightCard := e.Estimate(join.Right)

	// Join selectivity: assume shared variables create ~1% selectivity
	// (only 1% of the cross product is in the result)
	selectivity := 0.01

	result := leftCard * rightCard * selectivity

	// Bound by maximum
	if result > e.MaxResultSize {
		result = e.MaxResultSize
	}

	return result
}

// estimateFilter estimates cardinality after filtering.
func (e *DefaultCardinalityEstimator) estimateFilter(filter *AlgebraFilter) float64 {
	inputCard := e.Estimate(filter.Input)
	selectivity := e.EstimateSelectivity(filter)
	return inputCard * selectivity
}

// estimateLeftJoin estimates cardinality of a left join (OPTIONAL).
// Heuristic: Result includes all input rows plus matching rows from optional side.
func (e *DefaultCardinalityEstimator) estimateLeftJoin(lj *AlgebraLeftJoin) float64 {
	inputCard := e.Estimate(lj.Input)

	// In a left join, we always keep all input rows.
	// Some may be extended with optional pattern matches (~30% match rate)
	// but cardinality stays roughly the same as input
	return inputCard * 1.0 // No cardinality increase
}

// estimateUnion estimates cardinality of a union.
func (e *DefaultCardinalityEstimator) estimateUnion(union *AlgebraUnion) float64 {
	if len(union.Alternatives) == 0 {
		return 0
	}

	// Rough estimate: sum of all alternatives (assuming 10% overlap)
	total := 0.0
	for _, alt := range union.Alternatives {
		total += e.Estimate(alt)
	}

	// Subtract overlap (rough: 10% is duplicates)
	total *= 0.9

	if total > e.MaxResultSize {
		total = e.MaxResultSize
	}

	return total
}

// estimateProject estimates cardinality after projection.
// Heuristic: Projection can reduce cardinality if variables are removed.
func (e *DefaultCardinalityEstimator) estimateProject(proj *AlgebraProject) float64 {
	// Projection doesn't change the number of solution rows
	return e.Estimate(proj.Input)
}

// estimateAgg estimates cardinality of aggregation.
// Heuristic: GROUP BY cardinality is the number of distinct groups.
func (e *DefaultCardinalityEstimator) estimateAgg(agg *AlgebraAgg) float64 {
	inputCard := e.Estimate(agg.Input)

	if len(agg.Group) == 0 {
		// No GROUP BY: aggregates to a single row
		return 1
	}

	// GROUP BY: assume uniform distribution over groups
	// Heuristic: Number of groups is sqrt(input cardinality) * num_group_vars
	numGroupVars := float64(len(agg.Group))
	groupCount := math.Sqrt(inputCard) * numGroupVars

	// Bound by input cardinality
	if groupCount > inputCard {
		groupCount = inputCard
	}

	return groupCount
}

// estimateBind estimates cardinality after variable binding.
// Heuristic: BIND doesn't change cardinality.
func (e *DefaultCardinalityEstimator) estimateBind(bind *AlgebraBind) float64 {
	return e.Estimate(bind.Input)
}

// estimateDistinct estimates cardinality of distinct operation.
// Heuristic: DISTINCT can remove up to 90% of duplicates (depends on data).
func (e *DefaultCardinalityEstimator) estimateDistinct(distinct *AlgebraDistinct) float64 {
	inputCard := e.Estimate(distinct.Input)

	// Conservative estimate: DISTINCT removes some duplicates
	// Assume 30% of rows are duplicates (70% pass through)
	return inputCard * 0.7
}

// estimateOrderBy estimates cardinality after ordering.
// Heuristic: ORDER BY doesn't change cardinality.
func (e *DefaultCardinalityEstimator) estimateOrderBy(orderBy *AlgebraOrderBy) float64 {
	return e.Estimate(orderBy.Input)
}

// estimateLimit estimates cardinality after LIMIT.
func (e *DefaultCardinalityEstimator) estimateLimit(limit *AlgebraLimit) float64 {
	inputCard := e.Estimate(limit.Input)
	limitCard := float64(limit.Count)

	// Result is minimum of input and limit
	if limitCard < inputCard {
		return limitCard
	}
	return inputCard
}

// estimateValues estimates cardinality of VALUES clause.
func (e *DefaultCardinalityEstimator) estimateValues(values *AlgebraValues) float64 {
	return float64(len(values.Rows))
}

// EstimateSelectivity implements CardinalityEstimator.
func (e *DefaultCardinalityEstimator) EstimateSelectivity(filter *AlgebraFilter) float64 {
	if filter == nil || filter.Expr == "" {
		return 1.0 // No filter, all rows pass
	}

	// Analyze the filter expression for selectivity
	return e.analyzeFilterSelectivity(filter.Expr)
}

// analyzeFilterSelectivity analyzes a filter expression and estimates its selectivity.
func (e *DefaultCardinalityEstimator) analyzeFilterSelectivity(expr string) float64 {
	expr = strings.TrimSpace(expr)

	// Handle compound expressions
	if strings.Contains(expr, "&&") {
		// AND: multiply selectivities
		parts := strings.Split(expr, "&&")
		selectivity := 1.0
		for _, part := range parts {
			selectivity *= e.analyzeFilterSelectivity(strings.TrimSpace(part))
		}
		return selectivity
	}

	if strings.Contains(expr, "||") {
		// OR: sum selectivities (rough approximation)
		parts := strings.Split(expr, "||")
		selectivity := 0.0
		for _, part := range parts {
			selectivity += e.analyzeFilterSelectivity(strings.TrimSpace(part))
		}
		// Cap at 1.0
		if selectivity > 1.0 {
			selectivity = 1.0
		}
		return selectivity
	}

	// Handle specific filter types
	if strings.Contains(expr, "=") {
		// Equality: medium selectivity (10%)
		return 0.1
	}

	if strings.Contains(expr, ">") || strings.Contains(expr, "<") {
		// Range: medium-low selectivity (25%)
		return 0.25
	}

	if strings.Contains(expr, "REGEX") {
		// Regex: low selectivity (5%)
		return 0.05
	}

	if strings.Contains(expr, "CONTAINS") || strings.Contains(expr, "STR") {
		// String operations: medium-low (20%)
		return 0.2
	}

	// Default for unknown filters
	return e.FilterSelectivity
}

// CardinalityCostEstimator extends basic cardinality estimation with cost models.
type CardinalityCostEstimator struct {
	*DefaultCardinalityEstimator

	// CostPerRow is the estimated cost to process one row (arbitrary units).
	CostPerRow float64

	// CostPerJoin is the overhead cost for each join operation.
	CostPerJoin float64

	// CostPerSort is the overhead cost for ordering operations.
	CostPerSort float64
}

// NewCardinalityCostEstimator creates a cost estimator with reasonable defaults.
func NewCardinalityCostEstimator() *CardinalityCostEstimator {
	return &CardinalityCostEstimator{
		DefaultCardinalityEstimator: NewDefaultCardinalityEstimator(),
		CostPerRow:                  1.0,
		CostPerJoin:                 10.0,
		CostPerSort:                 5.0,
	}
}

// EstimateCost estimates the total cost of executing an expression.
// Lower cost is better.
func (e *CardinalityCostEstimator) EstimateCost(expr AlgebraExpr) float64 {
	return e.estimateCostRecursive(expr)
}

func (e *CardinalityCostEstimator) estimateCostRecursive(expr AlgebraExpr) float64 {
	if expr == nil {
		return 0
	}

	card := e.Estimate(expr)

	switch node := expr.(type) {
	case *AlgebraJoin:
		// Cost = cost of left + cost of right + join overhead + rows processed
		leftCost := e.estimateCostRecursive(node.Left)
		rightCost := e.estimateCostRecursive(node.Right)
		return leftCost + rightCost + e.CostPerJoin + (card * e.CostPerRow)

	case *AlgebraOrderBy:
		// Cost = input cost + sort overhead * log(cardinality)
		inputCost := e.estimateCostRecursive(node.Input)
		sortCost := e.CostPerSort * math.Log(card+1)
		return inputCost + sortCost

	case *AlgebraFilter:
		// Cost = input cost + filter processing
		inputCost := e.estimateCostRecursive(node.Input)
		return inputCost + (e.Estimate(node.Input) * e.CostPerRow)

	default:
		// Default: cost is cardinality * cost per row
		return card * e.CostPerRow
	}
}
