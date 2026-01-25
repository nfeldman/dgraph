package sparql

import (
	"fmt"
	"strings"
)

// AlgebraExpr is the base interface for all SPARQL algebra expressions.
// All algebra operators implement this interface, enabling a visitor pattern
// for traversal and transformation.
// Reference: https://www.w3.org/TR/sparql11-query/#sparqlAlgebra
type AlgebraExpr interface {
	// Accept implements the visitor pattern for this expression.
	Accept(visitor AlgebraVisitor) interface{}

	// String returns a human-readable representation for debugging.
	// Format: OperatorName(child1, child2, ...)
	String() string

	// Variables returns all variables referenced anywhere in this expression.
	Variables() []string

	// DefinedVars returns variables that are defined/bound by this expression.
	DefinedVars() []string
}

// AlgebraVisitor is the visitor interface for traversing algebra expressions.
type AlgebraVisitor interface {
	VisitAlgebraBGP(*AlgebraBGP) interface{}
	VisitAlgebraJoin(*AlgebraJoin) interface{}
	VisitAlgebraFilter(*AlgebraFilter) interface{}
	VisitAlgebraLeftJoin(*AlgebraLeftJoin) interface{}
	VisitAlgebraUnion(*AlgebraUnion) interface{}
	VisitAlgebraProject(*AlgebraProject) interface{}
	VisitAlgebraAgg(*AlgebraAgg) interface{}
	VisitAlgebraBind(*AlgebraBind) interface{}
	VisitAlgebraDistinct(*AlgebraDistinct) interface{}
	VisitAlgebraOrderBy(*AlgebraOrderBy) interface{}
	VisitAlgebraLimit(*AlgebraLimit) interface{}
}

// AlgebraBGP represents a Basic Graph Pattern in SPARQL algebra.
// AlgebraBGP is the leaf node of algebra expressions - it contains triples to match.
type AlgebraBGP struct {
	Triples []*Triple
}

func (b *AlgebraBGP) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraBGP(b)
}

func (b *AlgebraBGP) String() string {
	if len(b.Triples) == 0 {
		return "BGP()"
	}
	patterns := make([]string, len(b.Triples))
	for i, t := range b.Triples {
		patterns[i] = fmt.Sprintf("%s %s %s", t.Subject, t.Predicate, t.Object)
	}
	return fmt.Sprintf("BGP(%s)", strings.Join(patterns, ", "))
}

func (b *AlgebraBGP) Variables() []string {
	vars := make(map[string]bool)
	for _, t := range b.Triples {
		if isVar(t.Subject) {
			vars[t.Subject] = true
		}
		if isVar(t.Predicate) {
			vars[t.Predicate] = true
		}
		if t.ObjectIsVar && isVar(t.Object) {
			vars[t.Object] = true
		}
	}
	return mapToSlice(vars)
}

func (b *AlgebraBGP) DefinedVars() []string {
	return b.Variables()
}

// AlgebraJoin represents a join of two algebra expressions.
// Join combines two expressions by finding solutions that satisfy both.
type AlgebraJoin struct {
	Left  AlgebraExpr
	Right AlgebraExpr
}

func (j *AlgebraJoin) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraJoin(j)
}

func (j *AlgebraJoin) String() string {
	return fmt.Sprintf("Join(%s, %s)", j.Left.String(), j.Right.String())
}

func (j *AlgebraJoin) Variables() []string {
	vars := make(map[string]bool)
	for _, v := range j.Left.Variables() {
		vars[v] = true
	}
	for _, v := range j.Right.Variables() {
		vars[v] = true
	}
	return mapToSlice(vars)
}

func (j *AlgebraJoin) DefinedVars() []string {
	return j.Variables()
}

// AlgebraFilter represents a FILTER operation.
// Filter restricts solutions by applying a boolean expression condition.
type AlgebraFilter struct {
	Expr  string // String representation of filter expression
	Input AlgebraExpr
}

func (f *AlgebraFilter) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraFilter(f)
}

func (f *AlgebraFilter) String() string {
	return fmt.Sprintf("Filter(%s, %s)", f.Expr, f.Input.String())
}

func (f *AlgebraFilter) Variables() []string {
	return f.Input.Variables()
}

func (f *AlgebraFilter) DefinedVars() []string {
	return f.Input.DefinedVars()
}

// AlgebraLeftJoin represents an optional join (OPTIONAL patterns).
// Solutions from Input are preserved even if they don't match the optional patterns.
type AlgebraLeftJoin struct {
	Input    AlgebraExpr
	Patterns []*Triple
	Filter   string // Optional filter expression
}

func (lj *AlgebraLeftJoin) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraLeftJoin(lj)
}

func (lj *AlgebraLeftJoin) String() string {
	filterStr := ""
	if lj.Filter != "" {
		filterStr = fmt.Sprintf(", Filter: %s", lj.Filter)
	}
	patterns := make([]string, len(lj.Patterns))
	for i, p := range lj.Patterns {
		patterns[i] = fmt.Sprintf("%s %s %s", p.Subject, p.Predicate, p.Object)
	}
	return fmt.Sprintf("LeftJoin(%s, [%s]%s)", lj.Input.String(), strings.Join(patterns, "; "), filterStr)
}

func (lj *AlgebraLeftJoin) Variables() []string {
	vars := make(map[string]bool)
	for _, v := range lj.Input.Variables() {
		vars[v] = true
	}
	for _, p := range lj.Patterns {
		if isVar(p.Subject) {
			vars[p.Subject] = true
		}
		if isVar(p.Predicate) {
			vars[p.Predicate] = true
		}
		if p.ObjectIsVar && isVar(p.Object) {
			vars[p.Object] = true
		}
	}
	return mapToSlice(vars)
}

func (lj *AlgebraLeftJoin) DefinedVars() []string {
	return lj.Variables()
}

// AlgebraUnion represents a UNION operation.
// Union combines solutions from multiple alternative expressions.
type AlgebraUnion struct {
	Alternatives []AlgebraExpr
}

func (u *AlgebraUnion) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraUnion(u)
}

func (u *AlgebraUnion) String() string {
	alts := make([]string, len(u.Alternatives))
	for i, alt := range u.Alternatives {
		alts[i] = alt.String()
	}
	return fmt.Sprintf("Union(%s)", strings.Join(alts, " | "))
}

func (u *AlgebraUnion) Variables() []string {
	vars := make(map[string]bool)
	for _, alt := range u.Alternatives {
		for _, v := range alt.Variables() {
			vars[v] = true
		}
	}
	return mapToSlice(vars)
}

func (u *AlgebraUnion) DefinedVars() []string {
	return u.Variables()
}

// AlgebraProject represents a SELECT/projection operation.
// Project restricts the output to specific variables.
type AlgebraProject struct {
	Vars  []string
	Input AlgebraExpr
}

func (p *AlgebraProject) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraProject(p)
}

func (p *AlgebraProject) String() string {
	return fmt.Sprintf("Project([%s], %s)", strings.Join(p.Vars, ", "), p.Input.String())
}

func (p *AlgebraProject) Variables() []string {
	return p.Vars
}

func (p *AlgebraProject) DefinedVars() []string {
	return p.Vars
}

// AlgebraAgg represents GROUP BY and aggregate functions.
// Supported operations: COUNT, SUM, MIN, MAX, AVG.
type AlgebraAgg struct {
	Op    string // "COUNT", "SUM", "MIN", "MAX", "AVG"
	Var   string // Variable to aggregate (empty for COUNT(*))
	Group []string
	Input AlgebraExpr
}

func (a *AlgebraAgg) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraAgg(a)
}

func (a *AlgebraAgg) String() string {
	varStr := ""
	if a.Var != "" {
		varStr = fmt.Sprintf("(%s)", a.Var)
	}
	return fmt.Sprintf("Aggregate(%s%s, GROUP BY [%s], %s)", a.Op, varStr, strings.Join(a.Group, ", "), a.Input.String())
}

func (a *AlgebraAgg) Variables() []string {
	vars := make(map[string]bool)
	for _, v := range a.Group {
		vars[v] = true
	}
	return mapToSlice(vars)
}

func (a *AlgebraAgg) DefinedVars() []string {
	return a.Variables()
}

// AlgebraBind represents a BIND operation.
// Bind evaluates an expression and binds the result to a variable.
type AlgebraBind struct {
	Var   string
	Expr  string // Expression string
	Input AlgebraExpr
}

func (b *AlgebraBind) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraBind(b)
}

func (b *AlgebraBind) String() string {
	return fmt.Sprintf("Bind(%s = %s, %s)", b.Var, b.Expr, b.Input.String())
}

func (b *AlgebraBind) Variables() []string {
	vars := make(map[string]bool)
	vars[b.Var] = true
	for _, v := range b.Input.Variables() {
		vars[v] = true
	}
	return mapToSlice(vars)
}

func (b *AlgebraBind) DefinedVars() []string {
	return []string{b.Var}
}

// AlgebraDistinct represents a DISTINCT modifier.
// Distinct removes duplicate solutions.
type AlgebraDistinct struct {
	Input AlgebraExpr
}

func (d *AlgebraDistinct) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraDistinct(d)
}

func (d *AlgebraDistinct) String() string {
	return fmt.Sprintf("Distinct(%s)", d.Input.String())
}

func (d *AlgebraDistinct) Variables() []string {
	return d.Input.Variables()
}

func (d *AlgebraDistinct) DefinedVars() []string {
	return d.Input.DefinedVars()
}

// AlgebraOrderBy represents an ORDER BY modifier.
// OrderBy sorts solutions according to one or more ordering expressions.
type AlgebraOrderBy struct {
	Exprs     []string // Sort expressions
	Ascending []bool   // true for ASC, false for DESC
	Input     AlgebraExpr
}

func (o *AlgebraOrderBy) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraOrderBy(o)
}

func (o *AlgebraOrderBy) String() string {
	orders := make([]string, len(o.Exprs))
	for i, expr := range o.Exprs {
		dir := "ASC"
		if !o.Ascending[i] {
			dir = "DESC"
		}
		orders[i] = fmt.Sprintf("%s %s", expr, dir)
	}
	return fmt.Sprintf("OrderBy([%s], %s)", strings.Join(orders, ", "), o.Input.String())
}

func (o *AlgebraOrderBy) Variables() []string {
	return o.Input.Variables()
}

func (o *AlgebraOrderBy) DefinedVars() []string {
	return o.Input.DefinedVars()
}

// AlgebraLimit represents LIMIT and OFFSET modifiers.
// Limit restricts the number of solutions returned and skips an offset number.
type AlgebraLimit struct {
	Count  int
	Offset int
	Input  AlgebraExpr
}

func (l *AlgebraLimit) Accept(visitor AlgebraVisitor) interface{} {
	return visitor.VisitAlgebraLimit(l)
}

func (l *AlgebraLimit) String() string {
	if l.Offset > 0 {
		return fmt.Sprintf("Limit(%d, %d, %s)", l.Count, l.Offset, l.Input.String())
	}
	return fmt.Sprintf("Limit(%d, %s)", l.Count, l.Input.String())
}

func (l *AlgebraLimit) Variables() []string {
	return l.Input.Variables()
}

func (l *AlgebraLimit) DefinedVars() []string {
	return l.Input.DefinedVars()
}

// Helper function to check if a string is a SPARQL variable (starts with ?)
func isVar(s string) bool {
	return len(s) > 0 && s[0] == '?'
}

// Helper function to convert a map of strings to a slice
func mapToSlice(m map[string]bool) []string {
	result := make([]string, 0, len(m))
	for v := range m {
		result = append(result, v)
	}
	return result
}
