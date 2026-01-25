# Phase 3 Week 2: Advanced Schema Integration - IN PROGRESS

**Status**: STARTING  
**Date**: January 25, 2026  
**Week 1 Status**: COMPLETE (3/3 components)  
**Week 2 Target**: 2/2 components

---

## Executive Summary

Phase 3 Week 2 focuses on **advanced schema-aware optimizations** that leverage schema information
for intelligent decision-making:

1. **Join Order Optimizer** - Determine optimal join order using selectivity
2. **Index Selection Optimizer** - Select best indexes for query execution

These optimizers make cost-based decisions using schema statistics to improve query performance.

---

## Week 1 Foundation (COMPLETE ✓)

| Component                | Status   | Tests  | LOC       |
| ------------------------ | -------- | ------ | --------- |
| Schema Analyzer          | DONE     | 21     | 9,140     |
| Predicate Optimizer      | DONE     | 12     | 175       |
| Type Constraint Analyzer | DONE     | 18     | 314       |
| **Week 1 Total**         | **DONE** | **51** | **9,629** |

---

## Week 2 Components

### Component 1: Join Order Optimizer

**Purpose**: Use schema info to determine optimal join order

**Responsibilities**:

- Analyze join structure
- Calculate selectivity per triple
- Order joins by selectivity (most selective first)
- Handle multi-way joins
- Consider join correlations

**Key Methods**:

```go
func (joo *JoinOrderOptimizer) OptimizeJoinOrder(expr AlgebraExpr) (AlgebraExpr, error)
func (joo *JoinOrderOptimizer) calculateSelectivity(triple *Triple) float64
func (joo *JoinOrderOptimizer) reorderJoins(bgp *AlgebraBGP) AlgebraExpr
```

**Test Cases**: 10+ comprehensive tests

**Selectivity Calculation**:

- Indexed predicates: High selectivity (0.1-0.3)
- Type constraints: Medium selectivity (0.3-0.5)
- Unindexed: Low selectivity (0.5-0.9)

**Algorithm**:

1. Extract all triples
2. Calculate selectivity for each
3. Sort by selectivity (most selective first)
4. Reorder joins in algebra tree
5. Preserve semantics

### Component 2: Index Selection Optimizer

**Purpose**: Select optimal indexes for query execution

**Responsibilities**:

- Analyze available indexes
- Match predicates to indexes
- Calculate index cost
- Select best index
- Handle composite indexes

**Key Methods**:

```go
func (iso *IndexSelectionOptimizer) SelectIndexes(expr AlgebraExpr) map[string]*IndexInfo
func (iso *IndexSelectionOptimizer) getBestIndex(pred string) *IndexInfo
func (iso *IndexSelectionOptimizer) calculateIndexCost(idx *IndexInfo) float64
```

**Test Cases**: 8+ comprehensive tests

**Index Types**:

- Hash index (exact match)
- Range index (comparisons)
- Tokenizer index (text search)
- Reverse index (inverse predicates)
- Geo index (spatial queries)

**Cost Calculation**:

- Selectivity of index
- Storage overhead
- Query execution cost
- Index maintenance cost

---

## Implementation Plan

### Day 1: Join Order Optimizer

**Steps**:

1. Create `sparql/join_order_optimizer.go`
2. Implement selectivity calculation
3. Implement join reordering logic
4. Create comprehensive tests (10+ cases)
5. Verify no Phase 2 regressions

**Estimated**: ~250 LOC + tests

### Day 2: Index Selection Optimizer

**Steps**:

1. Create `sparql/index_selection_optimizer.go`
2. Implement index analysis
3. Implement cost calculation
4. Implement best index selection
5. Create comprehensive tests (8+ cases)

**Estimated**: ~200 LOC + tests

### Day 3: Integration & Testing

**Steps**:

1. Integration tests for both optimizers
2. Performance validation
3. Documentation
4. Week 2 summary

---

## Architecture

### Join Order Optimizer Location

```
Algebra Expression
      ↓
[Predicate Optimizer] (Phase 3 Week 1)
      ↓
[Type Constraint Analyzer] (Phase 3 Week 1)
      ↓
[Join Order Optimizer] ← YOU ARE HERE
      ↓
[Index Selection Optimizer]
      ↓
Optimized Algebra + Metadata
```

### Selectivity Factors

```
selectivity = index_factor × type_factor × filter_factor

index_factor:
  - Indexed: 0.1-0.3 (high)
  - Unindexed: 0.5-0.9 (low)

type_factor:
  - With type constraint: 0.5-0.7
  - Without: 1.0

filter_factor:
  - With FILTER: 0.1-0.9 (depends on condition)
  - Without: 1.0
```

---

## Success Criteria

### Functionality

- [ ] Join Order Optimizer correctly analyzes joins
- [ ] Selectivity calculation is accurate
- [ ] Join reordering improves performance
- [ ] Index Selection Optimizer identifies indexes
- [ ] Cost calculation is reasonable
- [ ] Best index selection works correctly

### Quality

- [ ] 18+ comprehensive test cases (10 + 8)
- [ ] 100% test pass rate
- [ ] No regressions from Phase 2/3 Week 1
- [ ] Clean code with full documentation
- [ ] Proper error handling

### Performance

- [ ] Join ordering improves query time
- [ ] Index selection reduces I/O
- [ ] Optimization overhead acceptable
- [ ] Selectivity estimates reasonable

---

## Testing Strategy

### Join Order Optimizer Tests

1. Simple two-way join
2. Multi-way join (3+ triples)
3. Selectivity calculation
4. Reordering verification
5. Filter integration
6. Type constraints
7. Index consideration
8. Correlated joins
9. Complex patterns
10. No-op cases

### Index Selection Optimizer Tests

1. Select hash index
2. Select range index
3. Multiple indexes
4. No suitable index
5. Cost calculation
6. Predicate matching
7. Index selectivity
8. Composite indexes

---

## Known Challenges

1. **Join Correlation**: Some joins share variables - need to preserve join structure
2. **Cardinality Estimates**: May be inaccurate - could cache better estimates in Phase 4
3. **Multi-way Join Optimization**: NP-hard problem - use heuristics
4. **Index Composition**: Composite indexes need special handling

---

## Risk Mitigation

### Challenge: Incorrect Join Reordering

**Mitigation**:

- Extensive testing with various patterns
- Preserve semantics validation
- Fallback to original order if invalid

### Challenge: Poor Selectivity Estimates

**Mitigation**:

- Use conservative estimates
- Support custom statistics in future
- Track actual vs estimated

### Challenge: Performance Regression

**Mitigation**:

- Baseline performance tests
- Compare before/after optimization
- Disable optimizer if slower

---

## Files to Create

### Implementation Files

- `sparql/join_order_optimizer.go` (~250 LOC)
- `sparql/index_selection_optimizer.go` (~200 LOC)

### Test Files

- `sparql/join_order_optimizer_test.go` (~400 LOC)
- `sparql/index_selection_optimizer_test.go` (~300 LOC)

### Total Week 2: ~1,150 LOC

---

## Metrics to Track

### Implementation

- [ ] LOC per component
- [ ] Test cases per component
- [ ] Code coverage %
- [ ] Time per component

### Quality

- [ ] Test pass rate
- [ ] Regression count
- [ ] Code complexity
- [ ] Documentation coverage

### Performance

- [ ] Optimization time
- [ ] Selectivity accuracy
- [ ] Query improvement %
- [ ] Cache hit rate

---

## Timeline

### Saturday (Jan 25)

- [x] Phase 3 Week 1 Complete (3/3 components)
- [ ] Start Join Order Optimizer

### Sunday (Jan 26)

- [ ] Join Order Optimizer complete + tested
- [ ] Index Selection Optimizer started

### Monday (Jan 27)

- [ ] Index Selection Optimizer complete + tested
- [ ] Integration testing
- [ ] Week 2 summary

---

## Dependencies

### On Phase 3 Week 1

- ✅ Schema Analyzer (read schema info)
- ✅ Predicate Optimizer (already reordered)
- ✅ Type Constraint Analyzer (type info)

### External

- Standard Go libraries
- Existing Dgraph packages

---

## Next After Week 2

### Phase 3 Week 3: Integration & Pipeline

1. **Statistics Collector** - Track execution stats
2. **Schema-Aware Pipeline** - Orchestrate all optimizers
3. **Full Integration Tests** - All optimizers together

---

## Notes for Implementation

1. **Selectivity is Key**: Most important factor for join ordering
2. **Conservative Estimates**: Better to be safe than sorry
3. **Preserve Semantics**: Don't reorder if semantics change
4. **Cache Metadata**: Store index info for repeated queries
5. **Graceful Degradation**: Work without schema if needed

---

**Status**: WEEK 2 STARTING

Next: Implement Join Order Optimizer
