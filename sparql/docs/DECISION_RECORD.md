# SPARQL Architecture Decision Record: Algebra-Driven Rewriting

**Date**: January 2026  
**Decision**: Introduce SPARQL Algebra intermediate representation with schema-aware compilation  
**Status**: Approved for Phase 1 implementation

---

## Problem Statement

The current SPARQL implementation has three significant limitations:

1. **No Query Rewriting**: SPARQL AST is translated directly to DQL without algebraic simplification
   - No filter pushdown optimization
   - No join reordering based on selectivity
   - Missed optimization opportunities

2. **Authorization Gap**: SPARQL queries bypass type-level RBAC rules that GraphQL queries receive
   - GraphQL queries use `addCommonRules()` to inject type-level auth filters
   - SPARQL queries only get predicate-level ACL filtering
   - Security implications: cannot enforce type-based access control

3. **Schema Disconnection**: SPARQL translation happens without schema context
   - Cannot validate predicates against schema during translation
   - Cannot use type information for optimization decisions
   - Blocks future semantic/ontology reasoning

4. **Redundant Parsing**: Generated DQL is serialized to strings and re-parsed
   - Loss of semantic information between layers
   - Performance overhead
   - Fragile: relies on string round-trip correctness

---

## Architectural Vision

**Core Idea**: Treat SPARQL like GraphQL - process through REST layer with schema awareness

**Current Pipeline**:

```
SPARQL String
  ↓ (antlr_adapter)
AST (SPARQLQueryImpl)
  ↓ (translator_extended)
[]*dql.GraphQuery
  ↓ (serialize)
DQL String
  ↓ (edgraph/server:parseRequest)
dql.ParseWithNeedVars
  ↓
dql.Result {Query: []*GraphQuery}
```

**Proposed Pipeline**:

```
SPARQL String
  ↓ (antlr_adapter)
AST (SPARQLQueryImpl)
  ↓ (NEW: algebra compiler)
SPARQL Algebra Tree
  ↓ (NEW: algebra optimizer)
Optimized Algebra Tree
  ↓ (NEW: schema-aware DQL compiler)
[]*dql.GraphQuery (with type metadata)
  ↓ (NEW: auth rule injector)
[]*dql.GraphQuery (with auth rules applied)
  ↓ (direct construction, no re-parsing)
dql.Result {Query: []*GraphQuery}
  ↓ (existing path)
ToSubGraph → ProcessGraph
```

**Key Benefits**:

- Schema information flows through entire pipeline
- Structured rewriting opportunities at algebra level
- Authorization rules applied before execution (like GraphQL)
- No string round-trip; semantic information preserved
- Foundation for ontology reasoning (Phase 5)

---

## Why SPARQL Algebra?

SPARQL Algebra is the W3C-specified intermediate form for SPARQL query semantics:

- https://www.w3.org/TR/sparql11-query/#sparqlAlgebra

**Why Use It?**:

1. **Standard**: Defined by W3C, not proprietary
2. **Compositional**: Each operator has well-defined semantics
3. **Rewritable**: Enabling algebraic optimization rules
4. **Bridging**: Natural intermediate between AST (parse form) and DQL (execution form)
5. **Future-Proof**: Enables advanced optimizations (caching, federated queries, etc.)

**Example Operators**:

- `BGP` - Basic Graph Pattern
- `Join` - Pattern combination
- `LeftJoin` - OPTIONAL semantics
- `Union` - UNION semantics
- `Filter` - WHERE constraints
- `Project` - SELECT variables
- `Aggregate` - GROUP BY with COUNT/SUM/etc.

---

## Implementation Strategy

### Phase 1: Foundation (Weeks 1-3)

Build the algebra type system and AST→Algebra converter. This is the critical foundation that all
later phases depend on.

**Deliverables**:

- Algebra operators (BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, etc.)
- Visitor pattern for traversal
- AST→Algebra converter
- Algebra validator
- 100+ comprehensive tests

**Isolated**: Can be developed independently; doesn't touch existing query execution.

### Phase 2: Optimization (Weeks 4-6)

Implement query rewriting rules: filter pushdown, join reordering, OPTIONAL simplification, dead
variable elimination.

**Deliverables**:

- AlgebraOptimizer infrastructure
- Filter pushdown rewriter
- Join reordering with cardinality estimates
- OPTIONAL simplification
- Comprehensive optimization tests

**Value**: 30%+ improvement in query plan complexity.

### Phase 3: Schema Integration (Weeks 7-8)

Connect algebra to Dgraph schema; compile to DQL with type information; eliminate re-parsing.

**Deliverables**:

- Schema-aware DQL compiler
- Type inference from patterns
- Direct dql.Result construction (no re-parsing)
- Integration with edgraph/server.go

**Impact**: Schema information available for optimization; performance improvement.

### Phase 4: Authorization (Weeks 9-10)

Apply RBAC rules to SPARQL queries like GraphQL does.

**Deliverables**:

- ApplyAuthRulesToQuery() function
- Type detection for auth context
- Auth rule injection
- Parity testing with GraphQL

**Security**: SPARQL queries now receive same type-level authorization as GraphQL.

### Phase 5: Ontology Foundation (Weeks 11-14)

Build infrastructure for ontology reasoning (OWL/RDFS); extend schema synthesis.

**Deliverables**:

- Ontology model and loader
- Schema synthesis from ontology
- Basic RDFS reasoning rules
- Ontology integration in optimization

**Future**: Enables semantic reasoning and knowledge graph features.

---

## Risk Assessment

| Risk                   | Likelihood | Impact | Mitigation                                             |
| ---------------------- | ---------- | ------ | ------------------------------------------------------ |
| Phase 1 too complex    | Low        | High   | Break into smaller PRs; start with minimal operators   |
| Integration issues     | Medium     | High   | Phase-by-phase; comprehensive tests; CI coverage       |
| Performance regression | Low        | High   | Benchmark before/after; fallback to old path if needed |
| Schema dependency      | Low        | Medium | Graceful degradation; work without schema in Phase 3   |
| Auth complexity        | Medium     | Medium | Study existing auth rules; iterate in Phase 4          |

---

## Alternative Approaches Considered

### 1. "Optimize at DQL Level"

- **Idea**: Apply optimizations after DQL parsing instead of before
- **Rejected**: DQL parsing already loses structure; too late for many optimizations
- **Trade-off**: Simpler to implement but fewer optimization opportunities

### 2. "Keep AST, Add Rewriters"

- **Idea**: Build rewriters that work directly on AST instead of algebra
- **Rejected**: AST is parse-specific; less natural for rewriting; harder to reason about
- **Trade-off**: Could work but less elegant

### 3. "Extend Existing Translator"

- **Idea**: Add optimizations to translator_extended.go incrementally
- **Rejected**: Ties optimizations to translation; harder to reason independently
- **Trade-off**: Faster initial win but limits future flexibility

### 4. "Use GraphQL Schema for SPARQL" (Phase 4+)

- **Idea**: Map SPARQL patterns to GraphQL schema types directly
- **Status**: Explored but deferred; algebra foundation enables this later
- **Reason**: Need algebra layer first for clean separation of concerns

---

## Success Criteria

### Phase 1 Success

- ✅ All algebra operators implemented and tested
- ✅ 100+ test cases, all passing
- ✅ Zero regressions in existing SPARQL tests
- ✅ Code review approved with <10 style comments

### Full Success (All Phases)

- ✅ SPARQL queries receive same optimizations as GraphQL
- ✅ SPARQL queries receive same authorization as GraphQL
- ✅ 30%+ improvement in query plan complexity
- ✅ Zero performance regression vs. baseline
- ✅ Schema information integrated throughout pipeline
- ✅ Foundation ready for ontology reasoning

---

## Decision Rationale

We chose this architecture because:

1. **Proven Pattern**: GraphQL resolver already does schema-aware optimization; we're applying the
   same approach
2. **Standards-Based**: SPARQL Algebra is W3C-specified, not proprietary
3. **Scalable**: Each phase is independent; can ship incrementally
4. **Future-Ready**: Algebra layer enables advanced features (caching, federation, reasoning)
5. **Performance**: Direct DQL construction avoids string round-trip
6. **Security**: Closes authorization gap between SPARQL and GraphQL

---

## Implementation Timeline

```
Week 1-3    Phase 1: Algebra Foundation (CRITICAL PATH)
  ├─ Types & operators
  ├─ AST→Algebra conversion
  └─ Comprehensive testing

Week 4-6    Phase 2: Query Optimization
  ├─ Filter pushdown
  ├─ Join reordering
  └─ Optimization tests

Week 7-8    Phase 3: Schema Integration
  ├─ Schema-aware compilation
  └─ Direct DQL construction

Week 9-10   Phase 4: Authorization
  ├─ Auth rule application
  └─ Parity testing

Week 11-14  Phase 5: Ontology Foundation (can parallel)
  ├─ Ontology loading
  └─ Reasoning rules

Total: ~14 weeks (3.5 months)
Can parallelize Phase 5 with Phases 3-4.
```

---

## Approval

| Role           | Status      | Date     | Notes                              |
| -------------- | ----------- | -------- | ---------------------------------- |
| Architecture   | ✅ Approved | Jan 2026 | Align SPARQL with GraphQL patterns |
| Implementation | 👤 Pending  | -        | Phase 1 starting                   |
| Review         | 👤 Pending  | -        | Code review process                |

---

## References

- [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - Full architectural specification
- [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) - Detailed implementation tasks
- [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) - Week 1 guide
- [W3C SPARQL Algebra](https://www.w3.org/TR/sparql11-query/#sparqlAlgebra) - Standard reference
- [graphql/resolve/query_rewriter.go](graphql/resolve/query_rewriter.go) - Reference implementation
  pattern
