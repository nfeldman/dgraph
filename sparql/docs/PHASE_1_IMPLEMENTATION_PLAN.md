# Phase 1 Implementation Plan: SPARQL Algebra Foundation

**Prepared**: January 2026  
**Duration**: 8-10 days  
**Target Completion**: Mergedable branch with 100% test coverage

---

## Executive Summary

This document provides a detailed, step-by-step implementation plan for Phase 1 of the SPARQL
Algebra project. The plan breaks down the work into 4 sequential steps, each with concrete
deliverables, design decisions, and testing strategies.

**Key Outcomes**:

- ✅ `sparql/algebra.go` (500-600 LOC) - All algebra operators + converter + validator
- ✅ `sparql/algebra_visitor.go` (100-150 LOC) - Visitor pattern infrastructure
- ✅ `sparql/context.go` (100 LOC) - Execution context carrying
- ✅ `sparql/algebra_test.go` (400-500 LOC) - 100+ comprehensive tests

---

## Part 1: Algebra Type System (3-4 Days)

### 1.1 Design Goals

**Purpose**: Define a complete set of algebraic expression types that represent all SPARQL query
patterns in a normalized form.

**Key Requirements**:

- Follow W3C SPARQL Algebra semantics (https://www.w3.org/TR/sparql11-query/#sparqlAlgebra)
- Composable: each operator accepts other AlgebraExpr as input
- Debuggable: all types implement String() for human-readable output
- Pattern-based: enable visitor pattern for rewriting and optimization
- Support variable tracking: track which variables are in scope at each node

**Design Philosophy**:

```
- Separate syntax (AST) from semantics (Algebra)
- Algebra = simplified, normalized representation
- Each operator = one semantic concept
- Composable = tree structure with clear data flow
```

### 1.2 Algebra Operators to Implement

**File**: `sparql/algebra.go`

#### Base Interface

```go
type AlgebraExpr interface {
	// Accept implements the Visitor pattern
	// Returns the visitor's result (for traversal and transformation)
	Accept(visitor AlgebraVisitor) interface{}

	// String returns a human-readable representation
	// Used for debugging and optimization verification
	// Format: OperatorName(details)
	// Example: Filter(eq(?name, "Alice"), Join(BGP(...), BGP(...)))
	String() string
}
```

**Why this interface?**

- `Accept()` enables the Visitor pattern for read operations and rewrites
- `String()` provides visibility into algebra structure for debugging
- Minimal interface = maximum flexibility for implementations

---

#### Algebra Operator Types

All operators should be **struct types** implementing `AlgebraExpr`:

**1. BGP (Basic Graph Pattern)**

```go
type BGP struct {
	// Triples is the set of triple patterns in this BGP
	Triples []*Triple
}

// Semantics:
// - Represents a set of triples that must ALL match
// - Variables in triples must be consistent across triples
// - Example: ?person rdf:type Person . ?person foaf:name ?name .
// - String output: BGP([?person:rdf:type:Person, ?person:foaf:name:?name])

// String() format: "BGP(T0, T1, T2)" where each Ti is a triple representation
```

**Design Notes**:

- Triple struct already exists in `ast.go`
- No need to redefine; reuse existing
- String() should format as: `BGP([t1, t2, ...])`

---

**2. Join**

```go
type Join struct {
	// Left represents the left operand
	Left AlgebraExpr

	// Right represents the right operand
	Right AlgebraExpr
}

// Semantics:
// - Represents pattern combination (cartesian product + filter by variable consistency)
// - All solutions from Left joined with solutions from Right on shared variables
// - In SPARQL Algebra: Join is the primary composition operator
// - Example: Join(BGP(?s ?p ?o), BGP(?s rdf:type ?c))
//           = patterns from first joined to patterns from second via shared ?s

// String() format: "Join(left_string, right_string)"
```

**Design Notes**:

- Join is fundamental to how SPARQL patterns combine
- Multiple triples = a chain of Join operators in algebra
- Example: triple1, triple2, triple3 → Join(triple1, Join(triple2, triple3))

---

**3. Filter**

```go
type Filter struct {
	// Expr is the SPARQL filter expression (boolean condition)
	// Type: string for now (since expressions are strings in current AST)
	// Later: could be refactored to structured expression AST
	Expr string

	// Input is the algebra expression being filtered
	Input AlgebraExpr
}

// Semantics:
// - Filters solutions from Input by the boolean condition in Expr
// - Example: Filter("?name = 'Alice'", BGP(...))
//           = only return solutions where name equals 'Alice'
// - Lazy execution: can be pushed down to improve performance

// String() format: "Filter(condition, input_string)"
```

**Design Notes**:

- Expression storage: use string for now to match existing FilterPattern
- Variables in Expr must be defined by Input
- String format: `Filter("?name = 'Alice'", BGP(...))`

---

**4. LeftJoin**

```go
type LeftJoin struct {
	// Input is the left side (must-match patterns)
	Input AlgebraExpr

	// Patterns are the optional patterns (may or may not match)
	Patterns []*Triple

	// Filter is an optional filter applied only to the right side
	Filter string
}

// Semantics:
// - SPARQL OPTIONAL { ... } → LeftJoin
// - Keep all solutions from Input
// - For each solution, try to match Patterns
// - If match succeeds, add matched variables to solution
// - If match fails, solution still appears (with unbound optional variables)
// - Example: LeftJoin(BGP(?s ?p ?o), [?s foaf:age ?age], "")
//           = all solutions from BGP, plus ?age if person has age

// String() format: "LeftJoin(input, patterns, filter)"
```

**Design Notes**:

- Patterns stored as Triple list (not AlgebraExpr) for simplicity
- Filter is optional (may be empty string if no FILTER on optional)
- Key semantic: preserves solutions even if right side doesn't match

---

**5. Union**

```go
type Union struct {
	// Alternatives is the list of query alternatives
	Alternatives []AlgebraExpr
}

// Semantics:
// - SPARQL UNION { ... } { ... } → Union
// - Represents disjunction: match any alternative
// - Solutions = union of solutions from all alternatives
// - Example: Union(BGP(?s rdf:type Person), BGP(?s rdf:type Agent))
//           = all things that are either Person or Agent

// String() format: "Union(alt1_string | alt2_string | alt3_string)"
```

**Design Notes**:

- All alternatives are composable AlgebraExpr
- Result is set union (no duplicates in SPARQL semantics)
- Separators in String(): use " | " to show alternatives

---

**6. Project**

```go
type Project struct {
	// Vars is the list of variables to keep in results
	// Example: ["?person", "?name"]
	Vars []string

	// Input is the algebra expression whose results to project
	Input AlgebraExpr
}

// Semantics:
// - SELECT ?person ?name → Project(vars=[?person, ?name], input=...)
// - Filters solution bindings to only keep specified variables
// - Example: Project([?person], Join(...))
//           = keep only ?person from solutions; drop everything else

// String() format: "Project([var1, var2, ...], input_string)"
```

**Design Notes**:

- Variables stored as []string (with ? prefix to match SPARQL convention)
- Maps to SELECT clause in SPARQL
- Applied at top level of algebra tree

---

**7. Aggregate**

```go
type Aggregate struct {
	// Op is the aggregation operation: "COUNT", "SUM", "MIN", "MAX", "AVG"
	Op string

	// Expr is the expression to aggregate (often a variable)
	Expr string

	// Group is the list of GROUP BY variables
	// If empty, aggregate all solutions into single result
	Group []string

	// Input is the algebra expression to aggregate
	Input AlgebraExpr
}

// Semantics:
// - GROUP BY ?type; COUNT(*) → Aggregate(COUNT, *, group=[?type], input=...)
// - Combines multiple solutions into summary statistics
// - GROUP BY variables determine grouping; others are aggregated
// - Example: Aggregate(COUNT, *count, group=[?type], input=BGP(...))
//           = count solutions grouped by type

// String() format: "Aggregate(COUNT, *count, GROUP_BY[?type], input_string)"
```

**Design Notes**:

- Expr = what to aggregate (usually "\*" for COUNT, variable name for others)
- Group = empty for non-grouped aggregates (COUNT(\*) with no GROUP BY)
- Aggregate must appear at specific points in algebra tree

---

**8. Bind**

```go
type Bind struct {
	// Var is the output variable name (with ? prefix)
	Var string

	// Expr is the expression to evaluate (math or function)
	Expr string

	// Input is the algebra expression providing variable bindings
	Input AlgebraExpr
}

// Semantics:
// - BIND (?expr AS ?result) → Bind(var=?result, expr=?expr, input=...)
// - Evaluates an expression and binds result to a new variable
// - Expression can reference variables from Input
// - Example: Bind(var=?age, expr="?birthYear - 2024", input=...)
//           = compute age from birth year

// String() format: "Bind(?var AS expression, input_string)"
```

**Design Notes**:

- Expr is string (can be math like "?x + 1" or function call)
- New variable added to solution bindings
- Can be chained: multiple Bind operators

---

**9. Distinct**

```go
type Distinct struct {
	// Input is the algebra expression whose results to deduplicate
	Input AlgebraExpr
}

// Semantics:
// - SELECT DISTINCT → Distinct
// - Removes duplicate solutions (same variable bindings)
// - Example: Distinct(Join(...))
//           = keep only unique solution combinations

// String() format: "Distinct(input_string)"
```

**Design Notes**:

- Simple wrapper around input
- Applied when DISTINCT modifier is present in SELECT

---

**10. OrderBy**

```go
type OrderBy struct {
	// Expressions is the list of ORDER BY expressions
	// Each is a variable name or computed expression
	Expressions []string

	// Ascending indicates sort direction for each expression
	// Ascending[i] = true means ASC, false means DESC
	Ascending []bool

	// Input is the algebra expression to sort
	Input AlgebraExpr
}

// Semantics:
// - ORDER BY ?name ASC, ?age DESC → OrderBy(...)
// - Sorts solutions by specified expressions
// - Multiple expressions = multi-level sort
// - Example: OrderBy(exprs=[?name, ?age], asc=[true, false], input=...)
//           = sort by name ascending, then by age descending

// String() format: "OrderBy([?name ASC, ?age DESC], input_string)"
```

**Design Notes**:

- Expressions stored as strings (variable names or computed expressions)
- Ascending[i] corresponds to Expressions[i]
- Applied before Limit if both present

---

**11. Limit**

```go
type Limit struct {
	// Count is the maximum number of solutions to return
	Count int

	// Offset is the number of solutions to skip
	Offset int

	// Input is the algebra expression to limit
	Input AlgebraExpr
}

// Semantics:
// - LIMIT n OFFSET m → Limit(count=n, offset=m, input=...)
// - Returns at most Count solutions, skipping first Offset
// - Example: Limit(count=10, offset=5, input=...)
//           = return solutions 6-15 (skip first 5, return next 10)

// String() format: "Limit(count=n, offset=m, input_string)"
```

**Design Notes**:

- Typically applied at top level
- Offset can be 0 if no OFFSET clause
- Should be applied AFTER OrderBy if both present

---

### 1.3 Implementation Checklist: Algebra Types

**Step 1: Define Interface**

- [ ] Create `AlgebraExpr` interface with `Accept()` and `String()`
- [ ] Document expected String() format

**Step 2: Implement All Operators**

- [ ] BGP struct + methods
- [ ] Join struct + methods
- [ ] Filter struct + methods
- [ ] LeftJoin struct + methods
- [ ] Union struct + methods
- [ ] Project struct + methods
- [ ] Aggregate struct + methods
- [ ] Bind struct + methods
- [ ] Distinct struct + methods
- [ ] OrderBy struct + methods
- [ ] Limit struct + methods

**Step 3: String() Methods**

- [ ] Each operator has clear, parseable String() format
- [ ] Can reconstruct algebra structure from String() output (for debugging)
- [ ] Handle nested operators properly (indent and parentheses)

**Step 4: Testing (Part 1)**

- [ ] Create all types successfully
- [ ] String() produces expected output
- [ ] Can nest operators (e.g., Filter wrapping Join)
- [ ] 15+ test cases for type construction

**Example Test**:

```go
func TestBGPCreation(t *testing.T) {
	bgp := &BGP{
		Triples: []*Triple{
			{Subject: "?s", Predicate: "rdf:type", Object: "Person"},
			{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
		},
	}

	result := bgp.String()
	require.Contains(t, result, "BGP")
	require.Contains(t, result, "?s")
	require.Contains(t, result, "foaf:name")
}
```

---

## Part 2: Visitor Pattern Infrastructure (2-3 Days)

### 2.1 Design Goals

**Purpose**: Implement the Visitor pattern to enable non-destructive traversal and transformation of
algebra trees.

**Benefits**:

- Separates concern: algebra structure vs. operations on structure
- Enables multiple operations (printing, validation, optimization) without modifying operators
- Foundation for future optimization passes

**Design Philosophy**:

```
- Visitor = interface defining operations on all operator types
- Concrete Visitor = implementation doing specific work
- AlgebraExpr.Accept(visitor) = dispatch to correct visitor method
- Result can be anything (bool, string, algebra expr, error, etc.)
```

### 2.2 Visitor Interface Definition

**File**: `sparql/algebra_visitor.go`

```go
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

**Key Design Decisions**:

1. **Return type is `interface{}`**: Allows visitor to return anything (error, bool, new
   AlgebraExpr, string, etc.)
2. **Method per operator**: Each operator has dedicated Visit method
3. **Naming convention**: VisitOperatorName (e.g., VisitBGP, VisitJoin)

---

### 2.3 AlgebraExpr Accept Methods

Each algebra operator must implement:

```go
func (bgp *BGP) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitBGP(bgp)
}
```

**Pattern**: All Accept() implementations follow this same pattern - just dispatch to correct
visitor method.

---

### 2.4 Concrete Visitor: AlgebraPrinter

**Purpose**: Print algebra tree in human-readable format

**File**: `sparql/algebra_visitor.go`

```go
type AlgebraPrinter struct {
	indent int  // Track indentation level
	output strings.Builder
}

// NewAlgebraPrinter creates a new printer
func NewAlgebraPrinter() *AlgebraPrinter {
	return &AlgebraPrinter{indent: 0}
}

// Print returns formatted algebra tree
func (p *AlgebraPrinter) Print(expr AlgebraExpr) string {
	p.Visit(expr)
	return p.output.String()
}

// Each VisitXxx method:
// 1. Writes operator name and details
// 2. Increments indent
// 3. Visits child expressions
// 4. Decrements indent

// Example:
func (p *AlgebraPrinter) VisitFilter(f *Filter) interface{} {
	p.writeLine(fmt.Sprintf("Filter(%s)", f.Expr))
	p.indent++
	f.Input.Accept(p)
	p.indent--
	return nil
}
```

**Output Format**:

```
Project([?person, ?name])
  Filter(?name = "Alice")
    Join
      BGP([?person:rdf:type:Person])
      BGP([?person:foaf:name:?name])
```

**Implementation Details**:

- Use `strings.Builder` for efficiency
- Helper method `writeLine()` to handle indentation
- Each operator prints itself, then calls Accept on children
- Children traverse recursively

---

### 2.5 Example Visitor Tests

```go
// Test 1: Simple traversal
func TestVisitorSimpleTraversal(t *testing.T) {
	bgp := &BGP{...}
	printer := NewAlgebraPrinter()
	result := printer.Print(bgp)
	require.Contains(t, result, "BGP")
}

// Test 2: Nested structure traversal
func TestVisitorNestedStructure(t *testing.T) {
	algebra := &Project{
		Vars: []string{"?s"},
		Input: &Filter{
			Expr: "?name = 'Alice'",
			Input: &Join{
				Left: &BGP{...},
				Right: &BGP{...},
			},
		},
	}

	printer := NewAlgebraPrinter()
	result := printer.Print(algebra)

	// Verify structure in output
	require.Contains(t, result, "Project")
	require.Contains(t, result, "Filter")
	require.Contains(t, result, "Join")
	require.Contains(t, result, "BGP")
}
```

---

### 2.6 Implementation Checklist: Visitor Pattern

**Step 1: Define Interface**

- [ ] Create `AlgebraVisitor` interface
- [ ] One Visit method per operator type
- [ ] Document visitor contract

**Step 2: Implement Accept() on All Operators**

- [ ] BGP.Accept()
- [ ] Join.Accept()
- [ ] Filter.Accept()
- [ ] LeftJoin.Accept()
- [ ] Union.Accept()
- [ ] Project.Accept()
- [ ] Aggregate.Accept()
- [ ] Bind.Accept()
- [ ] Distinct.Accept()
- [ ] OrderBy.Accept()
- [ ] Limit.Accept()

**Step 3: Implement AlgebraPrinter**

- [ ] AlgebraPrinter struct with indent tracking
- [ ] All Visit methods (11 total)
- [ ] Helper method for indented output
- [ ] Print() method to get final string

**Step 4: Testing**

- [ ] Test simple expressions
- [ ] Test deeply nested expressions
- [ ] Verify indentation and formatting
- [ ] 10+ test cases

---

## Part 3: AST to Algebra Converter (4-5 Days)

### 3.1 Design Goals

**Purpose**: Convert SPARQLQueryImpl AST (from parser) into normalized AlgebraExpr form.

**Key Challenge**: SPARQLQueryImpl represents query syntactically; we need to convert to semantic
algebra form.

**Example Transformation**:

```
Input AST:
  SPARQLQueryImpl{
    Qtype: "SELECT",
    Projs: ["?person", "?name"],
    Patterns: [BGP, OptionalPattern, FilterPattern, ...],
    Distinct: true,
    OrderBy: ["?name"],
    Limit: 10,
  }

Output Algebra:
  Limit(count=10,
    OrderBy(exprs=[?name], asc=[true],
      Distinct(
        Project(vars=[?person, ?name],
          Filter(expr=...,
            Join(BGP(...), LeftJoin(...), ...)
          )
        )
      )
    )
  )
```

### 3.2 ASTToAlgebra Function Signature

**File**: `sparql/algebra.go`

```go
// ASTToAlgebra converts a SPARQL AST query to an algebra expression
//
// Parameters:
//   query - SPARQLQueryImpl from parser (contains Patterns, Filter, etc.)
//
// Returns:
//   AlgebraExpr - Normalized algebra tree ready for optimization
//   error - If conversion fails (e.g., invalid patterns, undefined vars)
//
// Algorithm:
//   1. Extract pattern list (Patterns or fall back to Bgps)
//   2. Convert patterns to algebra (recursive helper)
//   3. Apply filters
//   4. Apply aggregates
//   5. Apply BIND expressions
//   6. Apply DISTINCT
//   7. Apply ORDER BY
//   8. Apply LIMIT
//   9. Apply projection (SELECT variables)
//
// Supported pattern types:
//   - BGP: identity (BGP → BGP algebra)
//   - OptionalPattern: → LeftJoin algebra
//   - UnionPattern: → Union algebra
//   - FilterPattern: → Filter algebra
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	// Implementation...
}
```

### 3.3 Conversion Algorithm - Step by Step

#### Step 3.3.1: Pattern Conversion

**Function**: `convertPatterns(patterns []GraphPattern) (AlgebraExpr, error)`

**Algorithm**:

```
1. If patterns is empty:
   - Return error (no patterns)

2. If patterns has 1 element:
   - Convert that pattern
   - Return result

3. If patterns has >1 element:
   - Convert first pattern → expr1
   - Convert rest recursively → expr2
   - Return Join(expr1, expr2)
```

**Pattern Type Handling**:

```go
// BGP Pattern → BGP Algebra (identity)
case *BGP:
  return &BGP{Triples: pattern.Triples}, nil

// OptionalPattern → LeftJoin Algebra
case *OptionalPattern:
  // Convert the optional patterns to triples
  // Create LeftJoin with those triples
  return &LeftJoin{
    Patterns: extractTriplesFromPatterns(pattern.Patterns),
    Filter: "", // Will be set if FilterPattern is inside Optional
  }, nil

// UnionPattern → Union Algebra
case *UnionPattern:
  // Each alternative is a []GraphPattern
  // Convert each alternative to algebra
  // Combine into Union
  alternatives := make([]AlgebraExpr, len(pattern.Alternatives))
  for i, alt := range pattern.Alternatives {
    alternatives[i] = convertPatterns(alt)
  }
  return &Union{Alternatives: alternatives}, nil

// FilterPattern → Filter Algebra (applied to input)
case *FilterPattern:
  // This is tricky - FilterPattern usually wraps other patterns
  // In our AST structure, we need to extract filters separately
  // See next section for detail
```

**Key Decision**: How to handle FilterPattern?

In the AST, a FilterPattern might appear inline with other patterns. In algebra, we want filters
separate (wrapped around their input). Solution:

```go
// Two-pass approach:
// 1. Extract all FilterPatterns into a list
// 2. Extract all other patterns, build algebra from them
// 3. Wrap result in Filters
```

#### Step 3.3.2: Filter Application

**Function**: `applyFilters(root AlgebraExpr, filters []FilterPattern) AlgebraExpr`

**Algorithm**:

```
1. If no filters:
   - Return root unchanged

2. For each filter (in reverse order, so they nest correctly):
   - Create Filter(expr=filter.Expression, input=root)
   - Update root to new Filter

3. Return wrapped root
```

**Example**:

```
Input: root=Join(...), filters=[f1, f2]
Output: Filter(f2, Filter(f1, root))
         (f1 is "inner" filter, f2 is "outer")
```

#### Step 3.3.3: Projection Application

**Function**: `applyProjection(root AlgebraExpr, projVars []string) AlgebraExpr`

**Algorithm**:

```
1. If projVars is empty:
   - Return root (no projection)

2. If projVars is not empty:
   - Return Project(vars=projVars, input=root)
```

**Key Point**: Project is typically the outermost operator in SELECT queries.

#### Step 3.3.4: Distinct Application

**Function**: `applyDistinct(root AlgebraExpr, distinct bool) AlgebraExpr`

**Algorithm**:

```
1. If not distinct:
   - Return root

2. If distinct:
   - Return Distinct(input=root)
```

**Ordering Note**: Distinct should be applied AFTER Project, BEFORE OrderBy.

#### Step 3.3.5: OrderBy Application

**Function**: `applyOrderBy(root AlgebraExpr, orderVars []string) AlgebraExpr`

**Algorithm**:

```
1. If orderVars is empty:
   - Return root (no sorting)

2. Parse each order var:
   - Extract expression and direction (ASC/DESC)
   - Default is ASC

3. Return OrderBy(expressions=exprs, ascending=dirs, input=root)
```

**Parsing Order Variables**:

```
Input: ["?name", "DESC(?age)"]
Output:
  expressions: ["?name", "?age"]
  ascending: [true, false]
```

#### Step 3.3.6: Aggregate Application

**Function**: `applyAggregates(root AlgebraExpr, aggs []*Aggregate) AlgebraExpr`

**Algorithm**:

```
1. If no aggregates:
   - Return root

2. If aggregates present:
   - Extract GROUP BY variables from aggregates
   - Create Aggregate(op=..., expr=..., group=groupVars, input=root)
   - Return Aggregate
```

**Design Note**: In SPARQL, all aggregates in a query must have the same GROUP BY. If mixed with
non-aggregates, GROUP BY is implicit over all non-aggregates.

#### Step 3.3.7: Limit/Offset Application

**Function**: `applyLimitOffset(root AlgebraExpr, limit, offset int) AlgebraExpr`

**Algorithm**:

```
1. If limit <= 0 and offset == 0:
   - Return root (no limit)

2. If limit > 0 or offset > 0:
   - Return Limit(count=limit, offset=offset, input=root)
```

**Key Point**: LIMIT/OFFSET is typically the outermost operator (applied last).

### 3.4 Complete AST→Algebra Function Structure

```go
func ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) {
	// Validation
	if query == nil {
		return nil, fmt.Errorf("query is nil")
	}

	if query.Qtype != "SELECT" && query.Qtype != "ASK" {
		return nil, fmt.Errorf("unsupported query type: %s", query.Qtype)
	}

	// Extract pattern list (new format or fallback)
	var patterns []GraphPattern
	if len(query.Patterns) > 0 {
		patterns = query.Patterns
	} else if len(query.Bgps) > 0 {
		// Backwards compatibility: convert old BGP list
		for _, bgp := range query.Bgps {
			patterns = append(patterns, bgp)
		}
	} else {
		return nil, fmt.Errorf("no patterns provided")
	}

	// Separate patterns by type
	filterPatterns, otherPatterns := separateFilterPatterns(patterns)

	// Build algebra from non-filter patterns
	algebra, err := convertPatterns(otherPatterns)
	if err != nil {
		return nil, fmt.Errorf("converting patterns: %w", err)
	}

	// Apply filters
	algebra = applyFilters(algebra, filterPatterns)

	// Apply BIND expressions
	for _, bind := range query.Binds {
		algebra = &Bind{
			Var: bind.Variable,
			Expr: bind.Expression,
			Input: algebra,
		}
	}

	// Apply aggregates (if any)
	if len(query.Aggregates) > 0 {
		groupVars := extractGroupByVars(query.Aggregates)
		// Note: Only apply first aggregate; SPARQL typically has one per query
		agg := query.Aggregates[0]
		algebra = &Aggregate{
			Op: strings.ToUpper(agg.Function),
			Expr: agg.Variable,
			Group: groupVars,
			Input: algebra,
		}
	}

	// Apply DISTINCT
	if query.Distinct {
		algebra = &Distinct{Input: algebra}
	}

	// Apply ORDER BY
	if len(query.OrderBy) > 0 {
		algebra = applyOrderBy(algebra, query.OrderBy)
	}

	// Apply LIMIT/OFFSET
	if query.Limit > 0 || query.Offset > 0 {
		algebra = &Limit{
			Count: query.Limit,
			Offset: query.Offset,
			Input: algebra,
		}
	}

	// Apply projection (outermost)
	if len(query.Projs) > 0 {
		algebra = &Project{
			Vars: query.Projs,
			Input: algebra,
		}
	}

	return algebra, nil
}
```

### 3.5 Helper Functions

**separateFilterPatterns()**

```
Input: []GraphPattern (mixed BGP, Optional, Union, Filter)
Output: ([]FilterPattern, []GraphPattern)
  - Returns filters separately
  - Preserves order of non-filter patterns
```

**extractTriplesFromPatterns()**

```
Input: []GraphPattern
Output: []*Triple
  - Flattens all BGPs into single triple list
  - Used for OPTIONAL pattern conversion
```

**extractGroupByVars()**

```
Input: []*Aggregate
Output: []string
  - Extracts implicit GROUP BY variables
  - In SPARQL: non-aggregated variables in SELECT become GROUP BY
```

### 3.6 Implementation Checklist: AST→Algebra Converter

**Step 1: Implement Core Function**

- [ ] ASTToAlgebra() skeleton
- [ ] Input validation
- [ ] Pattern separation

**Step 2: Implement Conversion Helpers**

- [ ] convertPatterns() for recursion
- [ ] Pattern type matching (BGP, Optional, Union, Filter)
- [ ] separateFilterPatterns()
- [ ] extractTriplesFromPatterns()

**Step 3: Implement Application Functions**

- [ ] applyFilters()
- [ ] applyProjection()
- [ ] applyDistinct()
- [ ] applyOrderBy()
- [ ] applyAggregates()
- [ ] applyLimitOffset()

**Step 4: Error Handling**

- [ ] Validate no undefined variables
- [ ] Validate pattern structure
- [ ] Clear error messages

**Step 5: Testing (30+ test cases)**

- [ ] Simple SELECT with WHERE
- [ ] SELECT with FILTER
- [ ] SELECT with OPTIONAL
- [ ] SELECT with UNION
- [ ] SELECT with aggregates (COUNT, SUM, etc.)
- [ ] SELECT with DISTINCT
- [ ] SELECT with ORDER BY
- [ ] SELECT with LIMIT/OFFSET
- [ ] Combinations of above
- [ ] Complex nested patterns

**Example Test Case**:

```go
func TestASTToAlgebraSimpleSelect(t *testing.T) {
	ast := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?person", "?name"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?person", Predicate: "rdf:type", Object: "Person"},
					{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
				},
			},
		},
		Limit: 10,
	}

	algebra, err := ASTToAlgebra(ast)
	require.NoError(t, err)

	// Verify top-level is Limit
	limit, ok := algebra.(*Limit)
	require.True(t, ok)
	require.Equal(t, 10, limit.Count)

	// Verify next is Project
	project, ok := limit.Input.(*Project)
	require.True(t, ok)
	require.Equal(t, []string{"?person", "?name"}, project.Vars)

	// Verify input is Join
	join, ok := project.Input.(*Join)
	require.True(t, ok)
}
```

---

## Part 4: Algebra Validator (2-3 Days)

### 4.1 Design Goals

**Purpose**: Ensure algebra expressions are semantically valid before execution.

**Key Validations**:

1. **Variable Consistency**: All referenced variables are defined somewhere
2. **No Circular Dependencies**: Variable definitions don't form cycles
3. **Filter Expressions**: Only reference variables in scope
4. **Aggregate Functions**: Valid operations with valid inputs
5. **Projection**: Only projects defined variables

### 4.2 Validator Type and Methods

**File**: `sparql/algebra.go`

```go
type AlgebraValidator struct {
	// Tracking for validation
	definedVars map[string]bool    // Variables defined at current scope
	errors      []string           // Collected errors
	visitedExprs map[AlgebraExpr]bool  // For cycle detection
}

// Validate validates an algebra expression tree
// Returns error if any validation fails; nil if valid
func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
	// Implementation details below
}
```

### 4.3 Validation Logic - Variable Scope

**Key Concept**: Variables flow through algebra tree:

- BGP defines variables (in subjects/objects)
- Join combines definitions from left and right
- Filter references variables from input
- Project restricts which variables are output

**Algorithm**:

```
1. Traverse algebra tree using visitor pattern
2. Track which variables are defined at each point
3. When encountering reference (Filter, Project, etc.):
   - Check variable is in defined set
   - Report error if undefined
4. Return all errors at end
```

**Variable Scope Tracking**:

```
BGP(?s rdf:type ?c) defines: {?s, ?c}

Join(
  BGP(?s rdf:type Person),     defines: {?s}
  BGP(?s foaf:name ?name)      defines: {?s, ?name}
)
Combined defines: {?s, ?name}

Filter(?name = "Alice", Join(...))
  - Checks: is ?name in {?s, ?name}? YES ✓

Filter(?age > 18, Filter(?name = "Alice", Join(...)))
  - Checks: is ?age in {?s, ?name}? NO ✗ ERROR
```

### 4.4 Variable Extraction Helper

```go
// extractDefinedVariables(expr AlgebraExpr) → []string
// Returns all variables defined by this expression
func extractDefinedVariables(expr AlgebraExpr) []string {
	switch e := expr.(type) {
	case *BGP:
		// Variables in triple subjects and objects
		vars := make(map[string]bool)
		for _, triple := range e.Triples {
			if strings.HasPrefix(triple.Subject, "?") {
				vars[triple.Subject] = true
			}
			if triple.ObjectIsVar {
				vars[triple.Object] = true
			}
		}
		// Convert to slice
		...

	case *Join:
		// Combine variables from both sides
		left := extractDefinedVariables(e.Left)
		right := extractDefinedVariables(e.Right)
		return append(left, right...)

	case *Filter:
		// Filter doesn't define vars, passes through
		return extractDefinedVariables(e.Input)

	// ... etc for other operators
	}
}
```

### 4.5 Expression Variable Reference Extraction

```go
// extractReferencedVariables(expr string) → []string
// Parse a SPARQL expression to find variable references
// Examples:
//   "?name = 'Alice'" → ["?name"]
//   "?age > 18 && ?type = 'Person'" → ["?age", "?type"]
func extractReferencedVariables(expr string) []string {
	// Use regex to find all ?varname patterns
	re := regexp.MustCompile(`\?(\w+)`)
	matches := re.FindAllString(expr, -1)

	// Deduplicate
	varMap := make(map[string]bool)
	for _, m := range matches {
		varMap[m] = true
	}

	vars := make([]string, 0, len(varMap))
	for v := range varMap {
		vars = append(vars, v)
	}
	return vars
}
```

### 4.6 Full Validation Algorithm

```go
func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
	v.definedVars = make(map[string]bool)
	v.errors = make([]string, 0)

	// Traverse and validate
	v.validateExpr(expr)

	// Return all errors at once
	if len(v.errors) > 0 {
		return fmt.Errorf("algebra validation failed:\n  " +
			strings.Join(v.errors, "\n  "))
	}
	return nil
}

func (v *AlgebraValidator) validateExpr(expr AlgebraExpr) {
	switch e := expr.(type) {
	case *BGP:
		v.validateBGP(e)
	case *Join:
		v.validateJoin(e)
	case *Filter:
		v.validateFilter(e)
	case *LeftJoin:
		v.validateLeftJoin(e)
	case *Union:
		v.validateUnion(e)
	case *Project:
		v.validateProject(e)
	case *Aggregate:
		v.validateAggregate(e)
	case *Bind:
		v.validateBind(e)
	case *Distinct:
		v.validateDistinct(e)
	case *OrderBy:
		v.validateOrderBy(e)
	case *Limit:
		v.validateLimit(e)
	}
}

func (v *AlgebraValidator) validateBGP(bgp *BGP) {
	for _, triple := range bgp.Triples {
		if strings.HasPrefix(triple.Subject, "?") {
			v.definedVars[triple.Subject] = true
		}
		if triple.ObjectIsVar {
			v.definedVars[triple.Object] = true
		}
	}
}

func (v *AlgebraValidator) validateFilter(f *Filter) {
	// First validate input (defines variables)
	v.validateExpr(f.Input)

	// Then check filter references only defined vars
	refVars := extractReferencedVariables(f.Expr)
	for _, varRef := range refVars {
		if !v.definedVars[varRef] {
			v.errors = append(v.errors,
				fmt.Sprintf("Filter references undefined variable: %s", varRef))
		}
	}
}

func (v *AlgebraValidator) validateProject(p *Project) {
	// First validate input
	v.validateExpr(p.Input)

	// Then check all projected vars are defined
	for _, projVar := range p.Vars {
		if !v.definedVars[projVar] {
			v.errors = append(v.errors,
				fmt.Sprintf("Project references undefined variable: %s", projVar))
		}
	}
}

func (v *AlgebraValidator) validateJoin(j *Join) {
	// Validate left
	leftDefined := make(map[string]bool)
	v.validateExpr(j.Left)
	for k, v := range v.definedVars {
		leftDefined[k] = v
	}

	// Validate right (with left's vars in scope)
	v.validateExpr(j.Right)

	// Merge definitions
	for k, v := range leftDefined {
		v.definedVars[k] = v
	}
}

// Similar patterns for LeftJoin, Union, Aggregate, Bind, etc.
```

### 4.7 Error Messages

Good error messages should include:

- What variable is problematic
- Where it's used (which operator)
- What variables are in scope

**Examples**:

```
"Filter references undefined variable: ?age (defined: ?person, ?name)"
"Project references undefined variable: ?email (defined: ?person)"
"Aggregate operation COUNT requires numeric expression, got: ?name"
```

### 4.8 Implementation Checklist: Validator

**Step 1: Core Validation Type**

- [ ] AlgebraValidator struct
- [ ] Validate() method
- [ ] validateExpr() dispatcher

**Step 2: Validation Methods**

- [ ] validateBGP()
- [ ] validateJoin()
- [ ] validateFilter()
- [ ] validateLeftJoin()
- [ ] validateUnion()
- [ ] validateProject()
- [ ] validateAggregate()
- [ ] validateBind()
- [ ] validateDistinct()
- [ ] validateOrderBy()
- [ ] validateLimit()

**Step 3: Helper Functions**

- [ ] extractDefinedVariables()
- [ ] extractReferencedVariables()
- [ ] formatErrorMessage()

**Step 4: Testing**

- [ ] Valid expressions pass
- [ ] Undefined variable detected in Filter
- [ ] Undefined variable detected in Project
- [ ] Undefined variable detected in Aggregate
- [ ] Valid nested expressions pass
- [ ] Multiple errors collected and reported
- [ ] 15+ test cases

**Example Test**:

```go
func TestValidatorDetectsUndefinedVariable(t *testing.T) {
	algebra := &Filter{
		Expr: "?age > 18",
		Input: &BGP{
			Triples: []*Triple{
				{Subject: "?s", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			},
		},
	}

	validator := &AlgebraValidator{}
	err := validator.Validate(algebra)

	require.Error(t, err)
	require.Contains(t, err.Error(), "?age")
	require.Contains(t, err.Error(), "undefined")
}
```

---

## Part 5: Execution Context (1-2 Days)

### 5.1 Purpose and Design

**File**: `sparql/context.go`

**Purpose**: Carry metadata and configuration through the SPARQL query pipeline.

**Key Information to Carry**:

- Query context (cancellation, timeouts)
- Schema information (for Phase 3)
- Authentication information (for Phase 4)
- Translation options

### 5.2 SPARQLExecutionContext Type

```go
package sparql

import (
	"context"
	"github.com/dgraph-io/dgraph/v25/graphql/schema"
)

// SPARQLExecutionContext carries execution context through the SPARQL pipeline
type SPARQLExecutionContext struct {
	// Ctx is the request context (cancellation, timeout)
	Ctx context.Context

	// Schema is the GraphQL schema (for Phase 3: type validation)
	Schema *schema.Schema

	// Authentication information (for Phase 4: auth rule application)
	UserID    string   // Current user ID
	GroupIDs  []string // User's group memberships
	Namespace uint64   // Tenant/namespace for multi-tenancy

	// Translation options
	Options TranslateOptions

	// Query metadata
	QueryString string   // Original SPARQL query
	QueryHash   string   // For caching (optional)
}

// NewSPARQLExecutionContext creates a new context with defaults
func NewSPARQLExecutionContext(ctx context.Context) *SPARQLExecutionContext {
	return &SPARQLExecutionContext{
		Ctx:      ctx,
		GroupIDs: make([]string, 0),
	}
}

// WithSchema sets the schema for this context
func (c *SPARQLExecutionContext) WithSchema(s *schema.Schema) *SPARQLExecutionContext {
	c.Schema = s
	return c
}

// WithAuth sets authentication information
func (c *SPARQLExecutionContext) WithAuth(userID string, groupIDs []string) *SPARQLExecutionContext {
	c.UserID = userID
	c.GroupIDs = groupIDs
	return c
}

// WithOptions sets translation options
func (c *SPARQLExecutionContext) WithOptions(opts TranslateOptions) *SPARQLExecutionContext {
	c.Options = opts
	return c
}
```

**Design Notes**:

- Fluent API (builder pattern) for easy configuration
- Carries context through each pipeline stage
- Optional fields (can be nil if not needed)
- Extensible for future needs

### 5.3 Implementation Checklist: Context

- [ ] SPARQLExecutionContext struct definition
- [ ] NewSPARQLExecutionContext() factory
- [ ] WithSchema() fluent method
- [ ] WithAuth() fluent method
- [ ] WithOptions() fluent method
- [ ] 5+ test cases for context creation and configuration

---

## Part 6: Comprehensive Testing Strategy (2-3 Days)

### 6.1 Test File: algebra_test.go

**File Location**: `sparql/algebra_test.go`

**Test Organization**:

```go
package sparql

import (
	"testing"
	"github.com/stretchr/testify/require"
)

// Test suites:
// 1. Type Construction Tests (15 tests)
// 2. Visitor Pattern Tests (10 tests)
// 3. AST→Algebra Conversion Tests (50+ tests)
// 4. Validator Tests (15+ tests)
// 5. Context Tests (5 tests)
```

### 6.2 Test Coverage Matrix

| Area                      | Count    | Examples                                      |
| ------------------------- | -------- | --------------------------------------------- |
| Algebra Type Creation     | 15       | BGP, Join, Filter, Union, etc.                |
| Visitor Traversal         | 10       | Print simple, print nested, custom visitor    |
| Simple Conversions        | 15       | Single pattern, with filters, with projection |
| Complex Conversions       | 25       | OPTIONAL, UNION, aggregates, combinations     |
| Validator - Valid Cases   | 8        | Complex nested expressions, all operators     |
| Validator - Invalid Cases | 10       | Undefined vars, type mismatches, cycles       |
| Context                   | 5        | Creation, builder pattern, configurations     |
| **TOTAL**                 | **100+** | **Comprehensive coverage**                    |

### 6.3 Example Test Categories

**Category 1: Type Construction**

```go
func TestBGPConstruction(t *testing.T) { ... }
func TestJoinConstruction(t *testing.T) { ... }
func TestFilterConstruction(t *testing.T) { ... }
func TestProjectionConstruction(t *testing.T) { ... }
// ... etc
```

**Category 2: Visitor Pattern**

```go
func TestPrinterSimpleExpr(t *testing.T) { ... }
func TestPrinterNestedExpr(t *testing.T) { ... }
func TestPrinterFormatting(t *testing.T) { ... }
func TestCustomVisitor(t *testing.T) { ... }
```

**Category 3: AST to Algebra Conversion**

```go
func TestConvertSimpleSelect(t *testing.T) { ... }
func TestConvertSelectWithFilter(t *testing.T) { ... }
func TestConvertSelectWithOptional(t *testing.T) { ... }
func TestConvertSelectWithUnion(t *testing.T) { ... }
func TestConvertSelectWithAggregate(t *testing.T) { ... }
func TestConvertSelectWithDistinct(t *testing.T) { ... }
func TestConvertSelectWithOrderBy(t *testing.T) { ... }
func TestConvertSelectWithLimit(t *testing.T) { ... }
func TestConvertComplexQuery(t *testing.T) { ... }
// ... and many more combinations
```

**Category 4: Validator**

```go
func TestValidatorValidExpression(t *testing.T) { ... }
func TestValidatorDetectsUndefinedVariable(t *testing.T) { ... }
func TestValidatorDetectsInvalidProjection(t *testing.T) { ... }
func TestValidatorDetectsInvalidFilter(t *testing.T) { ... }
func TestValidatorComplexNesting(t *testing.T) { ... }
// ... and more
```

**Category 5: Context**

```go
func TestContextCreation(t *testing.T) { ... }
func TestContextBuilderPattern(t *testing.T) { ... }
func TestContextCancel(t *testing.T) { ... }
```

### 6.4 Test Quality Guidelines

**Each Test Should**:

1. Have clear name describing what's tested
2. Set up test input explicitly (no magic)
3. Call function under test
4. Assert expected output/behavior
5. Include error case (if applicable)

**Example Test Template**:

```go
func TestConvertSelectWithFilter(t *testing.T) {
	// Arrange: Create input AST
	ast := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?person", "?name"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?person", Predicate: "rdf:type", Object: "Person"},
					{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
				},
			},
			&FilterPattern{
				Expression: "?name = 'Alice'",
			},
		},
	}

	// Act: Convert to algebra
	algebra, err := ASTToAlgebra(ast)

	// Assert
	require.NoError(t, err)
	require.NotNil(t, algebra)

	// Verify structure
	limit, ok := algebra.(*Limit)
	require.True(t, ok, "Expected Limit at top level")

	project, ok := limit.Input.(*Project)
	require.True(t, ok, "Expected Project")

	filter, ok := project.Input.(*Filter)
	require.True(t, ok, "Expected Filter")
	require.Equal(t, "?name = 'Alice'", filter.Expr)

	join, ok := filter.Input.(*Join)
	require.True(t, ok, "Expected Join")
}
```

### 6.5 Running Tests

**Command**:

```bash
cd /Users/nfeldman/repos/dgraph
go test ./sparql -v
```

**Expected Output**:

```
=== RUN   TestBGPConstruction
--- PASS: TestBGPConstruction (0.00s)
=== RUN   TestJoinConstruction
--- PASS: TestJoinConstruction (0.00s)
...
=== RUN   TestConvertComplexQuery
--- PASS: TestConvertComplexQuery (0.01s)
...
PASS
ok      github.com/dgraph-io/dgraph/v25/sparql  2.345s
```

---

## Integration Points with Existing Code

### Phase 1 → Existing SPARQL Code

**No breaking changes to existing code**. Phase 1 is purely additive:

1. **ast.go**: Already has SPARQLQueryImpl, BGP, Triple, etc.
   - Algebra types reference these, don't modify
2. **translator_extended.go**: Current AST→DQL translation
   - Will add ASTToAlgebra() call in future phases
   - Currently unused; no impact

3. **antlr_adapter.go**: Parses SPARQL → AST
   - No changes needed
   - Algebra compiler accepts AST as input

4. **Existing tests**: translate_test.go, translator_extended_test.go
   - Run `go test ./sparql` to verify no regressions
   - All should continue to pass

**Integration Points for Future Phases**:

- Phase 2: Add optimizer that takes algebra tree and applies rewrites
- Phase 3: Modify translator_extended.go to use ASTToAlgebra() + Algebra→DQL
- Phase 4: Add auth rules applied to algebra before DQL compilation

---

## Key Design Decisions & Reasoning

### Decision 1: Expression as String, Not Structured AST

**Issue**: Should Filter.Expr be a string or a structured expression AST?

**Decision**: Use string for now (`Expr string`)

**Reasoning**:

- Existing code uses strings for expressions
- Structured AST would be refactoring scope (Phase 2)
- Less migration work for Phase 1
- Validator can parse strings with regex if needed
- Future phases can introduce structured expressions

**Trade-off**: Slightly less type-safe, but faster to implement

---

### Decision 2: Variable Naming Convention

**Issue**: Should variables keep ? prefix (?person) or use bare names (person)?

**Decision**: Keep ? prefix in algebra, strip in DQL generation

**Reasoning**:

- Matches SPARQL convention
- Clearer intent in debugging
- Less error-prone during conversion
- DQL conversion removes prefix when compiling

**Example**:

```
SPARQL: SELECT ?person WHERE { ?person rdf:type Person }
Algebra: Project(vars=[?person], ...)
DQL: Generated without ? prefix
```

---

### Decision 3: Where to Apply Operators

**Issue**: What order should operators wrap each other?

**Decision**: Follow SPARQL semantics execution order:

```
FROM (implicit)
→ WHERE (patterns, filters)
→ GROUP BY (aggregates)
→ HAVING (aggregate filters)
→ SELECT (projection)
→ DISTINCT
→ ORDER BY
→ LIMIT/OFFSET
```

**Corresponding Algebra Nesting** (outermost to innermost):

```
Limit
  OrderBy
    Distinct
      Project
        Aggregate (if present)
          Filter (WHERE)
            Join (patterns)
              BGP
```

---

### Decision 4: LeftJoin vs. OptionalJoin

**Issue**: How to represent OPTIONAL?

**Decision**: Use LeftJoin operator

**Reasoning**:

- W3C SPARQL Algebra uses LeftJoin
- Clearly semantics: left side is mandatory, right side optional
- Better optimization opportunities (can detect optional branches)
- Matches existing literature on SPARQL optimization

---

### Decision 5: Multiple Aggregates

**Issue**: How to handle multiple aggregates (SELECT COUNT(\*) as cnt, SUM(?amount) as total)?

**Decision**: Single Aggregate operator; extend in Phase 2 if needed

**Reasoning**:

- Simplifies Phase 1
- SPARQL typically groups these
- Can refactor to list-based in Phase 2 if needed
- Validation logic clearer

---

### Decision 6: Validation Approach

**Issue**: Early validation or lazy evaluation?

**Decision**: Early validation (fail fast)

**Reasoning**:

- Catches errors before execution
- Better error messages with context
- Enables safe optimization passes
- Follows best practices in compiler design

---

## Testing Strategy in Detail

### Test Phases

**Phase 1: Unit Tests (algebra_test.go)**

- Test individual algebra operators
- Test conversion logic
- Test visitor pattern
- Test validation rules
- ~100+ test cases
- Run: `go test ./sparql -v -run TestAlgebra`

**Phase 2: Integration Tests (algebra_test.go)**

- Test full query conversions
- Test complex nested patterns
- Test error handling
- Test edge cases
- ~20+ additional test cases

**Phase 3: Regression Tests**

- Run entire test suite
- Verify existing tests still pass
- Run: `go test ./sparql -v`
- Verify no performance regression

### Test Naming Convention

```
Test<Operator><Scenario>
Test<Function><Expected>
Test<Feature><Case>

Examples:
- TestBGPCreation
- TestJoinNesting
- TestFilterWithUndefinedVariable
- TestConvertSelectWithOptional
- TestValidatorDetectsCircularDep
```

---

## Build & Deployment Checklist

### Before Starting Implementation

- [ ] Create feature branch: `git checkout -b feature/sparql-algebra-foundation`
- [ ] Review this specification thoroughly
- [ ] Read reference code (ast.go, translator_extended.go, graphql/resolve/query_rewriter.go)

### During Implementation

- [ ] Write incrementally (type system → visitor → converter → validator)
- [ ] Test frequently (`go test ./sparql -v`)
- [ ] Keep commits small and focused
- [ ] Document design decisions in code

### Before Code Review

- [ ] All tests pass: `go test ./sparql -v`
- [ ] No regressions in existing SPARQL tests
- [ ] Code follows Dgraph conventions
- [ ] All functions have docstrings
- [ ] Examples in tests cover all operators

### Code Review Checklist

- [ ] Design decisions documented
- [ ] All error cases handled
- [ ] Variable naming consistent
- [ ] No panics (proper error returns)
- [ ] Performance acceptable (algebra generation <10ms)

---

## Estimated Timeline

| Phase                    | Duration        | Deliverables                                       |
| ------------------------ | --------------- | -------------------------------------------------- |
| 1. Type System           | 3-4 days        | All operators defined, String() methods, 15+ tests |
| 2. Visitor Pattern       | 2-3 days        | Interface, printer, Accept() methods, 10+ tests    |
| 3. AST→Algebra           | 4-5 days        | Full converter, 30+ conversion tests               |
| 4. Validator             | 2-3 days        | Validation logic, 15+ validator tests              |
| 5. Context               | 1-2 days        | Context type, 5+ context tests                     |
| 6. Integration & Testing | 2-3 days        | Full test suite, edge cases, documentation         |
| **TOTAL**                | **~10-15 days** | **All deliverables complete**                      |

---

## Blockers & Clarifications Needed

### No Current Blockers

Phase 1 is self-contained and can proceed independently:

- ✅ AST types already defined (ast.go)
- ✅ Expression representation understood (string-based)
- ✅ No external dependencies on schema (added in Phase 3)
- ✅ No auth system needs (added in Phase 4)

### Potential Questions

**Q1**: Should algebra support CONSTRUCT/DESCRIBE queries?

- **A**: No, Phase 1 focuses on SELECT/ASK. CONSTRUCT in Phase 2.

**Q2**: How to represent FILTER expressions in algebra long-term?

- **A**: Strings for Phase 1; refactor to structured AST in Phase 2 if needed.

**Q3**: Should we optimize immediately after conversion?

- **A**: No, Phase 1 = conversion only. Phase 2 = optimization passes.

**Q4**: How do we handle namespace/named graphs?

- **A**: Phase 1 ignores them (preserved in Triple.Graph). Phase 3 integrates with schema.

---

## References & Resources

### SPARQL Specification

- W3C SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
- SPARQL Query Language: https://www.w3.org/TR/sparql11-query/

### Design Patterns

- Visitor Pattern: Gang of Four, Chapter 12
- AST-based compilation: "Compilers: Principles, Techniques, and Tools"

### Code References in Dgraph

- Similar translator: `graphql/resolve/query_rewriter.go`
- AST handling: `sparql/translator_extended.go`
- Filter trees: `dql/parser.go` (FilterTree, GraphQuery)
- Expression parsing: `sparql/filter_extended.go`

### Go Testing Best Practices

- Table-driven tests: https://github.com/golang/go/wiki/TableDrivenTests
- Test fixtures and helpers: standard Go testing patterns

---

## Success Criteria

Phase 1 is complete when:

1. ✅ All algebra operators implemented with String() methods
2. ✅ Visitor pattern infrastructure complete
3. ✅ ASTToAlgebra() function converts all query types
4. ✅ AlgebraValidator checks semantic correctness
5. ✅ SPARQLExecutionContext carries metadata
6. ✅ 100+ test cases, all passing
7. ✅ Existing SPARQL tests still pass (no regressions)
8. ✅ Code review approved
9. ✅ Documentation complete
10. ✅ Branch mergeable to main

---

## Next Steps

After Phase 1 completion:

1. **Merge and Review**: Get code reviewed and merged
2. **Plan Phase 2**: Design optimization passes (filter pushdown, etc.)
3. **Benchmark Baseline**: Measure current query performance
4. **Phase 2 Kickoff**: Implement algebraic rewrites and optimizations

---

**End of Implementation Plan**
