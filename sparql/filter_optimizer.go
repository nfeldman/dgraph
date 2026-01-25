package sparql

import (
	"regexp"
	"strings"
)

// FilterExpressionOptimizer simplifies and optimizes SPARQL FILTER expressions.
// It applies transformations like:
// - Constant folding: 1 + 2 > 2 → true
// - Boolean simplification: !(a && b) → !a || !b (De Morgan's laws)
// - Remove always-true filters
// - Detect always-false filters (convert to Empty)
// - Combine adjacent filters
type FilterExpressionOptimizer struct{}

// NewFilterExpressionOptimizer creates a new filter expression optimizer.
func NewFilterExpressionOptimizer() *FilterExpressionOptimizer {
	return &FilterExpressionOptimizer{}
}

// Optimize applies filter expression optimizations to the algebra tree.
func (o *FilterExpressionOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraFilter:
		return o.optimizeFilter(node)
	default:
		return o.optimizeChildren(expr)
	}
}

// optimizeFilter handles a single filter expression.
func (o *FilterExpressionOptimizer) optimizeFilter(filter *AlgebraFilter) (AlgebraExpr, error) {
	// Optimize the filter expression
	optimizedExpr := o.simplifyExpression(filter.Expr)

	// Check for special cases
	if optimizedExpr == "true" {
		// Always-true filter can be removed
		return o.Optimize(filter.Input)
	}

	if optimizedExpr == "false" {
		// Always-false filter converts to Empty
		return &AlgebraEmpty{}, nil
	}

	// Recursively optimize the input
	optimizedInput, err := o.Optimize(filter.Input)
	if err != nil {
		return nil, err
	}

	// Handle cascaded filters
	if nextFilter, ok := optimizedInput.(*AlgebraFilter); ok {
		// Merge adjacent filters with AND
		combined := o.combineFilters(optimizedExpr, nextFilter.Expr)
		return &AlgebraFilter{
			Expr:  combined,
			Input: nextFilter.Input,
		}, nil
	}

	return &AlgebraFilter{
		Expr:  optimizedExpr,
		Input: optimizedInput,
	}, nil
}

// optimizeChildren recursively optimizes child expressions.
func (o *FilterExpressionOptimizer) optimizeChildren(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraProject:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraProject{
			Vars:  node.Vars,
			Input: optimizedInput,
		}, nil

	case *AlgebraJoin:
		optimizedLeft, err := o.Optimize(node.Left)
		if err != nil {
			return nil, err
		}
		optimizedRight, err := o.Optimize(node.Right)
		if err != nil {
			return nil, err
		}
		return &AlgebraJoin{
			Left:  optimizedLeft,
			Right: optimizedRight,
		}, nil

	case *AlgebraAgg:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraAgg{
			Op:    node.Op,
			Var:   node.Var,
			Group: node.Group,
			Input: optimizedInput,
		}, nil

	case *AlgebraDistinct:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraDistinct{
			Input: optimizedInput,
		}, nil

	case *AlgebraOrderBy:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraOrderBy{
			Exprs:     node.Exprs,
			Ascending: node.Ascending,
			Input:     optimizedInput,
		}, nil

	case *AlgebraLimit:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraLimit{
			Count:  node.Count,
			Offset: node.Offset,
			Input:  optimizedInput,
		}, nil

	default:
		return expr, nil
	}
}

// simplifyExpression applies logical simplifications to a filter expression.
func (o *FilterExpressionOptimizer) simplifyExpression(expr string) string {
	expr = strings.TrimSpace(expr)

	// Check for trivial comparisons
	if o.isTrivialComparison(expr) {
		return o.evaluateTrivialComparison(expr)
	}

	// Apply De Morgan's laws: !(A && B) → !A || !B
	if strings.HasPrefix(expr, "!(") && strings.Contains(expr, "&&") {
		expr = o.applyDeMorgan(expr)
	}

	// Remove double negation: !!A → A
	for strings.HasPrefix(expr, "!!") {
		expr = strings.TrimPrefix(expr, "!!")
	}

	// Simplify identity conditions: A && true → A
	if strings.Contains(expr, "&&") {
		expr = o.simplifyIdentity(expr)
	}

	return expr
}

// isTrivialComparison checks if an expression is trivially evaluable.
func (o *FilterExpressionOptimizer) isTrivialComparison(expr string) bool {
	// Pattern: constant op constant
	// e.g., "1 > 2", "5 = 5", "10 < 20"
	parts := strings.FieldsFunc(expr, func(r rune) bool {
		return r == '>' || r == '<' || r == '=' || r == '!'
	})

	if len(parts) != 2 {
		return false
	}

	left := strings.TrimSpace(parts[0])
	right := strings.TrimSpace(parts[1])

	return o.isConstant(left) && o.isConstant(right)
}

// isConstant checks if a value is a constant (number or string literal).
func (o *FilterExpressionOptimizer) isConstant(s string) bool {
	s = strings.TrimSpace(s)

	// Check for number
	if matched, _ := regexp.MatchString(`^-?\d+(\.\d+)?$`, s); matched {
		return true
	}

	// Check for string literal
	if strings.HasPrefix(s, "\"") && strings.HasSuffix(s, "\"") {
		return true
	}

	// Check for single-quoted string
	if strings.HasPrefix(s, "'") && strings.HasSuffix(s, "'") {
		return true
	}

	return false
}

// evaluateTrivialComparison evaluates a trivial comparison.
func (o *FilterExpressionOptimizer) evaluateTrivialComparison(expr string) string {
	// Extract operator and operands
	var op string
	var left, right string

	if strings.Contains(expr, ">=") {
		parts := strings.Split(expr, ">=")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), ">="
		}
	} else if strings.Contains(expr, "<=") {
		parts := strings.Split(expr, "<=")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), "<="
		}
	} else if strings.Contains(expr, "!=") {
		parts := strings.Split(expr, "!=")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), "!="
		}
	} else if strings.Contains(expr, ">") {
		parts := strings.Split(expr, ">")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), ">"
		}
	} else if strings.Contains(expr, "<") {
		parts := strings.Split(expr, "<")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), "<"
		}
	} else if strings.Contains(expr, "=") {
		parts := strings.Split(expr, "=")
		if len(parts) == 2 {
			left, right, op = strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), "="
		}
	}

	// If we couldn't parse it, return as-is
	if op == "" {
		return expr
	}

	// For string comparison, only handle equality
	if (strings.HasPrefix(left, "\"") || strings.HasPrefix(left, "'")) &&
		(strings.HasPrefix(right, "\"") || strings.HasPrefix(right, "'")) {
		if op == "=" {
			if left == right {
				return "true"
			}
			return "false"
		}
		return expr
	}

	// Parse numbers
	leftNum, err1 := parseNumber(left)
	rightNum, err2 := parseNumber(right)

	if err1 != nil || err2 != nil {
		return expr
	}

	// Evaluate
	var result bool
	switch op {
	case ">":
		result = leftNum > rightNum
	case "<":
		result = leftNum < rightNum
	case ">=":
		result = leftNum >= rightNum
	case "<=":
		result = leftNum <= rightNum
	case "=":
		result = leftNum == rightNum
	case "!=":
		result = leftNum != rightNum
	}

	if result {
		return "true"
	}
	return "false"
}

// parseNumber parses a number from a string.
func parseNumber(s string) (float64, error) {
	var f float64
	_, err := parseFloat(s, &f)
	return f, err
}

// parseFloat is a simple numeric parser.
func parseFloat(s string, f *float64) (int, error) {
	s = strings.TrimSpace(s)
	i := 0
	hasDecimal := false
	sign := 1.0

	if i < len(s) && (s[i] == '-' || s[i] == '+') {
		if s[i] == '-' {
			sign = -1.0
		}
		i++
	}

	var num float64
	for i < len(s) && ((s[i] >= '0' && s[i] <= '9') || s[i] == '.') {
		if s[i] == '.' {
			hasDecimal = true
		} else {
			digit := float64(s[i] - '0')
			if hasDecimal {
				num = num + digit/10
			} else {
				num = num*10 + digit
			}
		}
		i++
	}

	*f = sign * num
	return i, nil
}

// applyDeMorgan applies De Morgan's laws to negate a conjunction.
func (o *FilterExpressionOptimizer) applyDeMorgan(expr string) string {
	// !(A && B) → (!A || !B)
	if strings.HasPrefix(expr, "!(") && strings.HasSuffix(expr, ")") {
		inner := expr[2 : len(expr)-1]
		if strings.Contains(inner, "&&") {
			parts := strings.Split(inner, "&&")
			if len(parts) == 2 {
				left := strings.TrimSpace(parts[0])
				right := strings.TrimSpace(parts[1])
				return "(" + "!" + left + " || " + "!" + right + ")"
			}
		}
	}
	return expr
}

// simplifyIdentity simplifies expressions with trivial identities.
func (o *FilterExpressionOptimizer) simplifyIdentity(expr string) string {
	// A && true → A
	if strings.Contains(expr, "&& true") {
		expr = strings.ReplaceAll(expr, "&& true", "")
		expr = strings.ReplaceAll(expr, "true &&", "")
	}

	// A || false → A
	if strings.Contains(expr, "|| false") {
		expr = strings.ReplaceAll(expr, "|| false", "")
		expr = strings.ReplaceAll(expr, "false ||", "")
	}

	return strings.TrimSpace(expr)
}

// combineFilters combines two adjacent filters with AND.
func (o *FilterExpressionOptimizer) combineFilters(expr1, expr2 string) string {
	return "(" + expr1 + " && " + expr2 + ")"
}
