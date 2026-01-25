# Phase 2: Advanced Query Optimization (Implementation Plan)

**Status**: Starting  
**Duration**: Weeks 4-6 (estimated)  
**Goal**: Enhance SPARQL query optimization with cost-based planning and performance improvements

---

## Phase 1 Review

Phase 1 successfully delivered:

- ✅ Complete algebra type system (13 operators)
- ✅ Visitor pattern framework
- ✅ 8 optimization rules in algebra_rewriter.go
- ✅ 250+ tests all passing
- ✅ Full documentation

**Current State**: algebra_rewriter.go has 1,072 LOC with:

- FilterPushdownOptimizer
- JoinReorderOptimizer
- DeadBindEliminationOptimizer
- OptionalSimplificationOptimizer
- ValuesProjectionSimplificationOptimizer
- IdentitySimplificationOptimizer
- UnionDeduplicationAndEmptyRemovalOptimizer
- FilterNormalizationOptimizer
- PathDesugaringOptimizer

---

## Phase 2 Objectives

### 2.1 Cardinality Estimation Framework ⭐

**Goal**: Estimate result set sizes for cost-based optimization

**Implementation**:

```go
type CardinalityEstimator interface {
    EstimateCardinality(expr AlgebraExpr) int
    EstimateSelectivity(filter AlgebraFilter) float64
}
```

**Estimations to implement**:

- BGP cardinality from pattern specificity
- Join output size from child cardinalities
- Filter selectivity from pattern analysis
- Aggregate output size
- OPTIONAL cardinality with null handling

**Testing**: Validate estimates against sample data

### 2.2 Cost-Based Join Reordering Enhancement

**Goal**: Improve join order selection using cardinality estimates

**Current**: JoinReorderOptimizer exists but could be enhanced with:

- Cost model for different join strategies
- Left-linear vs. bushy tree selection
- Nested loop vs. hash join selection
- Semi-join detection and optimization

**Implementation**:

- Use CardinalityEstimator to evaluate join costs
- Implement dynamic programming for join enumeration (Yannakakis)
- Add statistics tracking

### 2.3 Filter Selectivity Analysis

**Goal**: Better predict filter effectiveness

**Implementation**:

- Analyze filter expressions for selectivity
- Handle AND/OR combinations
- Range analysis for comparisons
- String pattern selectivity (REGEX, CONTAINS)

### 2.4 Early Termination Optimization

**Goal**: Optimize queries with LIMIT

**Implementation**:

- Detect LIMIT clauses and push down
- Reorder filters by selectivity
- Use sorted access for ORDER BY + LIMIT

### 2.5 Performance Benchmarking

**Goal**: Measure optimization impact

**Implementation**:

- Create benchmark suite with test queries
- Measure:
  - Algebra compilation time
  - Optimization pass time per rule
  - Memory usage
  - Result query quality improvement

### 2.6 Advanced Pattern Optimization

**Goal**: Optimize specific common patterns

**Implementation**:

- **Values pattern optimization**: Detect UNION of values patterns
- **Cartesian product detection**: Warn on cross products
- **Self-join detection**: Optimize reflexive patterns
- **Transitive closure**: Basic support for path patterns

---

## Deliverables

### Code (500+ LOC)

- `sparql/cardinality.go` - Cardinality estimator (NEW)
- `sparql/algebra_rewriter.go` - Enhanced with cost-based logic (MODIFIED)
- `sparql/cost_model.go` - Cost evaluation functions (NEW)

### Tests (200+ test cases)

- `sparql/cardinality_test.go` - Cardinality estimation tests (NEW)
- `sparql/algebra_optimizer_test.go` - Enhanced with cost-based tests (MODIFIED)
- `sparql/benchmark_test.go` - Performance benchmarks (NEW)

### Documentation

- `PHASE_2_IMPLEMENTATION_PLAN.md` - This document
- `COST_MODEL.md` - Cost model specification
- Code comments on all new functions

---

## Implementation Strategy

### Week 1: Foundation

1. Design cardinality estimation interface
2. Implement basic cardinality estimators
3. Write cardinality tests
4. Integrate with existing optimizers

### Week 2: Enhancement

1. Implement selectivity analysis
2. Enhance join reordering with costs
3. Add filter optimization
4. Write integration tests

### Week 3: Advanced & Performance

1. Implement advanced patterns
2. Add LIMIT optimization
3. Create benchmark suite
4. Performance profiling and tuning

---

## Success Criteria

- [ ] Cardinality estimator implemented and tested
- [ ] Cost-based join reordering working
- [ ] Filter selectivity analysis functional
- [ ] All new code fully tested (95%+ coverage)
- [ ] No regressions from Phase 1
- [ ] Benchmark suite showing improvements
- [ ] Documentation complete

---

## Dependencies & Integration

### Depends On

- Phase 1: Algebra types and basic optimizers ✅

### Enables

- Phase 3: Schema-aware compilation
- Phase 4: Authorization rule optimization
- Phase 5: Ontology reasoning optimization

### Integration Points

- **CardinalityEstimator**: Used by JoinReorderOptimizer
- **CostModel**: Consumed by optimization rules
- **Benchmarks**: Measure overall system performance

---

## Known Constraints

1. **No real statistics**: Will use pattern-based heuristics
   - Phase 3 will integrate with schema for real statistics
2. **Simple cardinality models**: More complex with:
   - Correlated attributes
   - Complex predicates
   - Aggregation selectivity

3. **Single-threaded optimization**: No parallel optimization passes

---

## Contingency Plans

If cardinality estimation is too complex:

- Fall back to simple heuristics (cardinality based on pattern only)
- Skip cost model, use selectivity rules

If benchmarking shows no improvement:

- Document findings
- Identify optimization opportunities for Phase 3
- Prepare for schema-aware approach

---

## Metrics to Track

| Metric              | Target | Phase 1 | Phase 2 Goal |
| ------------------- | ------ | ------- | ------------ |
| Lines of Code       | N/A    | 5,000+  | 5,500+       |
| Test Cases          | 250+   | 250+    | 450+         |
| Avg Algebra Compile | <1ms   | ✅      | <1ms         |
| Optimization Time   | <5ms   | ✅      | <10ms        |
| Code Coverage       | 95%+   | ✅      | 95%+         |
| Regressions         | 0      | ✅      | 0            |

---

## Timeline

- **Start**: January 25, 2026
- **Week 1 (Days 1-4)**: Cardinality estimation
- **Week 2 (Days 5-8)**: Cost-based optimization
- **Week 3 (Days 9-12)**: Advanced patterns & benchmarking
- **Target Completion**: ~February 8, 2026

---

## Next Action

Begin Phase 2 implementation:

1. Create `sparql/cardinality.go` with CardinalityEstimator interface
2. Implement basic cardinality estimates for each operator
3. Write comprehensive tests
4. Integrate with JoinReorderOptimizer

**Status**: 🚀 Ready to Start Phase 2
