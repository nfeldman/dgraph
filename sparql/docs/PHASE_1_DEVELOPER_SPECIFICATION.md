# Phase 1 SPARQL Algebra: Developer Implementation Specification

**Status**: Ready for Implementation  
**Branch**: `feature/sparql-phase1-dev`  
**Duration**: 10-15 days (1 developer) or 6-8 days (2 developers)  
**Target**: ~1,300 LOC across 4 files

---

## Overview

This document provides the detailed implementation specification for Phase 1 of the SPARQL Algebra
project. It covers Steps 1-4 sequentially with concrete design decisions, code structure, and
testing requirements.

**Read this document in order before starting implementation.**

---

## Implementation Steps (Sequential)

### Step 1: Algebra Type System (Days 1-4)

**Objective**: Define all algebraic expression types that will represent SPARQL queries.

**Files to Create**: `sparql/algebra.go` (~400-500 LOC)

**Core Types**:

```go
// Base interface that all algebra expressions implement
type AlgebraExpr interface {
    Accept(visitor AlgebraVisitor) interface{}
    String() string  // For debugging
}

// 11 Operator Types (W3C SPARQL Algebra)
type BGP struct {
    Patterns []*Triple  // Basic graph pattern
}

type Join struct {
    Left  AlgebraExpr
    Right AlgebraExpr
}

type Filter struct {
    Expr  Expression  // Reuse existing Expression type
    Input AlgebraExpr
}

type LeftJoin struct {
    Input    AlgebraExpr
    Patterns []*Triple
    Filter   Expression
}

type Union struct {
    Alternatives []AlgebraExpr
}

type Project struct {
    Vars  []string
    Input AlgebraExpr
}

type Aggregate struct {
    Op    string       // "COUNT", "SUM", "MIN", "MAX", "AVG"
    Expr  Expression   // What to aggregate
    Group []string     // GROUP BY variables
    Input AlgebraExpr
}

type Bind struct {
    Var   string
    Expr  Expression
    Input AlgebraExpr
}

type Distinct struct {
    Input AlgebraExpr
}

type OrderBy struct {
    Expressions []Expression
    Ascending   []bool
    Input       AlgebraExpr
}

type Limit struct {
    Count  int
    Offset int
    Input  AlgebraExpr
}
```

**Key Design Decisions**:

1. **Reuse Existing Expression Type**: Use existing `Expression` interface from filter_extended.go
   for filter expressions and BIND expressions. Minimizes new types.

2. **String() Implementation**: Each type implements `String()` for debugging. Format:
   `OperatorName(child1, child2, ...)`.

3. **Variable Tracking**: Add helper methods:

   ```go
   func (e *YourExpr) Variables() []string  // All referenced vars
   func (e *YourExpr) DefinedVars() []string // Vars defined by this expr
   ```

4. **No Optimization Yet**: Step 1 types are pure representation, not optimized.

**Testing Strategy**:

- Create 10+ types, verify each can:
  - Be constructed
  - Print via String()
  - Have variables extracted
  - Accept a visitor

---

### Step 2: Visitor Pattern Infrastructure (Days 5-7)

**Objective**: Implement visitor pattern for traversing algebra expressions.

**Files to Create**: `sparql/algebra_visitor.go` (~100-150 LOC)

**Core Interface**:

```go
// AlgebraVisitor follows the visitor pattern for algebra expressions
type AlgebraVisitor interface {
    VisitBGP(*BGP) interface{}
    VisitJoin(*Join) interface{}
    VisitFilter(*Filter) interface{}
    VisitLeftJoin(*LeftJoin) interface{}
    VisitUnion(*Union) interface{}
    VisitProject(*Project) interface{}
    VisitAggregate(*Aggregate) interface{}
    VisitBind(*Bind) interface{}
    VisitDistinct(*Distinct) interface{}
    VisitOrderBy(*OrderBy) interface{}
    VisitLimit(*Limit) interface{}
}
```

**Concrete Implementations**:

1. **AlgebraPrinter**: Prints algebra tree with indentation

   ```go
   type AlgebraPrinter struct {
       indent int
   }

   func (p *AlgebraPrinter) VisitBGP(bgp *BGP) interface{} {
       // Print BGP with variable names
   }
   ```

2. **VariableCollector**: Collects all variables used in expression
   ```go
   type VariableCollector struct {
       variables map[string]bool
   }
   ```

**Key Design Decisions**:

1. **Two Traversal Types**:
   - Top-down: Parent visits children (default)
   - Bottom-up: Children results flow to parent (for optimization)

2. **Generic Return Type**: Use `interface{}` for flexibility, but each visitor documents return
   type.

3. **Stateful Visitors**: Visitors can maintain state (e.g., variable map, indentation level).

**Testing Strategy**:

- Test each visitor type:
  - Visitor visits all expression types
  - Return values are correct
  - State accumulates properly
  - Complex nested expressions work

---

### Step 3: AST to Algebra Converter (Days 8-12)

**Objective**: Convert existing SPARQL AST to algebraic form.

**Files to Create**:

- `sparql/algebra.go` (add function, ~300-400 LOC)
- Comprehensive converter tests in `sparql/algebra_test.go`

**Main Function Signature**:

```go
// ASTToAlgebra converts SPARQL AST to algebraic form
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
    // 1. Start with patterns
    // 2. Build BGP and Join nodes
    // 3. Apply FILTER
    // 4. Apply OPTIONAL → LeftJoin
    // 5. Apply UNION → Union
    // 6. Apply GROUP BY → Aggregate
    // 7. Apply BIND
    // 8. Apply ORDER BY
    // 9. Apply LIMIT/OFFSET
    // 10. Apply SELECT → Project
}
```

**10-Step Conversion Algorithm**:

1. **Parse Patterns** → Extract BGPs, OPTIONAL, UNION separately
2. **Build BGP Nodes** → Each pattern becomes a BGP with triples
3. **Join BGPs** → Create Join nodes linking BGPs (left-associative)
4. **Apply Filters** → Wrap with Filter nodes at appropriate points
5. **Handle OPTIONAL** → Convert to LeftJoin operators
6. **Handle UNION** → Create Union operators with alternatives
7. **Build Filter Expressions** → From WHERE clause
8. **Apply Aggregates** → Wrap with Aggregate for COUNT, SUM, etc.
9. **Apply Projections** → Top-level Project node for SELECT variables
10. **Apply Modifiers** → OrderBy, Limit, Distinct nodes

**Key Helper Functions**:

```go
// Build Join chains from multiple expressions
func joinExpressions(exprs []AlgebraExpr) AlgebraExpr

// Convert OPTIONAL patterns to LeftJoin
func optionalToLeftJoin(main, optional AlgebraExpr, filter Expression) AlgebraExpr

// Wrap expression with filter if filter is provided
func applyFilter(expr AlgebraExpr, filter Expression) AlgebraExpr

// Get filter expression from WHERE clause
func extractFilter(filterTree *dql.FilterTree) Expression
```

**Example Conversion**:

Input SPARQL:

```sparql
SELECT ?person ?name WHERE {
  ?person rdf:type Person .
  ?person foaf:name ?name .
  FILTER (?name = "Alice")
}
```

Expected Output Algebra:

```
Project(vars: [?person, ?name],
  Filter(expr: (?name = "Alice"),
    Join(
      BGP(?person rdf:type Person),
      BGP(?person foaf:name ?name)
    )
  )
)
```

**Testing Strategy** (30+ test cases):

- Simple SELECT with one triple
- SELECT with multiple triples (join)
- SELECT with FILTER
- SELECT with OPTIONAL
- SELECT with UNION
- SELECT with COUNT/GROUP BY
- SELECT with BIND
- SELECT with ORDER BY
- SELECT with LIMIT/OFFSET
- Combinations of above
- Edge cases (empty graph pattern, multiple filters, nested unions)

---

### Step 4: Algebra Validator (Days 13-15)

**Objective**: Validate that algebra expressions are semantically correct.

**Files to Create**: `sparql/algebra.go` (add validator, ~150-200 LOC)

**Validator Type**:

```go
// AlgebraValidator checks algebra expression semantic correctness
type AlgebraValidator struct {
    errors []string
}

func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
    // Visit all nodes, collect errors
    // Return combined error or nil
}
```

**Validation Rules**:

1. **Variable Scoping**:
   - Variables referenced in Filter must be defined by input expressions
   - Variables in Project must be defined somewhere in the tree
   - Variables in ORDER BY must be defined
   - Variables in aggregate expressions must be in GROUP BY or be aggregated

2. **Circular Dependencies**:
   - Detect cycles in BIND expressions
   - BIND variables must not depend on themselves

3. **Filter Validity**:
   - Filter expressions must reference only defined variables
   - Filter boolean expressions valid

4. **Aggregate Validity**:
   - COUNT, SUM, etc. applied to valid expressions
   - GROUP BY variables must be defined

5. **Type Compatibility**:
   - Binary operators (=, <, etc.) have compatible operands

**Validator Implementation**:

```go
type AlgebraValidator struct {
    definedVars  map[string]bool  // Track defined variables at each scope
    errors       []string
}

func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
    v.definedVars = make(map[string]bool)
    expr.Accept(v)
    if len(v.errors) > 0 {
        return fmt.Errorf("algebra validation errors:\n%s", strings.Join(v.errors, "\n"))
    }
    return nil
}

// Implement VisitXXX methods that:
// 1. Check for undefined variables
// 2. Add newly defined variables to map
// 3. Check sub-expressions
```

**Error Messages** (clear and actionable):

- "Variable ?x referenced in filter but not defined"
- "Variable ?x defined in BIND but appears in circular dependency"
- "Aggregate function COUNT missing valid expression"
- "Variable ?x in ORDER BY not defined in projection"

**Testing Strategy** (25+ test cases):

- Valid expressions pass
- Missing variable detected
- Circular dependency detected
- Invalid filter caught
- Invalid aggregate caught
- Type errors detected
- Error messages are clear
- Complex nested expressions validated correctly

---

## Supporting Infrastructure

### `sparql/context.go` (~100 LOC)

Create execution context to carry metadata through compilation:

```go
type SPARQLExecutionContext struct {
    // Query context
    Ctx context.Context

    // Schema (for Phase 3)
    Schema *schema.Schema

    // Auth context (for Phase 4)
    UserID    string
    GroupIDs  []string
    Namespace uint64

    // Options
    Options TranslateOptions
}

// Builder pattern for ergonomics
func NewSPARQLExecutionContext(ctx context.Context) *SPARQLExecutionContext {
    return &SPARQLExecutionContext{Ctx: ctx}
}

func (c *SPARQLExecutionContext) WithSchema(s *schema.Schema) *SPARQLExecutionContext {
    c.Schema = s
    return c
}

func (c *SPARQLExecutionContext) WithUser(uid string, gids []string) *SPARQLExecutionContext {
    c.UserID = uid
    c.GroupIDs = gids
    return c
}
```

---

## File Structure

| File                 | LOC             | Purpose                               |
| -------------------- | --------------- | ------------------------------------- |
| `algebra.go`         | 600-700         | Algebra types, converter, validator   |
| `algebra_visitor.go` | 100-150         | Visitor interface and implementations |
| `context.go`         | 100             | SPARQLExecutionContext                |
| `algebra_test.go`    | 400-500         | 100+ comprehensive tests              |
| **Total**            | **1,200-1,450** | Phase 1 foundation                    |

---

## Design Decisions & Justification

### Decision 1: Reuse Expression Interface

**Choice**: Use existing `Expression` type from filter_extended.go instead of creating new algebra
expression types for filters.

**Justification**:

- Expression already parses FILTER syntax
- Minimal new code needed
- Integrates with existing FILTER handling

### Decision 2: Visitor Pattern Over Recursive Descent

**Choice**: Implement visitor pattern instead of embedding visit logic in each type.

**Justification**:

- Separates concerns (expression structure vs. operations)
- Easy to add new operations (optimizers, compilers) without modifying algebra types
- Standard pattern used in compiler design

### Decision 3: String() For Debugging Over Serialization

**Choice**: Implement human-readable String() for debugging, not serialization.

**Justification**:

- Phase 1 doesn't need to serialize algebra
- String() helps during development/debugging
- Actual serialization will be to DQL in Phase 3

### Decision 4: Helper Methods Over Visitor For Variable Collection

**Choice**: Add Variables() methods to types instead of requiring a visitor.

**Justification**:

- Variable collection is so common it deserves easy API
- Can implement as shorthand: expr.Accept(&VariableCollector{})
- Follows principle of least surprise

### Decision 5: Left-Associative Joins

**Choice**: When combining multiple BGPs, create left-associative join trees.

**Justification**:

- Standard in query processing (PostgreSQL, DuckDB)
- Simplifies implementation
- Can be optimized to bushy trees in Phase 2

### Decision 6: Include Limit/Offset in Algebra

**Choice**: Model LIMIT and OFFSET as algebraic operators, not query hints.

**Justification**:

- W3C SPARQL Algebra includes them
- Enables correct optimization of pagination
- Proper semantics (apply after ordering)

---

## Integration Points

### With Existing SPARQL Code

- **antlr_adapter.go**: Input (SPARQLQueryImpl)
- **translator_extended.go**: Should eventually call ASTToAlgebra (Phase 2)
- **filter_extended.go**: Reuses Expression interface
- **ast.go**: Input (Triple, Expression types)

### With Dgraph Core

- **dql/parser.go**: Target format (GraphQuery)
- **graphql/resolve/query_rewriter.go**: Reference pattern

---

## Testing Strategy

### Unit Tests by Category

1. **Type Construction** (15 tests)
   - Each type can be created
   - String() produces valid output
   - Variables() returns correct set

2. **Visitor Pattern** (20 tests)
   - All visitors work on all expression types
   - Traversal order correct
   - State accumulates properly

3. **AST Conversion** (40+ tests)
   - Simple SELECT → Project(Join(BGP, BGP))
   - FILTER → Filter node placement
   - OPTIONAL → LeftJoin
   - UNION → Union
   - COUNT → Aggregate
   - BIND → Bind
   - LIMIT/OFFSET → Limit
   - Complex combinations

4. **Validator** (25 tests)
   - Valid expressions pass
   - Undefined variable detected
   - Circular dependency detected
   - Invalid filter caught
   - Clear error messages

### Running Tests

```bash
# Run all Phase 1 tests
go test ./sparql -v -run "Algebra"

# Run specific test
go test ./sparql -v -run "TestASTToAlgebraSelect"

# Run with coverage
go test ./sparql -cover -run "Algebra"
```

**Coverage Target**: >85% on algebra.go, algebra_visitor.go, context.go

---

## Success Criteria Checklist

### Code Quality

- [ ] All algebra operators implemented (11 types)
- [ ] All types have String() and Accept() methods
- [ ] Helper methods for variable tracking
- [ ] Visitor interface with 11 Visit methods
- [ ] Concrete visitors (Printer, VariableCollector)
- [ ] Validator with semantic checks
- [ ] Code follows Dgraph conventions
- [ ] No compiler warnings

### Testing

- [ ] 100+ test cases (comprehensive)
- [ ] All tests pass (go test ./sparql -v)
- [ ] > 85% code coverage
- [ ] Tests cover happy path and error cases
- [ ] Error messages are clear

### Integration

- [ ] Compiles with existing SPARQL code
- [ ] Existing SPARQL tests still pass
- [ ] No modifications to existing code (additive only)
- [ ] Clear integration points with Phase 2

### Documentation

- [ ] All types have docstrings
- [ ] All functions have docstrings
- [ ] Example test cases show usage
- [ ] Design decisions documented

---

## Common Pitfalls to Avoid

1. **Don't serialize algebra to strings**: Phase 1 is in-memory representation only
2. **Don't optimize yet**: Stay focused on correct representation (Phase 2)
3. **Don't change existing code**: Only add new files and types
4. **Don't over-engineer**: Solver pattern more complex than needed for Phase 1
5. **Don't skip testing**: 100+ tests is not arbitrary; ensures foundation is solid

---

## Timeline Estimate (1 Developer)

| Step                 | Days | Cumulative |
| -------------------- | ---- | ---------- |
| 1. Algebra Types     | 4    | 4          |
| 2. Visitor Pattern   | 3    | 7          |
| 3. AST Converter     | 5    | 12         |
| 4. Validator         | 3    | 15         |
| Buffer for debugging | 2    | 17         |

**Realistic estimate**: 10-15 days (not 10-15 working days)

---

## Questions Before Starting?

If you have clarifications needed:

1. Variable scoping in OPTIONAL?
2. How to handle nested UNION?
3. Expression type details?
4. Visitor pattern specifics?

Refer to:

- [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - Full context
- [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) - Examples
- W3C SPARQL Algebra - https://www.w3.org/TR/sparql11-query/#sparqlAlgebra

---

## Next Steps After Phase 1

Once Phase 1 is complete:

1. Merge to feature/sparql-antlr
2. Create feature/sparql-phase2-optimization
3. Begin Phase 2: Algebra Rewriter (filter pushdown, join ordering)

See [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) for Phase 2 details.
