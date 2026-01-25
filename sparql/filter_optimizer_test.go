package sparql

import (
	"strings"
	"testing"
)

func TestFilterExpressionConstantFolding(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	// Test cases for constant folding
	tests := []struct {
		expr     string
		expected string
	}{
		{"1 > 2", "false"},
		{"5 = 5", "true"},
		{"10 < 20", "true"},
		{"3 >= 3", "true"},
		{"2 <= 1", "false"},
		{"7 != 7", "false"},
		{"4 != 5", "true"},
	}

	for _, test := range tests {
		result := optimizer.simplifyExpression(test.expr)
		if result != test.expected {
			t.Errorf("simplifyExpression(%q): expected %q, got %q", test.expr, test.expected, result)
		}
	}
}

func TestFilterExpressionRemoveTrivialTrue(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// Filter with always-true condition
	filter := &AlgebraFilter{
		Expr:  "1 = 1",
		Input: bgp,
	}

	optimized, err := optimizer.Optimize(filter)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return just the BGP (filter removed)
	if _, ok := optimized.(*AlgebraBGP); !ok {
		t.Errorf("expected AlgebraBGP, got %T", optimized)
	}
}

func TestFilterExpressionDetectFalse(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	// Filter with always-false condition
	filter := &AlgebraFilter{
		Expr:  "1 > 2",
		Input: bgp,
	}

	optimized, err := optimizer.Optimize(filter)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should return Empty
	if _, ok := optimized.(*AlgebraEmpty); !ok {
		t.Errorf("expected AlgebraEmpty, got %T", optimized)
	}
}

func TestFilterExpressionDoubleNegation(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	expr := "!!?alive"
	result := optimizer.simplifyExpression(expr)
	expected := "?alive"

	if result != expected {
		t.Errorf("simplifyExpression(%q): expected %q, got %q", expr, expected, result)
	}
}

func TestFilterExpressionDeMorgan(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	expr := "!(A && B)"
	result := optimizer.applyDeMorgan(expr)

	// Should apply De Morgan's law
	if !strings.Contains(result, "!A") || !strings.Contains(result, "!B") {
		t.Errorf("applyDeMorgan(%q): expected De Morgan form, got %q", expr, result)
	}
}

func TestFilterExpressionCombineFilters(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
		},
	}

	// Two adjacent filters
	filter1 := &AlgebraFilter{
		Expr: "?age > 18",
		Input: &AlgebraFilter{
			Expr:  "?age < 65",
			Input: bgp,
		},
	}

	optimized, err := optimizer.Optimize(filter1)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should combine into single filter with AND
	if filterNode, ok := optimized.(*AlgebraFilter); ok {
		if !strings.Contains(filterNode.Expr, "&&") {
			t.Errorf("expected combined filter with &&, got %q", filterNode.Expr)
		}
	}
}

func TestFilterExpressionWithProject(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
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

	optimized, err := optimizer.Optimize(project)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should still have project with filter
	if _, ok := optimized.(*AlgebraProject); !ok {
		t.Errorf("expected AlgebraProject, got %T", optimized)
	}
}

func TestFilterExpressionSimplifyIdentity(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	expr := "A && true"
	result := optimizer.simplifyIdentity(expr)

	if strings.Contains(result, "true") {
		t.Errorf("simplifyIdentity(%q): should remove 'true', got %q", expr, result)
	}
}

func TestFilterExpressionTrivialComparison(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	tests := []struct {
		expr      string
		isTrivial bool
		evaluates string
	}{
		{"1 > 2", true, "false"},
		{"5 = 5", true, "true"},
		{"?x > 18", false, ""},
		{"?age = ?age", false, ""},
	}

	for _, test := range tests {
		isTrivial := optimizer.isTrivialComparison(test.expr)
		if isTrivial != test.isTrivial {
			t.Errorf("isTrivialComparison(%q): expected %v, got %v", test.expr, test.isTrivial, isTrivial)
		}

		if isTrivial {
			result := optimizer.evaluateTrivialComparison(test.expr)
			if result != test.evaluates {
				t.Errorf("evaluateTrivialComparison(%q): expected %q, got %q", test.expr, test.evaluates, result)
			}
		}
	}
}

func TestFilterExpressionChain(t *testing.T) {
	optimizer := NewFilterExpressionOptimizer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
		},
	}

	// Chain of filters including one that's always true
	filter1 := &AlgebraFilter{
		Expr: "1 = 1", // Always true - should be removed
		Input: &AlgebraFilter{
			Expr:  "?age > 18",
			Input: bgp,
		},
	}

	optimized, err := optimizer.Optimize(filter1)
	if err != nil {
		t.Fatalf("optimization error: %v", err)
	}

	// Should only have the meaningful filter
	if filterNode, ok := optimized.(*AlgebraFilter); ok {
		if filterNode.Expr == "1 = 1" {
			t.Errorf("trivial filter should have been removed")
		}
	}
}
