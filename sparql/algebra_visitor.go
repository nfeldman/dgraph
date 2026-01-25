package sparql

import (
	"fmt"
	"strings"
)

// AlgebraPrinter is a concrete visitor that prints algebra expressions as formatted trees.
// It's useful for debugging and understanding query structures.
type AlgebraPrinter struct {
	indent int
}

// NewAlgebraPrinter creates a new algebra expression printer.
func NewAlgebraPrinter() *AlgebraPrinter {
	return &AlgebraPrinter{indent: 0}
}

// Print returns a formatted string representation of an algebra expression.
func (p *AlgebraPrinter) Print(expr AlgebraExpr) string {
	p.indent = 0
	return expr.String()
}

func (p *AlgebraPrinter) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	return bgp.String()
}

func (p *AlgebraPrinter) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	return j.String()
}

func (p *AlgebraPrinter) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	return f.String()
}

func (p *AlgebraPrinter) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	return lj.String()
}

func (p *AlgebraPrinter) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	return u.String()
}

func (p *AlgebraPrinter) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	return proj.String()
}

func (p *AlgebraPrinter) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	return agg.String()
}

func (p *AlgebraPrinter) VisitAlgebraBind(b *AlgebraBind) interface{} {
	return b.String()
}

func (p *AlgebraPrinter) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	return d.String()
}

func (p *AlgebraPrinter) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	return o.String()
}

func (p *AlgebraPrinter) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	return l.String()
}

// VariableCollector is a concrete visitor that collects all variables used in an expression.
// It's useful for semantic analysis and validation.
type VariableCollector struct {
	variables map[string]bool
}

// NewVariableCollector creates a new variable collector.
func NewVariableCollector() *VariableCollector {
	return &VariableCollector{
		variables: make(map[string]bool),
	}
}

// Collect returns a sorted slice of all unique variables in an algebra expression.
func (vc *VariableCollector) Collect(expr AlgebraExpr) []string {
	vc.variables = make(map[string]bool)
	expr.Accept(vc)
	return mapToSlice(vc.variables)
}

func (vc *VariableCollector) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	for _, v := range bgp.Variables() {
		vc.variables[v] = true
	}
	return nil
}

func (vc *VariableCollector) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	j.Left.Accept(vc)
	j.Right.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	f.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	lj.Input.Accept(vc)
	for _, v := range getTripleVars(lj.Patterns) {
		vc.variables[v] = true
	}
	return nil
}

func (vc *VariableCollector) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	for _, alt := range u.Alternatives {
		alt.Accept(vc)
	}
	return nil
}

func (vc *VariableCollector) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	for _, v := range proj.Vars {
		vc.variables[v] = true
	}
	proj.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	for _, v := range agg.Group {
		vc.variables[v] = true
	}
	agg.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraBind(b *AlgebraBind) interface{} {
	vc.variables[b.Var] = true
	b.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	d.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	o.Input.Accept(vc)
	return nil
}

func (vc *VariableCollector) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	l.Input.Accept(vc)
	return nil
}

// ExpressionDepthCalculator is a visitor that calculates the depth of an algebra tree.
// Depth = 1 for leaf nodes (BGP), depth = 1 + max(left, right) for inner nodes.
type ExpressionDepthCalculator struct{}

// Depth returns the depth of an algebra expression tree.
func (dc *ExpressionDepthCalculator) Depth(expr AlgebraExpr) int {
	result := expr.Accept(dc)
	if d, ok := result.(int); ok {
		return d
	}
	return 0
}

func (dc *ExpressionDepthCalculator) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	return 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	leftDepth := dc.Depth(j.Left)
	rightDepth := dc.Depth(j.Right)
	if leftDepth > rightDepth {
		return leftDepth + 1
	}
	return rightDepth + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	return dc.Depth(f.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	return dc.Depth(lj.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	maxDepth := 0
	for _, alt := range u.Alternatives {
		d := dc.Depth(alt)
		if d > maxDepth {
			maxDepth = d
		}
	}
	return maxDepth + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	return dc.Depth(proj.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	return dc.Depth(agg.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraBind(b *AlgebraBind) interface{} {
	return dc.Depth(b.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	return dc.Depth(d.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	return dc.Depth(o.Input) + 1
}

func (dc *ExpressionDepthCalculator) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	return dc.Depth(l.Input) + 1
}

// Helper function to extract variables from a slice of triples
func getTripleVars(triples []*Triple) []string {
	vars := make(map[string]bool)
	for _, t := range triples {
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

// ExpressionTreeFormatter formats an algebra expression tree with indentation.
// This is useful for pretty-printing complex query trees.
type ExpressionTreeFormatter struct {
	indent int
	output strings.Builder
}

// Format returns a formatted, indented representation of an algebra expression.
func (etf *ExpressionTreeFormatter) Format(expr AlgebraExpr) string {
	etf.indent = 0
	etf.output.Reset()
	expr.Accept(etf)
	return etf.output.String()
}

func (etf *ExpressionTreeFormatter) writeIndent() {
	for i := 0; i < etf.indent; i++ {
		etf.output.WriteString("  ")
	}
}

func (etf *ExpressionTreeFormatter) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("BGP (%d triples)\n", len(bgp.Triples)))
	etf.indent++
	for _, t := range bgp.Triples {
		etf.writeIndent()
		etf.output.WriteString(fmt.Sprintf("- %s %s %s\n", t.Subject, t.Predicate, t.Object))
	}
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	etf.writeIndent()
	etf.output.WriteString("JOIN\n")
	etf.indent++
	etf.writeIndent()
	etf.output.WriteString("Left:\n")
	etf.indent++
	j.Left.Accept(etf)
	etf.indent--
	etf.writeIndent()
	etf.output.WriteString("Right:\n")
	etf.indent++
	j.Right.Accept(etf)
	etf.indent--
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("FILTER: %s\n", f.Expr))
	etf.indent++
	f.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	etf.writeIndent()
	etf.output.WriteString("LEFT JOIN\n")
	etf.indent++
	etf.writeIndent()
	etf.output.WriteString("Input:\n")
	etf.indent++
	lj.Input.Accept(etf)
	etf.indent--
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("Optional patterns (%d):\n", len(lj.Patterns)))
	for _, p := range lj.Patterns {
		etf.indent++
		etf.writeIndent()
		etf.output.WriteString(fmt.Sprintf("- %s %s %s\n", p.Subject, p.Predicate, p.Object))
		etf.indent--
	}
	if lj.Filter != "" {
		etf.writeIndent()
		etf.output.WriteString(fmt.Sprintf("Filter: %s\n", lj.Filter))
	}
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("UNION (%d alternatives)\n", len(u.Alternatives)))
	etf.indent++
	for i, alt := range u.Alternatives {
		etf.writeIndent()
		etf.output.WriteString(fmt.Sprintf("Alternative %d:\n", i+1))
		etf.indent++
		alt.Accept(etf)
		etf.indent--
	}
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("PROJECT: [%s]\n", strings.Join(proj.Vars, ", ")))
	etf.indent++
	proj.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	etf.writeIndent()
	var varStr string
	if agg.Var != "" {
		varStr = fmt.Sprintf("(%s)", agg.Var)
	}
	etf.output.WriteString(fmt.Sprintf("AGGREGATE: %s%s, GROUP BY [%s]\n", agg.Op, varStr, strings.Join(agg.Group, ", ")))
	etf.indent++
	agg.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraBind(b *AlgebraBind) interface{} {
	etf.writeIndent()
	etf.output.WriteString(fmt.Sprintf("BIND: %s = %s\n", b.Var, b.Expr))
	etf.indent++
	b.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	etf.writeIndent()
	etf.output.WriteString("DISTINCT\n")
	etf.indent++
	d.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	etf.writeIndent()
	orders := make([]string, len(o.Exprs))
	for i, expr := range o.Exprs {
		dir := "ASC"
		if !o.Ascending[i] {
			dir = "DESC"
		}
		orders[i] = fmt.Sprintf("%s %s", expr, dir)
	}
	etf.output.WriteString(fmt.Sprintf("ORDER BY: [%s]\n", strings.Join(orders, ", ")))
	etf.indent++
	o.Input.Accept(etf)
	etf.indent--
	return nil
}

func (etf *ExpressionTreeFormatter) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	etf.writeIndent()
	if l.Offset > 0 {
		etf.output.WriteString(fmt.Sprintf("LIMIT %d OFFSET %d\n", l.Count, l.Offset))
	} else {
		etf.output.WriteString(fmt.Sprintf("LIMIT %d\n", l.Count))
	}
	etf.indent++
	l.Input.Accept(etf)
	etf.indent--
	return nil
}
