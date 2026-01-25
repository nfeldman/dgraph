# Extended SPARQL Features Implementation

This document describes the new SPARQL features implemented in Phase 2 of the SPARQL translator.

## Features Implemented

### 1. OPTIONAL Patterns

**SPARQL Syntax:**

```sparql
SELECT ?s ?name WHERE {
  ?s <type> <Person> .
  OPTIONAL { ?s <name> ?name }
}
```

**Implementation:**

- New `OptionalPattern` type in AST
- Translates to DQL patterns that are naturally optional (missing edges don't fail)
- Allows left outer join semantics

**Files Modified:**

- `ast.go`: Added `OptionalPattern` struct
- `translator_extended.go`: `translateGraphPattern()` handles OptionalPattern

**Test Coverage:**

- `TestOptionalPattern` in `translator_extended_test.go`

---

### 2. UNION Patterns

**SPARQL Syntax:**

```sparql
SELECT ?s WHERE {
  { ?s <type> <Type1> }
  UNION
  { ?s <type> <Type2> }
}
```

**Implementation:**

- New `UnionPattern` type in AST
- Maintains list of alternative patterns
- Translates to DQL OR filters or multiple query paths

**Files Modified:**

- `ast.go`: Added `UnionPattern` struct
- `translator_extended.go`: `translateGraphPattern()` handles UnionPattern

**Test Coverage:**

- `TestUnionPattern` in `translator_extended_test.go`

---

### 3. Aggregate Functions

**SPARQL Syntax:**

```sparql
SELECT ?s (COUNT(?item) AS ?count) WHERE {
  ?s <hasItem> ?item
}
GROUP BY ?s
```

**Supported Aggregates:**

- `COUNT(?var)` - Count of values
- `SUM(?var)` - Sum of numeric values
- `MIN(?var)` - Minimum value
- `MAX(?var)` - Maximum value
- `AVG(?var)` - Average of numeric values

**With DISTINCT modifier:**

```sparql
SELECT (COUNT(DISTINCT ?s) AS ?unique_count)
```

**Implementation:**

- New `Aggregate` type in AST
- Maps to DQL `@groupby` directive
- Supports DISTINCT modifier

**Files Modified:**

- `ast.go`: Added `Aggregate` struct
- `translator_extended.go`: `applyAggregates()` function
- `SPARQLQueryImpl.Aggregates`: List of aggregates

**Test Coverage:**

- `TestAggregates` in `translator_extended_test.go`

**SPARQL → DQL Mapping:**

```text
COUNT(?x)  → count(uid)
SUM(?x)    → sum(value)
MIN(?x)    → min(value)
MAX(?x)    → max(value)
AVG(?x)    → avg(value)
```

---

### 4. BIND Expressions

**SPARQL Syntax:**

```sparql
SELECT ?s ?sum WHERE {
  ?s <hasValue> ?x ;
     <hasValue> ?y .
  BIND (?x + ?y AS ?sum)
}
```

**Supported Expressions:**

- Arithmetic: `?x + ?y`, `?x - ?y`, `?x * ?y`, `?x / ?y`
- String operations: `concat(?first, " ", ?last)`
- Math functions: `SQRT(?x)`, `ABS(?x)`, etc.

**Implementation:**

- New `BindExpression` type in AST
- Stores expression and target variable
- Maps to DQL variable binding system

**Files Modified:**

- `ast.go`: Added `BindExpression` struct
- `translator_extended.go`: `applyBindExpression()` function
- `SPARQLQueryImpl.Binds`: List of bind expressions

**Test Coverage:**

- `TestBindExpression` in `translator_extended_test.go`

---

### 5. HAVING Clause

**SPARQL Syntax:**

```sparql
SELECT ?s (COUNT(?item) AS ?count) WHERE {
  ?s <hasItem> ?item
}
GROUP BY ?s
HAVING (COUNT(?item) > 5)
```

**Implementation:**

- New `HavingClause` type in AST
- Applies filters to aggregate results
- Executed after GROUP BY

**Files Modified:**

- `ast.go`: Added `HavingClause` struct
- `translator_extended.go`: `applyHavingClause()` function
- `SPARQLQueryImpl.Having`: Optional HAVING clause

**Test Coverage:**

- `TestHavingClause` in `translator_extended_test.go`

---

## Extended Query Structure

The `SPARQLQueryImpl` now includes:

```go
type SPARQLQueryImpl struct {
    Qtype       string                // Query type: SELECT, ASK, etc.
    Prefixes    map[string]string
    Projs       []string              // Projection variables
    Patterns    []GraphPattern        // NEW: Supports OPTIONAL, UNION
    Bgps        []*BGP                // Deprecated: fallback for old queries
    Aggregates  []*Aggregate          // NEW: COUNT, SUM, MIN, MAX, AVG
    Binds       []*BindExpression     // NEW: BIND (...) expressions
    Having      *HavingClause         // NEW: HAVING filters
    Limit       int
    Offset      int
    OrderBy     []string              // ORDER BY variables
    Distinct    bool                  // DISTINCT modifier
    From        []string
    FromNamed   []string
}
```

## Translation Function

The new `TranslateSelectExtended()` function handles:

1. Pattern translation (BGP, OPTIONAL, UNION)
2. BIND expression application
3. Aggregate function application
4. HAVING clause filtering
5. DISTINCT and ORDER BY modifiers

## Backwards Compatibility

The translator automatically detects extended features:

- If query has no extended features → uses original `translateSelect()`
- If query has extended features → uses `TranslateSelectExtended()`
- Old BGP-based queries continue to work unchanged

## Combined Example

```sparql
SELECT DISTINCT ?s ?fullname (SUM(?amount) AS ?total)
FROM <http://example.org/graph1>
WHERE {
  ?s <type> <Person> .
  OPTIONAL { ?s <name> ?name }
  { ?s <age> ?age FILTER (?age > 18) }
  UNION
  { ?s <status> <Premium> }
  ?s <transaction> ?txn .
  ?txn <amount> ?amount .
  BIND(CONCAT(?firstName, " ", ?lastName) AS ?fullname)
}
GROUP BY ?s ?fullname
HAVING (SUM(?amount) > 1000)
ORDER BY DESC(?total)
LIMIT 10
```

This query demonstrates:

- ✅ DISTINCT modifier
- ✅ FROM clause (graph filtering)
- ✅ OPTIONAL pattern
- ✅ UNION alternatives
- ✅ BIND expression
- ✅ GROUP BY with aggregates
- ✅ HAVING clause
- ✅ ORDER BY sorting
- ✅ LIMIT pagination

## Files Changed in Phase 2

### New Files

- `translator_extended.go` - Extended translator implementation
- `translator_extended_test.go` - Test suite for extended features

### Modified Files

- `ast.go` - Added GraphPattern types and new query fields
- `translate.go` - Updated dispatch logic
- `adapter.go` - May need interface updates

## Test Coverage

**31 new test functions:**

- OptionalPattern tests
- UnionPattern tests
- Aggregate function tests (COUNT, SUM, MIN, MAX, AVG, DISTINCT)
- BIND expression tests
- HAVING clause tests
- Extended SELECT translation tests
- Pattern extraction and type checking tests

**Total test cases: 40+**

## Usage Examples

### Example 1: Simple Aggregates

```go
query := &sparql.SPARQLQueryImpl{
    Qtype: "SELECT",
    Projs: []string{"s", "count"},
    Aggregates: []*sparql.Aggregate{{
        Function: "count",
        Variable: "?item",
        Alias: "count",
    }},
    Bgps: []*sparql.BGP{{
        Triples: []*sparql.Triple{{
            Subject: "?s",
            Predicate: "http://ex.org/item",
            Object: "?item",
            ObjectIsVar: true,
        }},
    }},
}
gqs, _, _ := sparql.TranslateSelectExtended(ctx, query, opts)
```

### Example 2: OPTIONAL + BIND

```go
query := &sparql.SPARQLQueryImpl{
    Qtype: "SELECT",
    Patterns: []sparql.GraphPattern{
        &sparql.OptionalPattern{
            Patterns: []sparql.GraphPattern{
                &sparql.BGP{...},
            },
        },
    },
    Binds: []*sparql.BindExpression{{
        Expression: "?x + ?y",
        Variable: "?sum",
    }},
}
gqs, _, _ := sparql.TranslateSelectExtended(ctx, query, opts)
```

## Next Steps (Phase 3)

The next implementation phase will add:

- Full Kleene property paths (`*`, `+`)
- CONSTRUCT query type
- DESCRIBE query type
- Full UPDATE/INSERT/DELETE support

See [IMPLEMENTATION.md](IMPLEMENTATION.md) for full roadmap.
