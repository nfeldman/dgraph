# Phase 1 SPARQL Algebra Implementation - Complete Specification Index

**Prepared**: January 25, 2026  
**Repository**: `/Users/nfeldman/repos/dgraph`  
**Branch**: `feature/sparql-phase1-dev`  
**Status**: ✅ Ready for Implementation

---

## Overview

This is a **complete, production-ready specification** for implementing Phase 1 of the SPARQL
Algebra project. The specification consists of **4 comprehensive documents** totaling ~11,000 words
of detailed planning, design decisions, code structure, and testing strategy.

### What Phase 1 Delivers

**Three new files totaling ~1,300 LOC**:

- `sparql/algebra.go` (500-600 LOC): All algebra types, converter, validator
- `sparql/algebra_visitor.go` (100-150 LOC): Visitor pattern infrastructure
- `sparql/context.go` (100 LOC): Execution context carrier

**Plus comprehensive tests**:

- `sparql/algebra_test.go` (400-500 LOC): 100+ test cases

### Timeline

- **Duration**: 10-15 days for one developer
- **Effort**: Moderate (~1,300 LOC, well-specified)
- **Parallel**: Can be split across 2 developers (6-8 days)

---

## The Four Documents

### 1. PHASE_1_QUICK_REFERENCE.md ⭐ START HERE

**Read Time**: 30 minutes  
**Type**: Executive summary and quick reference  
**Best for**: Getting oriented, understanding scope, quick lookup

**Contains**:

- High-level algorithm explanation
- 4-step overview with durations
- Key design decisions table
- Testing strategy matrix
- Rollout strategy (4-phase approach)
- FAQ (8 common questions)
- Links and resources

**Start here to**: Understand the big picture in 30 minutes

---

### 2. PHASE_1_IMPLEMENTATION_PLAN.md ⭐ MAIN DOCUMENT

**Read Time**: 2-3 hours  
**Type**: Detailed implementation specification  
**Best for**: Complete understanding, design rationale, implementation guidance

**Contains (6 main parts)**:

- **Part 1: Algebra Type System** (3-4 days)
  - AlgebraExpr interface (design purpose, semantics)
  - 11 concrete operators (BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, Bind, Distinct,
    OrderBy, Limit)
  - Each operator: Purpose, semantics, String() format, example
  - Implementation checklist with tests

- **Part 2: Visitor Pattern Infrastructure** (2-3 days)
  - Design philosophy and purpose
  - AlgebraVisitor interface definition
  - AlgebraPrinter concrete implementation
  - Visitor pattern semantics
  - Example tests

- **Part 3: AST to Algebra Converter** (4-5 days)
  - Design goals and algorithm overview
  - ASTToAlgebra function signature and detailed algorithm
  - 10-step conversion process
  - Pattern conversion (BGP, Optional, Union, Filter)
  - Helper functions (convertPatterns, applyFilters, etc.)
  - Aggregate, BIND, DISTINCT, ORDER BY, LIMIT handling
  - 30+ conversion test strategy
  - Example test case

- **Part 4: Algebra Validator** (2-3 days)
  - Design goals and key validations
  - AlgebraValidator type and methods
  - Variable scope tracking algorithm
  - Expression variable extraction
  - Full validation algorithm with pseudocode
  - Error message format and examples
  - Validator test strategy

- **Part 5: Execution Context** (1-2 days)
  - Purpose and design
  - SPARQLExecutionContext struct
  - Factory method and fluent builder pattern
  - Helper methods

- **Part 6: Comprehensive Testing** (2-3 days)
  - Test file organization
  - 100+ test case breakdown
  - Test coverage matrix
  - Test quality guidelines
  - Example test cases

- **Supporting sections**:
  - Integration points with existing code
  - 6 key design decisions with reasoning
  - Testing strategy in detail
  - Build & deployment checklist
  - Estimated timeline
  - Blockers & clarifications (none)
  - Success criteria

**Start here to**: Understand complete implementation approach

---

### 3. PHASE_1_CODE_STRUCTURE.md ⭐ REFERENCE WHILE CODING

**Read Time**: 1 hour  
**Type**: Concrete code examples and patterns  
**Best for**: During implementation, code reference, pattern examples

**Contains (5 main sections)**:

- **Section 1: algebra.go File Structure** (500-600 LOC outline)
  - AlgebraExpr interface
  - 11 operator types with complete implementations
  - Accept() method pattern
  - String() method pattern
  - ASTToAlgebra function with full pseudocode
  - Helper functions (convertPatterns, applyFilters, etc.)
  - AlgebraValidator with all methods
  - extractReferencedVariables function

- **Section 2: algebra_visitor.go File Structure** (150 LOC outline)
  - AlgebraVisitor interface
  - AlgebraPrinter class
  - All Visit methods (11 per operator)
  - Indentation handling

- **Section 3: context.go File Structure** (100 LOC outline)
  - SPARQLExecutionContext struct
  - Factory and fluent methods
  - Helper methods

- **Section 4: algebra_test.go Test Structure** (500 LOC outline)
  - 5 test categories
  - Example test for each category
  - Test quality template
  - Helper functions

- **Supporting sections**:
  - File size estimates
  - Implementation sequence (recommended order)
  - 7 common patterns & conventions

**Use this to**: Copy code structure while implementing

---

### 4. DELIVERABLES_SUMMARY.md

**Read Time**: 15 minutes  
**Type**: Roadmap and navigation guide  
**Best for**: Understanding what documents exist, how to use them, workflow

**Contains**:

- Overview of 4 documents
- How to use documents by role (PM, Developer, Reviewer)
- What's fully specified vs. intentionally excluded
- Implementation timeline with effort estimates
- Recommended pace and critical path
- How to execute (week-by-week breakdown)
- Complete success criteria checklist
- Reference document locations
- Pre-implementation checklist
- Support and escalation guidance

**Use this to**: Navigate the specification

---

## How to Use This Specification

### If you have 30 minutes

→ Read: **PHASE_1_QUICK_REFERENCE.md**

- Get the big picture
- Understand design decisions
- See testing strategy

### If you have 2-3 hours

→ Read in order:

1. **PHASE_1_QUICK_REFERENCE.md** (30 min)
2. **PHASE_1_IMPLEMENTATION_PLAN.md** (2 hours)

- Understand complete approach
- Review design rationale
- See detailed algorithms

### If you're implementing

→ Use in order:

1. **PHASE_1_QUICK_REFERENCE.md** (orientation)
2. **PHASE_1_IMPLEMENTATION_PLAN.md** (detailed design for each part)
3. **PHASE_1_CODE_STRUCTURE.md** (while coding each file)
4. **DELIVERABLES_SUMMARY.md** (navigation and status tracking)

### If you're reviewing code

→ Reference:

1. **PHASE_1_IMPLEMENTATION_PLAN.md** (design decisions)
2. **PHASE_1_CODE_STRUCTURE.md** (code patterns)
3. **PHASE_1_QUICK_REFERENCE.md** (success criteria)

### If you're planning Phase 2

→ Read:

1. **PHASE_1_QUICK_REFERENCE.md** (understand algebra structure)
2. **PHASE_1_IMPLEMENTATION_PLAN.md** (understand design decisions)

- You'll know where algebra fits in optimization pipeline

---

## Quick Facts

| Aspect                    | Details                                                                                     |
| ------------------------- | ------------------------------------------------------------------------------------------- |
| **Total LOC to write**    | ~1,300                                                                                      |
| **Total files to create** | 4 (algebra.go, algebra_visitor.go, context.go, algebra_test.go)                             |
| **Number of operators**   | 11 (BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, Bind, Distinct, OrderBy, Limit) |
| **Test cases**            | 100+ (15 type, 10 visitor, 50+ conversion, 15+ validator, 5+ context)                       |
| **Implementation steps**  | 4 (sequential)                                                                              |
| **Timeline**              | 10-15 days (1 developer) or 6-8 days (2 developers)                                         |
| **Complexity**            | Moderate (well-specified, no ambiguities)                                                   |
| **Dependencies**          | None (Phase 1 is standalone)                                                                |
| **Breaking changes**      | None (fully backward compatible)                                                            |

---

## Key Deliverables

### What You Get

✅ **Complete specification** - 11,000+ words, 4 documents  
✅ **Code structure** - File organization, function signatures, pseudocode  
✅ **Design decisions** - All major choices justified  
✅ **Testing strategy** - 100+ test cases specified  
✅ **Implementation path** - Step-by-step guidance  
✅ **Example conversions** - Real SPARQL→Algebra examples  
✅ **Integration plan** - How Phase 1 fits into larger system

### What You Implement

1. **algebra.go** - Types (11 operators), converter, validator
2. **algebra_visitor.go** - Visitor interface, printer implementation
3. **context.go** - Execution context carrier
4. **algebra_test.go** - Comprehensive test suite (100+ tests)

### What Comes Next (Phase 2+)

- Optimization passes (filter pushdown, join reordering)
- Schema integration
- Authorization rules
- Multi-phase execution

---

## Critical Path to Implementation

```
Start
  ↓
Read PHASE_1_QUICK_REFERENCE.md (30 min)
  ↓
Read PHASE_1_IMPLEMENTATION_PLAN.md (2 hrs)
  ↓
Study existing code (ast.go, translator_extended.go) (2 hrs)
  ↓
Create feature branch
  ↓
Implement (following PHASE_1_IMPLEMENTATION_PLAN.md):
  ├─ Step 1: Algebra types (Days 1-4)
  ├─ Step 2: Visitor pattern (Days 2-3)  [can overlap Step 1]
  ├─ Step 3: AST→Algebra (Days 4-8)
  ├─ Step 4: Validator (Days 9-11)
  └─ Integration & testing (Days 12-15)
  ↓
Run tests: go test ./sparql -v
  ↓
Code review
  ↓
Merge to main
```

---

## Document Contents Summary

### PHASE_1_QUICK_REFERENCE.md

- Big picture explanation
- Why we're doing this
- 4-step overview
- Design decisions table (1 page)
- Testing strategy
- Real query example
- Performance expectations
- Rollout strategy
- FAQ
- ~2000 words, 30 min read

### PHASE_1_IMPLEMENTATION_PLAN.md

- Executive summary
- **Part 1**: Algebra Type System (design + 11 operators)
- **Part 2**: Visitor Pattern Infrastructure
- **Part 3**: AST→Algebra Converter (10-step algorithm)
- **Part 4**: Algebra Validator
- **Part 5**: Execution Context
- **Part 6**: Comprehensive Testing
- Integration points
- Design decisions (detailed justification)
- Testing strategy in detail
- Success criteria
- ~4500 words, 2-3 hour read

### PHASE_1_CODE_STRUCTURE.md

- algebra.go structure (complete outline with pseudocode)
- algebra_visitor.go structure
- context.go structure
- algebra_test.go structure
- File size estimates
- Implementation sequence
- Common patterns & conventions
- ~2500 words, 1 hour read

### DELIVERABLES_SUMMARY.md

- What you're receiving
- How to use the documents
- What's specified vs. not
- Implementation timeline
- Rollout by week
- Success criteria checklist
- Pre-implementation checklist
- ~2000 words, 15 min read

---

## Success Criteria (At a Glance)

### Functionality ✓

- [ ] All 11 algebra operators working
- [ ] ASTToAlgebra converts all SPARQL patterns
- [ ] Visitor pattern traverses all operators
- [ ] Validator detects semantic errors
- [ ] Context carries metadata

### Code Quality ✓

- [ ] No panics (proper error handling)
- [ ] All public functions documented
- [ ] Error messages clear
- [ ] Follows Dgraph conventions
- [ ] Efficient (<10ms for typical queries)

### Testing ✓

- [ ] 100+ test cases
- [ ] All tests passing
- [ ] > 90% code coverage
- [ ] No flaky tests
- [ ] Edge cases covered

### Integration ✓

- [ ] No regressions in existing tests
- [ ] Backward compatible
- [ ] Clean Phase 2+ integration points
- [ ] Documented design decisions

---

## Getting Started

### Step 1: Read (1 hour)

1. Read PHASE_1_QUICK_REFERENCE.md (30 min)
2. Read DELIVERABLES_SUMMARY.md (15 min)
3. Skim PHASE_1_IMPLEMENTATION_PLAN.md (15 min for overview)

### Step 2: Study (2 hours)

1. Review sparql/ast.go (understand AST types)
2. Review sparql/translator_extended.go (understand current flow)
3. Review graphql/resolve/query_rewriter.go (understand visitor pattern)

### Step 3: Plan (30 min)

1. Create feature branch: `git checkout -b feature/sparql-algebra-foundation`
2. Plan calendar (10-15 days or pair with someone for 6-8 days)
3. Review success criteria

### Step 4: Implement (10-15 days)

Follow PHASE_1_IMPLEMENTATION_PLAN.md Part 1-6 in order

### Step 5: Review & Merge

1. Run tests: `go test ./sparql -v`
2. Get code review
3. Merge to main

---

## Document Locations

All files located in: `/Users/nfeldman/repos/dgraph/sparql/`

```
Phase 1 Implementation Specification:
├── DELIVERABLES_SUMMARY.md (this file)
├── PHASE_1_QUICK_REFERENCE.md ⭐ START HERE
├── PHASE_1_IMPLEMENTATION_PLAN.md ⭐ MAIN REFERENCE
├── PHASE_1_CODE_STRUCTURE.md ⭐ CODE REFERENCE
│
Existing Documentation:
├── PHASE_1_GETTING_STARTED.md (original requirements)
├── ARCHITECTURE_SPEC.md (system design)
│
To be Created (following this specification):
├── algebra.go
├── algebra_visitor.go
├── context.go
└── algebra_test.go
```

---

## FAQ

**Q: Which document should I read first?**  
A: PHASE_1_QUICK_REFERENCE.md (30 min)

**Q: How long is the full specification?**  
A: ~11,000 words across 4 documents (3-4 hours to read completely)

**Q: Can I implement without reading everything?**  
A: Not recommended. Read at least QUICK_REFERENCE + IMPLEMENTATION_PLAN parts relevant to what
you're working on.

**Q: Where do I find code examples?**  
A: PHASE_1_CODE_STRUCTURE.md - complete file structures with pseudocode

**Q: What if I find ambiguities?**  
A: Reference the relevant part of IMPLEMENTATION_PLAN.md, or note for code review discussion

**Q: Can I modify the plan?**  
A: Yes, but document changes. Specification is complete, adjustments should have clear rationale.

**Q: How many test cases do I really need to write?**  
A: Specification calls for 100+. This is comprehensive and necessary for quality.

**Q: What's the minimum viable Phase 1?**  
A: QUICK_REFERENCE.md section "Minimal MVP vs Full Implementation"

---

## Next Actions

### Immediate (This Week)

1. ✅ Read PHASE_1_QUICK_REFERENCE.md
2. ✅ Read PHASE_1_IMPLEMENTATION_PLAN.md (Parts 1-2)
3. ✅ Study existing code (ast.go, translator_extended.go)
4. ✅ Create feature branch

### Short-term (Week 1-2)

1. ✅ Implement Part 1: Algebra Type System
2. ✅ Implement Part 2: Visitor Pattern
3. ✅ Begin Part 3: AST→Algebra Converter

### Medium-term (Week 2-3)

1. ✅ Complete Part 3: AST→Algebra
2. ✅ Implement Part 4: Validator
3. ✅ Implement Part 5: Context
4. ✅ Complete Part 6: Testing

### Before Merge

1. ✅ All 100+ tests passing
2. ✅ No regressions
3. ✅ Code review approved
4. ✅ Documentation complete

---

## Support

### For Design Questions

→ Reference: PHASE_1_IMPLEMENTATION_PLAN.md (Part 1-5)

### For Code Structure

→ Reference: PHASE_1_CODE_STRUCTURE.md

### For Testing Strategy

→ Reference: PHASE_1_IMPLEMENTATION_PLAN.md Part 6

### For Timeline/Scope

→ Reference: PHASE_1_QUICK_REFERENCE.md

### For Integration

→ Reference: PHASE_1_IMPLEMENTATION_PLAN.md "Integration Points"

---

## Final Note

This specification is **complete and ready for implementation**. It provides:

✅ **What to build** - 11 algebra operators, 4 files  
✅ **Why we're building it** - Enable optimization and schema awareness  
✅ **How to build it** - 4 sequential steps with detailed algorithms  
✅ **How to test it** - 100+ test cases with strategy  
✅ **How to integrate it** - Clean interface for Phase 2+

Everything is specified in sufficient detail to implement without guessing.

**Begin with PHASE_1_QUICK_REFERENCE.md and proceed systematically.**

---

**Status**: ✅ READY FOR IMPLEMENTATION  
**Date**: January 25, 2026  
**Location**: `/Users/nfeldman/repos/dgraph/sparql/`
