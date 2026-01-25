# SPARQL front-end (ANTLR → Go)

This package provides scaffolding to parse SPARQL 1.1 with an ANTLR-generated Go parser and
translate SPARQL queries into Dgraph's internal dql.GraphQuery AST.

## Overview

- Generate a Go parser from the SPARQL grammar (grammars-v4).
- Implement a visitor/listener that creates a small intermediate SPARQL AST.
- Translate that AST into []\*dql.GraphQuery using TranslateToGraphQueries.
- Hook into the server (e.g., register a /sparql endpoint) and reuse existing execution pipeline
  (ToSubGraph, Request.ProcessQuery).

## Architecture & Future Direction

The current implementation translates SPARQL AST directly to DQL queries. We are planning an
architectural evolution to introduce a **SPARQL Algebra** intermediate representation that will:

- Enable **structured query rewriting** (filter pushdown, join reordering)
- **Integrate schema information** for type-aware optimizations
- **Apply authorization rules** (currently missing for SPARQL)
- Establish foundation for **ontology reasoning**

See the following documents for detailed information:

- **[ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)** - High-level architectural vision and design
- **[SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md)** - Detailed 5-phase implementation plan
- **[PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md)** - Immediate next steps for Phase 1

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

## Documentation Index

For developers working on or understanding the SPARQL implementation:

**Current Status & Features**:

- **[IMPLEMENTATION.md](IMPLEMENTATION.md)** - Test coverage and features (current state)
- **[EXTENDED_FEATURES.md](EXTENDED_FEATURES.md)** - Advanced SPARQL features (OPTIONAL, UNION,
  aggregates, BIND, HAVING)
- **[TEST_COVERAGE.md](TEST_COVERAGE.md)** - Test inventory and coverage status

**Architecture & Future Development**:

- **[DECISION_RECORD.md](DECISION_RECORD.md)** - Architecture decision: why algebra-driven rewriting
- **[ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)** - Proposed architecture with SPARQL Algebra layer
  (detailed)
- **[SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md)** - Implementation roadmap (5 phases, 14 weeks)
- **[SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md)** - How SPARQL fits in Dgraph's overall
  architecture
- **[PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md)** - Week 1 implementation guide for
  Phase 1

**Quick Links**:

- Want to understand the big picture? Start with [DECISION_RECORD.md](DECISION_RECORD.md)
- Ready to implement? Read [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md)
- Need full details? See [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)
