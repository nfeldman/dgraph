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
}

// BGP (basic graph pattern) is a sequence of triples that are joined.
type BGP struct {
	Triples []*Triple
}

// sparqlQueryImpl is a minimal implementation of SPARQLQuery used by the
// plain parser and translator.
type sparqlQueryImpl struct {
	qtype    string
	prefixes map[string]string
	projs    []string
	bgps     []*BGP
	limit    int
	offset   int
}

func (s *sparqlQueryImpl) Type() string { return s.qtype }
func (s *sparqlQueryImpl) GetPrefixes() map[string]string { return s.prefixes }
func (s *sparqlQueryImpl) RootGraphPatterns() []interface{} {
	out := make([]interface{}, 0, len(s.bgps))
	for _, b := range s.bgps {
		out = append(out, b)
	}
	return out
}
func (s *sparqlQueryImpl) ProjectionVars() []string { return s.projs }