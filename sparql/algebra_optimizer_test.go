package sparql

import (
	"strings"
	"testing"
)

func TestFilterPushdownIntoJoin(t *testing.T) {
	bgpLeft := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false}}}
	bgpRight := &AlgebraBGP{Triples: []*Triple{{Subject: "?y", Predicate: "q", Object: "?y", ObjectIsVar: true}}}
	join := &AlgebraJoin{Left: bgpLeft, Right: bgpRight}
	filter := &AlgebraFilter{Expr: "?y > 5", Input: join}

	op := NewAlgebraOptimizer(filterPushdownRule{})
	opt := op.Optimize(filter)

	outer, ok := opt.(*AlgebraJoin)
	if !ok {
		t.Fatalf("expected Join after pushdown, got %T", opt)
	}
	if _, ok := outer.Right.(*AlgebraFilter); !ok {
		t.Fatalf("expected filter pushed to right child")
	}
}

func TestFilterPushdownIntoUnion(t *testing.T) {
	u := &AlgebraUnion{Alternatives: []AlgebraExpr{
		&AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false}}},
		&AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "q", Object: "o", ObjectIsVar: false}}},
	}}
	filter := &AlgebraFilter{Expr: "?x = 1", Input: u}

	op := NewAlgebraOptimizer(filterPushdownRule{})
	opt := op.Optimize(filter)

	union, ok := opt.(*AlgebraUnion)
	if !ok {
		t.Fatalf("expected Union after pushdown, got %T", opt)
	}
	for i, alt := range union.Alternatives {
		if _, ok := alt.(*AlgebraFilter); !ok {
			t.Fatalf("alternative %d missing pushed filter", i)
		}
	}
}

func TestJoinReorderByCost(t *testing.T) {
	bgp3 := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?a", Predicate: "p", Object: "o1", ObjectIsVar: false},
		{Subject: "?a", Predicate: "p", Object: "o2", ObjectIsVar: false},
		{Subject: "?a", Predicate: "p", Object: "o3", ObjectIsVar: false},
	}}
	bgp1 := &AlgebraBGP{Triples: []*Triple{{Subject: "?b", Predicate: "r", Object: "o", ObjectIsVar: false}}}
	bgp2 := &AlgebraBGP{Triples: []*Triple{{Subject: "?c", Predicate: "s", Object: "o", ObjectIsVar: false}, {Subject: "?c", Predicate: "s2", Object: "o2", ObjectIsVar: false}}}

	join := &AlgebraJoin{Left: &AlgebraJoin{Left: bgp3, Right: bgp1}, Right: bgp2}

	op := NewAlgebraOptimizer(joinReorderRule{})
	opt := op.Optimize(join)

	order := collectJoinOrder(opt)
	expected := []AlgebraExpr{bgp1, bgp2, bgp3}
	if len(order) != len(expected) {
		t.Fatalf("unexpected join order length: %d", len(order))
	}
	for i := range order {
		if order[i] != expected[i] {
			t.Fatalf("join order mismatch at %d: got %T, want %T", i, order[i], expected[i])
		}
	}
}

func TestOptionalSimplificationWhenVarsUnused(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false}}}
	optPattern := []*Triple{{Subject: "?x", Predicate: "q", Object: "?y", ObjectIsVar: true}}
	lj := &AlgebraLeftJoin{Input: input, Patterns: optPattern}
	proj := &AlgebraProject{Vars: []string{"?x"}, Input: lj}

	op := NewAlgebraOptimizer(deadVariableEliminationRule{})
	opt := op.Optimize(proj)

	if _, ok := opt.(*AlgebraProject); !ok {
		t.Fatalf("expected Project at top, got %T", opt)
	}
	if _, ok := opt.(*AlgebraProject).Input.(*AlgebraLeftJoin); !ok {
		t.Fatalf("optional should be retained unless trivially empty")
	}
}

func TestDeadBindElimination(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false}}}
	bind := &AlgebraBind{Var: "?y", Expr: "?x + 1", Input: input}
	proj := &AlgebraProject{Vars: []string{"?x"}, Input: bind}

	op := NewAlgebraOptimizer(deadVariableEliminationRule{})
	opt := op.Optimize(proj)

	if _, ok := opt.(*AlgebraProject).Input.(*AlgebraBGP); !ok {
		t.Fatalf("bind should be removed when ?y not projected")
	}
}

func collectJoinOrder(expr AlgebraExpr) []AlgebraExpr {
	var out []AlgebraExpr
	switch n := expr.(type) {
	case *AlgebraJoin:
		out = append(out, collectJoinOrder(n.Left)...)
		out = append(out, collectJoinOrder(n.Right)...)
	default:
		out = append(out, n)
	}
	return out
}

func TestValuesProjectionSimplification(t *testing.T) {
	values := &AlgebraValues{Vars: []string{"?x", "?y"}, Rows: []map[string]string{
		{"?x": "1", "?y": "2"},
		{"?x": "1", "?y": "2"},
	}}
	proj := &AlgebraProject{Vars: []string{"?x"}, Input: values}

	op := NewAlgebraOptimizer(valuesPropagationRule{}, identitySimplificationRule{})
	opt := op.Optimize(proj)

	out, ok := opt.(*AlgebraValues)
	if !ok {
		t.Fatalf("expected AlgebraValues after simplification, got %T", opt)
	}
	if len(out.Vars) != 1 || out.Vars[0] != "?x" {
		t.Fatalf("unexpected projected vars: %v", out.Vars)
	}
	if len(out.Rows) != 1 {
		t.Fatalf("expected deduplicated single row, got %d", len(out.Rows))
	}
	if out.Rows[0]["?x"] != "1" {
		t.Fatalf("unexpected value for ?x: %s", out.Rows[0]["?x"])
	}
}

func TestIdentitySimplificationEliminatesEmptyValues(t *testing.T) {
	values := &AlgebraValues{Vars: []string{"?x"}, Rows: nil}

	op := NewAlgebraOptimizer(identitySimplificationRule{})
	opt := op.Optimize(values)

	if _, ok := opt.(*AlgebraEmpty); !ok {
		t.Fatalf("expected empty algebra, got %T", opt)
	}
}

func TestUnionDeduplicationAndEmptyRemoval(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?s", Predicate: "p", Object: "o", ObjectIsVar: false}}}
	bgp2 := &AlgebraBGP{Triples: []*Triple{{Subject: "?s", Predicate: "p", Object: "o", ObjectIsVar: false}}}
	union := &AlgebraUnion{Alternatives: []AlgebraExpr{
		&AlgebraEmpty{},
		bgp,
		bgp2,
	}}

	op := NewAlgebraOptimizer(identitySimplificationRule{})
	opt := op.Optimize(union)

	switch optimized := opt.(type) {
	case *AlgebraBGP:
		t.Logf("Simplified to single BGP after dedup and empty removal")
		if !strings.Contains(optimized.String(), "BGP") {
			t.Fatalf("unexpected BGP result: %s", optimized.String())
		}
	case *AlgebraUnion:
		if len(optimized.Alternatives) != 1 {
			t.Fatalf("expected 1 alternative after deduplication, got %d", len(optimized.Alternatives))
		}
	default:
		t.Fatalf("expected union or bgp after simplification, got %T", opt)
	}
}

func TestFilterNormalization(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?s", Predicate: "p", Object: "?o", ObjectIsVar: true}}}
	filter := &AlgebraFilter{Expr: "( ( ?o > 5 ) ) ", Input: bgp}

	op := NewAlgebraOptimizer(filterNormalizationRule{})
	opt := op.Optimize(filter)

	out := opt.(*AlgebraFilter)
	if out.Expr != "?o > 5" {
		t.Fatalf("expected normalized filter, got %q", out.Expr)
	}
}

func TestPathDesugarSequence(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?s", Predicate: "p1/p2", Object: "?o", ObjectIsVar: true}}}

	op := NewAlgebraOptimizer(pathDesugarRule{})
	opt := op.Optimize(bgp)

	out, ok := opt.(*AlgebraBGP)
	if !ok {
		t.Fatalf("expected BGP after desugaring, got %T", opt)
	}
	if len(out.Triples) != 2 {
		t.Fatalf("expected 2 triples after path expansion, got %d", len(out.Triples))
	}
	if out.Triples[0].Predicate != "p1" || out.Triples[1].Predicate != "p2" {
		t.Fatalf("unexpected predicates after desugaring: %v, %v", out.Triples[0].Predicate, out.Triples[1].Predicate)
	}
	if out.Triples[0].Object != out.Triples[1].Subject {
		t.Fatalf("expected chained intermediate variable, got %s vs %s", out.Triples[0].Object, out.Triples[1].Subject)
	}
}

func TestScopeAnalysis(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{{Subject: "?s", Predicate: "p", Object: "?o", ObjectIsVar: true}}}
	opt := &AlgebraLeftJoin{
		Input:    bgp,
		Patterns: []*Triple{{Subject: "?s", Predicate: "q", Object: "?y", ObjectIsVar: true}},
	}
	proj := &AlgebraProject{Vars: []string{"?s", "?o", "?y"}, Input: opt}

	maybe := MaybeBoundVars(proj)
	def := DefinitelyBoundVars(proj)

	if !containsVar(maybe, "?y") {
		t.Fatalf("expected optional var ?y to be maybe-bound, but got: %v", maybe)
	}
	if containsVar(def, "?y") {
		t.Fatalf("optional var ?y should not be definitely bound")
	}
	if !containsVar(def, "?s") || !containsVar(def, "?o") {
		t.Fatalf("expected projected vars to remain definitely bound: %v", def)
	}
}

func containsVar(vars []string, v string) bool {
	for _, x := range vars {
		if x == v {
			return true
		}
	}
	return false
}
