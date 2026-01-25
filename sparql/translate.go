package sparql

import (
	"context"
	"errors"
	"fmt"

	dql "github.com/nfeldman/dgraph/dql"
)

// TranslateOptions may include dataset and translation settings.
type TranslateOptions struct {
	// Any additional options go here, e.g. rewriting rules, query mode, etc.
}

// ParseAndTranslate: high-level entry, handles parsing and DQL translation.
func ParseAndTranslate(ctx context.Context, parser SPARQLParser, q string, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	sq, err := parser.Parse(ctx, q)
	if err != nil {
		return nil, nil, err
	}
	return TranslateToGraphQueries(ctx, sq, opts)
}

func TranslateToGraphQueries(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq == nil {
		return nil, nil, errors.New("nil SPARQLQuery")
	}

	switch sq.Type() {
	case "SELECT":
		// Try extended translator first (with OPTIONAL, UNION, aggregates, BIND, HAVING)
		// If the query has extended features, it will be handled here
		if hasExtendedFeatures(sq) {
			return TranslateSelectExtended(ctx, sq, opts)
		}
		// Fall back to basic translator for simple queries
		return translateSelect(ctx, sq, opts)
	case "ASK":
		return translateAsk(ctx, sq, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported SPARQL query type %q", sq.Type())
	}
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
