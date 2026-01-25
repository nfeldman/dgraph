# Phase 5 SPARQL Ontology Foundation: Researcher Specification

**Status**: Ready for Research & Design  
**Branch**: `feature/sparql-ontology-foundation`  
**Timeline**: 2-3 weeks research + design  
**Scope**: Foundation architecture for ontology support

---

## Research Assignment

Your task is to design how Dgraph will support SPARQL semantic web ontologies (OWL, RDFS). This is
Phase 5 of the SPARQL Algebra project, but can be researched independently while Phase 1-4
implementation happens in parallel.

---

## Deliverables (Due at End of Research Phase)

Produce a single comprehensive markdown document: `sparql/ONTOLOGY_FOUNDATION_DESIGN.md` with these
sections:

### Section 1: Research Summary (500-800 words)

Summarize your findings from investigating:

- How dgraph currently stores schema information (graphql/schema/schema.go)
- How dgraph handles type inference and validation
- How dgraph manages predicates and their properties
- What would be needed to layer ontology reasoning on top

Include:

- Current limitations for ontology support
- How other graph databases handle ontologies (Neo4j, Amazon Neptune, etc.)
- Why a hybrid storage approach is needed

### Section 2: Ontology Model Design (800-1000 words)

Design the semantic model for ontology support:

**Specify**:

- How OWL classes map to Dgraph types
- How OWL properties map to Dgraph predicates
- How RDFS subclass relationships are represented
- How class restrictions are stored
- How property ranges and domains are stored
- Type inference rules (if X:subClassOf Y, instances of X are also instances of Y)

**Design considerations**:

- Storage efficiency (don't duplicate data)
- Query efficiency (can we efficiently answer "is X instance of Y with inheritance?")
- Update efficiency (adding ontology statement shouldn't require recomputing everything)

**Example mappings**:

```
OWL: person:subClassOf agent:
Dgraph: Predicate `rdf:subClassOf` with domain Agent, range Agent

OWL: Person has subclass Student:
Dgraph: Triple: Student rdf:subClassOf Person

OWL: property domain/range:
Dgraph: Predicates `rdfs:domain` and `rdfs:range` with type metadata
```

### Section 3: Storage & Loading Strategy (600-800 words)

**Design for these scenarios**:

1. **Loading OWL/RDFS Files**
   - File formats: RDF/XML, Turtle, N-Triples (which to support first?)
   - Parser approach: reuse existing RDF parser or use Go library?
   - Storage: how to store ontology triples vs. instance triples?

2. **Persistent Storage**
   - Should ontologies be stored in dgraph itself (in predicates)?
   - Should they be cached in-memory at startup?
   - Should they be stored separately from instance data?
   - Recommendation: hybrid approach (explain rationale)

3. **In-Memory Caching**
   - What subset of ontology should be cached?
   - How is cache invalidated when ontology updates?
   - Memory budget considerations

4. **Query-Time Application**
   - When should ontology rules apply (at query compile-time or execution time)?
   - How do we track "instance of X (including inherited types)"?
   - Caching of inference results

### Section 4: Query Optimization with Ontologies (500-700 words)

**Research how ontologies improve query planning**:

1. **Type Narrowing**
   - If query asks for rdf:type Person, should include Students?
   - Should optimizer expand to include all subclasses?
   - Impact on query selectivity estimates

2. **Join Ordering**
   - How does class hierarchy affect cardinality estimation?
   - Should we reweight join costs based on inheritance?

3. **Filter Simplification**
   - Can we simplify filters using class restrictions?
   - Example: if domain of worksAt is Person, remove meaningless filters

4. **Property-Based Reasoning**
   - How to use property domains/ranges in optimization?
   - Example: if variable ?x must have property worksAt, infer ?x is Person

**Integration with existing optimizer** (worker/task.go, query/query.go):

- Minimal changes needed?
- Where would ontology-aware cardinality estimation go?
- How do we hook into query planning?

### Section 5: Implementation Plan (700-900 words)

Create a concrete plan for implementing Phase 5:

**Phase 5.1: Ontology Storage (Week 1-2)**

- [ ] Define OWL/RDFS predicate types
- [ ] Create ontology loading API
- [ ] Implement RDF parser (or choose library)
- [ ] Store loaded ontology in Dgraph

**Phase 5.2: Type Inference (Week 3-4)**

- [ ] Implement inheritance reasoning (subClassOf)
- [ ] Add domain/range reasoning
- [ ] Build ontology cache
- [ ] Implement "instance of with inheritance" query operator

**Phase 5.3: Query Integration (Week 5-6)**

- [ ] Hook ontology inference into SPARQL algebra compiler
- [ ] Modify query planner to use ontology cardinality
- [ ] Implement ontology-aware type checking
- [ ] Optimize queries based on class hierarchy

**Phase 5.4: SPARQL Reasoner (Week 7-8)**

- [ ] Implement SPARQL semantic entailment checking
- [ ] Add rule-based inference for common ontology patterns
- [ ] Optimize reasoning with caching

**For each sub-phase, specify**:

- Files to create/modify
- Function signatures
- Integration points
- Testing approach
- Success criteria

### Section 6: Reuse & Integration Decisions (400-600 words)

**Investigate and recommend reusing**:

1. **Dgraph Components**:
   - Can we reuse graphql/schema/schema.go infrastructure?
   - Should we extend or wrap the schema system?
   - How to integrate with existing type system?
   - Can we leverage posting/list.go for ontology index?

2. **External Libraries**:
   - Go RDF parsers: which is most suitable? (google/flatbuffers, rdf library, etc.)
   - OWL/RDFS reasoning: use existing Go library or implement?
   - Graph algorithms: leverage existing code?

3. **W3C Standards**:
   - SPARQL semantics: https://www.w3.org/TR/sparql11-semantics/
   - OWL: https://www.w3.org/OWL/
   - RDFS: https://www.w3.org/TR/rdf-schema/
   - Which features to implement (full OWL is complex)?

**Recommendation matrix**: | Component | Option A | Option B | Recommended | Rationale |
|-----------|----------|----------|-------------|-----------| | RDF Parser | Write from scratch |
Use Go library | ? | ? | | Schema reuse | Extend schema.go | Create new ontology/ | ? | ? | | etc. |
| | | |

### Section 7: Open Questions & Future Extensions (300-500 words)

**Questions to answer**:

1. How much OWL reasoning is actually needed?
   - Full OWL is undecidable (complex reasoning required)
   - RDFS is tractable (simple hierarchy)
   - What subset should Phase 5 cover?

2. When should reasoning happen?
   - At query compile time (static)
   - At query execution time (dynamic)
   - Pre-computed (materialization)
   - Trade-offs?

3. Ontology versioning?
   - Can we have multiple ontology versions?
   - How to handle ontology updates in production?

4. Performance expectations?
   - What queries should complete in <100ms with ontology?
   - How large ontologies should be supported?
   - Caching strategy to meet performance goals?

**Future extensions**:

- [ ] SHACL validation (shape constraints)
- [ ] SPARQL 1.2 entailment regimes
- [ ] Custom inference rules
- [ ] Ontology explanation (why is X instance of Y?)
- [ ] Ontology debugging tools

---

## Research Methodology

1. **Code Investigation** (Days 1-3)
   - Read graphql/schema/schema.go (how types are stored)
   - Read edgraph/server.go (how queries are processed)
   - Read worker/task.go (query planning)
   - Read posting/list.go (index structure)
   - Identify reusable abstractions

2. **Literature Review** (Days 2-4)
   - W3C SPARQL Algebra spec (you already have context)
   - RDFS semantics (simple subset)
   - OWL overview (understand complexity)
   - How other databases handle (skim documentation)

3. **Design Work** (Days 5-8)
   - Sketch ontology storage approaches
   - Design type inference algorithm
   - Map to existing Dgraph concepts
   - Create implementation plan

4. **Validation** (Days 8-9)
   - Review against Phase 1-4 architecture
   - Check for conflicts with existing code
   - Identify integration points
   - Validate assumptions with examples

---

## Success Criteria

Your research is complete when the markdown document demonstrates:

1. ✅ **Deep Understanding**
   - Shows you've read relevant Dgraph code
   - References specific types and functions
   - Makes design decisions based on codebase patterns

2. ✅ **Concrete Design**
   - Specific file names for new code
   - Function signatures with realistic parameters
   - Data structures for ontology storage
   - Clear algorithms for reasoning

3. ✅ **Realistic Implementation Plan**
   - Phased approach (5.1, 5.2, 5.3, 5.4)
   - Estimated effort for each phase
   - Clear success criteria
   - Integration tests specified

4. ✅ **Well-Reasoned Decisions**
   - Explains trade-offs considered
   - Justifies library choices
   - Addresses performance concerns
   - Considers Dgraph's architecture patterns

5. ✅ **Integrated with Phase 1-4**
   - Shows how Phase 5 connects to algebra
   - Minimal modifications to existing code
   - Leverages algebra infrastructure
   - Compatible with query planning

6. ✅ **Standards-Compliant**
   - References W3C specs appropriately
   - Explains which OWL/RDFS features are in scope
   - Clarifies why some features are out of scope

7. ✅ **Professional Quality**
   - Well-organized sections
   - Clear writing
   - Proper markdown formatting
   - Examples and diagrams where helpful

---

## Key Files to Investigate

```
graphql/schema/schema.go          # Type and predicate storage
graphql/resolve/query_rewriter.go # Query optimization reference
edgraph/server.go                 # Query entry point
edgraph/access.go                 # How RBAC works (pattern for ontology rules)
worker/task.go                    # Query planning and cardinality
posting/list.go                   # Index and storage abstractions
dql/parser.go                      # DQL intermediate representation
```

---

## Helpful Resources

**Included in workspace**:

- ARCHITECTURE_SPEC.md - Full Phase 1-5 overview
- PHASE_1_GETTING_STARTED.md - How algebra will work
- SPARQL_ALGEBRA_TODO.md - Task breakdown for all phases

**W3C Standards**:

- SPARQL Semantics: https://www.w3.org/TR/sparql11-semantics/
- SPARQL Algebra: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
- OWL 2 Specification: https://www.w3.org/TR/owl2-overview/
- RDFS: https://www.w3.org/TR/rdf-schema/

**Other Resources**:

- Dgraph schema documentation: See docs/ folder
- Similar implementations: Neo4j APOC, Amazon Neptune, Virtuoso

---

## Timeline

**Total research time**: 8-10 business days for experienced developer

- **Days 1-3**: Code investigation
- **Days 2-4**: Literature review (overlaps with code investigation)
- **Days 5-8**: Design and documentation
- **Days 8-10**: Validation and refinement

Can overlap with Phase 1-4 implementation on feature/sparql-antlr.

---

## After Research: Implementation

Once this design document is complete and reviewed:

1. Create `sparql/ontology/` directory structure
2. Begin implementing Phase 5.1 (ontology storage)
3. Each sub-phase creates pull request to `feature/sparql-ontology-mergeback`
4. When complete, squash and merge to `feature/sparql-antlr`

See ARCHITECTURE_SPEC.md for full timeline and context.

---

## Questions?

If you need clarification:

1. About Dgraph architecture → Read SYSTEM_INTEGRATION.md
2. About Phase 1-4 context → Read ARCHITECTURE_SPEC.md
3. About implementation details → Read PHASE_1_DEVELOPER_SPECIFICATION.md
4. About W3C standards → See linked resources above

Good luck! This is cutting-edge work that will make Dgraph the first graph database with true
SPARQL + ontology reasoning.
