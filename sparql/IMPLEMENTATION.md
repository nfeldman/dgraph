# SPARQL Named Graph Implementation Summary

## What Was Implemented

This implementation adds support for SPARQL named graph queries in Dgraph, following the design
specified in `docs/sparql-syntax-and-named-graph-design.md`.

### Core Components

#### 1. Interface Extensions ([sparql/adapter.go](sparql/adapter.go))

- Added `FROM()` and `FROMNamed()` methods to `SPARQLQuery` interface
- Enables translators to access dataset specification from queries

#### 2. AST Structures ([sparql/ast.go](sparql/ast.go))

- `Triple.Graph`: Tracks which named graph a triple pattern belongs to
- `sparqlQueryImpl.from`: Stores FROM clause graph IRIs
- `sparqlQueryImpl.fromNamed`: Stores FROM NAMED clause graph IRIs
- `SPARQLQueryImpl`: Exported version for testing and examples

#### 3. Translator ([sparql/translator_impl.go](sparql/translator_impl.go))

- **translateSelect()**: Converts SPARQL SELECT to DQL GraphQuery
- **translateAsk()**: Converts SPARQL ASK to DQL GraphQuery
- **buildGraphFilter()**: Core logic implementing SPARQL graph semantics:
  - For GRAPH patterns: Adds `@filter(eq(graph, "g"))`
  - For default graph with FROM: Adds `@filter(eq(graph, "g1") OR eq(graph, "g2") ...)`
  - Validates GRAPH patterns against FROM NAMED list
  - Returns nil for patterns that should be skipped

#### 4. RDF Quad Loader ([sparql/loader.go](sparql/loader.go))

- `QuadLoader`: Processes N-Quads format
- For each quad `<s> <p> <o> <g>`, generates:
  1. Data triple: `<s> <p> <o>`
  2. Graph membership triple: `<s> <graph> "g"`
- `ToMutationJSON()`: Converts triples to Dgraph mutation format

### Test Coverage

#### Unit Tests ([sparql/translate_test.go](sparql/translate_test.go))

1. **TestBuildGraphFilter_NamedGraph**: GRAPH pattern filtering
   - Matches when in FROM NAMED
   - Skips when not in FROM NAMED
   - Works without FROM NAMED specified

2. **TestBuildGraphFilter_DefaultGraph**: Default graph filtering
   - Single FROM clause
   - Multiple FROM clauses (OR logic)
   - No FROM clause (match all)

3. **TestTranslateSelect_WithGraphFilters**: Full query translation
   - FROM clause handling
   - FROM NAMED + GRAPH combination
   - Multiple FROM graphs

4. **TestTranslateAsk**: ASK query translation
5. **TestBuildOrFilter**: OR filter tree construction
6. **TestQuoteString**: String quoting utility

#### Examples ([sparql/examples_test.go](sparql/examples_test.go))

1. **Example_quadLoading**: Demonstrates N-Quads loading
2. **Example_sparqlTranslation**: Shows query translation workflow
3. **Example_namedGraphQuery**: Full documentation example

## SPARQL Semantics Implemented

### FROM Clause

```sparql
FROM <http://g1> FROM <http://g2>
```

- Defines the "default graph" as union of g1 and g2
- Triple patterns outside GRAPH blocks get: `@filter(eq(graph, "g1") OR eq(graph, "g2"))`

### FROM NAMED Clause

```sparql
FROM NAMED <http://g3>
```

- Makes g3 available for use in GRAPH patterns
- GRAPH patterns referencing graphs not in FROM NAMED are skipped

### GRAPH Patterns

```sparql
GRAPH <http://g3> { ?s ?p ?o }
```

- Triple patterns get: `@filter(eq(graph, "http://g3"))`
- Only works if g3 is in FROM NAMED (or no FROM NAMED specified)

### Default Behavior

- No FROM, no FROM NAMED: All graphs are accessible
- Patterns outside GRAPH blocks match all graphs

## Data Model

For RDF quad: `<s> <p> <o> <g>`

Store as:

```
<s> <p> <o> .
<s> <graph> "g" .
```

The `graph` predicate enables filtering by graph membership in DQL queries.

## Files Changed/Created

### New Files

- `sparql/translator_impl.go` - Core translation logic
- `sparql/loader.go` - RDF quad loader
- `sparql/examples_test.go` - Usage examples

### Modified Files

- `sparql/adapter.go` - Added FROM/FROM NAMED accessors
- `sparql/ast.go` - Added exported SPARQLQueryImpl
- `sparql/translate_test.go` - Comprehensive test suite
- `docs/sparql-syntax-and-named-graph-design.md` - Updated roadmap and implementation details

## Next Steps (Not Yet Implemented)

1. **Variable Graph Patterns**: `GRAPH ?g { ... }` binding
2. **ANTLR Visitor Integration**: Parse actual SPARQL text
3. **Full Triple Translation**: Proper subject/object handling with UID resolution
4. **PREFIX Expansion**: Expand prefixed names to full IRIs
5. **OPTIONAL, UNION**: Complex graph pattern operators
6. **FILTER Expressions**: Boolean constraints on variables
7. **Integration Testing**: End-to-end tests with actual Dgraph instance

## Usage Example

```go
// 1. Load RDF quads with graph predicates
loader := sparql.NewQuadLoader()
triples, _ := loader.LoadQuads(nquadsReader)

// 2. Create/Parse SPARQL query
query := &sparql.SPARQLQueryImpl{
    Qtype: "SELECT",
    Projs: []string{"s", "name"},
    From: []string{"http://g1"},
    FromNamed: []string{"http://g2"},
    Bgps: []*sparql.BGP{{
        Triples: []*sparql.Triple{{
            Subject: "?s",
            Predicate: "ex:name",
            Object: "?name",
            ObjectIsVar: true,
            Graph: "", // default graph - will filter on g1
        }},
    }},
}

// 3. Translate to DQL
gqs, prefixes, _ := sparql.TranslateToGraphQueries(ctx, query, opts)

// 4. gqs now contains DQL with @filter(eq(graph, "http://g1"))
```

## Architecture Alignment

This implementation follows the design document:

- ✅ Graph predicate emulation
- ✅ FROM/FROM NAMED dataset specification
- ✅ Filter generation per SPARQL semantics
- ✅ Loader example for quad ingestion
- ✅ Comprehensive test coverage

The foundation is now in place for full SPARQL 1.1 support with named graphs.
