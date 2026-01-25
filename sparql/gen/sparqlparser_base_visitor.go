// Code generated from SparqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SparqlParser

import "github.com/antlr4-go/antlr/v4"

type BaseSparqlParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseSparqlParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPrologue(ctx *PrologueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBaseDecl(ctx *BaseDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPrefixDecl(ctx *PrefixDeclContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitSelectQuery(ctx *SelectQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitSelectProjection(ctx *SelectProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitAggregateProjection(ctx *AggregateProjectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitAggregateFunction(ctx *AggregateFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConstructQuery(ctx *ConstructQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitDescribeQuery(ctx *DescribeQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitAskQuery(ctx *AskQueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitDatasetClause(ctx *DatasetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitDefaultGraphClause(ctx *DefaultGraphClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNamedGraphClause(ctx *NamedGraphClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitSourceSelector(ctx *SourceSelectorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitWhereClause(ctx *WhereClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitSolutionModifier(ctx *SolutionModifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGroupClause(ctx *GroupClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGroupCondition(ctx *GroupConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitHavingClause(ctx *HavingClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitHavingCondition(ctx *HavingConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitLimitOffsetClauses(ctx *LimitOffsetClausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitOrderClause(ctx *OrderClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitOrderCondition(ctx *OrderConditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitLimitClause(ctx *LimitClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitOffsetClause(ctx *OffsetClauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGroupGraphPattern(ctx *GroupGraphPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitTriplesBlock(ctx *TriplesBlockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBind(ctx *BindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitOptionalGraphPattern(ctx *OptionalGraphPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGraphGraphPattern(ctx *GraphGraphPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitFilter_(ctx *Filter_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConstraint(ctx *ConstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitFunctionCall(ctx *FunctionCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitArgList(ctx *ArgListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConstructTemplate(ctx *ConstructTemplateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConstructTriples(ctx *ConstructTriplesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitTriplesSameSubject(ctx *TriplesSameSubjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPropertyList(ctx *PropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitObjectList(ctx *ObjectListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitObject_(ctx *Object_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitVerb(ctx *VerbContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitTriplesNode(ctx *TriplesNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBlankNodePropertyList(ctx *BlankNodePropertyListContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitCollection(ctx *CollectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGraphNode(ctx *GraphNodeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitVarOrTerm(ctx *VarOrTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitVarOrIRIref(ctx *VarOrIRIrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitVar_(ctx *Var_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitGraphTerm(ctx *GraphTermContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitValueLogical(ctx *ValueLogicalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitRelationalExpression(ctx *RelationalExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNumericExpression(ctx *NumericExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitUnaryExpression(ctx *UnaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBrackettedExpression(ctx *BrackettedExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBuiltInCall(ctx *BuiltInCallContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitRegexExpression(ctx *RegexExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitIriRefOrFunction(ctx *IriRefOrFunctionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitRdfLiteral(ctx *RdfLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNumericLiteral(ctx *NumericLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNumericLiteralPositive(ctx *NumericLiteralPositiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitNumericLiteralNegative(ctx *NumericLiteralNegativeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitString_(ctx *String_Context) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitIriRef(ctx *IriRefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitPrefixedName(ctx *PrefixedNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseSparqlParserVisitor) VisitBlankNode(ctx *BlankNodeContext) interface{} {
	return v.VisitChildren(ctx)
}
