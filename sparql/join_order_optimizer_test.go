package sparql

import (
	"testing"
)

func TestJoinOrderOptimizerSimpleReordering(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
			"age": {
				Name:          "age",
				PredicateType: "int",
				IndexType:     "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimized := result.(*AlgebraBGP)

	// Indexed "name" should come before unindexed "age"
	if optimized.Triples[0].Predicate != "name" {
		t.Errorf("expected 'name' first, got %s", optimized.Triples[0].Predicate)
	}

	if optimized.Triples[1].Predicate != "age" {
		t.Errorf("expected 'age' second, got %s", optimized.Triples[1].Predicate)
	}
}

func TestJoinOrderOptimizerSelectivityCalculation(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	// Test indexed predicate with bound object
	triple1 := &Triple{Subject: "?x", Predicate: "name", Object: "alice", ObjectIsVar: false}
	sel1 := joo.calculateSelectivity(triple1)

	// Test unindexed predicate with variable object
	unindexedSchema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"type": {
				Name:          "type",
				PredicateType: "string",
				IndexType:     "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}
	sa2 := NewSchemaAnalyzer(unindexedSchema)
	joo2 := NewJoinOrderOptimizer(sa2)
	triple2 := &Triple{Subject: "?x", Predicate: "type", Object: "?type", ObjectIsVar: true}
	sel2 := joo2.calculateSelectivity(triple2)

	// Indexed with bound object should be more selective
	if sel1 >= sel2 {
		t.Errorf("expected indexed bound to be more selective: %f >= %f", sel1, sel2)
	}
}

func TestJoinOrderOptimizerMultiWayJoin(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
			"age": {
				Name:          "age",
				PredicateType: "int",
				IndexType:     "range",
			},
			"email": {
				Name:          "email",
				PredicateType: "string",
				IndexType:     "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimized := result.(*AlgebraBGP)

	// Indexed predicates should come first, in order of specificity
	if optimized.Triples[0].Predicate != "name" && optimized.Triples[0].Predicate != "age" {
		t.Errorf("expected indexed predicate first, got %s", optimized.Triples[0].Predicate)
	}

	if optimized.Triples[2].Predicate != "email" {
		t.Errorf("expected unindexed 'email' last, got %s", optimized.Triples[2].Predicate)
	}
}

func TestJoinOrderOptimizerBoundVariables(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	joo := NewJoinOrderOptimizer(sa)

	// Triple with bound subject and variable object - more selective
	triple1 := &Triple{Subject: "alice", Predicate: "name", Object: "?name", ObjectIsVar: true}
	sel1 := joo.calculateSelectivity(triple1)

	// Triple with variable subject and variable object - less selective
	triple2 := &Triple{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true}
	sel2 := joo.calculateSelectivity(triple2)

	if sel1 >= sel2 {
		t.Errorf("expected bound subject to be more selective: %f >= %f", sel1, sel2)
	}
}

func TestJoinOrderOptimizerEmptyBGP(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{Triples: []*Triple{}}
	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error for empty BGP")
	}

	if result == nil {
		t.Errorf("expected non-nil result")
	}
}

func TestJoinOrderOptimizerSingleTriple(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error for single triple")
	}

	optimized := result.(*AlgebraBGP)
	if len(optimized.Triples) != 1 {
		t.Errorf("expected 1 triple, got %d", len(optimized.Triples))
	}
}

func TestJoinOrderOptimizerNilInput(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	joo := NewJoinOrderOptimizer(sa)

	result, err := joo.OptimizeJoinOrder(nil)
	if err == nil {
		t.Errorf("expected error for nil input")
	}

	if result != nil {
		t.Errorf("expected nil result for nil input")
	}
}

func TestJoinOrderOptimizerOptimizeJoin(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	leftBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	rightBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	join := &AlgebraJoin{
		Left:  leftBGP,
		Right: rightBGP,
	}

	result, err := joo.OptimizeJoinOrder(join)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result == nil {
		t.Errorf("expected non-nil result")
	}
}

func TestJoinOrderOptimizerOptimizeUnion(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	joo := NewJoinOrderOptimizer(sa)

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	result, err := joo.OptimizeJoinOrder(union)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result == nil {
		t.Errorf("expected non-nil result")
	}
}

func TestJoinOrderOptimizerOptimizeFilter(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	filter := &AlgebraFilter{
		Input: bgp,
		Expr:  "?name = \"alice\"",
	}

	result, err := joo.OptimizeJoinOrder(filter)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if result == nil {
		t.Errorf("expected non-nil result")
	}
}

func TestJoinOrderOptimizerGetStatistics(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
			"age": {
				Name:          "age",
				PredicateType: "int",
				IndexType:     "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	stats := joo.GetJoinStatistics(bgp)

	if stats == nil {
		t.Errorf("expected non-nil statistics")
	}

	if stats.TotalJoins != 1 {
		t.Errorf("expected 1 total join, got %d", stats.TotalJoins)
	}

	if stats.IndexedJoins != 1 {
		t.Errorf("expected 1 indexed join, got %d", stats.IndexedJoins)
	}
}

func TestJoinOrderOptimizerNoSchema(t *testing.T) {
	joo := NewJoinOrderOptimizer(nil)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error with nil schema")
	}

	if result == nil {
		t.Errorf("expected non-nil result")
	}
}

func TestJoinOrderOptimizerReverseIndex(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"friend": {
				Name:          "friend",
				PredicateType: "uid",
				ReverseIndex:  true,
			},
			"knows": {
				Name:          "knows",
				PredicateType: "uid",
				ReverseIndex:  false,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	joo := NewJoinOrderOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "knows", Object: "?y", ObjectIsVar: true},
			{Subject: "?x", Predicate: "friend", Object: "?z", ObjectIsVar: true},
		},
	}

	result, err := joo.OptimizeJoinOrder(bgp)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimized := result.(*AlgebraBGP)

	// Reverse-indexed "friend" should come before non-indexed "knows"
	if optimized.Triples[0].Predicate != "friend" {
		t.Errorf("expected 'friend' (reverse-indexed) first, got %s", optimized.Triples[0].Predicate)
	}
}
