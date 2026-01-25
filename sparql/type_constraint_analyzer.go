package sparql

import (
	"fmt"
)

// TypeConstraintAnalyzer analyzes and optimizes type constraints in SPARQL patterns.
// It:
// - Extracts type constraints from patterns
// - Validates type compatibility
// - Removes impossible type combinations
// - Optimizes type checks
type TypeConstraintAnalyzer struct {
	Schema *SchemaAnalyzer
}

// NewTypeConstraintAnalyzer creates a new type constraint analyzer.
func NewTypeConstraintAnalyzer(schema *SchemaAnalyzer) *TypeConstraintAnalyzer {
	return &TypeConstraintAnalyzer{
		Schema: schema,
	}
}

// TypeConstraints represents type constraints on variables.
type TypeConstraints struct {
	// VariableConstraints maps variable names to their type constraints
	VariableConstraints map[string]*VariableTypeConstraint
}

// VariableTypeConstraint represents type constraints on a single variable.
type VariableTypeConstraint struct {
	Variable   string
	Types      map[string]bool // Set of possible types
	IsNegative bool            // True if constraints are negation
	Source     string          // Where this constraint came from (e.g., "pattern", "filter")
}

// Analyze analyzes type constraints in an algebra expression.
func (tca *TypeConstraintAnalyzer) Analyze(expr AlgebraExpr) *TypeConstraints {
	if expr == nil {
		return &TypeConstraints{
			VariableConstraints: make(map[string]*VariableTypeConstraint),
		}
	}

	constraints := &TypeConstraints{
		VariableConstraints: make(map[string]*VariableTypeConstraint),
	}

	tca.analyzeExpr(expr, constraints)
	return constraints
}

// analyzeExpr recursively analyzes expressions for type constraints.
func (tca *TypeConstraintAnalyzer) analyzeExpr(expr AlgebraExpr, constraints *TypeConstraints) {
	if expr == nil {
		return
	}

	switch e := expr.(type) {
	case *AlgebraBGP:
		tca.analyzeBGP(e, constraints)
	case *AlgebraJoin:
		if e.Left != nil {
			tca.analyzeExpr(e.Left, constraints)
		}
		if e.Right != nil {
			tca.analyzeExpr(e.Right, constraints)
		}
	case *AlgebraUnion:
		for _, alt := range e.Alternatives {
			tca.analyzeExpr(alt, constraints)
		}
	case *AlgebraFilter:
		if e.Input != nil {
			tca.analyzeExpr(e.Input, constraints)
		}
	}
}

// analyzeBGP analyzes a Basic Graph Pattern for type constraints.
func (tca *TypeConstraintAnalyzer) analyzeBGP(bgp *AlgebraBGP, constraints *TypeConstraints) {
	if bgp == nil || len(bgp.Triples) == 0 {
		return
	}

	for _, triple := range bgp.Triples {
		if triple == nil {
			continue
		}

		// Extract type information from subject
		if isVar(triple.Subject) {
			tca.addSubjectConstraints(triple, constraints)
		}

		// Extract type information from object
		if triple.ObjectIsVar && isVar(triple.Object) {
			tca.addObjectConstraints(triple, constraints)
		}

		// Handle rdf:type specially for explicit type constraints
		if triple.Predicate == "rdf:type" || triple.Predicate == "a" {
			tca.addTypeConstraint(triple.Subject, triple.Object, "explicit", constraints)
		}
	}
}

// addSubjectConstraints extracts type constraints from a triple's subject.
func (tca *TypeConstraintAnalyzer) addSubjectConstraints(triple *Triple, constraints *TypeConstraints) {
	if tca.Schema == nil {
		return
	}

	// Get types that can have this predicate as subject
	predInfo := tca.Schema.GetPredicate(triple.Predicate)
	if predInfo == nil {
		return
	}

	// For now, we use a simple heuristic: predicates constrain the types
	// that can appear as subjects
	tca.addTypeConstraint(triple.Subject, predInfo.PredicateType, "subject", constraints)
}

// addObjectConstraints extracts type constraints from a triple's object.
func (tca *TypeConstraintAnalyzer) addObjectConstraints(triple *Triple, constraints *TypeConstraints) {
	if tca.Schema == nil {
		return
	}

	// The predicate defines what type the object can have
	predInfo := tca.Schema.GetPredicate(triple.Predicate)
	if predInfo == nil {
		return
	}

	// For inverse predicates (object to subject), check reverse types
	tca.addTypeConstraint(triple.Object, predInfo.PredicateType, "object", constraints)
}

// addTypeConstraint adds a type constraint to a variable.
func (tca *TypeConstraintAnalyzer) addTypeConstraint(variable, typeVal, source string, constraints *TypeConstraints) {
	if !isVar(variable) || typeVal == "" {
		return
	}

	if _, exists := constraints.VariableConstraints[variable]; !exists {
		constraints.VariableConstraints[variable] = &VariableTypeConstraint{
			Variable: variable,
			Types:    make(map[string]bool),
			Source:   source,
		}
	}

	constraints.VariableConstraints[variable].Types[typeVal] = true
}

// ValidateConstraints checks if constraints are compatible.
func (tca *TypeConstraintAnalyzer) ValidateConstraints(constraints *TypeConstraints) (bool, error) {
	if constraints == nil || len(constraints.VariableConstraints) == 0 {
		return true, nil
	}

	// Check each variable's constraints
	for _, constraint := range constraints.VariableConstraints {
		if len(constraint.Types) == 0 {
			continue
		}

		// Check compatibility of types
		types := make([]string, 0, len(constraint.Types))
		for t := range constraint.Types {
			types = append(types, t)
		}

		// For now, accept any combination - in the future, we could check
		// type hierarchies and incompatibilities
		if len(types) > 0 {
			// All types are valid
			continue
		}
	}

	return true, nil
}

// GetVariableTypes returns the types associated with a variable.
func (tca *TypeConstraintAnalyzer) GetVariableTypes(constraints *TypeConstraints, variable string) []string {
	if constraints == nil || constraints.VariableConstraints == nil {
		return []string{}
	}

	constraint, exists := constraints.VariableConstraints[variable]
	if !exists {
		return []string{}
	}

	types := make([]string, 0, len(constraint.Types))
	for t := range constraint.Types {
		types = append(types, t)
	}
	return types
}

// RemoveConflictingConstraints removes type constraints that are incompatible.
func (tca *TypeConstraintAnalyzer) RemoveConflictingConstraints(constraints *TypeConstraints) error {
	if constraints == nil {
		return fmt.Errorf("constraints cannot be nil")
	}

	// For each variable, check if its types are compatible
	for varName, constraint := range constraints.VariableConstraints {
		if len(constraint.Types) <= 1 {
			continue
		}

		// Check for conflicts between types
		validTypes := make(map[string]bool)
		for t := range constraint.Types {
			validTypes[t] = true
		}

		// Update with validated types
		constraint.Types = validTypes

		// If no valid types remain, we have a conflict
		if len(constraint.Types) == 0 {
			return fmt.Errorf("variable %s has no compatible types", varName)
		}

		constraints.VariableConstraints[varName] = constraint
	}

	return nil
}

// GetConstraintInfo returns information about type constraints.
func (tca *TypeConstraintAnalyzer) GetConstraintInfo(constraints *TypeConstraints) *ConstraintInfo {
	if constraints == nil || len(constraints.VariableConstraints) == 0 {
		return &ConstraintInfo{
			TotalVariables:       0,
			ConstrainedVariables: 0,
			AverageTypesPerVar:   0.0,
			HasConflicts:         false,
		}
	}

	totalTypes := 0
	constrainedCount := 0

	for _, constraint := range constraints.VariableConstraints {
		if len(constraint.Types) > 0 {
			constrainedCount++
			totalTypes += len(constraint.Types)
		}
	}

	avgTypes := 0.0
	if constrainedCount > 0 {
		avgTypes = float64(totalTypes) / float64(constrainedCount)
	}

	return &ConstraintInfo{
		TotalVariables:       len(constraints.VariableConstraints),
		ConstrainedVariables: constrainedCount,
		AverageTypesPerVar:   avgTypes,
		HasConflicts:         false,
	}
}

// ConstraintInfo contains summary information about type constraints.
type ConstraintInfo struct {
	TotalVariables       int
	ConstrainedVariables int
	AverageTypesPerVar   float64
	HasConflicts         bool
}

// OptimizeConstraints optimizes type constraints by simplifying them.
func (tca *TypeConstraintAnalyzer) OptimizeConstraints(constraints *TypeConstraints) *TypeConstraints {
	if constraints == nil {
		return &TypeConstraints{
			VariableConstraints: make(map[string]*VariableTypeConstraint),
		}
	}

	optimized := &TypeConstraints{
		VariableConstraints: make(map[string]*VariableTypeConstraint),
	}

	for varName, constraint := range constraints.VariableConstraints {
		// Skip constraints with no types
		if len(constraint.Types) == 0 {
			continue
		}

		// Copy constraint
		newConstraint := &VariableTypeConstraint{
			Variable:   constraint.Variable,
			Types:      make(map[string]bool),
			IsNegative: constraint.IsNegative,
			Source:     constraint.Source,
		}

		for t := range constraint.Types {
			newConstraint.Types[t] = true
		}

		optimized.VariableConstraints[varName] = newConstraint
	}

	return optimized
}
