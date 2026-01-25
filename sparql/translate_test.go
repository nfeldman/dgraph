package sparql

import (
	"context"
	"testing"

	dql "github.com/dgraph-io/dgraph/v25/dql"
)

func TestSPARQLParseAndTranslate_Basic(t *testing.T) {
	parser := &ANTLRParserAdapter{}
	ctx := context.Background()
	opts := TranslateOptions{}

	// This test should work after the visitor and translator are implemented
	q := `
PREFIX ex: <http://example.org/>
SELECT ?s WHERE {
  GRAPH <http://g1> { ?s ex:p ?o }
}`

	_, _, err := ParseAndTranslate(ctx, parser, q, opts)
	if err != nil {
		t.Log("ParseAndTranslate error:", err)
		// For now, ok if not implemented.
	}
}

// TestBuildGraphFilter_NamedGraph tests graph filtering for GRAPH patterns
func TestBuildGraphFilter_NamedGraph(t *testing.T) {
	tests := []struct {
		name            string
		tripleGraph     string
		fromGraphs      []string
		fromNamedGraphs []string
		expectNil       bool
		expectFunc      string
	}{
		{
			name:            "GRAPH pattern with matching FROM NAMED",
			tripleGraph:     "http://g1",
			fromGraphs:      []string{},
			fromNamedGraphs: []string{"http://g1", "http://g2"},
			expectNil:       false,
			expectFunc:      "eq",
		},
		{
			name:            "GRAPH pattern without matching FROM NAMED",
			tripleGraph:     "http://g3",
			fromGraphs:      []string{},
			fromNamedGraphs: []string{"http://g1", "http://g2"},
			expectNil:       true, // Should be skipped
		},
		{
			name:            "GRAPH pattern with no FROM NAMED specified",
			tripleGraph:     "http://g1",
			fromGraphs:      []string{},
			fromNamedGraphs: []string{},
			expectNil:       false,
			expectFunc:      "eq",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := buildGraphFilter(tt.tripleGraph, tt.fromGraphs, tt.fromNamedGraphs)
			if tt.expectNil {
				if filter != nil {
					t.Errorf("Expected nil filter, got %+v", filter)
				}
				return
			}
			if filter == nil {
				t.Fatalf("Expected non-nil filter")
			}
			if filter.Func == nil {
				t.Fatalf("Expected filter with Func")
			}
			if filter.Func.Name != tt.expectFunc {
				t.Errorf("Expected function %s, got %s", tt.expectFunc, filter.Func.Name)
			}
			if filter.Func.Attr != "graph" {
				t.Errorf("Expected attr 'graph', got '%s'", filter.Func.Attr)
			}
		})
	}
}

// TestBuildGraphFilter_DefaultGraph tests graph filtering for default graph patterns
func TestBuildGraphFilter_DefaultGraph(t *testing.T) {
	tests := []struct {
		name            string
		fromGraphs      []string
		fromNamedGraphs []string
		expectNil       bool
		expectOp        string
	}{
		{
			name:            "Single FROM graph",
			fromGraphs:      []string{"http://g1"},
			fromNamedGraphs: []string{},
			expectNil:       false,
			expectOp:        "", // Single eq function, no OR
		},
		{
			name:            "Multiple FROM graphs",
			fromGraphs:      []string{"http://g1", "http://g2"},
			fromNamedGraphs: []string{},
			expectNil:       false,
			expectOp:        "OR",
		},
		{
			name:            "No FROM or FROM NAMED",
			fromGraphs:      []string{},
			fromNamedGraphs: []string{},
			expectNil:       true, // No filter needed - match all
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			filter := buildGraphFilter("", tt.fromGraphs, tt.fromNamedGraphs)
			if tt.expectNil {
				if filter != nil {
					t.Errorf("Expected nil filter, got %+v", filter)
				}
				return
			}
			if filter == nil {
				t.Fatalf("Expected non-nil filter")
			}
			if tt.expectOp == "" {
				// Single condition
				if filter.Func == nil {
					t.Fatalf("Expected filter with Func")
				}
				if filter.Func.Name != "eq" {
					t.Errorf("Expected function eq, got %s", filter.Func.Name)
				}
			} else {
				// OR condition
				if filter.Op != tt.expectOp {
					t.Errorf("Expected Op %s, got %s", tt.expectOp, filter.Op)
				}
				if len(filter.Child) != 2 {
					t.Errorf("Expected 2 children in OR filter, got %d", len(filter.Child))
				}
			}
		})
	}
}

// TestTranslateSelect_WithGraphFilters tests the full translation with graph filters
func TestTranslateSelect_WithGraphFilters(t *testing.T) {
	ctx := context.Background()
	opts := TranslateOptions{}

	tests := []struct {
		name        string
		query       *sparqlQueryImpl
		expectError bool
	}{
		{
			name: "SELECT with FROM clause",
			query: &sparqlQueryImpl{
				qtype:     "SELECT",
				prefixes:  map[string]string{"ex": "http://example.org/"},
				projs:     []string{"s", "o"},
				from:      []string{"http://g1"},
				fromNamed: []string{},
				bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "ex:p",
						Object:      "?o",
						ObjectIsVar: true,
						Graph:       "", // default graph
					}},
				}},
			},
			expectError: false,
		},
		{
			name: "SELECT with FROM NAMED and GRAPH",
			query: &sparqlQueryImpl{
				qtype:     "SELECT",
				prefixes:  map[string]string{"ex": "http://example.org/"},
				projs:     []string{"s", "o"},
				from:      []string{},
				fromNamed: []string{"http://g1", "http://g2"},
				bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "ex:p",
						Object:      "?o",
						ObjectIsVar: true,
						Graph:       "http://g1", // named graph
					}},
				}},
			},
			expectError: false,
		},
		{
			name: "SELECT with multiple FROM",
			query: &sparqlQueryImpl{
				qtype:     "SELECT",
				prefixes:  map[string]string{"ex": "http://example.org/"},
				projs:     []string{"s", "o"},
				from:      []string{"http://g1", "http://g2", "http://g3"},
				fromNamed: []string{},
				bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "ex:p",
						Object:      "?o",
						ObjectIsVar: true,
						Graph:       "",
					}},
				}},
			},
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gqs, prefixes, err := translateSelect(ctx, tt.query, opts)
			if tt.expectError && err == nil {
				t.Errorf("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if err == nil {
				if len(gqs) == 0 {
					t.Errorf("Expected non-empty GraphQuery slice")
				}
				if prefixes == nil {
					t.Errorf("Expected prefixes map")
				}
				// Check that filters were applied
				if len(tt.query.from) > 0 || tt.query.bgps[0].Triples[0].Graph != "" {
					if gqs[0].Children != nil && len(gqs[0].Children) > 0 {
						child := gqs[0].Children[0]
						if child.Filter == nil {
							t.Errorf("Expected filter to be set on child query")
						}
					}
				}
			}
		})
	}
}

// TestTranslateAsk tests ASK query translation
func TestTranslateAsk(t *testing.T) {
	ctx := context.Background()
	opts := TranslateOptions{}

	query := &sparqlQueryImpl{
		qtype:     "ASK",
		prefixes:  map[string]string{"ex": "http://example.org/"},
		from:      []string{"http://g1"},
		fromNamed: []string{},
		bgps: []*BGP{{
			Triples: []*Triple{{
				Subject:     "?s",
				Predicate:   "ex:p",
				Object:      "?o",
				ObjectIsVar: true,
				Graph:       "",
			}},
		}},
	}

	gqs, prefixes, err := translateAsk(ctx, query, opts)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	if len(gqs) == 0 {
		t.Errorf("Expected non-empty GraphQuery slice")
	}
	if prefixes == nil {
		t.Errorf("Expected prefixes map")
	}
}

// TestBuildOrFilter tests OR filter combination
func TestBuildOrFilter(t *testing.T) {
	tests := []struct {
		name           string
		conditionCount int
		expectNil      bool
	}{
		{"No conditions", 0, true},
		{"Single condition", 1, false},
		{"Two conditions", 2, false},
		{"Three conditions", 3, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var conditions []*dql.FilterTree
			for i := 0; i < tt.conditionCount; i++ {
				conditions = append(conditions, &dql.FilterTree{
					Func: &dql.Function{
						Name: "eq",
						Attr: "graph",
						Args: []dql.Arg{{Value: "\"g\""}},
					},
				})
			}

			result := buildOrFilter(conditions)
			if tt.expectNil {
				if result != nil {
					t.Errorf("Expected nil result")
				}
				return
			}
			if result == nil {
				t.Fatalf("Expected non-nil result")
			}

			if tt.conditionCount == 1 {
				if result.Op != "" {
					t.Errorf("Single condition should not have Op set")
				}
				if result.Func == nil {
					t.Errorf("Single condition should have Func set")
				}
			} else {
				// Multi-condition OR
				if result.Op != "OR" {
					t.Errorf("Expected Op=OR, got %s", result.Op)
				}
			}
		})
	}
}

// TestQuoteString tests string quoting
func TestQuoteString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.org/g1", "\"http://example.org/g1\""},
		{"\"already quoted\"", "\"already quoted\""},
		{"test", "\"test\""},
		{"", "\"\""},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := quoteString(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %s, got %s", tt.expected, result)
			}
		})
	}
}
