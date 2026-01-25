package sparql

import (
	"context"
)

// SPARQLParser is an adapter interface for pluggable SPARQL parsers.
// Implementations produce a SPARQLQuery wrapper that the translator consumes.
type SPARQLParser interface {
	// Parse the raw SPARQL text and return a SPARQLQuery wrapper (intermediate AST).
	Parse(ctx context.Context, query string) (SPARQLQuery, error)
}

// SPARQLQuery is an abstract wrapper around the parser-specific AST.
// Implement methods needed by TranslateToGraphQueries (the implementation here
// can be expanded as we implement visitor/walker).
type SPARQLQuery interface {
	// Type returns top-level query type: "SELECT", "ASK", "CONSTRUCT", etc.
	Type() string
	// GetPrefixes returns the map of declared prefixes in the query.
	GetPrefixes() map[string]string
	// RootGraphPatterns returns a list of top-level graph pattern nodes.
	// Concrete types for GraphPattern are defined in ast.go.
	RootGraphPatterns() []interface{}
	// ProjectionVars returns the projection variable names for SELECT queries.
	ProjectionVars() []string
	// FROM returns the list of graph IRIs that form the default graph.
	FROM() []string
	// FROMNamed returns the list of named graph IRIs available for GRAPH patterns.
	FROMNamed() []string
}
