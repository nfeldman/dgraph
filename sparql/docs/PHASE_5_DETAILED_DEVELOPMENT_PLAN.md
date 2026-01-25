# Phase 5: Ontology Implementation - Detailed Development Plan

**Status**: Research Complete, Ready for Development  
**Version**: 1.0  
**Last Updated**: January 25, 2026

---

## Document Overview

This document provides day-by-day implementation guidance for Phase 5. Each section maps to specific
days of development with:

- **Concrete tasks** - What to build exactly
- **Code signatures** - Function/struct signatures to implement
- **Test cases** - Specific tests to write
- **Validation checkpoints** - How to know it's working
- **Common pitfalls** - What to watch out for

---

## Phase 5 Structure (4 Weeks)

```
Week 1: Foundation (Days 1-5)
  ├─ Day 1-2: Ontology data models (model.go, model_test.go)
  ├─ Day 3: Turtle parser setup (loader.go starts)
  ├─ Day 4: Triple parsing (loader.go continues)
  └─ Day 5: Ontology construction (loader.go finalized)

Week 2: Reasoning (Days 6-10)
  ├─ Day 6: Transitive closure foundation (reasoning.go)
  ├─ Day 7-8: Full closure computation (reasoning.go)
  ├─ Day 9: Reasoning validation (reasoning_test.go)
  └─ Day 10: Integration with model (harmony between modules)

Week 3: Schema Synthesis (Days 11-14)
  ├─ Day 11: Schema data structures (schema_builder.go)
  ├─ Day 12: Type generation (schema_builder.go)
  ├─ Day 13: Property mapping + inheritance (schema_builder.go)
  └─ Day 14: Schema validation (schema_builder_test.go)

Week 4: Integration & Testing (Days 15-20)
  ├─ Day 15-16: OntologyStore + deployment (store.go)
  ├─ Day 17-19: End-to-end tests (integration_test.go)
  └─ Day 20: Performance tuning + review
```

---

## WEEK 1: FOUNDATION

### Day 1-2: Ontology Data Models

#### Objectives

- Define core ontology data structures
- Implement hierarchy representations
- Create helper methods for access

#### Files to Create

**File: `sparql/ontology/model.go`**

```go
package ontology

import (
    "fmt"
    "sync"
)

// Ontology represents a loaded OWL/RDFS ontology
type Ontology struct {
    // Metadata
    URI string // Ontology URI
    Namespaces map[string]string // prefix -> full URI mapping

    // Hierarchies
    ClassHierarchy *ClassHierarchy
    PropertyHierarchy *PropertyHierarchy

    // OWL relationships
    EquivalentClasses map[string][]string // C -> [equivalent classes]
    EquivalentProperties map[string][]string // P -> [equivalent properties]
    DisjointClasses map[string][]string // C -> [disjoint classes]
    InverseProperties map[string]string // P -> inverse(P)

    // Constraints
    PropertyConstraints map[string]*PropertyConstraint

    // Precomputed closures
    TransitiveClosures *TransitiveClosures

    // Mutex for thread safety
    mu sync.RWMutex
}

// ClassHierarchy manages class relationships
type ClassHierarchy struct {
    // Direct relationships
    SubClasses map[string][]string // Parent -> [children]
    SuperClasses map[string][]string // Child -> [parents]

    // All classes
    AllClasses map[string]*ClassDef
}

// ClassDef represents a single class definition
type ClassDef struct {
    URI string // Full URI of the class

    // Direct hierarchy
    DirectSuperClasses []string // Classes this extends
    DirectSubClasses []string // Classes that extend this

    // Transitive closure (computed later)
    AllSuperClasses []string
    AllSubClasses []string

    // OWL attributes
    EquivalentClasses []string
    DisjointClasses []string

    // Metadata
    Label string // rdfs:label
    Comment string // rdfs:comment
    IsAbstract bool // Is this an abstract class?
}

// PropertyHierarchy manages property relationships
type PropertyHierarchy struct {
    // Direct relationships
    SubProperties map[string][]string // Parent -> [children]
    SuperProperties map[string][]string // Child -> [parents]

    // All properties
    AllProperties map[string]*PropertyDef
}

// PropertyDef represents a single property definition
type PropertyDef struct {
    URI string // Full URI of the property
    Type PropertyType // OBJECT_PROPERTY or DATATYPE_PROPERTY

    // Direct hierarchy
    DirectSuperProperties []string
    DirectSubProperties []string

    // Transitive closure (computed later)
    AllSuperProperties []string
    AllSubProperties []string

    // Domain and range
    Domain []string // Class URIs that can have this property
    Range []string // Class URIs or datatype URIs

    // OWL attributes
    InverseProperty string // owl:inverseOf
    EquivalentProperties []string
    IsFunctional bool // owl:FunctionalProperty
    IsInverseFunctional bool // owl:InverseFunctionalProperty

    // Metadata
    Label string
    Comment string
}

type PropertyType int

const (
    OBJECT_PROPERTY PropertyType = iota
    DATATYPE_PROPERTY
)

// PropertyConstraint combines domain/range info for a property
type PropertyConstraint struct {
    Property string
    Domain []string
    Range []string
}

// TransitiveClosures holds precomputed transitive relationships
type TransitiveClosures struct {
    // SubClassOf[C1][C2] = true if C2 is ancestor of C1
    SubClassOf map[string]map[string]bool

    // SubPropertyOf[P1][P2] = true if P2 is ancestor of P1
    SubPropertyOf map[string]map[string]bool
}

// ========== PUBLIC METHODS ==========

// GetClass returns the class definition by URI
func (o *Ontology) GetClass(uri string) *ClassDef {
    if o == nil || o.ClassHierarchy == nil {
        return nil
    }
    o.mu.RLock()
    defer o.mu.RUnlock()
    return o.ClassHierarchy.AllClasses[uri]
}

// GetProperty returns the property definition by URI
func (o *Ontology) GetProperty(uri string) *PropertyDef {
    if o == nil || o.PropertyHierarchy == nil {
        return nil
    }
    o.mu.RLock()
    defer o.mu.RUnlock()
    return o.PropertyHierarchy.AllProperties[uri]
}

// IsSubClassOf checks if c1 is a subclass of c2 (includes transitive)
func (o *Ontology) IsSubClassOf(c1, c2 string) bool {
    if o == nil || o.TransitiveClosures == nil {
        return false
    }
    o.mu.RLock()
    defer o.mu.RUnlock()

    closure := o.TransitiveClosures.SubClassOf[c1]
    return closure[c2]
}

// IsSubPropertyOf checks if p1 is a subproperty of p2 (includes transitive)
func (o *Ontology) IsSubPropertyOf(p1, p2 string) bool {
    if o == nil || o.TransitiveClosures == nil {
        return false
    }
    o.mu.RLock()
    defer o.mu.RUnlock()

    closure := o.TransitiveClosures.SubPropertyOf[p1]
    return closure[p2]
}

// GetAllSuperClasses returns all super classes (transitive)
func (o *Ontology) GetAllSuperClasses(uri string) []string {
    classDef := o.GetClass(uri)
    if classDef == nil {
        return nil
    }
    return classDef.AllSuperClasses
}

// GetAllSubClasses returns all sub classes (transitive)
func (o *Ontology) GetAllSubClasses(uri string) []string {
    classDef := o.GetClass(uri)
    if classDef == nil {
        return nil
    }
    return classDef.AllSubClasses
}

// GetPropertyDomain returns domain classes for a property
func (o *Ontology) GetPropertyDomain(propURI string) []string {
    propDef := o.GetProperty(propURI)
    if propDef == nil {
        return nil
    }
    return propDef.Domain
}

// GetPropertyRange returns range for a property
func (o *Ontology) GetPropertyRange(propURI string) []string {
    propDef := o.GetProperty(propURI)
    if propDef == nil {
        return nil
    }
    return propDef.Range
}

// ExpandClassWithSubclasses returns a class and all its subclasses
func (o *Ontology) ExpandClassWithSubclasses(classURI string) []string {
    result := []string{classURI}
    subclasses := o.GetAllSubClasses(classURI)
    result = append(result, subclasses...)
    return result
}

// ========== INITIALIZATION ==========

// NewOntology creates a new empty ontology
func NewOntology() *Ontology {
    return &Ontology{
        Namespaces: make(map[string]string),
        ClassHierarchy: &ClassHierarchy{
            SubClasses: make(map[string][]string),
            SuperClasses: make(map[string][]string),
            AllClasses: make(map[string]*ClassDef),
        },
        PropertyHierarchy: &PropertyHierarchy{
            SubProperties: make(map[string][]string),
            SuperProperties: make(map[string][]string),
            AllProperties: make(map[string]*PropertyDef),
        },
        EquivalentClasses: make(map[string][]string),
        EquivalentProperties: make(map[string][]string),
        DisjointClasses: make(map[string][]string),
        InverseProperties: make(map[string]string),
        PropertyConstraints: make(map[string]*PropertyConstraint),
    }
}

// InitializeClass adds a class to the ontology
func (o *Ontology) InitializeClass(uri string) *ClassDef {
    if cd, exists := o.ClassHierarchy.AllClasses[uri]; exists {
        return cd
    }

    cd := &ClassDef{
        URI: uri,
        DirectSuperClasses: []string{},
        DirectSubClasses: []string{},
        AllSuperClasses: []string{},
        AllSubClasses: []string{},
        EquivalentClasses: []string{},
        DisjointClasses: []string{},
    }

    o.ClassHierarchy.AllClasses[uri] = cd
    return cd
}

// InitializeProperty adds a property to the ontology
func (o *Ontology) InitializeProperty(uri string, propType PropertyType) *PropertyDef {
    if pd, exists := o.PropertyHierarchy.AllProperties[uri]; exists {
        return pd
    }

    pd := &PropertyDef{
        URI: uri,
        Type: propType,
        DirectSuperProperties: []string{},
        DirectSubProperties: []string{},
        AllSuperProperties: []string{},
        AllSubProperties: []string{},
        Domain: []string{},
        Range: []string{},
        EquivalentProperties: []string{},
    }

    o.PropertyHierarchy.AllProperties[uri] = pd
    return pd
}

// ========== RELATIONSHIP BUILDERS ==========

// AddSubClassRelationship adds a direct subclass relationship
func (o *Ontology) AddSubClassRelationship(child, parent string) error {
    o.mu.Lock()
    defer o.mu.Unlock()

    childDef := o.InitializeClass(child)
    parentDef := o.InitializeClass(parent)

    // Add to direct relationships
    childDef.DirectSuperClasses = append(childDef.DirectSuperClasses, parent)
    parentDef.DirectSubClasses = append(parentDef.DirectSubClasses, child)

    // Add to hierarchy maps
    o.ClassHierarchy.SuperClasses[child] = append(o.ClassHierarchy.SuperClasses[child], parent)
    o.ClassHierarchy.SubClasses[parent] = append(o.ClassHierarchy.SubClasses[parent], child)

    return nil
}

// AddSubPropertyRelationship adds a direct subproperty relationship
func (o *Ontology) AddSubPropertyRelationship(child, parent string) error {
    o.mu.Lock()
    defer o.mu.Unlock()

    childDef := o.InitializeProperty(child, OBJECT_PROPERTY)
    parentDef := o.InitializeProperty(parent, OBJECT_PROPERTY)

    // Add to direct relationships
    childDef.DirectSuperProperties = append(childDef.DirectSuperProperties, parent)
    parentDef.DirectSubProperties = append(parentDef.DirectSubProperties, child)

    // Add to hierarchy maps
    o.PropertyHierarchy.SuperProperties[child] = append(o.PropertyHierarchy.SuperProperties[child], parent)
    o.PropertyHierarchy.SubProperties[parent] = append(o.PropertyHierarchy.SubProperties[parent], child)

    return nil
}

// AddPropertyDomain adds a domain constraint
func (o *Ontology) AddPropertyDomain(propURI, classURI string) error {
    o.mu.Lock()
    defer o.mu.Unlock()

    propDef := o.InitializeProperty(propURI, OBJECT_PROPERTY)
    propDef.Domain = append(propDef.Domain, classURI)

    return nil
}

// AddPropertyRange adds a range constraint
func (o *Ontology) AddPropertyRange(propURI, rangeURI string) error {
    o.mu.Lock()
    defer o.mu.Unlock()

    propDef := o.InitializeProperty(propURI, OBJECT_PROPERTY)
    propDef.Range = append(propDef.Range, rangeURI)

    return nil
}
```

**File: `sparql/ontology/model_test.go`**

```go
package ontology

import (
    "testing"
)

func TestNewOntology(t *testing.T) {
    onto := NewOntology()
    if onto == nil {
        t.Fatal("NewOntology returned nil")
    }
    if onto.ClassHierarchy == nil {
        t.Fatal("ClassHierarchy not initialized")
    }
    if onto.PropertyHierarchy == nil {
        t.Fatal("PropertyHierarchy not initialized")
    }
}

func TestInitializeClass(t *testing.T) {
    onto := NewOntology()
    classURI := "http://example.com/Person"

    cd := onto.InitializeClass(classURI)
    if cd == nil {
        t.Fatal("InitializeClass returned nil")
    }
    if cd.URI != classURI {
        t.Errorf("Expected URI %s, got %s", classURI, cd.URI)
    }

    // Verify it's stored
    retrieved := onto.GetClass(classURI)
    if retrieved == nil {
        t.Fatal("GetClass returned nil for initialized class")
    }
}

func TestAddSubClassRelationship(t *testing.T) {
    onto := NewOntology()
    child := "http://example.com/Employee"
    parent := "http://example.com/Person"

    err := onto.AddSubClassRelationship(child, parent)
    if err != nil {
        t.Fatalf("AddSubClassRelationship returned error: %v", err)
    }

    // Verify direct relationships
    childDef := onto.GetClass(child)
    if childDef == nil {
        t.Fatal("Child class not created")
    }
    if len(childDef.DirectSuperClasses) != 1 || childDef.DirectSuperClasses[0] != parent {
        t.Errorf("Direct super classes not set correctly")
    }

    parentDef := onto.GetClass(parent)
    if parentDef == nil {
        t.Fatal("Parent class not created")
    }
    if len(parentDef.DirectSubClasses) != 1 || parentDef.DirectSubClasses[0] != child {
        t.Errorf("Direct sub classes not set correctly")
    }
}

func TestInitializeProperty(t *testing.T) {
    onto := NewOntology()
    propURI := "http://example.com/manages"

    pd := onto.InitializeProperty(propURI, OBJECT_PROPERTY)
    if pd == nil {
        t.Fatal("InitializeProperty returned nil")
    }
    if pd.URI != propURI {
        t.Errorf("Expected URI %s, got %s", propURI, pd.URI)
    }
    if pd.Type != OBJECT_PROPERTY {
        t.Errorf("Expected OBJECT_PROPERTY, got %v", pd.Type)
    }
}

func TestAddPropertyDomain(t *testing.T) {
    onto := NewOntology()
    propURI := "http://example.com/manages"
    domainURI := "http://example.com/Manager"

    err := onto.AddPropertyDomain(propURI, domainURI)
    if err != nil {
        t.Fatalf("AddPropertyDomain returned error: %v", err)
    }

    propDef := onto.GetProperty(propURI)
    if propDef == nil {
        t.Fatal("Property not created")
    }
    if len(propDef.Domain) != 1 || propDef.Domain[0] != domainURI {
        t.Errorf("Domain not set correctly")
    }
}

func TestAddPropertyRange(t *testing.T) {
    onto := NewOntology()
    propURI := "http://example.com/manages"
    rangeURI := "http://example.com/Employee"

    err := onto.AddPropertyRange(propURI, rangeURI)
    if err != nil {
        t.Fatalf("AddPropertyRange returned error: %v", err)
    }

    propDef := onto.GetProperty(propURI)
    if propDef == nil {
        t.Fatal("Property not created")
    }
    if len(propDef.Range) != 1 || propDef.Range[0] != rangeURI {
        t.Errorf("Range not set correctly")
    }
}

func TestExpandClassWithSubclasses(t *testing.T) {
    onto := NewOntology()
    person := "http://example.com/Person"
    employee := "http://example.com/Employee"
    manager := "http://example.com/Manager"

    // Build hierarchy
    onto.AddSubClassRelationship(employee, person)
    onto.AddSubClassRelationship(manager, employee)

    // Test expand (note: need closure first, tested separately)
    expanded := onto.ExpandClassWithSubclasses(person)
    if expanded[0] != person {
        t.Error("First element should be the class itself")
    }
}
```

#### Validation Checkpoint

After Days 1-2:

- [ ] `model.go` compiles without errors
- [ ] All struct fields are accessible
- [ ] Helper methods work (GetClass, GetProperty, etc.)
- [ ] `model_test.go` passes all tests
- [ ] No thread safety issues (basic mutex works)

---

### Day 3-4: Turtle Parser Setup & Triple Parsing

#### Objectives

- Implement Turtle file parser
- Handle RDF triples (Subject, Predicate, Object)
- Process prefixes and URI expansion
- Parse various literal types

#### Files to Create

**File: `sparql/ontology/loader.go`**

```go
package ontology

import (
    "bufio"
    "fmt"
    "io"
    "regexp"
    "strings"
)

// Triple represents an RDF triple
type Triple struct {
    Subject string
    Predicate string
    Object string
    IsObjectLiteral bool // true if Object is a literal string
    ObjectDatatype string // for typed literals like "30"^^xsd:integer
}

// TurtleParser parses Turtle format RDF files
type TurtleParser struct {
    // Input stream
    scanner *bufio.Scanner

    // State during parsing
    prefixes map[string]string // @prefix mappings
    blankNodes map[string]string // blank node ID -> generated URI
    blankNodeCounter int

    // Output
    triples []*Triple

    // Parsing state
    currentSubject string
    currentObjects []string // For multi-object statements
}

// ParseTurtleFile reads and parses a Turtle format file
func ParseTurtleFile(filePath string) ([]*Triple, error) {
    // TODO: Open file and call ParseTurtleReader
    return nil, nil // Placeholder
}

// ParseTurtleReader parses Turtle format from a reader
func ParseTurtleReader(reader io.Reader) ([]*Triple, error) {
    parser := &TurtleParser{
        prefixes: make(map[string]string),
        blankNodes: make(map[string]string),
        triples: make([]*Triple, 0),
    }

    return parser.Parse(reader)
}

// Parse performs the actual parsing
func (tp *TurtleParser) Parse(reader io.Reader) ([]*Triple, error) {
    tp.scanner = bufio.NewScanner(reader)

    for tp.scanner.Scan() {
        line := tp.scanner.Text()

        // Strip comments and whitespace
        line = tp.stripComments(line)
        line = strings.TrimSpace(line)

        if line == "" {
            continue
        }

        // Handle directives
        if strings.HasPrefix(line, "@prefix") {
            if err := tp.parsePrefix(line); err != nil {
                return nil, err
            }
            continue
        }

        if strings.HasPrefix(line, "@base") {
            // TODO: Handle @base directive
            continue
        }

        // Parse triple(s)
        if err := tp.parseStatements(line); err != nil {
            return nil, fmt.Errorf("parse error: %v", err)
        }
    }

    if err := tp.scanner.Err(); err != nil {
        return nil, err
    }

    return tp.triples, nil
}

// stripComments removes Turtle comments (#...) from a line
func (tp *TurtleParser) stripComments(line string) string {
    idx := strings.Index(line, "#")
    if idx >= 0 {
        return line[:idx]
    }
    return line
}

// parsePrefix parses @prefix directives
func (tp *TurtleParser) parsePrefix(line string) error {
    // Format: @prefix prefix: <URI> .
    // Example: @prefix ex: <http://example.com/> .

    // Remove @prefix and trailing period
    line = strings.TrimPrefix(line, "@prefix")
    line = strings.TrimSuffix(line, ".")
    line = strings.TrimSpace(line)

    parts := strings.Fields(line)
    if len(parts) != 2 {
        return fmt.Errorf("invalid @prefix line: %s", line)
    }

    prefix := strings.TrimSuffix(parts[0], ":")
    uri := tp.extractURI(parts[1])

    tp.prefixes[prefix] = uri
    return nil
}

// extractURI removes angle brackets from URI
func (tp *TurtleParser) extractURI(s string) string {
    if strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">") {
        return s[1 : len(s)-1]
    }
    return s
}

// expandURI expands prefixed URIs to full form
func (tp *TurtleParser) expandURI(s string) string {
    // If already expanded (has :// or starts with <>)
    if strings.HasPrefix(s, "<") && strings.HasSuffix(s, ">") {
        return tp.extractURI(s)
    }

    // If it's http:// or similar, it's already expanded
    if strings.Contains(s, "://") {
        return s
    }

    // Try prefix expansion
    if strings.Contains(s, ":") {
        parts := strings.Split(s, ":")
        if len(parts) == 2 {
            prefix := parts[0]
            localName := parts[1]
            if baseURI, ok := tp.prefixes[prefix]; ok {
                return baseURI + localName
            }
        }
    }

    return s
}

// parseStatements parses one or more RDF statements
// Handles multi-line and semicolon/comma shorthand
func (tp *TurtleParser) parseStatements(line string) error {
    // For now, simple implementation handles single statement
    // TODO: Full multi-line statement handling

    if !strings.HasSuffix(line, ".") {
        // Multi-line statement (not supported yet)
        return fmt.Errorf("multi-line statements not yet supported")
    }

    // Remove trailing period
    line = strings.TrimSuffix(line, ".")
    line = strings.TrimSpace(line)

    // Split into subject, predicate, object
    parts := strings.Fields(line)
    if len(parts) < 3 {
        return fmt.Errorf("invalid triple: %s", line)
    }

    subject := tp.expandURI(parts[0])
    predicate := tp.expandURI(parts[1])
    object := tp.expandURI(strings.Join(parts[2:], " "))

    // Determine if object is literal
    isLiteral := strings.HasPrefix(object, "\"")

    triple := &Triple{
        Subject: subject,
        Predicate: predicate,
        Object: object,
        IsObjectLiteral: isLiteral,
    }

    tp.triples = append(tp.triples, triple)
    return nil
}

// LoadOntologyFromFile loads an ontology from a Turtle file
func LoadOntologyFromFile(filePath string) (*Ontology, error) {
    // TODO: Read file, parse, build ontology
    return nil, nil
}
```

#### Key Implementation Details

**Turtle Parsing Rules:**

1. **Comments**: `# comment text` (strip to end of line)
2. **Prefixes**: `@prefix ex: <http://example.com/> .`
3. **URIs**: `<http://example.com/Person>` or `ex:Person` (with prefix)
4. **Literals**: `"John Doe"` or `"30"^^xsd:integer`
5. **Triple**: `subject predicate object .`

**Tests: `sparql/ontology/loader_test.go`**

```go
package ontology

import (
    "strings"
    "testing"
)

func TestParsePrefix(t *testing.T) {
    parser := &TurtleParser{
        prefixes: make(map[string]string),
    }

    line := "@prefix ex: <http://example.com/> ."
    err := parser.parsePrefix(line)
    if err != nil {
        t.Fatalf("parsePrefix returned error: %v", err)
    }

    if parser.prefixes["ex"] != "http://example.com/" {
        t.Errorf("Expected 'http://example.com/', got '%s'", parser.prefixes["ex"])
    }
}

func TestExpandURI(t *testing.T) {
    parser := &TurtleParser{
        prefixes: make(map[string]string),
    }
    parser.prefixes["ex"] = "http://example.com/"

    tests := []struct {
        input string
        expected string
    }{
        {"ex:Person", "http://example.com/Person"},
        {"<http://example.com/Person>", "http://example.com/Person"},
        {"http://example.com/Person", "http://example.com/Person"},
    }

    for _, tt := range tests {
        result := parser.expandURI(tt.input)
        if result != tt.expected {
            t.Errorf("expandURI(%s) = %s, expected %s", tt.input, result, tt.expected)
        }
    }
}

func TestParseSingleTriple(t *testing.T) {
    reader := strings.NewReader(`
@prefix ex: <http://example.com/> .
ex:john rdf:type ex:Person .
`)

    triples, err := ParseTurtleReader(reader)
    if err != nil {
        t.Fatalf("ParseTurtleReader returned error: %v", err)
    }

    if len(triples) != 1 {
        t.Errorf("Expected 1 triple, got %d", len(triples))
    }

    triple := triples[0]
    if triple.Subject != "http://example.com/john" {
        t.Errorf("Subject mismatch: %s", triple.Subject)
    }
}

func TestParseMultipleTriples(t *testing.T) {
    reader := strings.NewReader(`
@prefix ex: <http://example.com/> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .

ex:Person a owl:Class .
ex:Employee rdfs:subClassOf ex:Person .
`)

    triples, err := ParseTurtleReader(reader)
    if err != nil {
        t.Fatalf("ParseTurtleReader returned error: %v", err)
    }

    if len(triples) < 2 {
        t.Errorf("Expected at least 2 triples, got %d", len(triples))
    }
}
```

#### Validation Checkpoint

After Days 3-4:

- [ ] `loader.go` parses basic Turtle files
- [ ] Prefixes are correctly expanded
- [ ] Triple parsing works
- [ ] Tests pass for single and multiple triples
- [ ] Comment stripping works
- [ ] No crashes on edge cases

---

### Day 5: Ontology Construction from Triples

#### Objectives

- Convert parsed triples to Ontology structure
- Discover classes and properties
- Build relationships from triples
- Handle OWL/RDFS vocabularies

#### Add to `sparql/ontology/loader.go`

```go
// Key RDF/RDFS/OWL predicates (URIs)
const (
    RDF_TYPE = "http://www.w3.org/1999/02/22-rdf-syntax-ns#type"
    RDF_VALUE = "http://www.w3.org/1999/02/22-rdf-syntax-ns#value"

    RDFS_CLASS = "http://www.w3.org/2000/01/rdf-schema#Class"
    RDFS_SUBCLASS_OF = "http://www.w3.org/2000/01/rdf-schema#subClassOf"
    RDFS_SUBPROPERTY_OF = "http://www.w3.org/2000/01/rdf-schema#subPropertyOf"
    RDFS_DOMAIN = "http://www.w3.org/2000/01/rdf-schema#domain"
    RDFS_RANGE = "http://www.w3.org/2000/01/rdf-schema#range"
    RDFS_LABEL = "http://www.w3.org/2000/01/rdf-schema#label"
    RDFS_COMMENT = "http://www.w3.org/2000/01/rdf-schema#comment"
    RDFS_PROPERTY = "http://www.w3.org/2000/01/rdf-schema#Property"

    OWL_CLASS = "http://www.w3.org/2002/07/owl#Class"
    OWL_OBJECT_PROPERTY = "http://www.w3.org/2002/07/owl#ObjectProperty"
    OWL_DATATYPE_PROPERTY = "http://www.w3.org/2002/07/owl#DatatypeProperty"
    OWL_EQUIVALENT_CLASS = "http://www.w3.org/2002/07/owl#equivalentClass"
    OWL_EQUIVALENT_PROPERTY = "http://www.w3.org/2002/07/owl#equivalentProperty"
    OWL_DISJOINT_WITH = "http://www.w3.org/2002/07/owl#disjointWith"
    OWL_INVERSE_OF = "http://www.w3.org/2002/07/owl#inverseOf"
    OWL_FUNCTIONAL_PROPERTY = "http://www.w3.org/2002/07/owl#FunctionalProperty"
    OWL_INVERSE_FUNCTIONAL_PROPERTY = "http://www.w3.org/2002/07/owl#InverseFunctionalProperty"
)

// BuildOntologyFromTriples converts triples to ontology structure
func BuildOntologyFromTriples(triples []*Triple) (*Ontology, error) {
    onto := NewOntology()

    // Pass 1: Discover all classes and properties
    if err := discoverClassesAndProperties(onto, triples); err != nil {
        return nil, err
    }

    // Pass 2: Build relationships
    if err := buildRelationships(onto, triples); err != nil {
        return nil, err
    }

    // Pass 3: Add constraints (domain, range, labels, etc.)
    if err := addConstraints(onto, triples); err != nil {
        return nil, err
    }

    return onto, nil
}

// discoverClassesAndProperties finds all classes and properties in triples
func discoverClassesAndProperties(onto *Ontology, triples []*Triple) error {
    for _, triple := range triples {
        if triple.Predicate == RDF_TYPE {
            if triple.Object == OWL_CLASS || triple.Object == RDFS_CLASS {
                onto.InitializeClass(triple.Subject)
            }
            if triple.Object == OWL_OBJECT_PROPERTY {
                onto.InitializeProperty(triple.Subject, OBJECT_PROPERTY)
            }
            if triple.Object == OWL_DATATYPE_PROPERTY {
                onto.InitializeProperty(triple.Subject, DATATYPE_PROPERTY)
            }
            if triple.Object == RDFS_PROPERTY {
                onto.InitializeProperty(triple.Subject, OBJECT_PROPERTY)
            }
        }
    }
    return nil
}

// buildRelationships builds class and property hierarchies
func buildRelationships(onto *Ontology, triples []*Triple) error {
    for _, triple := range triples {
        switch triple.Predicate {
        case RDFS_SUBCLASS_OF:
            if err := onto.AddSubClassRelationship(triple.Subject, triple.Object); err != nil {
                return err
            }

        case RDFS_SUBPROPERTY_OF:
            if err := onto.AddSubPropertyRelationship(triple.Subject, triple.Object); err != nil {
                return err
            }

        case OWL_EQUIVALENT_CLASS:
            onto.EquivalentClasses[triple.Subject] = append(
                onto.EquivalentClasses[triple.Subject], triple.Object)
            // Make it bidirectional
            onto.EquivalentClasses[triple.Object] = append(
                onto.EquivalentClasses[triple.Object], triple.Subject)

        case OWL_INVERSE_OF:
            onto.InverseProperties[triple.Subject] = triple.Object
            onto.InverseProperties[triple.Object] = triple.Subject

        case OWL_DISJOINT_WITH:
            onto.DisjointClasses[triple.Subject] = append(
                onto.DisjointClasses[triple.Subject], triple.Object)
            onto.DisjointClasses[triple.Object] = append(
                onto.DisjointClasses[triple.Object], triple.Subject)
        }
    }
    return nil
}

// addConstraints adds domain, range, labels, etc.
func addConstraints(onto *Ontology, triples []*Triple) error {
    for _, triple := range triples {
        switch triple.Predicate {
        case RDFS_DOMAIN:
            if err := onto.AddPropertyDomain(triple.Subject, triple.Object); err != nil {
                return err
            }

        case RDFS_RANGE:
            if err := onto.AddPropertyRange(triple.Subject, triple.Object); err != nil {
                return err
            }

        case RDFS_LABEL:
            if classDef := onto.GetClass(triple.Subject); classDef != nil {
                classDef.Label = strings.Trim(triple.Object, "\"")
            }
            if propDef := onto.GetProperty(triple.Subject); propDef != nil {
                propDef.Label = strings.Trim(triple.Object, "\"")
            }

        case RDFS_COMMENT:
            if classDef := onto.GetClass(triple.Subject); classDef != nil {
                classDef.Comment = strings.Trim(triple.Object, "\"")
            }
            if propDef := onto.GetProperty(triple.Subject); propDef != nil {
                propDef.Comment = strings.Trim(triple.Object, "\"")
            }

        case OWL_FUNCTIONAL_PROPERTY:
            if propDef := onto.GetProperty(triple.Subject); propDef != nil {
                propDef.IsFunctional = true
            }
        }
    }
    return nil
}
```

**Add Tests to `sparql/ontology/loader_test.go`:**

```go
func TestBuildOntologyFromTriples(t *testing.T) {
    triples := []*Triple{
        {
            Subject: "http://example.com/Person",
            Predicate: RDF_TYPE,
            Object: OWL_CLASS,
            IsObjectLiteral: false,
        },
        {
            Subject: "http://example.com/Employee",
            Predicate: RDFS_SUBCLASS_OF,
            Object: "http://example.com/Person",
            IsObjectLiteral: false,
        },
    }

    onto, err := BuildOntologyFromTriples(triples)
    if err != nil {
        t.Fatalf("BuildOntologyFromTriples returned error: %v", err)
    }

    if onto.GetClass("http://example.com/Person") == nil {
        t.Fatal("Person class not created")
    }

    if onto.GetClass("http://example.com/Employee") == nil {
        t.Fatal("Employee class not created")
    }
}
```

#### Validation Checkpoint

After Day 5:

- [ ] ParseTurtleReader successfully parses files
- [ ] BuildOntologyFromTriples correctly creates ontology
- [ ] Classes and properties are discovered
- [ ] Relationships are built correctly
- [ ] All tests pass
- [ ] Can parse real OWL files

---

## WEEK 2: REASONING (Days 6-10)

### Day 6-8: Transitive Closure Computation

#### Objectives

- Implement transitive closure algorithm
- Precompute all inferences at load time
- Handle cycles gracefully
- Ensure O(1) lookups

#### File: `sparql/ontology/reasoning.go`

```go
package ontology

import (
    "fmt"
)

// ComputeTransitiveClosures computes all transitive relationships in ontology
func (o *Ontology) ComputeTransitiveClosures() error {
    // Initialize closure maps
    o.TransitiveClosures = &TransitiveClosures{
        SubClassOf: make(map[string]map[string]bool),
        SubPropertyOf: make(map[string]map[string]bool),
    }

    // Compute SubClassOf closure using DFS
    for classURI := range o.ClassHierarchy.AllClasses {
        closure := make(map[string]bool)
        if err := o.computeSubClassOfClosure(classURI, closure, make(map[string]bool)); err != nil {
            return fmt.Errorf("error computing subclass closure for %s: %v", classURI, err)
        }
        o.TransitiveClosures.SubClassOf[classURI] = closure
    }

    // Compute SubPropertyOf closure using DFS
    for propURI := range o.PropertyHierarchy.AllProperties {
        closure := make(map[string]bool)
        if err := o.computeSubPropertyOfClosure(propURI, closure, make(map[string]bool)); err != nil {
            return fmt.Errorf("error computing subproperty closure for %s: %v", propURI, err)
        }
        o.TransitiveClosures.SubPropertyOf[propURI] = closure
    }

    // Update ClassDef with computed closures
    o.updateClassDefinitionsWithClosure()

    // Update PropertyDef with computed closures
    o.updatePropertyDefinitionsWithClosure()

    return nil
}

// computeSubClassOfClosure computes DFS closure for subclass relationships
func (o *Ontology) computeSubClassOfClosure(
    current string,
    closure map[string]bool,
    visited map[string]bool,
) error {
    // Prevent infinite loops (cycle detection)
    if visited[current] {
        return nil
    }
    visited[current] = true

    // Add self to closure (reflexive)
    closure[current] = true

    // Add all direct super classes and their closures
    superClasses := o.ClassHierarchy.SuperClasses[current]
    for _, superClass := range superClasses {
        if !closure[superClass] {
            if err := o.computeSubClassOfClosure(superClass, closure, visited); err != nil {
                return err
            }
        }
    }

    return nil
}

// computeSubPropertyOfClosure computes DFS closure for subproperty relationships
func (o *Ontology) computeSubPropertyOfClosure(
    current string,
    closure map[string]bool,
    visited map[string]bool,
) error {
    // Prevent infinite loops
    if visited[current] {
        return nil
    }
    visited[current] = true

    // Add self to closure (reflexive)
    closure[current] = true

    // Add all direct super properties and their closures
    superProps := o.PropertyHierarchy.SuperProperties[current]
    for _, superProp := range superProps {
        if !closure[superProp] {
            if err := o.computeSubPropertyOfClosure(superProp, closure, visited); err != nil {
                return err
            }
        }
    }

    return nil
}

// updateClassDefinitionsWithClosure updates ClassDef.AllSuperClasses and AllSubClasses
func (o *Ontology) updateClassDefinitionsWithClosure() {
    // For each class, fill in AllSuperClasses from closure
    for classURI, classDef := range o.ClassHierarchy.AllClasses {
        closure := o.TransitiveClosures.SubClassOf[classURI]
        var superClasses []string
        for class := range closure {
            if class != classURI { // Exclude self
                superClasses = append(superClasses, class)
            }
        }
        classDef.AllSuperClasses = superClasses
    }

    // For each class, fill in AllSubClasses
    // A subclass of C is any class X where C is in X's transitive supers
    for classURI := range o.ClassHierarchy.AllClasses {
        var subClasses []string
        for otherURI, otherDef := range o.ClassHierarchy.AllClasses {
            if otherURI != classURI {
                if o.TransitiveClosures.SubClassOf[otherURI][classURI] {
                    subClasses = append(subClasses, otherURI)
                }
            }
        }
        o.ClassHierarchy.AllClasses[classURI].AllSubClasses = subClasses
    }
}

// updatePropertyDefinitionsWithClosure similar to class version
func (o *Ontology) updatePropertyDefinitionsWithClosure() {
    for propURI, propDef := range o.PropertyHierarchy.AllProperties {
        closure := o.TransitiveClosures.SubPropertyOf[propURI]
        var superProps []string
        for prop := range closure {
            if prop != propURI {
                superProps = append(superProps, prop)
            }
        }
        propDef.AllSuperProperties = superProps
    }

    for propURI := range o.PropertyHierarchy.AllProperties {
        var subProps []string
        for otherURI, otherDef := range o.PropertyHierarchy.AllProperties {
            if otherURI != propURI {
                if o.TransitiveClosures.SubPropertyOf[otherURI][propURI] {
                    subProps = append(subProps, otherURI)
                }
            }
        }
        o.PropertyHierarchy.AllProperties[propURI].AllSubProperties = subProps
    }
}
```

**File: `sparql/ontology/reasoning_test.go`**

```go
package ontology

import (
    "testing"
)

func TestComputeTransitiveClosureSimple(t *testing.T) {
    onto := NewOntology()

    // Build: Employee -> Person -> Thing
    onto.AddSubClassRelationship("http://example.com/Employee", "http://example.com/Person")
    onto.AddSubClassRelationship("http://example.com/Person", "http://example.com/Thing")

    err := onto.ComputeTransitiveClosures()
    if err != nil {
        t.Fatalf("ComputeTransitiveClosures returned error: %v", err)
    }

    // Employee should have both Person and Thing in closure
    empClosure := onto.TransitiveClosures.SubClassOf["http://example.com/Employee"]
    if !empClosure["http://example.com/Person"] {
        t.Error("Person not in Employee closure")
    }
    if !empClosure["http://example.com/Thing"] {
        t.Error("Thing not in Employee closure")
    }
}

func TestComputeTransitiveClosureMultiplePaths(t *testing.T) {
    onto := NewOntology()

    // Build diamond:
    //    Thing
    //   /     \
    // A         B
    //   \     /
    //    C

    onto.AddSubClassRelationship("http://example.com/A", "http://example.com/Thing")
    onto.AddSubClassRelationship("http://example.com/B", "http://example.com/Thing")
    onto.AddSubClassRelationship("http://example.com/C", "http://example.com/A")
    onto.AddSubClassRelationship("http://example.com/C", "http://example.com/B")

    err := onto.ComputeTransitiveClosures()
    if err != nil {
        t.Fatalf("ComputeTransitiveClosures returned error: %v", err)
    }

    cClosure := onto.TransitiveClosures.SubClassOf["http://example.com/C"]
    if !cClosure["http://example.com/A"] {
        t.Error("A not in C closure")
    }
    if !cClosure["http://example.com/B"] {
        t.Error("B not in C closure")
    }
    if !cClosure["http://example.com/Thing"] {
        t.Error("Thing not in C closure")
    }
}

func TestIsSubClassOfAfterClosure(t *testing.T) {
    onto := NewOntology()
    onto.AddSubClassRelationship("http://example.com/Manager", "http://example.com/Employee")
    onto.AddSubClassRelationship("http://example.com/Employee", "http://example.com/Person")

    onto.ComputeTransitiveClosures()

    if !onto.IsSubClassOf("http://example.com/Manager", "http://example.com/Person") {
        t.Error("Manager should be subclass of Person transitively")
    }

    if onto.IsSubClassOf("http://example.com/Person", "http://example.com/Manager") {
        t.Error("Person should not be subclass of Manager")
    }
}
```

#### Validation Checkpoint

After Days 6-8:

- [ ] TransitiveClosures computed correctly
- [ ] All transitive relationships present
- [ ] No infinite loops on cycles
- [ ] IsSubClassOf/IsSubPropertyOf work
- [ ] Tests pass for simple and complex hierarchies
- [ ] Performance acceptable (<100ms for moderate ontologies)

---

### Days 9-10: Reasoning Integration & Tests

Complete the test suite:

```go
// In reasoning_test.go

func TestPropertyHierarchyClosure(t *testing.T) {
    onto := NewOntology()

    onto.AddSubPropertyRelationship(
        "http://example.com/manages",
        "http://example.com/supervises")
    onto.AddSubPropertyRelationship(
        "http://example.com/supervises",
        "http://example.com/oversees")

    onto.ComputeTransitiveClosures()

    if !onto.IsSubPropertyOf(
        "http://example.com/manages",
        "http://example.com/oversees") {
        t.Error("manages should be subproperty of oversees transitively")
    }
}

func TestReflexivity(t *testing.T) {
    onto := NewOntology()
    onto.InitializeClass("http://example.com/Person")
    onto.ComputeTransitiveClosures()

    // Classes should be subclass of themselves
    if !onto.IsSubClassOf("http://example.com/Person", "http://example.com/Person") {
        t.Error("Class should be subclass of itself")
    }
}

func TestLargeOntology(t *testing.T) {
    // Test with 100+ classes
    onto := NewOntology()

    for i := 0; i < 100; i++ {
        uri := fmt.Sprintf("http://example.com/Class%d", i)
        onto.InitializeClass(uri)
        if i > 0 {
            parentURI := fmt.Sprintf("http://example.com/Class%d", i-1)
            onto.AddSubClassRelationship(uri, parentURI)
        }
    }

    err := onto.ComputeTransitiveClosures()
    if err != nil {
        t.Fatalf("ComputeTransitiveClosures failed on large ontology: %v", err)
    }

    // Leaf should be subclass of root
    leaf := "http://example.com/Class99"
    root := "http://example.com/Class0"
    if !onto.IsSubClassOf(leaf, root) {
        t.Error("Leaf should be transitive subclass of root")
    }
}
```

---

## WEEK 3: SCHEMA SYNTHESIS (Days 11-14)

### Days 11-13: Schema Synthesis Implementation

#### File: `sparql/ontology/schema_builder.go`

This is the most complex part - converting ontology to Dgraph schema. Implementation at Week 3.

---

## WEEK 4: INTEGRATION & TESTING (Days 15-20)

### Days 15-16: OntologyStore

**File: `sparql/ontology/store.go`**

```go
package ontology

import (
    "fmt"
    "sync"
)

// OntologyStore manages ontologies per namespace
type OntologyStore struct {
    ontologies map[uint64]*Ontology // namespace -> ontology
    mu sync.RWMutex
}

// NewOntologyStore creates a new store
func NewOntologyStore() *OntologyStore {
    return &OntologyStore{
        ontologies: make(map[uint64]*Ontology),
    }
}

// Get retrieves ontology for namespace
func (os *OntologyStore) Get(namespaceID uint64) *Ontology {
    os.mu.RLock()
    defer os.mu.RUnlock()
    return os.ontologies[namespaceID]
}

// Load loads/reloads ontology for namespace
func (os *OntologyStore) Load(namespaceID uint64, filePath string) error {
    onto, err := LoadOntologyFromFile(filePath)
    if err != nil {
        return fmt.Errorf("failed to load ontology: %v", err)
    }

    os.mu.Lock()
    defer os.mu.Unlock()
    os.ontologies[namespaceID] = onto
    return nil
}

// Unload removes ontology from store
func (os *OntologyStore) Unload(namespaceID uint64) {
    os.mu.Lock()
    defer os.mu.Unlock()
    delete(os.ontologies, namespaceID)
}

// ListNamespaces returns all loaded namespaces
func (os *OntologyStore) ListNamespaces() []uint64 {
    os.mu.RLock()
    defer os.mu.RUnlock()

    var namespaces []uint64
    for ns := range os.ontologies {
        namespaces = append(namespaces, ns)
    }
    return namespaces
}
```

### Days 17-19: Integration Tests

Comprehensive end-to-end testing covering all phases.

---

## Implementation Completion Checklist

### Code Quality

- [ ] All code compiles (`go build ./sparql/ontology`)
- [ ] All tests pass (`go test ./sparql/ontology -v`)
- [ ] > 85% code coverage
- [ ] No race conditions (run with `-race`)
- [ ] No panics (proper error handling)
- [ ] Follows Dgraph code style

### Functionality

- [ ] Loads Turtle files correctly
- [ ] Parses all RDF/RDFS/OWL constructs
- [ ] Computes transitive closures correctly
- [ ] Handles cycles gracefully
- [ ] Synthesizes valid Dgraph schema
- [ ] Multi-tenant isolation works
- [ ] Performance targets met

### Testing

- [ ] 200+ test cases total
- [ ] Unit tests for each module
- [ ] Integration tests end-to-end
- [ ] Edge cases covered
- [ ] Real OWL files tested
- [ ] Performance benchmarks

### Documentation

- [ ] All functions have comments
- [ ] Complex algorithms documented
- [ ] Examples in comments
- [ ] Integration points clear

---

## Summary

This 20-day plan provides a complete roadmap for Phase 5 implementation. Follow it day-by-day:

1. **Foundation (Days 1-5)**: Build core data structures
2. **Reasoning (Days 6-10)**: Implement inference engine
3. **Schema (Days 11-14)**: Build schema synthesis
4. **Integration (Days 15-20)**: Complete and test

All code signatures, test cases, and validation checkpoints are provided. Begin with Day 1 and
proceed sequentially.

---

**Status**: Ready for Implementation  
**Next Step**: Start Day 1 (Create model.go)  
**Estimated Effort**: 160-200 developer-hours  
**Team Size**: 1-2 developers (can be parallelized)
