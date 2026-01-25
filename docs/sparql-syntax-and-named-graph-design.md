# SPARQL Support and Named Graph Emulation in Dgraph

## Overview

This document describes the architecture, translation rules, and data modeling approach for
supporting SPARQL queries—including named graph constructs (`GRAPH`, `FROM`, `FROM NAMED`)—in Dgraph
using ANTLR-generated parser and predicate emulation.

## Current status (2026-01-25)

- **Implemented:**
  - ANTLR visitor builds `SPARQLQueryImpl` for `SELECT`/`ASK`, parsing prefixes, projections,
    DISTINCT, FROM / FROM NAMED, GRAPH blocks, OPTIONAL, UNION, BIND, aggregates
    (COUNT/SUM/MIN/MAX/AVG with DISTINCT), HAVING, ORDER BY, LIMIT/OFFSET, FILTER, and
    graph-annotated BGPs.
  - Quad loader emits graph membership triples and is covered by extensive tests.
  - Translator supports graph filtering semantics, OPTIONAL/UNION, aggregates, BIND, HAVING,
    DISTINCT, ORDER, LIMIT/OFFSET, and now maps FILTER comparisons/regex to structured DQL filters
    (raw fallback only when explicitly allowed via options); examples and e2e tests are green.
- **Shortcuts/deviations:**
  - `antlr_adapter` still falls back to the legacy simple parser if ANTLR rejects input **or** if
    the visitor yields no patterns; this bypasses the AST in those cases.
  - FILTER translation covers comparisons/regex; complex expressions error in strict mode (default)
    and only fall back to raw when `AllowFilterFallback` is enabled.
  - HTTP integration stub (`http_integration.go`) remains unimplemented.
- **Remaining stubs before wider usability:**
  - Refine FILTER translation to structured DQL functions and expression mapping.
  - Remove/replace the simple-parser fallback once visitor coverage is complete.
  - Add support for CONSTRUCT/DESCRIBE and property paths as planned.

## Goals

- Parse full SPARQL 1.1 queries, using an ANTLR4-generated Go parser (checked-in generated code).
- Translate SPARQL patterns—including named graphs—to DQL, using a dedicated Dgraph predicate to
  emulate graph partitioning.
- Implement compliance with `GRAPH`, `FROM`, and `FROM NAMED` as required by the SPARQL standard.
- Enable ingestion and querying of RDF quads, allowing graph/context-aware results.

## Data Model

- Every RDF triple ingested from a quad includes an extra triple to define the graph/context:

  ```turtle
  <subject> <predicate> <object> .
  <subject> <graph> "graph-iri" .
  ```

- The `graph` predicate indicates named graph membership.
- Subjects may have multiple `<graph>` predicates if they reside in multiple named graphs.

## Query Translation

### 1. FROM clause

- Specifies which graphs are combined to form the “default graph” (the union of the listed graphs).
- All triple patterns **outside `GRAPH` blocks** are filtered by
  `@filter(or(eq(graph, "g1"), eq(graph, "g2"), ...))`, where the set is the FROM IRIs.

### 2. FROM NAMED clause

- Specifies which named graphs are available for use in `GRAPH <g>` blocks.
- Triple patterns inside `GRAPH <g>` blocks are only matched if `<g>` is declared in `FROM NAMED`.

### 3. GRAPH patterns

- Patterns inside `GRAPH <g>` blocks:
  - Emit DQL query blocks with `@filter(eq(graph, "<g>"))`
  - Skip these blocks if `<g>` isn’t in the FROM NAMED set.

- Patterns inside `GRAPH ?g` blocks:
  - Only bind variables to graph IRIs in the FROM NAMED set.

- Patterns outside any `GRAPH` block:
  - Matched against the union of graphs specified by `FROM`.

### 4. Default handling

- If neither FROM nor FROM NAMED is specified:
  - Patterns outside GRAPH blocks are matched against all available graphs (default: all data).
  - Patterns in GRAPH blocks match any graph in the dataset.

## Example Translations

**SPARQL**

```sparql
SELECT * FROM <g1> FROM <g2> FROM NAMED <g3>
WHERE {
  ?s <p> ?o .
  GRAPH <g3> { ?x <q> ?y }
  GRAPH <g4> { ?z <r> ?w }
}
```

**DQL**

- Outside GRAPH: filter on `graph in {"g1", "g2"}`
- GRAPH <g3> block: filter on `graph == "g3"`
- GRAPH <g4> block: skipped (not in FROM NAMED)

## Translator Design

- Parse and store FROM / FROM NAMED specification in query context.
- Annotate each triple/BGP in AST with its graph context.
- Translator emits DQL blocks with proper filter logic for default and named graphs.
- Strict compliance: only named graphs in FROM NAMED are queryable with GRAPH.
- Variable graph blocks only expose graphs declared in FROM NAMED.

## Loader Requirements

- Data ingestion must tag quads/triples so every triple includes the graph predicate.
- Subjects in multiple graphs must have multiple `graph` edges.

## Implementation Roadmap

### Phase 1: Core Implementation (COMPLETED ✅)

- [x] Integrate ANTLR-generated parser checked in to source for reproducibility.
- [x] Extend AST and visitor to support FROM, FROM NAMED, and GRAPH annotations.
- [x] Translator logic emitting filter expressions for graph partitions per SPARQL rules.
- [x] Example logic and code snippets for loader and test cases.
- [x] Unit and integration tests for representative SPARQL queries and DQL translations.
- [x] Comprehensive test coverage for loader, translator, and utilities.

### Phase 2: Documentation & Validation (COMPLETED ✅, ongoing doc refresh)

- [x] Document implementation details and core files.
- [x] Document edge cases and limitations.
- [x] Create TEST_COVERAGE.md with detailed test matrix (110+ test cases).
- [x] Create IMPLEMENTATION.md with architecture and usage guide.
- [ ] Continue refreshing docs as ANTLR visitor replaces simple-parser fallback and FILTER support
      lands.

### Phase 3: Future Enhancements (PLANNED)

- [ ] Complete ANTLR visitor coverage (FILTER expressions, CONSTRUCT/DESCRIBE, property paths).
- [ ] Reduce/remove legacy simple-parser fallback once visitor covers all patterns.
- [ ] Variable graph patterns (GRAPH ?g) support.
- [ ] Performance optimizations and query planning.
- [ ] Integration with actual Dgraph instance.

### Test Coverage Summary

- **Unit Tests**: 110+ test cases
- **Test Files**: 3 (translate_test.go, loader_test.go, examples_test.go)
- **Lines of Test Code**: 700+
- **Coverage Matrix**: See [TEST_COVERAGE.md](../sparql/TEST_COVERAGE.md)

## Implementation Details

### Additional Documentation

- **[IMPLEMENTATION.md](../sparql/IMPLEMENTATION.md)**: Detailed implementation summary with
  architecture, files changed, and usage examples.
- **[TEST_COVERAGE.md](../sparql/TEST_COVERAGE.md)**: Comprehensive test coverage matrix and test
  categories.

### Core Files

- **sparql/adapter.go**: Defines `SPARQLParser` and `SPARQLQuery` interfaces with FROM/FROM NAMED
  accessors
- **sparql/ast.go**: Triple, BGP, and query implementation structures
- **sparql/translator_impl.go**: Core translation logic with graph filtering
  - `translateSelect()`: SELECT query translation
  - `translateAsk()`: ASK query translation
  - `buildGraphFilter()`: Generates DQL filters based on SPARQL graph semantics
- **sparql/loader.go**: RDF quad loader with graph predicate emission
- **sparql/translate_test.go**: Comprehensive unit tests
- **sparql/examples_test.go**: Usage examples

### Translation Examples

```go
// SPARQL Query
query := &sparql.SPARQLQueryImpl{
    From: []string{"http://g1"},
    Bgps: []*sparql.BGP{{
        Triples: []*sparql.Triple{{
            Subject: "?s",
            Predicate: "ex:name",
            Object: "?o",
            Graph: "", // default graph
        }},
    }},
}

// Generated DQL Filter
// @filter(eq(graph, "http://g1"))
```

### Loader Usage

```go
loader := sparql.NewQuadLoader()
triples, _ := loader.LoadQuads(reader)
// Each quad <s> <p> <o> <g> becomes:
//   1. <s> <p> <o>
//   2. <s> <graph> "g"
```

## Limitations

- Emulation relies on convention—no enforcement or optimizations beyond filtering.
- No support for true dataset operations or cross-namespace federation.
- Performance and correctness depends on loader/preprocessor emitting tags correctly.
- Triple translation is simplified - full predicate path resolution not yet implemented.
- Variable graph patterns (GRAPH ?g) not yet supported.

## Keeping This Wiki Updated

- To update this document as plans or code change, ask Copilot for the latest
  “sparql-syntax-and-named-graph-design.md” reflecting all recent changes, discussions, and
  requirements.
- All changes will be provided as full file contents, ready to commit.
