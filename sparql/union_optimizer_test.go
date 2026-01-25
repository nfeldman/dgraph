package sparql

import (
	"testing"
)

func TestUnionOptimizationSingleBranch(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// UNION with single branch should be removed
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return just the BGP
	if _, ok := optimized.(*AlgebraBGP); !ok {
		t.Errorf("expected AlgebraBGP, got %T", optimized)
	}
}

func TestUnionOptimizationRemoveDuplicates(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp3 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	// UNION with duplicate branches
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2, bgp3},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should have 2 branches (duplicates removed)
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 2 {
			t.Errorf("expected 2 branches after removing duplicates, got %d", len(unionNode.Alternatives))
		}
	} else {
		t.Errorf("expected AlgebraUnion, got %T", optimized)
	}
}

func TestUnionOptimizationFlattenNested(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	bgp3 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Agent"},
		},
	}

	// Nested UNIONs: UNION(UNION(A, B), C)
	innerUnion := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	outerUnion := &AlgebraUnion{
		Alternatives: []AlgebraExpr{innerUnion, bgp3},
	}

	optimized, err := optimizer.Optimize(outerUnion)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should flatten to UNION(A, B, C) with 3 branches
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 3 {
			t.Errorf("expected 3 branches after flattening, got %d", len(unionNode.Alternatives))
		}
	} else {
		t.Errorf("expected AlgebraUnion, got %T", optimized)
	}
}

func TestUnionOptimizationEmptyBranches(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	empty := &AlgebraEmpty{}

	// UNION with empty branch
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp, empty},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Empty branches should be handled (kept in union)
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 2 {
			t.Errorf("expected 2 branches, got %d", len(unionNode.Alternatives))
		}
	}
}

func TestUnionOptimizationWithProject(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: union,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should still have project
	if projNode, ok := optimized.(*AlgebraProject); ok {
		if _, ok := projNode.Input.(*AlgebraUnion); !ok {
			t.Errorf("expected UNION inside PROJECT")
		}
	} else {
		t.Errorf("expected AlgebraProject, got %T", optimized)
	}
}

func TestUnionOptimizationWithFilter(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	filter := &AlgebraFilter{
		Expr:  "?x = 'John'",
		Input: union,
	}

	optimized, err := optimizer.Optimize(filter)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should still have filter with union inside
	if filterNode, ok := optimized.(*AlgebraFilter); ok {
		if _, ok := filterNode.Input.(*AlgebraUnion); !ok {
			t.Errorf("expected UNION inside FILTER")
		}
	} else {
		t.Errorf("expected AlgebraFilter, got %T", optimized)
	}
}

func TestUnionOptimizationDuplicateFilters(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	filter1 := &AlgebraFilter{
		Expr:  "?x = 'John'",
		Input: bgp,
	}

	// Create a separate copy of the same filter
	filter2 := &AlgebraFilter{
		Expr:  "?x = 'John'",
		Input: bgp,
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{filter1, filter2},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should remove duplicate filter branches (same signature)
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 1 {
			t.Errorf("expected 1 branch after removing duplicates, got %d", len(unionNode.Alternatives))
		}
	} else if _, ok := optimized.(*AlgebraFilter); ok {
		// If only one branch, it should return just the filter
		// This is correct behavior
	} else {
		t.Errorf("expected AlgebraUnion or AlgebraFilter, got %T", optimized)
	}
}

func TestUnionOptimizationExpressionSignature(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	sig1 := optimizer.getExpressionSignature(bgp1)
	sig2 := optimizer.getExpressionSignature(bgp2)

	if sig1 != sig2 {
		t.Errorf("identical BGPs should have same signature: %q vs %q", sig1, sig2)
	}
}

func TestUnionOptimizationMultipleLevelsNesting(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p1", Object: "o1"}}}
	bgp2 := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p2", Object: "o2"}}}
	bgp3 := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p3", Object: "o3"}}}
	bgp4 := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p4", Object: "o4"}}}

	// UNION(UNION(UNION(A, B), C), D)
	level1 := &AlgebraUnion{Alternatives: []AlgebraExpr{bgp1, bgp2}}
	level2 := &AlgebraUnion{Alternatives: []AlgebraExpr{level1, bgp3}}
	level3 := &AlgebraUnion{Alternatives: []AlgebraExpr{level2, bgp4}}

	optimized, err := optimizer.Optimize(level3)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should flatten to 4 branches
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 4 {
			t.Errorf("expected 4 branches after flattening nested unions, got %d", len(unionNode.Alternatives))
		}
	} else {
		t.Errorf("expected AlgebraUnion, got %T", optimized)
	}
}

func TestUnionOptimizationOnlyDuplicates(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// UNION with only duplicate branches
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp, bgp, bgp},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should reduce to single branch and remove UNION
	if _, ok := optimized.(*AlgebraBGP); !ok {
		t.Errorf("expected AlgebraBGP, got %T", optimized)
	}
}

func TestUnionOptimizationWithJoin(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1a := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p", Object: "o"}}}
	bgp1b := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "name", Object: "?n", ObjectIsVar: true}}}
	join1 := &AlgebraJoin{Left: bgp1a, Right: bgp1b}

	bgp2a := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "type", Object: "Organization"}}}
	bgp2b := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "name", Object: "?n", ObjectIsVar: true}}}
	join2 := &AlgebraJoin{Left: bgp2a, Right: bgp2b}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{join1, join2},
	}

	optimized, err := optimizer.Optimize(union)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should keep union with 2 join branches
	if unionNode, ok := optimized.(*AlgebraUnion); ok {
		if len(unionNode.Alternatives) != 2 {
			t.Errorf("expected 2 branches, got %d", len(unionNode.Alternatives))
		}
	}
}

func TestUnionOptimizationRecursive(t *testing.T) {
	optimizer := NewUnionPatternOptimizer()

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	project := &AlgebraProject{
		Vars: []string{"?x"},
		Input: &AlgebraFilter{
			Expr: "true",
			Input: &AlgebraDistinct{
				Input: union,
			},
		},
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should recursively optimize all levels
	if _, ok := optimized.(*AlgebraProject); !ok {
		t.Errorf("expected AlgebraProject, got %T", optimized)
	}
}
