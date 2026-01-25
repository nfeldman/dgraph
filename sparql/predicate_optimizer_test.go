package sparql

import (
	"testing"
)

func TestPredicateOptimizerIdentifyIndexedPredicates(t *testing.T) {
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
	po := NewPredicateOptimizer(sa)

	// Test indexed predicate
	if !po.isIndexed("name") {
		t.Errorf("expected 'name' to be indexed")
	}

	// Test another indexed predicate
	if !po.isIndexed("age") {
		t.Errorf("expected 'age' to be indexed")
	}

	// Test unindexed predicate
	if po.isIndexed("email") {
		t.Errorf("expected 'email' to not be indexed")
	}

	// Test non-existent predicate
	if po.isIndexed("nonexistent") {
		t.Errorf("expected 'nonexistent' to not be indexed")
	}
}

func TestPredicateOptimizerReverseIndex(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"friend": {
				Name:          "friend",
				PredicateType: "uid",
				ReverseIndex:  true,
				Cardinality:   5000,
			},
			"parent": {
				Name:          "parent",
				PredicateType: "uid",
				ReverseIndex:  false,
				Cardinality:   1000,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	po := NewPredicateOptimizer(sa)

	// Test reverse indexed predicate
	if !po.hasReverseIndex("friend") {
		t.Errorf("expected 'friend' to have reverse index")
	}

	// Test predicate without reverse index
	if po.hasReverseIndex("parent") {
		t.Errorf("expected 'parent' to not have reverse index")
	}

	// Test non-existent predicate
	if po.hasReverseIndex("nonexistent") {
		t.Errorf("expected 'nonexistent' to not have reverse index")
	}
}

func TestPredicateOptimizerOptimizeBGP(t *testing.T) {
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
			"friend": {
				Name:          "friend",
				PredicateType: "uid",
				ReverseIndex:  true,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	po := NewPredicateOptimizer(sa)

	// Create BGP with mixed indexed and unindexed predicates
	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "friend", Object: "?friend", ObjectIsVar: true},
		},
	}

	result, err := po.Optimize(bgp)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimizedBGP, ok := result.(*AlgebraBGP)
	if !ok {
		t.Errorf("expected AlgebraBGP result")
	}

	// Check reordering: indexed first, then reverse, then unindexed
	if len(optimizedBGP.Triples) != 3 {
		t.Errorf("expected 3 triples, got %d", len(optimizedBGP.Triples))
	}

	// First should be indexed (name)
	if optimizedBGP.Triples[0].Predicate != "name" {
		t.Errorf("expected first triple to have 'name', got %s", optimizedBGP.Triples[0].Predicate)
	}

	// Second should be reverse indexed (friend)
	if optimizedBGP.Triples[1].Predicate != "friend" {
		t.Errorf("expected second triple to have 'friend', got %s", optimizedBGP.Triples[1].Predicate)
	}

	// Third should be unindexed (age)
	if optimizedBGP.Triples[2].Predicate != "age" {
		t.Errorf("expected third triple to have 'age', got %s", optimizedBGP.Triples[2].Predicate)
	}
}

func TestPredicateOptimizerMultipleIndexed(t *testing.T) {
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
	po := NewPredicateOptimizer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	result, err := po.Optimize(bgp)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimizedBGP := result.(*AlgebraBGP)

	// Both indexed predicates should come first
	if optimizedBGP.Triples[0].Predicate != "email" && optimizedBGP.Triples[0].Predicate != "name" {
		t.Errorf("expected first indexed predicate, got %s", optimizedBGP.Triples[0].Predicate)
	}

	if optimizedBGP.Triples[1].Predicate != "email" && optimizedBGP.Triples[1].Predicate != "name" {
		t.Errorf("expected second indexed predicate, got %s", optimizedBGP.Triples[1].Predicate)
	}

	// Unindexed should be last
	if optimizedBGP.Triples[2].Predicate != "age" {
		t.Errorf("expected unindexed predicate last, got %s", optimizedBGP.Triples[2].Predicate)
	}
}

func TestPredicateOptimizerEmptyBGP(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	po := NewPredicateOptimizer(sa)

	// Test empty BGP
	bgp := &AlgebraBGP{Triples: []*Triple{}}
	result, err := po.Optimize(bgp)
	if err != nil {
		t.Errorf("expected no error for empty BGP, got %v", err)
	}

	if result == nil {
		t.Errorf("expected non-nil result for empty BGP")
	}
}

func TestPredicateOptimizerNilInput(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	po := NewPredicateOptimizer(sa)

	result, err := po.Optimize(nil)
	if err == nil {
		t.Errorf("expected error for nil input")
	}
	if result != nil {
		t.Errorf("expected nil result for nil input")
	}
}

func TestPredicateOptimizerOptimizeJoin(t *testing.T) {
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
	po := NewPredicateOptimizer(sa)

	// Create nested join structure
	leftBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	rightBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	join := &AlgebraJoin{
		Left:  leftBGP,
		Right: rightBGP,
	}

	result, err := po.Optimize(join)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimizedJoin, ok := result.(*AlgebraJoin)
	if !ok {
		t.Errorf("expected AlgebraJoin result")
	}

	// Check that both sides are optimized
	if optimizedJoin.Left == nil || optimizedJoin.Right == nil {
		t.Errorf("expected both join sides to be non-nil")
	}
}

func TestPredicateOptimizerOptimizeUnion(t *testing.T) {
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
	po := NewPredicateOptimizer(sa)

	// Create union with BGPs
	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?y", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	result, err := po.Optimize(union)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimizedUnion, ok := result.(*AlgebraUnion)
	if !ok {
		t.Errorf("expected AlgebraUnion result")
	}

	if len(optimizedUnion.Alternatives) != 2 {
		t.Errorf("expected 2 union expressions, got %d", len(optimizedUnion.Alternatives))
	}
}

func TestPredicateOptimizerOptimizeFilter(t *testing.T) {
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
	po := NewPredicateOptimizer(sa)

	// Create filter with BGP
	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	filter := &AlgebraFilter{
		Input: bgp,
		Expr:  "?age > 18",
	}

	result, err := po.Optimize(filter)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	optimizedFilter, ok := result.(*AlgebraFilter)
	if !ok {
		t.Errorf("expected AlgebraFilter result")
	}

	// Check that inner BGP is optimized
	if optimizedFilter.Input == nil {
		t.Errorf("expected non-nil inner expression")
	}

	innerBGP, ok := optimizedFilter.Input.(*AlgebraBGP)
	if !ok {
		t.Errorf("expected AlgebraBGP as inner expression")
	}

	// Check reordering happened
	if innerBGP.Triples[0].Predicate != "name" {
		t.Errorf("expected 'name' first after optimization, got %s", innerBGP.Triples[0].Predicate)
	}
}

func TestPredicateOptimizerGetOptimizationInfo(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
				ReverseIndex:  false,
				Cardinality:   1000,
			},
			"friend": {
				Name:          "friend",
				PredicateType: "uid",
				IndexType:     "",
				ReverseIndex:  true,
				Cardinality:   5000,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	po := NewPredicateOptimizer(sa)

	// Test indexed predicate info
	info := po.GetOptimizationInfo("name")
	if !info.IsIndexed {
		t.Errorf("expected 'name' to be indexed in info")
	}
	if info.IndexType != "hash" {
		t.Errorf("expected index type 'hash', got %s", info.IndexType)
	}
	if info.Selectivity != 0.3 {
		t.Errorf("expected selectivity 0.3 for indexed, got %f", info.Selectivity)
	}

	// Test reverse indexed predicate info
	info = po.GetOptimizationInfo("friend")
	if info.HasReverse != true {
		t.Errorf("expected 'friend' to have reverse index")
	}

	// Test non-existent predicate info
	info = po.GetOptimizationInfo("nonexistent")
	if info.IsIndexed {
		t.Errorf("expected 'nonexistent' to not be indexed")
	}
	if info.Selectivity != 0.5 {
		t.Errorf("expected selectivity 0.5 for unknown, got %f", info.Selectivity)
	}
}

func TestPredicateOptimizerNoSchema(t *testing.T) {
	po := NewPredicateOptimizer(nil)

	// Test with nil schema
	if po.isIndexed("name") {
		t.Errorf("expected no indexed predicates with nil schema")
	}

	if po.hasReverseIndex("name") {
		t.Errorf("expected no reverse indexes with nil schema")
	}

	// Test optimization info with nil schema
	info := po.GetOptimizationInfo("name")
	if info.IsIndexed || info.HasReverse {
		t.Errorf("expected no optimization info with nil schema")
	}
}

func TestPredicateOptimizerBuiltinPredicates(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"rdf:type": {
				Name:          "rdf:type",
				PredicateType: "uid",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	po := NewPredicateOptimizer(sa)

	// Built-in predicates should not be indexed in optimization context
	if po.isIndexed("rdf:type") {
		t.Errorf("expected built-in predicate 'rdf:type' to not be indexed in optimization")
	}

	if po.isIndexed("a") {
		t.Errorf("expected built-in predicate 'a' to not be indexed in optimization")
	}
}
