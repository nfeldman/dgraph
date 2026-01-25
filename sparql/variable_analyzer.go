package sparql

// VariableScopeAnalyzer detects unused variables in SPARQL patterns and removes unnecessary projections.
// When a variable is defined in a pattern but never used in SELECT, FILTER, GROUP BY, etc.,
// the entire pattern that only defines that variable can be removed (in certain cases).
// This is particularly useful for detecting overly broad patterns.
type VariableScopeAnalyzer struct {
	projectedVars map[string]bool
	usedVars      map[string]bool
	definedVars   map[string]bool
}

// NewVariableScopeAnalyzer creates a new variable scope analyzer.
func NewVariableScopeAnalyzer() *VariableScopeAnalyzer {
	return &VariableScopeAnalyzer{
		projectedVars: make(map[string]bool),
		usedVars:      make(map[string]bool),
		definedVars:   make(map[string]bool),
	}
}

// Analyze performs variable scope analysis on the algebra expression.
func (a *VariableScopeAnalyzer) Analyze(expr AlgebraExpr) *VariableScopeInfo {
	if expr == nil {
		return &VariableScopeInfo{
			ProjectedVars: make(map[string]bool),
			UsedVars:      make(map[string]bool),
			DefinedVars:   make(map[string]bool),
			UnusedVars:    make(map[string]bool),
		}
	}

	// First pass: collect projected variables
	a.collectProjectedVariables(expr)

	// Second pass: collect all used variables
	a.collectUsedVariables(expr)

	// Third pass: collect all defined variables
	a.collectDefinedVariables(expr)

	// Calculate unused variables
	unusedVars := make(map[string]bool)
	for v := range a.definedVars {
		if !a.usedVars[v] && !a.projectedVars[v] {
			unusedVars[v] = true
		}
	}

	return &VariableScopeInfo{
		ProjectedVars: a.projectedVars,
		UsedVars:      a.usedVars,
		DefinedVars:   a.definedVars,
		UnusedVars:    unusedVars,
	}
}

// VariableScopeInfo contains the results of variable scope analysis.
type VariableScopeInfo struct {
	ProjectedVars map[string]bool
	UsedVars      map[string]bool
	DefinedVars   map[string]bool
	UnusedVars    map[string]bool
}

// IsVariableUnused returns true if a variable is defined but never used.
func (info *VariableScopeInfo) IsVariableUnused(varName string) bool {
	return info.UnusedVars[varName]
}

// HasUnusedVariables returns true if there are any unused variables.
func (info *VariableScopeInfo) HasUnusedVariables() bool {
	return len(info.UnusedVars) > 0
}

// GetUnusedVariables returns a list of unused variables.
func (info *VariableScopeInfo) GetUnusedVariables() []string {
	var unused []string
	for v := range info.UnusedVars {
		unused = append(unused, v)
	}
	return unused
}

// collectProjectedVariables finds all variables that appear in SELECT/PROJECT.
func (a *VariableScopeAnalyzer) collectProjectedVariables(expr AlgebraExpr) {
	if expr == nil {
		return
	}

	switch node := expr.(type) {
	case *AlgebraProject:
		for _, v := range node.Vars {
			if isVariable(v) {
				a.projectedVars[v] = true
			}
		}
	}
}

// collectUsedVariables finds all variables that are actually referenced.
func (a *VariableScopeAnalyzer) collectUsedVariables(expr AlgebraExpr) {
	if expr == nil {
		return
	}

	switch node := expr.(type) {
	case *AlgebraProject:
		// Variables in SELECT are used
		for _, v := range node.Vars {
			if isVariable(v) {
				a.usedVars[v] = true
			}
		}
		a.collectUsedVariables(node.Input)

	case *AlgebraFilter:
		// Variables in FILTER are used
		extractVariablesFromFilter(node.Expr, a.usedVars)
		a.collectUsedVariables(node.Input)

	case *AlgebraAgg:
		// Variables in GROUP BY are used
		for _, g := range node.Group {
			if isVariable(g) {
				a.usedVars[g] = true
			}
		}
		a.collectUsedVariables(node.Input)

	case *AlgebraOrderBy:
		// Variables in ORDER BY are used
		for _, expr := range node.Exprs {
			if isVariable(expr) {
				a.usedVars[expr] = true
			}
		}
		a.collectUsedVariables(node.Input)

	case *AlgebraBind:
		// Variables in BIND expressions are used (for defining new vars)
		extractVariablesFromFilter(node.Expr, a.usedVars)
		a.collectUsedVariables(node.Input)

	case *AlgebraLeftJoin, *AlgebraDistinct, *AlgebraLimit:
		a.collectUsedVariables(getInput(node))

	case *AlgebraJoin:
		a.collectUsedVariables(node.Left)
		a.collectUsedVariables(node.Right)

	case *AlgebraUnion:
		for _, alt := range node.Alternatives {
			a.collectUsedVariables(alt)
		}
	}
}

// collectDefinedVariables finds all variables defined in the patterns.
func (a *VariableScopeAnalyzer) collectDefinedVariables(expr AlgebraExpr) {
	if expr == nil {
		return
	}

	switch node := expr.(type) {
	case *AlgebraBGP:
		// Variables in BGP patterns are defined
		for _, triple := range node.Triples {
			if isVariable(triple.Subject) {
				a.definedVars[triple.Subject] = true
			}
			if triple.ObjectIsVar && isVariable(triple.Object) {
				a.definedVars[triple.Object] = true
			}
		}

	case *AlgebraValues:
		// Variables in VALUES are defined
		for _, v := range node.Vars {
			if isVariable(v) {
				a.definedVars[v] = true
			}
		}

	case *AlgebraBind:
		// BIND defines a variable
		a.definedVars[node.Var] = true
		a.collectDefinedVariables(node.Input)

	case *AlgebraProject:
		a.collectDefinedVariables(node.Input)

	case *AlgebraFilter:
		a.collectDefinedVariables(node.Input)

	case *AlgebraOrderBy:
		a.collectDefinedVariables(node.Input)

	case *AlgebraDistinct:
		a.collectDefinedVariables(node.Input)

	case *AlgebraLimit:
		a.collectDefinedVariables(node.Input)

	case *AlgebraAgg:
		a.collectDefinedVariables(node.Input)

	case *AlgebraJoin:
		a.collectDefinedVariables(node.Left)
		a.collectDefinedVariables(node.Right)

	case *AlgebraUnion:
		for _, alt := range node.Alternatives {
			a.collectDefinedVariables(alt)
		}

	case *AlgebraLeftJoin:
		a.collectDefinedVariables(node.Input)
	}
}
