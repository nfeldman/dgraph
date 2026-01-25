# SPARQL Algebra Foundation: Coordinated Development Setup

**Project**: Adding SPARQL Algebra with Ontology Support to Dgraph  
**Setup Date**: 2024  
**Status**: Ready for Implementation

---

## 📋 Quick Start

You have two coordinated development efforts:

1. **Phase 1 Implementation** (Steps 1-4 of algebra foundation)
   - Read: [PHASE_1_DEVELOPER_SPECIFICATION.md](PHASE_1_DEVELOPER_SPECIFICATION.md)
   - Branch: `feature/sparql-phase1-dev`
   - Duration: 10-15 days (1 developer)

2. **Phase 5 Research** (Ontology foundation design)
   - Read:
     [PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md](PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md)
   - Branch: `feature/sparql-ontology-foundation`
   - Duration: 8-10 days (research + design)

Both can proceed in parallel with **no blocking dependencies**.

---

## 📚 Documentation Map

| Document                                                                                     | Purpose                        | Audience            | Read Time |
| -------------------------------------------------------------------------------------------- | ------------------------------ | ------------------- | --------- |
| [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md)                                                 | Full 5-phase architecture      | Everyone (context)  | 30 min    |
| [PHASE_1_DEVELOPER_SPECIFICATION.md](PHASE_1_DEVELOPER_SPECIFICATION.md)                     | Detailed implementation guide  | Phase 1 Dev         | 45 min    |
| [PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md](PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md) | Research assignment            | Ontology Researcher | 30 min    |
| [GIT_WORKFLOW.md](GIT_WORKFLOW.md)                                                           | Branch structure & workflow    | Both teams          | 20 min    |
| [SPARQL_ALGEBRA_TODO.md](SPARQL_ALGEBRA_TODO.md)                                             | Task breakdown all phases      | Reference           | 15 min    |
| [DECISION_RECORD.md](DECISION_RECORD.md)                                                     | Architecture rationale         | Reference           | 15 min    |
| [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md)                                               | Integration with existing code | Reference           | 25 min    |

---

## 🎯 What You're Building

### Phase 1: Algebra Foundation (This Round)

**Goal**: Build intermediate representation for SPARQL queries using W3C SPARQL Algebra

**What gets built**:

- Algebra type system (11 operators: BGP, Join, Filter, LeftJoin, Union, etc.)
- Visitor pattern infrastructure for traversing algebra
- AST to Algebra converter (translates SPARQL AST to algebra form)
- Algebra validator (checks semantic correctness)

**Why it matters**:

- Current SPARQL path: AST → string → re-parse to DQL (loses information, no optimization)
- New SPARQL path: AST → Algebra → optimize → DQL (schema-aware, optimized)
- Matches how GraphQL queries work in Dgraph

**Output**: ~1,300 LOC across 4 files with 100+ tests

### Phase 5: Ontology Foundation (Research Phase)

**Goal**: Design how Dgraph will support OWL/RDFS semantic web ontologies

**What gets researched**:

- How to store ontology metadata in Dgraph
- How to perform type inference with inheritance
- How to optimize queries using ontology information
- Implementation architecture (storage, caching, reasoning)

**Why it matters**:

- Enables semantic queries: "show all People including Students"
- Enables ontology-aware optimization (cardinality estimation, type narrowing)
- Positions Dgraph as enterprise-grade knowledge graph

**Output**: Detailed design document with implementation plan for Phase 5.1-5.4

---

## 🌳 Git Branch Structure

```
feature/sparql-antlr
├── feature/sparql-phase1-dev          ← Phase 1 implementation work happens here
│   └── feature/sparql-phase1-mergeback ← Squashed commit merges here → feature/sparql-antlr
│
└── feature/sparql-ontology-foundation  ← Phase 5 research/design happens here
    └── feature/sparql-ontology-mergeback ← Squashed commit merges here → feature/sparql-antlr
```

**Key rules**:

- Work on: `feature/sparql-phase1-dev` or `feature/sparql-ontology-foundation`
- Merge to: `feature/sparql-phase1-mergeback` or `feature/sparql-ontology-mergeback` (via squash
  rebase)
- Final destination: `feature/sparql-antlr` (main feature branch)
- Clean up: Local working branches deleted after merge

---

## 👥 Team Roles

### Role 1: Phase 1 Developer

**Responsible for**: Steps 1-4 of algebra implementation

**Timeline**: Days 1-15

**What to do**:

1. Read [PHASE_1_DEVELOPER_SPECIFICATION.md](PHASE_1_DEVELOPER_SPECIFICATION.md) (45 min)
2. Start with Step 1: Create `sparql/algebra.go` with 11 operator types
3. Step 2: Create `sparql/algebra_visitor.go` with visitor pattern
4. Step 3: Create converter function in `sparql/algebra.go`
5. Step 4: Create `sparql/algebra_test.go` with 100+ tests

**Success criteria**:

- ✅ All algebra types implemented
- ✅ 100+ comprehensive tests
- ✅ >85% code coverage
- ✅ Integrates with existing SPARQL code
- ✅ Clear git history with logical commits

**Branch**: `feature/sparql-phase1-dev` → `feature/sparql-phase1-mergeback` → `feature/sparql-antlr`

---

### Role 2: Phase 5 Ontology Researcher

**Responsible for**: Research and design specification for ontology support

**Timeline**: Days 1-10

**What to do**:

1. Read [PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md](PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md)
   (30 min)
2. Investigate Dgraph codebase (graphql/schema/schema.go, worker/task.go, etc.)
3. Do literature review (W3C specs, other databases' approaches)
4. Design Phase 5 architecture
5. Write comprehensive design document: `sparql/ONTOLOGY_FOUNDATION_DESIGN.md`

**Success criteria**:

- ✅ Deep understanding of Dgraph's architecture
- ✅ Concrete design with specific file names and functions
- ✅ Realistic implementation plan (Phase 5.1-5.4)
- ✅ Design decisions with rationale
- ✅ Integrated with Phase 1-4 algebra foundation

**Branch**: `feature/sparql-ontology-foundation` → `feature/sparql-ontology-mergeback` →
`feature/sparql-antlr`

---

## 📋 Pre-Implementation Checklist

Before either team starts coding/research:

### For Phase 1 Developer

- [ ] Read ARCHITECTURE_SPEC.md (context)
- [ ] Read PHASE_1_DEVELOPER_SPECIFICATION.md (detailed spec)
- [ ] Read GIT_WORKFLOW.md (workflow)
- [ ] Verify on `feature/sparql-phase1-dev` branch
- [ ] Have existing SPARQL code context (sparql/ast.go, sparql/translator_extended.go)
- [ ] Understand W3C SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra

### For Phase 5 Researcher

- [ ] Read ARCHITECTURE_SPEC.md (context)
- [ ] Read PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md (spec)
- [ ] Read GIT_WORKFLOW.md (workflow)
- [ ] Verify on `feature/sparql-ontology-foundation` branch
- [ ] Have access to Dgraph codebase for investigation
- [ ] Understand W3C standards (SPARQL, RDFS, OWL overview)

---

## 🚀 Getting Started

### Phase 1 Developer

```bash
# Verify branch
git branch

# Should show: * feature/sparql-phase1-dev

# Read spec
cat sparql/PHASE_1_DEVELOPER_SPECIFICATION.md

# Start with Step 1: Create algebra.go
# Follow instructions in spec exactly

# Test frequently
go test ./sparql -v -run "Algebra"
```

### Phase 5 Researcher

```bash
# Verify branch
git branch

# Should show: * feature/sparql-ontology-foundation

# Read spec
cat sparql/PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md

# Start investigating code
# Days 1-3: Code archaeology
# Days 2-4: Literature review
# Days 5-8: Write ONTOLOGY_FOUNDATION_DESIGN.md
```

---

## ✅ Success Metrics

### Phase 1 Completion

**Code Quality**:

- [ ] 1,200-1,450 LOC across 4 files
- [ ] All 11 algebra operators implemented
- [ ] All types have String() and Accept() methods
- [ ] Visitor interface with 11 visit methods
- [ ] Converter handles 10+ query patterns
- [ ] Validator checks 5+ semantic rules
- [ ] Code follows Dgraph conventions
- [ ] No compiler warnings

**Testing**:

- [ ] 100+ test cases
- [ ] All tests pass (`go test ./sparql -v`)
- [ ] > 85% code coverage
- [ ] Tests cover happy path and error cases
- [ ] Clear error messages

**Integration**:

- [ ] Compiles with existing code
- [ ] Existing SPARQL tests still pass
- [ ] Only new files created (no modifications to existing code)
- [ ] Clear integration points for Phase 2

### Phase 5 Completion

**Document Quality**:

- [ ] 7 comprehensive sections
- [ ] 3,500-5,000 words
- [ ] Codebase investigation evident
- [ ] Design decisions justified
- [ ] Implementation plan detailed

**Completeness**:

- [ ] Research Summary (500-800 words)
- [ ] Ontology Model Design (800-1,000 words)
- [ ] Storage & Loading Strategy (600-800 words)
- [ ] Query Optimization integration (500-700 words)
- [ ] Implementation Plan (700-900 words)
- [ ] Reuse & Integration Decisions (400-600 words)
- [ ] Open Questions & Future Extensions (300-500 words)

**Quality**:

- [ ] Consistent with Phase 1-4 architecture
- [ ] Shows deep understanding of Dgraph
- [ ] All success criteria answered
- [ ] Professional, clear writing
- [ ] Proper markdown formatting

---

## 📞 Communication & Questions

### Documentation Links

- **Overall architecture**: ARCHITECTURE_SPEC.md
- **Phase 1 implementation**: PHASE_1_DEVELOPER_SPECIFICATION.md
- **Phase 5 research**: PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md
- **Workflow & git**: GIT_WORKFLOW.md
- **All tasks**: SPARQL_ALGEBRA_TODO.md
- **Architecture decisions**: DECISION_RECORD.md
- **System integration**: SYSTEM_INTEGRATION.md

### Common Questions

**Q: Can Phase 1 and Phase 5 work in parallel?**  
A: Yes! They have no blocking dependencies. Phase 1 lays algebra foundation, Phase 5 designs
ontology storage. Can work simultaneously.

**Q: When do I commit?**  
A: Commit logical chunks (e.g., "Step 1: Algebra types complete"). Aim for 4-5 commits per phase
before merging to mergeback branch.

**Q: How do I merge to feature/sparql-antlr?**  
A: Create mergeback branch, squash commits, create PR, merge to feature/sparql-antlr. See
GIT_WORKFLOW.md for details.

**Q: What if I need clarification on the spec?**  
A: Refer to ARCHITECTURE_SPEC.md for context. Read relevant sections multiple times. The
specifications are comprehensive.

**Q: How much testing is "enough"?**  
A: For Phase 1: 100+ test cases, >85% coverage. For Phase 5: Just the design document (no code yet).

---

## 📊 Project Timeline

```
Today: Setup & documentation complete
↓
Days 1-15: Phase 1 Dev (algebra implementation)
Days 1-10: Phase 5 Dev (ontology research) [parallel]
↓
Day 15: Phase 1 mergeback branch created
Day 10: Phase 5 mergeback branch created
↓
Day 16: Phase 1 PR to feature/sparql-antlr
Day 11: Phase 5 PR to feature/sparql-antlr
↓
Day 17: Both merged to feature/sparql-antlr ✅
↓
Day 18+: Phase 2-4 implementation (follows Phase 1 success)
Future: Phase 5 implementation (based on research)
```

---

## 🎓 Learning Resources

### W3C Standards

- SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
- SPARQL Semantics: https://www.w3.org/TR/sparql11-semantics/
- RDFS: https://www.w3.org/TR/rdf-schema/
- OWL: https://www.w3.org/OWL/

### Dgraph Code

- graphql/schema/schema.go - Type system
- graphql/resolve/query_rewriter.go - Query optimization patterns
- worker/task.go - Query planning
- edgraph/server.go - Query entry point
- posting/list.go - Storage abstractions

### Related Documentation

- [ARCHITECTURE_SPEC.md](ARCHITECTURE_SPEC.md) - Full context
- [DECISION_RECORD.md](DECISION_RECORD.md) - Why these choices
- [SYSTEM_INTEGRATION.md](SYSTEM_INTEGRATION.md) - Integration patterns

---

## ✨ Next Steps

1. **Assign roles**: Who is Phase 1 developer? Who is Phase 5 researcher?
2. **Start work**: Each person reads their specification and begins
3. **Commit regularly**: Phase 1 every logical step, Phase 5 when document complete
4. **Create mergeback branches** when work is done
5. **Create PRs** to feature/sparql-antlr
6. **Merge & celebrate** ✅
7. **Plan Phase 2**: Once Phase 1 merged

---

## 📝 Document Versions

| Document                                     | Version | Date | Status                   |
| -------------------------------------------- | ------- | ---- | ------------------------ |
| PHASE_1_DEVELOPER_SPECIFICATION.md           | 1.0     | 2024 | Ready for implementation |
| PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md | 1.0     | 2024 | Ready for research       |
| GIT_WORKFLOW.md                              | 1.0     | 2024 | Ready for use            |
| ARCHITECTURE_SPEC.md                         | 1.0     | 2024 | Reference                |

---

## 🏁 Success = You're Building

After this phase completes:

✅ **Phase 1**: SPARQL queries use W3C Algebra intermediate representation  
✅ **Phase 1**: Algebra can be optimized, inspected, and debugged  
✅ **Phase 1**: Foundation for schema integration (Phase 3)  
✅ **Phase 5**: Complete design for ontology support in Dgraph

This is the foundation that enables:

- ✅ Schema-aware query optimization (Phase 2-3)
- ✅ RBAC integration (Phase 4)
- ✅ OWL/RDFS reasoning (Phase 5)

Good luck! You're building something great. 🚀
