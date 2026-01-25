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

// sparqlQueryImpl is a minimal implementation of SPARQLQuery used by the
// plain parser and translator.
type sparqlQueryImpl struct {
	qtype     string
	prefixes  map[string]string
	projs     []string
	bgps      []*BGP
	limit     int
	offset    int
	from      []string // active dataset IRIs (FROM)
	fromNamed []string // available named graphs (FROM NAMED)
}

// SPARQLQueryImpl is an exported version for testing and examples
type SPARQLQueryImpl struct {
	Qtype     string
	Prefixes  map[string]string
	Projs     []string
	Bgps      []*BGP
	Limit     int
	Offset    int
	From      []string // active dataset IRIs (FROM)
	FromNamed []string // available named graphs (FROM NAMED)
}

func (s *sparqlQueryImpl) Type() string                   { return s.qtype }
func (s *sparqlQueryImpl) GetPrefixes() map[string]string { return s.prefixes }
func (s *sparqlQueryImpl) RootGraphPatterns() []interface{} {
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
	out := make([]interface{}, 0, len(s.Bgps))
	for _, b := range s.Bgps {
		out = append(out, b)
	}
	return out
}
func (s *SPARQLQueryImpl) ProjectionVars() []string { return s.Projs }
func (s *SPARQLQueryImpl) FROM() []string           { return s.From }
func (s *SPARQLQueryImpl) FROMNamed() []string      { return s.FromNamed }
