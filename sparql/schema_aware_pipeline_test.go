package sparql

import (
	"testing"
)

func TestSchemaAwareOptimizerPipelineCreate(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	if pipeline == nil {
		t.Errorf("expected non-nil pipeline")
	}

	if pipeline.Config == nil {
		t.Errorf("expected non-nil config")
	}

	if !pipeline.Config.EnablePredicateOptimization {
		t.Errorf("expected predicate optimization to be enabled")
	}
}

func TestSchemaAwareOptimizerPipelineOptimize(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	optimized, result, err := pipeline.Optimize(bgp)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if optimized == nil {
		t.Errorf("expected optimized expression")
	}

	if result == nil {
		t.Errorf("expected optimization result")
	}

	if len(result.StepsApplied) == 0 {
		t.Errorf("expected steps to be applied")
	}
}

func TestSchemaAwareOptimizerPipelineOptimizationResult(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	_, result, _ := pipeline.Optimize(bgp)

	if result.OriginalExpr == "" {
		t.Errorf("expected original expression")
	}

	if result.OptimizedExpr == "" {
		t.Errorf("expected optimized expression")
	}

	if len(result.StepsApplied) == 0 {
		t.Errorf("expected steps applied")
	}
}

func TestSchemaAwareOptimizerPipelineDisableOptimizers(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	// Disable all optimizers
	pipeline.EnableOptimizer("predicate", false)
	pipeline.EnableOptimizer("type_constraint", false)
	pipeline.EnableOptimizer("join_order", false)
	pipeline.EnableOptimizer("index_selection", false)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	_, result, _ := pipeline.Optimize(bgp)

	if len(result.StepsApplied) != 0 {
		t.Errorf("expected no steps when all disabled, got %d", len(result.StepsApplied))
	}
}

func TestSchemaAwareOptimizerPipelineOptimizationLevel(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	// Set to basic optimization
	pipeline.SetOptimizationLevel(1)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	_, result, _ := pipeline.Optimize(bgp)

	// Basic level should run some optimizers but not all
	if len(result.StepsApplied) == 0 {
		t.Errorf("expected some optimization steps at level 1")
	}
}

func TestSchemaAwareOptimizerPipelineGetStatus(t *testing.T) {
	schema := &SchemaInfo{
		Predicates: make(map[string]*PredicateInfo),
		Types:      make(map[string]*TypeInfo),
	}

	sa := NewSchemaAnalyzer(schema)
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	status := pipeline.GetPipelineStatus()

	if !status.SchemaLoaded {
		t.Errorf("expected schema to be loaded")
	}

	if !status.PredicateOptimizationEnabled {
		t.Errorf("expected predicate optimization to be enabled")
	}

	if status.OptimizationLevel != 2 {
		t.Errorf("expected level 2, got %d", status.OptimizationLevel)
	}
}

func TestSchemaAwareOptimizerPipelineCache(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	pipeline.Optimize(bgp)

	status := pipeline.GetPipelineStatus()
	if status.CachedResults != 1 {
		t.Errorf("expected 1 cached result, got %d", status.CachedResults)
	}

	pipeline.ClearCache()
	status = pipeline.GetPipelineStatus()
	if status.CachedResults != 0 {
		t.Errorf("expected cache to be cleared")
	}
}

func TestSchemaAwareOptimizerPipelineRecordExecution(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	pipeline.RecordExecution("name", 100)
	pipeline.RecordExecution("age", 50)

	stats := pipeline.GetStatistics()

	if stats.TrackedPredicates != 2 {
		t.Errorf("expected 2 tracked predicates, got %d", stats.TrackedPredicates)
	}
}

func TestSchemaAwareOptimizerPipelineGetOptimizationInfo(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
		},
	}

	info := pipeline.GetOptimizationInfo(bgp)

	if info == nil {
		t.Errorf("expected optimization info")
	}

	if info.UnindexedPredicates != 1 {
		t.Errorf("expected 1 unindexed predicate, got %d", info.UnindexedPredicates)
	}

	if len(info.Recommendations) != 1 {
		t.Errorf("expected 1 recommendation, got %d", len(info.Recommendations))
	}
}

func TestSchemaAwareOptimizerPipelineUpdateSchema(t *testing.T) {
	schema1 := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"name": {
				Name:          "name",
				PredicateType: "string",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa1 := NewSchemaAnalyzer(schema1)
	pipeline := NewSchemaAwareOptimizerPipeline(sa1)

	// Update schema
	schema2 := &SchemaInfo{
		Predicates: map[string]*PredicateInfo{
			"email": {
				Name:          "email",
				PredicateType: "string",
				IndexType:     "hash",
			},
		},
		Types: make(map[string]*TypeInfo),
	}

	sa2 := NewSchemaAnalyzer(schema2)
	pipeline.UpdateSchema(sa2)

	// Check that schema was updated
	status := pipeline.GetPipelineStatus()
	if !status.SchemaLoaded {
		t.Errorf("expected updated schema to be loaded")
	}
}

func TestSchemaAwareOptimizerPipelineNilExpression(t *testing.T) {
	sa := NewSchemaAnalyzer(nil)
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	optimized, result, err := pipeline.Optimize(nil)

	if optimized != nil {
		t.Errorf("expected nil for nil expression")
	}

	if result != nil {
		t.Errorf("expected nil result for nil expression")
	}

	if err != nil {
		t.Errorf("expected no error for nil expression")
	}
}

func TestSchemaAwareOptimizerPipelineNoSchema(t *testing.T) {
	pipeline := NewSchemaAwareOptimizerPipeline(nil)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	optimized, result, err := pipeline.Optimize(bgp)

	if err != nil {
		t.Errorf("expected no error with nil schema")
	}

	if optimized == nil {
		t.Errorf("expected optimization to work without schema")
	}

	if result == nil {
		t.Errorf("expected result")
	}
}

func TestSchemaAwareOptimizerPipelineIntegrationFull(t *testing.T) {
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
	pipeline := NewSchemaAwareOptimizerPipeline(sa)

	// Set aggressive optimization
	pipeline.SetOptimizationLevel(2)

	bgp := &AlgebraBGP{
		Triples: []*Triple{
			{Subject: "?x", Predicate: "email", Object: "?email", ObjectIsVar: true},
			{Subject: "?x", Predicate: "age", Object: "?age", ObjectIsVar: true},
			{Subject: "?x", Predicate: "name", Object: "?name", ObjectIsVar: true},
		},
	}

	// Run optimization
	optimized, result, err := pipeline.Optimize(bgp)

	if err != nil {
		t.Errorf("expected no error, got %v", err)
	}

	if optimized == nil {
		t.Errorf("expected optimized expression")
	}

	if result == nil {
		t.Errorf("expected optimization result")
	}

	// Should apply multiple optimizers
	if len(result.StepsApplied) < 2 {
		t.Errorf("expected multiple optimization steps, got %d", len(result.StepsApplied))
	}

	// Should have index info
	if len(result.IndexInfo) == 0 {
		t.Errorf("expected index information")
	}

	// Record execution
	pipeline.RecordExecution("name", 1000)
	stats := pipeline.GetStatistics()

	if stats.TotalExecutions < 1 {
		t.Errorf("expected execution to be recorded")
	}

	// Get optimization info
	info := pipeline.GetOptimizationInfo(bgp)
	if info == nil {
		t.Errorf("expected optimization info")
	}
}
