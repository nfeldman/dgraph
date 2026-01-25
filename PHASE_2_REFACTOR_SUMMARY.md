# Phase 2 Refactor: Removing Redundant Optimizers

**Date**: January 25, 2026  
**Status COMPLETE **: **Result**: 19,299 LOC removed, architecture improved

---

## The Problem

We were implementing optimizations at the SPARQL level that are **already done better by the DQL
query engine**:

1. **Join Reordering** - We used heuristics based on cardinality estimates
   - Problem: DQL engine has real schema statistics
   - Our approach: Redundant and inferior

2. **Filter Selectivity Optimization** - We reordered filters by selectivity
   - Problem: DQL engine already does this with real data
   - Our approach: Duplicate work

3. **Cost-Based Selection** - We estimated execution costs
   - Problem: Without schema stats, estimates are unreliable
   - Our approach: Inferior to DQL's real statistics

## The Insight

When a query goes through the full pipeline:

```
 Execution
```

The DQL engine **re-optimizes everything** using real schema statistics.

Therefore, doing optimization at the SPARQL level for:

- Join ordering
- Filter ordering
- Cost-based decisions

...is **wasted effort** - the DQL engine will override our decisions anyway.

---

## What We Removed

### Redundant Code (19,299 LOC)

- `cost_optimizer.go` (9,106 LOC)
  - CostBasedJoinReorderer
  - SelectiveFilterOptimizer
  - CostBasedOptimizationPipeline

- `cost_optimizer_test.go` (10,193 LOC)
  - 15+ tests for the above
  - All redundant

### These Were Redundant Because:

- Join reordering: DQL engine already does this
- Filter ordering: DQL engine already does this
- Cost modeling: Needs real stats (Phase 3+)

---

## What We Kept

### Cardinality System (Still Useful)

- `cardinality.go` - Reference implementation
- `cardinality_test.go` - Tests
- **Purpose**: Introspection and analysis, NOT decision-making
- **Use**: For future Phase 3 schema integration

### Phase 1 Algebra System (Core)

- All 13 algebra operators
- Visitor pattern
- All optimization rules that are SPARQL-specific

---

## What We'll Implement Instead (Phase 2 Week 3)

**SPARQL-Only Optimizations** that CANNOT be done at DQL level:

### 1. OPTIONAL Pattern Optimization

- Merge compatible OPTIONAL patterns
- Detect when OPTIONAL can become INNER JOIN
- **Why SPARQL-only**: LEFT JOIN semantics specific to SPARQL

### 2. FILTER Expression Rewriting

`true` `!a || !b`

- Remove trivial filters
- **Why SPARQL-only**: SPARQL filter syntax/semantics

### 3. BIND Expression Optimization

- Remove unused bindings
- Dead binding elimination
- **Why SPARQL-only**: BIND is SPARQL-specific feature

### 4. Variable Scope Analysis

- Detect unused variables
- Remove unnecessary projections
- **Why SPARQL-only**: Graph pattern variable scoping

### 5. UNION Pattern Optimization

- Merge compatible branches
- Remove duplicates
- **Why SPARQL-only**: UNION semantics are SPARQL-specific

### 6. Property Path Desugaring

- Convert `a/b/c` to joins
- Optimize path patterns
- **Why SPARQL-only**: Property paths are SPARQL 1.1 feature

---

## Metrics

### Before Refactor

```
Total LOC: 40,349
  - Implementation: 20,091
  - Tests: 20,258
Tests: 287+
Files with redundancy: cost_optimizer.go + test file
```

### After Refactor

```
Total LOC: 21,050
  - Implementation: 10,985 (cardinality only)
  - Tests: 10,065 (cardinality only)
Tests: ~265
Redundant files: Removed
```

### Reduction

- **LOC Removed**: 19,299 (47.8%)
- **Tests Removed**: ~22 (7.7%)
- **Files Deleted**: 2
- **Cleanliness**: Significantly improved

---

## Architecture Improvement

### OLD Architecture (Redundant)

```

  SPARQL Optimizer
  (Cost- Uses heuristicsbased  decisions)
  ( RedundantJoin  reordering)
  ( RedundantFilter  optimization)

 (outputs DQL)

  DQL Query Engine
  (RE- With real statisticsOPTIMIZES  SAME)
  (Join reordering)
  (Filter optimization)
  (Cost-based decisions)

```

### NEW Architecture (Clean)

```

  SPARQL Optimizer
  ( SPARQL-onlyOPTIONAL  merging)
  ( SPARQL-onlyFILTER  rewriting)
  ( SPARQL-onlyBIND  optimization)
  ( SPARQL-onlyVariable  analysis)

 (outputs DQL)

  DQL Query Engine
  (Handles all DQL
   optimization with
   real statistics)

```

---

## Why This Is Better

1. **No Redundancy**
   - Don't waste effort optimizing what DQL will re-optimize
   - Each level focuses on what it does best

2. **Better Results**
   - DQL uses real schema statistics
   - Heuristics at SPARQL level are less accurate

3. **Cleaner Code**
   - Removed 19K LOC of redundant code
   - Clear separation of concerns
   - Each component has one responsibility

4. **Faster Development**
   - Don't implement what already exists
   - Focus on SPARQL-unique features
   - Avoid duplicate maintenance

5. **Better Performance**
   - No wasted optimization cycles
   - Each level optimizes what it's good at
   - Use statistics when available (DQL level)

---

## Validation

### Tests Still Pass

```
 All 265+ tests passing
 Zero regressions
 Algebra tests: PASS
 Cardinality tests: PASS
 Translator tests: PASS
 E2E tests: PASS
```

### Files Verified

- `cost_optimizer.go`: Removed
- `cost_optimizer_test.go`: Removed
- `cardinality.go`: Present and working
- `algebra.go`: Present and working
- All other core files: Intact

---

## Commit

**Hash**: `45952cb07`  
**Message**: "Refactor Phase 2: Remove redundant optimizers, focus on SPARQL-only optimizations"

**Changes**:

- Deleted: `cost_optimizer.go` (9,106 LOC)
- Deleted: `cost_optimizer_test.go` (10,193 LOC)
- Added: `PHASE_2_REVISED_PLAN.md` (explanation)
- Modified: `PHASE_2_WEEK_2_PROGRESS.md` (updated)

---

## Next Steps (Phase 2 Week 3)

With the redundant code removed and architecture clarified, we'll implement:

1. **OptionalPatternOptimizer** - Merge OPTIONAL patterns
2. **FilterExpressionOptimizer** - Rewrite FILTER expressions
3. **BindExpressionOptimizer** - Eliminate unused bindings
4. **VariableScopeAnalyzer** - Remove unused projections
5. **UnionPatternOptimizer** - Merge UNION branches
6. **PropertyPathDesugaringOptimizer** - Convert paths to joins

Each of these is **impossible to do at DQL level** because they require SPARQL-specific semantics.

---

## Lessons Learned

1. **Architecture Clarity**
   - Each component should optimize what's unique to it
   - Don't replicate work done downstream

2. **Understanding the Pipeline**
   - Know what DQL will do with your output
   - Avoid optimizations that will be overridden

3. **Focus on Unique Value**
   - SPARQL optimizer: SPARQL semantics
   - DQL optimizer: DQL semantics with statistics
   - Clear division of labor

---

## Conclusion

This refactor **improves both code quality and architecture**:

Removes 19,299 LOC of redundant code  
 Eliminates duplicate optimization work  
 Clarifies component responsibilities  
 Improves overall performance  
 Sets up Phase 2 Week 3 for success

The key insight: **Optimize at the right level with the right tools.**

- SPARQL level: Use SPARQL-specific semantics
- DQL level: Use real schema statistics

This approach produces better results with less code.

---

**Status Refactor Complete **: **Next**: Phase 2 Week 3 Implementation (SPARQL-only optimizers)
