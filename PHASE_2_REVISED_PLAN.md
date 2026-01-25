# Phase 2: SPARQL-Level Optimizations (Revised Plan)

**Status**: Refocused on SPARQL-only optimizations  
**Date**: January 25, 2026  
**Rationale**: Avoid duplication with DQL query engine

---

## Optimization Classification

### ❌ Optimizations to REMOVE (Redundant with DQL)

These should be deleted - they duplicate work the DQL engine already does:

1. **CostBasedJoinReorderer** ❌
   - Reason: DQL query engine already does join reordering
   - DQL has real statistics from schema
   - Heuristic estimates are redundant without schema data
   - Will be re-optimized when query goes through DQL

2. **SelectiveFilterOptimizer** ❌
   - Reason: DQL query engine already reorders filters
   - DQL can use index information
   - Filter selectivity ordering is basic optimizer feature
   - Will be re-optimized at DQL level

3. **CardinalityCostEstimator** ⚠️ Partial Keep
   - Keep heuristic-based cardinality estimation for SPARQL analysis
   - Remove cost-based decision making (use schema stats instead)
   - Keep only for introspection/debugging purposes

---

## ✅ Optimizations to KEEP (SPARQL-Only)

These are unique to SPARQL and CANNOT be done at DQL level:

### 1. **OPTIONAL Pattern Optimization**

**What**: Merge compatible OPTIONAL patterns and detect convertible patterns

**Examples**:

```sparql
# Before: Two OPTIONAL blocks
SELECT ?x ?y ?z WHERE {
  ?x rdf:type Person .
  OPTIONAL { ?x foaf:name ?y }
  OPTIONAL { ?x foaf:age ?z }
}

# After: Merged (but still LEFT JOIN semantics)
SELECT ?x ?y ?z WHERE {
  ?x rdf:type Person .
  OPTIONAL {
    ?x foaf:name ?y .
    ?x foaf:age ?z .
  }
}
```

**Why Only SPARQL**: SPARQL OPTIONAL has specific LEFT JOIN semantics that only exist at SPARQL
level.

### 2. **UNION Pattern Optimization**

**What**:

- Merge compatible UNION branches
- Remove duplicate branches
- Convert UNION to UNION ALL when possible
- Eliminate unreachable branches

**Examples**:

```sparql
# Before: Duplicate branches
SELECT ?x WHERE {
  { ?x rdf:type Person } UNION
  { ?x rdf:type Person }  # Duplicate
}

# After: Single branch
SELECT ?x WHERE {
  ?x rdf:type Person
}
```

**Why Only SPARQL**: UNION semantics are SPARQL-specific. DQL doesn't have UNION.

### 3. **FILTER Expression Rewriting**

**What**: Simplify and optimize SPARQL filter expressions

**Operations**:

- Constant folding: `1 + 2 > 2` → `true`
- De Morgan's laws: `!(a && b)` → `!a || !b`
- Simplify comparisons: `?x = ?x` → `true`
- Remove always-true filters
- Remove always-false filters (convert to empty)
- Combine adjacent filters: `FILTER(?x > 5) FILTER(?x < 10)` → `FILTER(?x > 5 && ?x < 10)`

**Why Only SPARQL**: Filter syntax and semantics are SPARQL-specific.

### 4. **Variable Scope Analysis & Dead Variable Elimination**

**What**: Identify and remove unused variables

**Examples**:

```sparql
# Before: ?unused is not in SELECT
SELECT ?x WHERE {
  ?x foaf:name ?y .
  ?unused foaf:age ?age .
}

# After: Remove unused pattern
SELECT ?x WHERE {
  ?x foaf:name ?y .
}
```

**Why Only SPARQL**: Variable scoping rules are specific to SPARQL graph patterns.

### 5. **Property Path Desugaring & Optimization**

**What**: Convert SPARQL property paths to joins and optimize

**Examples**:

```sparql
# Before: Property path
SELECT ?x WHERE {
  ?x foaf:knows/foaf:name ?name .
}

# After: Desugared to joins
SELECT ?x WHERE {
  ?x foaf:knows ?y .
  ?y foaf:name ?name .
}
```

**Paths to support**:

- Simple sequences: `a/b`
- Inverse: `^a`
- Alternatives: `a|b`
- Zero-or-one: `a?`
- Zero-or-more: `a*`
- One-or-more: `a+`

**Why Only SPARQL**: Property paths are SPARQL 1.1 feature, don't exist at DQL level.

### 6. **BIND Expression Optimization**

**What**:

- Remove unused BIND expressions
- Inline simple bindings
- Detect dead bindings

**Examples**:

```sparql
# Before: Unused binding
SELECT ?x WHERE {
  ?x foaf:name ?name .
  BIND(?name AS ?unused) .
}

# After: Remove binding
SELECT ?x WHERE {
  ?x foaf:name ?name .
}
```

**Why Only SPARQL**: BIND is SPARQL-specific, bindings exist only in SPARQL algebra.

### 7. **Aggregate Optimization**

**What**:

- Remove unnecessary GROUP BY when single group
- Optimize COUNT(\*) vs COUNT(?var)
- Push aggregate conditions earlier when possible

**Why Only SPARQL**: Aggregate semantics in SPARQL differ from SQL/DQL.

---

## Removed Components

### CostBasedJoinReorderer ❌

**File**: `sparql/cost_optimizer.go` lines 1-115  
**Reason**: DQL engine does this better with real statistics

### SelectiveFilterOptimizer ❌

**File**: `sparql/cost_optimizer.go` lines 156-270  
**Reason**: DQL engine already reorders filters optimally

### CardinalityCostEstimator ❌ (mostly)

**File**: `sparql/cardinality.go` lines 285-359  
**Reason**: Cost models need schema statistics to be accurate

---

## Implementation Priority

### Phase 2 Week 3 (Revised)

**HIGH PRIORITY** (Week 3):

1. OPTIONAL Pattern Optimization - High impact, SPARQL-only
2. FILTER Expression Rewriting - High impact, SPARQL-only
3. BIND Expression Elimination - Medium impact, simple to implement
4. Variable Scope Analysis - Medium impact, valuable optimization

**MEDIUM PRIORITY** (Future phases): 5. UNION Pattern Optimization - Medium impact 6. Property Path
Desugaring - Medium impact, larger implementation 7. Aggregate Optimization - Lower priority

**NOT IMPLEMENTING**:

- Join reordering (redundant)
- Filter selectivity ordering (redundant)
- Cost-based optimization (needs schema stats)

---

## Code Changes Needed

### Remove Redundant Code

1. Delete `CostBasedJoinReorderer` implementation
2. Delete `SelectiveFilterOptimizer` implementation
3. Delete `CostBasedOptimizationPipeline` (no longer needed)
4. Delete associated tests for above

### Keep for Reference/Debugging

- `CardinalityEstimator` interface (for analysis only)
- `cardinality.go` estimates (for introspection)
- Tests for cardinality (as reference)

### New Implementations Needed

- `OptionalPatternOptimizer`
- `FilterExpressionOptimizer`
- `BindExpressionOptimizer`
- `VariableScopeAnalyzer`

---

## Rationale

### Why DQL Optimizations Are Redundant

When a SPARQL query is:

1. Parsed to AST
2. Converted to DQL
3. Executed by DQL engine

The DQL engine will:

- Reorder joins based on actual schema statistics ✅
- Reorder filters based on selectivity ✅
- Push down limits and filters ✅
- Select appropriate join algorithms ✅

**Therefore**: Doing these optimizations at SPARQL level is wasted effort.

### Why SPARQL-Only Optimizations Are Necessary

Some SPARQL features disappear after DQL conversion:

- OPTIONAL patterns become part of the query structure
- FILTER expressions get converted to DQL predicates
- BIND expressions create aliases
- Property paths are expanded

**Therefore**: These must be optimized BEFORE DQL conversion.

---

## Benefits of This Approach

1. **No Redundancy**: Don't duplicate DQL engine work
2. **Better Results**: Use actual statistics at DQL level, not heuristics
3. **Smaller Codebase**: Remove ~20K lines of redundant code
4. **Focused Optimization**: Only optimize what can't be done later
5. **Clear Separation**: SPARQL handles SPARQL semantics, DQL handles DQL

---

## Implementation Plan for Week 3 (Revised)

### Step 1: Remove Redundant Code (1-2 hours)

- Delete `cost_optimizer.go`
- Delete `cost_optimizer_test.go`
- Delete selective filter tests
- Keep `cardinality.go` for reference

### Step 2: Implement OPTIONAL Optimization (2-3 hours)

- Detect compatible OPTIONAL patterns
- Merge non-conflicting OPTIONALs
- Detect convertible OPTIONALs (→ INNER JOIN)

### Step 3: Implement FILTER Rewriting (2-3 hours)

- Constant folding
- Boolean simplification
- Merge adjacent filters
- Remove trivial filters

### Step 4: Implement BIND Optimization (1-2 hours)

- Detect unused bindings
- Remove dead bindings
- Inline simple expressions

### Step 5: Implement Variable Scope Analysis (1-2 hours)

- Track variable definitions
- Identify unused variables
- Remove unnecessary projections

### Testing & Integration (2-3 hours)

- Test each optimizer independently
- Integration tests
- Regression testing

---

## Expected Outcome

**Before**:

- 40K+ LOC of code (including redundant optimizers)
- 287+ tests (including redundant tests)
- Duplicate optimization work

**After**:

- 20K LOC (removed redundant parts)
- ~200 tests (removed redundant tests)
- Clear focus on SPARQL-only optimizations
- Better performance (use DQL's real optimization)

---

## Conclusion

This revised approach:

1. ✅ Eliminates redundant work
2. ✅ Focuses on SPARQL-unique features
3. ✅ Leverages DQL's optimization capabilities
4. ✅ Reduces code complexity
5. ✅ Produces better optimizations overall

The key insight: **Don't optimize at SPARQL level what DQL will optimize anyway.**

---

**Status**: Plan Updated - Ready to Implement
