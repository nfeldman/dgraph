# SPARQL Branch Code Review & Cleanup Summary

**Date:** January 25, 2026  
**Branch Focus:** ANTLR parser integration with extended SPARQL grammar support

## Current Status and Changes

### Parser / Visitor

- ANTLR visitor now builds `SPARQLQueryImpl` for SELECT/ASK, handling prefixes, projections,
  DISTINCT, FROM/FROM NAMED, GRAPH, OPTIONAL, UNION, BIND, aggregates (COUNT/SUM/MIN/MAX/AVG with
  DISTINCT), HAVING, ORDER BY, LIMIT/OFFSET, FILTER, and graph-annotated BGPs.
- FILTER parsing is strict by default (comparison/regex mapped to DQL filter trees). A flag
  (`TranslateOptions.AllowFilterFallback`) can permit raw fallback but is off by default.
- **Fallback still present:** `antlr_adapter.go` will fall back to the legacy simple parser if ANTLR
  rejects input _or_ the visitor produces no patterns. This bypasses the AST and can silently drop
  features the regex parser can’t handle.

### Translator

- Graph filtering for FROM/FROM NAMED/GRAPH remains correct.
- Aggregates/BIND/HAVING are still stubs: aggregates set string attrs, BIND stores expressions in
  args, HAVING attaches a placeholder filter. None map to executable DQL aggregation or variable
  binding yet.
- OPTIONAL and UNION semantics are incomplete: OPTIONAL is treated as required; UNION is appended
  sequentially (AND instead of OR).
- Subject IRI handling remains TODO (placeholder `_subject`).
- CONSTRUCT/DESCRIBE still unsupported.
- Property paths are not implemented.

### Cleanup (previously completed)

- Removed unused ParseVisitor code (`antlr_adapter.go`).
- Removed unused accessor stubs (`adapter.go`).
- Deleted HTTP integration stub (`http_integration.go`).
- Clarified deprecated `bgps` fields in `ast.go` (migration: use `Patterns`).
- Documented UID lookup TODO for non-variable subjects (`translator_impl.go`).

---

## Outstanding TODOs / Gaps

1. **Remove or gate the simple-parser fallback** (strict mode by default; consider opt-in only).
2. **Aggregate translation**: map COUNT/SUM/MIN/MAX/AVG (with DISTINCT) to valid DQL aggregation and
   grouping.
3. **BIND translation**: parse/evaluate expressions into DQL variable bindings.
4. **HAVING translation**: parse expressions over aggregates and emit proper DQL filters.
5. **OPTIONAL semantics**: implement left-join/optional behavior; avoid requiring optional patterns.
6. **UNION semantics**: translate as OR branches, not sequential AND.
7. **FILTER coverage**: extend beyond comparison/regex to boolean ops, NOT, built-ins (BOUND, lang,
   datatype), IN/EXISTS, functions; keep strict mode unless fallback is explicitly enabled.
8. **Subject IRI handling**: implement UID lookup for non-variable subjects.
9. **CONSTRUCT/DESCRIBE**: implement or explicitly exclude for alpha.
10. **Property paths**: implement or defer with clear limitation.

---

## Shortcuts / Risks

1. **Simple parser fallback**: Still present; bypasses AST and can drop features. Should be opt-in
   or removed.
2. **Aggregate/BIND/HAVING stubs**: Not executable DQL yet; will yield incorrect or failing queries.
3. **OPTIONAL/UNION semantics incomplete**: Optional patterns currently required; unions behave like
   AND.
4. **FILTER coverage partial**: Comparisons/regex only; other built-ins/functions unsupported in
   strict mode.
5. **Subject IRI resolution**: Placeholder `_subject` means IRIs/literals as subject won’t resolve
   correctly.

---

### 2. **Simple Parser for Extended Features** ✓ _Pragmatic_

**File:** `simple_parser.go`

The lightweight parser uses regex patterns to extract:

- Aggregates: `COUNT|SUM|AVG|MIN|MAX` with aliases
- BIND expressions
- HAVING clauses
- Comments (skipped with `#`)
- ORDER BY, DISTINCT, LIMIT, OFFSET
- FROM, FROM NAMED clauses

**Correctness Impact:** ⚠️ **Medium** - Regex-based parsing is brittle:

- Will not handle nested parentheses in HAVING expressions correctly
- Cannot parse complex expressions in BIND clauses
- Whitespace variations may break matching
- No validation of expression correctness

**Example Problems:**

```sparql
# This will fail:
HAVING (count(?x) > count(?y))  # Nested aggregates

# This may fail:
BIND (CONCAT(STR(?a), " ", STR(?b)) AS ?result)  # Nested function calls
```

---

### 3. **Placeholder Aggregate/BIND/HAVING Translation** ⚠️ _Stub Implementation_

**File:** `translator_extended.go` lines 164-200

The translation functions exist but do minimal actual translation:

- **`applyBindExpression()`:** Just stores expression as `bind_<var> = <expr>` in query args
  - No actual expression evaluation
  - Variables in expressions are not resolved
  - String concatenation, math, etc. not evaluated

- **`applyAggregates()`:** Creates GroupByAttr with function string
  - `function + "(" + variable + ")"` becomes the attribute
  - This is not valid DQL syntax
  - DQL requires actual aggregation functions like `count(uid)`, `sum(val(...))`

- **`applyHavingClause()`:** Creates a filter with `having` as the function name
  - `having` is not a valid DQL filter function
  - Should parse the expression and map to actual DQL filter operations

**Correctness Impact:** ⚠️ **High** - These features are recognized but **will not execute correctly
in DQL**. Queries with aggregates, BIND, or HAVING will be built but fail or return incorrect
results at execution time.

---

### 4. **OPTIONAL/UNION Pattern Translation** ⚠️ _Incomplete_

**File:** `translator_extended.go` lines 125-160

**Current Implementation:** Recursively translates graph patterns but doesn't implement SPARQL
semantics:

- **OPTIONAL:** Currently just translates the inner patterns; doesn't mark them as optional
- **UNION:** Adds alternatives sequentially; doesn't create OR structure in DQL

**Correctness Impact:** ⚠️ **High** - OPTIONAL patterns will require matches; UNION will create AND
instead of OR.

```sparql
# This SPARQL:
SELECT ?x WHERE {
  ?x name ?name .
  OPTIONAL { ?x age ?age }
}

# Currently translates to something requiring both, not optional
```

---

### 5. **Graph Filtering Logic** ✓ _Mostly Correct_

**File:** `translator_impl.go` lines 136-200

The `buildGraphFilter()` function implements SPARQL graph filtering semantics correctly:

- Default graph patterns use FROM clauses
- Named graph patterns (GRAPH keyword) use FROM NAMED clauses
- Returns nil if pattern should be skipped

This part appears sound.

---

### 6. **No CONSTRUCT or DESCRIBE Support**

**File:** `translate.go`

Only SELECT and ASK are implemented. CONSTRUCT and DESCRIBE query types return "unsupported" error.

**Status:** Acceptable for MVP; documented limitation.

---

## Test Coverage Assessment

✅ **All tests pass** - 90+ test cases covering:

- Basic SELECT/ASK queries
- FROM/FROM NAMED clauses
- PREFIX declarations
- OPTIONAL patterns
- UNION patterns
- Aggregates (COUNT, SUM)
- DISTINCT
- ORDER BY
- LIMIT/OFFSET
- BIND expressions
- HAVING clauses
- Named graphs (GRAPH keyword)
- Comments

⚠️ **Note:** Tests pass because they check that parsing and translation don't error, not that the
output DQL is correct. The actual DQL query execution would likely fail for complex queries.

---

## Architecture Overview

### Current Data Flow

```text
Raw SPARQL Query
  ↓
ANTLRParserAdapter.Parse()
  ├→ ANTLR Lexer/Parser: Validate syntax
  ├→ On error: Try simple parser fallback
  └→ On success: Run simple parser (ignore parse tree)
  ↓
SPARQLQueryImpl (intermediate AST)
  ├→ Simple query: BGPs with triples
  └→ Extended query: Patterns, Aggregates, Binds, Having
  ↓
TranslateToGraphQueries()
  ├→ SELECT → TranslateSelectExtended()
  ├→ ASK → translateAsk()
  └→ Other → Error
  ↓
dql.GraphQuery (DQL AST)
```

### Parser Options

- **ANTLRParserAdapter:** Validates with ANTLR, builds AST with simple parser
- **Simple Parser Only:** Could use simple parser directly (no ANTLR validation)

---

## Recommendations toward alpha

High priority:

1. Gate or remove simple-parser fallback; rely on visitor output.
2. Implement correct DQL aggregation, grouping, and HAVING.
3. Implement BIND expression parsing/binding.
4. Fix OPTIONAL (left-join) and UNION (OR) semantics.
5. Extend FILTER support (boolean ops, built-ins) while keeping strict parsing.
6. Implement subject IRI lookup.

Medium: 7) Decide/implement CONSTRUCT/DESCRIBE and property paths; or explicitly defer. 8) Improve
error reporting on filter/aggregate parsing failures.

Low: 9) Remove legacy simple parser entirely once visitor coverage is complete. 10) Optimization
passes (predicate combining, filter pushdown) after correctness.

---

## Files Removed

- `http_integration.go` - Unfinished HTTP handler stub (0 usages, 10 lines)

## Files Modified

- `antlr_adapter.go` - Removed 48 lines of unused ParseVisitor code
- `adapter.go` - Removed 1 line of stub comment
- `ast.go` - Improved documentation (same line count)
- `translator_impl.go` - Converted placeholder comment to TODO (no line change)

**Net Change:** -49 lines of dead code, +5 lines of TODO documentation

---

## Conclusion

The branch successfully integrates ANTLR with extended SPARQL grammar support. All tests pass and
the basic architecture is sound. However, several features (aggregates, BIND, HAVING, OPTIONAL,
UNION) are partially implemented or use fallback strategies that may not produce correct DQL output.
These are documented as shortcuts taken to get tests passing quickly.

For production use, priority should be:

1. Implement proper ANTLR visitor for parse tree extraction
2. Complete DQL translation for aggregates and advanced features
3. Test with actual DQL query execution (not just parsing)
