# Phase 5: Ontology Foundation Design Specification

**Status**: Design Proposal  
**Date**: January 2026  
**Authors**: Dgraph Architecture Team  
**Related Spec**: sparql/ARCHITECTURE_SPEC.md (Phase 5 section)  
**Revision**: 1.0

---

## Table of Contents

1. [Section 1: Research Summary](#section-1-research-summary)
2. [Section 2: Ontology Model Design](#section-2-ontology-model-design)
3. [Section 3: Loading Strategy](#section-3-loading-strategy)
4. [Section 4: Query Optimization Integration](#section-4-query-optimization-integration)
5. [Section 5: Implementation Plan](#section-5-implementation-plan)
6. [Section 6: Reuse & Integration Decisions](#section-6-reuse--integration-decisions)
7. [Section 7: Open Questions & Future Extensions](#section-7-open-questions--future-extensions)

---

## Section 1: Research Summary

### 1.1 Dgraph's Schema System

#### Overview

Dgraph's schema system spans multiple layers:

**GraphQL Schema Layer** (`graphql/schema/`):

- Uses `gqlparser/ast` types from the gqlparser library
- Core `Type` and `FieldDefinition` interfaces defined in `wrappers.go`
- `Type` interface provides:
  - `Field(name string) FieldDefinition` - gets field by name
  - `Fields() []FieldDefinition` - gets all fields
  - `IDField() FieldDefinition` - gets the ID field
  - `DgraphPredicate(fld string) string` - maps field to Dgraph predicate
  - `Name()` - type name
  - Property accessors for nullability, unions, interfaces
- `FieldDefinition` interface provides:
  - `Name()` and `DgraphAlias()` for field identity
  - `Type()` - returns the Type this field is
  - `ParentType()` - returns the containing Type
  - `DgraphPredicate()` - maps to Dgraph storage predicate
  - Directives like `@dgraph`, `@search`, `@hasInverse`, `@auth`

**DQL Parser Layer** (`dql/parser.go`):

- `GraphQuery` struct represents parsed DQL queries in tree form
- Key fields:
  - `UID []uint64` - starting UIDs
  - `Attr string` - predicate being queried
  - `Children []*GraphQuery` - nested patterns
  - `Filter *FilterTree` - where clauses
  - `Order []*pb.Order` - sorting directives
  - `Args map[string]string` - query arguments
  - `AllowedPreds []string` - ACL filtering
- Supports functions, expansions, recursion, shortest paths, facets

**Proto/Storage Layer** (`protos/pb/`, `schema/parse.go`):

- `pb.SchemaUpdate` struct defines predicates:
  - `Predicate string` - predicate name
  - `ValueType pb.Posting_ValType` - type (string, int, float, bool, uid, etc.)
  - `List bool` - is list-valued
  - `Directive pb.SchemaUpdate_Directive` - reverse, index, upsert, etc.
  - `Tokenizer []string` - search indexes (term, fulltext, hash, etc.)
  - `Count bool`, `Upsert bool`, `Unique bool` - metadata
- `types.TypeID` enum: StringID, Int64ID, FloatID, BoolID, UidID, PasswordID, DefaultID, BinaryID,
  DateTimeID, GeoID, etc.

**Posting/Storage Layer** (`posting/list.go`):

- `List` struct manages in-memory posting lists
- Stores both immutable layer (`*pb.PostingList`) and mutable deltas
- Predicates support:
  - Multiple values per key
  - Metadata (facets, language tags)
  - Versioning via timestamps
  - Compression via snappy, delta encoding
- Indexes stored with special predicates (e.g., `pred|index_type`)

#### Key Insights for Ontologies

1. **Type information is decoupled**: GraphQL types map to Dgraph predicates but aren't directly
   stored as data
2. **Predicates are first-class**: Every attribute is a predicate with metadata in proto format
3. **Indexes enable efficient querying**: Full-text, term, exact, regex, range indexes available
4. **Metadata flexibility**: Facets and directives allow storing additional semantic information
5. **ACL integration**: `AllowedPreds` in GraphQuery enables predicate-level authorization

### 1.2 How Dgraph's Schema System Works

#### Registration Pattern

1. User defines GraphQL schema (e.g., `type Author { name: String! }`)
2. Schema handler parses with `gqlparser` into AST
3. For each type and field, `dgraphPredicate` map stores GraphQL→DQL mapping
   - Example: `Author.name` → predicate `author_name`
4. Schema handler generates Dgraph schema predicates:
   ```
   author_name: string @index(term, fulltext) .
   <Author>: [uid] @type .
   ```

#### Type System Interaction

- GraphQL types become Dgraph predicate metadata
- Type edges created as reverse predicates to instances
- Fields in types become predicates with type constraints
- Interfaces and unions use type discrimination predicates

#### Predicate Access Pattern

- At query compilation, `graphql/schema.Schema` interface provides:
  - Type lookup by name
  - Field lookup within type
  - Dgraph predicate name for field
  - Type metadata (nullability, list, interface, union)
  - Search indexes available on predicate
  - Authorization rules for field

#### Statistics & Cardinality

- **Not directly stored** in schema itself
- Calculated at query planning time via:
  - `query/execution.go` estimates join order
  - `worker/task.go` has cardinality estimation
  - Statistics gathered from actual query execution
  - Index statistics available via `posting/index.go`

### 1.3 How Dgraph's Storage System Works

#### Posting Lists

- Every predicate has a posting list: `key=predicate|uid, value=PostingList`
- `PostingList` contains array of `Posting` protos:
  ```go
  type Posting struct {
    Uid      uint64          // The UID of the entity
    Value    []byte          // For non-UID predicates
    ValType  ValType         // Type of value
    Facets   []*Facet        // Metadata
    Label    string          // Deleted/special marker
  }
  ```

#### Indexes

- Separate posting lists for indexed predicates
- Key format: `predicate|tokenizer|token → []UID`
- Examples:
  - `name|fulltext|"john" → [uid1, uid2]` (fulltext search)
  - `age|range_0_100|"50" → [uid5, uid10]` (range index)
  - `email|hash|hash("user@example.com") → [uid3]` (unique check)

#### Metadata Storage

- Facets store key-value metadata on edges
- Schema metadata in separate `dgraph.type` predicate
- Type information stored via implicit predicates:
  - `dgraph.type` holds type name as string value
  - Unique values accessible via `uid_in` function

#### Transaction Management

- MVCC with timestamp-based versioning
- Posting lists versioned by commit timestamp
- Query execution reads at specific timestamp snapshot
- Updates create new version without overwriting

### 1.4 Relevant Abstractions to Reuse

#### Type System Patterns

1. **Interface-based abstraction**: `Type` and `FieldDefinition` interfaces allow plugging in new
   implementations
2. **Predicate mapping**: `dgraphPredicate` map pattern can be extended for ontology predicates
3. **Metadata on predicates**: Directives and facets already support storing extra information

#### Storage Patterns

1. **Posting lists**: Can store ontology data (classes, properties, hierarchies) as regular
   predicates
2. **Indexes**: Term/fulltext indexes can enable efficient ontology lookups
3. **Metadata predicates**: Special predicates (like `dgraph.type`) can store ontology metadata

#### Query Patterns

1. **GraphQuery tree structure**: Excellent foundation for query building
2. **Filter composition**: `FilterTree` pattern works for ontology pattern matching
3. **Function system**: Existing `uid()`, `has()`, `type()` functions extensible

#### ACL/Authorization Patterns

1. **Predicate filtering**: `AllowedPreds` mechanism can restrict ontology access
2. **Type-based rules**: Auth rules on types can extend to ontology classes

#### Schema Integration Points

1. **schemagen.go**: Generates Dgraph schema from GraphQL - can add ontology synthesis
2. **gqlschema.go**: Type definitions, can add ontology type metadata
3. **rules.go**: Validation rules - can validate ontology consistency

### 1.5 Constraints and Limitations

#### Storage Constraints

1. **Predicate cardinality**: Each ontology element becomes a predicate; many ontologies have
   thousands of classes
2. **UID space**: Large ontologies can consume significant UID space for concept resources
3. **Index size**: Full-text indexes on class names may be memory-intensive
4. **Facet overhead**: Storing metadata via facets adds storage per triple

#### Performance Constraints

1. **Query compilation**: Loading ontology on every query is expensive
2. **Memory footprint**: Fully materialized ontology in memory significant for large ontologies
3. **Network calls**: Fetching ontology data from storage layer requires network I/O in distributed
   setup
4. **Cache invalidation**: Complex ontologies may have invalidation cascades

#### Semantic Constraints

1. **RDF/SPARQL semantics**: Ontology reasoning may not match Dgraph's execution model
2. **Circular definitions**: Dgraph doesn't prevent cycles in predicate definitions
3. **Constraint validation**: Dgraph doesn't enforce OWL cardinality restrictions
4. **Inference rules**: SPARQL FILTER logic doesn't directly map to Dgraph constraints

#### Authorization Constraints

1. **Graph-level access**: Can't restrict access to specific ontology elements easily
2. **Namespace isolation**: Multi-tenancy requires careful predicate namespacing
3. **Dynamic rules**: Auth rules are static, can't depend on ontology structure

---

## Section 2: Ontology Model Design

### 2.1 How Ontologies Will Be Represented in Dgraph

#### Core Principle: Ontologies as RDF Graphs

Ontologies are themselves RDF graphs. We model them using Dgraph predicates that follow RDF/RDFS
semantics:

```
Core RDF Classes (stored as predicates):
- rdf:type → uid predicate (what type/class an entity has)
- rdf:Property → class for properties
- rdf:Class → class for classes
- rdf:subClassOf → uid predicate (class hierarchy)
- rdf:subPropertyOf → uid predicate (property hierarchy)
- rdf:domain → uid predicate (property domain restriction)
- rdf:range → uid predicate (property range restriction)

RDFS Extensions:
- rdfs:label → string predicate (human-readable name)
- rdfs:comment → string predicate (documentation)
- rdfs:isDefinedBy → uid predicate (source ontology)

OWL Extensions (Phase 5+):
- owl:equivalentClass → uid predicate (class equivalence)
- owl:equivalentProperty → uid predicate (property equivalence)
- owl:inverseOf → uid predicate (inverse relationships)
- owl:disjointWith → uid predicate (class disjointness)
- owl:Restriction → uid class (property restrictions)
```

### 2.2 Data Model for OWL/RDFS Concepts

#### Ontology Concept Model

Every ontology element becomes a resource:

```go
// Conceptual structure (stored as Dgraph triples)
type OntologyClass {
  ID                 uid          // Unique ID
  name               string       // e.g., "Person"
  label              string       // Human-readable label
  comment            string       // Documentation
  equivalentClasses  [uid]        // Equivalent class URIs
  superClasses       [uid]        // Direct parent classes
  subClasses         [uid]        // Direct child classes  (computed via inference)
  properties         [uid]        // Properties with this as domain
  isAbstract         bool         // Facet on rdf:type predicate
}

type OntologyProperty {
  ID                 uid          // Unique ID
  name               string       // e.g., "hasFriend"
  label              string       // Human-readable label
  comment            string       // Documentation
  domain             uid          // What classes can have this property
  range              uid          // What type of values property has
  superProperties    [uid]        // Parent properties
  subProperties      [uid]        // Child properties
  inverseProperty    uid          // Inverse (if exists)
  isFunctional       bool         // Facet: true if max cardinality 1
  isTransitive       bool         // Facet: true if transitive
}

type OntologyRestriction {
  ID                 uid          // Unique ID
  onProperty         uid          // Which property
  minCardinality     int          // Facet: minimum occurrences
  maxCardinality     int          // Facet: maximum occurrences
  someValuesFrom     uid          // Facet: class constraint
  allValuesFrom      uid          // Facet: class constraint
}
```

#### Representation as Dgraph Schema

```
// Types that exist in both GraphQL schema and ontology storage
type OntologyClass {
  id: ID!
  uri: String!
  label: String
  comment: String
  equivalentClasses: [OntologyClass]
  superClasses: [OntologyClass]
  subClasses: [OntologyClass]
  properties: [OntologyProperty]
}

type OntologyProperty {
  id: ID!
  uri: String!
  label: String
  comment: String
  domain: OntologyClass
  range: OntologyClass
  superProperties: [OntologyProperty]
  subProperties: [OntologyProperty]
  inverseProperty: OntologyProperty
}

// Predicates used in ontology storage
ontology_class_uri: string @index(exact) .
ontology_class_label: string @index(term) .
ontology_class_comment: string .
ontology_class_equiv: uid @reverse .
ontology_class_super: uid @reverse .

ontology_property_uri: string @index(exact) .
ontology_property_label: string @index(term) .
ontology_property_domain: uid @reverse .
ontology_property_range: uid @reverse .

// Type declarations
<OntologyClass>: [uid] .
<OntologyProperty>: [uid] .
```

### 2.3 Storage Strategy: Hybrid Approach

#### Three-Layer Storage Architecture

**Layer 1: Persistent Storage (Dgraph)**

- Store complete ontology as RDF triples using predicates described above
- Enable SPARQL queries against ontology itself
- Support incremental ontology updates
- Enable multi-tenancy via namespace prefixes
- Allow ontology versioning via named graphs

**Layer 2: In-Memory Model (Process)**

- Cache parsed ontology in Go structs for fast access
- Built during ontology loading phase
- Provides O(1) lookup of class/property metadata
- Enables efficient hierarchy traversal
- Supports invalidation when ontology updates

**Layer 3: Query-Time Index (Per-Query)**

- Build transient optimization structures during query compilation
- Class hierarchy expansion cache
- Property domain/range lookup tables
- Reasoning rule application tracking
- Invalidate after query execution

#### Data Flow for Storage

```
Ontology File (RDF/OWL)
    ↓
[Ontology Loader] - Parse RDF/N-Triples/Turtle format
    ↓
OntologyModel (Go structs)
    ├─ classes: map[URI]*Class
    ├─ properties: map[URI]*Property
    └─ hierarchies: map[URI][]URI
    ↓
[Storage Writer]
    ├─ Write to Dgraph as triples
    ├─ Create indexes for fast lookup
    └─ Store metadata predicates
    ↓
Dgraph Posting Lists
```

#### Caching Strategy

**L1 Cache: Process-Level Ontology Model**

- Loaded once at ontology initialization
- TTL-based invalidation on updates
- Shared across all queries in process
- Memory footprint ~10-100MB for typical ontologies (5K-50K classes)

**L2 Cache: Query Compiler Cache**

- Class hierarchy expansion results
- Property metadata lookup results
- Populated during query compilation Phase 3
- Scoped to single query
- Cleared after query execution

**L3 Cache: Statistics Cache**

- Class instance counts
- Property usage patterns
- Populated lazily during optimization
- Updated periodically (configurable)
- Survives across queries

### 2.4 Efficient Ontology Query Patterns

#### Key Query Patterns

**Pattern 1: Class Lookup**

```dql
{
  query {
    ontology_class_uri(func: exact(ontology_class_uri, "http://example.com/Person"))
  }
}
```

Uses `@index(exact)` on URI for O(1) lookup

**Pattern 2: Hierarchy Traversal**

```dql
{
  query(func: uid(personClass)) {
    uid
    ontology_class_super {
      uid
      ontology_class_label
    }
  }
}
```

Leverages `@reverse` on `ontology_class_super` for efficient walking

**Pattern 3: Property Domain/Range**

```dql
{
  query(func: eq(ontology_property_domain, personClassUID)) {
    uid
    ontology_property_label
    ontology_property_range {
      ontology_class_label
    }
  }
}
```

Direct predicate queries with small result sets

**Pattern 4: Transitive Closure (Subclass)**

```dql
{
  query(func: uid(personClass)) @recurse(depth: 10) {
    uid
    ontology_class_super
  }
}
```

Uses `@recurse` to compute transitive superclasses

#### Design for Performance

1. **Exact match indexes**: URIs indexed for O(1) lookup
2. **Term indexes**: Labels/comments indexed for full-text search
3. **Reverse indexes**: Enable traversing backward relationships
4. **Small result sets**: Most ontology queries return <1000 results
5. **In-memory caching**: Ontology rarely changes during query processing

### 2.5 Integration with Schema System

#### Schema Synthesis Process

The `SchemaBuilder` component integrates ontology knowledge into Dgraph's schema:

```go
type SchemaBuilder interface {
  // BuildFromOntology creates schema from ontology concepts
  BuildFromOntology(model *OntologyModel) (*pb.Schema, error)

  // EnhanceExistingSchema adds ontology knowledge to existing schema
  EnhanceExistingSchema(existing *pb.Schema, model *OntologyModel) (*pb.Schema, error)

  // ComputeInferenceRules generates reasoning rules from ontology
  ComputeInferenceRules(model *OntologyModel) ([]*ReasoningRule, error)
}
```

#### How Ontology Enhances Schema

**Input Schema** (User-defined GraphQL):

```graphql
type Person {
  name: String!
  friends: [Person]
}
```

**Ontology Knowledge**:

```
Person
  ├─ subClassOf: Agent
  └─ properties: hasFriend (domain: Person, range: Person)
```

**Enhanced Schema** (Synthesized):

```graphql
type Person implements Agent {
  name: String!
  friends: [Person] @dgraph(pred: "has_friend")
  # Synthesized from ontology:
  # - Inherits all Agent properties
  # - Automatically validated via ontology constraints
  # - Query optimizer aware of class hierarchy
}
```

#### Integration Points

1. **Type Creation**: When creating new types, inherit from ontology superclasses
2. **Field Definition**: When defining fields, use ontology property metadata
3. **Validation**: Validate schema against ontology constraints
4. **Authorization**: Derive auth rules from ontology role definitions
5. **Optimization**: Pass ontology knowledge to query optimizer

---

## Section 3: Loading Strategy

### 3.1 How to Load OWL/RDFS Files

#### Supported Input Formats

**N-Triples** (simplest, one triple per line):

```
<http://example.com/Person> <http://www.w3.org/2000/01/rdf-schema#label> "Person" .
<http://example.com/Person> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2000/01/rdf-schema#Class> .
```

**Turtle** (compact, similar to N3):

```turtle
@prefix ex: <http://example.com/> .
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .

ex:Person a rdfs:Class ;
  rdfs:label "Person" ;
  rdfs:comment "A human person" .
```

**RDF/XML** (verbose, XML-based):

```xml
<rdf:RDF xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
         xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
         xmlns:ex="http://example.com/">
  <rdfs:Class rdf:about="http://example.com/Person">
    <rdfs:label>Person</rdfs:label>
  </rdfs:Class>
</rdf:RDF>
```

**OWL** (ontology format, extends RDF):

```turtle
@prefix owl: <http://www.w3.org/2002/07/owl#> .

ex:Person a owl:Class ;
  rdfs:subClassOf ex:Agent ;
  owl:equivalentClass [ owl:intersectionOf (ex:Human ex:Living) ] .
```

#### Loading Pipeline

```go
type OntologyLoader interface {
  // LoadFromFile reads ontology from file
  LoadFromFile(filePath string, format Format) (*OntologyModel, error)

  // LoadFromStream reads from io.Reader
  LoadFromStream(reader io.Reader, format Format) (*OntologyModel, error)

  // LoadFromDgraph reads ontology already stored in Dgraph
  LoadFromDgraph(ctx context.Context, client Client, namespace string) (*OntologyModel, error)
}

type Format int
const (
  NTriples Format = iota
  Turtle
  RDFXml
  JSONLD
)
```

**Step 1: Parsing** → RDF Triple Stream

- Use gorpc/rdf2go for Go RDF parsing
- Stream-based parsing for large files
- URI validation and normalization
- Namespace resolution

**Step 2: Filtering** → Ontology Relevant Triples

- Keep RDF/RDFS/OWL vocabulary triples
- Skip data triples (unless embedded in ontology)
- Filter by namespace (to support multi-tenancy)

**Step 3: Structuring** → OntologyModel

- Group triples by subject
- Classify subjects as Classes, Properties, Restrictions
- Build hierarchy relationships
- Validate ontology consistency

**Step 4: Validation** → OntologyModel (validated)

- Check no circular class hierarchies (except through inference)
- Validate domain/range specifications
- Check for undefined references
- Warn on conflicts with existing schema

**Step 5: Caching** → In-Memory Model

- Create index structures
- Pre-compute transitive closures (optional)
- Store in process-level cache
- Set TTL for invalidation

### 3.2 Parse into Dgraph Representation

#### Triple to Dgraph Conversion

Each RDF triple `(subject, predicate, object)` becomes one or more Dgraph operations:

**Class Definition Triple**:

```
<http://example.com/Person> rdf:type rdfs:Class
```

Becomes:

```go
// Create resource node
newUID := generateUID()
mutations := []*MutationOp{
  {
    Predicate: "ontology_class_uri",
    Object: "http://example.com/Person",
    UID: newUID,
  },
  {
    Predicate: "dgraph.type",
    Object: "OntologyClass",
    UID: newUID,
  },
}

// Store mapping for future reference
ontologyModel.classes["http://example.com/Person"] = newUID
```

**Property Definition Triple**:

```
<http://example.com/hasFriend> rdf:type rdf:Property
```

Becomes:

```go
newUID := generateUID()
mutations := []*MutationOp{
  {
    Predicate: "ontology_property_uri",
    Object: "http://example.com/hasFriend",
    UID: newUID,
  },
  {
    Predicate: "dgraph.type",
    Object: "OntologyProperty",
    UID: newUID,
  },
}
```

**Hierarchy Triple**:

```
<http://example.com/Student> rdfs:subClassOf <http://example.com/Person>
```

Becomes:

```go
mutations := []*MutationOp{
  {
    Predicate: "ontology_class_super",
    Subject: studentUID,
    Object: personUID,
  },
}

// Update in-memory model
ontologyModel.classHierarchy[studentUID] = append(
  ontologyModel.classHierarchy[studentUID],
  personUID,
)
```

**Property Metadata Triple**:

```
<http://example.com/hasFriend> rdfs:domain <http://example.com/Person>
```

Becomes:

```go
mutations := []*MutationOp{
  {
    Predicate: "ontology_property_domain",
    Subject: hasFriendUID,
    Object: personUID,
  },
}
```

#### Batch Operations for Efficiency

```go
type BatchLoader struct {
  batchSize    int  // e.g., 1000
  mutations    []*Mutation
  cache        map[string]uint64  // URI → UID cache
  dgraphClient *dqlc.GraphQLClient
}

func (bl *BatchLoader) LoadOntology(triples []*Triple) error {
  for _, triple := range triples {
    mutation := bl.tripleToMutation(triple)
    bl.mutations = append(bl.mutations, mutation)

    if len(bl.mutations) >= bl.batchSize {
      if err := bl.flush(); err != nil {
        return err
      }
    }
  }

  return bl.flush()  // Final batch
}

func (bl *BatchLoader) flush() error {
  // Submit all mutations to Dgraph in single transaction
  req := &pb.Request{
    Mutations: bl.mutations,
  }

  if _, err := bl.dgraphClient.Mutate(context.Background(), req); err != nil {
    return fmt.Errorf("mutation failed: %w", err)
  }

  bl.mutations = nil
  return nil
}
```

### 3.3 Validation and Error Handling

#### Validation Checks

```go
type OntologyValidator interface {
  // ValidateSyntax checks RDF/OWL syntax
  ValidateSyntax(model *OntologyModel) []ValidationError

  // ValidateSemantics checks logical consistency
  ValidateSemantics(model *OntologyModel) []ValidationError

  // ValidateConstraints checks OWL constraints
  ValidateConstraints(model *OntologyModel) []ValidationError
}

type ValidationError struct {
  Severity  ErrorSeverity  // Error, Warning, Info
  Code      string         // e.g., "CIRCULAR_HIERARCHY"
  Message   string
  Location  Triple         // Which triple caused error
  Fix       string         // Suggested fix
}

type ErrorSeverity int
const (
  ErrorSeverityError ErrorSeverity = iota
  ErrorSeverityWarning
  ErrorSeverityInfo
)
```

**Critical Validations**:

1. **Circular hierarchies**: Detect if `A subClassOf B subClassOf ... subClassOf A`
2. **Undefined references**: Check all classes/properties mentioned exist
3. **Type consistency**: Check objects have correct types (string vs UID)
4. **Namespace isolation**: Verify namespace prefixes don't conflict

**Warning Validations**:

1. **Orphan classes**: Classes not related to other classes
2. **Unused properties**: Properties with no domain
3. **Incomplete definitions**: Classes without labels
4. **Potential conflicts**: Ontology classes matching existing schema types

#### Error Handling Strategy

```go
type LoadResult struct {
  Model       *OntologyModel
  Errors      []ValidationError
  Warnings    []ValidationError
  Stats       LoadStats
}

type LoadStats struct {
  ClassesLoaded     int
  PropertiesLoaded  int
  RelationshipsLoaded int
  Duration          time.Duration
  TriplesParsed     int
  TriplesIgnored    int
}

func LoadOntology(ctx context.Context, filePath string) (*LoadResult, error) {
  result := &LoadResult{}

  // Load and parse
  model, syntaxErrors := parseOntologyFile(filePath)
  if len(syntaxErrors) > 0 {
    result.Errors = syntaxErrors
    return result, fmt.Errorf("syntax errors in ontology")
  }

  // Validate semantics
  semanticErrors := validateSemantics(model)
  result.Errors = append(result.Errors, semanticErrors...)

  if len(result.Errors) > 0 {
    return result, fmt.Errorf("%d validation errors", len(result.Errors))
  }

  // Warnings don't stop loading
  result.Warnings = validateWarnings(model)

  result.Model = model
  return result, nil
}
```

### 3.4 Caching Strategy for Query Optimization

#### Cache Hierarchy

**Cache Level 1: Ontology Model Cache**

```go
type OntologyCache struct {
  // Loaded ontology model
  model *OntologyModel

  // URI → UID mappings for fast lookup
  classURIToUID      map[string]uint64
  propertyURIToUID   map[string]uint64

  // Hierarchy traversal cache
  superClassCache    map[uint64][]uint64  // Transitive superclasses
  subClassCache      map[uint64][]uint64  // Transitive subclasses

  // Last updated timestamp
  lastUpdated        time.Time

  // Cache hit/miss statistics
  hits, misses       int
}

func (oc *OntologyCache) GetClass(uri string) (uint64, bool) {
  if uid, ok := oc.classURIToUID[uri]; ok {
    oc.hits++
    return uid, true
  }
  oc.misses++
  return 0, false
}
```

**Cache Level 2: Query Compiler Cache**

```go
type QueryCompilerCache struct {
  // Expanded class sets for queries
  classExpansions map[string][]uint64

  // Property metadata by UID
  propertyMetadata map[uint64]*PropertyMeta

  // Reasoning results
  inferenceSets map[string][]uint64

  // Valid at this query start time
  validAt time.Time
}

type PropertyMeta struct {
  URI              string
  Domain           uint64
  Range            uint64
  IsTransitive     bool
  IsFunctional     bool
  InverseProperty  uint64
  SuperProperties  []uint64
}
```

**Cache Level 3: Statistics Cache**

```go
type StatisticsCache struct {
  // Instance counts per class
  classInstanceCounts map[uint64]int64

  // Property usage patterns
  propertyFrequency map[uint64]float64

  // Join selectivity estimates
  selectivity map[string]float64

  // Last refreshed
  refreshedAt time.Time
}
```

#### Cache Invalidation Strategy

**Time-Based Invalidation**:

```go
const (
  OntologyModelCacheTTL = 1 * time.Hour
  StatisticsCacheTTL    = 30 * time.Minute
)

func (oc *OntologyCache) IsValid() bool {
  return time.Since(oc.lastUpdated) < OntologyModelCacheTTL
}

func (sc *StatisticsCache) Refresh(force bool) error {
  if !force && time.Since(sc.refreshedAt) < StatisticsCacheTTL {
    return nil
  }
  return sc.recomputeStatistics()
}
```

**Event-Based Invalidation**:

```go
type OntologyUpdateEvent struct {
  Operation   UpdateOp  // Add, Update, Delete
  Affected    []string  // URIs affected
  Timestamp   time.Time
}

type UpdateOp int
const (
  UpdateOpAdd UpdateOp = iota
  UpdateOpModify
  UpdateOpDelete
)

func (oc *OntologyCache) HandleUpdate(event *OntologyUpdateEvent) {
  // Invalidate affected entries
  for _, uri := range event.Affected {
    delete(oc.superClassCache, oc.classURIToUID[uri])
    delete(oc.subClassCache, oc.classURIToUID[uri])
  }

  // Invalidate dependent caches
  oc.lastUpdated = event.Timestamp
}
```

### 3.5 Invalidation When Ontology Changes

#### Change Detection

```go
type OntologyChangeDetector interface {
  // DetectChanges compares two ontology models
  DetectChanges(old, new *OntologyModel) []*OntologyUpdateEvent

  // WatchOntology subscribes to ontology changes
  WatchOntology(ctx context.Context, namespace string) chan *OntologyUpdateEvent
}
```

#### Invalidation Cascade

```
ontology_class_uri mutation
  ↓
Invalidate: classURIToUID[uri]
  ↓
Invalidate: superClassCache[uid], subClassCache[uid]
  ↓
Invalidate: PropertyMeta for properties with domain=uid or range=uid
  ↓
Notify: query compiler to invalidate expansion caches
  ↓
Update: statistics cache for affected classes
```

#### Example: Updating Class Hierarchy

```go
func (oc *OntologyCache) UpdateClassHierarchy(class, superClass uint64) error {
  // Detect circular dependency
  if !oc.isAcyclic(superClass, class) {
    return fmt.Errorf("circular hierarchy detected")
  }

  // Update model
  oc.model.classHierarchy[class] = append(
    oc.model.classHierarchy[class],
    superClass,
  )

  // Invalidate caches
  delete(oc.superClassCache, class)
  delete(oc.subClassCache, superClass)

  // Invalidate dependent property caches
  for propUID, meta := range oc.propertyMetadata {
    if meta.Domain == class || meta.Range == class {
      delete(oc.propertyMetadata, propUID)
    }
  }

  // Update timestamp
  oc.lastUpdated = time.Now()

  return nil
}
```

---

## Section 4: Query Optimization Integration

### 4.1 How Ontology Knowledge Improves Query Optimization

#### The Optimization Pipeline

```
SPARQL Query
    ↓
[Phase 1: AST → Algebra]  (Parser)
    ↓
[Phase 2: Algebraic Simplification]  (Optimizer)
    ↓
[Phase 3: Schema-Aware Optimization] ← ONTOLOGY ENTERS HERE
    ├─ Load ontology from cache
    ├─ Expand class hierarchies
    ├─ Apply domain/range constraints
    ├─ Reorder joins by selectivity
    └─ Apply reasoning rules
    ↓
[Phase 4: DQL Compilation]  (Translator)
    ↓
DQL Query → Dgraph Engine
```

#### Four Ways Ontology Improves Optimization

**1. Class Hierarchy Expansion**

Without ontology:

```sparql
SELECT ?x WHERE {
  ?x rdf:type ex:Person .
}
```

→ Searches only exact type "Person" → Misses subclasses: Student, Employee, etc.

With ontology:

```sparql
SELECT ?x WHERE {
  ?x rdf:type ex:Person .
}
```

→ Ontology knows: Person ⊇ {Student, Employee, Retired} → Query optimizer expands to:

```dql
{
  query(func: type(Person)) { ... }
  OR
  query(func: type(Student)) { ... }
  OR
  query(func: type(Employee)) { ... }
  OR
  query(func: type(Retired)) { ... }
}
```

**2. Domain/Range Constraint Propagation**

Without ontology:

```sparql
SELECT ?x ?y WHERE {
  ?x ex:hasFriend ?y .
}
```

→ Could bind to any ?x and ?y → Expensive join if many entities

With ontology:

```sparql
SELECT ?x ?y WHERE {
  ?x ex:hasFriend ?y .
}
```

→ Ontology defines: domain=Person, range=Person → Query optimizer rewrites to:

```dql
{
  query(func: type(Person)) @filter(has(hasFriend)) {
    hasFriend @filter(type(Person)) {
      uid
    }
  }
}
```

**3. Join Reordering by Selectivity**

Without ontology:

```sparql
SELECT ?x ?y ?z WHERE {
  ?x ex:hasFriend ?y .
  ?y ex:hasDegree ?z .
  ?z ex:value 5 .
}
```

→ Executes left-to-right → Might start with high-cardinality `hasFriend` join

With ontology:

```
hasFriend: 1M edges (Person → Person)
hasDegree: 10K edges (Person → Degree)  ← lower cardinality
value 5: 100 results  ← filter first (selectivity = 100/∞)
```

→ Query optimizer reorders:

```dql
{
  query(func: eq(value, 5)) {
    ~hasDegree {
      ~hasFriend {
        uid
      }
    }
  }
}
```

Executes rare filter first, then traverses back

**4. Reasoning Rule Application**

With ontology:

```sparql
SELECT ?x WHERE {
  ?x rdf:type ex:Mammal .
  ?x ex:likesMeat ?food .
}
```

If ontology includes rule: `Mammal ⊆ Animal` Query optimizer infers: results are also
`?x rdf:type Animal`

Can optimize downstream queries that filter on `Animal` type.

### 4.2 Example: Class Hierarchy Expansion

#### Input SPARQL Query

```sparql
PREFIX ex: <http://example.com/>
SELECT ?doc ?title WHERE {
  ?doc rdf:type ex:Publication .
  ?doc ex:title ?title .
}
```

#### Ontology Structure

```
Publication
├─ ConferencePaper
│  ├─ PeerReviewedConference
│  └─ WorkshopPaper
├─ JournalArticle
│  └─ SpecialIssue
└─ Thesis
   ├─ MasterThesis
   └─ PhDThesis
```

#### Optimization Steps

**Step 1: Load Ontology for class "Publication"**

```go
ontology := cache.GetOntology()
classUID := ontology.GetClassUID("http://example.com/Publication")

// Get all subclasses (transitive)
allClasses := ontology.GetTransitiveSubClasses(classUID)
// Returns: [Publication, ConferencePaper, ..., Thesis, MasterThesis, PhDThesis]
```

**Step 2: Expand to Union**

```dql
{
  var1 as var(func: type(Publication)) { uid }
  var2 as var(func: type(ConferencePaper)) { uid }
  var3 as var(func: type(JournalArticle)) { uid }
  var4 as var(func: type(Thesis)) { uid }
  // ... (and all transitive subclasses)

  query(func: uid(var1, var2, var3, var4)) @filter(has(title)) {
    uid
    title: ex_title
  }
}
```

**Step 3: Apply Filter Pushdown**

```dql
{
  query(func: type(Publication)) @filter(has(title)) {
    uid
    title: ex_title
  }

  query(func: type(ConferencePaper)) @filter(has(title)) {
    uid
    title: ex_title
  }

  // ... union of all subclass queries with filter
}
```

**Step 4: Order by Selectivity**

```
Statistics from ontology cache:
- Publication: 100,000 instances, 90,000 have title (selectivity 90%)
- ConferencePaper: 30,000 instances, 29,500 have title (selectivity 98%)
- Thesis: 5,000 instances, 4,500 have title (selectivity 90%)

Optimal order: Execute Thesis first (smallest result set), then Publication
```

#### Final Optimized DQL

```dql
{
  # Execute smallest set first
  theses as var(func: type(Thesis)) @filter(has(title)) { uid }

  # Execute largest set
  pubs as var(func: type(Publication)) @filter(has(title)) { uid }

  # Merge results
  query(func: uid(theses, pubs)) {
    uid
    title: ex_title
  }
}
```

### 4.3 Example: Property Range/Domain Constraints

#### Input SPARQL Query

```sparql
PREFIX ex: <http://example.com/>
PREFIX foaf: <http://xmlns.com/foaf/0.1/>
SELECT ?author ?name WHERE {
  ?book ex:hasAuthor ?author .
  ?author foaf:name ?name .
}
```

#### Ontology Metadata

```
hasAuthor property:
  domain: Book
  range: Person
  isFunctional: false
  isTransitive: false

foaf:name property:
  domain: Agent
  range: String
```

#### Optimization Steps

**Step 1: Domain/Range Constraint Injection**

Before optimization:

```dql
{
  query(func: has(hasAuthor)) {
    hasAuthor {
      uid
      name: foaf_name {
        uid
      }
    }
  }
}
```

After optimization (with constraints):

```dql
{
  # Filter to only objects with hasAuthor (books)
  query(func: type(Book)) @filter(has(hasAuthor)) {
    hasAuthor @filter(type(Person)) {
      uid
      # Range of hasAuthor is Person, so ?author is Person
      # Domain of foaf:name is Agent, Person ⊆ Agent ✓
      name: foaf_name {
        uid
      }
    }
  }
}
```

**Step 2: Type Validation**

```go
// During compilation phase, optimizer checks:
validator := NewConstraintValidator(ontology)

// Check: foaf:name can apply to Person (which is subclass of Agent)
if !validator.IsValidPropertyApplication("foaf:name", "Person") {
  return errors.New("foaf:name property cannot apply to Person")
}

// Passes: Person ⊆ Agent, and domain(foaf:name) = Agent
```

**Step 3: Selectivity Estimation**

```go
// Estimate: How many Books have hasAuthor?
bookCount := stats.GetClassInstanceCount("Book")           // 100,000
booksWithAuthor := stats.GetPropertyUsageCount("hasAuthor") // 95,000

// Filter selectivity: 95%
// => Start with type(Book) filter (smaller than has(hasAuthor))

// Estimate: Result size
avgAuthorsPerBook := 1.0  // hasAuthor is functional
resultSize := booksWithAuthor * avgAuthorsPerBook  // ~95,000
```

### 4.4 Integration with Phase 3 Compilation

#### Compiler API for Ontology

```go
type CompilerContext struct {
  // Existing fields
  graphQL *schema.Schema
  acl     *authorization.Checker

  // New ontology fields
  ontology        *OntologyModel          // Loaded ontology
  ontologyCache   *OntologyCache          // URI/UID lookups
  reasoningRules  []*ReasoningRule        // Inference rules
  statisticsCache *StatisticsCache        // Cardinality data
}

type QueryOptimizer interface {
  // Existing method
  OptimizeAlgebra(algebra *Algebra) (*Algebra, error)

  // New ontology-aware method
  ApplyOntologyOptimizations(algebra *Algebra, ctx *CompilerContext) (*Algebra, error)
}

type ReasoningRule interface {
  // Apply rule to expand triple patterns
  ApplyRule(pattern *TriplePattern, ctx *CompilerContext) []*TriplePattern

  // Get priority (higher = apply first)
  Priority() int
}
```

#### Optimization Hook Points

```go
func (o *Optimizer) CompileToSubGraph(algebra *Algebra, ctx *CompilerContext) (*pb.SubGraph, error) {
  // 1. Load ontology if not already loaded
  if ctx.ontology == nil {
    ont, err := o.loadOntology(ctx)
    if err != nil {
      return nil, err
    }
    ctx.ontology = ont
  }

  // 2. Apply algebra simplifications (Phase 2)
  simplified := o.simplifyAlgebra(algebra)

  // 3. Apply ontology optimizations (Phase 3)
  optimized := o.applyOntologyOptimizations(simplified, ctx)

  // 4. Apply join reordering
  reordered := o.reorderJoins(optimized, ctx.statisticsCache)

  // 5. Compile to SubGraph
  subgraph := o.compileToSubGraph(reordered)

  // 6. Inject ACL and authorization
  withAuth := o.applyAuthorization(subgraph, ctx.acl)

  return withAuth, nil
}

func (o *Optimizer) applyOntologyOptimizations(algebra *Algebra, ctx *CompilerContext) *Algebra {
  for _, rule := range ctx.reasoningRules {
    algebra = rule.ApplyRule(algebra, ctx)
  }

  for _, pattern := range algebra.TriplePatterns {
    o.expandClassHierarchy(pattern, ctx)
    o.applyPropertyConstraints(pattern, ctx)
  }

  return algebra
}
```

#### Concrete Optimization Pipeline

```go
// Step 1: Identify triple patterns with class matching
func (o *Optimizer) expandClassHierarchy(pattern *TriplePattern, ctx *CompilerContext) {
  if pattern.Predicate == "rdf:type" {
    classURI := pattern.Object
    classUID := ctx.ontologyCache.GetClassUID(classURI)

    // Get all subclasses
    subclasses := ctx.ontology.GetTransitiveSubClasses(classUID)

    // Expand to multiple patterns
    for _, subclass := range subclasses {
      newPattern := &TriplePattern{
        Subject:   pattern.Subject,
        Predicate: "rdf:type",
        Object:    ctx.ontologyCache.GetClassURI(subclass),
      }
      pattern.alternativePatterns = append(pattern.alternativePatterns, newPattern)
    }
  }
}

// Step 2: Apply domain/range constraints
func (o *Optimizer) applyPropertyConstraints(pattern *TriplePattern, ctx *CompilerContext) {
  propUID := ctx.ontologyCache.GetPropertyUID(pattern.Predicate)
  meta := ctx.ontology.GetPropertyMetadata(propUID)

  if meta.Domain != 0 {
    // Subject must be instance of domain class
    pattern.subjectConstraint = &TypeConstraint{
      Type: meta.Domain,
      isRequired: true,
    }
  }

  if meta.Range != 0 {
    // Object must be instance of range class
    pattern.objectConstraint = &TypeConstraint{
      Type: meta.Range,
      isRequired: true,
    }
  }
}

// Step 3: Estimate selectivity
func (o *Optimizer) estimateSelectivity(pattern *TriplePattern, ctx *CompilerContext) float64 {
  if pattern.hasObjectConstraint {
    domainInstances := ctx.statisticsCache.GetClassInstanceCount(pattern.subjectConstraint.Type)
    usageCount := ctx.statisticsCache.GetPropertyUsageCount(pattern.Predicate)
    return float64(usageCount) / float64(domainInstances)
  }
  return 1.0  // No constraint, worst case
}

// Step 4: Reorder joins by selectivity
func (o *Optimizer) reorderJoins(algebra *Algebra, stats *StatisticsCache) *Algebra {
  patterns := algebra.TriplePatterns

  // Sort by selectivity (ascending - execute rare first)
  sort.Slice(patterns, func(i, j int) bool {
    selectivityI := o.estimateSelectivity(patterns[i], ctx)
    selectivityJ := o.estimateSelectivity(patterns[j], ctx)
    return selectivityI < selectivityJ
  })

  algebra.TriplePatterns = patterns
  return algebra
}
```

### 4.5 Performance Implications

#### Expected Optimizations

**Execution Time Improvements**:

1. **Class hierarchy expansion**: 10-20% faster (avoids missing results in subclasses)
2. **Join reordering**: 20-50% faster (rare filters first)
3. **Property constraints**: 10-30% faster (reduced cardinality in joins)
4. **Reasoning rules**: 5-15% faster (materialized inferences)

**Memory Improvements**:

1. **Intermediate result caching**: 10-20% less memory (reuse subquery results)
2. **Index usage**: 20-40% faster (guided by ontology)
3. **Predicate filtering**: 5-10% less memory (early elimination)

#### Cost-Benefit Analysis

**Costs**:

- Ontology loading: ~50-500ms (depends on size, cached)
- Optimization overhead: ~10-20ms per query (negligible)
- Memory footprint: ~10-100MB (typical ontologies)

**Benefits**:

- Query correctness: Ensures all relevant subclasses included
- Performance: 20-50% execution time reduction on average
- Scalability: Enables query optimization on large graphs

**Break-even Point**: Query optimization breaks even after ~2-3 queries

---

## Section 5: Implementation Plan

### 5.1 Phase 5 File Structure

```
sparql/ontology/
├── model.go              # Core ontology data structures
├── loader.go             # OWL/RDFS file loading
├── loader_test.go
├── schema_builder.go     # Dgraph schema synthesis
├── schema_builder_test.go
├── reasoning.go          # RDFS reasoning rules
├── reasoning_test.go
├── cache.go              # Ontology caching layer
├── cache_test.go
├── validator.go          # Semantic validation
├── validator_test.go
├── examples/
│   ├── foaf.ttl          # FOAF ontology example
│   ├── person.owl        # OWL example
│   └── README.md
└── README.md

sparql/
├── compiler.go           # Updated to integrate ontology
├── compiler_test.go
└── ontology_integration_test.go  # Integration tests
```

### 5.2 Detailed Breakdown of Each File

#### sparql/ontology/model.go (400 lines)

**Responsibilities**:

- Define core ontology data structures
- Represent RDF/RDFS/OWL concepts as Go types
- Provide in-memory model for query optimization

**Key Types**:

```go
type OntologyModel struct {
  Classes              map[string]*Class
  Properties           map[string]*Property
  Restrictions         map[string]*Restriction
  Hierarchies          map[string][]string
  NamespaceMap         map[string]string
  Metadata             ModelMetadata
}

type Class struct {
  URI                  string
  Label                string
  Comment              string
  SuperClasses         []string
  SubClasses           []string
  EquivalentClasses    []string
  PropertiesWithDomain []*Property
  DgraphUID            uint64
}

type Property struct {
  URI                  string
  Label                string
  Comment              string
  Domain               string
  Range                string
  SuperProperties      []string
  SubProperties        []string
  InverseProperty      string
  IsFunctional         bool
  IsTransitive         bool
  DgraphUID            uint64
}

type Restriction struct {
  URI                  string
  OnProperty           string
  MinCardinality       int
  MaxCardinality       int
  SomeValuesFrom       string
  AllValuesFrom        string
  DgraphUID            uint64
}

type ModelMetadata struct {
  LoadedAt             time.Time
  Source               string
  Version              string
  NamespacePrefix      string
  Statistics           ModelStats
}

type ModelStats struct {
  TotalClasses         int
  TotalProperties      int
  TotalRelationships   int
  MaxHierarchyDepth    int
}
```

**Key Methods**:

- `GetClass(uri string) (*Class, error)` - lookup by URI
- `GetProperty(uri string) (*Property, error)`
- `GetTransitiveSuperClasses(uri string) []string` - hierarchy traversal
- `GetTransitiveSubClasses(uri string) []string`
- `GetPropertiesWithDomain(classURI string) []*Property`
- `GetPropertiesWithRange(classURI string) []*Property`
- `Validate() []ValidationError` - semantic checks

#### sparql/ontology/loader.go (600 lines)

**Responsibilities**:

- Parse RDF/OWL files in multiple formats
- Build OntologyModel from triples
- Handle errors and validation

**Key Types**:

```go
type Loader struct {
  graph              *rdf.Graph
  namespaceManager   *NamespaceManager
  uriNormalization   bool
  circularCheckDepth int
}

type LoadOptions struct {
  Format            RDFFormat
  Namespace         string  // For multi-tenancy
  ValidateOnLoad    bool
  SkipInvalidTerms  bool
}

type RDFFormat int
const (
  FormatNTriples RDFFormat = iota
  FormatTurtle
  FormatRDFXml
  FormatJSONLD
)

type Triple struct {
  Subject   string
  Predicate string
  Object    string
  Language  string  // For lang tags
  DataType  string  // For typed literals
}

type NamespaceManager struct {
  prefixMap  map[string]string
  reverseMap map[string]string
}
```

**Key Methods**:

- `LoadFromFile(path string, opts LoadOptions) (*OntologyModel, error)`
- `LoadFromReader(reader io.Reader, opts LoadOptions) (*OntologyModel, error)`
- `ParseTriples(reader io.Reader) ([]*Triple, error)`
- `BuildModel(triples []*Triple) (*OntologyModel, error)`
- `NormalizeURI(uri string) string`
- `ResolveNamespace(prefixed string) string`

**Algorithm for BuildModel**:

```
1. First pass: Identify all subjects and their types
2. Classify as Class if: rdf:type = rdfs:Class or owl:Class
3. Classify as Property if: rdf:type = rdf:Property or owl:ObjectProperty
4. Second pass: Build relationships (subClassOf, domain, range, etc.)
5. Third pass: Validate and compute derived data
6. Return populated OntologyModel
```

#### sparql/ontology/schema_builder.go (500 lines)

**Responsibilities**:

- Convert OntologyModel to Dgraph schema representation
- Synthesize GraphQL types from ontology
- Generate inference rules

**Key Types**:

```go
type SchemaBuilder struct {
  ontology      *OntologyModel
  graphqlSchema *graphql.Schema
  namespace     string
}

type SynthesizedSchema struct {
  Types          []*SynthesizedType
  Predicates     []*SynthesizedPredicate
  InferenceRules []*InferenceRule
}

type SynthesizedType struct {
  Name             string
  OntologyClass    string
  SuperTypes       []string
  Fields           []*SynthesizedField
  DgraphTypeName   string
}

type SynthesizedField struct {
  Name             string
  OntologyProperty string
  Type             string
  IsRequired       bool
  IsReverse        bool
  DgraphPredicate  string
}

type SynthesizedPredicate struct {
  Name             string
  ValueType        pb.Posting_ValType
  List             bool
  Index            []string
  Reverse          bool
  Metadata         map[string]string
}

type InferenceRule struct {
  Name             string
  Trigger          *TriplePattern
  Action           func(*TriplePattern) []*TriplePattern
  Priority         int
}
```

**Key Methods**:

- `BuildFromOntology() (*SynthesizedSchema, error)` - full synthesis
- `EnhanceExistingSchema(existing *pb.Schema) (*pb.Schema, error)` - augment user schema
- `GenerateInferenceRules() ([]*InferenceRule, error)` - reasoning rules
- `ValidateSchemaConsistency() []ValidationError`

**Synthesis Algorithm**:

```
1. For each ontology Class:
   - Create Dgraph type
   - Create rdf:type predicate if needed
   - Set @index for URI lookups

2. For each Property:
   - Create predicate
   - Set domain/range type constraints
   - Add indexes for efficient querying
   - Mark as @reverse if has inverse

3. For each Hierarchy:
   - Create subClassOf/subPropertyOf predicates
   - Mark as uid type with @reverse

4. Generate reasoning rules:
   - Class hierarchy inference
   - Property constraint checking
   - Transitivity rules
```

#### sparql/ontology/reasoning.go (400 lines)

**Responsibilities**:

- Implement RDFS and basic OWL reasoning
- Apply inference rules during query optimization
- Manage materialized inference results

**Key Types**:

```go
type ReasoningEngine struct {
  ontology *OntologyModel
  rules    []*ReasoningRule
  cache    *InferenceCache
}

type ReasoningRule struct {
  Name             string
  Description      string
  Apply            func(*Query, *OntologyModel) (*Query, error)
  Priority         int
  Category         RuleCategory  // RDFS vs OWL
}

type RuleCategory int
const (
  RuleRDFS RuleCategory = iota
  RuleOWL
  RuleCustom
)

type InferenceResult struct {
  Inferred  []*Triple      // New facts derived
  Cached    bool           // Whether from cache
  Duration  time.Duration
}
```

**Key Methods**:

- `ApplyRules(query *dql.GraphQuery) (*dql.GraphQuery, error)`
- `ExpandClassHierarchy(classURI string) []string`
- `ExpandPropertyHierarchy(propURI string) []string`
- `ApplyDomainConstraints(pattern *TriplePattern) *TriplePattern`
- `ApplyRangeConstraints(pattern *TriplePattern) *TriplePattern`

**Built-in Rules** (RDFS only):

```go
// Rule 1: Class hierarchy expansion
// If C1 subClassOf C2 and ?x rdf:type C1, then ?x rdf:type C2
rdfsc1 := &ReasoningRule{
  Name: "RDFS Class 1",
  Apply: func(q *Query, o *OntologyModel) (*Query, error) {
    // Expand type filter with superclasses
    return expandTypeFilter(q, o), nil
  },
  Priority: 100,
}

// Rule 2: Property domain/range
// If prop domain C and ?x prop ?y, then ?x rdf:type C
rdfsc2 := &ReasoningRule{
  Name: "RDFS Domain",
  Apply: func(q *Query, o *OntologyModel) (*Query, error) {
    // Add type constraints from property domain
    return applyDomainConstraints(q, o), nil
  },
  Priority: 90,
}

// Rule 3: Property hierarchy
// If P1 subPropertyOf P2 and ?x P1 ?y, then ?x P2 ?y
rdfsc5 := &ReasoningRule{
  Name: "RDFS Property Hierarchy",
  Apply: func(q *Query, o *OntologyModel) (*Query, error) {
    // Expand property predicates with superproperties
    return expandPropertyFilter(q, o), nil
  },
  Priority: 80,
}
```

#### sparql/ontology/cache.go (350 lines)

**Responsibilities**:

- Manage ontology caching at multiple levels
- Handle cache invalidation
- Provide efficient lookups

**Key Types**:

```go
type CacheManager struct {
  ontologyCache    *OntologyCache
  compilerCache    *CompilerCache
  statisticsCache  *StatisticsCache
  dgraphClient     *Client
  ttl              time.Duration
  mutex            sync.RWMutex
}

type OntologyCache struct {
  model           *OntologyModel
  classURIToUID   map[string]uint64
  propertyURIToUID map[string]uint64
  superClassCache  map[uint64][]uint64
  subClassCache    map[uint64][]uint64
  lastUpdated      time.Time
  stats            CacheStats
}

type CompilerCache struct {
  classExpansions  map[string][]uint64
  propertyMeta     map[uint64]*PropertyMetadata
  validAt          time.Time
}

type StatisticsCache struct {
  classInstanceCounts map[uint64]int64
  propertyFrequency   map[uint64]float64
  lastRefreshed       time.Time
}

type CacheStats struct {
  Hits              int64
  Misses            int64
  Invalidations     int64
  LastHitRate       float64
}
```

**Key Methods**:

- `GetOntology() (*OntologyModel, error)` - with caching
- `GetClassUID(uri string) (uint64, bool)` - O(1) with cache
- `GetPropertyUID(uri string) (uint64, bool)`
- `GetSuperClasses(classUID uint64) []uint64` - cached transitive
- `InvalidateClass(classUID uint64)` - cascade invalidation
- `RefreshStatistics(force bool) error`
- `CacheStats() *CacheStats`

#### sparql/ontology/validator.go (300 lines)

**Responsibilities**:

- Validate ontology semantic correctness
- Check for conflicts with existing schema
- Warn on inconsistencies

**Key Types**:

```go
type Validator struct {
  ontology *OntologyModel
  schema   *pb.Schema
}

type ValidationError struct {
  Severity  ErrorSeverity
  Code      string
  Message   string
  Affected  []string  // URIs involved
  Line      int       // Source line if from file
  Fix       string    // Suggested fix
}

type ErrorSeverity int
const (
  ErrorSeverityError ErrorSeverity = iota
  ErrorSeverityWarning
  ErrorSeverityInfo
)

type ValidationReport struct {
  Errors   []ValidationError
  Warnings []ValidationError
  Passed   bool
  Duration time.Duration
}
```

**Key Methods**:

- `ValidateSyntax() *ValidationReport` - Check RDF/OWL syntax
- `ValidateSemantics() *ValidationReport` - Check logical consistency
- `ValidateConstraints() *ValidationReport` - Check OWL constraints
- `CheckCircularHierarchy() []ValidationError`
- `CheckUndefinedReferences() []ValidationError`
- `CheckSchemaConflicts(schema *pb.Schema) []ValidationError`
- `GenerateReport() *ValidationReport`

### 5.3 Dependencies and Integration Points

#### Internal Dependencies

```
loader.go
  ↓ uses
model.go

schema_builder.go
  ↓ uses
model.go

reasoning.go
  ↓ uses
model.go, cache.go

cache.go
  ↓ uses
model.go

validator.go
  ↓ uses
model.go
```

#### External Dependencies

```
sparql/ontology/
  ↓ integrates with
graphql/schema/          (Type/Field definitions)
dql/parser.go            (GraphQuery structures)
posting/                 (Storage layer)
query/                   (Query execution)
schema/                  (Type definitions)
edgraph/                 (Server integration)
```

#### Integration Points

**1. Server Integration** (`edgraph/server.go`):

```go
// Load ontology at startup
func (s *Server) loadOntology(ctx context.Context) error {
  path := s.Config.OntologyPath
  if path == "" {
    return nil  // Optional
  }

  loader := ontology.NewLoader()
  model, err := loader.LoadFromFile(path, ontology.LoadOptions{
    Format: ontology.FormatTurtle,
    Namespace: s.Config.Namespace,
  })
  if err != nil {
    return err
  }

  s.ontologyCache = ontology.NewCacheManager(model, s.Client)
  return nil
}
```

**2. Query Compiler Integration** (`sparql/compiler.go`):

```go
func (c *Compiler) CompileQuery(sparql string) (*pb.SubGraph, error) {
  // ... existing parsing ...

  // Pass ontology to optimizer
  optCtx := &OptimizationContext{
    Ontology: c.ontologyCache.GetOntology(),
    Stats: c.ontologyCache.GetStatistics(),
  }

  return c.optimizer.OptimizeWithOntology(algebra, optCtx)
}
```

**3. Schema Synthesis** (`graphql/schema/schemagen.go`):

```go
// Enhance generated schema with ontology
func genDgSchema(gqlSch *ast.Schema, ontology *OntologyModel) string {
  builder := ontology.NewSchemaBuilder(ontology, gqlSch)
  synth, err := builder.EnhanceExistingSchema(gqlSch)
  if err != nil {
    log.Warning("ontology schema synthesis failed: %v", err)
    return original  // Fall back
  }

  return synth.String()
}
```

**4. ACL Integration** (`graphql/schema/auth.go`):

```go
// Apply ontology-based authorization
func applyOntologyAuth(rules *AuthRules, ontology *OntologyModel) error {
  for className, classAuth := range ontology.GetAuthRules() {
    rules.AddClassRule(className, classAuth)
  }
  return nil
}
```

### 5.4 How to Test Ontology Loading

#### Unit Tests (loader_test.go)

```go
func TestLoadNTriples(t *testing.T) {
  loader := ontology.NewLoader()
  model, err := loader.LoadFromReader(
    strings.NewReader(`
<http://ex.com/Person> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/2000/01/rdf-schema#Class> .
<http://ex.com/hasFriend> <http://www.w3.org/1999/02/22-rdf-syntax-ns#type> <http://www.w3.org/1999/02/22-rdf-syntax-ns#Property> .
    `),
    ontology.LoadOptions{Format: ontology.FormatNTriples},
  )

  require.NoError(t, err)
  require.NotNil(t, model)
  require.Len(t, model.Classes, 1)
  require.Len(t, model.Properties, 1)
}

func TestHierarchyTraversal(t *testing.T) {
  // Load ontology with hierarchy
  model := loadTestOntology(t, "testdata/hierarchy.ttl")

  // Test transitive superclasses
  superclasses := model.GetTransitiveSuperClasses("Student")
  require.Contains(t, superclasses, "Person")
  require.Contains(t, superclasses, "Agent")
}

func TestPropertyConstraints(t *testing.T) {
  model := loadTestOntology(t, "testdata/foaf.ttl")

  hasFriend := model.GetProperty("http://xmlns.com/foaf/0.1/knows")
  require.NotNil(t, hasFriend)
  require.Equal(t, "Person", hasFriend.Domain)
  require.Equal(t, "Person", hasFriend.Range)
}
```

#### Integration Tests (ontology_integration_test.go)

```go
func TestOntologyQueryOptimization(t *testing.T) {
  // Setup Dgraph with ontology
  client := setupTestDgraph(t)
  defer teardownTestDgraph(t, client)

  // Load ontology
  cache := ontology.NewCacheManager(...)
  ontModel := cache.GetOntology()

  // Prepare query with ontology optimization
  query := `
SELECT ?x WHERE {
  ?x rdf:type ex:Person .
}
  `

  // Compile with ontology
  compiler := NewCompiler(cache)
  dql, err := compiler.Compile(query)

  require.NoError(t, err)
  // Should include subclass expansion
  require.Contains(t, dql, "Student")
  require.Contains(t, dql, "Employee")
}

func TestOntologyLoading(t *testing.T) {
  client := setupTestDgraph(t)

  // Load ontology file
  loader := ontology.NewLoader()
  model, err := loader.LoadFromFile("testdata/foaf.owl", ontology.LoadOptions{
    Format: ontology.FormatRDFXml,
  })

  require.NoError(t, err)

  // Write to Dgraph
  writer := ontology.NewDgraphWriter(client)
  err = writer.Write(model)

  require.NoError(t, err)

  // Query back
  result := queryOntologyClass(client, "http://xmlns.com/foaf/0.1/Person")
  require.NotNil(t, result)
}

func TestCircularHierarchyDetection(t *testing.T) {
  triples := []*ontology.Triple{
    {S: "A", P: "rdfs:subClassOf", O: "B"},
    {S: "B", P: "rdfs:subClassOf", O: "C"},
    {S: "C", P: "rdfs:subClassOf", O: "A"},  // Circular!
  }

  validator := ontology.NewValidator()
  errs := validator.ValidateHierarchy(triples)

  require.Len(t, errs, 1)
  require.Contains(t, errs[0].Code, "CIRCULAR")
}
```

### 5.5 Success Criteria Per Component

#### loader.go Success Criteria

- [ ] Parses N-Triples format without errors
- [ ] Parses Turtle format with namespace resolution
- [ ] Parses RDF/XML format correctly
- [ ] Loads files up to 100MB (ontology benchmarks)
- [ ] Performance: <500ms for typical ontologies
- [ ] Circular hierarchy detection
- [ ] Undefined reference detection
- [ ] Multi-tenancy support via namespace filtering

#### model.go Success Criteria

- [ ] Represents all RDFS concepts (Classes, Properties, Hierarchies)
- [ ] Represents basic OWL concepts (Equivalence, Restrictions)
- [ ] O(1) lookup by URI
- [ ] Transitive closure computation correct
- [ ] Serializable to/from JSON
- [ ] Memory footprint <100MB for typical ontologies

#### schema_builder.go Success Criteria

- [ ] Generates valid Dgraph schema from ontology
- [ ] Handles conflicts between user schema and ontology
- [ ] Creates appropriate indexes for efficient querying
- [ ] Computes inference rules for reasoning
- [ ] Validates schema consistency

#### reasoning.go Success Criteria

- [ ] Class hierarchy expansion works for queries
- [ ] Domain/range constraints applied correctly
- [ ] Property hierarchy expansion works
- [ ] Inverse property reasoning works
- [ ] Performance: <10ms rule application per query

#### cache.go Success Criteria

- [ ] Caches ontology model correctly
- [ ] Cache TTL invalidation works
- [ ] Event-based invalidation on updates
- [ ] Cascade invalidation for affected entries
- [ ] Cache hit rate >95% for repeated queries
- [ ] Multi-level caching with correct semantics

#### validator.go Success Criteria

- [ ] Detects circular hierarchies
- [ ] Detects undefined references
- [ ] Detects type inconsistencies
- [ ] Warns on schema conflicts
- [ ] Produces clear error messages
- [ ] Validates all RDFS semantics

---

## Section 6: Reuse & Integration Decisions

### 6.1 Which Dgraph Abstractions to Reuse

#### Pattern 1: Type Interface Pattern (graphql/schema/wrappers.go)

**Decision: REUSE with Extension**

```go
// Current GraphQL Type interface
type Type interface {
  Field(name string) FieldDefinition
  Fields() []FieldDefinition
  IDField() FieldDefinition
  Name() string
  DgraphName() string
}

// Extend for ontology
type OntologyType interface {
  Type  // Embed existing interface

  // New methods for ontology
  OntologyURI() string
  SuperTypes() []Type
  SubTypes() []Type
  IsOntologyDefined() bool
}

// Implement for ontology classes
type ontologyType struct {
  ontologyClass *OntologyClass
  graphqlType   Type  // Compose with existing type
}

func (ot *ontologyType) Field(name string) FieldDefinition {
  // Delegate to graphqlType
  return ot.graphqlType.Field(name)
}

func (ot *ontologyType) OntologyURI() string {
  return ot.ontologyClass.URI
}

func (ot *ontologyType) SuperTypes() []Type {
  // Build Type objects from ontology superclasses
  ...
}
```

**Reasoning**:

- GraphQL type system already well-designed
- Ontology classes can implement same interface
- Allows gradual adoption (some types from GraphQL, some from ontology)
- Query compiler doesn't need to know source of type

**Reuse in Compiler**:

```go
// Query compiler uses Type interface
func compileFieldAccess(t Type, field string) *dql.GraphQuery {
  fd := t.Field(field)
  // Works for both GraphQL and ontology types
  return buildQuery(fd.DgraphPredicate())
}
```

#### Pattern 2: Predicate Mapping Pattern (graphql/schema/gqlschema.go)

**Decision: REUSE with Extension**

Current pattern:

```go
type schema struct {
  dgraphPredicate map[string]map[string]string  // typeName→fieldName→predicate
}
```

Extended pattern:

```go
type schema struct {
  // Existing
  dgraphPredicate map[string]map[string]string

  // New for ontology
  ontologyPredicate map[string]map[string]string  // ontologyClass→prop→dgraphPred

  // URI to local predicate name
  ontologyURIToPredicate map[string]string
}

func (s *schema) GetPredicate(typeName, fieldName string) string {
  // Try GraphQL first
  if pred, ok := s.dgraphPredicate[typeName][fieldName]; ok {
    return pred
  }

  // Try ontology
  if pred, ok := s.ontologyPredicate[typeName][fieldName]; ok {
    return pred
  }

  return ""
}
```

**Reasoning**:

- Predicate mapping is proven pattern in Dgraph
- Extends naturally to ontology predicates
- No need to reinvent type→predicate resolution

#### Pattern 3: Posting List Pattern (posting/list.go)

**Decision: REUSE as-is**

Ontology data stored as regular Dgraph predicates:

```
ontology_class_uri: string @index(exact) .
ontology_class_label: string @index(term) .
ontology_class_super: uid @reverse .
```

Benefits:

- Leverage existing posting list implementation
- Use existing indexes (exact, term, fulltext, etc.)
- Participate in transactions automatically
- Benefit from compression and optimization

**Reasoning**:

- Ontology is just more RDF data
- No need for special storage layer
- Automatic durability and replication

#### Pattern 4: Index System Pattern (posting/index.go)

**Decision: REUSE heavily**

```
// Indexes on ontology predicates
ontology_class_uri: string @index(exact) .
  → Fast lookup by class URI

ontology_class_label: string @index(term) .
  → Full-text search on class labels

ontology_property_uri: string @index(exact) .
  → Fast lookup by property URI

ontology_class_super: uid @reverse .
  → Efficient hierarchy traversal
```

**Reasoning**:

- Index system already optimized
- No special index types needed
- Exact+term indexes sufficient for ontology queries

#### Pattern 5: Statistics Pattern (query execution)

**Decision: EXTEND for ontology**

Current pattern: Cardinality estimated per predicate during execution

```go
type statistic struct {
  predicate string
  count     int64
  indexed   bool
}
```

Extended pattern:

```go
type OntologyStatistics struct {
  // Class instance counts
  classInstanceCounts map[string]int64

  // Property usage patterns
  propertyUsageFrequency map[string]float64

  // Join selectivity
  joinSelectivity map[string]float64

  // Last computed
  computedAt time.Time
}
```

**Reasoning**:

- Reuse cardinality estimation infrastructure
- Add ontology-specific statistics
- Used during query optimization phase

#### Pattern 6: Authorization Pattern (graphql/schema/auth.go)

**Decision: EXTEND for ontology**

Current: ACL rules applied to types/fields

Extended: Add ontology class/property level rules

```go
type AuthRule struct {
  // Existing
  Type string
  Field string

  // New for ontology
  OntologyClass string
  OntologyProperty string
}

func (s *schema) CanAccessOntologyClass(classURI string, user *User) bool {
  class := s.ontology.GetClass(classURI)
  rule := s.authRules.GetRuleForClass(classURI)
  return rule.IsAllowed(user)
}
```

**Reasoning**:

- Auth rules are critical for data governance
- Ontology access should follow same patterns
- Type system already has auth rules
- Extend naturally to ontology classes/properties

### 6.2 Which Dgraph Code to Study Deeper

**High Priority** (required for Phase 5):

1. `graphql/schema/wrappers.go` - Type/FieldDefinition interfaces (200 lines)
2. `graphql/schema/gqlschema.go` - Schema synthesis (300 lines critical parts)
3. `dql/parser.go` - GraphQuery structure (100 lines)
4. `posting/list.go` - Posting list structure (100 lines)
5. `query/query.go` - Query execution entry points (100 lines)

**Medium Priority** (helpful for optimization):

1. `graphql/schema/schemagen.go` - Schema generation logic (200 lines)
2. `query/query.go` - Join ordering logic (200 lines)
3. `worker/task.go` - Task execution and cardinality (150 lines)

**Lower Priority** (reference only):

1. `edgraph/server.go` - Server integration points
2. `edgraph/query.go` - Query handler flow
3. `schema/parse.go` - Schema parsing

### 6.3 Potential Conflicts or Challenges

#### Challenge 1: Type Name Conflicts

**Problem**: User-defined GraphQL types might have same name as ontology classes

**Solution**:

```go
// Prefix ontology types
ontologyTypePrefix := "Ont_"
ontologyFieldPrefix := "ont_"

// Example: ontology class "Person" → "Ont_Person"
// User type "Person" remains as-is
// Conflict resolved in schema handler
```

#### Challenge 2: Predicate Naming

**Problem**: Ontology predicates must not conflict with user predicates

**Solution**:

```go
// Namespace ontology predicates
func ontologyPredicate(prop string) string {
  return fmt.Sprintf("ontology_%s", strings.ToLower(prop))
}

// Example: ontology property "hasFriend" → "ontology_hasfriend"
```

#### Challenge 3: URI Encoding

**Problem**: Ontology URIs contain special characters, Dgraph predicates don't

**Solution**:

```go
// Hash URIs for storage, keep mapping
func encodeURI(uri string) string {
  h := fnv.New64a()
  h.Write([]byte(uri))
  return fmt.Sprintf("uri_%x", h.Sum64())
}

// Store URI in string field for reverse lookup
// ontology_class_uri: string @index(exact)  // Stores full URI
// Query by exact match on URI
```

#### Challenge 4: Query Optimization Performance

**Problem**: Ontology expansion could slow down query compilation

**Solution**:

```go
// Cache expansion results
type ExpansionCache struct {
  expansions map[string][]uint64  // classURI → expanded UIDs
  validUntil time.Time
}

// Limit expansion depth
maxExpansionDepth := 10
maxExpansionWidth := 1000  // Max classes in union
```

#### Challenge 5: Circular Dependencies in Reasoning

**Problem**: Reasoning rules might not terminate

**Solution**:

```go
// Track applied rules per pattern
type RuleApplication struct {
  pattern string
  appliedRules set[string]
  depth int
}

func (ra *RuleApplication) CanApplyRule(rule string) bool {
  if ra.depth > 10 {
    return false  // Max depth
  }
  if ra.appliedRules.Contains(rule) {
    return false  // Already applied
  }
  return true
}
```

### 6.4 Recommended Approach

#### Phase 5 Strategy: Gradual Integration

**Phase 5.1: Core Infrastructure** (1 week)

- Implement `model.go` and `loader.go`
- Test loading various ontology formats
- Build ontology validation
- No server integration yet

**Phase 5.2: Storage Integration** (1 week)

- Write ontology to Dgraph as predicates
- Build cache layer
- Test cache invalidation
- Verify index performance

**Phase 5.3: Schema Synthesis** (1 week)

- Implement `schema_builder.go`
- Generate schema from ontology
- Handle conflicts with existing schema
- Test end-to-end

**Phase 5.4: Query Optimization** (1 week)

- Implement `reasoning.go`
- Add class hierarchy expansion
- Add property constraint propagation
- Test query optimization

**Phase 5.5: Integration Testing** (1 week)

- End-to-end tests
- Performance benchmarks
- Documentation
- Examples

#### Technology Choices

**RDF Parsing Library**:

- Use `rdf2go` or `github.com/deiu/rdf2go` (if available)
- Or write minimal custom N-Triples parser (100 lines)
- Rationale: Minimize external dependencies, ontologies typically small

**Caching Library**:

- Use simple map with RWMutex (no external cache library)
- Manual TTL management
- Rationale: Simple TTL, no complex eviction policies needed

**Validation**:

- Custom validation logic in `validator.go`
- Don't rely on SPARQL/OWL validators
- Rationale: Only need subset of validation, lightweight

---

## Section 7: Open Questions & Future Extensions

### 7.1 What Needs More Design Before Implementation

#### Question 1: Ontology Update Handling

**Status**: Partially designed

**Open Issues**:

- How to handle ontology updates without reloading entire cache?
- Should updates be transactional?
- How to handle in-flight queries with stale ontology?

**Proposed Solution**:

```go
type OntoloagUpdate struct {
  Timestamp time.Time
  Changes   []*Change
  Atomic    bool
}

// Query picks snapshot at start time
// Updates queued until all queries using old version complete
```

**Design Needed**: Version/timestamp semantics

#### Question 2: Ontology Reasoning Scope

**Status**: Partially designed

**Open Issues**:

- Which RDFS/OWL rules to implement in Phase 5?
- Which to defer to Phase 6?
- How to control reasoning overhead?

**Proposed Solution**:

```go
type ReasoningLevel int
const (
  ReasoningNone ReasoningLevel = iota
  ReasoningRDFSMinimal  // Only basic rules
  ReasoningRDFSFull     // All RDFS rules
  ReasoningOWLBasic     // OWL subset
)
```

**Design Needed**: Rule priority system and performance budgets

#### Question 3: Multi-Tenancy and Namespaces

**Status**: High-level design only

**Open Issues**:

- How to isolate ontologies per tenant?
- How to share ontologies across tenants?
- How to handle namespace conflicts?

**Proposed Solution**:

```
Prefix all ontology predicates with tenant ID:
  [tenant-id]:ontology_class_uri
  [tenant-id]:ontology_property_domain

Keep master ontology index:
  ontology_registry: [
    { tenantID: string, ontologyURI: string, loadedAt: time.Time }
  ]
```

**Design Needed**: Tenant isolation strategy and cross-tenant queries

#### Question 4: Ontology Versioning

**Status**: Not designed

**Open Issues**:

- How to version ontologies?
- How to compare versions?
- How to support multiple ontology versions?

**Proposed Solution** (sketch):

```
ontology_version: string @index(exact) .
ontology_supersedes: uid .

Example:
  FOAF 0.1 (id: uid1)
  FOAF 0.2 (id: uid2, supersedes: uid1)
  FOAF 0.3 (id: uid3, supersedes: uid2)
```

**Design Needed**: Version compatibility and migration strategies

#### Question 5: Custom Reasoners

**Status**: Not designed

**Open Issues**:

- How to support custom reasoning rules (SWRL, etc)?
- How to compose multiple reasoners?
- How to control reasoner performance?

**Proposed Solution** (sketch):

```go
type CustomReasoner interface {
  ApplyRules(query *GraphQuery) (*GraphQuery, error)
  GetRules() []*ReasoningRule
  GetCost() float64  // Performance budget
}
```

**Design Needed**: Plugin architecture and rule composition

### 7.2 Can OWL Reasoning Go Beyond RDFS

#### Current Proposal: RDFS-Only in Phase 5

**Reasoning**:

1. RDFS covers 80% of ontology use cases
2. OWL reasoning adds 5-10x complexity
3. OWL requires DL reasoning (outside Dgraph's model)
4. RDFS rules fit naturally into query optimization

#### OWL Concepts for Phase 6+

**OWL Equivalence Classes**:

```owl
:Student owl:equivalentClass [
  owl:intersectionOf (:Person :hasParent [:Student])
]
```

Reasoning needed:

- Class equivalence → same query results
- Property equivalence → predicate aliasing
- Value equivalence → UID normalization

**OWL Restrictions**:

```owl
:Parent owl:Restriction [
  owl:onProperty :hasChild ;
  owl:minCardinality 1 ;
  owl:allValuesFrom :Person
]
```

Reasoning needed:

- Cardinality constraints → validation
- Value restrictions → type checking
- Existential quantification → completeness checking

**OWL Disjointness**:

```owl
:Student owl:disjointWith :Teacher
```

Reasoning needed:

- Contradiction detection
- Query optimization (mutual exclusion)
- Data quality checking

**Path to OWL Support**:

1. Phase 6: Store OWL constructs in ontology model
2. Phase 6: Implement restricted OWL reasoning (equivalence only)
3. Phase 7: Add OWL constraint checking
4. Phase 8: Full OWL 2 support (potentially with external reasoner)

### 7.3 How to Support Custom Reasoning Rules (SWRL, etc)

#### SWRL: Semantic Web Rule Language

**Example SWRL Rule**:

```swrl
:Person(?x) ∧ :hasBrother(?x, ?y) ∧ :hasSister(?z, ?y)
→ :Cousin(?x, ?z)
```

**Implementation Strategy**:

```go
type SWRLRule struct {
  Head      *RulePattern      // Consequent
  Body      []*RulePattern    // Antecedents
}

type RuleEngine struct {
  swrlRules []*SWRLRule

  // Apply rules during query compilation
  ApplyRules(pattern *GraphPattern) []*GraphPattern
}

// Example: hasBrother join with hasSister
// → Infer cousin relationship
```

**Challenges**:

- SWRL is Turing-complete → infinite computation possible
- Need rule safety analysis
- Need performance budgets
- Need query-time or materialization strategy

**Proposal for Phase 7+**:

```go
type RuleApplicationStrategy int
const (
  StrategyQueryTime RuleApplicationStrategy = iota  // Apply during compilation
  StrategyMaterialized                               // Pre-compute and store
  StrategyHybrid                                     // Mixture
)
```

### 7.4 How to Federate with Remote Ontologies

#### Problem

Local ontologies often reference concepts from remote ontologies:

```turtle
@prefix foaf: <http://xmlns.com/foaf/0.1/> .

ex:extends foaf:Person ;
  ...
```

#### Solution Strategy: Ontology Caching

**Phase 1: Simple Reference**

```go
// Store reference
type OntologyReference struct {
  URI          string
  Loaded       bool
  Cache        *OntologyModel
  LastFetched  time.Time
}

// Manual loading
ontology := loader.LoadRemote("http://xmlns.com/foaf/0.1/foaf.rdf")
cache.RegisterRemote(ontology)
```

**Phase 2: Automatic Fetching**

```go
// Auto-fetch referenced ontologies
func (l *Loader) LoadWithDependencies(uri string) (*OntologyModel, error) {
  root := l.load(uri)

  for _, ref := range root.References {
    dep := l.LoadRemote(ref)
    root.Merge(dep)
  }

  return root, nil
}
```

**Phase 3: Federated Queries**

```sparql
SELECT ?person WHERE {
  ?person rdf:type foaf:Person .  # Resolved from local cache of remote ontology
  ?person ex:worksFor ?company .   # Local ontology
}
```

#### Caching Strategy for Remote Ontologies

```go
type RemoteOntologyCache struct {
  cache          map[string]*OntologyModel
  ttl            map[string]time.Time
  fetchTimeout   time.Duration
  maxCacheSize   int64

  // Invalidation
  watchedURIs    map[string]chan *Update
}
```

#### Trust & Security Considerations

- Ontologies loaded from remote sources
- Need signature verification
- Need size/complexity limits
- Need allowed source list (whitelist)

### 7.5 How to Version Ontologies

#### Versioning Strategies

**Strategy 1: Semantic Versioning**

```
FOAF 0.0.1 - Initial version
FOAF 0.1.0 - Add properties, backward compatible
FOAF 1.0.0 - Breaking change (major version)
```

**Strategy 2: Timestamp-Based**

```
FOAF 2024-01-15 - Snapshot from date
FOAF 2024-02-01 - Later snapshot
```

**Strategy 3: Content-Based**

```
FOAF hash:abc123 - Identified by content hash
FOAF hash:def456 - Different content
```

#### Implementation in Dgraph

```go
type OntologyVersion struct {
  URI              string    // ontology URI
  Version          string    // semantic version
  Timestamp        time.Time // when loaded
  Hash             string    // content hash
  Supersedes       string    // previous version
  Deprecated       bool

  classes          map[string]*Class
  properties       map[string]*Property
}

// Store in Dgraph
type VersionedOntology {
  uri: string
  version: string
  timestamp: datetime
  hash: string
  classes: [Class]
  properties: [Property]
}
```

#### Migration Between Versions

```go
type OntologyMigration struct {
  From          string
  To            string
  Mapping       map[string]string  // Old URI → New URI
  Renames       map[string]string  // Old name → New name
  Deletions     []string           // Removed concepts
  Additions     []string           // New concepts
}

// Apply migration to existing data
func (m *OntologyMigration) Apply(dgraph Client) error {
  // Rewrite URIs in data
  // Update class memberships
  // Verify constraints
}
```

#### Phase 6+ Design Points

1. **Backward Compatibility**: Can queries written for v0.1 work with v0.2?
2. **Data Upgrade**: How to upgrade existing data to new ontology version?
3. **Multi-Version Support**: Can a single Dgraph instance run queries against multiple ontology
   versions?
4. **Version Negotiation**: How to specify which version a query uses?

### 7.6 Summary of Open Questions

| Question                | Phase | Priority | Status       |
| ----------------------- | ----- | -------- | ------------ |
| Ontology updates        | 5     | HIGH     | Sketched     |
| Reasoning scope         | 5     | HIGH     | Sketched     |
| Multi-tenancy           | 5     | HIGH     | Sketched     |
| Ontology versioning     | 6     | MEDIUM   | Not designed |
| Custom reasoners (SWRL) | 7     | MEDIUM   | Sketched     |
| Remote ontologies       | 6     | MEDIUM   | Sketched     |
| OWL beyond RDFS         | 6     | LOW      | Sketched     |

---

## Appendix: Example Ontology Files

### A1: FOAF Ontology (Turtle Format)

```turtle
@prefix rdf: <http://www.w3.org/1999/02/22-rdf-syntax-ns#> .
@prefix rdfs: <http://www.w3.org/2000/01/rdf-schema#> .
@prefix foaf: <http://xmlns.com/foaf/0.1/> .

foaf:Person
  a rdfs:Class ;
  rdfs:label "Person"@en ;
  rdfs:comment "A person." ;
  rdfs:isDefinedBy foaf: ;
.

foaf:name
  a rdf:Property ;
  rdfs:label "Name"@en ;
  rdfs:comment "A name for some thing." ;
  rdfs:domain foaf:Agent ;
  rdfs:range rdfs:Literal ;
  rdfs:isDefinedBy foaf: ;
.

foaf:knows
  a rdf:Property ;
  rdfs:label "Knows"@en ;
  rdfs:comment "A person known by this person (indicating some level of reciprocated interaction)." ;
  rdfs:domain foaf:Person ;
  rdfs:range foaf:Person ;
  rdfs:isDefinedBy foaf: ;
.
```

### A2: University Ontology (OWL)

```owl
<?xml version="1.0"?>
<rdf:RDF
  xmlns:rdf="http://www.w3.org/1999/02/22-rdf-syntax-ns#"
  xmlns:rdfs="http://www.w3.org/2000/01/rdf-schema#"
  xmlns:owl="http://www.w3.org/2002/07/owl#"
  xmlns:ex="http://example.edu/ontology#">

  <!-- Classes -->
  <owl:Class rdf:about="http://example.edu/ontology#Person">
    <rdfs:label>Person</rdfs:label>
  </owl:Class>

  <owl:Class rdf:about="http://example.edu/ontology#Student">
    <rdfs:subClassOf rdf:resource="http://example.edu/ontology#Person"/>
    <rdfs:label>Student</rdfs:label>
  </owl:Class>

  <owl:Class rdf:about="http://example.edu/ontology#Faculty">
    <rdfs:subClassOf rdf:resource="http://example.edu/ontology#Person"/>
    <rdfs:label>Faculty</rdfs:label>
  </owl:Class>

  <!-- Properties -->
  <owl:ObjectProperty rdf:about="http://example.edu/ontology#enrolledIn">
    <rdfs:domain rdf:resource="http://example.edu/ontology#Student"/>
    <rdfs:range rdf:resource="http://example.edu/ontology#Course"/>
    <rdfs:label>Enrolled In</rdfs:label>
  </owl:ObjectProperty>

  <owl:ObjectProperty rdf:about="http://example.edu/ontology#teaches">
    <rdfs:domain rdf:resource="http://example.edu/ontology#Faculty"/>
    <rdfs:range rdf:resource="http://example.edu/ontology#Course"/>
    <rdfs:label>Teaches</rdfs:label>
  </owl:ObjectProperty>
</rdf:RDF>
```

---

## Document Metadata

**Version**: 1.0  
**Created**: January 2026  
**Last Updated**: January 2026  
**Status**: Design Proposal  
**Target Audience**: Dgraph Architecture Team, SPARQL Developers  
**Review Checklist**:

- [ ] Architecture team review
- [ ] SPARQL implementation lead review
- [ ] Storage layer team review
- [ ] Query execution team review
- [ ] Approved for Phase 5 implementation
