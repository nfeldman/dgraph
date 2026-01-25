package sparql

// SchemaAwareOptimizerPipeline orchestrates all Phase 3 optimizers.
// It:
// - Coordinates all schema-aware optimizers
// - Manages optimization order
// - Handles optimizer interactions
// - Provides configuration options
// - Caches optimization results
type SchemaAwareOptimizerPipeline struct {
	SchemaAnalyzer          *SchemaAnalyzer
	PredicateOptimizer      *PredicateOptimizer
	TypeConstraintAnalyzer  *TypeConstraintAnalyzer
	JoinOrderOptimizer      *JoinOrderOptimizer
	IndexSelectionOptimizer *IndexSelectionOptimizer
	StatisticsCollector     *StatisticsCollector
	Config                  *OptimizerConfig
	OptimizationResults     map[string]AlgebraExpr
}

// OptimizerConfig contains configuration for optimization.
type OptimizerConfig struct {
	EnablePredicateOptimization  bool
	EnableTypeConstraintAnalysis bool
	EnableJoinOrderOptimization  bool
	EnableIndexSelection         bool
	EnableStatisticsCollection   bool
	CacheResults                 bool
	OptimizationLevel            int // 0=none, 1=basic, 2=aggressive
}

// NewSchemaAwareOptimizerPipeline creates a new pipeline with default configuration.
func NewSchemaAwareOptimizerPipeline(schema *SchemaAnalyzer) *SchemaAwareOptimizerPipeline {
	return &SchemaAwareOptimizerPipeline{
		SchemaAnalyzer:          schema,
		PredicateOptimizer:      NewPredicateOptimizer(schema),
		TypeConstraintAnalyzer:  NewTypeConstraintAnalyzer(schema),
		JoinOrderOptimizer:      NewJoinOrderOptimizer(schema),
		IndexSelectionOptimizer: NewIndexSelectionOptimizer(schema),
		StatisticsCollector:     NewStatisticsCollector(),
		Config: &OptimizerConfig{
			EnablePredicateOptimization:  true,
			EnableTypeConstraintAnalysis: true,
			EnableJoinOrderOptimization:  true,
			EnableIndexSelection:         true,
			EnableStatisticsCollection:   true,
			CacheResults:                 true,
			OptimizationLevel:            2,
		},
		OptimizationResults: make(map[string]AlgebraExpr),
	}
}

// Optimize optimizes an algebra expression using the full pipeline.
// Returns the optimized expression and optimization metadata.
func (pipeline *SchemaAwareOptimizerPipeline) Optimize(expr AlgebraExpr) (AlgebraExpr, *OptimizationResult, error) {
	if expr == nil {
		return nil, nil, nil
	}

	result := &OptimizationResult{
		OriginalExpr:    expr.String(),
		OptimizedExpr:   "",
		StepsApplied:    make([]string, 0),
		IndexInfo:       make(map[string]*IndexInfo),
		TypeConstraints: nil,
	}

	current := expr

	// Step 1: Predicate Optimization
	if pipeline.Config.EnablePredicateOptimization && pipeline.Config.OptimizationLevel >= 1 {
		optimized, err := pipeline.PredicateOptimizer.Optimize(current)
		if err == nil && optimized != nil {
			current = optimized
			result.StepsApplied = append(result.StepsApplied, "predicate_optimization")
		}
	}

	// Step 2: Type Constraint Analysis
	if pipeline.Config.EnableTypeConstraintAnalysis && pipeline.Config.OptimizationLevel >= 1 {
		constraints := pipeline.TypeConstraintAnalyzer.Analyze(current)
		if constraints != nil && len(constraints.VariableConstraints) > 0 {
			result.TypeConstraints = constraints
			result.StepsApplied = append(result.StepsApplied, "type_constraint_analysis")
		}
	}

	// Step 3: Join Order Optimization
	if pipeline.Config.EnableJoinOrderOptimization && pipeline.Config.OptimizationLevel >= 2 {
		optimized, err := pipeline.JoinOrderOptimizer.OptimizeJoinOrder(current)
		if err == nil && optimized != nil {
			current = optimized
			result.StepsApplied = append(result.StepsApplied, "join_order_optimization")
			result.JoinStats = pipeline.JoinOrderOptimizer.GetJoinStatistics(current)
		}
	}

	// Step 4: Index Selection
	if pipeline.Config.EnableIndexSelection && pipeline.Config.OptimizationLevel >= 2 {
		indexes := pipeline.IndexSelectionOptimizer.SelectIndexes(current)
		result.IndexInfo = indexes
		result.StepsApplied = append(result.StepsApplied, "index_selection")
	}

	// Store result
	result.OptimizedExpr = current.String()

	if pipeline.Config.CacheResults {
		pipeline.OptimizationResults[result.OriginalExpr] = current
	}

	return current, result, nil
}

// SetConfig updates the optimization configuration.
func (pipeline *SchemaAwareOptimizerPipeline) SetConfig(config *OptimizerConfig) {
	if config != nil {
		pipeline.Config = config
	}
}

// EnableOptimizer enables a specific optimizer by name.
func (pipeline *SchemaAwareOptimizerPipeline) EnableOptimizer(name string, enabled bool) {
	switch name {
	case "predicate":
		pipeline.Config.EnablePredicateOptimization = enabled
	case "type_constraint":
		pipeline.Config.EnableTypeConstraintAnalysis = enabled
	case "join_order":
		pipeline.Config.EnableJoinOrderOptimization = enabled
	case "index_selection":
		pipeline.Config.EnableIndexSelection = enabled
	case "statistics":
		pipeline.Config.EnableStatisticsCollection = enabled
	}
}

// SetOptimizationLevel sets the optimization level (0=none, 1=basic, 2=aggressive).
func (pipeline *SchemaAwareOptimizerPipeline) SetOptimizationLevel(level int) {
	if level >= 0 && level <= 2 {
		pipeline.Config.OptimizationLevel = level
	}
}

// GetPipelineStatus returns the current pipeline status and configuration.
func (pipeline *SchemaAwareOptimizerPipeline) GetPipelineStatus() *PipelineStatus {
	return &PipelineStatus{
		SchemaLoaded:                  pipeline.SchemaAnalyzer != nil,
		PredicateOptimizationEnabled:  pipeline.Config.EnablePredicateOptimization,
		TypeConstraintAnalysisEnabled: pipeline.Config.EnableTypeConstraintAnalysis,
		JoinOrderOptimizationEnabled:  pipeline.Config.EnableJoinOrderOptimization,
		IndexSelectionEnabled:         pipeline.Config.EnableIndexSelection,
		StatisticsCollectionEnabled:   pipeline.Config.EnableStatisticsCollection,
		OptimizationLevel:             pipeline.Config.OptimizationLevel,
		CachedResults:                 len(pipeline.OptimizationResults),
	}
}

// OptimizationResult contains information about an optimization run.
type OptimizationResult struct {
	OriginalExpr    string
	OptimizedExpr   string
	StepsApplied    []string
	IndexInfo       map[string]*IndexInfo
	TypeConstraints *TypeConstraints
	JoinStats       *JoinStatistics
}

// PipelineStatus contains information about the pipeline status.
type PipelineStatus struct {
	SchemaLoaded                  bool
	PredicateOptimizationEnabled  bool
	TypeConstraintAnalysisEnabled bool
	JoinOrderOptimizationEnabled  bool
	IndexSelectionEnabled         bool
	StatisticsCollectionEnabled   bool
	OptimizationLevel             int
	CachedResults                 int
}

// GetStatistics returns statistics from the pipeline.
func (pipeline *SchemaAwareOptimizerPipeline) GetStatistics() *StatisticsInfo {
	return pipeline.StatisticsCollector.GetStatisticsInfo()
}

// RecordExecution records execution statistics for a predicate.
func (pipeline *SchemaAwareOptimizerPipeline) RecordExecution(predicate string, cardinality uint64) {
	if pipeline.Config.EnableStatisticsCollection {
		pipeline.StatisticsCollector.RecordExecution(predicate, cardinality)
	}
}

// ClearCache clears the optimization result cache.
func (pipeline *SchemaAwareOptimizerPipeline) ClearCache() {
	pipeline.OptimizationResults = make(map[string]AlgebraExpr)
}

// GetOptimizationInfo returns optimization recommendations for an expression.
func (pipeline *SchemaAwareOptimizerPipeline) GetOptimizationInfo(expr AlgebraExpr) *OptimizationInfo {
	if expr == nil {
		return &OptimizationInfo{
			Recommendations:  make([]*IndexRecommendation, 0),
			OptimizationGain: 0.0,
		}
	}

	recommendations := pipeline.IndexSelectionOptimizer.GetIndexRecommendations(expr)
	analysis := pipeline.IndexSelectionOptimizer.AnalyzeIndexUsage(expr)

	// Calculate potential optimization gain
	// More unindexed predicates = higher potential gain
	gain := float64(analysis.UnindexedPredicates) / float64(analysis.TotalPredicates)

	return &OptimizationInfo{
		Recommendations:     recommendations,
		OptimizationGain:    gain,
		IndexedPredicates:   analysis.IndexedPredicates,
		UnindexedPredicates: analysis.UnindexedPredicates,
		AverageIndexCost:    analysis.AverageIndexCost,
	}
}

// OptimizationInfo contains information about optimization potential.
type OptimizationInfo struct {
	Recommendations     []*IndexRecommendation
	OptimizationGain    float64
	IndexedPredicates   int
	UnindexedPredicates int
	AverageIndexCost    float64
}

// UpdateSchema updates the schema used by all optimizers.
func (pipeline *SchemaAwareOptimizerPipeline) UpdateSchema(schema *SchemaAnalyzer) {
	pipeline.SchemaAnalyzer = schema
	pipeline.PredicateOptimizer = NewPredicateOptimizer(schema)
	pipeline.TypeConstraintAnalyzer = NewTypeConstraintAnalyzer(schema)
	pipeline.JoinOrderOptimizer = NewJoinOrderOptimizer(schema)
	pipeline.IndexSelectionOptimizer = NewIndexSelectionOptimizer(schema)
}
