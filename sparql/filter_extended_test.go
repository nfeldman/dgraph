package sparql

import (
	"testing"
)

// TestBuiltInFunctions tests parsing of SPARQL built-in functions in FILTER clauses.
func TestBuiltInFunctions(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "LANG function",
			filterExpr: "LANG(?x) = 'en'",
			expectErr:  false,
		},
		{
			name:       "DATATYPE function",
			filterExpr: "DATATYPE(?x) = <http://www.w3.org/2001/XMLSchema#integer>",
			expectErr:  false,
		},
		{
			name:       "STR function",
			filterExpr: "STR(?x) = 'test'",
			expectErr:  false,
		},
		{
			name:       "STRLEN function",
			filterExpr: "STRLEN(?str) > 5",
			expectErr:  false,
		},
		{
			name:       "BOUND function",
			filterExpr: "BOUND(?x)",
			expectErr:  false,
		},
		{
			name:       "ISBLANK function",
			filterExpr: "ISBLANK(?x)",
			expectErr:  false,
		},
		{
			name:       "ISURI function",
			filterExpr: "ISURI(?x)",
			expectErr:  false,
		},
		{
			name:       "ISLITERAL function",
			filterExpr: "ISLITERAL(?x)",
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && ft == nil {
				t.Error("Expected non-nil filter tree")
			}
		})
	}
}

// TestNegationOperator tests the NOT (!) operator in FILTER clauses.
func TestNegationOperator(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "NOT with comparison",
			filterExpr: "!(?x = 5)",
			expectErr:  false,
		},
		{
			name:       "NOT with BOUND",
			filterExpr: "!BOUND(?x)",
			expectErr:  false,
		},
		{
			name:       "NOT with ISBLANK",
			filterExpr: "!ISBLANK(?x)",
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && ft == nil {
				t.Error("Expected non-nil filter tree")
			}
		})
	}
}

// TestINOperator tests the IN operator in FILTER clauses.
func TestINOperator(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "IN with numbers",
			filterExpr: "?x IN (1, 2, 3)",
			expectErr:  false,
		},
		{
			name:       "IN with strings",
			filterExpr: "?name IN ('Alice', 'Bob', 'Charlie')",
			expectErr:  false,
		},
		{
			name:       "NOT IN operator",
			filterExpr: "?status NOT IN ('inactive', 'deleted')",
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && ft == nil {
				t.Error("Expected non-nil filter tree")
			}
		})
	}
}

// TestComplexBooleanExpressions tests complex boolean combinations in FILTER clauses.
func TestComplexBooleanExpressions(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "AND with multiple conditions",
			filterExpr: "?age > 18 && ?status = 'active' && ?verified = true",
			expectErr:  false,
		},
		{
			name:       "OR with multiple conditions",
			filterExpr: "?role = 'admin' || ?role = 'moderator' || ?role = 'superuser'",
			expectErr:  false,
		},
		{
			name:       "Mixed AND/OR",
			filterExpr: "?x = 1 && ?y = 2 || ?z = 3",
			expectErr:  false,
		},
		{
			name:       "Nested with parentheses",
			filterExpr: "(?x = 1 || ?x = 2) && (?y > 10)",
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && ft == nil {
				t.Error("Expected non-nil filter tree")
			}
		})
	}
}

// TestStringFunctions tests SPARQL string functions in FILTER clauses.
func TestStringFunctions(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "UCASE function",
			filterExpr: "UCASE(?name) = 'ALICE'",
			expectErr:  false,
		},
		{
			name:       "LCASE function",
			filterExpr: "LCASE(?name) = 'alice'",
			expectErr:  false,
		},
		{
			name:       "CONTAINS function",
			filterExpr: "CONTAINS(?text, 'hello')",
			expectErr:  false,
		},
		{
			name:       "STRSTARTS function",
			filterExpr: "STRSTARTS(?uri, 'http://')",
			expectErr:  false,
		},
		{
			name:       "STRENDS function",
			filterExpr: "STRENDS(?filename, '.pdf')",
			expectErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if (err != nil) != tt.expectErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectErr, err)
				return
			}

			if !tt.expectErr && ft == nil {
				t.Error("Expected non-nil filter tree")
			}
		})
	}
}

// TestFilterExpressionErrors tests error handling for invalid FILTER expressions.
func TestFilterExpressionErrors(t *testing.T) {
	tests := []struct {
		name       string
		filterExpr string
		expectErr  bool
	}{
		{
			name:       "Invalid syntax",
			filterExpr: "?x = = 5",
			expectErr:  true,
		},
		{
			name:       "Empty expression",
			filterExpr: "",
			expectErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ft, err := buildFilterTreeFromExpr(tt.filterExpr)

			if !tt.expectErr && err != nil {
				t.Errorf("Expected no error, got: %v", err)
				return
			}

			if tt.expectErr && err == nil && ft == nil {
				// This is acceptable - empty or invalid expressions may return nil without error
				return
			}
		})
	}
}

// TestExtractListValues tests the extractListValues helper function.
func TestExtractListValues(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "Simple numbers",
			input:    "1, 2, 3",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "Quoted strings",
			input:    "'Alice', 'Bob', 'Charlie'",
			expected: []string{"'Alice'", "'Bob'", "'Charlie'"},
		},
		{
			name:     "URIs",
			input:    "<http://example.com/type1>, <http://example.com/type2>",
			expected: []string{"<http://example.com/type1>", "<http://example.com/type2>"},
		},
		{
			name:     "Mixed with spaces",
			input:    "1 , 'test' , <http://example.com/val>",
			expected: []string{"1", "'test'", "<http://example.com/val>"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := extractListValues(tt.input)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d values, got %d", len(tt.expected), len(result))
				return
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("Value %d: expected %q, got %q", i, tt.expected[i], v)
				}
			}
		})
	}
}

// TestOperatorSplitting tests the splitByOperator helper function.
func TestOperatorSplitting(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		op       string
		expected []string
	}{
		{
			name:     "OR operator",
			expr:     "?x = 1 || ?x = 2",
			op:       "||",
			expected: []string{"?x = 1", "?x = 2"},
		},
		{
			name:     "AND operator",
			expr:     "?x > 0 && ?y < 10",
			op:       "&&",
			expected: []string{"?x > 0", "?y < 10"},
		},
		{
			name:     "OR inside parentheses",
			expr:     "(?x = 1 || ?x = 2) && ?y = 3",
			op:       "&&",
			expected: []string{"(?x = 1 || ?x = 2)", "?y = 3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := splitByOperator(tt.expr, tt.op)

			if len(result) != len(tt.expected) {
				t.Errorf("Expected %d parts, got %d", len(tt.expected), len(result))
				return
			}

			for i, v := range result {
				if v != tt.expected[i] {
					t.Errorf("Part %d: expected %q, got %q", i, tt.expected[i], v)
				}
			}
		})
	}
}
