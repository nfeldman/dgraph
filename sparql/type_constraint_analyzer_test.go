package sparql

import (
	"testing"
)

func TestTypeConstraintAnalyzerAnalyzeBGP(t *testing.T) {
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
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
		},
	}

	constraints := tca.Analyze(bgp)

	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}

	if len(constraints.VariableConstraints) == 0 {
		t.Errorf("expected constraints for variables")
	}
}

func TestTypeConstraintAnalyzerExplicitTypeConstraint(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{},
		Types: map[string]*TypeInfo{
			"Person": {
				Name:   "Person",
				Fields: make(map[string]*PredicateInfo),
			},
		},
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
		},
	}

	constraints := tca.Analyze(bgp)

	// Check that type constraint was added
	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}

	// Get the types for ?x
	types := tca.GetVariableTypes(constraints, "?x")
	if len(types) == 0 {
		t.Errorf("expected type constraint for ?x")
	}
}

func TestTypeConstraintAnalyzerValidateConstraints(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	tca := NewTypeConstraintAnalyzer(sa)

	constraints := &TypeConstraints{
		VariableConstraints: map[string]*VariableTypeConstraint{
			"?x": {
				Variable: "?x",
				Types: map[string]bool{
					"string": true,
					"int":    true,
				},
			},
		},
	}

	valid, err := tca.ValidateConstraints(constraints)
	if !valid {
		t.Errorf("expected valid constraints, got invalid")
	}
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}
}

func TestTypeConstraintAnalyzerNilConstraints(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	tca := NewTypeConstraintAnalyzer(sa)

	valid, err := tca.ValidateConstraints(nil)
	if !valid {
		t.Errorf("expected nil constraints to be valid")
	}
	if err != nil {
		t.Errorf("expected no error for nil constraints")
	}
}

func TestTypeConstraintAnalyzerGetVariableTypes(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := &TypeConstraints{
		VariableConstraints: map[string]*VariableTypeConstraint{
			"?x": {
				Variable: "?x",
				Types: map[string]bool{
					"string": true,
					"int":    true,
				},
			},
		},
	}

	types := tca.GetVariableTypes(constraints, "?x")
	if len(types) != 2 {
		t.Errorf("expected 2 types, got %d", len(types))
	}

	// Check that both types are present
	typeSet := make(map[string]bool)
	for _, t := range types {
		typeSet[t] = true
	}

	if !typeSet["string"] || !typeSet["int"] {
		t.Errorf("expected 'string' and 'int' types")
	}
}

func TestTypeConstraintAnalyzerGetVariableTypesNotFound(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := &TypeConstraints{
		VariableConstraints: make(map[string]*VariableTypeConstraint),
	}

	types := tca.GetVariableTypes(constraints, "?x")
	if len(types) != 0 {
		t.Errorf("expected no types for non-existent variable")
	}
}

func TestTypeConstraintAnalyzerRemoveConflicting(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	tca := NewTypeConstraintAnalyzer(sa)

	constraints := &TypeConstraints{
		VariableConstraints: map[string]*VariableTypeConstraint{
			"?x": {
				Variable: "?x",
				Types: map[string]bool{
					"string": true,
				},
			},
		},
	}

	err := tca.RemoveConflictingConstraints(constraints)
	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	// Check that constraint still exists
	if _, exists := constraints.VariableConstraints["?x"]; !exists {
		t.Errorf("expected constraint for ?x to remain")
	}
}

func TestTypeConstraintAnalyzerRemoveConflictingNil(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	tca := NewTypeConstraintAnalyzer(sa)

	err := tca.RemoveConflictingConstraints(nil)
	if err == nil {
		t.Errorf("expected error for nil constraints")
	}
}

func TestTypeConstraintAnalyzerGetConstraintInfo(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := &TypeConstraints{
		VariableConstraints: map[string]*VariableTypeConstraint{
			"?x": {
				Variable: "?x",
				Types: map[string]bool{
					"string": true,
					"int":    true,
				},
			},
			"?y": {
				Variable: "?y",
				Types: map[string]bool{
					"bool": true,
				},
			},
		},
	}

	info := tca.GetConstraintInfo(constraints)

	if info.TotalVariables != 2 {
		t.Errorf("expected 2 total variables, got %d", info.TotalVariables)
	}

	if info.ConstrainedVariables != 2 {
		t.Errorf("expected 2 constrained variables, got %d", info.ConstrainedVariables)
	}

	if info.AverageTypesPerVar != 1.5 {
		t.Errorf("expected average 1.5 types per var, got %f", info.AverageTypesPerVar)
	}
}

func TestTypeConstraintAnalyzerGetConstraintInfoEmpty(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := &TypeConstraints{
		VariableConstraints: make(map[string]*VariableTypeConstraint),
	}

	info := tca.GetConstraintInfo(constraints)

	if info.TotalVariables != 0 {
		t.Errorf("expected 0 total variables, got %d", info.TotalVariables)
	}

	if info.ConstrainedVariables != 0 {
		t.Errorf("expected 0 constrained variables, got %d", info.ConstrainedVariables)
	}
}

func TestTypeConstraintAnalyzerOptimize(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := &TypeConstraints{
		VariableConstraints: map[string]*VariableTypeConstraint{
			"?x": {
				Variable: "?x",
				Types: map[string]bool{
					"string": true,
					"int":    true,
				},
				Source: "pattern",
			},
			"?empty": {
				Variable: "?empty",
				Types:    make(map[string]bool),
			},
		},
	}

	optimized := tca.OptimizeConstraints(constraints)

	if len(optimized.VariableConstraints) != 1 {
		t.Errorf("expected 1 constraint after optimization, got %d", len(optimized.VariableConstraints))
	}

	if _, exists := optimized.VariableConstraints["?x"]; !exists {
		t.Errorf("expected ?x in optimized constraints")
	}

	if _, exists := optimized.VariableConstraints["?empty"]; exists {
		t.Errorf("expected ?empty to be removed from optimized constraints")
	}
}

func TestTypeConstraintAnalyzerAnalyzeJoin(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	leftBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	rightBGP := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
		},
	}

	join := &AlgebraJoin{
		Left:  leftBGP,
		Right: rightBGP,
	}

	constraints := tca.Analyze(join)

	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}

	// Both sides should be analyzed
	if len(constraints.VariableConstraints) == 0 {
		t.Errorf("expected constraints from join")
	}
}

func TestTypeConstraintAnalyzerAnalyzeUnion(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	bgp1 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	bgp2 := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
		},
	}

	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{bgp1, bgp2},
	}

	constraints := tca.Analyze(union)

	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}
}

func TestTypeConstraintAnalyzerAnalyzeFilter(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	filter := &AlgebraFilter{
		Input: bgp,
		Expr:  "?x > 0",
	}

	constraints := tca.Analyze(filter)

	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}
}

func TestTypeConstraintAnalyzerNoSchema(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	constraints := tca.Analyze(bgp)

	if constraints == nil {
		t.Errorf("expected non-nil constraints even with nil schema")
	}

	// With no schema, we still get empty constraints structure
	if len(constraints.VariableConstraints) != 0 {
		t.Errorf("expected no constraints with nil schema")
	}
}

func TestTypeConstraintAnalyzerMultipleTypes(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
			},
			"parent": {
				Name:          "parent",
				PredicateType: "uid",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	tca := NewTypeConstraintAnalyzer(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "parent", Object: "?parent", ObjectIsVar: true},
			{Subject: "?x", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
		},
	}

	constraints := tca.Analyze(bgp)

	if constraints == nil {
		t.Errorf("expected non-nil constraints")
	}
}

func TestTypeConstraintAnalyzerNilExpression(t *testing.T) {
	tca := NewTypeConstraintAnalyzer(nil)

	constraints := tca.Analyze(nil)

	if constraints == nil {
		t.Errorf("expected non-nil constraints structure")
	}

	if len(constraints.VariableConstraints) != 0 {
		t.Errorf("expected no variables in constraints")
	}
}
