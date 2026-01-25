package sparql

import (
	"testing"
)

func TestSchemaAnalyzerGetPredicate(t *testing.T) {
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

	// Test get existing predicate
	predInfo := sa.GetPredicate("name")
	if predInfo == nil {
		t.Errorf("expected predicate info for 'name', got nil")
	}
	if predInfo.Name != "name" {
		t.Errorf("expected name 'name', got %s", predInfo.Name)
	}
	if predInfo.IndexType != "hash" {
		t.Errorf("expected index type 'hash', got %s", predInfo.IndexType)
	}

	// Test get non-existing predicate
	predInfo = sa.GetPredicate("nonexistent")
	if predInfo != nil {
		t.Errorf("expected nil for non-existent predicate, got %v", predInfo)
	}

	// Test nil schema
	sa = NewSchemaAnalyzer(nil)
	predInfo = sa.GetPredicate("name")
	if predInfo != nil {
		t.Errorf("expected nil from nil schema, got %v", predInfo)
	}
}

func TestSchemaAnalyzerHasPredicate(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {Name: "name"},
			"age":  {Name: "age"},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	if !sa.HasPredicate("name") {
		t.Errorf("expected HasPredicate('name') to be true")
	}

	if sa.HasPredicate("nonexistent") {
		t.Errorf("expected HasPredicate('nonexistent') to be false")
	}
}

func TestSchemaAnalyzerIsIndexed(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:      "name",
				IndexType: "hash",
			},
			"age": {
				Name:      "age",
				IndexType: "range",
			},
			"email": {
				Name:      "email",
				IndexType: "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	// Test hash index
	if !sa.IsIndexed("name", "hash") {
		t.Errorf("expected IsIndexed('name', 'hash') to be true")
	}

	// Test range index
	if !sa.IsIndexed("age", "range") {
		t.Errorf("expected IsIndexed('age', 'range') to be true")
	}

	// Test wrong index type
	if sa.IsIndexed("name", "range") {
		t.Errorf("expected IsIndexed('name', 'range') to be false")
	}

	// Test unindexed predicate
	if sa.IsIndexed("email", "hash") {
		t.Errorf("expected IsIndexed('email', 'hash') to be false")
	}

	// Test any index
	if !sa.IsIndexed("name", "") {
		t.Errorf("expected IsIndexed('name', '') to be true")
	}

	if sa.IsIndexed("email", "") {
		t.Errorf("expected IsIndexed('email', '') to be false")
	}
}

func TestSchemaAnalyzerGetCardinality(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {Name: "name", Cardinality: 1000},
			"age":  {Name: "age", Cardinality: 100},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	if sa.GetCardinality("name") != 1000 {
		t.Errorf("expected cardinality 1000 for 'name'")
	}

	if sa.GetCardinality("age") != 100 {
		t.Errorf("expected cardinality 100 for 'age'")
	}

	if sa.GetCardinality("nonexistent") != 0 {
		t.Errorf("expected cardinality 0 for non-existent predicate")
	}
}

func TestSchemaAnalyzerGetType(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: make(map[string]*PredicateInfo),
		Types: map[string]*TypeInfo{
			"Person": {
				Name: "Person",
				Fields: map[string]*PredicateInfo{
					"name": {Name: "name"},
				},
			},
		},
	}

	sa := NewSchemaAnalyzer(schema)

	typeInfo := sa.GetType("Person")
	if typeInfo == nil {
		t.Errorf("expected type info for 'Person', got nil")
	}
	if typeInfo.Name != "Person" {
		t.Errorf("expected type name 'Person', got %s", typeInfo.Name)
	}

	typeInfo = sa.GetType("NonExistent")
	if typeInfo != nil {
		t.Errorf("expected nil for non-existent type, got %v", typeInfo)
	}
}

func TestSchemaAnalyzerHasType(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: make(map[string]*PredicateInfo),
		Types: map[string]*TypeInfo{
			"Person":   {Name: "Person"},
			"Company":  {Name: "Company"},
			"Location": {Name: "Location"},
		},
	}

	sa := NewSchemaAnalyzer(schema)

	if !sa.HasType("Person") {
		t.Errorf("expected HasType('Person') to be true")
	}

	if sa.HasType("NonExistent") {
		t.Errorf("expected HasType('NonExistent') to be false")
	}
}

func TestSchemaAnalyzerGetTypes(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: make(map[string]*PredicateInfo),
		Types: map[string]*TypeInfo{
			"Person":   {Name: "Person"},
			"Company":  {Name: "Company"},
			"Location": {Name: "Location"},
		},
	}

	sa := NewSchemaAnalyzer(schema)

	types := sa.GetTypes()
	if len(types) != 3 {
		t.Errorf("expected 3 types, got %d", len(types))
	}

	// Check all types are present
	typeMap := make(map[string]bool)
	for _, t := range types {
		typeMap[t] = true
	}

	if !typeMap["Person"] || !typeMap["Company"] || !typeMap["Location"] {
		t.Errorf("expected Person, Company, Location types")
	}
}

func TestSchemaAnalyzerGetPredicates(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name":  {Name: "name"},
			"age":   {Name: "age"},
			"email": {Name: "email"},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	predicates := sa.GetPredicates()
	if len(predicates) != 3 {
		t.Errorf("expected 3 predicates, got %d", len(predicates))
	}

	// Check all predicates are present
	predMap := make(map[string]bool)
	for _, p := range predicates {
		predMap[p] = true
	}

	if !predMap["name"] || !predMap["age"] || !predMap["email"] {
		t.Errorf("expected name, age, email predicates")
	}
}

func TestSchemaAnalyzerHasReverseIndex(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"parent": {
				Name:         "parent",
				ReverseIndex: true,
			},
			"name": {
				Name:         "name",
				ReverseIndex: false,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	if !sa.HasReverseIndex("parent") {
		t.Errorf("expected HasReverseIndex('parent') to be true")
	}

	if sa.HasReverseIndex("name") {
		t.Errorf("expected HasReverseIndex('name') to be false")
	}
}

func TestSchemaAnalyzerGetSelectivity(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"indexed": {
				Name:      "indexed",
				IndexType: "hash",
			},
			"unindexed": {
				Name:      "unindexed",
				IndexType: "",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	// Indexed predicates should have higher selectivity (lower value)
	indexedSelectivity := sa.GetSelectivity("indexed")
	if indexedSelectivity != 0.3 {
		t.Errorf("expected selectivity 0.3 for indexed, got %f", indexedSelectivity)
	}

	// Unindexed predicates should have lower selectivity (higher value)
	unindexedSelectivity := sa.GetSelectivity("unindexed")
	if unindexedSelectivity != 0.7 {
		t.Errorf("expected selectivity 0.7 for unindexed, got %f", unindexedSelectivity)
	}

	// Non-existent predicates should get default selectivity
	defaultSelectivity := sa.GetSelectivity("nonexistent")
	if defaultSelectivity != 0.5 {
		t.Errorf("expected default selectivity 0.5, got %f", defaultSelectivity)
	}
}

func TestSchemaAnalyzerAddPredicate(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)

	newPred := &PredicateInfo{
		Name:      "newpred",
		IndexType: "hash",
	}

	sa.AddPredicate(newPred)

	if !sa.HasPredicate("newpred") {
		t.Errorf("expected predicate 'newpred' to be added")
	}

	if sa.GetPredicate("newpred").IndexType != "hash" {
		t.Errorf("expected index type 'hash'")
	}
}

func TestSchemaAnalyzerAddType(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)

	newType := &TypeInfo{
		Name: "NewType",
		Fields: map[string]*PredicateInfo{
			"field1": {Name: "field1"},
		},
		Predicates: []string{"field1"},
	}

	sa.AddType(newType)

	if !sa.HasType("NewType") {
		t.Errorf("expected type 'NewType' to be added")
	}
}

func TestSchemaAnalyzerUpdateCardinality(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {Name: "name", Cardinality: 100},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	sa.UpdateCardinality("name", 500)

	if sa.GetCardinality("name") != 500 {
		t.Errorf("expected cardinality 500, got %d", sa.GetCardinality("name"))
	}
}

func TestSchemaAnalyzerCopy(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {Name: "name", Cardinality: 100},
		},
		Types: map[string]*TypeInfo{
			"Person": {Name: "Person"},
		},
	}

	sa := NewSchemaAnalyzer(schema)
	saCopy := sa.Copy()

	// Verify copy has same data
	if !saCopy.HasPredicate("name") {
		t.Errorf("expected copied predicate 'name'")
	}

	if !saCopy.HasType("Person") {
		t.Errorf("expected copied type 'Person'")
	}

	// Modify copy
	saCopy.UpdateCardinality("name", 500)

	// Original should be unchanged
	if sa.GetCardinality("name") != 100 {
		t.Errorf("expected original cardinality to be 100, got %d", sa.GetCardinality("name"))
	}
}

func TestSchemaAnalyzerIsEmpty(t *testing.T) {
	// Empty schema
	sa := NewSchemaAnalyzer(nil)
	if !sa.IsEmpty() {
		t.Errorf("expected empty schema")
	}

	// Non-empty schema
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {Name: "name"},
		},
		Types: make(map[string]*TypeInfo),
	}
	sa = NewSchemaAnalyzer(schema)
	if sa.IsEmpty() {
		t.Errorf("expected non-empty schema")
	}
}

func TestSchemaAnalyzerPredicateAndTypeCount(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name":  {Name: "name"},
			"age":   {Name: "age"},
			"email": {Name: "email"},
		},
		Types: map[string]*TypeInfo{
			"Person":   {Name: "Person"},
			"Company":  {Name: "Company"},
			"Location": {Name: "Location"},
			"Post":     {Name: "Post"},
		},
	}

	sa := NewSchemaAnalyzer(schema)

	if sa.PredicateCount() != 3 {
		t.Errorf("expected 3 predicates, got %d", sa.PredicateCount())
	}

	if sa.TypeCount() != 4 {
		t.Errorf("expected 4 types, got %d", sa.TypeCount())
	}
}

func TestSchemaAnalyzerGetIndexedAndUnindexedPredicates(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name":    {Name: "name", IndexType: "hash"},
			"age":     {Name: "age", IndexType: "range"},
			"email":   {Name: "email", IndexType: ""},
			"address": {Name: "address", IndexType: ""},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	indexed := sa.GetIndexedPredicates()
	if len(indexed) != 2 {
		t.Errorf("expected 2 indexed predicates, got %d", len(indexed))
	}

	unindexed := sa.GetUnindexedPredicates()
	if len(unindexed) != 2 {
		t.Errorf("expected 2 unindexed predicates, got %d", len(unindexed))
	}
}

func TestSchemaAnalyzerGetPredicatesWithReverseIndex(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"parent": {Name: "parent", ReverseIndex: true},
			"friend": {Name: "friend", ReverseIndex: true},
			"name":   {Name: "name", ReverseIndex: false},
			"age":    {Name: "age", ReverseIndex: false},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	reverse := sa.GetPredicatesWithReverseIndex()
	if len(reverse) != 2 {
		t.Errorf("expected 2 predicates with reverse index, got %d", len(reverse))
	}

	// Verify correct predicates
	reverseMap := make(map[string]bool)
	for _, p := range reverse {
		reverseMap[p] = true
	}

	if !reverseMap["parent"] || !reverseMap["friend"] {
		t.Errorf("expected parent and friend predicates")
	}
}

func TestSchemaAnalyzerPredicateProperties(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				ListType:      false,
				Lang:          false,
				Count:         true,
			},
			"titles": {
				Name:          "titles",
				PredicateType: "string",
				ListType:      true,
				Lang:          true,
				Count:         false,
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	if sa.GetPredicateType("name") != "string" {
		t.Errorf("expected predicate type 'string'")
	}

	if sa.IsListType("name") {
		t.Errorf("expected 'name' to not be list type")
	}

	if !sa.IsListType("titles") {
		t.Errorf("expected 'titles' to be list type")
	}

	if sa.IsLanguageSupported("name") {
		t.Errorf("expected 'name' to not support language")
	}

	if !sa.IsLanguageSupported("titles") {
		t.Errorf("expected 'titles' to support language")
	}

	if !sa.IsCountSupported("name") {
		t.Errorf("expected 'name' to support count")
	}

	if sa.IsCountSupported("titles") {
		t.Errorf("expected 'titles' to not support count")
	}
}

func TestSchemaAnalyzerGetTypeFields(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: make(map[string]*PredicateInfo),
		Types: map[string]*TypeInfo{
			"Person": {
				Name: "Person",
				Fields: map[string]*PredicateInfo{
					"name":  {Name: "name"},
					"email": {Name: "email"},
					"age":   {Name: "age"},
				},
				Predicates: []string{"name", "email", "age"},
			},
		},
	}

	sa := NewSchemaAnalyzer(schema)

	fields := sa.GetTypeFields("Person")
	if len(fields) != 3 {
		t.Errorf("expected 3 fields, got %d", len(fields))
	}

	// Fields for non-existent type
	fields = sa.GetTypeFields("NonExistent")
	if len(fields) != 0 {
		t.Errorf("expected 0 fields for non-existent type, got %d", len(fields))
	}
}

func TestSchemaAnalyzerHasCardinalityEstimate(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name":   {Name: "name", Cardinality: 1000},
			"id":     {Name: "id", Cardinality: 0},
			"unused": {Name: "unused"},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)

	if !sa.HasCardinalityEstimate("name") {
		t.Errorf("expected cardinality estimate for 'name'")
	}

	if sa.HasCardinalityEstimate("id") {
		t.Errorf("expected no cardinality estimate for 'id' with 0 cardinality")
	}

	if sa.HasCardinalityEstimate("nonexistent") {
		t.Errorf("expected no cardinality estimate for non-existent predicate")
	}
}
