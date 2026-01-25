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
		// translateSelect is implemented in translator_impl.go (first-pass / fallback)
		return translateSelect(ctx, sq, opts)
	case "ASK":
		return translateAsk(ctx, sq, opts)
	default:
		return nil, nil, fmt.Errorf("unsupported SPARQL query type %q", sq.Type())
	}
}
