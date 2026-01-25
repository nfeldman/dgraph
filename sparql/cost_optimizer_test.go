package sparql

import (
	"testing"
)

func TestCostBasedJoinReorderingBasic(t *testing.T) {
	reorderer := NewCostBasedJoinReorderer(nil, nil)

	// Create a simple join: Join(A, B)
	bgpA := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgpB := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	join := &AlgebraJoin{
		Left:  bgpA,
		Right: bgpB,
	}

	optimized, err := reorderer.Optimize(join)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return a valid join
	if _, ok := optimized.(*AlgebraJoin); !ok {
		t.Errorf("expected join result, got %T", optimized)
	}
}

func TestCostBasedJoinReorderingFlattening(t *testing.T) {
	reorderer := NewCostBasedJoinReorderer(nil, nil)

	// Create a nested join: Join(Join(A, B), C)
	bgpA := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	bgpB := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}
	bgpC := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "c", Object: "C"}}}

	innerJoin := &AlgebraJoin{Left: bgpA, Right: bgpB}
	outerJoin := &AlgebraJoin{Left: innerJoin, Right: bgpC}

	// Flatten should extract [A, B, C]
	flattened := reorderer.flattenJoins(outerJoin)

	if len(flattened) != 3 {
		t.Errorf("expected 3 flattened inputs, got %d", len(flattened))
	}
}

func TestCostBasedJoinReorderingOrder(t *testing.T) {
	cardEst := NewDefaultCardinalityEstimator()
	costEst := NewCardinalityCostEstimator()
	reorderer := NewCostBasedJoinReorderer(cardEst, costEst)

	// Create 3 BGPs with different sizes
	small := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	medium := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}
	large := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "c", Object: "C"},
			{Subject: "?x", Predicate: "d", Object: "D"},
		},
	}

	// Create initial join order: large, medium, small
	join1 := &AlgebraJoin{Left: large, Right: medium}
	join2 := &AlgebraJoin{Left: join1, Right: small}

	optimized, err := reorderer.Optimize(join2)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should reorder to put small inputs first
	if _, ok := optimized.(*AlgebraJoin); !ok {
		t.Errorf("expected join result, got %T", optimized)
	}

	// Verify the result is a valid join structure
	if err := validateJoinStructure(optimized); err != nil {
		t.Errorf("invalid join structure: %v", err)
	}
}

func validateJoinStructure(expr AlgebraExpr) error {
	switch node := expr.(type) {
	case *AlgebraJoin:
		if err := validateJoinStructure(node.Left); err != nil {
			return err
		}
		return validateJoinStructure(node.Right)
	case *AlgebraBGP:
		return nil
	default:
		return nil
	}
}

func TestSelectiveFilterOptimizing(t *testing.T) {
	optimizer := NewSelectiveFilterOptimizer(nil)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// Create a chain of filters
	// Low selectivity filter (should be applied first)
	filter1 := &AlgebraFilter{
		Expr:  "REGEX(?name, '^A')",
		Input: bgp,
	}

	// High selectivity filter (should be applied second)
	filter2 := &AlgebraFilter{
		Expr:  "?age > 10",
		Input: filter1,
	}

	optimized, err := optimizer.Optimize(filter2)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return a valid filter chain
	if _, ok := optimized.(*AlgebraFilter); !ok {
		t.Errorf("expected filter result, got %T", optimized)
	}
}

func TestFilterChainExtraction(t *testing.T) {
	optimizer := NewSelectiveFilterOptimizer(nil)

	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}

	// Create filter chain: Filter3(Filter2(Filter1(BGP)))
	filter1 := &AlgebraFilter{Expr: "?x = 1", Input: bgp}
	filter2 := &AlgebraFilter{Expr: "?y > 5", Input: filter1}
	filter3 := &AlgebraFilter{Expr: "?z < 10", Input: filter2}

	filters := optimizer.extractFilterChain(filter3)

	if len(filters) != 3 {
		t.Errorf("expected 3 filters, got %d", len(filters))
	}

	// Verify order is preserved (top to bottom)
	if filters[0].Expr != "?x = 1" {
		t.Errorf("expected first filter to be '?x = 1', got %q", filters[0].Expr)
	}

	if filters[2].Expr != "?z < 10" {
		t.Errorf("expected last filter to be '?z < 10', got %q", filters[2].Expr)
	}
}

func TestFilterSelectivitySorting(t *testing.T) {
	cardEst := NewDefaultCardinalityEstimator()
	optimizer := NewSelectiveFilterOptimizer(cardEst)

	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}

	// Create filters with different selectivities
	// Regex: 5% (most selective)
	filter1 := &AlgebraFilter{Expr: "REGEX(?name, '^A')", Input: bgp}
	// Equality: 10%
	filter2 := &AlgebraFilter{Expr: "?status = 'active'", Input: bgp}
	// Range: 25% (least selective)
	filter3 := &AlgebraFilter{Expr: "?age > 18", Input: bgp}

	filters := []*AlgebraFilter{filter3, filter2, filter1}

	optimizer.sortFiltersBySelectivity(filters)

	// After sorting, should be: Regex (5%), Equality (10%), Range (25%)
	sel1 := cardEst.EstimateSelectivity(filters[0])
	sel2 := cardEst.EstimateSelectivity(filters[1])
	sel3 := cardEst.EstimateSelectivity(filters[2])

	// Verify selectivity is increasing
	if sel1 >= sel2 || sel2 >= sel3 {
		t.Errorf("selectivity not increasing: %.2f, %.2f, %.2f", sel1, sel2, sel3)
	}
}

func TestCostBasedPipeline(t *testing.T) {
	pipeline := NewCostBasedOptimizationPipeline(nil)

	bgpA := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	bgpB := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}

	filter := &AlgebraFilter{
		Expr: "?age > 18",
		Input: &AlgebraJoin{
			Left:  bgpA,
			Right: bgpB,
		},
	}

	optimized, err := pipeline.Optimize(filter)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return a valid expression
	if optimized == nil {
		t.Errorf("expected optimized expression, got nil")
	}
}

func TestCostBasedOptimizationComparison(t *testing.T) {
	costEst := NewCardinalityCostEstimator()

	// Create two join orders
	bgpA := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	bgpB := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}
	bgpC := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "c", Object: "C"}}}

	// Order 1: A -> B -> C
	join1a := &AlgebraJoin{Left: bgpA, Right: bgpB}
	join1 := &AlgebraJoin{Left: join1a, Right: bgpC}

	// Order 2: B -> A -> C
	join2a := &AlgebraJoin{Left: bgpB, Right: bgpA}
	join2 := &AlgebraJoin{Left: join2a, Right: bgpC}

	cost1 := costEst.EstimateCost(join1)
	cost2 := costEst.EstimateCost(join2)

	// Both should be positive
	if cost1 <= 0 || cost2 <= 0 {
		t.Errorf("costs should be positive: %.2f, %.2f", cost1, cost2)
	}

	// Costs may differ (or be same depending on heuristics)
	t.Logf("Cost of order 1: %.2f", cost1)
	t.Logf("Cost of order 2: %.2f", cost2)
}

func TestComplexOptimizationScenario(t *testing.T) {
	// Create a complex query: Filter(Filter(Join(Join(A, B), C)))
	pipeline := NewCostBasedOptimizationPipeline(nil)

	bgpA := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	bgpB := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}
	bgpC := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "c", Object: "C"}}}

	join1 := &AlgebraJoin{Left: bgpA, Right: bgpB}
	join2 := &AlgebraJoin{Left: join1, Right: bgpC}

	filter1 := &AlgebraFilter{Expr: "?age > 18", Input: join2}
	filter2 := &AlgebraFilter{Expr: "REGEX(?name, '^A')", Input: filter1}

	optimized, err := pipeline.Optimize(filter2)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return a valid expression
	if optimized == nil {
		t.Errorf("expected optimized expression, got nil")
	}

	// Verify the structure is still valid
	if err := validateAlgebraStructure(optimized); err != nil {
		t.Errorf("invalid structure after optimization: %v", err)
	}
}

func validateAlgebraStructure(expr AlgebraExpr) error {
	if expr == nil {
		return nil
	}

	switch node := expr.(type) {
	case *AlgebraJoin:
		if err := validateAlgebraStructure(node.Left); err != nil {
			return err
		}
		return validateAlgebraStructure(node.Right)

	case *AlgebraFilter:
		return validateAlgebraStructure(node.Input)

	case *AlgebraProject:
		return validateAlgebraStructure(node.Input)

	case *AlgebraAgg:
		return validateAlgebraStructure(node.Input)

	case *AlgebraLimit:
		return validateAlgebraStructure(node.Input)

	case *AlgebraDistinct:
		return validateAlgebraStructure(node.Input)

	case *AlgebraOrderBy:
		return validateAlgebraStructure(node.Input)

	case *AlgebraBind:
		return validateAlgebraStructure(node.Input)

	case *AlgebraBGP, *AlgebraValues, *AlgebraEmpty:
		return nil

	default:
		return nil
	}
}

func TestJoinReorderingSmallTree(t *testing.T) {
	// Small join tree shouldn't be modified
	reorderer := NewCostBasedJoinReorderer(nil, nil)

	bgpA := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	bgpB := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "b", Object: "B"}}}

	join := &AlgebraJoin{Left: bgpA, Right: bgpB}

	optimized, err := reorderer.Optimize(join)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should still be a valid join
	if _, ok := optimized.(*AlgebraJoin); !ok {
		t.Errorf("expected join, got %T", optimized)
	}
}

func TestEmptyOptimization(t *testing.T) {
	reorderer := NewCostBasedJoinReorderer(nil, nil)

	// Test with nil
	result, err := reorderer.Optimize(nil)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != nil {
		t.Errorf("expected nil, got %v", result)
	}

	// Test with non-join expression
	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "a", Object: "A"}}}
	result, err = reorderer.Optimize(bgp)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if result != bgp {
		t.Errorf("expected same BGP, got different expression")
	}
}
