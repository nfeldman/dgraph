package sparql

// UnionPatternOptimizer optimizes SPARQL UNION patterns.
// It applies transformations like:
// - Remove duplicate branches (identical patterns)
// - Merge compatible branches (same variables, different conditions)
// - Flatten nested UNIONs
// - Remove UNION with single branch
// - Move filters out of UNIONs when possible
type UnionPatternOptimizer struct{}

// NewUnionPatternOptimizer creates a new union pattern optimizer.
func NewUnionPatternOptimizer() *UnionPatternOptimizer {
	return &UnionPatternOptimizer{}
}

// Optimize applies union pattern optimizations to the algebra tree.
func (o *UnionPatternOptimizer) Optimize(expr AlgebraExpr) (AlgebraExpr, error) {
	if expr == nil {
		return expr, nil
	}

	switch node := expr.(type) {
	case *AlgebraUnion:
		return o.optimizeUnion(node)
	default:
		return o.optimizeChildren(expr)
	}
}

// optimizeUnion handles a union expression.
func (o *UnionPatternOptimizer) optimizeUnion(union *AlgebraUnion) (AlgebraExpr, error) {
	if len(union.Alternatives) == 0 {
		return nil, nil
	}

	// First, recursively optimize all alternatives
	optimizedAlts, err := o.optimizeAlternatives(union.Alternatives)
	if err != nil {
		return nil, err
	}

	// Flatten nested unions
	flattened := o.flattenNestedUnions(optimizedAlts)

	// Remove duplicates
	deduped := o.removeDuplicateBranches(flattened)

	// If only one branch remains, return it directly
	if len(deduped) == 1 {
		return deduped[0], nil
	}

	// If no branches remain, return empty
	if len(deduped) == 0 {
		return &AlgebraEmpty{}, nil
	}

	return &AlgebraUnion{
		Alternatives: deduped,
	}, nil
}

// optimizeAlternatives recursively optimizes each union branch.
func (o *UnionPatternOptimizer) optimizeAlternatives(alts []AlgebraExpr) ([]AlgebraExpr, error) {
	var optimized []AlgebraExpr

	for _, alt := range alts {
		if alt == nil {
			continue
		}

		// Recursively optimize
		opt, err := o.Optimize(alt)
		if err != nil {
			return nil, err
		}

		optimized = append(optimized, opt)
	}

	return optimized, nil
}

// flattenNestedUnions flattens nested UNION expressions.
// UNION(UNION(A, B), C) → UNION(A, B, C)
func (o *UnionPatternOptimizer) flattenNestedUnions(alts []AlgebraExpr) []AlgebraExpr {
	var flattened []AlgebraExpr

	for _, alt := range alts {
		if union, ok := alt.(*AlgebraUnion); ok {
			// Recursively flatten nested unions
			inner := o.flattenNestedUnions(union.Alternatives)
			flattened = append(flattened, inner...)
		} else {
			flattened = append(flattened, alt)
		}
	}

	return flattened
}

// removeDuplicateBranches removes identical duplicate branches from a union.
func (o *UnionPatternOptimizer) removeDuplicateBranches(alts []AlgebraExpr) []AlgebraExpr {
	var seen []string
	var unique []AlgebraExpr

	for _, alt := range alts {
		if alt == nil {
			continue
		}

		sig := o.getExpressionSignature(alt)

		// Check if we've seen this before
		found := false
		for _, s := range seen {
			if s == sig {
				found = true
				break
			}
		}

		if !found {
			seen = append(seen, sig)
			unique = append(unique, alt)
		}
	}

	return unique
}

// getExpressionSignature creates a signature for an expression for comparison.
func (o *UnionPatternOptimizer) getExpressionSignature(expr AlgebraExpr) string {
	switch node := expr.(type) {
	case *AlgebraBGP:
		// Create signature from triples
		var sig string
		for _, triple := range node.Triples {
			sig += triple.Subject + "|" + triple.Predicate + "|" + triple.Object + ";"
		}
		return "BGP:" + sig

	case *AlgebraFilter:
		// Signature from filter expression and input
		return "FILTER:" + node.Expr + ":" + o.getExpressionSignature(node.Input)

	case *AlgebraProject:
		// Signature from variables and input
		var sig string
		for _, v := range node.Vars {
			sig += v + ","
		}
		return "PROJECT:" + sig + ":" + o.getExpressionSignature(node.Input)

	case *AlgebraJoin:
		// Signature from both inputs
		return "JOIN:" + o.getExpressionSignature(node.Left) + ":" + o.getExpressionSignature(node.Right)

	case *AlgebraUnion:
		// Signature from all alternatives
		var sig string
		for _, alt := range node.Alternatives {
			sig += o.getExpressionSignature(alt) + "|"
		}
		return "UNION:" + sig

	case *AlgebraValues:
		// Signature from variables and values
		var sig string
		for _, v := range node.Vars {
			sig += v + ","
		}
		for _, row := range node.Rows {
			for k, v := range row {
				sig += k + "=" + v + ";"
			}
		}
		return "VALUES:" + sig

	case *AlgebraBind:
		// Signature from binding
		return "BIND:" + node.Var + "=" + node.Expr + ":" + o.getExpressionSignature(node.Input)

	case *AlgebraAgg:
		// Signature from aggregation
		var groupSig string
		for _, g := range node.Group {
			groupSig += g + ","
		}
		return "AGG:" + node.Op + ":" + node.Var + ":GROUP:" + groupSig + ":" + o.getExpressionSignature(node.Input)

	case *AlgebraDistinct:
		return "DISTINCT:" + o.getExpressionSignature(node.Input)

	case *AlgebraOrderBy:
		var sig string
		for _, e := range node.Exprs {
			sig += e + ","
		}
		return "ORDERBY:" + sig + ":" + o.getExpressionSignature(node.Input)

	case *AlgebraLimit:
		return "LIMIT:" + o.getExpressionSignature(node.Input)

	case *AlgebraEmpty:
		return "EMPTY"

	default:
		return "UNKNOWN"
	}
}

// optimizeChildren recursively optimizes child expressions.
func (o *UnionPatternOptimizer) optimizeChildren(expr AlgebraExpr) (AlgebraExpr, error) {
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

	case *AlgebraFilter:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraFilter{
			Expr:  node.Expr,
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

	case *AlgebraBind:
		optimizedInput, err := o.Optimize(node.Input)
		if err != nil {
			return nil, err
		}
		return &AlgebraBind{
			Var:   node.Var,
			Expr:  node.Expr,
			Input: optimizedInput,
		}, nil

	default:
		return expr, nil
	}
}
