# SPARQL front-end (ANTLR → Go)

This package provides scaffolding to parse SPARQL 1.1 with an ANTLR-generated Go parser and
translate SPARQL queries into Dgraph's internal dql.GraphQuery AST.

## Overview

- ANTLR-generated Go parser for SPARQL 1.1 (checked in under `sparql/gen/`).
- Visitor builds a small SPARQL AST (`SPARQLQueryImpl`).
- Translator maps that AST into `[]*dql.GraphQuery` for execution in Dgraph.
- Named graph support is implemented by storing the graph IRI in a `graph` predicate and filtering
  via FROM / FROM NAMED / GRAPH.

## Current Status (2026-01)

- **Implemented and exercised**: SELECT/ASK over basic graph patterns; FILTER (common functions),
  PREFIX expansion; dataset scoping via FROM/FROM NAMED/GRAPH on fixed IRIs; quad loader that
  materializes `graph` membership.
- **Present but partial/experimental**: OPTIONAL, UNION, aggregates, BIND, HAVING, ORDER BY,
  DISTINCT are wired in the AST/translator but semantics are simplified and not validated against
  the full SPARQL spec or schema-aware execution. Property paths are minimal (no Kleene).
- **Not implemented**: CONSTRUCT, DESCRIBE, UPDATE, SERVICE/federation, variable GRAPH, full
  property paths (`*`, `+`).
- **Planned**: SPARQL Algebra IR to enable correct OPTIONAL semantics, schema/authorization
  integration, and optimizer passes. ANTLR visitor updates to cover HAVING and property paths.

## Architecture & Future Direction

Today the pipeline is direct AST → DQL translation. The intended evolution inserts a SPARQL Algebra
layer to enable rewriting, schema-awareness, and auth enforcement before DQL generation. See:

- **[ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)** – High-level architectural vision
- **[SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md)** – 5-phase implementation plan
- **[PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md)** – Immediate next steps for Phase 1

Development flow

1. Place `SPARQL.g4` (from https://github.com/antlr/grammars-v4/tree/master/sparql) into this
   directory.
2. Run `./generate_parser.sh` to regenerate `sparql/gen/` (parser is committed for reproducibility).
3. Implement/extend the ANTLR visitor to populate `SPARQLQueryImpl` (HAVING and property paths need
   coverage).
4. Improve the translator to route through the planned algebra layer; today it directly maps AST →
   DQL.
5. Add/maintain tests; CI should regenerate the parser for PR builds.

Supported today (stable path)

- SELECT / ASK
- Basic graph patterns (triple patterns)
- PREFIX resolution
- FILTER with common functions (regex, comparisons, langMatches, boolean ops)
- FROM / FROM NAMED / GRAPH (static IRI) via `graph` predicate filtering

Experimental / partial

- OPTIONAL, UNION, aggregates, BIND, HAVING, ORDER BY, DISTINCT (translator hooks exist; semantics
  need validation and schema-aware execution)
- Property-path sequence and inverse (limited)

Out of scope currently

- CONSTRUCT, DESCRIBE, UPDATE
- SERVICE / federation
- Full Kleene property paths (`*`, `+`)

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
