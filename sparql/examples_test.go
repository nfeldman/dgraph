package sparql_test

import (
	"context"
	"fmt"
	"strings"

	"github.com/dgraph-io/dgraph/v25/sparql"
)

// Example_quadLoading demonstrates loading RDF quads with named graph support
func Example_quadLoading() {
	// Example N-Quads data
	nquads := `
<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/graph1> .
<http://example.org/alice> <http://example.org/age> "30" <http://example.org/graph1> .
<http://example.org/bob> <http://example.org/name> "Bob" <http://example.org/graph2> .
<http://example.org/bob> <http://example.org/friend> <http://example.org/alice> <http://example.org/graph2> .
`

	loader := sparql.NewQuadLoader()
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		fmt.Printf("Error loading quads: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d triples (includes graph membership triples)\n", len(triples))

	// Show first few triples
	for i, t := range triples[:4] {
		fmt.Printf("Triple %d: <%s> <%s> %s\n", i+1, t.Subject, t.Predicate, t.Object)
	}

	// Output:
	// Loaded 8 triples (includes graph membership triples)
	// Triple 1: <http://example.org/alice> <http://example.org/name> "Alice"
	// Triple 2: <http://example.org/alice> <graph> "http://example.org/graph1"
	// Triple 3: <http://example.org/alice> <http://example.org/age> "30"
	// Triple 4: <http://example.org/alice> <graph> "http://example.org/graph1"
}

// Example_sparqlTranslation demonstrates SPARQL to DQL translation with graph filtering
func Example_sparqlTranslation() {
	ctx := context.Background()

	// Create a mock SPARQL query with FROM and FROM NAMED
	query := &sparql.SPARQLQueryImpl{
		Qtype:     "SELECT",
		Prefixes:  map[string]string{"ex": "http://example.org/"},
		Projs:     []string{"s", "name"},
		From:      []string{"http://example.org/graph1"},
		FromNamed: []string{"http://example.org/graph2"},
		Bgps: []*sparql.BGP{
			{
				Triples: []*sparql.Triple{
					{
						Subject:     "?s",
						Predicate:   "ex:name",
						Object:      "?name",
						ObjectIsVar: true,
						Graph:       "", // Default graph pattern
					},
				},
			},
		},
	}

	gqs, prefixes, err := sparql.TranslateToGraphQueries(ctx, query, sparql.TranslateOptions{})
	if err != nil {
		fmt.Printf("Translation error: %v\n", err)
		return
	}

	fmt.Printf("Translated to %d DQL queries\n", len(gqs))
	fmt.Printf("Prefixes: %v\n", prefixes)
	fmt.Printf("Root query attr: %s\n", gqs[0].Attr)

	if len(gqs[0].Children) > 0 && gqs[0].Children[0].Filter != nil {
		filter := gqs[0].Children[0].Filter
		if filter.Func != nil {
			fmt.Printf("Filter function: %s on attr: %s\n", filter.Func.Name, filter.Func.Attr)
		}
	}

	// Output:
	// Translated to 1 DQL queries
	// Prefixes: map[ex:http://example.org/]
	// Root query attr: query
	// Filter function: eq on attr: graph
}

// Example_namedGraphQuery demonstrates a full workflow
func Example_namedGraphQuery() {
	fmt.Println("SPARQL Named Graph Support")
	fmt.Println("==========================")
	fmt.Println()
	fmt.Println("Data Model:")
	fmt.Println("For each RDF quad: <s> <p> <o> <g>")
	fmt.Println("We store:")
	fmt.Println("  1. <s> <p> <o>")
	fmt.Println("  2. <s> <graph> \"g\"")
	fmt.Println()
	fmt.Println("Query Translation:")
	fmt.Println("- FROM <g1> <g2>: Filters default patterns to graphs g1, g2")
	fmt.Println("- FROM NAMED <g3>: Makes g3 available for GRAPH patterns")
	fmt.Println("- GRAPH <g3> { ... }: Adds filter eq(graph, \"g3\")")
	fmt.Println()
	fmt.Println("Example SPARQL:")
	fmt.Println("  SELECT ?s ?name")
	fmt.Println("  FROM <http://example.org/graph1>")
	fmt.Println("  FROM NAMED <http://example.org/graph2>")
	fmt.Println("  WHERE {")
	fmt.Println("    ?s <http://example.org/name> ?name .")
	fmt.Println("    GRAPH <http://example.org/graph2> {")
	fmt.Println("      ?s <http://example.org/verified> true .")
	fmt.Println("    }")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("DQL Translation:")
	fmt.Println("  - First pattern gets: @filter(eq(graph, \"http://example.org/graph1\"))")
	fmt.Println("  - GRAPH pattern gets: @filter(eq(graph, \"http://example.org/graph2\"))")

	// Output:
	// SPARQL Named Graph Support
	// ==========================
	//
	// Data Model:
	// For each RDF quad: <s> <p> <o> <g>
	// We store:
	//   1. <s> <p> <o>
	//   2. <s> <graph> "g"
	//
	// Query Translation:
	// - FROM <g1> <g2>: Filters default patterns to graphs g1, g2
	// - FROM NAMED <g3>: Makes g3 available for GRAPH patterns
	// - GRAPH <g3> { ... }: Adds filter eq(graph, "g3")
	//
	// Example SPARQL:
	//   SELECT ?s ?name
	//   FROM <http://example.org/graph1>
	//   FROM NAMED <http://example.org/graph2>
	//   WHERE {
	//     ?s <http://example.org/name> ?name .
	//     GRAPH <http://example.org/graph2> {
	//       ?s <http://example.org/verified> true .
	//     }
	//   }
	//
	// DQL Translation:
	//   - First pattern gets: @filter(eq(graph, "http://example.org/graph1"))
	//   - GRAPH pattern gets: @filter(eq(graph, "http://example.org/graph2"))
}
