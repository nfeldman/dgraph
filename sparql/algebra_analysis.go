package sparql

// InScopeVars returns variables that are visible after evaluating expr.
// This is a conservative approximation using maybe-bound analysis.
func InScopeVars(expr AlgebraExpr) []string {
	return mapToSlice(inferMaybeBound(expr))
}

// MaybeBoundVars returns variables that might be bound after evaluating expr.
func MaybeBoundVars(expr AlgebraExpr) []string {
	return mapToSlice(inferMaybeBound(expr))
}

// DefinitelyBoundVars returns variables guaranteed to be bound after evaluating expr.
func DefinitelyBoundVars(expr AlgebraExpr) []string {
	return mapToSlice(inferDefBound(expr))
}

func inferMaybeBound(expr AlgebraExpr) map[string]bool {
	switch n := expr.(type) {
	case *AlgebraBGP:
		return toSet(n.Variables())
	case *AlgebraValues:
		return toSet(n.Vars)
	case *AlgebraEmpty:
		return map[string]bool{}
	case *AlgebraJoin:
		return unionSets(inferMaybeBound(n.Left), inferMaybeBound(n.Right))
	case *AlgebraLeftJoin:
		base := inferMaybeBound(n.Input)
		optVars := toSet(getTripleVars(n.Patterns))
		return unionSets(base, optVars)
	case *AlgebraUnion:
		acc := make(map[string]bool)
		for _, alt := range n.Alternatives {
			acc = unionSets(acc, inferMaybeBound(alt))
		}
		return acc
	case *AlgebraProject:
		child := inferMaybeBound(n.Input)
		res := make(map[string]bool)
		for _, v := range n.Vars {
			if child[v] {
				res[v] = true
			}
		}
		return res
	case *AlgebraAgg:
		res := toSet(n.Group)
		if n.Var != "" {
			res[n.Var] = true
		}
		return res
	case *AlgebraBind:
		res := inferMaybeBound(n.Input)
		res[n.Var] = true
		return res
	case *AlgebraFilter:
		return inferMaybeBound(n.Input)
	case *AlgebraDistinct:
		return inferMaybeBound(n.Input)
	case *AlgebraOrderBy:
		return inferMaybeBound(n.Input)
	case *AlgebraLimit:
		return inferMaybeBound(n.Input)
	default:
		return map[string]bool{}
	}
}

func inferDefBound(expr AlgebraExpr) map[string]bool {
	switch n := expr.(type) {
	case *AlgebraBGP:
		return toSet(n.Variables())
	case *AlgebraValues:
		if len(n.Rows) == 0 {
			return map[string]bool{}
		}
		return toSet(n.Vars)
	case *AlgebraEmpty:
		return map[string]bool{}
	case *AlgebraJoin:
		return unionSets(inferDefBound(n.Left), inferDefBound(n.Right))
	case *AlgebraLeftJoin:
		return inferDefBound(n.Input)
	case *AlgebraUnion:
		if len(n.Alternatives) == 0 {
			return map[string]bool{}
		}
		acc := inferDefBound(n.Alternatives[0])
		for _, alt := range n.Alternatives[1:] {
			acc = intersectionSets(acc, inferDefBound(alt))
		}
		return acc
	case *AlgebraProject:
		child := inferDefBound(n.Input)
		res := make(map[string]bool)
		for _, v := range n.Vars {
			if child[v] {
				res[v] = true
			}
		}
		return res
	case *AlgebraAgg:
		res := toSet(n.Group)
		if n.Var != "" {
			res[n.Var] = true
		}
		return res
	case *AlgebraBind:
		res := inferDefBound(n.Input)
		res[n.Var] = true
		return res
	case *AlgebraFilter:
		return inferDefBound(n.Input)
	case *AlgebraDistinct:
		return inferDefBound(n.Input)
	case *AlgebraOrderBy:
		return inferDefBound(n.Input)
	case *AlgebraLimit:
		return inferDefBound(n.Input)
	default:
		return map[string]bool{}
	}
}

func intersectionSets(a, b map[string]bool) map[string]bool {
	res := make(map[string]bool)
	for k := range a {
		if b[k] {
			res[k] = true
		}
	}
	return res
}
