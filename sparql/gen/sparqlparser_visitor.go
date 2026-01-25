// Code generated from SparqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SparqlParser

import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by SparqlParser.
type SparqlParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by SparqlParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by SparqlParser#prologue.
	VisitPrologue(ctx *PrologueContext) interface{}

	// Visit a parse tree produced by SparqlParser#baseDecl.
	VisitBaseDecl(ctx *BaseDeclContext) interface{}

	// Visit a parse tree produced by SparqlParser#prefixDecl.
	VisitPrefixDecl(ctx *PrefixDeclContext) interface{}

	// Visit a parse tree produced by SparqlParser#selectQuery.
	VisitSelectQuery(ctx *SelectQueryContext) interface{}

	// Visit a parse tree produced by SparqlParser#selectProjection.
	VisitSelectProjection(ctx *SelectProjectionContext) interface{}

	// Visit a parse tree produced by SparqlParser#aggregateProjection.
	VisitAggregateProjection(ctx *AggregateProjectionContext) interface{}

	// Visit a parse tree produced by SparqlParser#aggregateFunction.
	VisitAggregateFunction(ctx *AggregateFunctionContext) interface{}

	// Visit a parse tree produced by SparqlParser#constructQuery.
	VisitConstructQuery(ctx *ConstructQueryContext) interface{}

	// Visit a parse tree produced by SparqlParser#describeQuery.
	VisitDescribeQuery(ctx *DescribeQueryContext) interface{}

	// Visit a parse tree produced by SparqlParser#askQuery.
	VisitAskQuery(ctx *AskQueryContext) interface{}

	// Visit a parse tree produced by SparqlParser#datasetClause.
	VisitDatasetClause(ctx *DatasetClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#defaultGraphClause.
	VisitDefaultGraphClause(ctx *DefaultGraphClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#namedGraphClause.
	VisitNamedGraphClause(ctx *NamedGraphClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#sourceSelector.
	VisitSourceSelector(ctx *SourceSelectorContext) interface{}

	// Visit a parse tree produced by SparqlParser#whereClause.
	VisitWhereClause(ctx *WhereClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#solutionModifier.
	VisitSolutionModifier(ctx *SolutionModifierContext) interface{}

	// Visit a parse tree produced by SparqlParser#groupClause.
	VisitGroupClause(ctx *GroupClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#groupCondition.
	VisitGroupCondition(ctx *GroupConditionContext) interface{}

	// Visit a parse tree produced by SparqlParser#havingClause.
	VisitHavingClause(ctx *HavingClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#havingCondition.
	VisitHavingCondition(ctx *HavingConditionContext) interface{}

	// Visit a parse tree produced by SparqlParser#limitOffsetClauses.
	VisitLimitOffsetClauses(ctx *LimitOffsetClausesContext) interface{}

	// Visit a parse tree produced by SparqlParser#orderClause.
	VisitOrderClause(ctx *OrderClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#orderCondition.
	VisitOrderCondition(ctx *OrderConditionContext) interface{}

	// Visit a parse tree produced by SparqlParser#limitClause.
	VisitLimitClause(ctx *LimitClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#offsetClause.
	VisitOffsetClause(ctx *OffsetClauseContext) interface{}

	// Visit a parse tree produced by SparqlParser#groupGraphPattern.
	VisitGroupGraphPattern(ctx *GroupGraphPatternContext) interface{}

	// Visit a parse tree produced by SparqlParser#triplesBlock.
	VisitTriplesBlock(ctx *TriplesBlockContext) interface{}

	// Visit a parse tree produced by SparqlParser#graphPatternNotTriples.
	VisitGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) interface{}

	// Visit a parse tree produced by SparqlParser#bind.
	VisitBind(ctx *BindContext) interface{}

	// Visit a parse tree produced by SparqlParser#optionalGraphPattern.
	VisitOptionalGraphPattern(ctx *OptionalGraphPatternContext) interface{}

	// Visit a parse tree produced by SparqlParser#graphGraphPattern.
	VisitGraphGraphPattern(ctx *GraphGraphPatternContext) interface{}

	// Visit a parse tree produced by SparqlParser#groupOrUnionGraphPattern.
	VisitGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) interface{}

	// Visit a parse tree produced by SparqlParser#filter_.
	VisitFilter_(ctx *Filter_Context) interface{}

	// Visit a parse tree produced by SparqlParser#constraint.
	VisitConstraint(ctx *ConstraintContext) interface{}

	// Visit a parse tree produced by SparqlParser#functionCall.
	VisitFunctionCall(ctx *FunctionCallContext) interface{}

	// Visit a parse tree produced by SparqlParser#argList.
	VisitArgList(ctx *ArgListContext) interface{}

	// Visit a parse tree produced by SparqlParser#constructTemplate.
	VisitConstructTemplate(ctx *ConstructTemplateContext) interface{}

	// Visit a parse tree produced by SparqlParser#constructTriples.
	VisitConstructTriples(ctx *ConstructTriplesContext) interface{}

	// Visit a parse tree produced by SparqlParser#triplesSameSubject.
	VisitTriplesSameSubject(ctx *TriplesSameSubjectContext) interface{}

	// Visit a parse tree produced by SparqlParser#propertyListNotEmpty.
	VisitPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) interface{}

	// Visit a parse tree produced by SparqlParser#propertyList.
	VisitPropertyList(ctx *PropertyListContext) interface{}

	// Visit a parse tree produced by SparqlParser#objectList.
	VisitObjectList(ctx *ObjectListContext) interface{}

	// Visit a parse tree produced by SparqlParser#object_.
	VisitObject_(ctx *Object_Context) interface{}

	// Visit a parse tree produced by SparqlParser#verb.
	VisitVerb(ctx *VerbContext) interface{}

	// Visit a parse tree produced by SparqlParser#triplesNode.
	VisitTriplesNode(ctx *TriplesNodeContext) interface{}

	// Visit a parse tree produced by SparqlParser#blankNodePropertyList.
	VisitBlankNodePropertyList(ctx *BlankNodePropertyListContext) interface{}

	// Visit a parse tree produced by SparqlParser#collection.
	VisitCollection(ctx *CollectionContext) interface{}

	// Visit a parse tree produced by SparqlParser#graphNode.
	VisitGraphNode(ctx *GraphNodeContext) interface{}

	// Visit a parse tree produced by SparqlParser#varOrTerm.
	VisitVarOrTerm(ctx *VarOrTermContext) interface{}

	// Visit a parse tree produced by SparqlParser#varOrIRIref.
	VisitVarOrIRIref(ctx *VarOrIRIrefContext) interface{}

	// Visit a parse tree produced by SparqlParser#var_.
	VisitVar_(ctx *Var_Context) interface{}

	// Visit a parse tree produced by SparqlParser#graphTerm.
	VisitGraphTerm(ctx *GraphTermContext) interface{}

	// Visit a parse tree produced by SparqlParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#conditionalOrExpression.
	VisitConditionalOrExpression(ctx *ConditionalOrExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#conditionalAndExpression.
	VisitConditionalAndExpression(ctx *ConditionalAndExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#valueLogical.
	VisitValueLogical(ctx *ValueLogicalContext) interface{}

	// Visit a parse tree produced by SparqlParser#relationalExpression.
	VisitRelationalExpression(ctx *RelationalExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#numericExpression.
	VisitNumericExpression(ctx *NumericExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#additiveExpression.
	VisitAdditiveExpression(ctx *AdditiveExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#multiplicativeExpression.
	VisitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#unaryExpression.
	VisitUnaryExpression(ctx *UnaryExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#primaryExpression.
	VisitPrimaryExpression(ctx *PrimaryExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#brackettedExpression.
	VisitBrackettedExpression(ctx *BrackettedExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#builtInCall.
	VisitBuiltInCall(ctx *BuiltInCallContext) interface{}

	// Visit a parse tree produced by SparqlParser#regexExpression.
	VisitRegexExpression(ctx *RegexExpressionContext) interface{}

	// Visit a parse tree produced by SparqlParser#iriRefOrFunction.
	VisitIriRefOrFunction(ctx *IriRefOrFunctionContext) interface{}

	// Visit a parse tree produced by SparqlParser#rdfLiteral.
	VisitRdfLiteral(ctx *RdfLiteralContext) interface{}

	// Visit a parse tree produced by SparqlParser#numericLiteral.
	VisitNumericLiteral(ctx *NumericLiteralContext) interface{}

	// Visit a parse tree produced by SparqlParser#numericLiteralUnsigned.
	VisitNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) interface{}

	// Visit a parse tree produced by SparqlParser#numericLiteralPositive.
	VisitNumericLiteralPositive(ctx *NumericLiteralPositiveContext) interface{}

	// Visit a parse tree produced by SparqlParser#numericLiteralNegative.
	VisitNumericLiteralNegative(ctx *NumericLiteralNegativeContext) interface{}

	// Visit a parse tree produced by SparqlParser#booleanLiteral.
	VisitBooleanLiteral(ctx *BooleanLiteralContext) interface{}

	// Visit a parse tree produced by SparqlParser#string_.
	VisitString_(ctx *String_Context) interface{}

	// Visit a parse tree produced by SparqlParser#iriRef.
	VisitIriRef(ctx *IriRefContext) interface{}

	// Visit a parse tree produced by SparqlParser#prefixedName.
	VisitPrefixedName(ctx *PrefixedNameContext) interface{}

	// Visit a parse tree produced by SparqlParser#blankNode.
	VisitBlankNode(ctx *BlankNodeContext) interface{}
}
