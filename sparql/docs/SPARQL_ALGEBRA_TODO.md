# SPARQL Algebra & Schema Integration: Implementation TODO

**Epic**: Bridge SPARQL and GraphQL query execution paths with shared schema-aware optimization  
**Milestone**: Q1/Q2 2026  
**Owner**: Dgraph Team

---

## Overview

Implement a SPARQL→Algebra→DQL pipeline that:

1. Converts SPARQL AST to algebraic form for query rewriting
2. Integrates schema information during compilation
3. Applies the same authorization rules as GraphQL queries
4. Eliminates redundant DQL re-parsing
5. Provides foundation for future ontology reasoning

See [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) for detailed design.

---

## Phase 1: SPARQL Algebra Foundation (Weeks 1-3)

### 1.1 Design & Core Types

- [ ] Define algebraic expression types (Join, LeftJoin, Union, Filter, Project, Aggregate, Bind,
      etc.)
  - File: `sparql/algebra.go`
  - Based on W3C SPARQL algebra specification
  - Support visitor pattern for traversal
  - ~200 LOC

- [ ] Create visitor pattern infrastructure
  - File: `sparql/algebra_visitor.go`
  - Interfaces: `AlgebraVisitor`, `AlgebraExpr`
  - Support both top-down and bottom-up traversal
  - ~100 LOC

### 1.2 AST → Algebra Conversion

- [ ] Implement `astToAlgebra()` converter
  - File: `sparql/algebra.go` (extension)
  - Convert BGP patterns to algebra
  - Handle FILTER placement
  - Transform graph patterns (OPTIONAL, UNION)
  - Process variables and projections
  - ~400 LOC

- [ ] Unit tests for astToAlgebra
  - File: `sparql/algebra_test.go`
  - Test each algebra operator creation
  - Verify variable tracking
  - Pattern combination logic
  - ~300 test cases

### 1.3 Algebra Validator

- [ ] Build `AlgebraValidator`
  - File: `sparql/algebra.go` (extension)
  - Validate variable consistency
  - Check domain/range constraints
  - Detect free variables
  - ~150 LOC

- [ ] Validator test suite
  - File: `sparql/algebra_test.go` (extension)
  - Valid and invalid algebra expressions
  - Error reporting clarity
  - ~100 test cases

### 1.4 Integration Point

- [ ] Create `SPARQLExecutionContext`
  - File: `sparql/context.go` (NEW)
  - Hold schema, auth, namespace
  - Pass through compilation pipeline
  - ~100 LOC

---

## Phase 2: Query Optimization (Weeks 4-6)

### 2.1 Algebra Rewriter Framework

- [ ] Build `AlgebraOptimizer` infrastructure
  - File: `sparql/algebra_rewriter.go` (NEW)
  - Rewriting rule interface
  - Fixed-point iteration
  - Profiling hooks
  - ~200 LOC

### 2.2 Filter Pushdown

- [ ] Implement filter pushdown rewriter
  - File: `sparql/algebra_rewriter.go` (extension)
  - Move filters toward source patterns
  - Combine filter conditions
  - ~150 LOC

- [ ] Filter pushdown test suite
  - File: `sparql/algebra_optimizer_test.go` (NEW)
  - Verify filter movement
  - Constraint preservation
  - Join compatibility
  - ~150 test cases

### 2.3 Join Reordering

- [ ] Implement cardinality-based join ordering
  - File: `sparql/algebra_rewriter.go` (extension)
  - Use schema stats for cardinality
  - Greedy join ordering
  - Handle bushy vs. left-linear
  - ~200 LOC

- [ ] Join reordering tests
  - File: `sparql/algebra_optimizer_test.go` (extension)
  - Reordering correctness
  - Cardinality estimates
  - Complex join chains
  - ~100 test cases

### 2.4 OPTIONAL Simplification

- [ ] Simplify OPTIONAL patterns
  - File: `sparql/algebra_rewriter.go` (extension)
  - Convert to LeftJoin where unnecessary
  - Merge compatible OPTIONALs
  - ~100 LOC

- [ ] OPTIONAL tests
  - File: `sparql/algebra_optimizer_test.go` (extension)
  - Simplification correctness
  - Result preservation
  - ~50 test cases

### 2.5 Dead Variable Elimination

- [ ] Remove unused variables
  - File: `sparql/algebra_rewriter.go` (extension)
  - Identify unreachable projections
  - Eliminate unused bindings
  - ~100 LOC

- [ ] DVE tests
  - File: `sparql/algebra_optimizer_test.go` (extension)
  - ~50 test cases

### 2.6 Integration Tests

- [ ] End-to-end optimization tests
  - File: `sparql/algebra_optimizer_test.go` (extension)
  - Combine multiple optimizations
  - Verify no regressions
  - Performance benchmarks
  - ~100 test cases

---

## Phase 3: Schema-Aware DQL Compilation (Weeks 7-8)

### 3.1 Schema Context Loading

- [ ] Load schema in SPARQL execution path
  - File: `sparql/context.go` (extension)
  - Extract from `schema.State()`
  - Build predicate→type mappings
  - Cache schema info
  - ~100 LOC

- [ ] Schema loading tests
  - File: `sparql/compiler_dql_test.go` (NEW)
  - Various schema configurations
  - ~50 test cases

### 3.2 DQL Compiler with Schema

- [ ] Implement schema-aware DQL compiler
  - File: `sparql/compiler_dql.go` (NEW)
  - Algebra → DQL GraphQuery with type info
  - Validate predicates against schema
  - Attach type metadata
  - Optimize based on cardinality hints
  - ~300 LOC

- [ ] Type inference from patterns
  - File: `sparql/compiler_dql.go` (extension)
  - Infer subject types from predicates
  - Identify implicit type filters
  - Build type hierarchies
  - ~150 LOC

### 3.3 DQL Compilation Tests

- [ ] Compiler test suite
  - File: `sparql/compiler_dql_test.go` (extension)
  - Algebra → DQL conversion
  - Type metadata attachment
  - Schema validation
  - ~200 test cases

### 3.4 Direct DQL Result Construction

- [ ] Avoid re-parsing optimization
  - File: `sparql/compiler_dql.go` (extension) + `query/query.go` (modification)
  - Construct `dql.Result` directly from compiled queries
  - Skip `dql.ParseWithNeedVars` for SPARQL
  - ~50 LOC

- [ ] Integration with edgraph
  - File: `edgraph/server.go` (modification)
  - Detect SPARQL query path
  - Call new compilation pipeline
  - Pass Result to execution
  - ~100 LOC

- [ ] Integration tests
  - File: `sparql/compiler_dql_test.go` (extension)
  - End-to-end with query execution
  - Performance comparison
  - ~50 test cases

### 3.5 Regression Tests

- [ ] Verify existing SPARQL tests still pass
  - File: `sparql/e2e_test.go` (modification)
  - All existing test cases should work
  - No behavior changes
  - ~200+ existing tests

---

## Phase 4: Authorization Rule Integration (Weeks 9-10)

### 4.1 SPARQL Auth Rules Module

- [ ] Create `ApplyAuthRulesToQuery()` function
  - File: `sparql/auth_rules.go` (NEW)
  - Similar to GraphQL's `addCommonRules`
  - Detect primary types from query
  - Apply RBAC rules
  - ~200 LOC

- [ ] Type inference for auth
  - File: `sparql/auth_rules.go` (extension)
  - Infer type context from DQL queries
  - Handle ambiguous types
  - ~100 LOC

### 4.2 Auth Rule Injection

- [ ] Implement static rule evaluation
  - File: `sparql/auth_rules.go` (extension)
  - Call auth rewriter functions
  - Attach auth variables
  - ~100 LOC

- [ ] Add auth filter queries
  - File: `sparql/auth_rules.go` (extension)
  - Generate @cascade queries
  - Combine with existing filters
  - ~150 LOC

### 4.3 Authorization Tests

- [ ] Auth rule application tests
  - File: `sparql/auth_rules_test.go` (NEW)
  - Type detection correctness
  - Rule evaluation
  - Filter injection
  - ~100 test cases

- [ ] Parity tests with GraphQL
  - File: `sparql/auth_rules_test.go` (extension)
  - Compare SPARQL vs. GraphQL auth behavior
  - Same query in both languages
  - Verify identical results
  - ~50 test cases

### 4.4 Integration with Execution

- [ ] Hook auth rules into compilation pipeline
  - File: `sparql/compiler_dql.go` (modification)
  - Call `ApplyAuthRulesToQuery()` in pipeline
  - Handle auth failures gracefully
  - ~50 LOC

- [ ] End-to-end auth tests
  - File: `sparql/auth_rules_test.go` (extension)
  - Full query execution with auth
  - Permission denied scenarios
  - Multi-user scenarios
  - ~100 test cases

---

## Phase 5: Ontology Foundation (Weeks 11-14)

### 5.1 Ontology Data Model

- [ ] Design ontology model
  - File: `sparql/ontology/model.go` (NEW)
  - Classes, properties, relationships
  - Hierarchy representation
  - Equivalence relations
  - ~200 LOC

### 5.2 Ontology Loader

- [ ] Build OWL/RDFS loader
  - File: `sparql/ontology/loader.go` (NEW)
  - Parse ontology triples
  - Extract class hierarchies
  - Build property definitions
  - Handle equivalence classes
  - ~300 LOC

- [ ] Ontology loader tests
  - File: `sparql/ontology/loader_test.go` (NEW)
  - Standard OWL/RDFS files
  - Various ontology patterns
  - ~100 test cases

### 5.3 Ontology → Schema Synthesis

- [ ] Build schema synthesis engine
  - File: `sparql/ontology/schema_builder.go` (NEW)
  - Convert ontology to Dgraph schema
  - Generate type predicates
  - Create relationship indexes
  - ~250 LOC

- [ ] Schema synthesis tests
  - File: `sparql/ontology/schema_builder_test.go` (NEW)
  - Various ontology structures
  - Schema correctness
  - ~100 test cases

### 5.4 Reasoning Infrastructure

- [ ] Design reasoning rule engine
  - File: `sparql/ontology/reasoning.go` (NEW)
  - Rule evaluation framework
  - Built-in rules (RDFS, OWL-lite)
  - Custom rule support
  - ~200 LOC

- [ ] Implement basic RDFS rules
  - File: `sparql/ontology/reasoning.go` (extension)
  - rdfs:subClassOf expansion
  - rdfs:subPropertyOf expansion
  - Type inheritance
  - ~150 LOC

### 5.5 Ontology Integration

- [ ] Load ontology at startup (optional config)
  - File: `sparql/ontology/loader.go` (modification)
  - Cache in memory
  - Invalidation strategy
  - ~50 LOC

- [ ] Use ontology in query optimization
  - File: `sparql/compiler_dql.go` (modification)
  - Apply class hierarchy knowledge
  - Expand implicit type relationships
  - Optimize based on ontology cardinality
  - ~100 LOC

### 5.6 Ontology Tests

- [ ] Reasoning correctness tests
  - File: `sparql/ontology/reasoning_test.go` (NEW)
  - RDFS inference rules
  - Custom rules
  - ~100 test cases

- [ ] End-to-end ontology tests
  - File: `sparql/ontology/integration_test.go` (NEW)
  - Load ontology → use in queries
  - Verify reasoning application
  - ~50 test cases

---

## Cross-Cutting Concerns

### Documentation

- [ ] Update README.md with architecture overview
  - High-level pipeline description
  - Point to ARCHITECTURE_SPEC.md
- [ ] Add inline code documentation
  - Each major function
  - Algebra operator semantics
  - Compilation phases
- [ ] Create troubleshooting guide
  - Common optimization issues
  - Debugging algebra rewriting
  - Schema validation errors

### Testing Strategy

- [ ] Unit tests for each module
- [ ] Integration tests between phases
- [ ] Regression tests (all existing tests must pass)
- [ ] Performance benchmarks
  - Algebra compilation time
  - Query optimization impact
  - Total execution time
  - Comparison to current implementation

### Code Quality

- [ ] Code review checklist
  - Design patterns consistent with Dgraph
  - Error handling completeness
  - Test coverage >80%
- [ ] Linting and formatting
- [ ] Security review (especially auth rules)

### Monitoring & Observability

- [ ] Add tracing hooks
  - Algebra rewriting stages
  - Compilation phases
  - Auth rule application
- [ ] Metrics
  - Query optimization effectiveness
  - Auth rule injection frequency
  - Ontology reasoning latency

---

## Dependencies & Blockers

| Task    | Dependencies          | Blocker? | Notes                        |
| ------- | --------------------- | -------- | ---------------------------- |
| Phase 1 | None                  | No       | Can start immediately        |
| Phase 2 | Phase 1 complete      | No       | Depends on algebra types     |
| Phase 3 | Phase 1-2 complete    | No       | Needs optimizer working      |
| Phase 3 | Schema access         | No       | Can work with partial schema |
| Phase 4 | Phase 3 complete      | No       | Needs DQL compilation        |
| Phase 4 | Auth system knowledge | No       | Study existing auth rules    |
| Phase 5 | Phase 1-3 complete    | No       | Can be done in parallel      |

---

## Success Criteria

### Phase 1 Completion

- [ ] All algebra operators implemented and tested
- [ ] Algebra validator passes all cases
- [ ] AST→Algebra conversion correct for 100+ test cases
- [ ] Zero regressions in existing tests

### Phase 2 Completion

- [ ] All optimization rewriters implemented
- [ ] Optimization correctness proven mathematically
- [ ] 30%+ improvement in query plan complexity (measured)
- [ ] Performance not degraded for simple queries

### Phase 3 Completion

- [ ] DQL compilation working end-to-end
- [ ] Direct Result construction (no re-parsing)
- [ ] Schema information properly attached
- [ ] Zero performance regression vs. current implementation

### Phase 4 Completion

- [ ] SPARQL auth rules working correctly
- [ ] Parity testing shows same behavior as GraphQL
- [ ] All ACL scenarios covered

### Phase 5 Completion

- [ ] Ontology loading and synthesis working
- [ ] Basic RDFS reasoning rules functional
- [ ] Ontology knowledge improves query optimization

---

## Delivery Timeline

```
Week 1-3   Phase 1: Algebra Foundation
Week 4-6   Phase 2: Query Optimization
Week 7-8   Phase 3: Schema-Aware Compilation
Week 9-10  Phase 4: Authorization Integration
Week 11-14 Phase 5: Ontology Foundation (can overlap)

Estimated: 14 weeks (3.5 months)
Actual will depend on: complexity, team size, testing thoroughness
```

---

## Future Extensions (Post-Implementation)

1. **Query Caching**: Cache optimized algebra expressions keyed by SPARQL syntax
2. **Federated Queries**: Use algebra as bridge for multi-datasource queries
3. **Incremental Compilation**: Reuse compilation of query fragments
4. **Explanation Generation**: Generate human-readable optimization steps
5. **Machine Learning Integration**: Learn cost model from execution statistics
6. **Advanced Reasoning**: OWL full reasoning, rule engines (SWRL), etc.

---

## Notes for Reviewers

- This plan is aggressive but achievable with focused effort
- Each phase is self-contained; can ship independently
- Early phases provide value even without later phases
- Ontology work (Phase 5) is speculative; may need iteration
- Recommend weekly sync during execution to catch blockers early
