package sparql

import (
	"testing"
)

func TestVariableScopeAnalyzerBasic(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bgp,
	}

	info := analyzer.Analyze(project)

	// ?x is projected
	if !info.ProjectedVars["?x"] {
		t.Errorf("expected ?x to be in projected vars")
	}

	// ?name is defined but not projected
	if !info.DefinedVars["?name"] {
		t.Errorf("expected ?name to be defined")
	}

	if !info.UnusedVars["?name"] {
		t.Errorf("expected ?name to be unused")
	}
}

func TestVariableScopeAnalyzerUsedInFilter(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

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

	info := analyzer.Analyze(project)

	// ?age is used in FILTER
	if !info.UsedVars["?age"] {
		t.Errorf("expected ?age to be used (in FILTER)")
	}

	// ?age should NOT be unused
	if info.UnusedVars["?age"] {
		t.Errorf("expected ?age to NOT be unused (it's in FILTER)")
	}
}

func TestVariableScopeAnalyzerUsedInGroupBy(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "?type", ObjectIsVar: true},
		},
	}

	agg := &AlgebraAgg{
		Op:    "COUNT",
		Var:   "?count",
		Group: []string{"?type"},
		Input: bgp,
	}

	info := analyzer.Analyze(agg)

	// ?type is used in GROUP BY
	if !info.UsedVars["?type"] {
		t.Errorf("expected ?type to be used (in GROUP BY)")
	}

	// ?type should NOT be unused
	if info.UnusedVars["?type"] {
		t.Errorf("expected ?type to NOT be unused")
	}
}

func TestVariableScopeAnalyzerMultipleUnused(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "foaf:email", Object: "?email", ObjectIsVar: true},
		},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bgp,
	}

	info := analyzer.Analyze(project)

	// ?name, ?age, ?email should all be unused
	if !info.UnusedVars["?name"] {
		t.Errorf("expected ?name to be unused")
	}
	if !info.UnusedVars["?age"] {
		t.Errorf("expected ?age to be unused")
	}
	if !info.UnusedVars["?email"] {
		t.Errorf("expected ?email to be unused")
	}

	// Should have 3 unused variables
	if len(info.UnusedVars) != 3 {
		t.Errorf("expected 3 unused variables, got %d", len(info.UnusedVars))
	}
}

func TestVariableScopeAnalyzerBind(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	// BIND creates ?unused
	bind := &AlgebraBind{
		Var:   "?unused",
		Expr:  "CONCAT(?name, ' Smith')",
		Input: bgp,
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bind,
	}

	info := analyzer.Analyze(project)

	// ?name is used in BIND expression
	if !info.UsedVars["?name"] {
		t.Errorf("expected ?name to be used (in BIND expression)")
	}

	// ?unused is defined but not used
	if !info.DefinedVars["?unused"] {
		t.Errorf("expected ?unused to be defined")
	}

	if !info.UnusedVars["?unused"] {
		t.Errorf("expected ?unused to be unused")
	}
}

func TestVariableScopeAnalyzerOrderBy(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
		},
	}

	orderBy := &AlgebraOrderBy{
		Exprs:     []string{"?name"},
		Ascending: []bool{true},
		Input:     bgp,
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: orderBy,
	}

	info := analyzer.Analyze(project)

	// ?name is used in ORDER BY
	if !info.UsedVars["?name"] {
		t.Errorf("expected ?name to be used (in ORDER BY)")
	}

	// ?age is not used
	if !info.UnusedVars["?age"] {
		t.Errorf("expected ?age to be unused")
	}
}

func TestVariableScopeAnalyzerJoin(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgpA := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	bgpB := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:age", Object: "?age", ObjectIsVar: true},
		},
	}

	join := &AlgebraJoin{
		Left:  bgpA,
		Right: bgpB,
	}

	project := &AlgebraProject{
		Vars:  []string{"?name", "?age"},
		Input: join,
	}

	info := analyzer.Analyze(project)

	// ?name and ?age are projected
	if !info.ProjectedVars["?name"] {
		t.Errorf("expected ?name to be in projected vars")
	}
	if !info.ProjectedVars["?age"] {
		t.Errorf("expected ?age to be in projected vars")
	}

	// ?x is defined but not projected (unused in SELECT)
	if !info.DefinedVars["?x"] {
		t.Errorf("expected ?x to be defined")
	}

	// In SPARQL, if ?x is not in SELECT, it's unused (even if it's a join key)
	if !info.UnusedVars["?x"] {
		t.Errorf("expected ?x to be unused (not in SELECT)")
	}
}

func TestVariableScopeAnalyzerUnion(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgpA := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person"},
		},
	}

	bgpB := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Organization"},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgpA, bgpB},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: union,
	}

	info := analyzer.Analyze(project)

	// ?x should be defined and projected
	if !info.ProjectedVars["?x"] {
		t.Errorf("expected ?x to be in projected vars")
	}

	if info.UnusedVars["?x"] {
		t.Errorf("expected ?x to be used (in SELECT)")
	}
}

func TestVariableScopeAnalyzerValues(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	values := &AlgebraValues{
		Vars: []string{"?x", "?name"},
		Rows: []map[string]string{
			{"?x": "<uri1>", "?name": "John"},
			{"?x": "<uri2>", "?name": "Jane"},
		},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: values,
	}

	info := analyzer.Analyze(project)

	// Both variables should be defined
	if !info.DefinedVars["?x"] {
		t.Errorf("expected ?x to be defined")
	}

	if !info.DefinedVars["?name"] {
		t.Errorf("expected ?name to be defined")
	}

	// ?name should be unused
	if !info.UnusedVars["?name"] {
		t.Errorf("expected ?name to be unused")
	}
}

func TestVariableScopeAnalyzerHasUnusedVariables(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: bgp,
	}

	info := analyzer.Analyze(project)

	if !info.HasUnusedVariables() {
		t.Errorf("expected to have unused variables")
	}

	unusedList := info.GetUnusedVariables()
	if len(unusedList) != 1 {
		t.Errorf("expected 1 unused variable, got %d", len(unusedList))
	}
}

func TestVariableScopeAnalyzerNoUnused(t *testing.T) {
	analyzer := NewVariableScopeAnalyzer()

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	project := &AlgebraProject{
		Vars:  []string{"?x", "?name"},
		Input: bgp,
	}

	info := analyzer.Analyze(project)

	if info.HasUnusedVariables() {
		t.Errorf("expected no unused variables")
	}

	if len(info.GetUnusedVariables()) != 0 {
		t.Errorf("expected 0 unused variables")
	}
}
