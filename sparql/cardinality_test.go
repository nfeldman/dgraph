package sparql

import (
	"testing"
)

func TestCardinalityEstimatorBasic(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	tests := []struct {
		name     string
		expr     AlgebraExpr
		minCard  float64
		maxCard  float64
		testName string
	}{
		{
			name:     "Empty BGP",
			expr:     &AlgebraBGP{Triples: []*Triple{}},
			minCard:  1,
			maxCard:  1,
			testName: "single_empty_bgp",
		},
		{
			name: "Single triple BGP",
			expr: &AlgebraBGP{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
				},
			},
			minCard:  100,
			maxCard:  10000,
			testName: "single_triple",
		},
		{
			name: "Multiple triples BGP",
			expr: &AlgebraBGP{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
					{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
					{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
				},
			},
			minCard:  0.001,
			maxCard:  100,
			testName: "multiple_triples",
		},
		{
			name:     "Empty result",
			expr:     &AlgebraEmpty{},
			minCard:  0,
			maxCard:  0,
			testName: "empty_result",
		},
		{
			name: "VALUES with 5 rows",
			expr: &AlgebraValues{
				Vars: []string{"?x", "?y"},
				Rows: []map[string]string{
					{"?x": "a", "?y": "1"},
					{"?x": "b", "?y": "2"},
					{"?x": "c", "?y": "3"},
					{"?x": "d", "?y": "4"},
					{"?x": "e", "?y": "5"},
				},
			},
			minCard:  5,
			maxCard:  5,
			testName: "values_5_rows",
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			card := est.Estimate(tt.expr)
			if card < tt.minCard || card > tt.maxCard {
				t.Errorf("cardinality out of bounds: got %.2f, expected range [%.2f, %.2f]",
					card, tt.minCard, tt.maxCard)
			}
		})
	}
}

func TestCardinalityJoin(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	// Create two BGPs
	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	join := &AlgebraJoin{
		Left:  bgp1,
		Right: bgp2,
	}

	card := est.Estimate(join)

	// Join of 1000 * 1000 with 1% selectivity = 10,000
	if card < 100 || card > 100000 {
		t.Errorf("join cardinality out of range: got %.2f", card)
	}
}

func TestCardinalityFilter(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	filter := &AlgebraFilter{
		Expr:  "?age > 18",
		Input: bgp,
	}

	card := est.Estimate(filter)

	// BGP cardinality (1000) * filter selectivity (0.25 for range) = 250
	if card < 100 || card > 500 {
		t.Errorf("filter cardinality out of range: got %.2f", card)
	}
}

func TestCardinalityProject(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	proj := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bgp,
	}

	card := est.Estimate(proj)

	// Projection doesn't change cardinality
	bgpCard := est.Estimate(bgp)
	if card != bgpCard {
		t.Errorf("projection should not change cardinality: got %.2f, expected %.2f",
			card, bgpCard)
	}
}

func TestCardinalityAgg(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// GROUP BY with 2 variables
	agg := &AlgebraAgg{
		Op:    "COUNT",
		Group: []string{"?type", "?status"},
		Input: bgp,
	}

	card := est.Estimate(agg)

	// Should be less than input cardinality due to grouping
	bgpCard := est.Estimate(bgp)
	if card >= bgpCard {
		t.Errorf("aggregate cardinality should be less than input: got %.2f, input %.2f",
			card, bgpCard)
	}

	if card < 1 {
		t.Errorf("aggregate cardinality should be at least 1: got %.2f", card)
	}
}

func TestCardinalityAggNoGroupBy(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// Aggregate without GROUP BY
	agg := &AlgebraAgg{
		Op:    "COUNT",
		Group: []string{},
		Input: bgp,
	}

	card := est.Estimate(agg)

	// Should be exactly 1 (single group)
	if card != 1 {
		t.Errorf("aggregate without GROUP BY should return 1 row: got %.2f", card)
	}
}

func TestCardinalityLimit(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	limit := &AlgebraLimit{
		Count: 10,
		Input: bgp,
	}

	card := est.Estimate(limit)

	// Should be at most 10
	if card > 10 {
		t.Errorf("limit cardinality should be at most 10: got %.2f", card)
	}

	if card < 1 {
		t.Errorf("limit cardinality should be positive: got %.2f", card)
	}
}

func TestCardinalityDistinct(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	distinct := &AlgebraDistinct{
		Input: bgp,
	}

	card := est.Estimate(distinct)
	bgpCard := est.Estimate(bgp)

	// DISTINCT should reduce cardinality
	if card >= bgpCard {
		t.Errorf("distinct should reduce cardinality: got %.2f, input %.2f",
			card, bgpCard)
	}

	if card < bgpCard*0.5 || card > bgpCard {
		t.Errorf("distinct should reduce by 30 percent: got %.2f, input %.2f",
			card, bgpCard)
	}
}

func TestCardinalityUnion(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	alt1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	alt2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{alt1, alt2},
	}

	card := est.Estimate(union)

	// Union should be less than sum (due to overlap)
	alt1Card := est.Estimate(alt1)
	alt2Card := est.Estimate(alt2)
	maxCard := (alt1Card + alt2Card) * 0.9 * 1.1 // 10% margin

	if card > maxCard {
		t.Errorf("union cardinality too high: got %.2f", card)
	}

	if card < 100 {
		t.Errorf("union cardinality should be positive: got %.2f", card)
	}
}

func TestSelectivityAnalysis(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	tests := []struct {
		expr           string
		minSelectivity float64
		maxSelectivity float64
		name           string
	}{
		{
			expr:           "?age > 18",
			minSelectivity: 0.2,
			maxSelectivity: 0.3,
			name:           "range_query",
		},
		{
			expr:           "?status = 'active'",
			minSelectivity: 0.05,
			maxSelectivity: 0.15,
			name:           "equality",
		},
		{
			expr:           "REGEX(?name, '^A')",
			minSelectivity: 0.01,
			maxSelectivity: 0.1,
			name:           "regex",
		},
		{
			expr:           "?age > 18 && ?status = 'active'",
			minSelectivity: 0.001,
			maxSelectivity: 0.05,
			name:           "compound_and",
		},
		{
			expr:           "",
			minSelectivity: 0.99,
			maxSelectivity: 1.01,
			name:           "no_filter",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := &AlgebraFilter{
				Expr:  tt.expr,
				Input: nil,
			}
			sel := est.EstimateSelectivity(filter)

			if sel < tt.minSelectivity || sel > tt.maxSelectivity {
				t.Errorf("selectivity out of range: got %.3f, expected [%.3f, %.3f]",
					sel, tt.minSelectivity, tt.maxSelectivity)
			}
		})
	}
}

func TestCostEstimation(t *testing.T) {
	costEst := NewCardinalityCostEstimator()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	cost := costEst.EstimateCost(bgp)

	// Cost should be positive
	if cost <= 0 {
		t.Errorf("cost should be positive: got %.2f", cost)
	}

	// Cost should be reasonably bounded
	if cost > 100000 {
		t.Errorf("cost seems too high: got %.2f", cost)
	}
}

func TestCostEstimationJoinOrder(t *testing.T) {
	costEst := NewCardinalityCostEstimator()

	// Two alternatives: join(small, large) vs join(large, small)
	small := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	large := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?y", Predicate: "rdf:type", Object: "Thing"},
		},
	}

	join1 := &AlgebraJoin{Left: small, Right: large}
	join2 := &AlgebraJoin{Left: large, Right: small}

	cost1 := costEst.EstimateCost(join1)
	cost2 := costEst.EstimateCost(join2)

	// Both should be positive
	if cost1 <= 0 || cost2 <= 0 {
		t.Errorf("costs should be positive: got %.2f and %.2f", cost1, cost2)
	}

	// Costs should be similar (order shouldn't matter much for cost)
	// but both should be reasonable
	if cost1 > 100000 || cost2 > 100000 {
		t.Errorf("costs seem too high: got %.2f and %.2f", cost1, cost2)
	}
}

func TestCardinalityChainedOperations(t *testing.T) {
	est := NewDefaultCardinalityEstimator()

	// BGP -> Filter -> Project -> Limit
	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	filter := &AlgebraFilter{
		Expr:  "?age > 18",
		Input: bgp,
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: filter,
	}

	limit := &AlgebraLimit{
		Count: 100,
		Input: project,
	}

	bgpCard := est.Estimate(bgp)
	filterCard := est.Estimate(filter)
	projectCard := est.Estimate(project)
	limitCard := est.Estimate(limit)

	// Verify decreasing cardinalities
	if filterCard >= bgpCard {
		t.Errorf("filter should reduce cardinality: %.2f >= %.2f", filterCard, bgpCard)
	}

	if projectCard != filterCard {
		t.Errorf("project should not change cardinality: %.2f != %.2f", projectCard, filterCard)
	}

	if limitCard > projectCard {
		t.Errorf("limit should not increase cardinality: %.2f > %.2f", limitCard, projectCard)
	}

	if limitCard > 100 {
		t.Errorf("limit should cap at 100: got %.2f", limitCard)
	}
}
