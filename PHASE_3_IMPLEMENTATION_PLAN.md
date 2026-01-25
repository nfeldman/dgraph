# Phase 3: Schema-Aware Optimization & Integration

**Status**: Starting (January 25, 2026)  
**Objective**: Integrate schema information into SPARQL query optimization  
**Duration**: 2-3 weeks  
**Priority**: HIGH - Enables statistics-based optimization

---

## Executive Summary

Phase 3 extends the SPARQL optimization pipeline with **schema-aware optimizations** that leverage
Dgraph's schema information to make better optimization decisions. This phase bridges the SPARQL
algebra optimizations (Phase 2) with the DQL execution engine.

### Why Phase 3 Is Critical

With Phase 2 complete (4 SPARQL-only optimizers), we now need to:

1. **Connect to Dgraph schema** - Access predicates, types, indexes
2. **Apply schema constraints** - Use cardinality and selectivity info
3. **Join ordering** - Cost-based decisions using schema stats
4. **Predicate optimization** - Leverage index information
5. **Type constraints** - Use schema type definitions

This is the bridge between SPARQL algebra and DQL execution.

---

## Phase 3 Deliverables

### Week 1: Schema Integration Foundation

#### 1. Schema Analyzer (NEW)

**Purpose**: Extract and analyze Dgraph schema for optimization

**Responsibilities**:

- Read Dgraph schema information
- Extract predicate definitions
- Analyze type hierarchies
- Collect cardinality hints
- Identify indexes and their types

**Implementation**:

```go
type SchemaAnalyzer struct {
    Schema SchemaInfo
}

func (sa *SchemaAnalyzer) GetPredicate(name string) *PredicateInfo
func (sa *SchemaAnalyzer) IsIndexed(pred string, indexType string) bool
func (sa *SchemaAnalyzer) GetCardinality(pred string) uint64
func (sa *SchemaAnalyzer) GetTypes(pred string) []string
```

**Tests**: 8+ comprehensive test cases

#### 2. Predicate Optimizer (NEW)

**Purpose**: Optimize SPARQL patterns based on predicate information

**Responsibilities**:

- Detect indexed predicates
- Push index predicates early
- Optimize predicate patterns
- Handle inverse predicates
- Leverage reverse indexes

**Implementation**:

```go
type PredicateOptimizer struct {
    Schema *SchemaAnalyzer
}

func (po *PredicateOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error)
func (po *PredicateOptimizer) isIndexed(pred string) bool
func (po *PredicateOptimizer) shouldPushEarly(triple *Triple) bool
```

**Tests**: 8+ comprehensive test cases

#### 3. Type Constraint Analyzer (NEW)

**Purpose**: Analyze and apply type constraints from schema

**Responsibilities**:

- Extract type constraints from patterns
- Identify implicit type filters
- Validate type compatibility
- Optimize type checks
- Remove impossible type combinations

**Implementation**:

```go
type TypeConstraintAnalyzer struct {
    Schema *SchemaAnalyzer
}

func (tca *TypeConstraintAnalyzer) Analyze(expr AlgebraExpr) *TypeConstraints
func (tca *TypeConstraintAnalyzer) getTypes(pattern string) []string
func (tca *TypeConstraintAnalyzer) isCompatible(type1, type2 string) bool
```

**Tests**: 10+ comprehensive test cases

### Week 2: Advanced Schema Integration

#### 4. Join Order Optimizer (NEW)

**Purpose**: Use schema info to determine optimal join order

**Responsibilities**:

- Analyze join structure
- Calculate selectivity per triple
- Order joins by selectivity (most selective first)
- Handle multi-way joins
- Consider join correlations

**Implementation**:

```go
type JoinOrderOptimizer struct {
    Schema *SchemaAnalyzer
}

func (joo *JoinOrderOptimizer) OptimizeJoinOrder(expr AlgebraExpr) (AlgebraExpr, error)
func (joo *JoinOrderOptimizer) calculateSelectivity(triple *Triple) float64
func (joo *JoinOrderOptimizer) reorderJoins(bgp *AlgebraBGP) AlgebraExpr
```

**Tests**: 10+ comprehensive test cases

#### 5. Index Selection Optimizer (NEW)

**Purpose**: Select optimal indexes for query execution

**Responsibilities**:

- Analyze available indexes
- Match predicates to indexes
- Consider index selectivity
- Handle composite indexes
- Recommend index usage

**Implementation**:

```go
type IndexSelectionOptimizer struct {
    Schema *SchemaAnalyzer
}

func (iso *IndexSelectionOptimizer) SelectIndexes(expr AlgebraExpr) map[string]*IndexInfo
func (iso *IndexSelectionOptimizer) getBestIndex(pred string) *IndexInfo
func (iso *IndexSelectionOptimizer) calculateIndexCost(idx *IndexInfo) float64
```

**Tests**: 8+ comprehensive test cases

### Week 3: Schema-Aware Optimization Pipeline

#### 6. Statistics Collector (NEW)

**Purpose**: Collect and maintain statistics about query execution

**Responsibilities**:

- Track predicate cardinalities
- Monitor index usage
- Record join selectivity
- Measure actual vs. estimated costs
- Update statistics

**Implementation**:

```go
type StatisticsCollector struct {
    Stats map[string]*PredicateStats
}

func (sc *StatisticsCollector) RecordExecution(pred string, cardinality uint64)
func (sc *StatisticsCollector) GetStats(pred string) *PredicateStats
func (sc *StatisticsCollector) UpdateEstimates()
```

**Tests**: 6+ test cases

#### 7. Optimizer Pipeline Integration (NEW)

**Purpose**: Integrate all Phase 3 optimizers with Phase 2 optimizers

**Responsibilities**:

- Define optimization order
- Handle optimizer interactions
- Cache optimization results
- Provide configuration options
- Enable/disable optimizers

**Implementation**:

```go
type SchemaAwareOptimizerPipeline struct {
    SchemaAnalyzer *SchemaAnalyzer
    Optimizers []Optimizer
}

func (pipeline *SchemaAwareOptimizerPipeline) Optimize(expr AlgebraExpr) (AlgebraExpr, error)
func (pipeline *SchemaAwareOptimizerPipeline) AddOptimizer(opt Optimizer)
func (pipeline *SchemaAwareOptimizerPipeline) SetConfig(config OptimizerConfig)
```

**Tests**: 10+ integration tests

---

## Architecture Overview

### Phase 3 Architecture

```
SPARQL Query
     ↓
[SPARQL Parser & Translator]
     ↓
[Phase 2: SPARQL-Only Optimizers]
  • BIND Optimizer
  • Variable Analyzer
  • FILTER Optimizer
  • UNION Optimizer
     ↓
[Phase 3: Schema-Aware Optimizers]
  • Schema Analyzer (read schema)
  • Predicate Optimizer (use indexes)
  • Type Constraint Analyzer (type checks)
  • Join Order Optimizer (best order)
  • Index Selection Optimizer (best indexes)
  • Statistics Collector (track performance)
     ↓
[Optimized Algebra + Metadata]
     ↓
[DQL Query Engine]
     ↓
Results
```

### Schema Integration Points

```
Dgraph Schema
     ↓
[Schema Analyzer]
     ├─→ Predicate Info (type, indexes, cardinality)
     ├─→ Type Hierarchy (inheritance, compatibility)
     ├─→ Index Info (types, selectivity)
     └─→ Cardinality Stats (predicates, types)
          ↓
     [Used by all Phase 3 optimizers]
```

---

## Implementation Details

### Week 1: Schema Integration Foundation

#### SchemaAnalyzer

**Responsibilities**:

1. Extract schema from Dgraph
2. Build predicate index
3. Build type hierarchy
4. Index cardinality data
5. Track index definitions

**Key Methods**:

- `GetPredicate(name)` - Retrieve predicate info
- `IsIndexed(pred, type)` - Check if indexed
- `GetCardinality(pred)` - Get cardinality hint
- `GetTypes(pred)` - Get associated types
- `BuildIndex()` - Initialize from schema

**Data Structures**:

```go
type PredicateInfo struct {
    Name         string
    Type         string         // scalar, uid, etc.
    ReverseIndex bool
    Index        string         // hash, range, etc.
    Cardinality  uint64
    ListType     bool
    Lang         bool
}

type TypeInfo struct {
    Name       string
    Fields     map[string]*PredicateInfo
    Predicates []string
}

type SchemaInfo struct {
    Predicates map[string]*PredicateInfo
    Types      map[string]*TypeInfo
}
```

**Test Cases**:

1. Load schema from Dgraph
2. Access predicate info
3. Check index presence
4. Get cardinality estimates
5. Retrieve type info
6. Handle missing predicates
7. Cache schema data
8. Update schema

#### PredicateOptimizer

**Responsibilities**:

1. Analyze predicates in patterns
2. Identify indexed predicates
3. Push indexed patterns forward
4. Handle inverse predicates
5. Optimize pattern order

**Optimizations**:

- Push indexed predicates to beginning
- Combine related predicates
- Eliminate redundant patterns
- Use reverse indexes when beneficial

**Test Cases**:

1. Identify indexed predicates
2. Push indexed to front
3. Handle multiple indexes
4. Reverse index usage
5. Pattern combination
6. Filter integration
7. Complex patterns
8. No-op cases

#### TypeConstraintAnalyzer

**Responsibilities**:

1. Extract type constraints
2. Validate type compatibility
3. Remove impossible types
4. Detect type conflicts
5. Optimize type checks

**Constraints**:

- Subject type constraints
- Object type constraints
- Compatibility validation
- Conflict detection

**Test Cases**:

1. Extract type constraints
2. Single type constraints
3. Multiple type constraints
4. Type compatibility check
5. Conflict detection
6. Type hierarchy traversal
7. Implicit type filters
8. Remove redundant types
9. Type union handling
10. Invalid type combinations

### Week 2: Advanced Schema Integration

#### JoinOrderOptimizer

**Responsibilities**:

1. Analyze join structure
2. Calculate selectivity
3. Order joins optimally
4. Handle correlations
5. Consider statistics

**Algorithm**:

1. Extract all triples
2. Calculate selectivity for each
3. Sort by selectivity (most selective first)
4. Reorder joins in algebra tree
5. Preserve semantics

**Selectivity Factors**:

- Predicate cardinality
- Type constraints
- Filter conditions
- Index availability
- Join correlation

**Test Cases**:

1. Simple join reordering
2. Multi-way joins
3. Selectivity calculation
4. Filter integration
5. Type constraints
6. Index consideration
7. Correlated joins
8. Complex patterns
9. No-op cases
10. Performance comparison

#### IndexSelectionOptimizer

**Responsibilities**:

1. Analyze available indexes
2. Match predicates to indexes
3. Calculate index cost
4. Select best index
5. Handle composite indexes

**Index Types**:

- Hash index
- Range index
- Tokenizer index
- Reverse index
- Geo index

**Cost Calculation**:

- Selectivity of index
- Storage overhead
- Query execution cost
- Index maintenance cost

**Test Cases**:

1. Select best index
2. Hash vs range
3. Multiple indexes
4. Composite indexes
5. No suitable index
6. Cost calculation
7. Predicate matching
8. Index selectivity

### Week 3: Integration & Pipeline

#### StatisticsCollector

**Responsibilities**:

1. Track execution stats
2. Update cardinality estimates
3. Monitor index usage
4. Record actual costs
5. Provide statistics

**Statistics**:

- Predicate cardinality
- Type distribution
- Index selectivity
- Join selectivity
- Filter selectivity

**Test Cases**:

1. Record execution
2. Update estimates
3. Access statistics
4. Multiple predicates
5. Type stats
6. Index usage tracking

#### SchemaAwareOptimizerPipeline

**Responsibilities**:

1. Orchestrate optimizers
2. Define optimization order
3. Handle interactions
4. Cache results
5. Configuration

**Optimization Order**:

1. Schema analysis
2. Phase 2 optimizers (runs as is)
3. Schema-specific optimizers
4. Join ordering
5. Index selection
6. Statistics collection

**Configuration**:

- Enable/disable optimizers
- Set optimization level
- Configure statistics
- Set caching behavior

**Integration Tests**:

1. Full pipeline execution
2. All optimizers together
3. Complex queries
4. Edge cases
5. Performance testing
6. Configuration options
7. Error handling
8. Schema changes
9. Statistics accuracy
10. Integration with DQL

---

## Key Differences from Phase 2

### Phase 2: SPARQL-Only Optimizations

- Worked with SPARQL algebra
- No external dependencies
- Pure syntax optimization
- Could be applied independently

### Phase 3: Schema-Aware Optimizations

- Requires schema access
- Depends on external information
- Semantic optimization
- Coordinated application
- Statistics-based decisions

---

## Testing Strategy

### Unit Tests

- Individual optimizer tests
- Schema analyzer tests
- Statistics collector tests
- Index selection tests
- ~50+ individual test cases

### Integration Tests

- Full pipeline tests
- Multiple optimizers working together
- Complex query scenarios
- Schema-dependent optimizations
- ~20+ integration test cases

### End-to-End Tests

- Real SPARQL queries
- Against schema
- Verify correct results
- Measure performance improvement

---

## Success Criteria

### Functionality

✅ Schema analyzer correctly reads schema  
✅ Predicate optimizer identifies indexes  
✅ Type constraint analyzer works correctly  
✅ Join ordering improves query performance  
✅ Index selection picks optimal indexes  
✅ Statistics collector tracks performance  
✅ Pipeline integrates all optimizers

### Quality

✅ 70+ comprehensive test cases  
✅ 100% test pass rate  
✅ No regressions from Phase 2  
✅ Clean code with full documentation  
✅ Proper error handling

### Performance

✅ Measurable query improvement  
✅ Reasonable optimization time  
✅ Efficient schema lookups  
✅ Statistics accuracy

---

## Risk Mitigation

### Schema Dependency

**Risk**: Schema unavailable or invalid  
**Mitigation**: Graceful fallback to Phase 2 optimizations

### Performance Overhead

**Risk**: Optimization too slow  
**Mitigation**: Cache schema, limit analysis depth

### Incorrect Optimization

**Risk**: Schema optimization makes queries slower  
**Mitigation**: Comprehensive tests, validate results

### Integration Issues

**Risk**: Phase 3 breaks Phase 2  
**Mitigation**: Careful integration, run all Phase 2 tests

---

## File Structure

### New Files (Phase 3)

Implementation:

- `sparql/schema_analyzer.go`
- `sparql/predicate_optimizer.go`
- `sparql/type_constraint_analyzer.go`
- `sparql/join_order_optimizer.go`
- `sparql/index_selection_optimizer.go`
- `sparql/statistics_collector.go`
- `sparql/schema_aware_pipeline.go`

Tests:

- `sparql/schema_analyzer_test.go`
- `sparql/predicate_optimizer_test.go`
- `sparql/type_constraint_analyzer_test.go`
- `sparql/join_order_optimizer_test.go`
- `sparql/index_selection_optimizer_test.go`
- `sparql/statistics_collector_test.go`
- `sparql/schema_aware_pipeline_test.go`

### Modified Files (Phase 3)

None expected - Phase 3 is additive to Phase 2.

---

## Timeline

### Week 1: Foundation (Jan 25-31)

- Monday: Schema Analyzer
- Tuesday: Predicate Optimizer
- Wednesday-Thursday: Type Constraint Analyzer
- Friday: Review & testing

### Week 2: Advanced (Feb 1-7)

- Monday: Join Order Optimizer
- Tuesday-Wednesday: Index Selection Optimizer
- Thursday: Statistics Collector
- Friday: Review & testing

### Week 3: Integration (Feb 8-14)

- Monday-Tuesday: Pipeline Integration
- Wednesday-Thursday: Full integration testing
- Friday: Final review & documentation

---

## Metrics to Track

### Implementation

- Lines of code per component
- Test cases per component
- Code coverage percentage
- Time per component

### Quality

- Test pass rate
- Regression count
- Code complexity
- Documentation coverage

### Performance

- Optimization time
- Schema lookup time
- Cache hit rate
- Query improvement percentage

---

## Dependencies

### On Phase 2

✅ Uses all Phase 2 optimizers  
✅ Builds on algebra foundation  
✅ Extends optimization pipeline

### On Dgraph

- Schema access API
- Predicate metadata
- Index definitions
- Type information

### External

- Standard Go libraries
- Existing Dgraph packages

---

## Next After Phase 3

### Phase 4 Options

1. **Authorization Integration** - Consider permissions in optimization
2. **Caching Optimization** - Leverage query result caching
3. **Streaming Optimization** - Handle streaming results
4. **Performance Benchmarking** - Comprehensive perf suite
5. **Advanced Statistics** - Histogram-based cardinality

---

## Success Checklist

- [ ] Schema Analyzer implemented and tested
- [ ] Predicate Optimizer implemented and tested
- [ ] Type Constraint Analyzer implemented and tested
- [ ] Join Order Optimizer implemented and tested
- [ ] Index Selection Optimizer implemented and tested
- [ ] Statistics Collector implemented and tested
- [ ] Pipeline Integration implemented and tested
- [ ] All Phase 2 tests still passing
- [ ] 70+ Phase 3 tests passing
- [ ] Integration tests passing
- [ ] Documentation complete
- [ ] No regressions
- [ ] Performance validated

---

## Questions for Clarification

1. **Schema Access**: How should we access Dgraph schema in tests?
2. **Statistics**: Should we use actual cardinality or estimates?
3. **Caching**: Should schema be cached, and for how long?
4. **Configuration**: What configuration options are most important?
5. **Fallback**: How should we handle schema unavailability?

---

**Status**: Ready for Week 1 implementation  
**Next Step**: Implement Schema Analyzer
