# SPARQL Integration Map: System Architecture Context

**Purpose**: Show where SPARQL fits in Dgraph's overall architecture  
**Audience**: Architects, senior engineers, future maintainers

---

## High-Level System View

```
┌─────────────────────────────────────────────────────────────┐
│                         Client Layer                         │
│                    (HTTP, gRPC, etc)                         │
└────────┬────────────────────────────────────────────────────┘
         │
    ┌────▼────────────────────────────────────────────────────┐
    │              edgraph/server.go (REST Layer)              │
    │  ┌──────────────────────────────────────────────────────┤
    │  │ doQuery(ctx, req) - Main dispatcher                   │
    │  ├─► isGraphQL? → GraphQL path (existing)               │
    │  ├─► isSPARQL? → SPARQL path (NEW - this project)       │
    │  └─► else → DQL path (existing)                         │
    └────┬───────────────────────────────────────────────────┘
         │
    ┌────▼─────────────────────────────────────────────────────┐
    │         Query Compilation & Execution Engine             │
    │                                                           │
    │  ┌─────────────────────────────────────────────────────┐ │
    │  │ GraphQL Path (graphql/resolve/query_rewriter.go)    │ │
    │  │  • addCommonRules (RBAC injection)                  │ │
    │  │  • Query optimization (filter reordering, etc)      │ │
    │  │  • Schema-aware compilation                         │ │
    │  └──────────────────────┬──────────────────────────────┘ │
    │                         │                                │
    │  ┌──────────────────────▼──────────────────────────────┐ │
    │  │ [NEW] SPARQL Path (sparql/*.go)                     │ │
    │  │  • Phase 1: AST → Algebra                           │ │
    │  │  • Phase 2: Algebra Optimization                    │ │
    │  │  • Phase 3: Schema-aware DQL compilation            │ │
    │  │  • Phase 4: Auth rule injection (NEW)               │ │
    │  │  • Phase 5: Ontology reasoning (future)             │ │
    │  └──────────────────────┬──────────────────────────────┘ │
    │                         │                                │
    │  ┌──────────────────────▼──────────────────────────────┐ │
    │  │ DQL Path (dql/parser.go)                            │ │
    │  │  • Parse DQL strings                                │ │
    │  │  • Build GraphQuery trees                           │ │
    │  └──────────────────────┬──────────────────────────────┘ │
    │                         │                                │
    │  ┌──────────────────────▼──────────────────────────────┐ │
    │  │ Unified Execution (query/query.go)                  │ │
    │  │  • ToSubGraph (GraphQuery → SubGraph)               │ │
    │  │  • ProcessGraph (execute SubGraph)                  │ │
    │  └──────────────────────┬──────────────────────────────┘ │
    └────┬───────────────────┬──────────────────────────────────┘
         │                   │
    ┌────▼───────────────────▼──────────────────────────────────┐
    │              Execution Layer                              │
    │                                                            │
    │  ┌────────────────────────────────────────────────────┐  │
    │  │ query/algebra.go - UID set operations              │  │
    │  │  • MergeSorted, Intersect, etc                     │  │
    │  └────────────────────────────────────────────────────┘  │
    │                                                            │
    │  ┌────────────────────────────────────────────────────┐  │
    │  │ worker/task.go - Task execution                    │  │
    │  │  • planForEqFilter (cardinality-based planning)     │  │
    │  │  • Filter execution, function evaluation           │  │
    │  │  • Index selection decisions                       │  │
    │  └────────────────────────────────────────────────────┘  │
    │                                                            │
    │  ┌────────────────────────────────────────────────────┐  │
    │  │ posting/list.go - Posting list operations          │  │
    │  │  • Index scanning, filtering                       │  │
    │  │  • Cardinality estimation                          │  │
    │  └────────────────────────────────────────────────────┘  │
    └─────────────────────────┬────────────────────────────────┘
         │
    ┌────▼────────────────────────────────────────────────────┐
    │              Storage & Indexes (Badger)                  │
    └───────────────────────────────────────────────────────────┘
```

---

## Query Execution Paths

### Path 1: GraphQL → DQL (Existing)

```
GraphQL Query String
    ↓
[graphql/schema/schema.go]
  Load schema

    ↓
[graphql/resolve/query_rewriter.go]
  rewriteAsQuery/Get/Mutation
  • addCommonRules (RBAC injection)
  • rootQueryOptimization (filter reordering)
  • Selection set processing
  → []*dql.GraphQuery (schema-aware)

    ↓
[query/query.go]
  ToSubGraph(GraphQuery) → SubGraph
  ProcessGraph(SubGraph) → results
```

**Key Feature**: Type-level RBAC rules injected at rewriter level (before execution)

---

### Path 2: SPARQL → DQL (Current - Before Project)

```
SPARQL Query String
    ↓
[sparql/antlr_adapter.go]
  Parse → SPARQLQueryImpl (AST)

    ↓
[sparql/translator_extended.go]
  TranslateToGraphQueries
  • AST → []*dql.GraphQuery
  → []*dql.GraphQuery (no schema context!)

    ↓
[edgraph/server.go:parseRequest]
  Serialize to DQL string
  → dql.ParseWithNeedVars (RE-PARSE!)
  → dql.Result {Query: []*GraphQuery}

    ↓
[query/query.go]
  ToSubGraph(GraphQuery) → SubGraph
  ProcessGraph(SubGraph) → results
```

**Limitation**: No schema context, no type-level auth, redundant parsing

---

### Path 3: SPARQL → DQL (Proposed - After Project)

```
SPARQL Query String
    ↓
[sparql/antlr_adapter.go]
  Parse → SPARQLQueryImpl (AST)

    ↓
[sparql/context.go] NEW
  Load schema context
  Load auth context

    ↓
[sparql/algebra.go] NEW - Phase 1
  ASTToAlgebra()
  → Algebra Tree (Join, Filter, BGP, etc)

    ↓
[sparql/algebra_rewriter.go] NEW - Phase 2
  Optimize(Algebra)
  • Filter pushdown
  • Join reordering
  • OPTIONAL simplification
  → Optimized Algebra Tree

    ↓
[sparql/compiler_dql.go] NEW - Phase 3
  Compile(Algebra, SchemaContext)
  • Type inference from predicates
  • Type filter injection
  • Cardinality hints from schema
  → []*dql.GraphQuery (schema-aware!)

    ↓
[sparql/auth_rules.go] NEW - Phase 4
  ApplyAuthRulesToQuery()
  • Type detection
  • RBAC rule injection
  • Auth filter queries
  → []*dql.GraphQuery (with auth!)

    ↓
[edgraph/server.go:parseRequest] - MODIFIED
  Skip re-parsing
  Construct dql.Result directly
  → dql.Result {Query: []*GraphQuery}

    ↓
[query/query.go]
  ToSubGraph(GraphQuery) → SubGraph
  ProcessGraph(SubGraph) → results
```

**Gains**:

- ✅ Schema context throughout
- ✅ Type-level auth (like GraphQL)
- ✅ Query optimization (filter pushdown, join reordering)
- ✅ No redundant re-parsing
- ✅ Foundation for ontology reasoning

---

## Component Interactions

### Schema Access

```
SPARQL Translation Needs Schema:
  sparql/context.go
      ↓
  schema.State() [getter in edgraph]
      ↓
  graphql/schema/schema.go [main schema holder]
      ↓
  Predicate → Type mappings
  Cardinality statistics
  Authorization rules
```

### Authorization Integration

```
SPARQL Auth Rule Application:
  sparql/auth_rules.go
      ↓
  authRewriter.evaluateStaticRules() [borrow from GraphQL]
  authRewriter.addAuthQueries() [same as GraphQL]
      ↓
  graphql/resolve/query_rewriter.go [shared auth logic]
```

### Execution Integration

```
All paths converge:
  []*dql.GraphQuery
      ↓
  query/query.go:ToSubGraph
      ↓
  query/query.go:ProcessQuery
      ↓
  worker/task.go [task execution]
      ↓
  posting/ [index operations]
      ↓
  Results
```

---

## File Organization

### Current SPARQL Files (Before Project)

```
sparql/
├── antlr_adapter.go        - ANTLR parser integration
├── antlr_visitor.go        - AST building
├── adapter.go              - AST interface
├── ast.go                  - AST type definitions
├── simple_parser.go        - Legacy parser
├── translate.go            - Entry point
├── translator_impl.go      - Basic translation
├── translator_extended.go  - Extended features (OPTIONAL, UNION, aggregates)
├── filter_extended.go      - FILTER parsing & translation
├── loader.go               - Test data loading
├── IMPLEMENTATION.md       - Feature documentation
├── EXTENDED_FEATURES.md    - Advanced features
├── TEST_COVERAGE.md        - Test inventory
└── gen/                    - ANTLR-generated code
```

### New SPARQL Files (After Project - Phases 1-5)

```
sparql/
├── [existing files]
├──
├── ─── Phase 1: Algebra Foundation ───
├── algebra.go              - Algebra operators and types
├── algebra_visitor.go      - Visitor pattern for algebra
├── algebra_test.go         - Phase 1 tests (100+ cases)
├── context.go              - SPARQLExecutionContext
├──
├── ─── Phase 2: Optimization ───
├── algebra_rewriter.go     - Query optimization rules
├── algebra_optimizer_test.go
├──
├── ─── Phase 3: Schema Integration ───
├── compiler_dql.go         - Schema-aware DQL compilation
├── compiler_dql_test.go
├──
├── ─── Phase 4: Authorization ───
├── auth_rules.go           - SPARQL auth rule injection
├── auth_rules_test.go
├──
├── ─── Phase 5: Ontology ───
├── ontology/
│   ├── loader.go          - OWL/RDFS loading
│   ├── model.go           - Ontology data model
│   ├── schema_builder.go  - Ontology → Schema synthesis
│   ├── reasoning.go       - Reasoning rules
│   ├── loader_test.go
│   ├── schema_builder_test.go
│   └── reasoning_test.go
├──
├── ─── Documentation ───
├── ARCHITECTURE_SPEC.md     - [NEW] Full architectural spec
├── SPARQL_ALGEBRA_TODO.md   - [NEW] Implementation roadmap
├── DECISION_RECORD.md       - [NEW] Architecture decision
├── PHASE_1_GETTING_STARTED.md - [NEW] Week 1 guide
└── README.md               - Updated overview
```

---

## Critical Integration Points

### 1. REST Layer Entry (edgraph/server.go)

**Current**:

```go
func (s *Server) doQuery(ctx context.Context, req *Request) (*api.Response, error) {
    qc := &queryContext{req: req.req, ...}
    parseRequest(ctx, qc)           // Parse query to DQL
    processQuery(ctx, qc)           // Execute
}

func parseRequest(ctx context.Context, qc *queryContext) error {
    qc.dqlRes, err = dql.ParseWithNeedVars(dql.Request{
        Str: qc.req.Query,          // Parse string!
    })
}
```

**After Integration**:

```go
func parseRequest(ctx context.Context, qc *queryContext) error {
    if isSPARQL(qc.req.Query) {
        // NEW PATH: Direct compilation
        sqCtx := &sparql.SPARQLExecutionContext{...}
        queries, err := sparql.CompileToGraphQueries(sqCtx, ...)
        qc.dqlRes = sparql.DirectToDQLResult(queries, ...)
        return nil
    }

    // Existing path for DQL
    qc.dqlRes, err = dql.ParseWithNeedVars(...)
}
```

**Impact**: Minimal; new code path, no changes to existing DQL path.

---

### 2. Authorization Checking (edgraph/access.go)

**Current**:

```go
func authorizeQuery(ctx context.Context, parsedReq *dql.Result, graphql bool) error {
    // Filters predicates based on ACL
    // For GraphQL, adds user-specific filters
}
```

**After Integration**:

```go
// Works as-is! SPARQL queries now have auth rules pre-injected
// (unlike current SPARQL which skips addCommonRules)
// So authorizeQuery sees them just like GraphQL queries
```

**Impact**: None; existing function works with pre-injected auth rules.

---

### 3. Query Execution (query/query.go)

**Current**:

```go
func ToSubGraph(ctx context.Context, gq *dql.GraphQuery) (*SubGraph, error) {
    // Converts GraphQuery to SubGraph
    // Used by both GraphQL and current SPARQL paths
}

func ProcessGraph(ctx context.Context, sg, parent *SubGraph, ...) {
    // Executes SubGraph
    // Works with any input source (GraphQL, SPARQL, DQL)
}
```

**After Integration**:

```
// No changes needed!
// SPARQL queries produce []*dql.GraphQuery just like GraphQL
// So they flow through existing ToSubGraph/ProcessGraph automatically
```

**Impact**: None; seamless integration with existing execution engine.

---

## Data Flow Example

### Query: "Find all people named Alice with their email"

**SPARQL Input**:

```sparql
SELECT ?person ?email WHERE {
  ?person rdf:type Person .
  ?person foaf:name "Alice" .
  ?person foaf:email ?email .
}
```

**Phase 1 Output** (Algebra):

```
Project(
  vars: [?person, ?email],
  input: Filter(
    expr: (?name = "Alice"),
    input: Join(
      left: Join(
        left: BGP(?person rdf:type Person),
        right: BGP(?person foaf:name ?name)
      ),
      right: BGP(?person foaf:email ?email)
    )
  )
)
```

**Phase 2 Output** (Optimized Algebra):

```
Project(
  vars: [?person, ?email],
  input: Join(
    left: Join(
      left: BGP(?person rdf:type Person),
      right: Filter(
        expr: (?name = "Alice"),
        input: BGP(?person foaf:name ?name)
      )
    ),
    right: BGP(?person foaf:email ?email)
  )
)
[Filter pushed down to earliest point]
```

**Phase 3 Output** (DQL with Type Info):

```go
[]*dql.GraphQuery{
  {
    Attr: "query",
    Children: [
      {
        Attr: "person",
        Func: {Name: "type", Args: ["Person"]},
        Children: [
          {
            Attr: "foaf.name",
            Filter: {Func: {Name: "eq", Args: ["Alice"]}},
            Children: [
              {Attr: "foaf.email"}
            ]
          }
        ]
      }
    ]
  }
}
[Type information attached; filter positioned optimally]
```

**Phase 4 Output** (With Auth):

```go
// If Person type has auth rules:
[]*dql.GraphQuery{
  {Attr: "query", Var: "p1", ...},     // Auth rule query
  {Attr: "person", Func: {Name: "uid", Args: ["p1"]}, ...}, // Wrapped with UID filter
  // ... rest of query
}
```

**Final Execution**:

```
→ query/query.go:ToSubGraph → SubGraph
→ query/query.go:ProcessGraph → Execute
→ worker/task.go → Task execution
→ posting/list.go → Index operations
→ Results returned
```

---

## Testing Strategy

### Unit Tests (Per Phase)

- Phase 1: `sparql/algebra_test.go` (100+ tests)
- Phase 2: `sparql/algebra_optimizer_test.go` (150+ tests)
- Phase 3: `sparql/compiler_dql_test.go` (200+ tests)
- Phase 4: `sparql/auth_rules_test.go` (150+ tests)
- Phase 5: `sparql/ontology/*_test.go` (200+ tests)

### Integration Tests

- `sparql/e2e_test.go` - End-to-end with full execution
  - All existing tests must continue to pass
  - New schema-aware optimization tests
  - Authorization behavior tests

### Regression Tests

- Verify all existing SPARQL tests pass (200+ cases)
- No performance degradation
- Behavior identical to current implementation (for non-optimized cases)

### Performance Benchmarks

- Algebra compilation time
- Optimization overhead
- Query execution time (with/without optimization)
- Memory usage

---

## Rollout Plan

### Internal Development

1. Week 1-3: Phase 1 implementation in feature branch
2. Week 4-6: Phase 2 implementation
3. Continuous: Integration tests, performance testing
4. Week 14: Final review and merge

### Rollout Strategy

1. **Phase 1**: Land algebra foundation (no user impact)
2. **Phase 2-3**: Enable optimization (transparent improvement)
3. **Phase 4**: Enable auth (security enhancement)
4. **Phase 5**: Ontology (future feature, opt-in)

### Fallback Plan

- Keep old translator as fallback during transition
- Disable new path if regressions detected
- Gradual enablement via feature flags

---

## Future Extensions

Once foundation is solid (Phases 1-3), enable:

1. **Query Caching**: Cache optimized algebra expressions
2. **Federated Queries**: Use algebra as bridge for multi-DB queries
3. **Plan Explanation**: Generate human-readable optimization steps
4. **Cost Model Learning**: ML-based cardinality estimation
5. **Advanced Reasoning**: OWL-full, SWRL, custom rules
6. **Incremental Compilation**: Reuse compilation fragments

---

## Key Contacts & References

### Core Contributors

- @architect: Architecture design and oversight
- @implementation: Phase 1-3 implementation
- @ontology-expert: Phase 5 design

### Related Code

- [GraphQL Query Rewriter](graphql/resolve/query_rewriter.go) - Reference pattern
- [DQL Parser](dql/parser.go) - GraphQuery format
- [Query Execution](query/query.go) - Execution engine
- [Worker Tasks](worker/task.go) - Planning and execution

### External Resources

- [W3C SPARQL Spec](https://www.w3.org/TR/sparql11-query/)
- [SPARQL Algebra](https://www.w3.org/TR/sparql11-query/#sparqlAlgebra)
