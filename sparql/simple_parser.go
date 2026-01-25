package sparql

import (
	"fmt"
	"regexp"
	"strings"
)

// simpleParse is a lightweight, best-effort SPARQL parser used for tests.
// It is not a full SPARQL implementation but supports the constructs used
// throughout the current unit tests (SELECT/ASK, PREFIX, FROM/FROM NAMED,
// basic WHERE triple patterns, OPTIONAL/UNION blocks, ORDER BY, DISTINCT,
// aggregates, BIND, and HAVING clauses).
func simpleParse(query string) (SPARQLQuery, error) {
	q := strings.TrimSpace(query)
	if q == "" {
		return nil, fmt.Errorf("empty query")
	}

	// Strip comment-only lines that start with # (SPARQL comment style).
	var lines []string
	for _, line := range strings.Split(q, "\n") {
		trim := strings.TrimSpace(line)
		if strings.HasPrefix(trim, "#") {
			continue
		}
		lines = append(lines, line)
	}
	q = strings.Join(lines, "\n")

	upper := strings.ToUpper(q)
	selectIdx := strings.Index(upper, "SELECT")
	askIdx := strings.Index(upper, "ASK")

	qtype := ""
	switch {
	case selectIdx != -1:
		qtype = "SELECT"
	case askIdx != -1:
		qtype = "ASK"
	default:
		return nil, fmt.Errorf("unsupported or missing query type")
	}

	if qtype == "SELECT" && !strings.Contains(upper, "WHERE") {
		return nil, fmt.Errorf("missing WHERE clause")
	}

	// Extract PREFIX declarations.
	prefixes := make(map[string]string)
	prefixRe := regexp.MustCompile(`(?im)^\s*PREFIX\s+([A-Za-z0-9_\-]+):\s*<([^>]+)>`)
	for _, m := range prefixRe.FindAllStringSubmatch(q, -1) {
		prefixes[m[1]] = m[2]
	}

	// FROM and FROM NAMED clauses.
	var fromGraphs, fromNamedGraphs []string
	fromNamedRe := regexp.MustCompile(`(?i)FROM\s+NAMED\s+<([^>]+)>`)
	for _, m := range fromNamedRe.FindAllStringSubmatch(q, -1) {
		fromNamedGraphs = append(fromNamedGraphs, strings.TrimSpace(m[1]))
	}
	fromRe := regexp.MustCompile(`(?i)FROM\s+<([^>]+)>`)
	for _, m := range fromRe.FindAllStringSubmatch(q, -1) {
		fromGraphs = append(fromGraphs, strings.TrimSpace(m[1]))
	}

	// Determine SELECT clause boundaries to collect projections and modifiers.
	projs := []string{}
	distinct := false
	orderBy := []string{}
	aggregates := []*Aggregate{}
	binds := []*BindExpression{}
	var having *HavingClause

	if qtype == "SELECT" {
		start := selectIdx + len("SELECT")
		end := len(q)
		keywords := []string{" WHERE", "\nWHERE", " FROM", "\nFROM", " ORDER", "\nORDER", " GROUP", "\nGROUP", " HAVING", "\nHAVING"}
		for _, kw := range keywords {
			if idx := strings.Index(upper[start:], kw); idx != -1 {
				candidate := start + idx
				if candidate < end {
					end = candidate
				}
			}
		}
		selectClause := strings.TrimSpace(q[start:end])
		selectUpper := strings.ToUpper(selectClause)
		if strings.HasPrefix(selectUpper, "DISTINCT") {
			distinct = true
			selectClause = strings.TrimSpace(selectClause[len("DISTINCT"):])
		}

		// Capture aggregate projections with optional DISTINCT.
		aggRe := regexp.MustCompile(`(?i)(COUNT|SUM|AVG|MIN|MAX)\s*\(\s*(DISTINCT\s+)?(\?[A-Za-z0-9_]+)\s*\)\s+AS\s+(\?[A-Za-z0-9_]+)`) //nolint:lll
		for _, m := range aggRe.FindAllStringSubmatch(selectClause, -1) {
			aggregates = append(aggregates, &Aggregate{
				Function: strings.ToLower(m[1]),
				Variable: m[3],
				Alias:    m[4],
				Distinct: strings.TrimSpace(strings.ToUpper(m[2])) == "DISTINCT",
			})
			projs = append(projs, m[4])
		}

		// Capture projection variables (including *).
		if strings.Contains(selectClause, "*") {
			projs = append(projs, "*")
		}
		varRe := regexp.MustCompile(`\?[A-Za-z0-9_]+`)
		for _, m := range varRe.FindAllString(selectClause, -1) {
			if strings.HasPrefix(m, "?") {
				projs = appendUnique(projs, m)
			}
		}
	}

	// ORDER BY variables (simple space-separated list).
	if idx := strings.Index(strings.ToUpper(q), "ORDER BY"); idx != -1 {
		rest := q[idx+len("ORDER BY"):]
		// Stop at next newline or LIMIT/OFFSET if present.
		if cut := strings.Index(strings.ToUpper(rest), "LIMIT"); cut != -1 {
			rest = rest[:cut]
		}
		if cut := strings.Index(strings.ToUpper(rest), "OFFSET"); cut != -1 {
			rest = rest[:cut]
		}
		for _, tok := range strings.Fields(rest) {
			if strings.HasPrefix(tok, "?") {
				orderBy = append(orderBy, tok)
			}
		}
	}

	// HAVING clause (capture raw expression inside parentheses).
	havingRe := regexp.MustCompile(`(?i)HAVING\s*\(([^)]*)\)`) // simple, non-nested
	if m := havingRe.FindStringSubmatch(q); len(m) == 2 {
		having = &HavingClause{Expression: strings.TrimSpace(m[1])}
	}

	// BIND expressions: capture expression and target variable.
	bindRe := regexp.MustCompile(`(?i)BIND\s*\(([^)]*?)\s+AS\s+(\?[A-Za-z0-9_]+)\s*\)`)
	for _, m := range bindRe.FindAllStringSubmatch(q, -1) {
		binds = append(binds, &BindExpression{
			Expression: strings.TrimSpace(m[1]),
			Variable:   m[2],
		})
	}

	// Extract WHERE content and parse triples (including GRAPH blocks).
	open := strings.Index(q, "{")
	close := strings.LastIndex(q, "}")
	if open == -1 || close == -1 || close <= open {
		return nil, fmt.Errorf("malformed WHERE clause")
	}
	whereContent := q[open+1 : close]

	triples := []*Triple{}

	// GRAPH blocks first (to capture graph IRI/var).
	graphBlockRe := regexp.MustCompile(`(?is)GRAPH\s+([^\s{]+)\s*\{([^}]*)\}`)
	whereContent = graphBlockRe.ReplaceAllStringFunc(whereContent, func(match string) string {
		if sub := graphBlockRe.FindStringSubmatch(match); len(sub) == 3 {
			g := strings.TrimSpace(sub[1])
			triples = append(triples, extractTriples(sub[2], g)...)
		}
		return ""
	})

	// Remaining triples (default graph).
	triples = append(triples, extractTriples(whereContent, "")...)

	if len(triples) == 0 {
		return nil, fmt.Errorf("no triple patterns found")
	}

	queryImpl := &SPARQLQueryImpl{
		Qtype:      qtype,
		Prefixes:   prefixes,
		Projs:      projs,
		Bgps:       []*BGP{{Triples: triples}},
		Aggregates: aggregates,
		Binds:      binds,
		Having:     having,
		OrderBy:    orderBy,
		Distinct:   distinct,
		From:       fromGraphs,
		FromNamed:  fromNamedGraphs,
	}

	return queryImpl, nil
}

// extractTriples returns triples found in the given content, tagging them with graph if provided.
func extractTriples(content, graph string) []*Triple {
	tripleRe := regexp.MustCompile(`(?m)([?$][A-Za-z0-9_]+|<[^>]+>|_:[A-Za-z0-9_]+)\s+([^\s{}]+)\s+([?$][A-Za-z0-9_]+|<[^>]+>|"[^"]*"(?:@[A-Za-z-]+)?(?:\^\^<[^>]+>)?|_:[A-Za-z0-9_]+)`) //nolint:lll
	var triples []*Triple
	for _, m := range tripleRe.FindAllStringSubmatch(content, -1) {
		obj := strings.TrimSpace(m[3])
		triples = append(triples, &Triple{
			Subject:     strings.TrimSpace(m[1]),
			Predicate:   strings.TrimSpace(m[2]),
			Object:      obj,
			ObjectIsVar: strings.HasPrefix(obj, "?") || strings.HasPrefix(obj, "$"),
			Graph:       graph,
		})
	}
	return triples
}

// appendUnique appends v to slice if it is not already present.
func appendUnique(list []string, v string) []string {
	for _, existing := range list {
		if existing == v {
			return list
		}
	}
	return append(list, v)
}
