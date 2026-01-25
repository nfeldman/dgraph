package sparql

import (
	"context"
	"fmt"

	"github.com/antlr/antlr4/runtime/Go/antlr"
	gen "github.com/nfeldman/dgraph/sparql/gen"
)

// ANTLRParserAdapter implements SPARQLParser using the ANTLR-generated parser.
type ANTLRParserAdapter struct{}

func (a *ANTLRParserAdapter) Parse(ctx context.Context, query string) (SPARQLQuery, error) {
	if query == "" {
		return nil, fmt.Errorf("empty query")
	}

	// Create ANTLR input/lexer/parser
	input := antlr.NewInputStream(query)
	lexer := gen.NewSPARQLLexer(input)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)
	parser := gen.NewSPARQLParser(stream)
	parser.BuildParseTrees = true

	// Entry rule: Query (adjust this if generated entry differs)
	tree := parser.Query()

	// Build the visitor and walk.
	vis := NewParseVisitor()
	tree.Accept(vis)

	// Return the result. The visitor builds a sparqlQueryImpl.
	result := vis.Result()
	if result == nil {
		return nil, fmt.Errorf("failed to build SPARQL AST from parse tree")
	}
	return result, nil
}
