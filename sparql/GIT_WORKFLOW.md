# SPARQL Algebra Project: Git Workflow

**Status**: Initial Setup  
**Main Feature Branch**: `feature/sparql-antlr`  
**Date**: 2024

---

## Branch Structure

```
feature/sparql-antlr (main feature branch)
├── feature/sparql-phase1-dev (working branch - Phase 1 implementation)
│   └── feature/sparql-phase1-mergeback (squash branch - ready to merge)
│
└── feature/sparql-ontology-foundation (working branch - Phase 5 research)
    └── feature/sparql-ontology-mergeback (squash branch - ready to merge)
```

---

## Workflow Overview

### For Phase 1 Developer

**Working branch**: `feature/sparql-phase1-dev`

1. **Start work on Phase 1**

   ```bash
   # Already on feature/sparql-phase1-dev
   git log --oneline  # Verify current branch
   ```

2. **Implement Steps 1-4**
   - Create algebra.go (type system)
   - Create algebra_visitor.go (visitor pattern)
   - Create context.go (execution context)
   - Create algebra_test.go (comprehensive tests)
   - Make commits as you go (1 commit per logical chunk)

3. **Test constantly**

   ```bash
   go test ./sparql -v -run "Algebra"
   ```

4. **When Step 1 complete** (e.g., algebra types done):

   ```bash
   git add sparql/algebra.go
   git commit -m "feat(sparql): Add algebra type system

   - 11 operator types (BGP, Join, Filter, etc.)
   - String() and Accept() implementations
   - Helper methods for variable tracking
   - Matches W3C SPARQL Algebra spec"
   ```

5. **Repeat for Steps 2, 3, 4** with similar logical commits

6. **When Phase 1 complete** (all 4 steps done, all tests pass):

   ```bash
   # Create mergeback branch
   git checkout -b feature/sparql-phase1-mergeback feature/sparql-phase1-dev

   # View commits since feature/sparql-antlr
   git log --oneline feature/sparql-antlr..feature/sparql-phase1-mergeback

   # Interactive rebase to clean up commits
   git rebase -i feature/sparql-antlr
   # In editor: squash all commits into first one
   # Edit final commit message to summarize Phase 1

   # Final commit message format:
   # feat(sparql): SPARQL Algebra foundation (Phase 1)
   #
   # Implements W3C SPARQL Algebra as intermediate representation:
   # - Algebra type system (11 operators)
   # - Visitor pattern for traversal
   # - AST to algebra converter
   # - Algebra validator
   #
   # Enables future optimization phases by providing
   # schema-aware, rule-based query optimization.
   #
   # Tests: 100+ comprehensive tests, >85% coverage
   #
   # Co-authored-by: [Your Name] <email>
   ```

7. **Push mergeback branch**

   ```bash
   git push origin feature/sparql-phase1-mergeback
   ```

8. **Create PR to feature/sparql-antlr**
   - Title: "feat(sparql): SPARQL Algebra foundation (Phase 1)"
   - Description: Link to PHASE_1_DEVELOPER_SPECIFICATION.md
   - Request review

9. **After merge to feature/sparql-antlr**:

   ```bash
   # Switch back to working branch (don't delete yet)
   git checkout feature/sparql-phase1-dev

   # Sync with merged changes
   git pull origin feature/sparql-antlr

   # Wait for user to delete working branch
   # (user will clean up local branches after reviewing)
   ```

---

### For Phase 5 Ontology Researcher

**Working branch**: `feature/sparql-ontology-foundation`

1. **Start research**

   ```bash
   # Switch to ontology branch
   git checkout feature/sparql-ontology-foundation
   ```

2. **Investigate Dgraph code** (Days 1-3)
   - Read graphql/schema/schema.go
   - Read edgraph/server.go, worker/task.go, posting/list.go
   - Read existing SPARQL code (sparql/ast.go, sparql/translator_extended.go)
   - Take notes in markdown

3. **Do literature review** (Days 2-4)
   - W3C SPARQL Algebra spec (https://www.w3.org/TR/sparql11-query/#sparqlAlgebra)
   - RDFS spec (https://www.w3.org/TR/rdf-schema/)
   - OWL overview (https://www.w3.org/OWL/)
   - Other databases' approaches

4. **Write design document** (Days 5-8)

   ```bash
   # Create the research document
   vim sparql/ONTOLOGY_FOUNDATION_DESIGN.md
   ```

5. **Structure the document** with all 7 sections:
   - Section 1: Research Summary
   - Section 2: Ontology Model Design
   - Section 3: Storage & Loading Strategy
   - Section 4: Query Optimization with Ontologies
   - Section 5: Implementation Plan (Phase 5.1-5.4)
   - Section 6: Reuse & Integration Decisions
   - Section 7: Open Questions & Future Extensions

6. **Validate design against Phase 1-4**
   - Review ARCHITECTURE_SPEC.md for consistency
   - Ensure Phase 5 integrates with algebra foundation
   - Check that Phase 1-4 completion doesn't block Phase 5

7. **Commit research document**

   ```bash
   git add sparql/ONTOLOGY_FOUNDATION_DESIGN.md
   git commit -m "docs(sparql): Phase 5 ontology foundation design

   Comprehensive research and design specification for SPARQL ontology
   support in Dgraph. Covers:

   - OWL/RDFS storage and loading strategy
   - Type inference and reasoning
   - Query optimization with ontology awareness
   - Integration with Phase 1-4 algebra foundation
   - Implementation plan for Phase 5.1-5.4
   - Design decisions with rationale

   Ref: PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md

   Total research: 8-10 business days
   Document size: ~3,500-5,000 words"
   ```

8. **When research complete**:

   ```bash
   # Create mergeback branch
   git checkout -b feature/sparql-ontology-mergeback feature/sparql-ontology-foundation

   # View commits since feature/sparql-antlr
   git log --oneline feature/sparql-antlr..feature/sparql-ontology-mergeback

   # Interactive rebase to clean up commits
   git rebase -i feature/sparql-antlr

   # Squash into single commit with comprehensive message
   ```

9. **Push mergeback branch**

   ```bash
   git push origin feature/sparql-ontology-mergeback
   ```

10. **Create PR to feature/sparql-antlr**
    - Title: "docs(sparql): Phase 5 ontology foundation research"
    - Description: Summary of research findings
    - Note: This is research/design, not implementation yet

11. **After merge**:

    ```bash
    # Switch back to working branch
    git checkout feature/sparql-ontology-foundation

    # Sync with merged changes
    git pull origin feature/sparql-antlr

    # Ready to begin Phase 5 implementation in next cycle
    ```

---

## Important Workflow Rules

### Commit Message Format

All commits should follow conventional commits:

```
type(scope): short summary

Detailed explanation of what changed and why.
Can span multiple lines.

Ref: [filename] - reference related documentation
Related-To: [issue/task]
Co-authored-by: Name <email>
```

**Types**:

- `feat`: New feature or functionality
- `fix`: Bug fix
- `docs`: Documentation only
- `refactor`: Code reorganization without behavior change
- `test`: Test additions or fixes
- `perf`: Performance improvements
- `build`: Build system changes

**Scopes**:

- `sparql`: General SPARQL code
- `sparql(algebra)`: Algebra system specifically
- `sparql(parser)`: Parser changes
- etc.

**Example**:

```
feat(sparql): Add algebra visitor pattern

Implements visitor pattern for traversing SPARQL algebra expressions.
Enables separation of concerns between representation and operations.

- AlgebraVisitor interface with 11 visit methods
- AlgebraPrinter concrete visitor for debugging
- VariableCollector for semantic analysis

Tests: 20+ visitor tests
Ref: PHASE_1_DEVELOPER_SPECIFICATION.md
```

### When to Commit

**Do commit**:

- When a logical unit of work is complete
- Before starting a new major piece of work
- When all tests pass for that piece
- When code is reviewed and approved

**Don't commit**:

- Incomplete implementations
- Failing tests
- Debugging prints or temporary code
- Large refactorings mixed with new features

### Code Review Before Merge

**Mergeback branch PR process**:

1. Create PR from `feature/sparql-phase1-mergeback` → `feature/sparql-antlr`
2. Link to specification document (PHASE_1_DEVELOPER_SPECIFICATION.md)
3. Include test results:
   ```
   ✅ All tests pass: go test ./sparql -v
   ✅ Coverage >85%: go test ./sparql -cover
   ✅ No linter errors: golangci-lint ./sparql/...
   ✅ Integrates with existing SPARQL code
   ```
4. Get approval from project lead
5. Merge with "Squash and merge" option (GitHub)

---

## Troubleshooting

### "I messed up a commit, how do I fix it?"

**Option 1: Amend last commit** (if not pushed)

```bash
git add your_changes
git commit --amend
# Edit message if needed
```

**Option 2: Revert commit** (if already pushed)

```bash
git revert <commit_hash>
# Creates new commit that undoes changes
```

**Option 3: Hard reset** (careful! only if not pushed)

```bash
git reset --hard HEAD~1
# Goes back one commit
```

### "I committed to wrong branch"

```bash
# Save the commit
git log --oneline  # Note the hash
git reset --soft HEAD~1  # Undo commit but keep changes

# Switch to correct branch
git checkout feature/sparql-phase1-dev
git commit -m "your message"
```

### "Mergeback branch has conflicts"

```bash
# Update mergeback branch with latest feature/sparql-antlr
git checkout feature/sparql-phase1-mergeback
git rebase feature/sparql-antlr
# Resolve conflicts in editor
git add .
git rebase --continue
```

---

## Timeline Summary

| Milestone                     | Phase 1 Dev | Ontology Dev     | Status      |
| ----------------------------- | ----------- | ---------------- | ----------- |
| Start Work                    | Day 1       | Day 1            | Parallel    |
| Step 1 Complete               | Day 4       | Mid-research     | -           |
| Step 2 Complete               | Day 7       | Late-research    | -           |
| Step 3 Complete               | Day 12      | Done (mergeback) | -           |
| Step 4 Complete               | Day 15      | -                | -           |
| Code Review                   | Day 16-17   | Day 10           | -           |
| Merge to feature/sparql-antlr | Day 17      | Day 11           | Both merged |

**Can proceed in parallel** - no blocking dependencies.

---

## After Merge: Next Phases

### Phase 2: Algebra Optimization (feature/sparql-phase2-optimization)

- Filter pushdown
- Join reordering based on cardinality
- OPTIONAL simplification

### Phase 3: Schema Integration (feature/sparql-phase3-schema)

- Type inference from schema
- Direct DQL construction (skip string re-parsing)

### Phase 4: Authorization (feature/sparql-phase4-auth)

- RBAC rule injection
- Type-level permission checking

### Phase 5: Ontology Implementation (feature/sparql-phase5-ontology)

- Phase 5.1: Ontology storage
- Phase 5.2: Type inference
- Phase 5.3: Query integration
- Phase 5.4: SPARQL reasoner

Each phase follows the same workflow:

1. Create working branch from feature/sparql-antlr
2. Do work on working branch
3. Create mergeback branch
4. Squash and merge to feature/sparql-antlr

---

## Quick Reference

```bash
# View current branch
git branch

# Switch to Phase 1 branch
git checkout feature/sparql-phase1-dev

# Switch to ontology branch
git checkout feature/sparql-ontology-foundation

# View commits on your branch
git log --oneline origin/feature/sparql-antlr..HEAD

# Sync with feature/sparql-antlr
git pull origin feature/sparql-antlr

# Create mergeback branch
git checkout -b feature/sparql-phase1-mergeback feature/sparql-phase1-dev

# Push your work
git push origin feature/sparql-phase1-mergeback

# Rebase to clean up commits
git rebase -i feature/sparql-antlr
```

---

## Questions?

- **About git**: Ask for specific command examples
- **About workflow**: See this document (GIT_WORKFLOW.md)
- **About implementation**: See PHASE_1_DEVELOPER_SPECIFICATION.md
- **About research**: See PHASE_5_ONTOLOGY_RESEARCHER_SPECIFICATION.md
- **About overall project**: See ARCHITECTURE_SPEC.md
