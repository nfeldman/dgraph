package sparql

import (
	"fmt"
	"strings"
	"testing"
)

func TestNewQuadLoader(t *testing.T) {
	loader := NewQuadLoader()
	if loader == nil {
		t.Fatal("Expected non-nil loader")
	}
	if loader.DefaultGraph != "urn:x-arq:DefaultGraph" {
		t.Errorf("Expected default graph 'urn:x-arq:DefaultGraph', got '%s'", loader.DefaultGraph)
	}
}

func TestLoadQuads_SingleQuad(t *testing.T) {
	nquads := `<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .`
	loader := NewQuadLoader()
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should emit 2 triples: the data triple and the graph membership triple
	if len(triples) != 2 {
		t.Errorf("Expected 2 triples, got %d", len(triples))
		return
	}

	// Check data triple
	if triples[0].Subject != "http://example.org/alice" {
		t.Errorf("Expected subject 'http://example.org/alice', got '%s'", triples[0].Subject)
	}
	if triples[0].Predicate != "http://example.org/name" {
		t.Errorf("Expected predicate 'http://example.org/name', got '%s'", triples[0].Predicate)
	}

	// Check graph membership triple
	if triples[1].Predicate != "graph" {
		t.Errorf("Expected predicate 'graph', got '%s'", triples[1].Predicate)
	}
	if !strings.Contains(triples[1].Object, "http://example.org/g1") {
		t.Errorf("Expected object containing 'http://example.org/g1', got '%s'", triples[1].Object)
	}
}

func TestLoadQuads_MultipleQuads(t *testing.T) {
	nquads := `<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .
<http://example.org/bob> <http://example.org/name> "Bob" <http://example.org/g2> .
<http://example.org/alice> <http://example.org/age> "30" <http://example.org/g1> .`

	loader := NewQuadLoader()
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// 3 quads * 2 triples per quad = 6 triples
	if len(triples) != 6 {
		t.Errorf("Expected 6 triples, got %d", len(triples))
	}
}

func TestLoadQuads_EmptyAndComments(t *testing.T) {
	nquads := `# This is a comment
<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .

# Another comment
<http://example.org/bob> <http://example.org/age> "25" <http://example.org/g2> .`

	loader := NewQuadLoader()
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should skip comments and empty lines
	if len(triples) != 4 {
		t.Errorf("Expected 4 triples (2 quads), got %d", len(triples))
	}
}

func TestLoadQuads_InvalidFormat(t *testing.T) {
	nquads := `<http://example.org/alice> <http://example.org/name>`

	loader := NewQuadLoader()
	reader := strings.NewReader(nquads)

	_, err := loader.LoadQuads(reader)
	if err == nil {
		t.Error("Expected error for invalid format")
	}
}

func TestParseNQuad_Complete(t *testing.T) {
	line := `<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .`

	triple, err := parseNQuad(line)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if triple.Subject != "http://example.org/alice" {
		t.Errorf("Expected subject 'http://example.org/alice', got '%s'", triple.Subject)
	}
	if triple.Predicate != "http://example.org/name" {
		t.Errorf("Expected predicate 'http://example.org/name', got '%s'", triple.Predicate)
	}
	if triple.Graph != "http://example.org/g1" {
		t.Errorf("Expected graph 'http://example.org/g1', got '%s'", triple.Graph)
	}
}

func TestParseNQuad_TripleOnly(t *testing.T) {
	line := `<http://example.org/alice> <http://example.org/name> "Alice" .`

	triple, err := parseNQuad(line)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if triple.Subject != "http://example.org/alice" {
		t.Errorf("Expected subject 'http://example.org/alice', got '%s'", triple.Subject)
	}
	if triple.Graph != "" {
		t.Errorf("Expected empty graph for triple, got '%s'", triple.Graph)
	}
}

func TestParseNQuad_TrimsPeriod(t *testing.T) {
	line := `<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .`
	triple, err := parseNQuad(line)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Should not end with trailing period
	if strings.HasSuffix(triple.Subject, ".") {
		t.Error("Subject has trailing period - parsing did not trim properly")
	}
}

func TestRDFTriple_Structure(t *testing.T) {
	triple := &RDFTriple{
		Subject:   "http://example.org/s",
		Predicate: "http://example.org/p",
		Object:    "http://example.org/o",
		Graph:     "http://example.org/g",
	}

	if triple.Subject == "" {
		t.Error("Subject not set")
	}
	if triple.Predicate == "" {
		t.Error("Predicate not set")
	}
	if triple.Object == "" {
		t.Error("Object not set")
	}
	if triple.Graph == "" {
		t.Error("Graph not set")
	}
}

func TestLoadQuads_DefaultGraphHandling(t *testing.T) {
	nquads := `<http://example.org/alice> <http://example.org/name> "Alice" <http://example.org/g1> .`

	loader := NewQuadLoader()
	loader.DefaultGraph = "http://custom.default"
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Graph membership triple should use the specified graph
	if len(triples) >= 2 {
		graphTriple := triples[1]
		if graphTriple.Graph != "http://example.org/g1" {
			t.Errorf("Expected graph membership triple to have graph 'http://example.org/g1', got '%s'", graphTriple.Graph)
		}
	}
}

func TestToMutationJSON_SingleSubject(t *testing.T) {
	triples := []*RDFTriple{
		{
			Subject:   "http://example.org/alice",
			Predicate: "name",
			Object:    "\"Alice\"",
			Graph:     "http://example.org/g1",
		},
		{
			Subject:   "http://example.org/alice",
			Predicate: "age",
			Object:    "30",
			Graph:     "http://example.org/g1",
		},
	}

	loader := NewQuadLoader()
	json := loader.ToMutationJSON(triples)

	if !strings.Contains(json, "[") {
		t.Error("JSON should start with array bracket")
	}
	if !strings.Contains(json, "name") {
		t.Error("JSON should contain predicate 'name'")
	}
	if !strings.Contains(json, "age") {
		t.Error("JSON should contain predicate 'age'")
	}
	if !strings.Contains(json, "Alice") {
		t.Error("JSON should contain object value 'Alice'")
	}
}

func TestToMutationJSON_MultipleSubjects(t *testing.T) {
	triples := []*RDFTriple{
		{
			Subject:   "http://example.org/alice",
			Predicate: "name",
			Object:    "\"Alice\"",
			Graph:     "http://example.org/g1",
		},
		{
			Subject:   "http://example.org/bob",
			Predicate: "name",
			Object:    "\"Bob\"",
			Graph:     "http://example.org/g2",
		},
	}

	loader := NewQuadLoader()
	json := loader.ToMutationJSON(triples)

	if !strings.Contains(json, "Alice") {
		t.Error("JSON should contain 'Alice'")
	}
	if !strings.Contains(json, "Bob") {
		t.Error("JSON should contain 'Bob'")
	}
	// Should have 2 objects in array
	count := strings.Count(json, "uid")
	if count < 2 {
		t.Errorf("Expected at least 2 subjects in JSON, got %d", count)
	}
}

func TestToMutationJSON_EmptyTriples(t *testing.T) {
	loader := NewQuadLoader()
	json := loader.ToMutationJSON([]*RDFTriple{})

	if !strings.Contains(json, "[") || !strings.Contains(json, "]") {
		t.Error("JSON should be valid empty array")
	}
}

func TestEscapeSubject(t *testing.T) {
	tests := []struct {
		input    string
		hasValid bool
	}{
		{"http://example.org/alice", true},
		{"alice123", true},
		{"_:blank", true},
		{"alice-bob", false}, // hyphen gets escaped
		{"alice/bob", false}, // slash gets escaped
	}

	for _, tt := range tests {
		result := escapeSubject(tt.input)
		if result == "" {
			t.Errorf("escapeSubject(%s) returned empty string", tt.input)
		}
	}
}

func TestEscapePredicate(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"http://example.org/name", "name"},
		{"http://example.org/knows", "knows"},
		{"http://ex.org/isChildOf", "isChildOf"},
		{"name", "name"},
	}

	for _, tt := range tests {
		result := escapePredicate(tt.input)
		if result != tt.expected {
			t.Errorf("escapePredicate(%s): expected '%s', got '%s'", tt.input, tt.expected, result)
		}
	}
}

func TestEscapePredicateWithHash(t *testing.T) {
	input := "http://example.org#name"
	result := escapePredicate(input)
	if result != "name" {
		t.Errorf("escapePredicate should handle # separator, expected 'name', got '%s'", result)
	}
}

func TestLoadQuads_LargeDataset(t *testing.T) {
	// Create 100 quads
	var sb strings.Builder
	for i := 0; i < 100; i++ {
		sb.WriteString(fmt.Sprintf("<http://example.org/entity/%d> <http://example.org/prop> \"value%d\" <http://example.org/g1> .\n", i, i))
	}

	loader := NewQuadLoader()
	reader := strings.NewReader(sb.String())

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// 100 quads * 2 triples = 200
	if len(triples) != 200 {
		t.Errorf("Expected 200 triples from 100 quads, got %d", len(triples))
	}
}

func TestLoadQuads_SpecialCharactersInLiterals(t *testing.T) {
	// Test with special characters in literal values
	nquads := `<http://example.org/alice> <http://example.org/bio> "Alice \"Ally\" Smith" <http://example.org/g1> .`

	loader := NewQuadLoader()
	reader := strings.NewReader(nquads)

	triples, err := loader.LoadQuads(reader)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if len(triples) == 0 {
		t.Fatal("Expected triples")
	}

	// Object should preserve the literal
	if !strings.Contains(triples[0].Object, "Alice") {
		t.Errorf("Expected literal to be preserved in object")
	}
}
