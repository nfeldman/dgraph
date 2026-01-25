# Phase 5: Ontology Implementation - Quick Reference Guide

**Purpose**: Developer-focused guide for implementing Phase 5  
**Length**: ~30-45 minutes to read  
**Best for**: Active implementation reference

---

## TL;DR - Start Here

### What You're Building

A **complete ontology system** that:

1. **Loads** OWL/RDFS ontologies from Turtle files
2. **Reasons** using RDFS inference rules (11 core rules)
3. **Synthesizes** Dgraph schema from ontology structure
4. **Integrates** with SPARQL query compilation for optimization

### Why It Matters

- SPARQL queries can reference classes/properties not explicitly in schema
- Ontology provides semantic understanding of relationships
- Enables automatic type expansion and property inference
- Foundation for future advanced reasoning

### 3-Minute Architecture

```
OWL/RDFS File (Turtle format)
    ↓
TurtleParser → [Subject, Predicate, Object] triples
    ↓
BuildOntology → Ontology structure
    ├─ ClassHierarchy (classes + subclass relationships)
    ├─ PropertyHierarchy (properties + subproperty relationships)
    ├─ Equivalences (owl:equivalentClass, owl:equivalentProperty)
    ├─ Inverses (owl:inverseOf bidirectional mapping)
    └─ Constraints (domain, range, disjointness)
    ↓
ComputeTransitiveClosures → Precomputed all inferences
    ├─ SubClassOf[C1][C2] = true if C2 is ancestor of C1
    └─ SubPropertyOf[P1][P2] = true if P2 is ancestor of P1
    ↓
SynthesizeSchema → Dgraph type definitions
    ├─ Create types from classes
    ├─ Create predicates from properties
    ├─ Map inheritance (rdfs:subClassOf → type composition)
    └─ Map domain/range to field constraints
    ↓
Integration Point → Use in SPARQL compilation
    ├─ Expand class filters with subclasses
    ├─ Infer types from property domain/range
    └─ Add cardinality hints to optimizer
```

---

## Quick Reference: File Structure

### 6 Go Files to Create

```
sparql/ontology/
├── model.go (250 LOC)
│   - Ontology struct
│   - ClassHierarchy & PropertyHierarchy
│   - ClassDef & PropertyDef
│
├── loader.go (300 LOC)
│   - TurtleParser struct
│   - LoadOntology() function
│   - buildOntology() converter
│
├── reasoning.go (200 LOC)
│   - ComputeTransitiveClosures()
│   - ReasoningEngine struct
│   - Rule implementations
│
├── schema_builder.go (250 LOC)
│   - SynthesizeSchema() function
│   - Type/predicate generation
│   - Inheritance mapping
│
├── store.go (100 LOC)
│   - OntologyStore (multi-tenant)
│   - Load/Get operations
│
└── [6 test files] (500+ LOC combined)
    - model_test.go
    - loader_test.go
    - reasoning_test.go
    - schema_builder_test.go
    - store_test.go
    - integration_test.go
```

---

## Implementation Sequence

### Day 1-2: Data Models

**Start**: `model.go`

```go
// Core structs to define
type Ontology struct {}
type ClassHierarchy struct {}
type PropertyHierarchy struct {}
type ClassDef struct {}
type PropertyDef struct {}
type TransitiveClosures struct {}

// Helper methods
func (o *Ontology) GetClass(uri string) *ClassDef
func (o *Ontology) GetProperty(uri string) *PropertyDef
func (o *Ontology) IsSubClassOf(c1, c2 string) bool
func (o *Ontology) IsSubPropertyOf(p1, p2 string) bool
```

**Tests**: `model_test.go`

- Construction of each struct
- Basic hierarchy operations
- Method correctness

### Day 3-4: Loading & Parsing

**Start**: `loader.go`

```go
// Main function
func LoadOntology(filePath string) (*Ontology, error)

// Parser
type TurtleParser struct {}
func (tp *TurtleParser) Parse(reader io.Reader) ([]*Triple, error)

// Converter
func buildOntology(triples []*Triple) (*Ontology, error)

// Helpers
func expandURI(prefixedURI string, prefixes map[string]string) string
func parseTriple(line string) (*Triple, error)
```

**Sample Input** (Turtle):

```turtle
@prefix ex: <http://example.com/> .
@prefix owl: <http://www.w3.org/2002/07/owl#> .

ex:Person a owl:Class .
ex:Employee rdfs:subClassOf ex:Person .
```

**Output**: Ontology structure ready for reasoning

**Tests**: `loader_test.go`

- Parse simple/complex Turtle files
- Handle prefixes correctly
- Create proper Triple structure
- Convert triples to ontology

### Day 5: Reasoning

**Start**: `reasoning.go`

```go
// Main function
func (o *Ontology) ComputeTransitiveClosures() error

// Implementation
func (o *Ontology) computeSubClassOfClosure(start, current string, closure map[string]bool)
func (o *Ontology) computeSubPropertyOfClosure(start, current string, closure map[string]bool)

// Inference helpers
func (o *Ontology) GetAllSuperClasses(c string) []string
func (o *Ontology) GetAllSubClasses(c string) []string
func (o *Ontology) GetAllSuperProperties(p string) []string
```

**Algorithm**: Depth-first transitive closure computation (Warshall's)

**Tests**: `reasoning_test.go`

- Verify closure computation is correct
- Test class hierarchy expansion
- Test property hierarchy expansion
- Check for cycles (should be OK or detected)

### Day 6: Schema Synthesis

**Start**: `schema_builder.go`

```go
// Main function
func (o *Ontology) SynthesizeSchema() (*DgraphSchema, error)

// Implementation phases:
// Phase 1: Create types from classes
// Phase 2: Create predicates from properties
// Phase 3: Add fields to types (domain/range)
// Phase 4: Handle inheritance (subClassOf)

// Result
type DgraphSchema struct {
    Types map[string]*TypeDef
    Predicates map[string]*PredicateDef
}

type TypeDef struct {
    Name string
    Fields map[string]*FieldDef
    Implements []string // parent types
}
```

**Sample Output** (DQL):

```dql
type Person {
    name: string
    age: int
}

type Manager implements Employee {
    manages: [Employee] @reverse("managedBy")
}
```

**Tests**: `schema_builder_test.go`

- Basic class → type mapping
- Property → predicate mapping
- Inheritance mapping
- Domain/range constraints

### Day 7: Store & Integration

**Start**: `store.go`

```go
// Multi-tenant ontology management
type OntologyStore struct {
    ontologies map[uint64]*Ontology
    mu sync.RWMutex
}

func (os *OntologyStore) Load(namespaceID uint64, filePath string) error
func (os *OntologyStore) Get(namespaceID uint64) *Ontology
func (os *OntologyStore) Unload(namespaceID uint64)
```

**Tests**: `store_test.go`

- Load multiple ontologies
- Namespace isolation
- Thread safety

### Day 8: Integration Tests

**Start**: `integration_test.go`

```go
// End-to-end tests
func TestLoadParseReason(t *testing.T) {}
func TestOntologyToSchema(t *testing.T) {}
func TestComplexOntology(t *testing.T) {}
```

---

## RDFS Reasoning Rules (Quick Reference)

### Core 11 Rules (Memorize These)

| #   | Rule                     | Formula                                                                                      | Example                                         |
| --- | ------------------------ | -------------------------------------------------------------------------------------------- | ----------------------------------------------- | --- |
| 1   | Class Extension          | `(S, rdf:type, C1) ∧ (C1, rdfs:subClassOf, C2) → (S, rdf:type, C2)`                          | Employee is a Person                            |
| 2   | Subclass Reflexivity     | `(C, rdf:type, rdfs:Class) → (C, rdfs:subClassOf, C)`                                        | Class subclasses itself                         |
| 3   | Subclass Transitivity    | `(C1, rdfs:subClassOf, C2) ∧ (C2, rdfs:subClassOf, C3) → (C1, rdfs:subClassOf, C3)`          | Manager is an Employee is a Person              |
| 4   | Property Domain          | `(P, rdfs:domain, C) ∧ (X, P, Y) → (X, rdf:type, C)`                                         | Manager has manages → Manager is type           |
| 5   | Property Range           | `(P, rdfs:range, C) ∧ (X, P, Y) → (Y, rdf:type, C)`                                          | manages has range Employee → object is Employee |
| 6   | Property Inheritance     | `(P1, rdfs:subPropertyOf, P2) ∧ (X, P1, Y) → (X, P2, Y)`                                     | supervises is subprop of manages                |
| 7   | Subproperty Reflexivity  | `(P, rdf:type, rdfs:Property) → (P, rdfs:subPropertyOf, P)`                                  | Property subproperties itself                   |
| 8   | Subproperty Transitivity | `(P1, rdfs:subPropertyOf, P2) ∧ (P2, rdfs:subPropertyOf, P3) → (P1, rdfs:subPropertyOf, P3)` | Chains of properties                            |
| 9   | Class Type Inference     | `(C, rdfs:subClassOf, D) → (C, rdf:type, rdfs:Class)`                                        | Subclass declarations imply class               |
| 10  | Property Type Inference  | `(P, rdfs:subPropertyOf, Q) → (P, rdf:type, rdfs:Property)`                                  | Subproperty declarations                        |
| 11  | Resource Typing          | Rarely used; skip for Phase 5                                                                | N/A                                             | N/A |

### For Phase 5, Focus On

- **Rules 1, 3, 6, 8** - These are most important (transitivity, inheritance)
- **Rules 2, 7** - Trivial (reflexivity) - handle at initialization
- **Rules 4, 5** - Domain/range; used in schema synthesis

---

## Critical Implementation Details

### Turtle Parser Key Points

```go
// 1. Handle prefixes
@prefix ex: <http://example.com/> .
// → map["ex"] = "http://example.com/"

// 2. Expand prefixed URIs
ex:Person → "http://example.com/Person"
<http://example.com/Person> → "http://example.com/Person"

// 3. Handle literals
"string literal" → object is string
"30"^^xsd:integer → object is typed

// 4. Multi-line statements (semicolons, commas)
ex:john
    rdf:type ex:Person ;    # semicolon = new property
    ex:name "John" ;
    ex:age "30" .           # period = end of subject

// 5. Blank nodes
[ rdf:type ex:Person ] → Generate synthetic URI
```

### Transitive Closure Algorithm

**Use Depth-First Search (DFS):**

```
function computeClosure(start, current, visited, closure):
    visited[current] = true
    closure[current] = true

    for each parent of current:
        if parent not in visited:
            computeClosure(start, parent, visited, closure)
```

**Why DFS?**

- Simple to understand
- Natural recursive structure
- Works well for ontology graphs (usually sparse)
- Total: O(V + E) per node, O(V²) to compute all

### Schema Synthesis Key Mapping

```
OWL/RDFS          →  Dgraph Schema
─────────────────────────────────
owl:Class         →  type TypeName
owl:ObjectProperty → predicate: uid
owl:DatatypeProperty → predicate: scalar (string/int/etc)
rdfs:subClassOf   →  type Child implements Parent
rdfs:domain       →  Field belongs to Type
rdfs:range        →  Field value type
owl:inverseOf     →  @reverse("inverse_name")
```

---

## Common Patterns & Code Snippets

### Pattern 1: Class Hierarchy Traversal

```go
// Find all instances of a class and its subclasses
func findAllInstancesOfClassAndSubclasses(className string) {
    var results []string

    // Add the class itself
    results = append(results, className)

    // Add all subclasses
    if classDef, ok := onto.ClassHierarchy.AllClasses[className]; ok {
        results = append(results, classDef.AllSubClasses...)
    }

    return results
}
```

### Pattern 2: Property Chain Following

```go
// Follow a chain of subproperties
func findPropertyChain(prop string) []string {
    var chain []string

    current := prop
    for current != "" {
        chain = append(chain, current)
        superProps := onto.PropertyHierarchy.SuperProperties[current]
        if len(superProps) == 0 {
            break
        }
        current = superProps[0] // Usually just one super
    }

    return chain
}
```

### Pattern 3: Domain/Range Inference

```go
// What types can have property P?
func getPropertyDomainClasses(propURI string) []string {
    if propDef, ok := onto.PropertyHierarchy.AllProperties[propURI]; ok {
        var results []string
        for _, domain := range propDef.Domain {
            results = append(results, domain)
            // Also add subclasses of domain
            if classDef, ok := onto.ClassHierarchy.AllClasses[domain]; ok {
                results = append(results, classDef.AllSubClasses...)
            }
        }
        return results
    }
    return nil
}
```

---

## Testing Checklist

### Unit Tests (Per-Module)

- [ ] Model construction (structs, fields)
- [ ] Hierarchy operations (add relationships)
- [ ] Turtle parsing (prefixes, literals, multi-line)
- [ ] Triple conversion (correct structure)
- [ ] Transitive closure (correctness, no infinite loops)
- [ ] Schema synthesis (correct type mapping)
- [ ] Store operations (load, get, isolation)

### Integration Tests

- [ ] Load real OWL ontology → parse → reason
- [ ] Verify closure computation matches manual calculation
- [ ] Schema synthesis produces valid DQL
- [ ] Multi-tenant isolation (different namespaces)
- [ ] Error handling (invalid files, circular refs)

### Edge Cases

- [ ] Empty ontology
- [ ] Ontology with only classes, no properties
- [ ] Circular subclass relationships (should handle gracefully)
- [ ] Properties with no domain/range
- [ ] Equivalent classes/properties
- [ ] Disjoint classes
- [ ] Very large ontology (1000+ classes)

---

## Debugging Tips

### Common Issues & Solutions

| Issue                           | Cause                                | Fix                        |
| ------------------------------- | ------------------------------------ | -------------------------- |
| Parsing fails                   | Wrong prefix expansion               | Check @prefix declarations |
| Closure never completes         | Circular subclass relationships      | Add cycle detection        |
| Type mismatch                   | URI vs string confusion              | Always use full URIs       |
| Schema synthesis missing fields | Domain/range not found               | Check property definitions |
| Multi-tenant collision          | Same ontology in multiple namespaces | Use namespace ID in key    |

### Debug Output

```go
// Print class hierarchy
func (o *Ontology) DebugPrintHierarchy() {
    for className, classDef := range o.ClassHierarchy.AllClasses {
        fmt.Printf("Class: %s\n", className)
        fmt.Printf("  Direct parents: %v\n", classDef.DirectSuperClasses)
        fmt.Printf("  All parents: %v\n", classDef.AllSuperClasses)
        fmt.Printf("  Direct children: %v\n", classDef.DirectSubClasses)
    }
}

// Verify closure computation
func (o *Ontology) VerifyClosureSanity() error {
    for c1 := range o.ClassHierarchy.AllClasses {
        for c2 := range o.TransitiveClosures.SubClassOf[c1] {
            // c2 should be ancestor of c1
            if !o.isAncestor(c1, c2) {
                return fmt.Errorf("closure inconsistency: %s not ancestor of %s", c2, c1)
            }
        }
    }
    return nil
}
```

---

## Performance Targets

### Acceptable Performance

| Operation                          | Target    | How to Achieve         |
| ---------------------------------- | --------- | ---------------------- |
| Load 1000-triple ontology          | <500ms    | Optimize parser, cache |
| Type lookup                        | <1ms      | Hash map O(1)          |
| Subclass query                     | <0.1ms    | Precomputed closure    |
| Schema synthesis                   | <100ms    | Single pass            |
| Full pipeline (load→reason→schema) | <1 second | All of above           |

### Benchmarks to Write

```go
func BenchmarkLoadOntology(b *testing.B) {
    // Load ontology from file
    for i := 0; i < b.N; i++ {
        LoadOntology("test_ontology.ttl")
    }
}

func BenchmarkTransitiveClosure(b *testing.B) {
    onto := loadTestOntology()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        onto.ComputeTransitiveClosures()
    }
}

func BenchmarkSchemaS synthesis(b *testing.B) {
    onto := loadTestOntology()
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        onto.SynthesizeSchema()
    }
}
```

---

## Integration Points

### Where Ontology Plugs Into SPARQL Pipeline

**1. Query Planning** (in compiler_dql.go)

```go
// When converting algebra to DQL:
// If query filters on class C, expand to C ∪ subclasses(C)
expandedClasses := onto.GetAllSubClasses(filterClass)
```

**2. Type Inference** (in compiler_dql.go)

```go
// When property P appears, infer types from domain/range:
domain := onto.GetPropertyDomain(P)
range_ := onto.GetPropertyRange(P)
```

**3. Optimization** (in optimizer)

```go
// Use ontology cardinality hints for join ordering:
card := onto.GetPropertyCardinality(P)
```

---

## Success Criteria for Phase 5

### Code Quality

- [ ] No panics (proper error handling)
- [ ] All functions documented with comments
- [ ] Error messages are clear and actionable
- [ ] Follows Dgraph code conventions
- [ ] <10ms for small ontology load

### Functionality

- [ ] All 6 files implemented
- [ ] RDFS reasoning working correctly
- [ ] Schema synthesis produces valid DQL
- [ ] Multi-tenant isolation works
- [ ] Integrates with SPARQL pipeline

### Testing

- [ ] 200+ test cases total
- [ ] All tests passing
- [ ] > 85% code coverage
- [ ] Real OWL file tested
- [ ] No regressions in existing SPARQL tests

---

## Resources & References

### Go Patterns Used

- **Visitor Pattern** (for ontology traversal)
- **Builder Pattern** (for Ontology construction)
- **Store Pattern** (for OntologyStore)
- **DFS algorithm** (for transitive closure)

### Files to Study First

- `/Users/nfeldman/repos/dgraph/schema/schema.go` (Dgraph schema model)
- `/Users/nfeldman/repos/dgraph/sparql/algebra.go` (SPARQL algebra types)
- `/Users/nfeldman/repos/dgraph/sparql/translator_extended.go` (SPARQL compilation)

### External Resources

- [W3C OWL Specification](https://www.w3.org/TR/owl2-overview/)
- [RDFS Semantics](https://www.w3.org/TR/rdf-schema/)
- [Turtle Format](https://www.w3.org/TR/turtle/)

---

## Next Steps

1. **Read PHASE_5_ONTOLOGY_IMPLEMENTATION_RESEARCH.md** (full details)
2. **Start with model.go** (Day 1-2)
3. **Follow file structure sequentially** (Days 3-8)
4. **Write tests alongside code** (continuous)
5. **Integration test at end** (Day 8)

---

**Status**: Ready to implement  
**Estimated Duration**: 10-12 days for one developer  
**Complexity**: Medium (new concepts but clear structure)

**Good luck! 🚀**
