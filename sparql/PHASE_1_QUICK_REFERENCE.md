# Phase 1 Implementation Summary & Quick Reference

**Status**: Ready for Implementation  
**Complexity**: Medium (4 sequential steps, ~1300 LOC total)  
**Timeline**: 10-15 days  
**Team Size**: 1-2 developers

---

## Quick Start

### What You're Building

A **SPARQL Algebra intermediate representation** that transforms SPARQL queries from syntactic AST
form into semantic algebra form, enabling optimization and schema-aware processing.

### Why It Matters

- **Current path**: SPARQL → AST → DQL String → (re-parse) → GraphQuery (slow, no optimization)
- **New path**: SPARQL → AST → **Algebra** → (optimize) → DQL GraphQuery (fast, schema-aware)
- **Foundation** for authorization rules, ontology reasoning, and query optimization

### High-Level Algorithm

```
SPARQLQueryImpl (from parser)
    ↓ ASTToAlgebra()
Algebra Tree (operators like BGP, Join, Filter, etc.)
    ↓ (Future: Optimizer applies rewrites)
DQL GraphQuery (with type/auth information)
    ↓ (Execute)
Results
```

---

## Four Implementation Steps (Sequential)

### Step 1: Algebra Type System (3-4 days)

**What**: Define 11 algebra operator types  
**File**: `sparql/algebra.go` (section 2)  
**Key Types**: BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, Bind, Distinct, OrderBy,
Limit  
**Deliverables**: Types + String() methods, ~200 LOC  
**Tests**: 15+ type construction tests

### Step 2: Visitor Pattern (2-3 days)

**What**: Implement visitor pattern for traversal  
**Files**: `sparql/algebra_visitor.go`  
**Key Types**: AlgebraVisitor interface, AlgebraPrinter implementation  
**Deliverables**: Visitor infrastructure + printer, ~150 LOC  
**Tests**: 10+ visitor/printer tests

### Step 3: AST→Algebra Converter (4-5 days)

**What**: Implement ASTToAlgebra() function  
**File**: `sparql/algebra.go` (section 3)  
**Key Functions**: convertPatterns(), convertSinglePattern(), applyFilters(), etc.  
**Deliverables**: Full conversion logic, ~250 LOC  
**Tests**: 50+ conversion tests (simple→complex queries)

### Step 4: Algebra Validator (2-3 days)

**What**: Ensure semantic correctness  
**File**: `sparql/algebra.go` (section 4)  
**Key Type**: AlgebraValidator  
**Deliverables**: Validation logic, ~150 LOC  
**Tests**: 15+ validator tests

---

## Minimal MVP vs Full Implementation

### Minimal MVP (5-7 days, ~800 LOC)

- ✅ Core algebra types (no Aggregate/Bind initially)
- ✅ Basic visitor (just printer)
- ✅ Simple converter (SELECT with WHERE, FILTER, OPTIONAL, UNION only)
- ✅ Basic validator (variable consistency only)
- ✅ 50+ tests

### Full Phase 1 (10-15 days, ~1300 LOC)

- ✅ All 11 algebra operators
- ✅ Full visitor infrastructure
- ✅ Complete converter (all SPARQL features)
- ✅ Comprehensive validator (all checks)
- ✅ 100+ tests
- ✅ Execution context carrier

**Recommendation**: Target Full Phase 1 (more useful foundation for Phase 2)

---

## Key Design Decisions

| Decision                 | Choice                      | Why                                 |
| ------------------------ | --------------------------- | ----------------------------------- |
| Expression format        | String (not structured AST) | Faster Phase 1; refactor Phase 2    |
| Variable naming          | Keep ? prefix               | Clearer, matches SPARQL convention  |
| Operator ordering        | W3C SPARQL execution order  | Correct semantics, matches spec     |
| LeftJoin vs OptionalJoin | LeftJoin                    | W3C standard, clear semantics       |
| Multiple aggregates      | Single operator per query   | Simpler Phase 1; extend Phase 2     |
| Validation timing        | Early (fail-fast)           | Better errors, enables optimization |
| Error handling           | Return all errors together  | User sees complete picture          |

---

## Testing Strategy

### Test Coverage Matrix

| Category             | Count    | Examples                                  |
| -------------------- | -------- | ----------------------------------------- |
| Type Construction    | 15       | BGP, Join, Filter, Union, Project, etc.   |
| Visitor Traversal    | 10       | Print simple, nested, formatting          |
| Conversion - Simple  | 20       | SELECT, SELECT+WHERE, SELECT+FILTER       |
| Conversion - Complex | 30       | OPTIONAL, UNION, aggregates, combinations |
| Validator - Valid    | 10       | All operators, nested, complex scopes     |
| Validator - Invalid  | 10       | Undefined vars, type errors, cycles       |
| Context              | 5        | Creation, builder pattern, auth           |
| **TOTAL**            | **100+** | **Comprehensive**                         |

### Running Tests

```bash
# Run all SPARQL tests
go test ./sparql -v

# Run only algebra tests
go test ./sparql -v -run TestAlgebra

# Run with verbose output
go test ./sparql -v -run Test -count=1

# Run with coverage
go test ./sparql -v -cover
```

---

## Code Structure Overview

```
sparql/algebra.go (500-600 LOC)
├── AlgebraExpr interface
├── Algebra Operators (11 types)
│   ├── BGP, Join, Filter, LeftJoin, Union
│   ├── Project, Aggregate, Bind, Distinct
│   └── OrderBy, Limit
├── ASTToAlgebra() function + helpers
│   ├── convertPatterns()
│   ├── convertSinglePattern()
│   ├── applyFilters(), applyOrderBy(), etc.
│   └── Extract* helper functions
└── AlgebraValidator + helper functions
    ├── validateExpr() dispatcher
    ├── validateXxx() per operator
    └── extractReferencedVariables()

sparql/algebra_visitor.go (100-150 LOC)
├── AlgebraVisitor interface
└── AlgebraPrinter implementation
    ├── Print() method
    ├── writeLine() helper
    └── VisitXxx() methods (11 per operator)

sparql/context.go (100 LOC)
└── SPARQLExecutionContext
    ├── Fluent builder methods (WithSchema, WithAuth, etc.)
    └── Helper methods (IsAuthenticated, IsMember)

sparql/algebra_test.go (400-500 LOC)
├── Type construction tests (15)
├── Visitor pattern tests (10)
├── AST→Algebra conversion tests (50+)
├── Validator tests (15+)
├── Context tests (5+)
└── Test helper functions
```

---

## Integration Points

### With Existing Code

- **No changes** to ast.go, translator_extended.go, antlr_adapter.go
- Algebra types **reference** existing Triple, BGP, Expression types
- Fully backward compatible

### With Future Phases

- **Phase 2**: Optimizer applies rewrites to algebra tree (visitor-based)
- **Phase 3**: Schema integration (executor uses context.Schema)
- **Phase 4**: Auth rules applied to algebra (visitor-based transformation)

---

## Common Pitfalls & Solutions

### Pitfall 1: Expression Parsing

**Problem**: Filtering/ordering expressions are strings; how to extract variables?  
**Solution**: Use regex: `\?(\w+)` to find all ?variables  
**Code**: `extractReferencedVariables()` function

### Pitfall 2: Variable Scope Tracking

**Problem**: Variables flow through tree; undefined vars hard to detect  
**Solution**: Visitor pattern with scope accumulation  
**Code**: Validator tracks `definedVars` map as it traverses

### Pitfall 3: Operator Nesting Order

**Problem**: Apply operators in wrong order → wrong semantics  
**Solution**: Follow W3C SPARQL execution order  
**Code**: AST→Algebra does steps 1-10 in correct order (see Section 3.3)

### Pitfall 4: Filter Patterns Mixed With Others

**Problem**: Filters can appear anywhere in pattern list  
**Solution**: Separate filters first, apply at end  
**Code**: `separateFilterPatterns()` helper

### Pitfall 5: Multiple Test Assertions

**Problem**: Test fails on first assertion; hard to see full picture  
**Solution**: Use require (assert immediately) or collect all errors  
**Code**: Tests assert each structural piece separately

---

## File Checklist

### Before Starting

- [ ] Read PHASE_1_GETTING_STARTED.md
- [ ] Read ARCHITECTURE_SPEC.md
- [ ] Read this document (PHASE_1_IMPLEMENTATION_PLAN.md)
- [ ] Read CODE_STRUCTURE_GUIDE.md
- [ ] Review ast.go (understand AST types)
- [ ] Review translator_extended.go (understand current flow)
- [ ] Review graphql/resolve/query_rewriter.go (understand visitor pattern in Dgraph)

### During Implementation

- [ ] Create algebra.go with all types
- [ ] Create algebra_visitor.go with interface + printer
- [ ] Create context.go with execution context
- [ ] Create algebra_test.go with comprehensive tests
- [ ] Run tests frequently: `go test ./sparql -v`
- [ ] Commit incrementally with clear messages

### Code Review Prep

- [ ] All tests pass
- [ ] No regressions in existing tests
- [ ] All functions have docstrings
- [ ] Error cases documented
- [ ] Design decisions in code comments
- [ ] Examples in tests cover all operators

---

## Example: Converting a Real Query

### Input SPARQL

```sparql
SELECT ?person ?name WHERE {
  ?person rdf:type Person .
  ?person foaf:name ?name .
  FILTER (?name = "Alice")
}
LIMIT 10
```

### AST (From Parser - Already Have)

```go
SPARQLQueryImpl{
  Qtype: "SELECT",
  Projs: ["?person", "?name"],
  Patterns: []GraphPattern{
    BGP{Triples: [
      Triple{Subject: "?person", Predicate: "rdf:type", Object: "Person"},
      Triple{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
    ]},
    FilterPattern{Expression: "?name = \"Alice\""},
  },
  Limit: 10,
}
```

### Algebra (What We Build)

```
Limit(count=10,
  Project(vars=[?person, ?name],
    Filter(expr="?name = \"Alice\"",
      Join(
        BGP([?person:rdf:type:<Person>]),
        BGP([?person:foaf:name:?name])
      )
    )
  )
)
```

### Code Execution Path

```
ASTToAlgebra(ast)
  ↓ extractPatterns() → [BGP, FilterPattern]
  ↓ separateFilterPatterns() → filters=[FilterPattern], others=[BGP]
  ↓ convertPatterns([BGP]) → Join(BGP, BGP)
  ↓ applyFilters(join, [FilterPattern]) → Filter(join)
  ↓ applyLimitOffset(filter, 10, 0) → Limit(Filter(Join(...)))
  ↓ applyProjection(limit, [?person, ?name]) → Project(Limit(...))
  ↓ return Project(...)
```

---

## Validation in Action

### Valid Query

```go
BGP{?s rdf:type Person} defines {?s}
Join(BGP1, BGP2) defines {?s, ?name} if both triples reference these
Filter("?name = 'Alice'", Join) ✓ OK (?name is defined)
Project([?person], Filter) ✓ OK (?person is defined, error: no such var!)
```

**Fix**: Variable names must match exactly - "?person" vs "?s"

### Invalid Query

```go
Filter("?age > 18", Join) ✗ ERROR: ?age undefined
  Defined vars in Join: {?s, ?name}
  Referenced in Filter: {?age}
  Error: "Filter references undefined variable ?age"
```

---

## Performance Expectations

### Compilation Time

- Simple query (1-3 patterns): <1ms
- Complex query (10+ patterns, filters): <5ms
- Target: algebra generation <10ms for typical queries

### Memory Usage

- BGP with 100 triples: ~10KB (small)
- Deeply nested tree (depth 10): ~50KB (very small)
- Overall Phase 1 very efficient

### Testing Time

- ~100 test cases: 1-2 seconds to run
- No performance regressions expected

---

## Rollout Strategy

### Phase 1A: Core (Days 1-5)

1. Create algebra.go with types
2. Create algebra_visitor.go with printer
3. Write first 20 tests
4. Verify types and visitor work
5. Get early code review feedback

### Phase 1B: Conversion (Days 6-9)

1. Implement ASTToAlgebra()
2. Write 50+ conversion tests
3. Test against real SPARQL queries
4. Fix bugs, edge cases
5. Get conversion review

### Phase 1C: Validation (Days 10-12)

1. Implement AlgebraValidator
2. Write 15+ validator tests
3. Implement context.go
4. Write remaining tests
5. Final code review

### Phase 1D: Integration (Days 13-15)

1. Run full test suite
2. Check for regressions
3. Update documentation
4. Prepare for merge
5. Address final review comments

---

## Success Criteria Checklist

### Functionality ✓

- [ ] All 11 algebra operators implemented and working
- [ ] ASTToAlgebra() converts all SPARQL patterns
- [ ] Visitor pattern traverses all operators
- [ ] Validator detects semantic errors
- [ ] Context carries metadata through pipeline

### Testing ✓

- [ ] 100+ test cases written
- [ ] All tests passing
- [ ] Code coverage >90%
- [ ] No flaky tests
- [ ] Edge cases covered

### Quality ✓

- [ ] No regressions in existing SPARQL tests
- [ ] All functions documented
- [ ] Error messages clear and helpful
- [ ] Code follows Dgraph conventions
- [ ] No panics or unhandled errors

### Performance ✓

- [ ] Algebra generation <10ms
- [ ] No memory leaks
- [ ] Efficient data structures
- [ ] Minimal allocations

### Documentation ✓

- [ ] Code comments explain intent
- [ ] Docstrings on all public functions
- [ ] Test cases show usage patterns
- [ ] Design decisions documented

---

## Links & References

### Documentation

- [PHASE_1_GETTING_STARTED.md](PHASE_1_GETTING_STARTED.md) - Requirements and examples
- [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - System design
- [PHASE_1_IMPLEMENTATION_PLAN.md](PHASE_1_IMPLEMENTATION_PLAN.md) - Detailed plan (THIS FILE)
- [PHASE_1_CODE_STRUCTURE.md](PHASE_1_CODE_STRUCTURE.md) - Code examples

### Standards & Specs

- [W3C SPARQL Algebra](https://www.w3.org/TR/sparql11-query/#sparqlAlgebra)
- [W3C SPARQL Query Language](https://www.w3.org/TR/sparql11-query/)

### Code References

- Existing SPARQL AST: `sparql/ast.go`
- Current translator: `sparql/translator_extended.go`
- GraphQL visitor pattern: `graphql/resolve/query_rewriter.go`
- DQL structures: `dql/parser.go`

---

## FAQ

**Q1: Should we implement all 11 operators or start with subset?**  
A: Start with all 11. They're simple and separate concerns. Faster to implement all than refactor
later.

**Q2: Can we reuse existing Expression types?**  
A: Yes! Expressions stay as strings for Phase 1. Refactor to structured AST in Phase 2 if needed.

**Q3: How do we handle CONSTRUCT/DESCRIBE queries?**  
A: Phase 1 focuses on SELECT/ASK. CONSTRUCT in Phase 2.

**Q4: What if pattern extraction fails?**  
A: Return error with context: "Pattern type X not supported at position Y"

**Q5: How to test Validator thoroughly?**  
A: Create both valid and invalid algebras; verify error messages.

**Q6: Can we parallelize test execution?**  
A: Yes, use `go test -parallel N`. ~100 tests should run in 1-2 seconds.

**Q7: What about variable case sensitivity?**  
A: SPARQL variables ARE case-sensitive. Preserve exact case.

**Q8: Should we optimize immediately?**  
A: No, Phase 1 = conversion only. Phase 2 = optimization passes.

---

## Next Steps After Phase 1

Once Phase 1 is merged and stable:

1. **Plan Phase 2**: Optimization passes (filter pushdown, join reordering)
2. **Benchmark**: Measure baseline query performance
3. **Design Phase 3**: Schema integration approach
4. **Study Phase 4**: GraphQL auth pattern for SPARQL

---

## Contact & Questions

For clarifications or design discussions:

- Review planning documents in `sparql/` directory
- Reference implementation examples in this guide
- Check code in `graphql/resolve/query_rewriter.go` for similar patterns

---

**This plan is ready for implementation. Start with Step 1 (Type System) and proceed sequentially.**
