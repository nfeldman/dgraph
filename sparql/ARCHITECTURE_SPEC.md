# SPARQL Architecture Specification: Algebra-Driven Query Rewriting with Schema Integration

**Status**: Design Proposal  
**Date**: January 2026  
**Authors**: Dgraph Team  
**Related Issues**: Schema-aware SPARQL optimization, Ontology reasoning foundation

---

## Executive Summary

This document proposes an architectural evolution of the SPARQL implementation to:

1. **Introduce SPARQL Algebra as an intermediate representation** between the AST and DQL, enabling
   structured query rewriting and optimization
2. **Integrate schema information** into the SPARQL→DQL translation pipeline, enabling type-aware
   optimizations and authorization rule injection (currently missing)
3. **Isolate SPARQL handling to the REST layer**, keeping SPARQL concerns separate from the core
   query execution engine
4. **Establish a foundation for ontology reasoning** by modeling ontologies as schema artifacts

This approach mirrors how GraphQL queries flow through the resolve layer with schema context,
providing parity for authorization, optimization, and future semantic reasoning capabilities.

---

## Current Architecture

### Query Execution Pipeline

```
User Request
    ↓
[REST Layer - edgraph/server.go]
  ├─ GraphQL: parseRequest → authorizeRequest → processQuery
  └─ SPARQL:  Parse SPARQL string → TranslateToGraphQueries → DQL strings → parseRequest
    ↓
[Parsing Layer - dql/parser.go]
  dql.ParseWithNeedVars(queryStr) → dql.Result {Query: []*GraphQuery}
    ↓
[Query Execution Layer - query/query.go]
  ToSubGraph(GraphQuery) → SubGraph → ProcessGraph
    ↓
[Worker Layer - worker/task.go]
  Task execution with cardinality-based planning
```

### SPARQL Translation Today

```
SPARQL Query String
    ↓
[antlr_adapter.go] Parse → SPARQLQueryImpl (AST)
    ↓
[translator_extended.go] TranslateToGraphQueries
    ├─ SELECT → Patterns → DQL GraphQuery trees
    ├─ FILTER expressions → DQL FilterTree
    ├─ BIND/Aggregates → DQL math/groupby
    └─ UNION/OPTIONAL → DQL child queries + filters
    ↓
[]*dql.GraphQuery (sent to DQL parser for validation)
    ↓
dql.ParseWithNeedVars re-parses as DQL strings (overhead!)
```

### Current Limitations

1. **No Schema Context**: SPARQL→DQL translation happens without access to schema information
   - Cannot apply type-aware optimizations
   - Cannot inject authorization rules (unlike GraphQL path)
   - Cannot validate predicates against schema during translation

2. **No Query Rewriting**: AST is directly translated without algebraic simplification
   - No filter pushdown optimization
   - No join reordering based on selectivity
   - No sharing of subexpression patterns

3. **Authorization Gap**: SPARQL queries bypass the `addCommonRules` logic that injects RBAC filters
   - GraphQL gets type-level auth rules automatically
   - SPARQL only gets predicate-level ACL filtering (in `authorizeQuery`)
   - No way to apply graph-pattern-level authorization

4. **Redundant Parsing**: Generated DQL GraphQuery objects are serialized to strings and re-parsed
   - Loss of schema information between layers
   - Performance overhead for intermediate serialization

---

## Proposed Architecture

### Phase 1: SPARQL Algebra Intermediate Representation

Introduce a SPARQL algebra layer that:

- Converts AST to a normalized algebraic form
- Represents operations as composable algebra operators
- Enables pattern-based rewriting before DQL generation

```
SPARQL AST
    ↓
[SPARQL Algebra Compiler] - NEW FILE: sparql/algebra.go
    ├─ Transform BGPs → Join operators
    ├─ Normalize FILTER expressions
    ├─ Hoist OPTIONAL → LeftJoin operators
    ├─ Expand UNION → Union operators
    └─ Collect projection/aggregation operators
    ↓
Algebra Query Tree
    (Join, Filter, LeftJoin, Union, Project, Aggregate, Bind, etc.)
    ↓
[Algebra Rewriter] - NEW FILE: sparql/algebra_rewriter.go
    ├─ Filter pushdown (move filters down the tree)
    ├─ Common subexpression elimination
    ├─ Join reordering (based on cardinality estimates)
    ├─ OPTIONAL simplification
    └─ Dead variable elimination
    ↓
Optimized Algebra Query Tree
```

### Phase 2: Schema-Aware Translation

Enhance translation to leverage schema information:

```
Optimized Algebra Query Tree + Schema Context
    ↓
[Schema-Aware DQL Compiler] - ENHANCED: sparql/translator_extended.go
    ├─ Resolve predicates to schema types
    ├─ Validate predicate compatibility
    ├─ Extract type information from patterns
    ├─ Identify implicit type relationships
    └─ Build type-filter functions where applicable
    ↓
Enhanced []*dql.GraphQuery with schema metadata
    ├─ Type information attached to nodes
    ├─ Cardinality hints from schema statistics
    └─ Authorization rules prepared (pending auth context)
```

### Phase 3: REST Layer Integration

Modify the REST layer to provide schema and auth context to SPARQL translation:

```
[edgraph/server.go] doQuery() - ENHANCED
    ├─ Detect query type (GraphQL vs DQL vs SPARQL)
    ├─ For SPARQL:
    │   ├─ Load schema context
    │   ├─ Parse SPARQL → AST
    │   ├─ Compile to SPARQL Algebra
    │   ├─ Apply Algebra rewrites
    │   ├─ Translate to enhanced DQL GraphQuery[]
    │   │   (with type info, schema validation done)
    │   ├─ Skip dql.ParseWithNeedVars (direct to Result)
    │   └─ Pass to processQuery (like GraphQL path)
    │
    ├─ For GraphQL: existing path (with addCommonRules)
    └─ For DQL: existing path (parse string)
    ↓
query.Request.Process(ctx)
```

### Phase 4: Authorization Rule Integration

Add SPARQL-specific auth rule application (similar to GraphQL's `addCommonRules`):

```
[sparql/auth_rules.go] - NEW FILE
    ├─ ApplyAuthRulesToQuery(ctx, queries, schema)
    ├─ For each query block:
    │   ├─ Determine primary type (if detectable from predicates)
    │   ├─ Evaluate auth rules for that type
    │   ├─ Apply static RBAC rules
    │   ├─ Inject auth filter queries (@cascade)
    │   └─ Add UID function wrapping where needed
    └─ Return auth-enhanced query tree
```

---

## Detailed Design: SPARQL Algebra

### Algebra Operators

```go
// sparql/algebra.go

type AlgebraExpr interface {
	Accept(visitor AlgebraVisitor) interface{}
	String() string
}

// Base operators
type (
	BGP struct {
		Patterns []*Triple  // Basic graph pattern
	}

	Filter struct {
		Expr   Expression  // Filter expression
		Input  AlgebraExpr // What to filter
	}

	Join struct {
		Left  AlgebraExpr
		Right AlgebraExpr
	}

	LeftJoin struct {
		Input    AlgebraExpr
		Patterns []*Triple  // Optional patterns
		Filter   Expression
	}

	Union struct {
		Alternatives []AlgebraExpr
	}

	Project struct {
		Vars  []string   // Variables to project
		Input AlgebraExpr
	}

	Aggregate struct {
		Op    string     // COUNT, SUM, MIN, MAX, AVG
		Expr  Expression
		Group []string   // GROUP BY variables
		Input AlgebraExpr
	}

	Bind struct {
		Var  string
		Expr Expression
		Input AlgebraExpr
	}

	Distinct struct {
		Input AlgebraExpr
	}

	OrderBy struct {
		Expressions []Expression
		Ascending   []bool
		Input       AlgebraExpr
	}

	Limit struct {
		Count  int
		Offset int
		Input  AlgebraExpr
	}
)

// Query optimizer
type AlgebraOptimizer struct {
	schema *Schema  // Schema for cardinality estimation
}

func (ao *AlgebraOptimizer) Optimize(expr AlgebraExpr) AlgebraExpr {
	// Apply rewriting rules
	// 1. Filter pushdown
	// 2. Join reordering
	// 3. OPTIONAL simplification
	// 4. Dead variable elimination
}

// DQL compiler
type DQLCompiler struct {
	schema *Schema
}

func (dc *DQLCompiler) Compile(expr AlgebraExpr, schemaCtx map[string]*schema.Type) (*dql.GraphQuery, error) {
	// Convert algebra to DQL GraphQuery with schema information
	// Attach type information to nodes
	// Inject type filters where beneficial
}
```

### Example: Filter Pushdown

**Input Algebra**:

```
Project(vars=[?person, ?name],
  Filter(expr: ?name = "Alice",
    Join(
      BGP(?person rdf:type Person),
      BGP(?person foaf:name ?name)
    )
  )
)
```

**After Pushdown**:

```
Project(vars=[?person, ?name],
  Join(
    BGP(?person rdf:type Person),
    Filter(expr: ?name = "Alice",
      BGP(?person foaf:name ?name)
    )
  )
)
```

---

## Integration Points with Dgraph Core

### 1. Schema Access

SPARQL translation needs access to Dgraph schema:

```go
// In edgraph/server.go or new sparql/context.go
type SPARQLExecutionContext struct {
	Schema    *schema.Schema        // GraphQL schema (has predicate info)
	Namespace uint64
	Auth      *AuthContext          // User, groups, ACLs
	Ctx       context.Context
}

// In translator_extended.go
func CompileToGraphQueries(
	ctx *SPARQLExecutionContext,
	ast *SPARQLQueryImpl,
) ([]*dql.GraphQuery, error) {
	// 1. Convert AST → Algebra
	algebra := astToAlgebra(ast)

	// 2. Optimize with schema info
	optimizer := NewAlgebraOptimizer(ctx.Schema)
	optimized := optimizer.Optimize(algebra)

	// 3. Compile to DQL with schema context
	compiler := NewDQLCompiler(ctx.Schema)
	queries, err := compiler.Compile(optimized, extractTypeContext(ctx.Schema))
	if err != nil {
		return nil, err
	}

	// 4. Apply authorization rules (if enabled)
	if ctx.Auth != nil {
		queries, err = applyAuthRules(ctx, queries)
	}

	return queries, err
}
```

### 2. Avoiding Re-parsing

Currently: `GraphQuery[] → serialize → dql.ParseWithNeedVars → GraphQuery[]`

**Proposed**: Construct `dql.Result` directly

```go
// In query/query.go or new wrapper
func DirectToDQLResult(gqs []*dql.GraphQuery, queryVars []*dql.Vars) dql.Result {
	return dql.Result{
		Query:     gqs,
		QueryVars: queryVars,
	}
}

// In edgraph/server.go
if isSPARQL {
	// Direct construction without re-parsing
	qc.dqlRes = DirectToDQLResult(sparqlQueries, sparqlVars)
} else {
	// Existing DQL parse path
	qc.dqlRes, err = dql.ParseWithNeedVars(...)
}
```

### 3. Authorization Integration

Apply auth rules at translation time, similar to GraphQL:

```go
// In sparql/auth_rules.go
func ApplyAuthRulesToQuery(
	ctx context.Context,
	queries []*dql.GraphQuery,
	schemaCtx *SchemaContext,
) ([]*dql.GraphQuery, error) {
	// Extract auth context
	authCtx := extractAuthContext(ctx)

	result := make([]*dql.GraphQuery, 0)

	for _, q := range queries {
		// Determine primary type if possible
		primaryType := inferTypeFromQuery(q, schemaCtx)

		if primaryType != nil {
			// Use GraphQL-style auth rule injection
			// (similar to addCommonRules logic)
			authQ := injectAuthRules(q, primaryType, authCtx)
			result = append(result, authQ...)
		} else {
			// No type detected - return as-is
			// (predicate-level auth still applied later)
			result = append(result, q)
		}
	}

	return result, nil
}
```

---

## Future: Ontology Reasoning

### Phase 5: Ontology as Schema

Once the algebra + schema framework is in place:

```
RDF Ontology Graph (OWL/RDFS)
    ↓
[Ontology Loader] - NEW: sparql/ontology/loader.go
    ├─ Parse OWL/RDFS triples
    ├─ Extract class hierarchies
    ├─ Extract property definitions
    └─ Build equivalence mappings
    ↓
Ontology Model
    ├─ Classes (subclass relationships)
    ├─ Properties (domain, range, subproperty)
    ├─ Equivalence classes
    └─ Restrictions (cardinality, inverseof, etc.)
    ↓
[Schema Synthesis] - NEW: sparql/ontology/schema_builder.go
    ├─ Convert ontology to dgraph schema format
    ├─ Create implicit type predicates
    ├─ Add inference rules
    └─ Build relationship indexes
    ↓
Enhanced Dgraph Schema with Ontology Knowledge
```

### Example: Class Hierarchy Optimization

```sparql
SELECT ?x WHERE {
  ?x rdf:type ex:Publication .
  ?x dc:title ?title .
}
```

With ontology knowledge:

```
Publication
  ├─ Journal (subclass)
  ├─ Conference (subclass)
  └─ Thesis (subclass)
```

**Optimized Query**:

```
Join(
  Union(
    BGP(?x rdf:type Journal),
    BGP(?x rdf:type Conference),
    BGP(?x rdf:type Thesis)
  ),
  BGP(?x dc:title ?title)
)
```

Or with reasoning:

```
BGP(?x rdf:type ex:Publication)  // Index expands to subclasses
BGP(?x dc:title ?title)
```

---

## Implementation Phases

### Phase 1: Foundation (2-3 weeks)

- [ ] Design SPARQL algebra operators
- [ ] Implement `astToAlgebra()` converter
- [ ] Build basic algebra validator
- [ ] Create comprehensive algebra unit tests

### Phase 2: Optimization (2-3 weeks)

- [ ] Implement filter pushdown rewriter
- [ ] Implement join ordering based on cardinality
- [ ] Implement OPTIONAL simplification
- [ ] Add algebra optimization tests

### Phase 3: DQL Integration (1-2 weeks)

- [ ] Create schema-aware DQL compiler
- [ ] Direct DQL Result construction (avoid re-parsing)
- [ ] Type information attachment to GraphQuery nodes
- [ ] Integration tests with existing DQL execution

### Phase 4: Authorization (1 week)

- [ ] Implement `ApplyAuthRulesToQuery()`
- [ ] SPARQL-specific auth rule tests
- [ ] Parity testing with GraphQL auth behavior

### Phase 5: Ontology Foundation (2-4 weeks)

- [ ] Design ontology model
- [ ] Build ontology loader (OWL/RDFS)
- [ ] Create schema synthesis from ontology
- [ ] Reasoning rule engine skeleton

---

## File Structure

```
sparql/
├── adapter.go              (existing)
├── antlr_adapter.go        (existing)
├── translate.go            (existing)
├── translator_extended.go  (existing)
├── translator_impl.go      (existing)
├── filter_extended.go      (existing)
│
├── ARCHITECTURE_SPEC.md    (this file)
├── algebra.go              (NEW - Phase 1)
├── algebra_rewriter.go     (NEW - Phase 2)
├── algebra_visitor.go      (NEW - Phase 1, pattern for traversal)
├── compiler_dql.go         (NEW - Phase 3, schema-aware compilation)
├── auth_rules.go           (NEW - Phase 4)
├── context.go              (NEW - execution context with schema)
│
├── ontology/               (NEW - Phase 5)
│   ├── loader.go
│   ├── model.go
│   ├── schema_builder.go
│   └── reasoning.go
│
└── tests/
    ├── algebra_test.go
    ├── algebra_optimizer_test.go
    ├── compiler_dql_test.go
    └── auth_rules_test.go
```

---

## Benefits

1. **Schema-Aware Optimization**: Same level of sophistication as GraphQL
2. **Authorization Parity**: SPARQL gets type-level auth like GraphQL
3. **Cleaner Integration**: SPARQL handling isolated to REST layer
4. **Performance**: Avoid string serialization round-trip
5. **Extensibility**: Algebra layer enables future transformations (query caching, federated
   queries, etc.)
6. **Ontology Ready**: Foundation for semantic reasoning and knowledge graph features

---

## Risk Mitigation

| Risk              | Mitigation                                             |
| ----------------- | ------------------------------------------------------ |
| Large refactor    | Phase incrementally; Phase 1-3 functional on their own |
| Backward compat   | Keep old translator as fallback during transition      |
| Schema dependency | Graceful degradation if schema unavailable             |
| Auth complexity   | Start with simple type inference; enhance over time    |

---

## Success Metrics

- [ ] Phase 1: All algebra operators tested, 100+ test cases
- [ ] Phase 2: 30%+ query optimization (measured by plan complexity reduction)
- [ ] Phase 3: Zero performance regression vs. current implementation
- [ ] Phase 4: SPARQL auth behavior matches GraphQL auth
- [ ] Phase 5: First ontology reasoning rules working

---

## References

- [SPARQL 1.1 Query Language Spec](https://www.w3.org/TR/sparql11-query/)
- [SPARQL Algebra (W3C)](https://www.w3.org/TR/sparql11-query/#sparqlAlgebra)
- [Dgraph GraphQL Schema Resolution](graphql/resolve/query_rewriter.go)
- [DQL Parser Architecture](dql/parser.go)
