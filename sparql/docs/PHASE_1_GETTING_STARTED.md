# Getting Started: SPARQL Algebra Implementation

**Goal**: Establish foundation for Phase 1 implementation  
**Timeline**: Week 1 prep

---

## What We're Building

Transform the SPARQL query pipeline from:

```
SPARQL String → AST → [serialize] → DQL String → [re-parse] → GraphQuery
```

To:

```
SPARQL String → AST → Algebra → [optimize] → DQL GraphQuery (with schema)
```

**Key Benefits**:

1. Schema-aware optimizations (like GraphQL)
2. Structured query rewriting (filter pushdown, join ordering)
3. Authorization rule integration (currently missing)
4. Foundation for ontology reasoning

---

## Phase 1: Algebra Foundation - What to Build First

### Step 1: Algebra Type System (3-4 days)

**Objective**: Define the algebra operators that represent SPARQL queries

**Create**: `sparql/algebra.go`

```go
package sparql

// AlgebraExpr is the base interface for all algebra operations
type AlgebraExpr interface {
	Accept(visitor AlgebraVisitor) interface{}
	String() string                              // For debugging
}

// Basic operators (follow W3C SPARQL algebra semantics)
type (
	// BGP represents a Basic Graph Pattern
	BGP struct {
		Triples []*Triple
	}

	// Join represents pattern combination
	Join struct {
		Left  AlgebraExpr
		Right AlgebraExpr
	}

	// Filter represents WHERE constraints
	Filter struct {
		Expr  Expression
		Input AlgebraExpr
	}

	// And more...
)

// AlgebraVisitor implements the visitor pattern
type AlgebraVisitor interface {
	VisitBGP(*BGP) interface{}
	VisitJoin(*Join) interface{}
	VisitFilter(*Filter) interface{}
	// ... etc
}
```

**Reference**: W3C SPARQL Algebra section:  
https://www.w3.org/TR/sparql11-query/#sparqlAlgebra

**Testing**:

- Can create algebra expressions
- Visitor pattern works
- String representation for debugging

---

### Step 2: AST to Algebra Converter (4-5 days)

**Objective**: Convert existing SPARQLQueryImpl AST to algebra form

**Extend**: `sparql/algebra.go`

```go
// ASTToAlgebra converts SPARQL AST to algebra
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	// Main translation logic
	// 1. Convert BGPs to Join chains
	// 2. Convert FILTER expressions
	// 3. Handle OPTIONAL → LeftJoin
	// 4. Handle UNION → Union operator
	// 5. Collect projections and aggregates

	root := buildFromPatterns(query.Patterns)

	if query.Filter != nil {
		root = &Filter{Expr: query.Filter, Input: root}
	}

	// ... apply aggregates, ordering, etc.

	return root, nil
}
```

**Testing**:

- Convert simple SELECT with WHERE
- Convert SELECT with FILTER
- Convert SELECT with OPTIONAL
- Convert SELECT with UNION
- Convert aggregates (COUNT, SUM, etc.)
- Compare to manual algebra specification

---

### Step 3: Algebra Validator (2-3 days)

**Objective**: Ensure algebra expressions are valid

**Extend**: `sparql/algebra.go`

```go
type AlgebraValidator struct{}

func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
	// Check:
	// 1. All referenced variables are defined
	// 2. No circular dependencies
	// 3. Filter expressions reference existing variables
	// 4. Aggregate functions use valid operators
	// 5. Projections only reference defined variables
}
```

**Testing**:

- Valid expressions pass
- Missing variable detected
- Circular reference detected
- Clear error messages

---

### Step 4: Integration Point (1-2 days)

**Objective**: Create context object to carry metadata through pipeline

**Create**: `sparql/context.go`

```go
package sparql

import (
	"context"
	"github.com/dgraph-io/dgraph/v25/graphql/schema"
)

type SPARQLExecutionContext struct {
	// Query context
	Ctx context.Context

	// Schema information (for Phase 3)
	Schema *schema.Schema

	// Auth context (for Phase 4)
	UserID    string
	GroupIDs  []string
	Namespace uint64

	// Compilation options
	Options TranslateOptions
}
```

**Testing**:

- Context creation
- Schema loading (mock)
- Auth info extraction (mock)

---

### Step 5: Comprehensive Testing (2-3 days)

**Create**: `sparql/algebra_test.go`

Test structure:

```go
// Test algebra creation from AST
func TestASTToAlgebraBasicSelect(t *testing.T) {}
func TestASTToAlgebraWithFilter(t *testing.T) {}
func TestASTToAlgebraWithOptional(t *testing.T) {}
func TestASTToAlgebraWithUnion(t *testing.T) {}
func TestASTToAlgebraAggregates(t *testing.T) {}

// Test validator
func TestValidatorDetectsUndefinedVariable(t *testing.T) {}
func TestValidatorDetectsCircularDeps(t *testing.T) {}

// Test visitor pattern
func TestAlgebraVisitorTraversal(t *testing.T) {}
func TestAlgebraStringRepresentation(t *testing.T) {}
```

Target: **100+ test cases** across all algebra operators

---

## Concrete Example: Converting a SPARQL Query

### Input SPARQL

```sparql
SELECT ?person ?name WHERE {
  ?person rdf:type Person .
  ?person foaf:name ?name .
  FILTER (?name = "Alice")
}
```

### AST (Already Have)

```
SPARQLQueryImpl{
  Type: "SELECT",
  Variables: ["person", "name"],
  Patterns: [
    BGP{
      Triples: [
        Triple{Subject: ?person, Predicate: rdf:type, Object: Person},
        Triple{Subject: ?person, Predicate: foaf:name, Object: ?name}
      ]
    }
  ],
  Filter: EqExpression{Left: Variable("name"), Right: "Alice"}
}
```

### Target Algebra

```
Project(
  vars: [?person, ?name],
  input: Filter(
    expr: (?name = "Alice"),
    input: Join(
      left: BGP(Triple(?person, rdf:type, Person)),
      right: BGP(Triple(?person, foaf:name, ?name))
    )
  )
)
```

### Code That Builds It

```go
func TestConvertSimpleQuery(t *testing.T) {
	ast := &SPARQLQueryImpl{
		Type:      "SELECT",
		Variables: []string{"person", "name"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?person", Predicate: "rdf:type", Object: "Person"},
					{Subject: "?person", Predicate: "foaf:name", Object: "?name"},
				},
			},
		},
		Filter: &EqExpression{
			Left:  Variable("?name"),
			Right: StringLiteral("Alice"),
		},
	}

	algebra, err := ASTToAlgebra(ast)
	require.NoError(t, err)

	// Verify structure
	project, ok := algebra.(*Project)
	require.True(t, ok)
	require.Equal(t, []string{"?person", "?name"}, project.Vars)

	filter, ok := project.Input.(*Filter)
	require.True(t, ok)

	join, ok := filter.Input.(*Join)
	require.True(t, ok)
	require.Equal(t, 2, len(join.Left.(*BGP).Triples) + len(join.Right.(*BGP).Triples))
}
```

---

## Implementation Checklist for Week 1

### Day 1-2: Algebra Types

- [ ] Define algebraic expression interfaces and types
- [ ] Implement String() methods for debugging
- [ ] Write 10+ tests for type construction

### Day 3: Visitor Pattern

- [ ] Implement visitor interface
- [ ] Create concrete visitor examples
- [ ] Test traversal on sample algebras

### Day 4-5: AST→Algebra Conversion

- [ ] Implement ASTToAlgebra function
- [ ] Handle BGPs and Joins
- [ ] Handle Filters
- [ ] Write 30+ conversion tests

### Day 6-7: Validator & Context

- [ ] Implement AlgebraValidator
- [ ] Create SPARQLExecutionContext
- [ ] Write validator tests
- [ ] Write context tests

### Day 8: Integration & Cleanup

- [ ] Integrate new code with existing tests
- [ ] Ensure all existing SPARQL tests still pass
- [ ] Code review prep
- [ ] Documentation

### End of Week 1: Deliverables

- ✅ `sparql/algebra.go` (400-500 LOC)
- ✅ `sparql/context.go` (100 LOC)
- ✅ `sparql/algebra_test.go` (400+ LOC, 100+ test cases)
- ✅ All existing tests pass
- ✅ No performance regression

---

## Key Files to Study First

Before implementing, understand:

1. **SPARQL AST Structure**
   - `sparql/ast.go` - SPARQLQueryImpl, BGP, Triple, etc.
   - `sparql/antlr_adapter.go` - How AST is created from parsing

2. **Current Translation**
   - `sparql/translator_extended.go` - Current AST→DQL conversion
   - Study how it handles each pattern type

3. **DQL Structure**
   - `dql/parser.go` - GraphQuery, FilterTree structures
   - Understand what we're targeting

4. **GraphQL Auth Pattern**
   - `graphql/resolve/query_rewriter.go` - addCommonRules, authRewriter
   - This is the pattern we'll apply to SPARQL in Phase 4

---

## Resources & References

### SPARQL Specification

- W3C SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
- SPARQL Query Language: https://www.w3.org/TR/sparql11-query/

### Design Patterns

- Visitor Pattern: https://en.wikipedia.org/wiki/Visitor_pattern
- Abstract Syntax Tree: https://en.wikipedia.org/wiki/Abstract_syntax_tree

### Code Examples

- Similar translator work: `graphql/resolve/query_rewriter.go`
- Existing AST handling: `sparql/translator_extended.go`

---

## Questions to Answer Before Starting

1. **Variable Naming**: Should algebra variables keep the `?` prefix (SPARQL convention) or match
   DQL style?
   - Decision: Keep `?` in algebra, strip when compiling to DQL (Phase 3)

2. **Expression Representation**: Reuse existing `Expression` interface or new algebra expressions?
   - Decision: Reuse existing for now; refactor in Phase 2 if needed

3. **Error Handling**: Early validation or lazy evaluation?
   - Decision: Validate immediately (fail fast); helps debugging

4. **Optimization Hooks**: Where can optimizers hook in?
   - Decision: Visitor pattern on algebra; each optimizer is a visitor

5. **Test Coverage**: How comprehensive?
   - Decision: Aim for 100+ test cases; cover all operators and combinations

---

## Success Metrics for Phase 1

- [ ] **Functionality**: All algebra operators working
- [ ] **Testing**: 100+ test cases, 100% passing
- [ ] **Quality**: No regressions in existing SPARQL tests
- [ ] **Code**: <10 code review comments for style/design
- [ ] **Documentation**: Each operator has docstring explaining semantics
- [ ] **Performance**: Algebra compilation <10ms for typical queries

---

## Next Steps After Phase 1

Once Phase 1 is complete and merged:

1. **Plan Phase 2**: Detail optimization rules
2. **Benchmark baseline**: Measure current query execution time
3. **Design Phase 3**: Schema integration approach
4. **Study Phase 4**: Deep dive into GraphQL auth implementation

---

## Contact & Questions

For clarifications on this roadmap:

- Review the [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) for full context
- Check [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md) for all phases
- Discuss design decisions in code review
