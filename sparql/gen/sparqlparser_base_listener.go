// Code generated from SparqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SparqlParser

import "github.com/antlr4-go/antlr/v4"

// BaseSparqlParserListener is a complete listener for a parse tree produced by SparqlParser.
type BaseSparqlParserListener struct{}

var _ SparqlParserListener = &BaseSparqlParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseSparqlParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseSparqlParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseSparqlParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseSparqlParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterQuery is called when production query is entered.
func (s *BaseSparqlParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseSparqlParserListener) ExitQuery(ctx *QueryContext) {}

// EnterPrologue is called when production prologue is entered.
func (s *BaseSparqlParserListener) EnterPrologue(ctx *PrologueContext) {}

// ExitPrologue is called when production prologue is exited.
func (s *BaseSparqlParserListener) ExitPrologue(ctx *PrologueContext) {}

// EnterBaseDecl is called when production baseDecl is entered.
func (s *BaseSparqlParserListener) EnterBaseDecl(ctx *BaseDeclContext) {}

// ExitBaseDecl is called when production baseDecl is exited.
func (s *BaseSparqlParserListener) ExitBaseDecl(ctx *BaseDeclContext) {}

// EnterPrefixDecl is called when production prefixDecl is entered.
func (s *BaseSparqlParserListener) EnterPrefixDecl(ctx *PrefixDeclContext) {}

// ExitPrefixDecl is called when production prefixDecl is exited.
func (s *BaseSparqlParserListener) ExitPrefixDecl(ctx *PrefixDeclContext) {}

// EnterSelectQuery is called when production selectQuery is entered.
func (s *BaseSparqlParserListener) EnterSelectQuery(ctx *SelectQueryContext) {}

// ExitSelectQuery is called when production selectQuery is exited.
func (s *BaseSparqlParserListener) ExitSelectQuery(ctx *SelectQueryContext) {}

// EnterSelectProjection is called when production selectProjection is entered.
func (s *BaseSparqlParserListener) EnterSelectProjection(ctx *SelectProjectionContext) {}

// ExitSelectProjection is called when production selectProjection is exited.
func (s *BaseSparqlParserListener) ExitSelectProjection(ctx *SelectProjectionContext) {}

// EnterAggregateProjection is called when production aggregateProjection is entered.
func (s *BaseSparqlParserListener) EnterAggregateProjection(ctx *AggregateProjectionContext) {}

// ExitAggregateProjection is called when production aggregateProjection is exited.
func (s *BaseSparqlParserListener) ExitAggregateProjection(ctx *AggregateProjectionContext) {}

// EnterAggregateFunction is called when production aggregateFunction is entered.
func (s *BaseSparqlParserListener) EnterAggregateFunction(ctx *AggregateFunctionContext) {}

// ExitAggregateFunction is called when production aggregateFunction is exited.
func (s *BaseSparqlParserListener) ExitAggregateFunction(ctx *AggregateFunctionContext) {}

// EnterConstructQuery is called when production constructQuery is entered.
func (s *BaseSparqlParserListener) EnterConstructQuery(ctx *ConstructQueryContext) {}

// ExitConstructQuery is called when production constructQuery is exited.
func (s *BaseSparqlParserListener) ExitConstructQuery(ctx *ConstructQueryContext) {}

// EnterDescribeQuery is called when production describeQuery is entered.
func (s *BaseSparqlParserListener) EnterDescribeQuery(ctx *DescribeQueryContext) {}

// ExitDescribeQuery is called when production describeQuery is exited.
func (s *BaseSparqlParserListener) ExitDescribeQuery(ctx *DescribeQueryContext) {}

// EnterAskQuery is called when production askQuery is entered.
func (s *BaseSparqlParserListener) EnterAskQuery(ctx *AskQueryContext) {}

// ExitAskQuery is called when production askQuery is exited.
func (s *BaseSparqlParserListener) ExitAskQuery(ctx *AskQueryContext) {}

// EnterDatasetClause is called when production datasetClause is entered.
func (s *BaseSparqlParserListener) EnterDatasetClause(ctx *DatasetClauseContext) {}

// ExitDatasetClause is called when production datasetClause is exited.
func (s *BaseSparqlParserListener) ExitDatasetClause(ctx *DatasetClauseContext) {}

// EnterDefaultGraphClause is called when production defaultGraphClause is entered.
func (s *BaseSparqlParserListener) EnterDefaultGraphClause(ctx *DefaultGraphClauseContext) {}

// ExitDefaultGraphClause is called when production defaultGraphClause is exited.
func (s *BaseSparqlParserListener) ExitDefaultGraphClause(ctx *DefaultGraphClauseContext) {}

// EnterNamedGraphClause is called when production namedGraphClause is entered.
func (s *BaseSparqlParserListener) EnterNamedGraphClause(ctx *NamedGraphClauseContext) {}

// ExitNamedGraphClause is called when production namedGraphClause is exited.
func (s *BaseSparqlParserListener) ExitNamedGraphClause(ctx *NamedGraphClauseContext) {}

// EnterSourceSelector is called when production sourceSelector is entered.
func (s *BaseSparqlParserListener) EnterSourceSelector(ctx *SourceSelectorContext) {}

// ExitSourceSelector is called when production sourceSelector is exited.
func (s *BaseSparqlParserListener) ExitSourceSelector(ctx *SourceSelectorContext) {}

// EnterWhereClause is called when production whereClause is entered.
func (s *BaseSparqlParserListener) EnterWhereClause(ctx *WhereClauseContext) {}

// ExitWhereClause is called when production whereClause is exited.
func (s *BaseSparqlParserListener) ExitWhereClause(ctx *WhereClauseContext) {}

// EnterSolutionModifier is called when production solutionModifier is entered.
func (s *BaseSparqlParserListener) EnterSolutionModifier(ctx *SolutionModifierContext) {}

// ExitSolutionModifier is called when production solutionModifier is exited.
func (s *BaseSparqlParserListener) ExitSolutionModifier(ctx *SolutionModifierContext) {}

// EnterGroupClause is called when production groupClause is entered.
func (s *BaseSparqlParserListener) EnterGroupClause(ctx *GroupClauseContext) {}

// ExitGroupClause is called when production groupClause is exited.
func (s *BaseSparqlParserListener) ExitGroupClause(ctx *GroupClauseContext) {}

// EnterGroupCondition is called when production groupCondition is entered.
func (s *BaseSparqlParserListener) EnterGroupCondition(ctx *GroupConditionContext) {}

// ExitGroupCondition is called when production groupCondition is exited.
func (s *BaseSparqlParserListener) ExitGroupCondition(ctx *GroupConditionContext) {}

// EnterHavingClause is called when production havingClause is entered.
func (s *BaseSparqlParserListener) EnterHavingClause(ctx *HavingClauseContext) {}

// ExitHavingClause is called when production havingClause is exited.
func (s *BaseSparqlParserListener) ExitHavingClause(ctx *HavingClauseContext) {}

// EnterHavingCondition is called when production havingCondition is entered.
func (s *BaseSparqlParserListener) EnterHavingCondition(ctx *HavingConditionContext) {}

// ExitHavingCondition is called when production havingCondition is exited.
func (s *BaseSparqlParserListener) ExitHavingCondition(ctx *HavingConditionContext) {}

// EnterLimitOffsetClauses is called when production limitOffsetClauses is entered.
func (s *BaseSparqlParserListener) EnterLimitOffsetClauses(ctx *LimitOffsetClausesContext) {}

// ExitLimitOffsetClauses is called when production limitOffsetClauses is exited.
func (s *BaseSparqlParserListener) ExitLimitOffsetClauses(ctx *LimitOffsetClausesContext) {}

// EnterOrderClause is called when production orderClause is entered.
func (s *BaseSparqlParserListener) EnterOrderClause(ctx *OrderClauseContext) {}

// ExitOrderClause is called when production orderClause is exited.
func (s *BaseSparqlParserListener) ExitOrderClause(ctx *OrderClauseContext) {}

// EnterOrderCondition is called when production orderCondition is entered.
func (s *BaseSparqlParserListener) EnterOrderCondition(ctx *OrderConditionContext) {}

// ExitOrderCondition is called when production orderCondition is exited.
func (s *BaseSparqlParserListener) ExitOrderCondition(ctx *OrderConditionContext) {}

// EnterLimitClause is called when production limitClause is entered.
func (s *BaseSparqlParserListener) EnterLimitClause(ctx *LimitClauseContext) {}

// ExitLimitClause is called when production limitClause is exited.
func (s *BaseSparqlParserListener) ExitLimitClause(ctx *LimitClauseContext) {}

// EnterOffsetClause is called when production offsetClause is entered.
func (s *BaseSparqlParserListener) EnterOffsetClause(ctx *OffsetClauseContext) {}

// ExitOffsetClause is called when production offsetClause is exited.
func (s *BaseSparqlParserListener) ExitOffsetClause(ctx *OffsetClauseContext) {}

// EnterGroupGraphPattern is called when production groupGraphPattern is entered.
func (s *BaseSparqlParserListener) EnterGroupGraphPattern(ctx *GroupGraphPatternContext) {}

// ExitGroupGraphPattern is called when production groupGraphPattern is exited.
func (s *BaseSparqlParserListener) ExitGroupGraphPattern(ctx *GroupGraphPatternContext) {}

// EnterTriplesBlock is called when production triplesBlock is entered.
func (s *BaseSparqlParserListener) EnterTriplesBlock(ctx *TriplesBlockContext) {}

// ExitTriplesBlock is called when production triplesBlock is exited.
func (s *BaseSparqlParserListener) ExitTriplesBlock(ctx *TriplesBlockContext) {}

// EnterGraphPatternNotTriples is called when production graphPatternNotTriples is entered.
func (s *BaseSparqlParserListener) EnterGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) {}

// ExitGraphPatternNotTriples is called when production graphPatternNotTriples is exited.
func (s *BaseSparqlParserListener) ExitGraphPatternNotTriples(ctx *GraphPatternNotTriplesContext) {}

// EnterBind is called when production bind is entered.
func (s *BaseSparqlParserListener) EnterBind(ctx *BindContext) {}

// ExitBind is called when production bind is exited.
func (s *BaseSparqlParserListener) ExitBind(ctx *BindContext) {}

// EnterOptionalGraphPattern is called when production optionalGraphPattern is entered.
func (s *BaseSparqlParserListener) EnterOptionalGraphPattern(ctx *OptionalGraphPatternContext) {}

// ExitOptionalGraphPattern is called when production optionalGraphPattern is exited.
func (s *BaseSparqlParserListener) ExitOptionalGraphPattern(ctx *OptionalGraphPatternContext) {}

// EnterGraphGraphPattern is called when production graphGraphPattern is entered.
func (s *BaseSparqlParserListener) EnterGraphGraphPattern(ctx *GraphGraphPatternContext) {}

// ExitGraphGraphPattern is called when production graphGraphPattern is exited.
func (s *BaseSparqlParserListener) ExitGraphGraphPattern(ctx *GraphGraphPatternContext) {}

// EnterGroupOrUnionGraphPattern is called when production groupOrUnionGraphPattern is entered.
func (s *BaseSparqlParserListener) EnterGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) {
}

// ExitGroupOrUnionGraphPattern is called when production groupOrUnionGraphPattern is exited.
func (s *BaseSparqlParserListener) ExitGroupOrUnionGraphPattern(ctx *GroupOrUnionGraphPatternContext) {
}

// EnterFilter_ is called when production filter_ is entered.
func (s *BaseSparqlParserListener) EnterFilter_(ctx *Filter_Context) {}

// ExitFilter_ is called when production filter_ is exited.
func (s *BaseSparqlParserListener) ExitFilter_(ctx *Filter_Context) {}

// EnterConstraint is called when production constraint is entered.
func (s *BaseSparqlParserListener) EnterConstraint(ctx *ConstraintContext) {}

// ExitConstraint is called when production constraint is exited.
func (s *BaseSparqlParserListener) ExitConstraint(ctx *ConstraintContext) {}

// EnterFunctionCall is called when production functionCall is entered.
func (s *BaseSparqlParserListener) EnterFunctionCall(ctx *FunctionCallContext) {}

// ExitFunctionCall is called when production functionCall is exited.
func (s *BaseSparqlParserListener) ExitFunctionCall(ctx *FunctionCallContext) {}

// EnterArgList is called when production argList is entered.
func (s *BaseSparqlParserListener) EnterArgList(ctx *ArgListContext) {}

// ExitArgList is called when production argList is exited.
func (s *BaseSparqlParserListener) ExitArgList(ctx *ArgListContext) {}

// EnterConstructTemplate is called when production constructTemplate is entered.
func (s *BaseSparqlParserListener) EnterConstructTemplate(ctx *ConstructTemplateContext) {}

// ExitConstructTemplate is called when production constructTemplate is exited.
func (s *BaseSparqlParserListener) ExitConstructTemplate(ctx *ConstructTemplateContext) {}

// EnterConstructTriples is called when production constructTriples is entered.
func (s *BaseSparqlParserListener) EnterConstructTriples(ctx *ConstructTriplesContext) {}

// ExitConstructTriples is called when production constructTriples is exited.
func (s *BaseSparqlParserListener) ExitConstructTriples(ctx *ConstructTriplesContext) {}

// EnterTriplesSameSubject is called when production triplesSameSubject is entered.
func (s *BaseSparqlParserListener) EnterTriplesSameSubject(ctx *TriplesSameSubjectContext) {}

// ExitTriplesSameSubject is called when production triplesSameSubject is exited.
func (s *BaseSparqlParserListener) ExitTriplesSameSubject(ctx *TriplesSameSubjectContext) {}

// EnterPropertyListNotEmpty is called when production propertyListNotEmpty is entered.
func (s *BaseSparqlParserListener) EnterPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) {}

// ExitPropertyListNotEmpty is called when production propertyListNotEmpty is exited.
func (s *BaseSparqlParserListener) ExitPropertyListNotEmpty(ctx *PropertyListNotEmptyContext) {}

// EnterPropertyList is called when production propertyList is entered.
func (s *BaseSparqlParserListener) EnterPropertyList(ctx *PropertyListContext) {}

// ExitPropertyList is called when production propertyList is exited.
func (s *BaseSparqlParserListener) ExitPropertyList(ctx *PropertyListContext) {}

// EnterObjectList is called when production objectList is entered.
func (s *BaseSparqlParserListener) EnterObjectList(ctx *ObjectListContext) {}

// ExitObjectList is called when production objectList is exited.
func (s *BaseSparqlParserListener) ExitObjectList(ctx *ObjectListContext) {}

// EnterObject_ is called when production object_ is entered.
func (s *BaseSparqlParserListener) EnterObject_(ctx *Object_Context) {}

// ExitObject_ is called when production object_ is exited.
func (s *BaseSparqlParserListener) ExitObject_(ctx *Object_Context) {}

// EnterVerb is called when production verb is entered.
func (s *BaseSparqlParserListener) EnterVerb(ctx *VerbContext) {}

// ExitVerb is called when production verb is exited.
func (s *BaseSparqlParserListener) ExitVerb(ctx *VerbContext) {}

// EnterTriplesNode is called when production triplesNode is entered.
func (s *BaseSparqlParserListener) EnterTriplesNode(ctx *TriplesNodeContext) {}

// ExitTriplesNode is called when production triplesNode is exited.
func (s *BaseSparqlParserListener) ExitTriplesNode(ctx *TriplesNodeContext) {}

// EnterBlankNodePropertyList is called when production blankNodePropertyList is entered.
func (s *BaseSparqlParserListener) EnterBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// ExitBlankNodePropertyList is called when production blankNodePropertyList is exited.
func (s *BaseSparqlParserListener) ExitBlankNodePropertyList(ctx *BlankNodePropertyListContext) {}

// EnterCollection is called when production collection is entered.
func (s *BaseSparqlParserListener) EnterCollection(ctx *CollectionContext) {}

// ExitCollection is called when production collection is exited.
func (s *BaseSparqlParserListener) ExitCollection(ctx *CollectionContext) {}

// EnterGraphNode is called when production graphNode is entered.
func (s *BaseSparqlParserListener) EnterGraphNode(ctx *GraphNodeContext) {}

// ExitGraphNode is called when production graphNode is exited.
func (s *BaseSparqlParserListener) ExitGraphNode(ctx *GraphNodeContext) {}

// EnterVarOrTerm is called when production varOrTerm is entered.
func (s *BaseSparqlParserListener) EnterVarOrTerm(ctx *VarOrTermContext) {}

// ExitVarOrTerm is called when production varOrTerm is exited.
func (s *BaseSparqlParserListener) ExitVarOrTerm(ctx *VarOrTermContext) {}

// EnterVarOrIRIref is called when production varOrIRIref is entered.
func (s *BaseSparqlParserListener) EnterVarOrIRIref(ctx *VarOrIRIrefContext) {}

// ExitVarOrIRIref is called when production varOrIRIref is exited.
func (s *BaseSparqlParserListener) ExitVarOrIRIref(ctx *VarOrIRIrefContext) {}

// EnterVar_ is called when production var_ is entered.
func (s *BaseSparqlParserListener) EnterVar_(ctx *Var_Context) {}

// ExitVar_ is called when production var_ is exited.
func (s *BaseSparqlParserListener) ExitVar_(ctx *Var_Context) {}

// EnterGraphTerm is called when production graphTerm is entered.
func (s *BaseSparqlParserListener) EnterGraphTerm(ctx *GraphTermContext) {}

// ExitGraphTerm is called when production graphTerm is exited.
func (s *BaseSparqlParserListener) ExitGraphTerm(ctx *GraphTermContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseSparqlParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseSparqlParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterConditionalOrExpression is called when production conditionalOrExpression is entered.
func (s *BaseSparqlParserListener) EnterConditionalOrExpression(ctx *ConditionalOrExpressionContext) {
}

// ExitConditionalOrExpression is called when production conditionalOrExpression is exited.
func (s *BaseSparqlParserListener) ExitConditionalOrExpression(ctx *ConditionalOrExpressionContext) {}

// EnterConditionalAndExpression is called when production conditionalAndExpression is entered.
func (s *BaseSparqlParserListener) EnterConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// ExitConditionalAndExpression is called when production conditionalAndExpression is exited.
func (s *BaseSparqlParserListener) ExitConditionalAndExpression(ctx *ConditionalAndExpressionContext) {
}

// EnterValueLogical is called when production valueLogical is entered.
func (s *BaseSparqlParserListener) EnterValueLogical(ctx *ValueLogicalContext) {}

// ExitValueLogical is called when production valueLogical is exited.
func (s *BaseSparqlParserListener) ExitValueLogical(ctx *ValueLogicalContext) {}

// EnterRelationalExpression is called when production relationalExpression is entered.
func (s *BaseSparqlParserListener) EnterRelationalExpression(ctx *RelationalExpressionContext) {}

// ExitRelationalExpression is called when production relationalExpression is exited.
func (s *BaseSparqlParserListener) ExitRelationalExpression(ctx *RelationalExpressionContext) {}

// EnterNumericExpression is called when production numericExpression is entered.
func (s *BaseSparqlParserListener) EnterNumericExpression(ctx *NumericExpressionContext) {}

// ExitNumericExpression is called when production numericExpression is exited.
func (s *BaseSparqlParserListener) ExitNumericExpression(ctx *NumericExpressionContext) {}

// EnterAdditiveExpression is called when production additiveExpression is entered.
func (s *BaseSparqlParserListener) EnterAdditiveExpression(ctx *AdditiveExpressionContext) {}

// ExitAdditiveExpression is called when production additiveExpression is exited.
func (s *BaseSparqlParserListener) ExitAdditiveExpression(ctx *AdditiveExpressionContext) {}

// EnterMultiplicativeExpression is called when production multiplicativeExpression is entered.
func (s *BaseSparqlParserListener) EnterMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// ExitMultiplicativeExpression is called when production multiplicativeExpression is exited.
func (s *BaseSparqlParserListener) ExitMultiplicativeExpression(ctx *MultiplicativeExpressionContext) {
}

// EnterUnaryExpression is called when production unaryExpression is entered.
func (s *BaseSparqlParserListener) EnterUnaryExpression(ctx *UnaryExpressionContext) {}

// ExitUnaryExpression is called when production unaryExpression is exited.
func (s *BaseSparqlParserListener) ExitUnaryExpression(ctx *UnaryExpressionContext) {}

// EnterPrimaryExpression is called when production primaryExpression is entered.
func (s *BaseSparqlParserListener) EnterPrimaryExpression(ctx *PrimaryExpressionContext) {}

// ExitPrimaryExpression is called when production primaryExpression is exited.
func (s *BaseSparqlParserListener) ExitPrimaryExpression(ctx *PrimaryExpressionContext) {}

// EnterBrackettedExpression is called when production brackettedExpression is entered.
func (s *BaseSparqlParserListener) EnterBrackettedExpression(ctx *BrackettedExpressionContext) {}

// ExitBrackettedExpression is called when production brackettedExpression is exited.
func (s *BaseSparqlParserListener) ExitBrackettedExpression(ctx *BrackettedExpressionContext) {}

// EnterBuiltInCall is called when production builtInCall is entered.
func (s *BaseSparqlParserListener) EnterBuiltInCall(ctx *BuiltInCallContext) {}

// ExitBuiltInCall is called when production builtInCall is exited.
func (s *BaseSparqlParserListener) ExitBuiltInCall(ctx *BuiltInCallContext) {}

// EnterRegexExpression is called when production regexExpression is entered.
func (s *BaseSparqlParserListener) EnterRegexExpression(ctx *RegexExpressionContext) {}

// ExitRegexExpression is called when production regexExpression is exited.
func (s *BaseSparqlParserListener) ExitRegexExpression(ctx *RegexExpressionContext) {}

// EnterIriRefOrFunction is called when production iriRefOrFunction is entered.
func (s *BaseSparqlParserListener) EnterIriRefOrFunction(ctx *IriRefOrFunctionContext) {}

// ExitIriRefOrFunction is called when production iriRefOrFunction is exited.
func (s *BaseSparqlParserListener) ExitIriRefOrFunction(ctx *IriRefOrFunctionContext) {}

// EnterRdfLiteral is called when production rdfLiteral is entered.
func (s *BaseSparqlParserListener) EnterRdfLiteral(ctx *RdfLiteralContext) {}

// ExitRdfLiteral is called when production rdfLiteral is exited.
func (s *BaseSparqlParserListener) ExitRdfLiteral(ctx *RdfLiteralContext) {}

// EnterNumericLiteral is called when production numericLiteral is entered.
func (s *BaseSparqlParserListener) EnterNumericLiteral(ctx *NumericLiteralContext) {}

// ExitNumericLiteral is called when production numericLiteral is exited.
func (s *BaseSparqlParserListener) ExitNumericLiteral(ctx *NumericLiteralContext) {}

// EnterNumericLiteralUnsigned is called when production numericLiteralUnsigned is entered.
func (s *BaseSparqlParserListener) EnterNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) {}

// ExitNumericLiteralUnsigned is called when production numericLiteralUnsigned is exited.
func (s *BaseSparqlParserListener) ExitNumericLiteralUnsigned(ctx *NumericLiteralUnsignedContext) {}

// EnterNumericLiteralPositive is called when production numericLiteralPositive is entered.
func (s *BaseSparqlParserListener) EnterNumericLiteralPositive(ctx *NumericLiteralPositiveContext) {}

// ExitNumericLiteralPositive is called when production numericLiteralPositive is exited.
func (s *BaseSparqlParserListener) ExitNumericLiteralPositive(ctx *NumericLiteralPositiveContext) {}

// EnterNumericLiteralNegative is called when production numericLiteralNegative is entered.
func (s *BaseSparqlParserListener) EnterNumericLiteralNegative(ctx *NumericLiteralNegativeContext) {}

// ExitNumericLiteralNegative is called when production numericLiteralNegative is exited.
func (s *BaseSparqlParserListener) ExitNumericLiteralNegative(ctx *NumericLiteralNegativeContext) {}

// EnterBooleanLiteral is called when production booleanLiteral is entered.
func (s *BaseSparqlParserListener) EnterBooleanLiteral(ctx *BooleanLiteralContext) {}

// ExitBooleanLiteral is called when production booleanLiteral is exited.
func (s *BaseSparqlParserListener) ExitBooleanLiteral(ctx *BooleanLiteralContext) {}

// EnterString_ is called when production string_ is entered.
func (s *BaseSparqlParserListener) EnterString_(ctx *String_Context) {}

// ExitString_ is called when production string_ is exited.
func (s *BaseSparqlParserListener) ExitString_(ctx *String_Context) {}

// EnterIriRef is called when production iriRef is entered.
func (s *BaseSparqlParserListener) EnterIriRef(ctx *IriRefContext) {}

// ExitIriRef is called when production iriRef is exited.
func (s *BaseSparqlParserListener) ExitIriRef(ctx *IriRefContext) {}

// EnterPrefixedName is called when production prefixedName is entered.
func (s *BaseSparqlParserListener) EnterPrefixedName(ctx *PrefixedNameContext) {}

// ExitPrefixedName is called when production prefixedName is exited.
func (s *BaseSparqlParserListener) ExitPrefixedName(ctx *PrefixedNameContext) {}

// EnterBlankNode is called when production blankNode is entered.
func (s *BaseSparqlParserListener) EnterBlankNode(ctx *BlankNodeContext) {}

// ExitBlankNode is called when production blankNode is exited.
func (s *BaseSparqlParserListener) ExitBlankNode(ctx *BlankNodeContext) {}
