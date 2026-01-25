# SPARQL Status – Single Source of Truth (2026-01-25)

Purpose: unify prior summaries and design decisions into one up-to-date view of what is implemented,
what changed, and what’s next.

## TL;DR

- **Stable**: Named graph plumbing via `graph` predicate; quad loader; FROM/FROM NAMED/GRAPH (fixed
  IRIs); SELECT/ASK over basic graph patterns with PREFIX + common FILTER; direct AST→DQL path.
- **Experimental**: OPTIONAL, UNION, aggregates, BIND, HAVING, ORDER BY, DISTINCT; property-path
  sequences/inverse; HAVING parsing; FILTER coverage beyond basics. Semantics are simplified and not
  schema- or auth-aware.
- **Not yet**: Variable GRAPH, full property paths (`*`, `+`), CONSTRUCT/DESCRIBE/UPDATE,
  SERVICE/federation, HTTP/SPARQL endpoint, auth/schema integration in planner.

## Implemented (current runtime path)

- Graph predicate model with quad loader emitting `graph` membership triples; unit-tested.
- Translator applies FROM/FROM NAMED/GRAPH filters on fixed IRIs for default and named graphs.
- SELECT/ASK over basic graph patterns; PREFIX resolution; FILTER with common functions
  (regex/comparisons/langMatches/boolean ops).
- Parser is ANTLR-generated and checked in; direct AST→DQL translation remains the execution path.

## Experimental / partial

- OPTIONAL and UNION are wired but use simplified semantics (not true left join/UNION planning).
- Aggregates, BIND, HAVING, ORDER BY, DISTINCT exist in AST/translator; parser coverage for HAVING
  and complex expressions is incomplete; execution is not schema-aware.
- Property paths: minimal support for simple sequences/inverse only.
- ANTLR visitor coverage gaps can trigger the legacy simple-parser fallback in some cases.

## Not implemented / out of scope right now

- Variable GRAPH bindings; full property paths with Kleene (`*`, `+`).
- CONSTRUCT, DESCRIBE, UPDATE; SERVICE/federation.
- HTTP/SPARQL endpoint wiring; auth and UID/IRI resolution in SPARQL path.

## Design decisions & changes

- **Graph predicate emulation**: Keep using a dedicated `graph` predicate to emulate named graphs in
  Dgraph; loader materializes graph membership per quad.
- **Direct AST→DQL today; move to Algebra**: Current translator bypasses the planned algebra layer;
  upcoming work will insert SPARQL Algebra IR for correct OPTIONAL/UNION semantics, optimizer
  passes, and schema/auth integration.
- **Parser strategy**: ANTLR grammar is canonical; simple-parser fallback is temporary and slated
  for removal once visitor coverage is complete (esp. HAVING/property paths/complex FILTER).
- **Optimization progress**: Phase 2B algebra optimizations landed (VALUES/EMPTY nodes, UNION
  deduplication, path desugaring, scope analysis, filter normalization). These live in the algebra
  code but are not yet on the main execution path because AST→DQL skips algebra.

## Recent milestones

- Phase 1 named-graph support: DONE (graph predicate model, loader, FROM/FROM NAMED/GRAPH filters).
- Phase 2 extended features: PARTIAL/EXPERIMENTAL (OPTIONAL/UNION/aggregates/BIND/HAVING wired, but
  semantics simplified and parser coverage incomplete).
- Phase 2B algebra optimizations: DONE in algebra codebase (see
  `OPTIMIZATION_IMPLEMENTATION_SUMMARY.md`).
- Documentation refresh (this doc, README, implementation/design docs) to reflect actual status.

## Next steps (priority)

1. Insert SPARQL Algebra layer into the execution path; route translator through it.
2. Make OPTIONAL/UNION semantics correct via schema-aware planning (true left joins/OR planning).
3. Extend ANTLR visitor: full HAVING/property paths/complex FILTER; eliminate simple-parser
   fallback.
4. Add service integrations: UID/IRI resolution + auth enforcement for SPARQL endpoint.
5. Implement full property paths (`*`, `+`) and validate CONSTRUCT/DESCRIBE/UPDATE coverage.
6. Refresh end-to-end tests against Dgraph with quad loader; update test matrix accordingly.

## Supersedes older summaries

For the authoritative snapshot, use this file. Superseded summaries have been removed from the repo
(available in git history if needed):

- `COMPLETION_SUMMARY.md` (Phase 2B completion)
- `OPTIMIZATION_IMPLEMENTATION_SUMMARY.md` (Phase 2B optimization details)
- `FINAL_SUMMARY.md` / `DELIVERY_SUMMARY.md` (earlier delivery inventories)
- `sparql/docs/DELIVERABLES_SUMMARY.md` and related phase summaries

## Pointers

- Status & features: `sparql/README.md`
- Implementation details: `sparql/docs/IMPLEMENTATION.md`, `sparql/docs/EXTENDED_FEATURES.md`
- Named-graph design: `docs/sparql-syntax-and-named-graph-design.md`
- Algebra/optimizer details: `sparql/algebra_*` (see git history for prior summaries)
