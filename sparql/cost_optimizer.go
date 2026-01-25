package sparql

import (
	"fmt"
)

// CostBasedJoinReorderer reorders joins using cardinality estimates and cost models.
// It selects join order that minimizes total execution cost.
type CostBasedJoinReorderer struct {
	CardinalityEstimator CardinalityEstimator
	CostEstimator        *CardinalityCostEstimator
}

// NewCostBasedJoinReorderer creates a new cost-based join reorderer.
func NewCostBasedJoinReorderer(cardEst CardinalityEstimator, costEst *CardinalityCostEstimator) *CostBasedJoinReorderer {
	if cardEst == nil {
		cardEst = NewDefaultCardinalityEstimator()
	}
	if costEst == nil {
		costEst = NewCardinalityCostEstimator()
	}
	return &CostBasedJoinReorderer{
		CardinalityEstimator: cardEst,
		CostEstimator:        costEst,
	}
}

// Optimize implements AlgebraOptimizer for cost-based join reordering.
func (r *CostBasedJoinReorderer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraJoin:
		return r.optimizeJoin(node)
	default:
		// Recursively optimize children for non-join expressions
		return r.optimizeExpression(expr)
	}
}

// optimizeJoin reorders a join to minimize cost.
func (r *CostBasedJoinReorderer) optimizeJoin(join *AlgebraJoin) (AlgebraExpr, error) {
	// Flatten the join tree into a list of inputs
	inputs := r.flattenJoins(join)

	if len(inputs) <= 2 {
		// No benefit to reordering just 2 inputs
		// Recursively optimize children instead
		left, err := r.optimizeExpression(join.Left)
		if err != nil {
			return nil, err
		}
		right, err := r.optimizeExpression(join.Right)
		if err != nil {
			return nil, err
		}
		return &AlgebraJoin{Left: left, Right: right}, nil
	}

	// Find the best join order using greedy algorithm
	bestOrder := r.greedyJoinOrdering(inputs)

	// Rebuild the join tree from the optimal order
	result := bestOrder[0]
	for i := 1; i < len(bestOrder); i++ {
		result = &AlgebraJoin{
			Left:  result,
			Right: bestOrder[i],
		}
	}

	return result, nil
}

// optimizeExpression recursively optimizes an expression.
func (r *CostBasedJoinReorderer) optimizeExpression(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraJoin:
		return r.optimizeJoin(node)
	case *AlgebraFilter:
		optimizedInput, err := r.optimizeExpression(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraFilter{
			Expr:  node.Expr,
			Input: optimizedInput,
		}, nil
	case *AlgebraProject:
		optimizedInput, err := r.optimizeExpression(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraProject{
			Vars:  node.Vars,
			Input: optimizedInput,
		}, nil
	default:
		return expr, nil
	}
}

// flattenJoins extracts all inputs from a join tree.
// Converts Join(Join(A, B), C) into [A, B, C]
func (r *CostBasedJoinReorderer) flattenJoins(join AlgebraExpr) []AlgebraExpr {
	var inputs []AlgebraExpr

	var flatten func(AlgebraExpr)
	flatten = func(expr AlgebraExpr) {
		if j, ok := expr.(*AlgebraJoin); ok {
			flatten(j.Left)
			flatten(j.Right)
		} else {
			inputs = append(inputs, expr)
		}
	}

	flatten(join)
	return inputs
}

// greedyJoinOrdering finds a good join order using greedy algorithm.
// Always joins the two inputs with smallest output cost.
func (r *CostBasedJoinReorderer) greedyJoinOrdering(inputs []AlgebraExpr) []AlgebraExpr {
	if len(inputs) <= 1 {
		return inputs
	}

	// Start with the smallest input
	remaining := make([]AlgebraExpr, len(inputs))
	copy(remaining, inputs)

	// Find starting input: the one with smallest cardinality
	startIdx := 0
	minCard := r.CardinalityEstimator.Estimate(remaining[0])
	for i := 1; i < len(remaining); i++ {
		card := r.CardinalityEstimator.Estimate(remaining[i])
		if card < minCard {
			minCard = card
			startIdx = i
		}
	}

	// Swap start input to front
	remaining[0], remaining[startIdx] = remaining[startIdx], remaining[0]

	// Greedy: repeatedly pick the next input that results in smallest join output
	ordered := []AlgebraExpr{remaining[0]}
	remaining = remaining[1:]

	for len(remaining) > 0 {
		// Find the input that minimizes cost when joined with current result
		bestIdx := 0
		bestCost := float64(1e100)

		currentResult := ordered[len(ordered)-1]

		for i, candidate := range remaining {
			// Estimate cost of joining current result with candidate
			testJoin := &AlgebraJoin{
				Left:  currentResult,
				Right: candidate,
			}

			cost := r.CostEstimator.EstimateCost(testJoin)
			if cost < bestCost {
				bestCost = cost
				bestIdx = i
			}
		}

		// Add the best candidate
		ordered = append(ordered, remaining[bestIdx])

		// Remove from remaining
		remaining = append(remaining[:bestIdx], remaining[bestIdx+1:]...)
	}

	return ordered
}

// SelectiveFilterOptimizer reorders filters by selectivity.
// Applies filters with high selectivity (low pass rates) first to reduce data flow.
type SelectiveFilterOptimizer struct {
	CardinalityEstimator CardinalityEstimator
}

// NewSelectiveFilterOptimizer creates a new selective filter optimizer.
func NewSelectiveFilterOptimizer(cardEst CardinalityEstimator) *SelectiveFilterOptimizer {
	if cardEst == nil {
		cardEst = NewDefaultCardinalityEstimator()
	}
	return &SelectiveFilterOptimizer{
		CardinalityEstimator: cardEst,
	}
}

// Optimize implements AlgebraOptimizer for filter selectivity optimization.
func (o *SelectiveFilterOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	// Only optimize chains of filters
	filters := o.extractFilterChain(expr)
	if len(filters) <= 1 {
		// Nothing to optimize - recursively optimize children
		return o.optimizeExpression(expr)
	}

	// Sort filters by selectivity (most selective first)
	o.sortFiltersBySelectivity(filters)

	// Rebuild the expression with reordered filters
	result := filters[len(filters)-1].Input
	for i := len(filters) - 1; i >= 0; i-- {
		result = &AlgebraFilter{
			Expr:  filters[i].Expr,
			Input: result,
		}
	}

	return result, nil
}

// optimizeExpression recursively optimizes an expression.
func (o *SelectiveFilterOptimizer) optimizeExpression(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraFilter:
		optimizedInput, err := o.optimizeExpression(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraFilter{
			Expr:  node.Expr,
			Input: optimizedInput,
		}, nil
	case *AlgebraJoin:
		optimizedLeft, err := o.optimizeExpression(node.Left)
		if err != nil {
			return nil, err
		}
		optimizedRight, err := o.optimizeExpression(node.Right)
		if err != nil {
			return nil, err
		}
		return &AlgebraJoin{
			Left:  optimizedLeft,
			Right: optimizedRight,
		}, nil
	case *AlgebraProject:
		optimizedInput, err := o.optimizeExpression(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraProject{
			Vars:  node.Vars,
			Input: optimizedInput,
		}, nil
	default:
		return expr, nil
	}
}

// extractFilterChain extracts a chain of consecutive filters.
func (o *SelectiveFilterOptimizer) extractFilterChain(expr AlgebraExpr) []*AlgebraFilter {
	var filters []*AlgebraFilter

	current := expr
	for {
		if f, ok := current.(*AlgebraFilter); ok {
			filters = append(filters, f)
			current = f.Input
		} else {
			break
		}
	}

	// Reverse to get filters in order from top to bottom
	for i := len(filters)/2 - 1; i >= 0; i-- {
		opp := len(filters) - 1 - i
		filters[i], filters[opp] = filters[opp], filters[i]
	}

	return filters
}

// sortFiltersBySelectivity sorts filters with lowest selectivity first.
// Low selectivity = fewer rows pass = apply first.
func (o *SelectiveFilterOptimizer) sortFiltersBySelectivity(filters []*AlgebraFilter) {
	// Bubble sort by selectivity (simple but sufficient)
	for i := 0; i < len(filters); i++ {
		for j := i + 1; j < len(filters); j++ {
			selI := o.CardinalityEstimator.EstimateSelectivity(filters[i])
			selJ := o.CardinalityEstimator.EstimateSelectivity(filters[j])

			// Lower selectivity should come first
			if selI > selJ {
				filters[i], filters[j] = filters[j], filters[i]
			}
		}
	}
}

// CostBasedOptimizationPipeline combines multiple cost-based optimizers.
type CostBasedOptimizationPipeline struct {
	JoinOptimizer   *CostBasedJoinReorderer
	FilterOptimizer *SelectiveFilterOptimizer
}

// NewCostBasedOptimizationPipeline creates a pipeline of cost-based optimizers.
func NewCostBasedOptimizationPipeline(cardEst CardinalityEstimator) *CostBasedOptimizationPipeline {
	if cardEst == nil {
		cardEst = NewDefaultCardinalityEstimator()
	}

	costEst := NewCardinalityCostEstimator()

	return &CostBasedOptimizationPipeline{
		JoinOptimizer:   NewCostBasedJoinReorderer(cardEst, costEst),
		FilterOptimizer: NewSelectiveFilterOptimizer(cardEst),
	}
}

// Optimize runs all optimizers in the pipeline.
func (p *CostBasedOptimizationPipeline) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	result := expr

	// Apply filter optimization first
	optimized, err := p.FilterOptimizer.Optimize(result)
	if err != nil {
		return nil, fmt.Errorf("filter optimization error: %w", err)
	}
	result = optimized

	// Then apply join optimization
	optimized, err = p.JoinOptimizer.Optimize(result)
	if err != nil {
		return nil, fmt.Errorf("join optimization error: %w", err)
	}
	result = optimized

	return result, nil
}
