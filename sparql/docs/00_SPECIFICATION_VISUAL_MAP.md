# Phase 1 SPARQL Algebra - Visual Specification Overview

## Document Map

```
┌─────────────────────────────────────────────────────────────────┐
│                    IMPLEMENTATION SPECIFICATION                  │
│                    SPARQL Algebra Phase 1                        │
│                   11,000+ words, 5 Documents                    │
└─────────────────────────────────────────────────────────────────┘

START HERE:
├─ 00_START_HERE.md (THIS FILE - Overview)
│  ├─ What you received
│  ├─ The 4-step task
│  ├─ How to get started
│  └─ Key takeaways
│
NAVIGATION & ORIENTATION:
├─ PHASE_1_INDEX.md (15 min)
│  ├─ Quick facts table
│  ├─ Critical path to implementation
│  ├─ Getting started roadmap
│  └─ Document locations
│
EXECUTIVE SUMMARY:
├─ PHASE_1_QUICK_REFERENCE.md (30 min) ⭐ BEST FOR EXECUTIVES
│  ├─ High-level algorithm
│  ├─ 4-step overview with durations
│  ├─ Key design decisions table
│  ├─ Testing strategy matrix
│  ├─ Real query example (SPARQL→Algebra)
│  ├─ Performance expectations
│  ├─ Rollout strategy
│  └─ FAQ
│
DETAILED SPECIFICATION:
├─ PHASE_1_IMPLEMENTATION_PLAN.md (2-3 hours) ⭐ BEST FOR DEVELOPERS
│  ├─ Part 1: Algebra Type System (3-4 days)
│  │  ├─ AlgebraExpr interface
│  │  ├─ 11 operator types (fully specified)
│  │  ├─ String() format for each
│  │  └─ Implementation checklist
│  │
│  ├─ Part 2: Visitor Pattern Infrastructure (2-3 days)
│  │  ├─ Design philosophy
│  │  ├─ AlgebraVisitor interface
│  │  ├─ AlgebraPrinter implementation
│  │  └─ Example tests
│  │
│  ├─ Part 3: AST to Algebra Converter (4-5 days)
│  │  ├─ 10-step conversion algorithm
│  │  ├─ ASTToAlgebra function signature
│  │  ├─ Pattern conversion logic
│  │  ├─ Helper functions
│  │  ├─ 30+ conversion test strategy
│  │  └─ Example test case
│  │
│  ├─ Part 4: Algebra Validator (2-3 days)
│  │  ├─ Variable scope tracking
│  │  ├─ Semantic validation logic
│  │  └─ Validator test strategy
│  │
│  ├─ Part 5: Execution Context (1-2 days)
│  │  └─ SPARQLExecutionContext definition
│  │
│  ├─ Part 6: Comprehensive Testing (2-3 days)
│  │  ├─ Test categories and counts
│  │  ├─ Coverage matrix
│  │  └─ Example test cases
│  │
│  └─ Supporting sections
│     ├─ Design decisions (6, all justified)
│     ├─ Success criteria
│     └─ Blockers & clarifications
│
CODE REFERENCE:
├─ PHASE_1_CODE_STRUCTURE.md (1 hour) ⭐ BEST WHILE CODING
│  ├─ algebra.go structure outline (500-600 LOC)
│  │  ├─ All 11 operators with implementations
│  │  ├─ ASTToAlgebra function + helpers
│  │  └─ AlgebraValidator
│  │
│  ├─ algebra_visitor.go structure (100-150 LOC)
│  │  ├─ AlgebraVisitor interface
│  │  └─ AlgebraPrinter class
│  │
│  ├─ context.go structure (100 LOC)
│  │  └─ SPARQLExecutionContext
│  │
│  ├─ algebra_test.go structure (400-500 LOC)
│  │  ├─ Type construction tests (15)
│  │  ├─ Visitor tests (10)
│  │  ├─ Conversion tests (50+)
│  │  └─ Validator tests (15+)
│  │
│  └─ Common patterns & conventions
│
WORKFLOW & TRACKING:
└─ DELIVERABLES_SUMMARY.md (15 min)
   ├─ What you received
   ├─ How to use documents
   ├─ Implementation timeline
   ├─ Success criteria checklist
   └─ Pre-implementation checklist
```

---

## Reading Paths by Role

### I'm an Executive (Need Quick Overview)

1. **This file** (00_START_HERE.md) - 5 min
2. PHASE_1_INDEX.md - 10 min
3. PHASE_1_QUICK_REFERENCE.md - 30 min
4. **Total: 45 minutes**

### I'm Implementing Phase 1

1. **This file** (00_START_HERE.md) - 5 min
2. PHASE_1_INDEX.md - 10 min
3. PHASE_1_QUICK_REFERENCE.md - 30 min
4. PHASE_1_IMPLEMENTATION_PLAN.md - 2 hours
5. Review existing code (ast.go, translator_extended.go) - 1 hour
6. Start implementing following PLAN + using CODE_STRUCTURE.md
7. **Total: 3.5 hours prep + 10-15 days implementation**

### I'm Reviewing Code

1. PHASE_1_IMPLEMENTATION_PLAN.md - 1 hour (focus on design decisions)
2. PHASE_1_CODE_STRUCTURE.md - 30 min (code patterns)
3. PHASE_1_QUICK_REFERENCE.md - 20 min (success criteria)
4. **Total: 1.5 hours**

### I'm Planning Phase 2

1. PHASE_1_QUICK_REFERENCE.md - 30 min
2. PHASE_1_IMPLEMENTATION_PLAN.md - 1 hour (focus on integration points)
3. PHASE_1_CODE_STRUCTURE.md - 30 min (understand algebra structure)
4. **Total: 2 hours**

---

## Quick Facts

```
┌─────────────────────────────────────────────────────────────────┐
│ IMPLEMENTATION TASK OVERVIEW                                    │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│ Timeline: 10-15 days (1 developer) or 6-8 days (2 developers)  │
│ Complexity: Moderate (well-specified, no ambiguities)          │
│ Total LOC: ~1,300 (highly specified patterns)                  │
│ Test Cases: 100+ (comprehensive coverage)                      │
│ Breaking Changes: NONE (fully backward compatible)             │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│ FILES TO CREATE (4)                                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│ sparql/algebra.go                    500-600 LOC               │
│   ├─ AlgebraExpr interface                                     │
│   ├─ 11 algebra operator types                                 │
│   ├─ ASTToAlgebra() converter                                  │
│   └─ AlgebraValidator                                          │
│                                                                 │
│ sparql/algebra_visitor.go            100-150 LOC               │
│   ├─ AlgebraVisitor interface                                  │
│   └─ AlgebraPrinter implementation                             │
│                                                                 │
│ sparql/context.go                    100 LOC                   │
│   └─ SPARQLExecutionContext struct                             │
│                                                                 │
│ sparql/algebra_test.go               400-500 LOC               │
│   └─ 100+ comprehensive test cases                             │
│                                                                 │
├─────────────────────────────────────────────────────────────────┤
│ IMPLEMENTATION STEPS (4 Sequential)                            │
├─────────────────────────────────────────────────────────────────┤
│                                                                 │
│ Step 1: Algebra Type System (3-4 days)                         │
│   └─ Define 11 operators with String() methods                │
│                                                                 │
│ Step 2: Visitor Pattern (2-3 days)                             │
│   └─ Implement AlgebraVisitor + AlgebraPrinter                 │
│                                                                 │
│ Step 3: AST→Algebra Converter (4-5 days)                       │
│   └─ Implement ASTToAlgebra() function (10 steps)              │
│                                                                 │
│ Step 4: Validator + Context (2-3 days)                         │
│   ├─ Implement AlgebraValidator                                │
│   └─ Implement SPARQLExecutionContext                          │
│                                                                 │
│ Step 5: Testing (2-3 days)                                     │
│   └─ Write 100+ test cases                                     │
│                                                                 │
└─────────────────────────────────────────────────────────────────┘
```

---

## Algebra Operators (All Specified)

```
┌────────────────────────────────────────────────────────────────┐
│ 11 ALGEBRA OPERATORS                                           │
├─────────────┬──────────────────────────────────────────────────┤
│ BGP         │ Basic graph pattern                              │
│ Join        │ Pattern combination                              │
│ Filter      │ WHERE constraints                                │
│ LeftJoin    │ OPTIONAL patterns                                │
│ Union       │ UNION alternatives                               │
│ Project     │ SELECT variables                                 │
│ Aggregate   │ GROUP BY / aggregates                            │
│ Bind        │ BIND expressions                                 │
│ Distinct    │ DISTINCT modifier                                │
│ OrderBy     │ ORDER BY clause                                  │
│ Limit       │ LIMIT/OFFSET                                     │
└─────────────┴──────────────────────────────────────────────────┘
```

---

## Algorithm at a Glance

```
Input:
  SPARQL: SELECT ?person ?name WHERE {
            ?person rdf:type Person .
            ?person foaf:name ?name .
            FILTER (?name = "Alice")
          } LIMIT 10

         ↓ (Parser)

AST:     SPARQLQueryImpl {
           Qtype: "SELECT",
           Projs: ["?person", "?name"],
           Patterns: [BGP, FilterPattern],
           Limit: 10
         }

         ↓ (ASTToAlgebra)

ALGEBRA: Limit(count=10,
           Project(vars=[?person, ?name],
             Filter(expr="?name = 'Alice'",
               Join(
                 BGP([?person:rdf:type:Person]),
                 BGP([?person:foaf:name:?name])
               )
             )
           )
         )

         ↓ (Future: Optimize)
         ↓ (Future: Compile to DQL)

Output: Results
```

---

## Design Decisions (All Made)

```
┌──────────────────────────────────────────────────────────────┐
│ KEY DESIGN DECISIONS                                         │
├──────────────────┬──────────────────────────────────────────┤
│ Expression Type  │ String (not structured AST)              │
│ Reason           │ Faster Phase 1, refactor Phase 2         │
│                  │                                          │
│ Variable Names   │ Keep ? prefix (?person, not person)      │
│ Reason           │ Clearer intent, SPARQL convention        │
│                  │                                          │
│ Op Ordering      │ W3C SPARQL execution order               │
│ Reason           │ Correct semantics, matches spec          │
│                  │                                          │
│ LeftJoin vs      │ Use LeftJoin for OPTIONAL                │
│ OptionalJoin     │ W3C standard, clear semantics            │
│                  │                                          │
│ Aggregates       │ Single operator per query                │
│ Reason           │ Simpler Phase 1, extend Phase 2          │
│                  │                                          │
│ Validation       │ Early, fail-fast approach                │
│ Reason           │ Better errors, enables optimization      │
│                  │                                          │
│ Error Handling   │ Collect all errors together              │
│ Reason           │ User sees complete picture               │
└──────────────────┴──────────────────────────────────────────┘
```

---

## Testing Strategy

```
┌───────────────────────────────────────────────────────┐
│ 100+ TEST CASES ORGANIZED BY CATEGORY                │
├─────────────────────────┬─────────────────────────────┤
│ Type Construction Tests │ 15 test cases               │
│                         │ BGP, Join, Filter, etc.     │
│                         │                             │
│ Visitor Pattern Tests   │ 10 test cases               │
│                         │ Print, traversal, format    │
│                         │                             │
│ Conversion Tests        │ 50+ test cases              │
│                         │ Simple to complex queries   │
│                         │ All SPARQL patterns         │
│                         │                             │
│ Validator Tests         │ 15+ test cases              │
│                         │ Valid + invalid expressions │
│                         │ Error detection             │
│                         │                             │
│ Context Tests           │ 5+ test cases               │
│                         │ Creation, builder, helpers  │
│                         │                             │
│ TOTAL                   │ 100+ comprehensive tests    │
└─────────────────────────┴─────────────────────────────┘
```

---

## Success Criteria

```
┌──────────────────────────────────────────────────────────┐
│ FUNCTIONALITY ✓                                          │
├──────────────────────────────────────────────────────────┤
│ ✓ All 11 operators working                             │
│ ✓ ASTToAlgebra converts all SPARQL patterns            │
│ ✓ Visitor pattern traverses all operators              │
│ ✓ Validator detects semantic errors                    │
│ ✓ Context carries metadata                             │
│                                                          │
├──────────────────────────────────────────────────────────┤
│ CODE QUALITY ✓                                           │
├──────────────────────────────────────────────────────────┤
│ ✓ No panics (proper error handling)                    │
│ ✓ All functions documented                             │
│ ✓ Error messages clear                                 │
│ ✓ Follows Dgraph conventions                           │
│ ✓ Efficient (<10ms algebra generation)                 │
│                                                          │
├──────────────────────────────────────────────────────────┤
│ TESTING ✓                                                │
├──────────────────────────────────────────────────────────┤
│ ✓ 100+ test cases                                      │
│ ✓ All tests passing                                    │
│ ✓ >90% code coverage                                   │
│ ✓ No flaky tests                                       │
│ ✓ Edge cases covered                                   │
│                                                          │
├──────────────────────────────────────────────────────────┤
│ INTEGRATION ✓                                            │
├──────────────────────────────────────────────────────────┤
│ ✓ No regressions in existing tests                     │
│ ✓ Backward compatible                                  │
│ ✓ Clean Phase 2+ integration points                    │
│ ✓ Design decisions documented                          │
└──────────────────────────────────────────────────────────┘
```

---

## Getting Started (Quick Roadmap)

```
TODAY (1 hour):
  1. Read 00_START_HERE.md (5 min)
  2. Read PHASE_1_INDEX.md (10 min)
  3. Read PHASE_1_QUICK_REFERENCE.md (30 min)
  4. Review DELIVERABLES_SUMMARY.md (15 min)

NEXT DAY (4 hours):
  1. Read PHASE_1_IMPLEMENTATION_PLAN.md Parts 1-2 (1 hour)
  2. Study sparql/ast.go (30 min)
  3. Study sparql/translator_extended.go (30 min)
  4. Study graphql/resolve/query_rewriter.go (30 min)
  5. Create feature branch

WEEK 1 (4-5 days):
  1. Implement Step 1: Algebra types (Days 1-2)
  2. Implement Step 2: Visitor pattern (Days 2-3)
  3. Write initial tests
  4. Review PHASE_1_IMPLEMENTATION_PLAN.md Part 3

WEEK 2 (4-5 days):
  1. Implement Step 3: AST→Algebra (Days 1-2)
  2. Write conversion tests (50+ cases)
  3. Implement Step 4: Validator (Days 3-4)
  4. Implement context.go (Day 5)

WEEK 3 (1-2 days):
  1. Complete testing
  2. Verify all tests pass
  3. Code review
  4. Merge
```

---

## Documents at a Glance

| File                           | Pages   | Words       | Read Time    | Best For          |
| ------------------------------ | ------- | ----------- | ------------ | ----------------- |
| 00_START_HERE.md               | 2       | 1000        | 5 min        | First orientation |
| PHASE_1_INDEX.md               | 3       | 1200        | 10 min       | Navigation        |
| PHASE_1_QUICK_REFERENCE.md     | 8       | 2000        | 30 min       | Executives        |
| PHASE_1_IMPLEMENTATION_PLAN.md | 20      | 4500        | 2-3 hrs      | Developers        |
| PHASE_1_CODE_STRUCTURE.md      | 10      | 2500        | 1 hr         | While coding      |
| DELIVERABLES_SUMMARY.md        | 8       | 2000        | 15 min       | Navigation        |
| **TOTAL**                      | **~50** | **~11,000** | **~4 hours** | **Full spec**     |

---

## The Bottom Line

✅ **You have a complete specification**  
✅ **Everything is detailed and justified**  
✅ **Code structure is provided**  
✅ **Testing strategy is specified**  
✅ **Timeline is realistic**  
✅ **Success criteria are measurable**

**You can start implementing immediately with confidence.**

---

## Next Action

👉 **Read PHASE_1_INDEX.md** (10 minutes)  
Then: **PHASE_1_QUICK_REFERENCE.md** (30 minutes)  
Then: **Start implementing following PHASE_1_IMPLEMENTATION_PLAN.md**

---

**Status**: ✅ SPECIFICATION COMPLETE  
**Location**: `/Users/nfeldman/repos/dgraph/sparql/`  
**Begin**: With 00_START_HERE.md (you are here!)
