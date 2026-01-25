package sparql

import (
	"testing"
)

func TestIndexSelectionOptimizerSelectIndexes(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				Cardinality:   1000,
			},
			"age": {
				Name:          "age",
				PredicateType: "int",
				IndexType:     "range",
				Cardinality:   100,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	indexes := iso.SelectIndexes(bgp)

	if len(indexes) != 2 {
		t.Errorf("expected 2 indexes, got %d", len(indexes))
	}

	if indexes["name"] == nil {
		t.Errorf("expected index info for 'name'")
	}

	if indexes["age"] == nil {
		t.Errorf("expected index info for 'age'")
	}

	if indexes["name"].IndexType != "hash" {
		t.Errorf("expected 'hash' index for name, got %s", indexes["name"].IndexType)
	}
}

func TestIndexSelectionOptimizerGetBestIndex(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				Cardinality:   1000,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	info := iso.getBestIndex("name")

	if info == nil {
		t.Errorf("expected index info")
	}

	if info.IndexType != "hash" {
		t.Errorf("expected 'hash' index type, got %s", info.IndexType)
	}

	if !info.Recommended {
		t.Errorf("expected indexed predicate to be recommended")
	}
}

func TestIndexSelectionOptimizerCalculateIndexCost(t *testing.T) {
	iso := NewIndexSelectionOptimizer(nil)

	// Indexed predicate should have lower cost
	indexedInfo := &IndexInfo{
		Predicate:   "name",
		IndexType:   "hash",
		Selectivity: 0.3,
	}

	// Unindexed predicate should have higher cost
	unindexedInfo := &IndexInfo{
		Predicate:   "unknown",
		IndexType:   "",
		Selectivity: 0.7,
	}

	indexedCost := iso.calculateIndexCost(indexedInfo)
	unindexedCost := iso.calculateIndexCost(unindexedInfo)

	if indexedCost >= unindexedCost {
		t.Errorf("expected indexed to have lower cost: %f >= %f", indexedCost, unindexedCost)
	}
}

func TestIndexSelectionOptimizerGetIndexRecommendations(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				Cardinality:   1000,
			},
			"email": {
				Name:          "email",
				PredicateType: "string",
				IndexType:     "",
				Cardinality:   500,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
		},
	}

	recommendations := iso.GetIndexRecommendations(bgp)

	if len(recommendations) != 1 {
		t.Errorf("expected 1 recommendation, got %d", len(recommendations))
	}

	if recommendations[0].Predicate != "email" {
		t.Errorf("expected recommendation for 'email', got %s", recommendations[0].Predicate)
	}

	if recommendations[0].RecommendType != "hash" {
		t.Errorf("expected 'hash' recommendation, got %s", recommendations[0].RecommendType)
	}
}

func TestIndexSelectionOptimizerAnalyzeIndexUsage(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				Cardinality:   1000,
			},
			"age": {
				Name:          "age",
				PredicateType: "int",
				IndexType:     "range",
				Cardinality:   100,
			},
			"email": {
				Name:          "email",
				PredicateType: "string",
				IndexType:     "",
				Cardinality:   500,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
		},
	}

	analysis := iso.AnalyzeIndexUsage(bgp)

	if analysis.TotalPredicates != 3 {
		t.Errorf("expected 3 total predicates, got %d", analysis.TotalPredicates)
	}

	if analysis.IndexedPredicates != 2 {
		t.Errorf("expected 2 indexed predicates, got %d", analysis.IndexedPredicates)
	}

	if analysis.UnindexedPredicates != 1 {
		t.Errorf("expected 1 unindexed predicate, got %d", analysis.UnindexedPredicates)
	}
}

func TestIndexSelectionOptimizerNilExpression(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	iso := NewIndexSelectionOptimizer(sa)

	indexes := iso.SelectIndexes(nil)

	if len(indexes) != 0 {
		t.Errorf("expected empty indexes for nil expression")
	}
}

func TestIndexSelectionOptimizerEmptyBGP(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	iso := NewIndexSelectionOptimizer(sa)

	bgp := &AlgebraBGP{Triples: []*Triple{}}

	indexes := iso.SelectIndexes(bgp)

	if len(indexes) != 0 {
		t.Errorf("expected empty indexes for empty BGP")
	}
}

func TestIndexSelectionOptimizerIsIndexedPredicate(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
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
	iso := NewIndexSelectionOptimizer(sa)

	if !iso.IsIndexedPredicate("name") {
		t.Errorf("expected 'name' to be indexed")
	}

	if iso.IsIndexedPredicate("email") {
		t.Errorf("expected 'email' to not be indexed")
	}

	if iso.IsIndexedPredicate("unknown") {
		t.Errorf("expected 'unknown' to not be indexed")
	}
}

func TestIndexSelectionOptimizerGetIndexCost(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				Cardinality:   1000,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	cost := iso.GetIndexCost("name")

	if cost <= 0 {
		t.Errorf("expected positive cost, got %f", cost)
	}

	if cost >= 1.0 {
		t.Errorf("expected indexed predicate to have cost < 1.0, got %f", cost)
	}
}

func TestIndexSelectionOptimizerMultipleIndexTypes(t *testing.T) {
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
			"bio": {
				Name:          "bio",
				PredicateType: "string",
				IndexType:     "fulltext",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	iso := NewIndexSelectionOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "bio", Object: "?bio", ObjectIsVar: true},
		},
	}

	indexes := iso.SelectIndexes(bgp)

	if len(indexes) != 3 {
		t.Errorf("expected 3 indexes, got %d", len(indexes))
	}

	if indexes["name"].IndexType != "hash" {
		t.Errorf("expected 'hash' for name")
	}

	if indexes["age"].IndexType != "range" {
		t.Errorf("expected 'range' for age")
	}

	if indexes["bio"].IndexType != "fulltext" {
		t.Errorf("expected 'fulltext' for bio")
	}
}

func TestIndexSelectionOptimizerNoSchema(t *testing.T) {
	iso := NewIndexSelectionOptimizer(nil)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	indexes := iso.SelectIndexes(bgp)

	if len(indexes) != 1 {
		t.Errorf("expected 1 index even with nil schema")
	}

	if indexes["name"].Recommended {
		t.Errorf("expected unindexed without schema")
	}
}

func TestIndexSelectionOptimizerJoinAnalysis(t *testing.T) {
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
	iso := NewIndexSelectionOptimizer(sa)

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

	indexes := iso.SelectIndexes(join)

	if len(indexes) != 2 {
		t.Errorf("expected 2 indexes from join")
	}
}
