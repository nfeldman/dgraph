package sparql

import (
	"context"
	"testing"

	dql "github.com/dgraph-io/dgraph/v25/dql"
)

// TestOptionalPattern tests OPTIONAL pattern handling
func TestOptionalPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern *OptionalPattern
		wantErr bool
	}{
		{
			name: "Basic optional pattern",
			pattern: &OptionalPattern{
				Patterns: []GraphPattern{
					&BGP{
						Triples: []*Triple{{
							Subject:     "?s",
							Predicate:   "http://ex.org/name",
							Object:      "?name",
							ObjectIsVar: true,
						}},
					},
				},
				Required: false,
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}

			err := translateGraphPattern(tt.pattern, []string{}, []string{}, rootQuery, TranslateOptions{})
			if (err != nil) != tt.wantErr {
				t.Errorf("translateGraphPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestUnionPattern tests UNION pattern handling
func TestUnionPattern(t *testing.T) {
	tests := []struct {
		name    string
		pattern *UnionPattern
		wantErr bool
	}{
		{
			name: "Basic union pattern",
			pattern: &UnionPattern{
				Alternatives: [][]GraphPattern{
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/Type1",
								ObjectIsVar: false,
							}},
						},
					},
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/Type2",
								ObjectIsVar: false,
							}},
						},
					},
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}

			err := translateGraphPattern(tt.pattern, []string{}, []string{}, rootQuery, TranslateOptions{})
			if (err != nil) != tt.wantErr {
				t.Errorf("translateGraphPattern() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// TestOptionalPatternSemantics verifies that OPTIONAL patterns create separate subqueries
func TestOptionalPatternSemantics(t *testing.T) {
	tests := []struct {
		name        string
		pattern     *OptionalPattern
		expectChild bool
		wantErr     bool
	}{
		{
			name: "Optional with single BGP",
			pattern: &OptionalPattern{
				Required: false,
				Patterns: []GraphPattern{
					&BGP{
						Triples: []*Triple{{
							Subject:     "?s",
							Predicate:   "http://ex.org/age",
							Object:      "?age",
							ObjectIsVar: true,
						}},
					},
				},
			},
			expectChild: true,
			wantErr:     false,
		},
		{
			name: "Optional with multiple patterns",
			pattern: &OptionalPattern{
				Required: false,
				Patterns: []GraphPattern{
					&BGP{
						Triples: []*Triple{{
							Subject:     "?s",
							Predicate:   "http://ex.org/name",
							Object:      "?name",
							ObjectIsVar: true,
						}},
					},
					&BGP{
						Triples: []*Triple{{
							Subject:     "?s",
							Predicate:   "http://ex.org/email",
							Object:      "?email",
							ObjectIsVar: true,
						}},
					},
				},
			},
			expectChild: true,
			wantErr:     false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := translateGraphPattern(tt.pattern, []string{}, []string{}, rootQuery, TranslateOptions{})

			if (err != nil) != tt.wantErr {
				t.Errorf("translateGraphPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.expectChild && len(rootQuery.Children) == 0 {
				t.Errorf("Expected child query for OPTIONAL pattern, got none")
			}

			// Verify that the child query is marked as optional
			if tt.expectChild && len(rootQuery.Children) > 0 {
				optQuery := rootQuery.Children[0]
				if optQuery.Attr != "_optional" {
					t.Errorf("Expected child Attr='_optional', got %q", optQuery.Attr)
				}
			}
		})
	}
}

// TestUnionPatternSemantics verifies that UNION patterns create OR semantics
func TestUnionPatternSemantics(t *testing.T) {
	tests := []struct {
		name         string
		pattern      *UnionPattern
		expectFilter bool
		expectOpOr   bool
		wantErr      bool
	}{
		{
			name: "Union with two alternatives",
			pattern: &UnionPattern{
				Alternatives: [][]GraphPattern{
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/TypeA",
								ObjectIsVar: false,
							}},
						},
					},
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/TypeB",
								ObjectIsVar: false,
							}},
						},
					},
				},
			},
			expectFilter: true,
			expectOpOr:   true,
			wantErr:      false,
		},
		{
			name: "Union with three alternatives",
			pattern: &UnionPattern{
				Alternatives: [][]GraphPattern{
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/Type1",
								ObjectIsVar: false,
							}},
						},
					},
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/Type2",
								ObjectIsVar: false,
							}},
						},
					},
					{
						&BGP{
							Triples: []*Triple{{
								Subject:     "?s",
								Predicate:   "http://ex.org/type",
								Object:      "http://ex.org/Type3",
								ObjectIsVar: false,
							}},
						},
					},
				},
			},
			expectFilter: true,
			expectOpOr:   true,
			wantErr:      false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := translateGraphPattern(tt.pattern, []string{}, []string{}, rootQuery, TranslateOptions{})

			if (err != nil) != tt.wantErr {
				t.Errorf("translateGraphPattern() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Verify that child queries are created for each alternative
			expectedChildren := len(tt.pattern.Alternatives)
			if len(rootQuery.Children) != expectedChildren {
				t.Errorf("Expected %d child queries (one per alternative), got %d", expectedChildren, len(rootQuery.Children))
			}

			// Verify that a filter is created
			if tt.expectFilter && rootQuery.Filter == nil {
				t.Errorf("Expected filter for UNION pattern, got nil")
			}

			// Verify OR operator in filter if expected
			if tt.expectOpOr && rootQuery.Filter != nil && rootQuery.Filter.Op != "OR" {
				t.Errorf("Expected OR operator in filter, got %q", rootQuery.Filter.Op)
			}
		})
	}
}

// TestAggregates tests aggregate function support
func TestAggregates(t *testing.T) {
	tests := []struct {
		name       string
		aggregates []*Aggregate
		wantErr    bool
	}{
		{
			name: "COUNT aggregate",
			aggregates: []*Aggregate{
				{
					Function: "count",
					Variable: "?x",
					Alias:    "count_x",
					Distinct: false,
				},
			},
			wantErr: false,
		},
		{
			name: "Multiple aggregates",
			aggregates: []*Aggregate{
				{Function: "sum", Variable: "?value", Alias: "total"},
				{Function: "max", Variable: "?value", Alias: "maximum"},
				{Function: "min", Variable: "?value", Alias: "minimum"},
				{Function: "avg", Variable: "?value", Alias: "average"},
			},
			wantErr: false,
		},
		{
			name: "DISTINCT count",
			aggregates: []*Aggregate{
				{
					Function: "count",
					Variable: "?s",
					Alias:    "distinct_count",
					Distinct: true,
				},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := applyAggregates(rootQuery, tt.aggregates)

			if (err != nil) != tt.wantErr {
				t.Errorf("applyAggregates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if !rootQuery.IsGroupby {
					t.Error("Expected IsGroupby to be true")
				}
				if len(rootQuery.GroupbyAttrs) != len(tt.aggregates) {
					t.Errorf("Expected %d groupby attrs, got %d", len(tt.aggregates), len(rootQuery.GroupbyAttrs))
				}
				// Verify child queries for aggregates
				if len(rootQuery.Children) != len(tt.aggregates) {
					t.Errorf("Expected %d child queries for aggregates, got %d", len(tt.aggregates), len(rootQuery.Children))
				}
				// Verify each aggregate child query
				for i, child := range rootQuery.Children {
					if child.Var != tt.aggregates[i].Alias {
						t.Errorf("Child query %d: expected var %s, got %s", i, tt.aggregates[i].Alias, child.Var)
					}
				}
			}
		})
	}
}

// TestAggregatesWithDQLStructure tests that aggregates produce correct DQL structure
func TestAggregatesWithDQLStructure(t *testing.T) {
	tests := []struct {
		name            string
		aggregates      []*Aggregate
		expectedAttrs   []string
		expectedAliases []string
		wantErr         bool
	}{
		{
			name: "COUNT produces count(uid)",
			aggregates: []*Aggregate{
				{Function: "count", Variable: "?s", Alias: "cnt", Distinct: false},
			},
			expectedAttrs:   []string{"count(uid)"},
			expectedAliases: []string{"cnt"},
			wantErr:         false,
		},
		{
			name: "SUM produces sum(value)",
			aggregates: []*Aggregate{
				{Function: "sum", Variable: "?amount", Alias: "total_amount", Distinct: false},
			},
			expectedAttrs:   []string{"sum(amount)"},
			expectedAliases: []string{"total_amount"},
			wantErr:         false,
		},
		{
			name: "DISTINCT count produces distinct count(uid)",
			aggregates: []*Aggregate{
				{Function: "count", Variable: "?x", Alias: "distinct_cnt", Distinct: true},
			},
			expectedAttrs:   []string{"distinct count(uid)"},
			expectedAliases: []string{"distinct_cnt"},
			wantErr:         false,
		},
		{
			name: "MIN/MAX aggregates",
			aggregates: []*Aggregate{
				{Function: "min", Variable: "?val", Alias: "min_val", Distinct: false},
				{Function: "max", Variable: "?val", Alias: "max_val", Distinct: false},
			},
			expectedAttrs:   []string{"min(val)", "max(val)"},
			expectedAliases: []string{"min_val", "max_val"},
			wantErr:         false,
		},
		{
			name: "AVG aggregate",
			aggregates: []*Aggregate{
				{Function: "avg", Variable: "?price", Alias: "avg_price", Distinct: false},
			},
			expectedAttrs:   []string{"avg(price)"},
			expectedAliases: []string{"avg_price"},
			wantErr:         false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := applyAggregates(rootQuery, tt.aggregates)

			if (err != nil) != tt.wantErr {
				t.Errorf("applyAggregates() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify GroupbyAttrs match expectations
				if len(rootQuery.GroupbyAttrs) != len(tt.expectedAttrs) {
					t.Errorf("Expected %d groupby attrs, got %d", len(tt.expectedAttrs), len(rootQuery.GroupbyAttrs))
				}
				for i, attr := range rootQuery.GroupbyAttrs {
					if attr.Attr != tt.expectedAttrs[i] {
						t.Errorf("Groupby attr %d: expected %q, got %q", i, tt.expectedAttrs[i], attr.Attr)
					}
					if attr.Alias != tt.expectedAliases[i] {
						t.Errorf("Groupby alias %d: expected %q, got %q", i, tt.expectedAliases[i], attr.Alias)
					}
				}

				// Verify child queries with variables
				if len(rootQuery.Children) != len(tt.expectedAliases) {
					t.Errorf("Expected %d child queries, got %d", len(tt.expectedAliases), len(rootQuery.Children))
				}
				for i, child := range rootQuery.Children {
					if child.Var != tt.expectedAliases[i] {
						t.Errorf("Child query %d var: expected %q, got %q", i, tt.expectedAliases[i], child.Var)
					}
					// For COUNT, verify IsCount flag
					if tt.aggregates[i].Function == "count" {
						if !child.IsCount {
							t.Errorf("Child query %d: expected IsCount=true, got false", i)
						}
					}
				}
			}
		})
	}
}

// TestBindExpression tests variable binding support
func TestBindExpression(t *testing.T) {
	tests := []struct {
		name    string
		bind    *BindExpression
		wantErr bool
	}{
		{
			name: "Simple arithmetic bind (addition)",
			bind: &BindExpression{
				Expression: "?x + ?y",
				Variable:   "?sum",
			},
			wantErr: false,
		},
		{
			name: "Arithmetic with multiplication",
			bind: &BindExpression{
				Expression: "?x * ?y",
				Variable:   "?product",
			},
			wantErr: false,
		},
		{
			name: "Math function (SQRT)",
			bind: &BindExpression{
				Expression: "sqrt(?value)",
				Variable:   "?result",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := applyBindExpression(rootQuery, tt.bind)

			if (err != nil) != tt.wantErr {
				t.Errorf("applyBindExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify that a child query was added with the variable
				if len(rootQuery.Children) != 1 {
					t.Errorf("Expected 1 child query, got %d", len(rootQuery.Children))
				}
				child := rootQuery.Children[0]
				if child.Var != tt.bind.Variable {
					t.Errorf("Expected child var %q, got %q", tt.bind.Variable, child.Var)
				}
				if child.MathExp == nil {
					t.Error("Expected MathExp to be set")
				}
			}
		})
	}
}

// TestBindExpressionParsing tests individual expression parsing
func TestBindExpressionParsing(t *testing.T) {
	tests := []struct {
		name     string
		expr     string
		wantFunc string
		wantErr  bool
	}{
		{
			name:     "Simple variable",
			expr:     "?x",
			wantFunc: "",
			wantErr:  false,
		},
		{
			name:     "Integer constant",
			expr:     "42",
			wantFunc: "",
			wantErr:  false,
		},
		{
			name:     "Float constant",
			expr:     "3.14",
			wantFunc: "",
			wantErr:  false,
		},
		{
			name:     "Addition",
			expr:     "?x + ?y",
			wantFunc: "+",
			wantErr:  false,
		},
		{
			name:     "Subtraction",
			expr:     "?x - ?y",
			wantFunc: "-",
			wantErr:  false,
		},
		{
			name:     "Multiplication",
			expr:     "?x * ?y",
			wantFunc: "*",
			wantErr:  false,
		},
		{
			name:     "Division",
			expr:     "?x / ?y",
			wantFunc: "/",
			wantErr:  false,
		},
		{
			name:     "Modulo",
			expr:     "?x % ?y",
			wantFunc: "%",
			wantErr:  false,
		},
		{
			name:     "SQRT function",
			expr:     "sqrt(?value)",
			wantFunc: "sqrt",
			wantErr:  false,
		},
		{
			name:     "ABS function",
			expr:     "abs(?value)",
			wantFunc: "abs",
			wantErr:  false,
		},
		{
			name:     "Parenthesized expression",
			expr:     "(?x + ?y)",
			wantFunc: "+",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mathTree, err := parseBindExpression(tt.expr)

			if (err != nil) != tt.wantErr {
				t.Errorf("parseBindExpression() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && mathTree == nil {
				t.Error("Expected non-nil MathTree")
			}

			if !tt.wantErr && tt.wantFunc != "" && mathTree.Fn != tt.wantFunc {
				t.Errorf("Expected Fn=%q, got %q", tt.wantFunc, mathTree.Fn)
			}
		})
	}
}

// TestHavingClause tests HAVING filter support
func TestHavingClause(t *testing.T) {
	tests := []struct {
		name    string
		having  *HavingClause
		wantErr bool
	}{
		{
			name: "Simple HAVING",
			having: &HavingClause{
				Expression: "count(?x) > 5",
			},
			wantErr: false,
		},
		{
			name: "HAVING with multiple conditions",
			having: &HavingClause{
				Expression: "sum(?value) > 100 && avg(?value) < 50",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rootQuery := &dql.GraphQuery{Attr: "query"}
			err := applyHavingClause(rootQuery, tt.having)

			if (err != nil) != tt.wantErr {
				t.Errorf("applyHavingClause() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr && rootQuery.Filter == nil {
				t.Error("Expected filter to be set")
			}
		})
	}
}

// TestTranslateSelectExtended tests full extended SELECT translation
func TestTranslateSelectExtended(t *testing.T) {
	tests := []struct {
		name    string
		query   *SPARQLQueryImpl
		wantErr bool
	}{
		{
			name: "SELECT with aggregates",
			query: &SPARQLQueryImpl{
				Qtype:    "SELECT",
				Prefixes: map[string]string{"ex": "http://example.org/"},
				Projs:    []string{"s", "count"},
				Aggregates: []*Aggregate{
					{Function: "count", Variable: "?item", Alias: "count"},
				},
				Bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "http://example.org/item",
						Object:      "?item",
						ObjectIsVar: true,
					}},
				}},
			},
			wantErr: false,
		},
		{
			name: "SELECT with BIND",
			query: &SPARQLQueryImpl{
				Qtype:    "SELECT",
				Prefixes: map[string]string{"ex": "http://example.org/"},
				Projs:    []string{"s", "sum"},
				Binds: []*BindExpression{
					{Expression: "?x + ?y", Variable: "?sum"},
				},
				Bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "http://example.org/value",
						Object:      "?x",
						ObjectIsVar: true,
					}},
				}},
			},
			wantErr: false,
		},
		{
			name: "SELECT with aggregates and HAVING",
			query: &SPARQLQueryImpl{
				Qtype:    "SELECT",
				Prefixes: map[string]string{"ex": "http://example.org/"},
				Projs:    []string{"s", "total"},
				Aggregates: []*Aggregate{
					{Function: "sum", Variable: "?amount", Alias: "total"},
				},
				Having: &HavingClause{
					Expression: "sum(?amount) > 1000",
				},
				Bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "http://example.org/amount",
						Object:      "?amount",
						ObjectIsVar: true,
					}},
				}},
			},
			wantErr: false,
		},
		{
			name: "SELECT with DISTINCT and ORDER BY",
			query: &SPARQLQueryImpl{
				Qtype:    "SELECT",
				Distinct: true,
				OrderBy:  []string{"?s", "?name"},
				Prefixes: map[string]string{"ex": "http://example.org/"},
				Projs:    []string{"s", "name"},
				Bgps: []*BGP{{
					Triples: []*Triple{{
						Subject:     "?s",
						Predicate:   "http://example.org/name",
						Object:      "?name",
						ObjectIsVar: true,
					}},
				}},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			gqs, prefixes, err := TranslateSelectExtended(ctx, tt.query, TranslateOptions{})

			if (err != nil) != tt.wantErr {
				t.Errorf("TranslateSelectExtended() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				if len(gqs) == 0 {
					t.Error("Expected non-empty GraphQuery slice")
				}
				if prefixes == nil {
					t.Error("Expected prefixes map")
				}
			}
		})
	}
}

// TestExtractBGPFromPattern tests pattern extraction helper
func TestExtractBGPFromPattern(t *testing.T) {
	bgp := &BGP{
		Triples: []*Triple{{
			Subject:     "?s",
			Predicate:   "http://ex.org/p",
			Object:      "?o",
			ObjectIsVar: true,
		}},
	}

	tests := []struct {
		name      string
		pattern   GraphPattern
		wantCount int
	}{
		{
			name:      "Direct BGP",
			pattern:   bgp,
			wantCount: 1,
		},
		{
			name: "Optional with BGP",
			pattern: &OptionalPattern{
				Patterns: []GraphPattern{bgp},
			},
			wantCount: 1,
		},
		{
			name: "Union with BGPs",
			pattern: &UnionPattern{
				Alternatives: [][]GraphPattern{
					{bgp},
					{bgp},
				},
			},
			wantCount: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ExtractBGPFromPattern(tt.pattern)
			if len(result) != tt.wantCount {
				t.Errorf("Expected %d BGPs, got %d", tt.wantCount, len(result))
			}
		})
	}
}

// TestIsOptionalPattern tests pattern type checking
func TestIsOptionalPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  GraphPattern
		expected bool
	}{
		{
			name:     "OptionalPattern",
			pattern:  &OptionalPattern{},
			expected: true,
		},
		{
			name:     "BGP",
			pattern:  &BGP{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsOptionalPattern(tt.pattern)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}

// TestIsUnionPattern tests UNION type checking
func TestIsUnionPattern(t *testing.T) {
	tests := []struct {
		name     string
		pattern  GraphPattern
		expected bool
	}{
		{
			name:     "UnionPattern",
			pattern:  &UnionPattern{},
			expected: true,
		},
		{
			name:     "BGP",
			pattern:  &BGP{},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsUnionPattern(tt.pattern)
			if result != tt.expected {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
		})
	}
}
