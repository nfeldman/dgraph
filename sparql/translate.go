package sparql

import (
	"context"
	"fmt"

	dql "github.com/dgraph-io/dgraph/v25/dql"
)

// TranslateOptions holds configuration for translation
type TranslateOptions struct {
	// AllowFilterFallback permits falling back to raw FILTER attachment when parsing fails.
	// Defaults to false (strict): parsing errors will surface.
	AllowFilterFallback bool
}

// TranslateToGraphQueries is the main entry point for translating SPARQL queries to DQL
func TranslateToGraphQueries(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq == nil {
		return nil, nil, fmt.Errorf("query is nil")
	}

	queryType := sq.Type()
	switch queryType {
	case "SELECT":
		return TranslateSelectExtended(ctx, sq, opts)
	case "ASK":
		return translateAsk(ctx, sq, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported query type: %s", queryType)
	}
}

// ParseAndTranslate parses a SPARQL query string and translates it to DQL
func ParseAndTranslate(ctx context.Context, parser SPARQLParser, query string, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	// Parse the SPARQL query
	sq, err := parser.Parse(ctx, query)
	if err != nil {
		return nil, nil, fmt.Errorf("parsing SPARQL: %w", err)
	}

	// Translate to DQL
	return TranslateToGraphQueries(ctx, sq, opts)
}

// hasExtendedFeatures checks if the query uses extended SPARQL features
func hasExtendedFeatures(sq SPARQLQuery) bool {
	// Check if this is an extended query with new features
	if extQuery, ok := sq.(*SPARQLQueryImpl); ok {
		return len(extQuery.Patterns) > 0 ||
			len(extQuery.Aggregates) > 0 ||
			len(extQuery.Binds) > 0 ||
			extQuery.Having != nil ||
			extQuery.Distinct ||
			len(extQuery.OrderBy) > 0
	}
	return false
}
