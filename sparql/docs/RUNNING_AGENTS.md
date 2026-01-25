# Running the Development Agents

**Guide for Launching Coordinated Development**

---

## Overview

You now have:

1. ✅ **Phase 1 Specification** - Detailed implementation guide with 4 sequential steps
2. ✅ **Phase 5 Specification** - Research assignment with 7 detailed sections
3. ✅ **Git Workflow** - Branch structure and coordination guide
4. ✅ **Supporting Docs** - Architecture, decisions, integration points

This guide explains how to actually execute this work.

---

## Option 1: Run With Human Developers

If you have human developers available:

### Phase 1 Developer

```bash
# Share these files:
- sparql/ARCHITECTURE_SPEC.md
- sparql/PHASE_1_DEVELOPER_SPECIFICATION.md
- sparql/GIT_WORKFLOW.md
- sparql/COORDINATED_DEVELOPMENT.md

# Give instructions:
1. Read the specs above
2. Switch to feature/sparql-phase1-dev branch
3. Follow Steps 1-4 in PHASE_1_DEVELOPER_SPECIFICATION.md
4. Implement ~1,300 LOC across 4 files
5. Write 100+ tests
6. Create mergeback branch when done
7. Create PR to feature/sparql-antlr

# Timeline: 10-15 days
```

### Phase 5 Researcher

```bash
# Share these files:
- sparql/ARCHITECTURE_SPEC.md
- sparql/PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md
- sparql/GIT_WORKFLOW.md
- sparql/COORDINATED_DEVELOPMENT.md

# Give instructions:
1. Read the specs above
2. Switch to feature/sparql-ontology-foundation branch
3. Follow research assignment in PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md
4. Investigate Dgraph codebase (Days 1-3)
5. Review literature (Days 2-4)
6. Write ONTOLOGY_FOUNDATION_DESIGN.md (Days 5-8)
7. Create mergeback branch when done
8. Create PR to feature/sparql-antlr

# Timeline: 8-10 days
```

---

## Option 2: Use AI Agents

If you want to use AI agents to perform the work:

### Launch Phase 1 Development Agent

Create a subagent with this prompt:

```
Role: SPARQL Algebra Implementation Developer

Assignment: Implement Phase 1 (Steps 1-4) of SPARQL Algebra foundation

Context:
- Dgraph is a distributed knowledge graph database
- We're adding SPARQL query support with W3C algebraic optimization
- Current SPARQL path: AST → string → re-parse (loses info, no optimization)
- New path: AST → Algebra → optimize → DQL (schema-aware)
- Phase 1 builds the algebraic intermediate representation

Working Branch: feature/sparql-phase1-dev (already created in git)
Target Branch: feature/sparql-phase1-mergeback, then feature/sparql-antlr

Your Task (do these sequentially):

STEP 1: Algebra Type System (Days 1-4)
- Create sparql/algebra.go (~400-500 LOC)
- Implement 11 operator types: BGP, Join, Filter, LeftJoin, Union, Project, Aggregate, Bind, Distinct, OrderBy, Limit
- Each type should have:
  * Accept(visitor AlgebraVisitor) interface{}
  * String() string for debugging
  * Variables() []string for variable tracking
- Make sure types match W3C SPARQL Algebra spec
- See PHASE_1_DEVELOPER_SPECIFICATION.md for detailed type definitions

STEP 2: Visitor Pattern (Days 5-7)
- Create sparql/algebra_visitor.go (~100-150 LOC)
- Define AlgebraVisitor interface with 11 Visit methods
- Implement AlgebraPrinter (prints algebra tree with indentation)
- Implement VariableCollector (collects all variables in expression)
- Support both top-down and bottom-up traversal patterns

STEP 3: AST to Algebra Converter (Days 8-12)
- Add ASTToAlgebra(query *SPARQLQueryImpl) (AlgebraExpr, error) function to algebra.go
- Implement 10-step conversion algorithm:
  1. Parse patterns
  2. Build BGP nodes
  3. Join BGPs
  4. Apply filters
  5. Handle OPTIONAL → LeftJoin
  6. Handle UNION → Union
  7. Build filter expressions
  8. Apply aggregates
  9. Apply projections
  10. Apply modifiers
- Add helper functions for join chains, optional conversion, filter wrapping
- Test with examples like SELECT ?person ?name WHERE { ... FILTER ... }

STEP 4: Algebra Validator (Days 13-15)
- Add AlgebraValidator type with semantic checks
- Validate variable scoping (vars referenced must be defined)
- Check circular dependencies in BIND expressions
- Validate filter expressions and aggregate functions
- Produce clear error messages
- Implement as visitor over algebra expressions

Testing:
- Create sparql/algebra_test.go (~400-500 LOC)
- Write 100+ test cases covering:
  * Each algebra type construction (10 tests)
  * Visitor traversal (20 tests)
  * AST conversion (40+ tests)
  * Validation rules (25 tests)
- Aim for >85% code coverage
- Run: go test ./sparql -v -run "Algebra"

Success Criteria:
- All 4 files created (~1,300 LOC total)
- 100+ tests all passing
- >85% code coverage
- Code integrates with existing SPARQL code
- Clear git history with logical commits

Reference Materials:
- PHASE_1_DEVELOPER_SPECIFICATION.md - Detailed implementation guide
- W3C SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
- ARCHITECTURE_SPEC.md - Full context

Git Workflow:
1. You're already on feature/sparql-phase1-dev
2. Make commits as you complete each step
3. When done, create feature/sparql-phase1-mergeback branch
4. Rebase to squash commits: git rebase -i feature/sparql-antlr
5. Create PR to feature/sparql-antlr
6. See GIT_WORKFLOW.md for detailed workflow

Do NOT:
- Modify existing SPARQL code (only add new files)
- Skip testing
- Commit broken code
- Change unrelated files

Output:
- Complete, working implementation of Phase 1
- All tests passing
- Ready to merge to feature/sparql-antlr
```

### Launch Phase 5 Research Agent

Create a subagent with this prompt:

```
Role: SPARQL Ontology Foundation Researcher & Designer

Assignment: Design Phase 5 (ontology support) for SPARQL algebra system

Context:
- Dgraph is implementing SPARQL Algebra (W3C standard intermediate representation)
- Phase 1-4 handle query parsing, optimization, schema integration, authorization
- Phase 5 will add support for OWL/RDFS semantic web ontologies
- This is research phase to design the architecture before implementation

Working Branch: feature/sparql-ontology-foundation (already created in git)
Target Branch: feature/sparql-ontology-mergeback, then feature/sparql-antlr

Your Task (research and design):

Research Days 1-3: Code Investigation
- Read graphql/schema/schema.go - how Dgraph stores types and predicates
- Read worker/task.go - query planning and cardinality estimation
- Read edgraph/access.go - how RBAC authorization works (pattern for ontology rules)
- Read posting/list.go - index structures and statistics
- Read sparql/ code - understand current architecture
- Document findings about:
  * Current schema system
  * Type inference capabilities
  * Authorization patterns
  * Where ontology would integrate

Research Days 2-4: Literature Review (overlaps with code investigation)
- W3C SPARQL: https://www.w3.org/TR/sparql11-query/
- RDFS: https://www.w3.org/TR/rdf-schema/ (simpler than OWL)
- OWL: https://www.w3.org/OWL/ (understand complexity)
- How other databases handle ontologies (skim Neo4j, Neptune documentation)
- Document what's feasible for Dgraph

Design Days 5-8: Write Comprehensive Specification

Create sparql/ONTOLOGY_FOUNDATION_DESIGN.md with these 7 sections:

SECTION 1: Research Summary (500-800 words)
- Summarize findings from investigating Dgraph code
- How does Dgraph currently store schema?
- How does type inference work?
- What would be needed for ontology support?
- Include limitations and solutions

SECTION 2: Ontology Model Design (800-1000 words)
- How do OWL classes map to Dgraph types?
- How do OWL properties map to Dgraph predicates?
- How are RDFS subclass relationships represented?
- How are class restrictions stored?
- How are property ranges/domains stored?
- Type inference rules with inheritance
- Include example mappings (OWL → Dgraph triples)

SECTION 3: Storage & Loading Strategy (600-800 words)
- How to load OWL/RDFS files (formats, parsers)
- Should ontologies be stored in Dgraph or separately?
- In-memory caching approach
- Query-time rule application
- Hybrid storage recommendation with rationale

SECTION 4: Query Optimization with Ontologies (500-700 words)
- Type narrowing with inheritance
- How class hierarchy affects cardinality
- Join ordering improvements
- Filter simplification with ontology knowledge
- Property-based reasoning
- Integration with worker/task.go (minimal changes needed?)

SECTION 5: Implementation Plan (700-900 words)
- Break Phase 5 into 4 sub-phases:
  * Phase 5.1: Ontology Storage (week 1-2)
  * Phase 5.2: Type Inference (week 3-4)
  * Phase 5.3: Query Integration (week 5-6)
  * Phase 5.4: SPARQL Reasoner (week 7-8)
- For each sub-phase: files to create, functions, integration points, tests
- Specific file names and directory structure

SECTION 6: Reuse & Integration Decisions (400-600 words)
- Can we reuse graphql/schema/schema.go? How?
- Which Go RDF libraries are best?
- Can we leverage posting/list.go for ontology indices?
- Recommend specific libraries with justification
- Create comparison table of options and recommendations

SECTION 7: Open Questions & Future Extensions (300-500 words)
- How much OWL reasoning is actually needed?
- When should reasoning happen (compile-time or execution)?
- Ontology versioning approach?
- Performance expectations?
- Future extensions (SHACL, custom rules, etc.)

Testing:
- No code to test, just design document
- Validate that all success criteria are met
- Proofread for clarity

Success Criteria:
- Document has all 7 sections
- 3,500-5,000 words total
- Shows deep investigation of Dgraph code
- Concrete file names and function signatures
- Realistic implementation plan
- Design decisions justified
- Integrated with Phase 1-4 architecture
- Professional quality writing

Reference Materials:
- PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md - Detailed assignment
- ARCHITECTURE_SPEC.md - Full Phase 1-5 context
- SYSTEM_INTEGRATION.md - Integration patterns

Git Workflow:
1. You're already on feature/sparql-ontology-foundation
2. Commit document when research is complete: git add sparql/ONTOLOGY_FOUNDATION_DESIGN.md && git commit ...
3. Create feature/sparql-ontology-mergeback branch when done
4. Rebase to squash commits
5. Create PR to feature/sparql-antlr
6. See GIT_WORKFLOW.md for detailed workflow

Output:
- ONTOLOGY_FOUNDATION_DESIGN.md complete with all sections
- Ready for Phase 5 implementation to begin
```

---

## How to Actually Run This

### Option 2A: Manual Agent Execution

If you want to run subagents manually:

```bash
# Make sure you're at the workspace root
cd /Users/nfeldman/repos/dgraph

# Verify branches
git branch -a | grep "feature/sparql"

# Verify you can see the specifications
cat sparql/PHASE_1_DEVELOPER_SPECIFICATION.md | head -50
cat sparql/PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md | head -50

# Launch Phase 1 agent with detailed prompt
# (Use the prompt template above with runSubagent tool)

# Launch Phase 5 agent with detailed prompt
# (Use the prompt template above with runSubagent tool)
```

### Option 2B: What Happens After Agent Launches

Once you launch an agent:

1. Agent receives detailed prompt
2. Agent reads specification files from workspace
3. Agent reads referenced Dgraph code
4. Agent creates/modifies files
5. Agent makes git commits
6. Agent creates mergeback branch
7. Agent returns summary of work completed

You then:

1. Review the work
2. Create PR to feature/sparql-antlr
3. Merge when approved
4. Delete working branches

---

## Implementation Timeline with Agents

```
Now: Specifications complete, agents ready to launch
↓
Agent 1 (Phase 1): Days 1-15
- Step 1: Algebra types (days 1-4)
- Step 2: Visitor pattern (days 5-7)
- Step 3: AST converter (days 8-12)
- Step 4: Validator (days 13-15)
- Testing throughout
↓
Agent 2 (Phase 5): Days 1-10 [PARALLEL]
- Investigate code (days 1-3)
- Literature review (days 2-4)
- Design document (days 5-8)
- Validation (days 8-10)
↓
Day 11: Phase 5 research complete
↓
Day 15: Phase 1 implementation complete
↓
Day 16-17: Code review and merge both
↓
Day 17: Both merged to feature/sparql-antlr ✅
```

---

## What's Ready for Agents

✅ **Phase 1 Developer needs**:

- [x] Detailed specification (PHASE_1_DEVELOPER_SPECIFICATION.md)
- [x] Code examples and type definitions
- [x] 10-step converter algorithm
- [x] Test case examples (100+ specified)
- [x] Git workflow documentation

✅ **Phase 5 Researcher needs**:

- [x] Research assignment (PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md)
- [x] Codebase investigation targets
- [x] 7-section document template
- [x] Example design patterns
- [x] Git workflow documentation

✅ **Both teams have**:

- [x] Context documentation (ARCHITECTURE_SPEC.md)
- [x] Integration guide (SYSTEM_INTEGRATION.md)
- [x] Workflow guide (GIT_WORKFLOW.md)
- [x] Reference specifications (SPARQL_ALGEBRA_TODO.md, DECISION_RECORD.md)

---

## Ready to Launch

Everything is prepared. You can:

1. **Use human developers**: Share the specifications and let them work
2. **Use AI agents**: Copy the prompts above and launch via runSubagent
3. **Hybrid**: Some parts human, some parts agent

All have the same outcome: Phase 1 + Phase 5 complete in parallel in 2-3 weeks.

Good luck! 🚀
