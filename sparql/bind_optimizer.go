package sparql

// BindExpressionOptimizer removes unused BIND expressions and inlines simple bindings.
// BIND creates variable bindings in SPARQL, but if the bound variable is not used
// in the SELECT clause or elsewhere, it's dead code and should be removed.
type BindExpressionOptimizer struct {
	// Tracks which variables are actually used in projections/filters
	usedVars map[string]bool
}

// NewBindExpressionOptimizer creates a new BIND expression optimizer.
func NewBindExpressionOptimizer() *BindExpressionOptimizer {
	return &BindExpressionOptimizer{
		usedVars: make(map[string]bool),
	}
}

// Optimize removes unused BIND expressions from the algebra tree.
func (o *BindExpressionOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	// First, collect which variables are actually used
	o.collectUsedVariables(expr)

	// Then optimize the tree
	return o.optimizeExpr(expr)
}

// collectUsedVariables walks the expression tree and collects all variables that are used.
func (o *BindExpressionOptimizer) collectUsedVariables(expr AlgebraExpr) {
	if expr == nil {
		return
	}

	switch node := expr.(type) {
	case *AlgebraProject:
		// Variables in SELECT are used
		for _, v := range node.Vars {
			if isVariable(v) {
				o.usedVars[v] = true
			}
		}
		o.collectUsedVariables(node.Input)

	case *AlgebraFilter:
		// Variables in FILTER are used
		extractVariablesFromFilter(node.Expr, o.usedVars)
		o.collectUsedVariables(node.Input)

	case *AlgebraAgg:
		// Variables in GROUP BY are used
		for _, g := range node.Group {
			if isVariable(g) {
				o.usedVars[g] = true
			}
		}
		o.collectUsedVariables(node.Input)

	case *AlgebraOrderBy:
		// Variables in ORDER BY are used
		for _, expr := range node.Exprs {
			if isVariable(expr) {
				o.usedVars[expr] = true
			}
		}
		o.collectUsedVariables(node.Input)

	case *AlgebraDistinct, *AlgebraLimit, *AlgebraLeftJoin:
		o.collectUsedVariables(getInput(node))

	case *AlgebraJoin:
		jnode := node
		o.collectUsedVariables(jnode.Left)
		o.collectUsedVariables(jnode.Right)

	case *AlgebraUnion:
		unode := node
		for _, alt := range unode.Alternatives {
			o.collectUsedVariables(alt)
		}

	case *AlgebraBGP:
		// Variables in BGP patterns are used
		for _, triple := range node.Triples {
			if isVariable(triple.Subject) {
				o.usedVars[triple.Subject] = true
			}
			if triple.ObjectIsVar && isVariable(triple.Object) {
				o.usedVars[triple.Object] = true
			}
		}
	}
}

// optimizeExpr recursively optimizes the expression tree, removing unused BIND expressions.
func (o *BindExpressionOptimizer) optimizeExpr(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraBind:
		// Check if this BIND's variable is used
		if !o.usedVars[node.Var] {
			// Variable is unused - skip this BIND and recurse into input
			return o.optimizeExpr(node.Input)
		}

		// Variable is used - keep the BIND but optimize its input
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}

		return &AlgebraBind{
			Var:   node.Var,
			Expr:  node.Expr,
			Input: optimizedInput,
		}, nil

	case *AlgebraFilter:
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraFilter{
			Expr:  node.Expr,
			Input: optimizedInput,
		}, nil

	case *AlgebraProject:
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraProject{
			Vars:  node.Vars,
			Input: optimizedInput,
		}, nil

	case *AlgebraAgg:
		optimizedInput, err := o.optimizeExpr(node.Input)
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
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraDistinct{
			Input: optimizedInput,
		}, nil

	case *AlgebraOrderBy:
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraOrderBy{
			Exprs:     node.Exprs,
			Ascending: node.Ascending,
			Input:     optimizedInput,
		}, nil

	case *AlgebraLimit:
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraLimit{
			Count:  node.Count,
			Offset: node.Offset,
			Input:  optimizedInput,
		}, nil

	case *AlgebraLeftJoin:
		optimizedInput, err := o.optimizeExpr(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraLeftJoin{
			Input:    optimizedInput,
			Patterns: node.Patterns,
			Filter:   node.Filter,
		}, nil

	case *AlgebraJoin:
		optimizedLeft, err := o.optimizeExpr(node.Left)
		if err != nil {
			return nil, err
		}
		optimizedRight, err := o.optimizeExpr(node.Right)
		if err != nil {
			return nil, err
		}
		return &AlgebraJoin{
			Left:  optimizedLeft,
			Right: optimizedRight,
		}, nil

	case *AlgebraUnion:
		var optimizedAlts []AlgebraExpr
		for _, alt := range node.Alternatives {
			optimized, err := o.optimizeExpr(alt)
			if err != nil {
				return nil, err
			}
			optimizedAlts = append(optimizedAlts, optimized)
		}
		return &AlgebraUnion{
			Alternatives: optimizedAlts,
		}, nil

	default:
		return expr, nil
	}
}

// Helper functions

func isVariable(s string) bool {
	if len(s) == 0 {
		return false
	}
	return s[0] == '?'
}

func extractVariablesFromFilter(expr string, vars map[string]bool) {
	// Simple variable extraction from filter string
	// Look for ?variable patterns
	for i := 0; i < len(expr); i++ {
		if expr[i] == '?' && i+1 < len(expr) {
			// Found start of variable
			j := i + 1
			for j < len(expr) && (isAlphaNumeric(expr[j]) || expr[j] == '_') {
				j++
			}
			if j > i+1 {
				vars[expr[i:j]] = true
			}
			i = j - 1
		}
	}
}

func isAlphaNumeric(b byte) bool {
	return (b >= 'a' && b <= 'z') || (b >= 'A' && b <= 'Z') || (b >= '0' && b <= '9')
}

func getInput(expr AlgebraExpr) AlgebraExpr {
	switch node := expr.(type) {
	case *AlgebraDistinct:
		return node.Input
	case *AlgebraLimit:
		return node.Input
	case *AlgebraLeftJoin:
		return node.Input
	default:
		return nil
	}
}
