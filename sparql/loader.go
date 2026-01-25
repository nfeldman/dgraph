package sparql

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// LoaderExample demonstrates how to convert RDF quads to the Dgraph data model
// with graph membership predicates.
//
// Example usage:
//
//	loader := NewQuadLoader()
//	triples, err := loader.LoadQuads(reader)
//	// ... insert triples into Dgraph using mutation API

// QuadLoader processes RDF quads and emits triples with graph predicates.
type QuadLoader struct {
	// DefaultGraph is used when a triple (not quad) is loaded
	DefaultGraph string
}

// NewQuadLoader creates a new loader with optional default graph
func NewQuadLoader() *QuadLoader {
	return &QuadLoader{
		DefaultGraph: "urn:x-arq:DefaultGraph",
	}
}

// RDFTriple represents a single RDF triple with optional graph context
type RDFTriple struct {
	Subject   string
	Predicate string
	Object    string
	Graph     string // Empty for default graph
}

// LoadQuads reads N-Quads format and produces triples with graph predicates.
//
// For each quad:
//
//	<subject> <predicate> <object> <graph> .
//
// We emit TWO triples:
//
//	<subject> <predicate> <object> .
//	<subject> <graph-predicate> "graph-iri" .
//
// This allows Dgraph to filter by graph membership using the translator's logic.
func (l *QuadLoader) LoadQuads(r io.Reader) ([]*RDFTriple, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return nil, fmt.Errorf("reading input: %w", err)
	}

	// Normalize the stream so that concatenated quads without newlines are split correctly.
	normalized := normalizeNQuadStream(string(data))
	scanner := bufio.NewScanner(strings.NewReader(normalized))
	var triples []*RDFTriple

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		quad, err := parseNQuad(line)
		if err != nil {
			return nil, fmt.Errorf("parsing line %q: %w", line, err)
		}

		// Add the main triple
		triples = append(triples, &RDFTriple{
			Subject:   quad.Subject,
			Predicate: quad.Predicate,
			Object:    quad.Object,
			Graph:     quad.Graph,
		})

		// Add the graph membership triple if graph is specified
		if quad.Graph != "" && quad.Graph != l.DefaultGraph {
			triples = append(triples, &RDFTriple{
				Subject:   quad.Subject,
				Predicate: "graph", // Special predicate for graph membership
				Object:    fmt.Sprintf("\"%s\"", quad.Graph),
				Graph:     quad.Graph,
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("reading normalized input: %w", err)
	}

	return triples, nil
}

// normalizeNQuadStream inserts line breaks between quads even when the input
// stream omits newlines (e.g., `... .<next>`). This keeps scanning simple while
// remaining tolerant of compacted datasets used in tests.
func normalizeNQuadStream(input string) string {
	var sb strings.Builder
	runes := []rune(input)
	for i := 0; i < len(runes); i++ {
		r := runes[i]
		sb.WriteRune(r)
		if r == '.' {
			// Look ahead to see if next non-space character starts another quad.
			j := i + 1
			for j < len(runes) {
				if runes[j] == ' ' || runes[j] == '\t' || runes[j] == '\n' || runes[j] == '\r' {
					j++
					continue
				}
				break
			}
			if j < len(runes) && runes[j] == '<' {
				sb.WriteRune('\n')
			}
		}
	}
	return sb.String()
}

// parseNQuad is a simplified N-Quads parser for demonstration.
// In production, use a proper RDF parsing library.
func parseNQuad(line string) (*RDFTriple, error) {
	// Very basic parser - expects: <s> <p> <o> [<g>] .
	line = strings.TrimSuffix(strings.TrimSpace(line), ".")
	parts := strings.Fields(line)

	if len(parts) < 3 {
		return nil, fmt.Errorf("insufficient parts in quad: %s", line)
	}

	triple := &RDFTriple{
		Subject:   parts[0],
		Predicate: parts[1],
		Object:    strings.Join(parts[2:len(parts)], " "),
	}

	// Check if there's a graph (4th component)
	// Simple heuristic: if we have exactly 4 parts and the 4th starts with <
	if len(parts) == 4 && strings.HasPrefix(parts[3], "<") {
		triple.Object = parts[2]
		triple.Graph = strings.Trim(parts[3], "<>")
	}

	// Clean up IRIs
	triple.Subject = strings.Trim(triple.Subject, "<>")
	triple.Predicate = strings.Trim(triple.Predicate, "<>")

	// Object might be IRI or literal - keep as-is for now

	return triple, nil
}

// ToMutationJSON converts triples to Dgraph JSON mutation format
func (l *QuadLoader) ToMutationJSON(triples []*RDFTriple) string {
	var sb strings.Builder
	sb.WriteString("[\n")

	// Group by subject
	subjectMap := make(map[string][]*RDFTriple)
	for _, t := range triples {
		subjectMap[t.Subject] = append(subjectMap[t.Subject], t)
	}

	first := true
	for subject, subjectTriples := range subjectMap {
		if !first {
			sb.WriteString(",\n")
		}
		first = false

		sb.WriteString("  {\n")
		sb.WriteString(fmt.Sprintf("    \"uid\": \"_:%s\",\n", escapeSubject(subject)))

		for i, t := range subjectTriples {
			pred := escapePredicate(t.Predicate)
			obj := t.Object

			sb.WriteString(fmt.Sprintf("    \"%s\": %s", pred, obj))
			if i < len(subjectTriples)-1 {
				sb.WriteString(",")
			}
			sb.WriteString("\n")
		}

		sb.WriteString("  }")
	}

	sb.WriteString("\n]")
	return sb.String()
}

func escapeSubject(s string) string {
	// Simple escape - replace non-alphanumeric with underscore
	return strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, s)
}

func escapePredicate(p string) string {
	// Extract local name from IRI. Prefer fragment (#) over path (/).
	if idx := strings.LastIndex(p, "#"); idx >= 0 {
		return p[idx+1:]
	}
	if idx := strings.LastIndex(p, "/"); idx >= 0 {
		return p[idx+1:]
	}
	return p
}
