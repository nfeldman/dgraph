package sparql

import (
	"testing"
)

func TestBindOptimizationRemoveUnused(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// BIND creates ?unused which is never referenced
	bind := &AlgebraBind{
		Var:   "?unused",
		Expr:  "CONCAT(?name, ' Smith')",
		Input: bgp,
	}

	// Project only ?x
	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bind,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// The result should skip the unused BIND
	// Should be Project(BGP), not Project(Bind(BGP))
	if projNode, ok := optimized.(*AlgebraProject); ok {
		if _, ok := projNode.Input.(*AlgebraBind); ok {
			t.Errorf("expected BIND to be removed, but it's still there")
		}
	} else {
		t.Errorf("expected AlgebraProject, got %T", optimized)
	}
}

func TestBindOptimizationKeepUsed(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// BIND creates ?fullName which IS referenced in SELECT
	bind := &AlgebraBind{
		Var:   "?fullName",
		Expr:  "CONCAT(?name, ' Smith')",
		Input: bgp,
	}

	// Project includes ?fullName (so it's used)
	project := &AlgebraProject{
		Vars:  []string{"?x", "?fullName"},
		Input: bind,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// The result should keep the BIND (it's used)
	if projNode, ok := optimized.(*AlgebraProject); ok {
		if _, ok := projNode.Input.(*AlgebraBind); !ok {
			t.Errorf("expected BIND to be kept (it's used in SELECT), but it was removed")
		}
	} else {
		t.Errorf("expected AlgebraProject, got %T", optimized)
	}
}

func TestBindOptimizationChain(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// Multiple BINDs: ?unused1, ?used, ?unused2
	bind1 := &AlgebraBind{
		Var:   "?unused1",
		Expr:  "CONCAT(?name, ' X')",
		Input: bgp,
	}

	bind2 := &AlgebraBind{
		Var:   "?used",
		Expr:  "CONCAT(?name, ' Y')",
		Input: bind1,
	}

	bind3 := &AlgebraBind{
		Var:   "?unused2",
		Expr:  "CONCAT(?name, ' Z')",
		Input: bind2,
	}

	// Project only ?used
	project := &AlgebraProject{
		Vars:  []string{"?x", "?used"},
		Input: bind3,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should remove unused BINDs but keep the one that's used
	validateBindChain(t, optimized, []string{"?used"}, []string{"?unused1", "?unused2"})
}

func TestBindOptimizationWithFilter(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// BIND creates ?age (used in FILTER)
	bind := &AlgebraBind{
		Var:   "?age",
		Expr:  "25",
		Input: bgp,
	}

	// FILTER uses ?age
	filter := &AlgebraFilter{
		Expr:  "?age > 18",
		Input: bind,
	}

	// Project without ?age (but it's used in FILTER)
	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: filter,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// The BIND should be kept because it's used in the FILTER
	if filterNode, ok := optimized.(*AlgebraProject).Input.(*AlgebraFilter); ok {
		if _, ok := filterNode.Input.(*AlgebraBind); !ok {
			t.Errorf("expected BIND to be kept (used in FILTER), but it was removed")
		}
	}
}

func TestBindOptimizationWithAgg(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// BIND creates ?type (used in GROUP BY)
	bind := &AlgebraBind{
		Var:   "?type",
		Expr:  "'Person'",
		Input: bgp,
	}

	// GROUP BY uses ?type
	agg := &AlgebraAgg{
		Op:    "COUNT",
		Var:   "?count",
		Group: []string{"?type"},
		Input: bind,
	}

	optimized, err := optimizer.Optimize(agg)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// The BIND should be kept because it's used in GROUP BY
	if aggNode, ok := optimized.(*AlgebraAgg); ok {
		if _, ok := aggNode.Input.(*AlgebraBind); !ok {
			t.Errorf("expected BIND to be kept (used in GROUP BY), but it was removed")
		}
	}
}

func TestBindOptimizationMultipleUnused(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// Create 3 unused BINDs
	bind1 := &AlgebraBind{Var: "?unused1", Expr: "1", Input: bgp}
	bind2 := &AlgebraBind{Var: "?unused2", Expr: "2", Input: bind1}
	bind3 := &AlgebraBind{Var: "?unused3", Expr: "3", Input: bind2}

	// Project only ?x
	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bind3,
	}

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// All BINDs should be removed
	if projNode, ok := optimized.(*AlgebraProject); ok {
		if _, ok := projNode.Input.(*AlgebraBGP); !ok {
			t.Errorf("expected all BINDs to be removed, got input type %T", projNode.Input)
		}
	}
}

func TestBindOptimizationEmptyProject(t *testing.T) {
	optimizer := NewBindExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	bind := &AlgebraBind{
		Var:   "?unused",
		Expr:  "CONCAT(?name, ' X')",
		Input: bgp,
	}

	optimized, err := optimizer.Optimize(bind)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Without PROJECT, all variables are potentially unused
	// Should still remove the BIND if it's truly unused
	if _, ok := optimized.(*AlgebraBGP); !ok {
		t.Errorf("expected BIND to be removed (no project to determine usage)")
	}
}

// Helper function to validate BIND chains
func validateBindChain(t *testing.T, expr AlgebraExpr, expectedUsed, expectedUnused []string) {
	// Walk the tree and count BINDs
	foundUsed := make(map[string]bool)
	_ = expectedUnused // Note: expectedUnused is reserved for future validation

	var walk func(AlgebraExpr)
	walk = func(e AlgebraExpr) {
		if e == nil {
			return
		}
		switch node := e.(type) {
		case *AlgebraBind:
			// This BIND was kept, so it must be in expectedUsed
			found := false
			for _, v := range expectedUsed {
				if node.Var == v {
					found = true
					foundUsed[v] = true
				}
			}
			if !found {
				t.Errorf("found kept BIND %s but it's not in expectedUsed list", node.Var)
			}
			walk(node.Input)
		case *AlgebraProject:
			walk(node.Input)
		default:
			walk(expr)
		}
	}

	walk(expr)

	// Check that expected used BINDs are present
	for _, v := range expectedUsed {
		if !foundUsed[v] {
			t.Errorf("expected BIND %s to be kept, but it wasn't found", v)
		}
	}
}
