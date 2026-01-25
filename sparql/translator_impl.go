package sparql

import (
	"context"
	"fmt"
	"strings"

	dql "github.com/nfeldman/dgraph/dql"
)

// translateSelect converts a SPARQL SELECT query to DQL GraphQuery.
func translateSelect(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq.Type() != "SELECT" {
		return nil, nil, fmt.Errorf("expected SELECT query, got %s", sq.Type())
	}

	// Get FROM and FROM NAMED clauses
	fromGraphs := sq.FROM()
	fromNamedGraphs := sq.FROMNamed()

	// Build the DQL query
	rootQuery := &dql.GraphQuery{
		Attr: "query",
	}

	// Get the graph patterns
	patterns := sq.RootGraphPatterns()

	// Process each basic graph pattern (BGP)
	for _, patternIntf := range patterns {
		bgp, ok := patternIntf.(*BGP)
		if !ok {
			continue
		}

		// For each triple in the BGP, create a DQL query block
		for _, triple := range bgp.Triples {
			childQuery, err := translateTriple(triple, fromGraphs, fromNamedGraphs)
			if err != nil {
				return nil, nil, fmt.Errorf("translating triple: %w", err)
			}
			if childQuery != nil {
				rootQuery.Children = append(rootQuery.Children, childQuery)
			}
		}
	}

	// Return the query and prefixes
	return []*dql.GraphQuery{rootQuery}, sq.GetPrefixes(), nil
}

// translateAsk converts a SPARQL ASK query to DQL GraphQuery.
func translateAsk(ctx context.Context, sq SPARQLQuery, opts TranslateOptions) ([]*dql.GraphQuery, map[string]string, error) {
	if sq.Type() != "ASK" {
		return nil, nil, fmt.Errorf("expected ASK query, got %s", sq.Type())
	}

	// ASK queries are similar to SELECT but just check existence
	// We can reuse the SELECT logic and wrap it appropriately
	fromGraphs := sq.FROM()
	fromNamedGraphs := sq.FROMNamed()

	rootQuery := &dql.GraphQuery{
		Attr: "query",
	}

	patterns := sq.RootGraphPatterns()
	for _, patternIntf := range patterns {
		bgp, ok := patternIntf.(*BGP)
		if !ok {
			continue
		}

		for _, triple := range bgp.Triples {
			childQuery, err := translateTriple(triple, fromGraphs, fromNamedGraphs)
			if err != nil {
				return nil, nil, fmt.Errorf("translating triple: %w", err)
			}
			if childQuery != nil {
				rootQuery.Children = append(rootQuery.Children, childQuery)
			}
		}
	}

	return []*dql.GraphQuery{rootQuery}, sq.GetPrefixes(), nil
}

// translateTriple converts a single SPARQL triple pattern to a DQL GraphQuery.
// It applies graph filtering based on FROM, FROM NAMED, and the triple's Graph field.
func translateTriple(triple *Triple, fromGraphs, fromNamedGraphs []string) (*dql.GraphQuery, error) {
	if triple == nil {
		return nil, nil
	}

	// Create base query for this triple pattern
	query := &dql.GraphQuery{
		Attr: triple.Predicate,
	}

	// Apply graph filtering according to SPARQL semantics
	filter := buildGraphFilter(triple.Graph, fromGraphs, fromNamedGraphs)
	if filter != nil {
		query.Filter = filter
	}

	// Handle the subject
	if strings.HasPrefix(triple.Subject, "?") || strings.HasPrefix(triple.Subject, "$") {
		// Subject is a variable - use as Var
		query.Var = strings.TrimPrefix(strings.TrimPrefix(triple.Subject, "?"), "$")
	} else {
		// Subject is an IRI or literal - would need UID resolution
		// For now, we'll need to handle this in a more sophisticated way
		// This is a simplified placeholder
		query.Var = "_subject"
	}

	// Handle the object
	if triple.ObjectIsVar {
		// Object is a variable
		objVar := strings.TrimPrefix(strings.TrimPrefix(triple.Object, "?"), "$")
		// Create a child query for the object
		childQuery := &dql.GraphQuery{
			Attr: "val(" + objVar + ")",
		}
		query.Children = append(query.Children, childQuery)
	} else {
		// Object is a literal or IRI
		// Add as a filter or arg depending on the context
		if query.Args == nil {
			query.Args = make(map[string]string)
		}
		query.Args["value"] = triple.Object
	}

	return query, nil
}

// buildGraphFilter creates a FilterTree for graph membership based on SPARQL semantics.
//
// Rules:
// - If triple.Graph is set (inside GRAPH block):
//   - Filter on eq(graph, triple.Graph)
//   - Skip if triple.Graph not in fromNamedGraphs (when fromNamedGraphs is specified)
//
// - If triple.Graph is empty (default graph patterns):
//   - Filter on graph IN fromGraphs (when fromGraphs is specified)
//   - No filter if fromGraphs is empty (match all graphs)
//
// Returns nil if the pattern should be skipped entirely.
func buildGraphFilter(tripleGraph string, fromGraphs, fromNamedGraphs []string) *dql.FilterTree {
	if tripleGraph != "" {
		// Triple is inside a GRAPH block
		// Check if this graph is allowed by FROM NAMED
		if len(fromNamedGraphs) > 0 {
			found := false
			for _, g := range fromNamedGraphs {
				if g == tripleGraph {
					found = true
					break
				}
			}
			if !found {
				// This GRAPH is not in FROM NAMED - skip this pattern
				return nil
			}
		}

		// Create filter: eq(graph, "tripleGraph")
		return &dql.FilterTree{
			Func: &dql.Function{
				Name: "eq",
				Attr: "graph",
				Args: []dql.Arg{
					{Value: quoteString(tripleGraph)},
				},
			},
		}
	}

	// Triple is in default graph (not inside GRAPH block)
	if len(fromGraphs) > 0 {
		// Filter on graph IN fromGraphs
		// For multiple graphs: @filter(eq(graph, "g1") OR eq(graph, "g2") OR ...)
		if len(fromGraphs) == 1 {
			return &dql.FilterTree{
				Func: &dql.Function{
					Name: "eq",
					Attr: "graph",
					Args: []dql.Arg{
						{Value: quoteString(fromGraphs[0])},
					},
				},
			}
		}

		// Build OR filter for multiple graphs
		var conditions []*dql.FilterTree
		for _, g := range fromGraphs {
			conditions = append(conditions, &dql.FilterTree{
				Func: &dql.Function{
					Name: "eq",
					Attr: "graph",
					Args: []dql.Arg{
						{Value: quoteString(g)},
					},
				},
			})
		}

		// Combine with OR
		return buildOrFilter(conditions)
	}

	// No FROM clause - match all graphs (no filter needed)
	return nil
}

// buildOrFilter combines multiple FilterTree nodes with OR logic.
func buildOrFilter(conditions []*dql.FilterTree) *dql.FilterTree {
	if len(conditions) == 0 {
		return nil
	}
	if len(conditions) == 1 {
		return conditions[0]
	}

	// Build a tree of OR conditions
	// For simplicity, we'll create a right-associative tree
	result := conditions[len(conditions)-1]
	for i := len(conditions) - 2; i >= 0; i-- {
		result = &dql.FilterTree{
			Op:    "OR",
			Child: []*dql.FilterTree{conditions[i], result},
		}
	}
	return result
}

// quoteString adds quotes to a string value for use in DQL filters.
func quoteString(s string) string {
	// Simple quoting - in production, would need proper escaping
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return s
	}
	return "\"" + s + "\""
}
