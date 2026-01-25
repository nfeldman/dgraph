# SPARQL front-end (ANTLR → Go)

This package provides scaffolding to parse SPARQL 1.1 with an ANTLR-generated Go parser and
translate SPARQL queries into Dgraph's internal dql.GraphQuery AST.

Overview

- Generate a Go parser from the SPARQL grammar (grammars-v4).
- Implement a visitor/listener that creates a small intermediate SPARQL AST.
- Translate that AST into []\*dql.GraphQuery using TranslateToGraphQueries.
- Hook into the server (e.g., register a /sparql endpoint) and reuse existing execution pipeline
  (ToSubGraph, Request.ProcessQuery).

Development flow

1. Place `SPARQL.g4` (from https://github.com/antlr/grammars-v4/tree/master/sparql) into this
   directory.
2. Run `./generate_parser.sh` (this downloads or uses the local ANTLR jar) which will produce
   generated parser code in `sparql/gen/`.
3. Implement an ANTLR visitor that builds a SPARQLQuery implementation used by the translator.
4. Implement the translator in `translate.go` (map triple patterns, FILTER, OPTIONAL, UNION,
   ORDER/LIMIT).
5. Add tests and rely on CI to run the parser generation step for PR builds.

Current: We check in the generated parser under `sparql/gen/` for stable builds and reproducibility.

Supported subset (v1)

- SELECT / ASK
- Basic graph patterns (triple patterns)
- PREFIX resolution
- FILTER with common functions (regex, comparisons, langMatches, boolean ops)
- OPTIONAL, UNION, LIMIT, OFFSET, ORDER BY, DISTINCT
- Basic COUNT aggregation (basic GROUP BY)
- Property-path sequence and inverse (limited support)

Not in v1

- CONSTRUCT, DESCRIBE, UPDATE
- SERVICE/federation
- Full Kleene property paths (`*`, `+`) (planned for v2)
