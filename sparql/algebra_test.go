package sparql

import (
	"strings"
	"testing"
)

// ============================================================================
// Step 1: Algebra Type System Tests
// ============================================================================

func TestAlgebraBGP(t *testing.T) {
	tests := []struct {
		name     string
		triples  []*Triple
		wantVars []string
	}{
		{
			name: "single triple with variables",
			triples: []*Triple{
				{Subject: "?person", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
			},
			wantVars: []string{"?person"},
		},
		{
			name: "multiple triples",
			triples: []*Triple{
				{Subject: "?person", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
				{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
			},
			wantVars: []string{"?person", "?name"},
		},
		{
			name: "no variables",
			triples: []*Triple{
				{Subject: "ex:Alice", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
			},
			wantVars: []string{},
		},
		{
			name: "predicate variable",
			triples: []*Triple{
				{Subject: "?s", Predicate: "?p", Object: "?o", ObjectIsVar: true},
			},
			wantVars: []string{"?s", "?p", "?o"},
		},
		{
			name:     "empty BGP",
			triples:  []*Triple{},
			wantVars: []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bgp := &AlgebraBGP{Triples: tt.triples}

			// Test String()
			str := bgp.String()
			if !strings.Contains(str, "BGP") {
				t.Errorf("String() should contain 'BGP', got: %s", str)
			}

			// Test Variables()
			vars := bgp.Variables()
			if !variablesMatch(vars, tt.wantVars) {
				t.Errorf("Variables() = %v, want %v", vars, tt.wantVars)
			}

			// Test DefinedVars()
			defined := bgp.DefinedVars()
			if !variablesMatch(defined, tt.wantVars) {
				t.Errorf("DefinedVars() = %v, want %v", defined, tt.wantVars)
			}

			// Test Accept()
			visitor := &testVisitor{}
			result := bgp.Accept(visitor)
			if result != "bgp" {
				t.Errorf("Accept() should return 'bgp', got %v", result)
			}
		})
	}
}

func TestAlgebraJoin(t *testing.T) {
	left := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	right := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?y", Predicate: "q", Object: "?x", ObjectIsVar: true},
	}}

	join := &AlgebraJoin{Left: left, Right: right}

	// Test String()
	str := join.String()
	if !strings.Contains(str, "Join") {
		t.Errorf("String() should contain 'Join', got: %s", str)
	}

	// Test Variables()
	vars := join.Variables()
	if !containsVariable(vars, "?x") || !containsVariable(vars, "?y") {
		t.Errorf("Variables() = %v, missing expected variables", vars)
	}

	// Test DefinedVars()
	defined := join.DefinedVars()
	if !containsVariable(defined, "?x") || !containsVariable(defined, "?y") {
		t.Errorf("DefinedVars() = %v, missing expected variables", defined)
	}
}

func TestAlgebraFilter(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	filter := &AlgebraFilter{
		Expr:  "?x > 5",
		Input: input,
	}

	// Test String()
	str := filter.String()
	if !strings.Contains(str, "Filter") {
		t.Errorf("String() should contain 'Filter', got: %s", str)
	}
	if !strings.Contains(str, "?x > 5") {
		t.Errorf("String() should contain filter expression, got: %s", str)
	}

	// Test Variables() - includes variables from input
	vars := filter.Variables()
	if !containsVariable(vars, "?x") {
		t.Errorf("Variables() = %v, missing ?x", vars)
	}

	// Test DefinedVars() - same as input's defined vars
	defined := filter.DefinedVars()
	if !containsVariable(defined, "?x") {
		t.Errorf("DefinedVars() = %v, missing ?x", defined)
	}
}

func TestAlgebraLeftJoin(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	patterns := []*Triple{
		{Subject: "?x", Predicate: "optional", Object: "?y", ObjectIsVar: true},
	}
	lj := &AlgebraLeftJoin{
		Input:    input,
		Patterns: patterns,
		Filter:   "?y != 0",
	}

	// Test String()
	str := lj.String()
	if !strings.Contains(str, "LeftJoin") {
		t.Errorf("String() should contain 'LeftJoin', got: %s", str)
	}

	// Test Variables()
	vars := lj.Variables()
	if !containsVariable(vars, "?x") || !containsVariable(vars, "?y") {
		t.Errorf("Variables() should include both ?x and ?y, got %v", vars)
	}

	// Test DefinedVars()
	defined := lj.DefinedVars()
	if !containsVariable(defined, "?x") || !containsVariable(defined, "?y") {
		t.Errorf("DefinedVars() should include both ?x and ?y, got %v", defined)
	}
}

func TestAlgebraUnion(t *testing.T) {
	alt1 := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	alt2 := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "q", Object: "o", ObjectIsVar: false},
	}}
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{alt1, alt2},
	}

	// Test String()
	str := union.String()
	if !strings.Contains(str, "Union") {
		t.Errorf("String() should contain 'Union', got: %s", str)
	}

	// Test Variables()
	vars := union.Variables()
	if !containsVariable(vars, "?x") {
		t.Errorf("Variables() = %v, missing ?x", vars)
	}

	// Test DefinedVars()
	defined := union.DefinedVars()
	if !containsVariable(defined, "?x") {
		t.Errorf("DefinedVars() = %v, missing ?x", defined)
	}
}

func TestAlgebraProject(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
	}}
	proj := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: input,
	}

	// Test String()
	str := proj.String()
	if !strings.Contains(str, "Project") {
		t.Errorf("String() should contain 'Project', got: %s", str)
	}
	if !strings.Contains(str, "?x") {
		t.Errorf("String() should contain projected variable, got: %s", str)
	}

	// Test Variables()
	vars := proj.Variables()
	if !containsVariable(vars, "?x") {
		t.Errorf("Variables() = %v, should only include ?x", vars)
	}

	// Test DefinedVars()
	defined := proj.DefinedVars()
	if !containsVariable(defined, "?x") {
		t.Errorf("DefinedVars() = %v, should only include ?x", defined)
	}
}

func TestAlgebraAgg(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?person", Predicate: "age", Object: "?age", ObjectIsVar: true},
	}}
	agg := &AlgebraAgg{
		Op:    "COUNT",
		Var:   "?age",
		Group: []string{"?person"},
		Input: input,
	}

	// Test String()
	str := agg.String()
	if !strings.Contains(str, "Agg") {
		t.Errorf("String() should contain 'Agg', got: %s", str)
	}

	// Test Variables() and DefinedVars()
	vars := agg.Variables()
	if !containsVariable(vars, "?person") {
		t.Errorf("Variables() = %v, missing group variable", vars)
	}

	defined := agg.DefinedVars()
	if !containsVariable(defined, "?person") && !containsVariable(defined, "?age") {
		t.Errorf("DefinedVars() = %v, should include group/aggregate vars", defined)
	}
}

func TestAlgebraBind(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	bind := &AlgebraBind{
		Var:   "?y",
		Expr:  "?x + 1",
		Input: input,
	}

	// Test String()
	str := bind.String()
	if !strings.Contains(str, "Bind") {
		t.Errorf("String() should contain 'Bind', got: %s", str)
	}

	// Test Variables()
	vars := bind.Variables()
	if !containsVariable(vars, "?x") || !containsVariable(vars, "?y") {
		t.Errorf("Variables() = %v, should include both ?x and ?y", vars)
	}

	// Test DefinedVars() - Bind adds its variable to what's already defined
	defined := bind.DefinedVars()
	if !containsVariable(defined, "?y") {
		t.Errorf("DefinedVars() = %v, should include ?y", defined)
	}
}

func TestAlgebraDistinct(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	distinct := &AlgebraDistinct{Input: input}

	// Test String()
	str := distinct.String()
	if !strings.Contains(str, "Distinct") {
		t.Errorf("String() should contain 'Distinct', got: %s", str)
	}

	// Test Variables()
	vars := distinct.Variables()
	if !containsVariable(vars, "?x") {
		t.Errorf("Variables() = %v, missing ?x", vars)
	}

	// Test DefinedVars()
	defined := distinct.DefinedVars()
	if !containsVariable(defined, "?x") {
		t.Errorf("DefinedVars() = %v, missing ?x", defined)
	}
}

func TestAlgebraOrderBy(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
	}}
	orderBy := &AlgebraOrderBy{
		Exprs:     []string{"?x", "?y"},
		Ascending: []bool{true, false},
		Input:     input,
	}

	// Test String()
	str := orderBy.String()
	if !strings.Contains(str, "OrderBy") {
		t.Errorf("String() should contain 'OrderBy', got: %s", str)
	}

	// Test Variables()
	vars := orderBy.Variables()
	if !containsVariable(vars, "?x") || !containsVariable(vars, "?y") {
		t.Errorf("Variables() = %v, missing expected variables", vars)
	}

	// Test DefinedVars()
	defined := orderBy.DefinedVars()
	if !containsVariable(defined, "?x") || !containsVariable(defined, "?y") {
		t.Errorf("DefinedVars() = %v, missing expected variables", defined)
	}
}

func TestAlgebraLimit(t *testing.T) {
	input := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	limit := &AlgebraLimit{
		Count:  10,
		Offset: 5,
		Input:  input,
	}

	// Test String()
	str := limit.String()
	if !strings.Contains(str, "Limit") {
		t.Errorf("String() should contain 'Limit', got: %s", str)
	}
	if !strings.Contains(str, "10") || !strings.Contains(str, "5") {
		t.Errorf("String() should contain count and offset, got: %s", str)
	}

	// Test Variables()
	vars := limit.Variables()
	if !containsVariable(vars, "?x") {
		t.Errorf("Variables() = %v, missing ?x", vars)
	}

	// Test DefinedVars()
	defined := limit.DefinedVars()
	if !containsVariable(defined, "?x") {
		t.Errorf("DefinedVars() = %v, missing ?x", defined)
	}
}

// ============================================================================
// Step 2: Visitor Infrastructure Tests
// ============================================================================

func TestAlgebraPrinter(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	printer := NewAlgebraPrinter()
	result := printer.Print(bgp)

	if !strings.Contains(result, "BGP") {
		t.Errorf("AlgebraPrinter.Print() should contain 'BGP', got: %s", result)
	}
}

func TestVariableCollector(t *testing.T) {
	tests := []struct {
		name     string
		expr     AlgebraExpr
		wantVars []string
	}{
		{
			name: "simple BGP",
			expr: &AlgebraBGP{Triples: []*Triple{
				{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
			}},
			wantVars: []string{"?x", "?y"},
		},
		{
			name: "join",
			expr: &AlgebraJoin{
				Left: &AlgebraBGP{Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				}},
				Right: &AlgebraBGP{Triples: []*Triple{
					{Subject: "?y", Predicate: "q", Object: "o", ObjectIsVar: false},
				}},
			},
			wantVars: []string{"?x", "?y"},
		},
		{
			name: "project",
			expr: &AlgebraProject{
				Vars: []string{"?x"},
				Input: &AlgebraBGP{Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
				}},
			},
			wantVars: []string{"?x", "?y"},
		},
		{
			name: "bind",
			expr: &AlgebraBind{
				Var:  "?z",
				Expr: "?x + ?y",
				Input: &AlgebraBGP{Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
				}},
			},
			wantVars: []string{"?x", "?y", "?z"},
		},
		{
			name: "filter",
			expr: &AlgebraFilter{
				Expr: "?x > 5",
				Input: &AlgebraBGP{Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				}},
			},
			wantVars: []string{"?x"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			collector := NewVariableCollector()
			vars := collector.Collect(tt.expr)
			if !variablesMatch(vars, tt.wantVars) {
				t.Errorf("Collect() = %v, want %v", vars, tt.wantVars)
			}
		})
	}
}

func TestExpressionDepthCalculator(t *testing.T) {
	tests := []struct {
		name      string
		expr      AlgebraExpr
		wantDepth int
	}{
		{
			name:      "single BGP",
			expr:      &AlgebraBGP{Triples: []*Triple{}},
			wantDepth: 1,
		},
		{
			name: "BGP wrapped in filter",
			expr: &AlgebraFilter{
				Expr:  "?x > 5",
				Input: &AlgebraBGP{Triples: []*Triple{}},
			},
			wantDepth: 2,
		},
		{
			name: "join of two BGPs",
			expr: &AlgebraJoin{
				Left:  &AlgebraBGP{Triples: []*Triple{}},
				Right: &AlgebraBGP{Triples: []*Triple{}},
			},
			wantDepth: 2,
		},
		{
			name: "nested joins",
			expr: &AlgebraJoin{
				Left: &AlgebraJoin{
					Left:  &AlgebraBGP{Triples: []*Triple{}},
					Right: &AlgebraBGP{Triples: []*Triple{}},
				},
				Right: &AlgebraBGP{Triples: []*Triple{}},
			},
			wantDepth: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			calc := &ExpressionDepthCalculator{}
			depth := calc.Depth(tt.expr)
			if depth != tt.wantDepth {
				t.Errorf("Depth() = %d, want %d", depth, tt.wantDepth)
			}
		})
	}
}

func TestExpressionTreeFormatter(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}
	formatter := &ExpressionTreeFormatter{}
	result := formatter.Format(bgp)

	if !strings.Contains(result, "BGP") {
		t.Errorf("Format() should contain 'BGP', got: %s", result)
	}
	if len(result) == 0 {
		t.Errorf("Format() should return non-empty string")
	}
}

// ============================================================================
// Step 3: AST to Algebra Converter Tests
// ============================================================================

func TestASTToAlgebraSimpleBGP(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x"},
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Result should be Project(BGP)
	proj, ok := expr.(*AlgebraProject)
	if !ok {
		t.Errorf("Expected AlgebraProject, got %T", expr)
		return
	}

	if !containsVariable(proj.Vars, "?x") {
		t.Errorf("Project vars = %v, want ?x", proj.Vars)
	}

	_, ok = proj.Input.(*AlgebraBGP)
	if !ok {
		t.Errorf("Expected BGP as input, got %T", proj.Input)
	}
}

func TestASTToAlgebraMultipleTriples(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x", "?y"},
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
					{Subject: "?y", Predicate: "q", Object: "o", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Result should be Project(BGP(...))
	proj, ok := expr.(*AlgebraProject)
	if !ok {
		t.Errorf("Expected AlgebraProject, got %T", expr)
		return
	}

	bgp, ok := proj.Input.(*AlgebraBGP)
	if !ok {
		t.Errorf("Expected BGP as input, got %T", proj.Input)
		return
	}

	if len(bgp.Triples) != 2 {
		t.Errorf("BGP should have 2 triples, got %d", len(bgp.Triples))
	}
}

func TestASTToAlgebraWithFilter(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
			&FilterPattern{Expression: "?x > 5"},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Navigate to filter
	proj := expr.(*AlgebraProject)
	filter := proj.Input.(*AlgebraFilter)
	if filter.Expr != "?x > 5" {
		t.Errorf("Filter expr = %s, want '?x > 5'", filter.Expr)
	}
}

func TestASTToAlgebraWithBind(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x", "?y"},
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
		},
		Binds: []*BindExpression{
			{Expression: "?x + 1", Variable: "?y"},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Result should be Project(Bind(...))
	proj := expr.(*AlgebraProject)
	bind, ok := proj.Input.(*AlgebraBind)
	if !ok {
		t.Errorf("Expected Bind, got %T", proj.Input)
		return
	}

	if bind.Var != "?y" || bind.Expr != "?x + 1" {
		t.Errorf("Bind = (%s, %s), want (?y, ?x + 1)", bind.Var, bind.Expr)
	}
}

func TestASTToAlgebraWithDistinct(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype:    "SELECT",
		Projs:    []string{"?x"},
		Distinct: true,
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Should have Distinct in the tree
	proj := expr.(*AlgebraProject)
	distinct, ok := proj.Input.(*AlgebraDistinct)
	if !ok {
		t.Errorf("Expected Distinct, got %T", proj.Input)
		return
	}

	_, ok = distinct.Input.(*AlgebraBGP)
	if !ok {
		t.Errorf("Expected BGP as input to Distinct, got %T", distinct.Input)
	}
}

func TestASTToAlgebraWithOrderBy(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype:   "SELECT",
		Projs:   []string{"?x"},
		OrderBy: []string{"?x"},
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Should have OrderBy before Project
	proj := expr.(*AlgebraProject)
	orderBy, ok := proj.Input.(*AlgebraOrderBy)
	if !ok {
		t.Errorf("Expected OrderBy, got %T", proj.Input)
		return
	}

	if len(orderBy.Exprs) != 1 || orderBy.Exprs[0] != "?x" {
		t.Errorf("OrderBy exprs = %v, want [?x]", orderBy.Exprs)
	}
}

func TestASTToAlgebraWithLimit(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype:  "SELECT",
		Projs:  []string{"?x"},
		Limit:  10,
		Offset: 5,
		Bgps: []*BGP{
			{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	proj := expr.(*AlgebraProject)
	limit, ok := proj.Input.(*AlgebraLimit)
	if !ok {
		t.Errorf("Expected Limit, got %T", proj.Input)
		return
	}

	if limit.Count != 10 || limit.Offset != 5 {
		t.Errorf("Limit = (%d, %d), want (10, 5)", limit.Count, limit.Offset)
	}
}

func TestASTToAlgebraNoPatterns(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x"},
	}

	_, err := ASTToAlgebra(query)
	if err == nil {
		t.Errorf("ASTToAlgebra() should error on no patterns")
	}
}

func TestASTToAlgebraWithPatterns(t *testing.T) {
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?x"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	if expr == nil {
		t.Errorf("ASTToAlgebra() returned nil")
	}
}

// ============================================================================
// Step 4: Algebra Validator Tests
// ============================================================================

func TestValidatorSimpleBGP(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
	}}

	validator := NewAlgebraValidator()
	err := validator.Validate(bgp)
	if err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}
}

func TestValidatorUndefinedVariable(t *testing.T) {
	filter := &AlgebraFilter{
		Expr: "?undefined > 5",
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(filter)
	if err == nil {
		t.Errorf("Validator.Validate() should error on undefined variable")
	}
	if !strings.Contains(err.Error(), "?undefined") {
		t.Errorf("Error should mention undefined variable, got: %v", err)
	}
}

func TestValidatorDefinedInInput(t *testing.T) {
	filter := &AlgebraFilter{
		Expr: "?x > 5",
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(filter)
	if err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}
}

func TestValidatorCircularBind(t *testing.T) {
	bind := &AlgebraBind{
		Var:  "?x",
		Expr: "?x + 1",
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?y", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(bind)
	if err == nil {
		t.Errorf("Validator.Validate() should error on circular BIND")
	}
	if !strings.Contains(err.Error(), "Circular") {
		t.Errorf("Error should mention circular dependency, got: %v", err)
	}
}

func TestValidatorValidBind(t *testing.T) {
	bind := &AlgebraBind{
		Var:  "?y",
		Expr: "?x + 1",
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(bind)
	if err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}
}

func TestValidatorProjectUndefined(t *testing.T) {
	proj := &AlgebraProject{
		Vars: []string{"?undefined"},
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(proj)
	if err == nil {
		t.Errorf("Validator.Validate() should error on undefined projected variable")
	}
}

func TestValidatorOrderByUndefined(t *testing.T) {
	orderBy := &AlgebraOrderBy{
		Exprs:     []string{"?undefined"},
		Ascending: []bool{true},
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(orderBy)
	if err == nil {
		t.Errorf("Validator.Validate() should error on undefined ORDER BY variable")
	}
}

func TestValidatorJoin(t *testing.T) {
	join := &AlgebraJoin{
		Left: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
		Right: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?y", Predicate: "q", Object: "o", ObjectIsVar: false},
		}},
	}

	proj := &AlgebraProject{
		Vars:  []string{"?x", "?y"},
		Input: join,
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(proj)
	if err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}
}

func TestValidatorLeftJoinUndefinedOptional(t *testing.T) {
	lj := &AlgebraLeftJoin{
		Input: &AlgebraBGP{Triples: []*Triple{
			{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
		}},
		Patterns: []*Triple{
			{Subject: "?x", Predicate: "q", Object: "?y", ObjectIsVar: true},
		},
		Filter: "?undefined > 5",
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(lj)
	if err == nil {
		t.Errorf("Validator.Validate() should error on undefined variable in LEFT JOIN filter")
	}
}

func TestValidatorUnion(t *testing.T) {
	union := &AlgebraUnion{
		Alternatives: []AlgebraExpr{
			&AlgebraBGP{Triples: []*Triple{
				{Subject: "?x", Predicate: "p", Object: "o", ObjectIsVar: false},
			}},
			&AlgebraBGP{Triples: []*Triple{
				{Subject: "?x", Predicate: "q", Object: "o", ObjectIsVar: false},
			}},
		},
	}

	proj := &AlgebraProject{
		Vars:  []string{"?x"},
		Input: union,
	}

	validator := NewAlgebraValidator()
	err := validator.Validate(proj)
	if err != nil {
		t.Errorf("Validator.Validate() error = %v, want nil", err)
	}
}

func TestValidationStats(t *testing.T) {
	bgp := &AlgebraBGP{Triples: []*Triple{
		{Subject: "?x", Predicate: "p", Object: "?y", ObjectIsVar: true},
	}}

	stats := GetValidationStats(bgp)
	if stats.TotalVariables != 2 {
		t.Errorf("stats.TotalVariables = %d, want 2", stats.TotalVariables)
	}
	if stats.OperatorCount != 1 {
		t.Errorf("stats.OperatorCount = %d, want 1", stats.OperatorCount)
	}
	if stats.MaxDepth != 1 {
		t.Errorf("stats.MaxDepth = %d, want 1", stats.MaxDepth)
	}
}

// ============================================================================
// Integration Tests
// ============================================================================

func TestFullConversionAndValidation(t *testing.T) {
	// Build complete query: SELECT ?name WHERE { ?x foaf:name ?name . FILTER (?name != "Alice") }
	query := &SPARQLQueryImpl{
		Qtype: "SELECT",
		Projs: []string{"?name"},
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?x", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
				},
			},
			&FilterPattern{Expression: "?name != \"Alice\""},
		},
	}

	// Convert AST to algebra
	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	// Validate algebra
	validator := NewAlgebraValidator()
	err = validator.Validate(expr)
	if err != nil {
		t.Fatalf("Validator.Validate() error = %v", err)
	}

	// Collect variables
	collector := NewVariableCollector()
	vars := collector.Collect(expr)
	if !containsVariable(vars, "?name") {
		t.Errorf("Collected variables = %v, missing ?name", vars)
	}

	// Format output
	formatted := AlgebraToString(expr)
	if len(formatted) == 0 {
		t.Errorf("AlgebraToString() returned empty string")
	}
}

func TestComplexQueryConversion(t *testing.T) {
	// SELECT ?person ?name WHERE {
	//   { ?person rdf:type Person . ?person foaf:name ?name }
	//   OPTIONAL { ?person foaf:age ?age . FILTER (?age > 18) }
	// } ORDER BY ?name LIMIT 10
	query := &SPARQLQueryImpl{
		Qtype:   "SELECT",
		Projs:   []string{"?person", "?name"},
		OrderBy: []string{"?name"},
		Limit:   10,
		Patterns: []GraphPattern{
			&BGP{
				Triples: []*Triple{
					{Subject: "?person", Predicate: "rdf:type", Object: "Person", ObjectIsVar: false},
					{Subject: "?person", Predicate: "foaf:name", Object: "?name", ObjectIsVar: true},
				},
			},
		},
	}

	expr, err := ASTToAlgebra(query)
	if err != nil {
		t.Fatalf("ASTToAlgebra() error = %v", err)
	}

	validator := NewAlgebraValidator()
	err = validator.Validate(expr)
	if err != nil {
		t.Fatalf("Validator.Validate() error = %v", err)
	}

	// Verify structure
	proj, ok := expr.(*AlgebraProject)
	if !ok {
		t.Errorf("Top level should be Project, got %T", expr)
	}
	if len(proj.Vars) != 2 {
		t.Errorf("Project vars = %d, want 2", len(proj.Vars))
	}
}

// ============================================================================
// Helper Functions
// ============================================================================

func variablesMatch(got, want []string) bool {
	if len(got) != len(want) {
		return false
	}
	gotMap := make(map[string]bool)
	for _, v := range got {
		gotMap[v] = true
	}
	for _, v := range want {
		if !gotMap[v] {
			return false
		}
	}
	return true
}

func containsVariable(vars []string, target string) bool {
	for _, v := range vars {
		if v == target {
			return true
		}
	}
	return false
}

// testVisitor is a simple visitor for testing
type testVisitor struct{}

func (tv *testVisitor) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	return "bgp"
}

func (tv *testVisitor) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	return "join"
}

func (tv *testVisitor) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	return "filter"
}

func (tv *testVisitor) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	return "leftjoin"
}

func (tv *testVisitor) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	return "union"
}

func (tv *testVisitor) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	return "project"
}

func (tv *testVisitor) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	return "agg"
}

func (tv *testVisitor) VisitAlgebraBind(b *AlgebraBind) interface{} {
	return "bind"
}

func (tv *testVisitor) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	return "distinct"
}

func (tv *testVisitor) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	return "orderby"
}

func (tv *testVisitor) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	return "limit"
}
