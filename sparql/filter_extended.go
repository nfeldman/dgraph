package sparql

import (
	"fmt"
	"regexp"
	"strings"

	dql "github.com/dgraph-io/dgraph/v25/dql"
)

// parseINOperator parses the IN operator expression.
// Example: ?x IN (1, 2, 3)
func parseINOperator(expr string) (*dql.FilterTree, error) {
	// Pattern: ?var IN (val1, val2, ...)
	inRe := regexp.MustCompile(`^([?$]\w+)\s+IN\s+\((.+)\)$`)
	m := inRe.FindStringSubmatch(expr)
	if len(m) != 3 {
		return nil, fmt.Errorf("invalid IN expression: %s", expr)
	}

	varName := strings.TrimLeft(m[1], "?$")
	valuesStr := m[2]

	// Parse values
	values := extractListValues(valuesStr)
	var dqlArgs []dql.Arg
	for _, val := range values {
		dqlArgs = append(dqlArgs, dql.Arg{Value: val})
	}

	return &dql.FilterTree{
		Func: &dql.Function{
			Name: "in",
			Attr: varName,
			Args: dqlArgs,
		},
	}, nil
}

// parseNOTINOperator parses the NOT IN operator expression.
// Example: ?x NOT IN (1, 2, 3)
func parseNOTINOperator(expr string) (*dql.FilterTree, error) {
	notInRe := regexp.MustCompile(`^([?$]\w+)\s+NOT\s+IN\s+\((.+)\)$`)
	m := notInRe.FindStringSubmatch(expr)
	if len(m) != 3 {
		return nil, fmt.Errorf("invalid NOT IN expression: %s", expr)
	}

	varName := strings.TrimLeft(m[1], "?$")
	valuesStr := m[2]

	// Parse values
	values := extractListValues(valuesStr)
	var dqlArgs []dql.Arg
	for _, val := range values {
		dqlArgs = append(dqlArgs, dql.Arg{Value: val})
	}

	return &dql.FilterTree{
		Func: &dql.Function{
			Name: "notin",
			Attr: varName,
			Args: dqlArgs,
		},
	}, nil
}

// extractListValues extracts comma-separated values from a list string, respecting quotes and URIs.
// Example: "'val1', 'val2', <http://example.com/val3>" -> ["'val1'", "'val2'", "<http://example.com/val3>"]
func extractListValues(listStr string) []string {
	var values []string
	var current strings.Builder
	inQuotes := false
	inURI := false
	quoteChar := rune(0)

	for _, ch := range listStr {
		if !inQuotes && !inURI {
			if ch == '\'' || ch == '"' {
				inQuotes = true
				quoteChar = ch
				current.WriteRune(ch)
			} else if ch == '<' {
				inURI = true
				current.WriteRune(ch)
			} else if ch == ',' {
				if current.Len() > 0 {
					values = append(values, strings.TrimSpace(current.String()))
					current.Reset()
				}
			} else if !isWhitespace(ch) {
				current.WriteRune(ch)
			}
		} else if inQuotes {
			current.WriteRune(ch)
			if ch == quoteChar {
				inQuotes = false
			}
		} else if inURI {
			current.WriteRune(ch)
			if ch == '>' {
				inURI = false
			}
		} else {
			current.WriteRune(ch)
		}
	}

	if current.Len() > 0 {
		values = append(values, strings.TrimSpace(current.String()))
	}

	return values
}

// isWhitespace checks if a rune is whitespace.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// parseBuiltInFunctionDirect tries to parse built-in functions from the expression.
// Supports: LANG(), DATATYPE(), STR(), STRLEN(), UCASE(), LCASE(), CONTAINS(),
// STRSTARTS(), STRENDS(), BOUND(), ISBLANK(), ISURI(), ISLITERAL(), ISNUMERIC()
func parseBuiltInFunctionDirect(expr string) *dql.FilterTree {
	exprUpper := strings.ToUpper(expr)

	// List of built-in functions and their DQL function names
	builtIns := map[string]string{
		"LANG(":      "lang",
		"DATATYPE(":  "datatype",
		"STR(":       "str",
		"STRLEN(":    "strlen",
		"UCASE(":     "ucase",
		"LCASE(":     "lcase",
		"CONTAINS(":  "contains",
		"STRSTARTS(": "strstarts",
		"STRENDS(":   "strends",
		"BOUND(":     "bound",
		"ISBLANK(":   "isblank",
		"ISURI(":     "isuri",
		"ISLITERAL(": "isliteral",
		"ISNUMERIC(": "isnumeric",
	}

	for fnPrefix, fnName := range builtIns {
		if strings.HasPrefix(exprUpper, fnPrefix) {
			// Extract the arguments
			if closeIdx := strings.LastIndex(exprUpper, ")"); closeIdx > 0 {
				argsStr := expr[len(fnPrefix):closeIdx]
				args := extractListValues(argsStr)

				var dqlArgs []dql.Arg
				for _, arg := range args {
					dqlArgs = append(dqlArgs, dql.Arg{Value: arg})
				}

				return &dql.FilterTree{
					Func: &dql.Function{
						Name: fnName,
						Args: dqlArgs,
					},
				}
			}
		}
	}

	return nil
}

// parseINOperatorDirect tries to parse an IN operator expression directly.
func parseINOperatorDirect(expr string) *dql.FilterTree {
	// Pattern: ?var IN (val1, val2, ...)
	inRe := regexp.MustCompile(`^\s*([?$]\w+)\s+IN\s+\((.+)\)\s*$`)
	m := inRe.FindStringSubmatch(expr)
	if len(m) != 3 {
		return nil
	}

	varName := strings.TrimLeft(m[1], "?$")
	valuesStr := m[2]

	// Parse values (respecting quotes)
	values := extractListValues(valuesStr)
	var dqlArgs []dql.Arg
	for _, val := range values {
		dqlArgs = append(dqlArgs, dql.Arg{Value: val})
	}

	return &dql.FilterTree{
		Func: &dql.Function{
			Name: "in",
			Attr: varName,
			Args: dqlArgs,
		},
	}
}

// parseNOTINOperatorDirect tries to parse a NOT IN operator expression directly.
func parseNOTINOperatorDirect(expr string) *dql.FilterTree {
	// Pattern: ?var NOT IN (val1, val2, ...)
	notInRe := regexp.MustCompile(`^\s*([?$]\w+)\s+NOT\s+IN\s+\((.+)\)\s*$`)
	m := notInRe.FindStringSubmatch(expr)
	if len(m) != 3 {
		return nil
	}

	varName := strings.TrimLeft(m[1], "?$")
	valuesStr := m[2]

	// Parse values (respecting quotes)
	values := extractListValues(valuesStr)
	var dqlArgs []dql.Arg
	for _, val := range values {
		dqlArgs = append(dqlArgs, dql.Arg{Value: val})
	}

	return &dql.FilterTree{
		Func: &dql.Function{
			Name: "notin",
			Attr: varName,
			Args: dqlArgs,
		},
	}
}

// findOperatorOutsideParens finds the index of an operator that's not inside parentheses.
// Returns -1 if not found.
func findOperatorOutsideParens(expr, op string) int {
	depth := 0
	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' {
			depth++
		} else if expr[i] == ')' {
			depth--
		} else if depth == 0 && i+len(op) <= len(expr) && expr[i:i+len(op)] == op {
			return i
		}
	}
	return -1
}

// splitByOperator splits an expression by an operator that's not inside parentheses.
func splitByOperator(expr, op string) []string {
	var parts []string
	var current strings.Builder
	depth := 0

	for i := 0; i < len(expr); i++ {
		if expr[i] == '(' {
			depth++
			current.WriteByte(expr[i])
		} else if expr[i] == ')' {
			depth--
			current.WriteByte(expr[i])
		} else if depth == 0 && i+len(op) <= len(expr) && expr[i:i+len(op)] == op {
			// Found the operator at depth 0
			parts = append(parts, strings.TrimSpace(current.String()))
			current.Reset()
			i += len(op) - 1 // Skip the operator
		} else {
			current.WriteByte(expr[i])
		}
	}

	if current.Len() > 0 {
		parts = append(parts, strings.TrimSpace(current.String()))
	}

	return parts
}
