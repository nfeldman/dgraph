package sparql

// Lightweight intermediate AST for SPARQL used by the translator and the
// fallback plain parser. This is intentionally minimal and designed to be
// easy to populate from a simple parser. The ANTLR visitor (later) will
// produce the same SPARQLQuery implementation.

// Triple represents a single RDF triple pattern.
type Triple struct {
	Subject     string // either a variable like "?s" or an IRI/literal
	Predicate   string // prefixed name or IRI (raw as in the query)
	Object      string // variable or literal/IRI
	ObjectIsVar bool
	Graph       string // named graph IRI or "" for default graph
}

// BGP (basic graph pattern) is a sequence of triples that are joined.
type BGP struct {
	Triples []*Triple
}

// GraphPattern is a union type for different SPARQL pattern types.
type GraphPattern interface {
	isGraphPattern()
}

func (b *BGP) isGraphPattern()             {}
func (o *OptionalPattern) isGraphPattern() {}
func (u *UnionPattern) isGraphPattern()    {}
func (f *FilterPattern) isGraphPattern()   {}

// OptionalPattern represents OPTIONAL { ... } pattern blocks
type OptionalPattern struct {
	Patterns []GraphPattern // Patterns to match optionally
	Required bool           // False for OPTIONAL, true for non-optional
}

// UnionPattern represents UNION { ... } { ... } alternatives
type UnionPattern struct {
	Alternatives [][]GraphPattern // Each alternative is a list of patterns
}

// FilterPattern represents a FILTER expression applied to the surrounding group pattern.
type FilterPattern struct {
	Expression string
}

// Aggregate represents an aggregate function like COUNT, SUM, etc.
type Aggregate struct {
	Function string // "count", "sum", "min", "max", "avg"
	Variable string // Variable to aggregate
	Alias    string // Output alias
	Distinct bool   // DISTINCT modifier
}

// BindExpression represents BIND (?expr AS ?var) binding
type BindExpression struct {
	Expression string // Math or string expression
	Variable   string // Output variable
}

// HavingClause represents HAVING filter on aggregates
type HavingClause struct {
	Expression string // Boolean expression for filtering
}

// sparqlQueryImpl is a minimal implementation of SPARQLQuery used by the
// plain parser and translator.
type sparqlQueryImpl struct {
	qtype      string
	prefixes   map[string]string
	projs      []string
	patterns   []GraphPattern    // New: supports OPTIONAL, UNION, and other pattern types
	bgps       []*BGP            // Deprecated: only for backwards compatibility with old simple parser output
	aggregates []*Aggregate      // Aggregate functions (COUNT, SUM, AVG, MIN, MAX)
	binds      []*BindExpression // BIND expressions
	having     *HavingClause     // HAVING clause
	limit      int
	offset     int
	orderBy    []string // ORDER BY variables
	distinct   bool     // DISTINCT modifier
	from       []string // active dataset IRIs (FROM)
	fromNamed  []string // available named graphs (FROM NAMED)
}

// SPARQLQueryImpl is an exported version for testing and examples.
// Prefer using patterns field for new code; bgps is maintained for backwards compatibility.
type SPARQLQueryImpl struct {
	Qtype      string
	Prefixes   map[string]string
	Projs      []string
	Patterns   []GraphPattern    // New: supports OPTIONAL, UNION, and other pattern types
	Bgps       []*BGP            // Deprecated: only for backwards compatibility with old simple parser output
	Aggregates []*Aggregate      // Aggregate functions (COUNT, SUM, AVG, MIN, MAX)
	Binds      []*BindExpression // BIND expressions
	Having     *HavingClause     // HAVING clause
	Limit      int
	Offset     int
	OrderBy    []string // ORDER BY variables
	Distinct   bool     // DISTINCT modifier
	From       []string // active dataset IRIs (FROM)
	FromNamed  []string // available named graphs (FROM NAMED)
}

func (s *sparqlQueryImpl) Type() string                   { return s.qtype }
func (s *sparqlQueryImpl) GetPrefixes() map[string]string { return s.prefixes }
func (s *sparqlQueryImpl) RootGraphPatterns() []interface{} {
	// Return patterns if available, otherwise fall back to bgps for backwards compatibility
	if len(s.patterns) > 0 {
		out := make([]interface{}, 0, len(s.patterns))
		for _, p := range s.patterns {
			out = append(out, p)
		}
		return out
	}
	out := make([]interface{}, 0, len(s.bgps))
	for _, b := range s.bgps {
		out = append(out, b)
	}
	return out
}
func (s *sparqlQueryImpl) ProjectionVars() []string { return s.projs }
func (s *sparqlQueryImpl) FROM() []string           { return s.from }
func (s *sparqlQueryImpl) FROMNamed() []string      { return s.fromNamed }

// Exported version implementations
func (s *SPARQLQueryImpl) Type() string                   { return s.Qtype }
func (s *SPARQLQueryImpl) GetPrefixes() map[string]string { return s.Prefixes }
func (s *SPARQLQueryImpl) RootGraphPatterns() []interface{} {
	// Return patterns if available, otherwise fall back to bgps for backwards compatibility
	if len(s.Patterns) > 0 {
		out := make([]interface{}, 0, len(s.Patterns))
		for _, p := range s.Patterns {
			out = append(out, p)
		}
		return out
	}
	out := make([]interface{}, 0, len(s.Bgps))
	for _, b := range s.Bgps {
		out = append(out, b)
	}
	return out
}
func (s *SPARQLQueryImpl) ProjectionVars() []string { return s.Projs }
func (s *SPARQLQueryImpl) FROM() []string           { return s.From }
func (s *SPARQLQueryImpl) FROMNamed() []string      { return s.FromNamed }
