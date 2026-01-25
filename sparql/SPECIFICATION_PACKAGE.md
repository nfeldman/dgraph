# SPARQL Algebra Project: Complete Specification Package

**Date**: January 2026  
**Project**: SPARQL→Algebra→DQL Pipeline with Schema-Aware Optimization  
**Status**: Specification Complete, Ready for Phase 1 Implementation

---

## What This Package Contains

A complete architectural specification and implementation roadmap for evolving SPARQL query
processing from a direct AST→DQL translator to a sophisticated algebra-based system with schema
awareness, query optimization, and authorization parity with GraphQL.

### Documents in This Package

| Document                                                 | Purpose                          | Audience                    | Read Time |
| -------------------------------------------------------- | -------------------------------- | --------------------------- | --------- |
| [README.md](README.md)                                   | Overview and documentation index | All                         | 5 min     |
| [DECISION_RECORD.md](DECISION_RECORD.md)                 | Why this architecture            | Architects, decision-makers | 10 min    |
| [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)             | Complete technical design        | Engineers, architects       | 30 min    |
| [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md)         | Detailed task breakdown          | Project leads, implementers | 20 min    |
| [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md)           | How SPARQL fits in Dgraph        | Senior engineers            | 25 min    |
| [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) | Week 1 implementation guide      | Phase 1 developers          | 15 min    |
| [IMPLEMENTATION.md](IMPLEMENTATION.md)                   | Current test coverage            | Existing (not new)          | 10 min    |
| [EXTENDED_FEATURES.md](EXTENDED_FEATURES.md)             | Current advanced features        | Existing (not new)          | 10 min    |

**Total Reading Time**: ~2-3 hours for complete understanding

---

## The Vision in 30 Seconds

Transform SPARQL from a simple AST→DQL translator into a **schema-aware query compiler** that:

1. **Translates** SPARQL to algebraic form (W3C standard)
2. **Optimizes** via algebraic rewriting (filter pushdown, join reordering)
3. **Compiles** with schema context (type inference, cardinality hints)
4. **Authenticates** with RBAC rules (like GraphQL already does)
5. **Executes** through existing Dgraph engine

**Result**: SPARQL gets the same sophistication as GraphQL queries.

---

## Quick Start Guide

### For Architects / Decision-Makers

1. Read [DECISION_RECORD.md](DECISION_RECORD.md) - Why this architecture (10 min)
2. Scan [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - Overview sections (10 min)
3. Review [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) - Timeline and phases (5 min)
4. Decision: Approve or request changes

### For Phase 1 Implementers

1. Read [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) - Complete (15 min)
2. Study [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - Algebra section (15 min)
3. Examine [graphql/resolve/query_rewriter.go](../graphql/resolve/query_rewriter.go) - Reference
   pattern (20 min)
4. Review [dql/parser.go](../dql/parser.go) - Target format (15 min)
5. Start implementing Phase 1: algebra.go

### For Project Leads

1. Read [DECISION_RECORD.md](DECISION_RECORD.md) - Context (10 min)
2. Read [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) - Complete (20 min)
3. Review [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md) - Integration points (20 min)
4. Use [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) as project tracking document

---

## Key Concepts

### SPARQL Algebra

Formal intermediate representation of SPARQL semantics (W3C standard):

- **Operators**: BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, etc.
- **Properties**: Composable, rewritable, semantics-preserving
- **Value**: Enables query optimization before execution

### Query Rewriting

Transforming algebra expressions to improve execution efficiency:

- **Filter Pushdown**: Move WHERE constraints closer to source data
- **Join Reordering**: Arrange joins by estimated cardinality (cheap first)
- **OPTIONAL Simplification**: Convert unnecessary LeftJoins to regular Joins
- **Result**: 30%+ improvement in query plan complexity

### Schema Integration

Using Dgraph's schema information in query compilation:

- **Type Inference**: Detect which types predicates apply to
- **Cardinality Hints**: Use schema statistics to guide optimization
- **Validation**: Catch invalid queries at compile time (not runtime)
- **Authorization**: Apply type-level RBAC rules (currently missing)

---

## Implementation Timeline

```
Week 1-3    Phase 1: Algebra Foundation ✅ Ready to start
            • Algebra operators and types
            • AST→Algebra conversion
            • Algebra validator
            • 100+ tests

Week 4-6    Phase 2: Query Optimization
            • Filter pushdown rewriter
            • Join reordering
            • OPTIONAL simplification
            • Dead variable elimination

Week 7-8    Phase 3: Schema Integration
            • Schema-aware DQL compiler
            • Type inference
            • Direct DQL Result construction

Week 9-10   Phase 4: Authorization
            • Auth rule application
            • Parity with GraphQL

Week 11-14  Phase 5: Ontology Foundation (can parallel)
            • Ontology loading (OWL/RDFS)
            • Schema synthesis
            • RDFS reasoning rules

Total: ~14 weeks (can parallelize Phase 5)
```

---

## Success Metrics

| Metric                         | Target                            | Phase |
| ------------------------------ | --------------------------------- | ----- |
| Algebra operators implemented  | All W3C operators                 | 1     |
| Test cases                     | 100+                              | 1     |
| AST→Algebra coverage           | 100% of SPARQL patterns           | 1     |
| Query optimization improvement | 30%+ reduction in plan complexity | 2     |
| Schema integration             | Type info attached to all nodes   | 3     |
| Performance regression         | <5% overhead                      | 3     |
| Authorization parity           | Same behavior as GraphQL          | 4     |
| Ontology reasoning             | RDFS rules functional             | 5     |

---

## Risk Mitigation

| Risk                   | Probability | Impact | Mitigation                         |
| ---------------------- | ----------- | ------ | ---------------------------------- |
| Phase 1 complexity     | Low         | High   | Break into PRs; start minimal      |
| Integration issues     | Medium      | High   | Comprehensive tests; feature flags |
| Performance regression | Low         | High   | Benchmark before/after             |
| Schema dependency      | Low         | Medium | Graceful degradation               |
| Auth complexity        | Medium      | Medium | Study existing auth rules          |

---

## How to Use This Package

### As a Developer (Implementing Phase 1)

1. Read [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md)
2. Create `sparql/algebra.go` with types and conversion logic
3. Reference [dql/parser.go](../dql/parser.go) for GraphQuery format
4. Study [graphql/resolve/query_rewriter.go](../graphql/resolve/query_rewriter.go) for optimization
   patterns
5. Write 100+ comprehensive tests
6. Use [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) to track progress

### As a Project Lead

1. Use [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) as master task list
2. Track phases and milestones
3. Monitor success criteria
4. Manage dependencies between phases
5. Report progress weekly

### As an Architect

1. Review [DECISION_RECORD.md](DECISION_RECORD.md) and [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)
2. Validate assumptions with team
3. Identify risks and blockers
4. Plan integration points with existing systems
5. Review [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md) for impact assessment

### As a Reviewer / Validator

1. Check against [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) design
2. Verify test coverage (target: 80%+)
3. Confirm no regressions in existing tests
4. Performance benchmarks within tolerance
5. Code quality standards maintained

---

## Key Integration Points

### REST Layer (edgraph/server.go)

- Detect SPARQL queries
- Create execution context with schema
- Call SPARQL compilation pipeline
- Construct dql.Result directly (bypass re-parsing)

### Authorization (edgraph/access.go)

- Works as-is with pre-injected auth rules
- No changes needed
- SPARQL queries now get same auth treatment as GraphQL

### Query Execution (query/query.go)

- SPARQL queries produce []\*dql.GraphQuery
- Seamless integration with ToSubGraph/ProcessGraph
- No changes needed to execution engine

---

## Deliverables by Phase

### Phase 1: Foundation

- [ ] `sparql/algebra.go` (~500 LOC)
- [ ] `sparql/context.go` (~100 LOC)
- [ ] `sparql/algebra_test.go` (~400 LOC, 100+ tests)
- [ ] Zero regressions in existing tests

### Phase 2: Optimization

- [ ] `sparql/algebra_rewriter.go` (~400 LOC)
- [ ] `sparql/algebra_optimizer_test.go` (~200 LOC, 150+ tests)
- [ ] 30%+ improvement in query plans

### Phase 3: Schema Integration

- [ ] `sparql/compiler_dql.go` (~400 LOC)
- [ ] `sparql/compiler_dql_test.go` (~200 LOC, 200+ tests)
- [ ] Modified `edgraph/server.go` (~50 LOC)
- [ ] Zero performance regression

### Phase 4: Authorization

- [ ] `sparql/auth_rules.go` (~300 LOC)
- [ ] `sparql/auth_rules_test.go` (~150 LOC, 150+ tests)
- [ ] Parity testing with GraphQL

### Phase 5: Ontology

- [ ] `sparql/ontology/loader.go` (~300 LOC)
- [ ] `sparql/ontology/schema_builder.go` (~250 LOC)
- [ ] `sparql/ontology/reasoning.go` (~350 LOC)
- [ ] Complete test suite

---

## Open Questions & Decisions

### Design Decisions Made

✅ Use W3C SPARQL Algebra (not custom form) ✅ Phase-by-phase implementation (not all-at-once) ✅
Direct DQL construction (not re-parsing) ✅ Leverage GraphQL schema (not separate schema) ✅
Standard auth rules (not SPARQL-specific auth)

### Decisions Pending

- [ ] Feature flag for new path (A/B testing)?
- [ ] Metrics/observability strategy?
- [ ] Fallback behavior for errors?
- [ ] Optimization rule prioritization?

---

## Getting Help

### Understanding the Architecture

- [DECISION_RECORD.md](DECISION_RECORD.md) - Why this design
- [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - How it works
- [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md) - Where it fits

### Getting Started with Implementation

- [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) - Week 1 guide
- [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) - Detailed tasks

### Reference Code

- [graphql/resolve/query_rewriter.go](../graphql/resolve/query_rewriter.go) - Query rewriting
  pattern
- [dql/parser.go](../dql/parser.go) - Target GraphQuery format
- [query/query.go](../query/query.go) - Query execution
- [worker/task.go](../worker/task.go) - Task planning

---

## Document Evolution

This specification package will be updated as:

- Phases complete and learnings emerge
- Integration points are validated
- Risks materialize and mitigations apply
- Future extensions are designed

**Current Version**: 1.0  
**Last Updated**: January 2026  
**Next Review**: After Phase 1 completion

---

## Executive Summary

**What**: Introduce SPARQL Algebra intermediate representation with schema-aware optimization  
**Why**: Close performance and security gaps vs. GraphQL; foundation for future ontology reasoning  
**When**: 14 weeks, 5 phases, can parallelize  
**Who**: Core team + architecture review  
**How**: Phase-by-phase implementation with comprehensive testing and integration validation  
**Impact**: SPARQL queries get same sophistication as GraphQL queries

**Ready to proceed with Phase 1?** ✅
