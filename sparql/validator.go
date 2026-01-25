package sparql

import (
	"fmt"
	"strings"
)

// AlgebraValidator performs semantic validation on algebra expressions.
// It checks for:
// - Undefined variable references
// - Variable scoping issues
// - Circular dependencies in BIND expressions
// - Filter validity
// - Aggregate validity
type AlgebraValidator struct {
	definedVars map[string]bool
	errors      []string
	context     validationContext
}

// validationContext tracks state during validation traversal.
type validationContext struct {
	currentScope     map[string]bool
	bindVars         map[string]string // Variable -> expression mapping for cycle detection
	inAggregateScope bool
}

// NewAlgebraValidator creates a new validator.
func NewAlgebraValidator() *AlgebraValidator {
	return &AlgebraValidator{
		definedVars: make(map[string]bool),
		errors:      []string{},
		context: validationContext{
			currentScope: make(map[string]bool),
			bindVars:     make(map[string]string),
		},
	}
}

// Validate checks an algebra expression for semantic correctness.
// Returns an error with all validation issues found, or nil if valid.
func (v *AlgebraValidator) Validate(expr AlgebraExpr) error {
	v.definedVars = make(map[string]bool)
	v.errors = []string{}
	v.context.currentScope = make(map[string]bool)
	v.context.bindVars = make(map[string]string)

	// Traverse expression tree and validate
	expr.Accept(v)

	if len(v.errors) > 0 {
		return fmt.Errorf("algebra validation failed:\n%s", strings.Join(v.errors, "\n"))
	}
	return nil
}

// Implement AlgebraVisitor interface

func (v *AlgebraValidator) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	// BGP defines variables from triple subjects/objects
	for _, t := range bgp.Triples {
		if isVar(t.Subject) {
			v.definedVars[t.Subject] = true
			v.context.currentScope[t.Subject] = true
		}
		if isVar(t.Predicate) {
			v.definedVars[t.Predicate] = true
			v.context.currentScope[t.Predicate] = true
		}
		if t.ObjectIsVar && isVar(t.Object) {
			v.definedVars[t.Object] = true
			v.context.currentScope[t.Object] = true
		}
	}
	return nil
}

func (v *AlgebraValidator) VisitAlgebraValues(val *AlgebraValues) interface{} {
	for _, name := range val.Vars {
		v.definedVars[name] = true
		v.context.currentScope[name] = true
	}
	return nil
}

func (v *AlgebraValidator) VisitAlgebraEmpty(_ *AlgebraEmpty) interface{} {
	return nil
}

func (v *AlgebraValidator) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	// Both sides of join are executed, variables from both are available
	j.Left.Accept(v)
	j.Right.Accept(v)
	return nil
}

func (v *AlgebraValidator) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	// Process input first to define variables
	f.Input.Accept(v)

	// Validate filter references variables that are defined
	if f.Expr != "" {
		refs := extractVariableReferences(f.Expr)
		for _, ref := range refs {
			if !v.definedVars[ref] {
				v.addError(fmt.Sprintf("Variable %s referenced in FILTER but not defined", ref))
			}
		}
	}

	return nil
}

func (v *AlgebraValidator) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	// Input must be processed first
	lj.Input.Accept(v)

	// Variables from the optional patterns
	for _, t := range lj.Patterns {
		if isVar(t.Subject) {
			v.definedVars[t.Subject] = true
		}
		if isVar(t.Predicate) {
			v.definedVars[t.Predicate] = true
		}
		if t.ObjectIsVar && isVar(t.Object) {
			v.definedVars[t.Object] = true
		}
	}

	// Validate optional filter
	if lj.Filter != "" {
		refs := extractVariableReferences(lj.Filter)
		for _, ref := range refs {
			if !v.definedVars[ref] {
				v.addError(fmt.Sprintf("Variable %s referenced in LEFT JOIN filter but not defined", ref))
			}
		}
	}

	return nil
}

func (v *AlgebraValidator) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	// All union alternatives are executed, variables from each must be consistent
	allVars := make(map[string]bool)

	for i, alt := range u.Alternatives {
		// Save current state
		savedVars := make(map[string]bool)
		for k, val := range v.definedVars {
			savedVars[k] = val
		}

		// Reset to initial state and validate this alternative
		v.definedVars = make(map[string]bool)
		alt.Accept(v)

		// Collect variables from this alternative
		altVars := make(map[string]bool)
		for k := range v.definedVars {
			altVars[k] = true
		}

		if i == 0 {
			allVars = altVars
		} else {
			// Each alternative must define the same variables
			for varName := range allVars {
				if !altVars[varName] {
					v.addError(fmt.Sprintf("Variable %s not defined in all UNION alternatives", varName))
				}
			}
			for varName := range altVars {
				if !allVars[varName] {
					v.addError(fmt.Sprintf("Variable %s not defined in all UNION alternatives", varName))
				}
			}
		}

		// Restore to saved state for next iteration
		v.definedVars = savedVars
	}

	// After union, all variables from all alternatives are available
	for varName := range allVars {
		v.definedVars[varName] = true
	}

	return nil
}

func (v *AlgebraValidator) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	// Process input first to define variables
	proj.Input.Accept(v)

	// All projected variables must be defined
	for _, varName := range proj.Vars {
		if !isVar(varName) {
			continue
		}
		if !v.definedVars[varName] {
			v.addError(fmt.Sprintf("Variable %s in SELECT not defined", varName))
		}
	}

	return nil
}

func (v *AlgebraValidator) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	// Process input first
	v.context.inAggregateScope = true
	agg.Input.Accept(v)
	v.context.inAggregateScope = false

	// GROUP BY variables must be defined
	for _, groupVar := range agg.Group {
		if !isVar(groupVar) {
			continue
		}
		if !v.definedVars[groupVar] {
			v.addError(fmt.Sprintf("Variable %s in GROUP BY not defined", groupVar))
		}
	}

	// Aggregate target variable must be defined (or is being defined)
	if agg.Var != "" && isVar(agg.Var) {
		// Validate it's being applied to appropriate expression
		// For now, accept any defined variable
		if !v.definedVars[agg.Var] && agg.Op != "count" {
			v.addError(fmt.Sprintf("Aggregate %s applied to undefined variable %s", agg.Op, agg.Var))
		}
	}

	// After aggregation, only GROUP BY variables and aggregates are available
	newDefined := make(map[string]bool)
	for _, groupVar := range agg.Group {
		newDefined[groupVar] = true
	}
	if agg.Var != "" {
		newDefined[agg.Var] = true
	}
	v.definedVars = newDefined

	return nil
}

func (v *AlgebraValidator) VisitAlgebraBind(b *AlgebraBind) interface{} {
	// Process input first to define its variables
	b.Input.Accept(v)

	// Check for circular dependency: variable being bound must not appear in its expression
	if b.Expr != "" {
		refs := extractVariableReferences(b.Expr)
		for _, ref := range refs {
			if ref == b.Var {
				v.addError(fmt.Sprintf("Circular dependency: Variable %s referenced in its own BIND expression", b.Var))
			}
			if !v.definedVars[ref] {
				v.addError(fmt.Sprintf("Variable %s referenced in BIND but not defined", ref))
			}
		}
	}

	// BIND defines a new variable
	v.definedVars[b.Var] = true
	v.context.bindVars[b.Var] = b.Expr

	return nil
}

func (v *AlgebraValidator) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	// DISTINCT doesn't affect variable scoping
	d.Input.Accept(v)
	return nil
}

func (v *AlgebraValidator) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	// Process input first
	o.Input.Accept(v)

	// All ORDER BY expressions must reference defined variables
	for _, expr := range o.Exprs {
		refs := extractVariableReferences(expr)
		for _, ref := range refs {
			if !v.definedVars[ref] {
				v.addError(fmt.Sprintf("Variable %s in ORDER BY not defined", ref))
			}
		}
	}

	return nil
}

func (v *AlgebraValidator) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	// LIMIT/OFFSET are pure restrictions, don't affect variables
	l.Input.Accept(v)
	return nil
}

// Helper methods

func (v *AlgebraValidator) addError(msg string) {
	v.errors = append(v.errors, msg)
}

// extractVariableReferences extracts all variable references from an expression string.
// Variables are identified as starting with "?".
func extractVariableReferences(expr string) []string {
	vars := make(map[string]bool)

	parts := strings.FieldsFunc(expr, func(r rune) bool {
		return r == ' ' || r == '(' || r == ')' || r == ',' || r == '=' || r == '<' || r == '>' || r == '!'
	})

	for _, part := range parts {
		if isVar(part) {
			vars[part] = true
		}
	}

	return mapToSlice(vars)
}

// ValidationStats provides statistics about a validated algebra expression.
type ValidationStats struct {
	TotalVariables   int
	DefinedVariables int
	OperatorCount    int
	MaxDepth         int
}

// GetValidationStats returns statistics about an algebra expression.
func GetValidationStats(expr AlgebraExpr) *ValidationStats {
	collector := NewVariableCollector()
	vars := collector.Collect(expr)

	depthCalc := &ExpressionDepthCalculator{}
	depth := depthCalc.Depth(expr)

	counter := &operatorCounter{}
	expr.Accept(counter)

	return &ValidationStats{
		TotalVariables:   len(vars),
		DefinedVariables: len(vars),
		OperatorCount:    counter.count,
		MaxDepth:         depth,
	}
}

// operatorCounter is an internal visitor that counts operators.
type operatorCounter struct {
	count int
}

func (oc *operatorCounter) VisitAlgebraBGP(bgp *AlgebraBGP) interface{} {
	oc.count++
	return nil
}

func (oc *operatorCounter) VisitAlgebraValues(val *AlgebraValues) interface{} {
	oc.count++
	return nil
}

func (oc *operatorCounter) VisitAlgebraEmpty(_ *AlgebraEmpty) interface{} {
	oc.count++
	return nil
}

func (oc *operatorCounter) VisitAlgebraJoin(j *AlgebraJoin) interface{} {
	oc.count++
	j.Left.Accept(oc)
	j.Right.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraFilter(f *AlgebraFilter) interface{} {
	oc.count++
	f.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraLeftJoin(lj *AlgebraLeftJoin) interface{} {
	oc.count++
	lj.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraUnion(u *AlgebraUnion) interface{} {
	oc.count++
	for _, alt := range u.Alternatives {
		alt.Accept(oc)
	}
	return nil
}

func (oc *operatorCounter) VisitAlgebraProject(proj *AlgebraProject) interface{} {
	oc.count++
	proj.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraAgg(agg *AlgebraAgg) interface{} {
	oc.count++
	agg.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraBind(b *AlgebraBind) interface{} {
	oc.count++
	b.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraDistinct(d *AlgebraDistinct) interface{} {
	oc.count++
	d.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraOrderBy(o *AlgebraOrderBy) interface{} {
	oc.count++
	o.Input.Accept(oc)
	return nil
}

func (oc *operatorCounter) VisitAlgebraLimit(l *AlgebraLimit) interface{} {
	oc.count++
	l.Input.Accept(oc)
	return nil
}
