package sparql

import (
	"strconv"
	"strings"

	"github.com/antlr4-go/antlr/v4"
	gen "github.com/dgraph-io/dgraph/v25/sparql/gen"
)

// SparqlVisitor walks the ANTLR parse tree and builds a SPARQLQuery.
// This is a simplified visitor that extracts the key SPARQL features
// (SELECT/ASK, WHERE clauses, aggregates, BIND, HAVING, etc.).
type SparqlVisitor struct {
	query *SPARQLQueryImpl
	err   error
}

// NewSparqlVisitor creates a visitor for building SPARQLQuery from a parse tree.
func NewSparqlVisitor() *SparqlVisitor {
	return &SparqlVisitor{
		query: &SPARQLQueryImpl{
			Prefixes:   make(map[string]string),
			Projs:      []string{},
			Patterns:   []GraphPattern{},
			Bgps:       []*BGP{},
			Aggregates: []*Aggregate{},
			Binds:      []*BindExpression{},
			From:       []string{},
			FromNamed:  []string{},
		},
	}
}

// GetQuery returns the built query and any error.
func (v *SparqlVisitor) GetQuery() (*SPARQLQueryImpl, error) {
	return v.query, v.err
}

// Visit dispatches to the correct visit method based on context type.
func (v *SparqlVisitor) Visit(tree antlr.ParseTree) interface{} {
	if tree == nil {
		return nil
	}
	return tree.Accept(v)
}

// VisitChildren visits all children of a rule context.
func (v *SparqlVisitor) VisitChildren(node antlr.RuleNode) interface{} {
	for i := 0; i < node.GetChildCount(); i++ {
		if child := node.GetChild(i); child != nil {
			child.(antlr.ParseTree).Accept(v)
		}
	}
	return nil
}

// VisitTerminal handles terminal nodes (tokens).
func (v *SparqlVisitor) VisitTerminal(node antlr.TerminalNode) interface{} {
	return node.GetText()
}

// VisitErrorNode handles error nodes.
func (v *SparqlVisitor) VisitErrorNode(node antlr.ErrorNode) interface{} {
	return nil
}

// VisitQuery processes the top-level query rule.
// This implements the SparqlParserVisitor interface method.
func (v *SparqlVisitor) VisitQuery(ctx *gen.QueryContext) interface{} {
	if ctx == nil {
		return nil
	}

	// Visit prologue for prefixes
	if ctx.Prologue() != nil {
		ctx.Prologue().Accept(v)
	}

	// Examine children to find the query type
	selectCtx := ctx.SelectQuery()
	askCtx := ctx.AskQuery()
	constructCtx := ctx.ConstructQuery()
	describeCtx := ctx.DescribeQuery()

	if selectCtx != nil {
		// Cast interface to concrete type
		if concreteCtx, ok := selectCtx.(*gen.SelectQueryContext); ok {
			v.VisitSelectQuery(concreteCtx)
		}
	} else if askCtx != nil {
		// Cast interface to concrete type
		if concreteCtx, ok := askCtx.(*gen.AskQueryContext); ok {
			v.VisitAskQuery(concreteCtx)
		}
	} else if constructCtx != nil {
		v.query.Qtype = "CONSTRUCT"
	} else if describeCtx != nil {
		v.query.Qtype = "DESCRIBE"
	}

	return v.query
}

// VisitPrologue processes prefix declarations.
func (v *SparqlVisitor) VisitPrologue(ctx *gen.PrologueContext) interface{} {
	if ctx == nil {
		return nil
	}

	// Process all prefix declarations
	prefixDeclCount := ctx.GetChildCount()
	for i := 0; i < prefixDeclCount; i++ {
		child := ctx.GetChild(i)
		if prefixDeclCtx, ok := child.(*gen.PrefixDeclContext); ok {
			v.VisitPrefixDecl(prefixDeclCtx)
		}
	}

	return nil
}

// VisitPrefixDecl processes a single prefix declaration.
func (v *SparqlVisitor) VisitPrefixDecl(ctx *gen.PrefixDeclContext) interface{} {
	if ctx == nil {
		return nil
	}

	// PREFIX pattern: PREFIX name: <IRI>
	// Find the PNAME_NS (prefix name with colon) and IRI_REF
	var prefixName, iriRef string

	for i := 0; i < ctx.GetChildCount(); i++ {
		child := ctx.GetChild(i)
		if terminal, ok := child.(antlr.TerminalNode); ok {
			text := terminal.GetText()
			if strings.HasSuffix(text, ":") {
				prefixName = strings.TrimSuffix(text, ":")
			}
		}
		if iriRefCtx, ok := child.(*gen.IriRefContext); ok {
			iriRef = extractIRI(iriRefCtx)
		}
	}

	if prefixName != "" && iriRef != "" {
		v.query.Prefixes[prefixName] = iriRef
	}

	return nil
}

// VisitSelectQuery implements SparqlParserVisitor interface for SELECT queries.
func (v *SparqlVisitor) VisitSelectQuery(ctx *gen.SelectQueryContext) interface{} {
	if ctx == nil {
		return nil
	}

	v.query.Qtype = "SELECT"

	if len(ctx.GetTokens(gen.SparqlParserDISTINCT)) > 0 {
		v.query.Distinct = true
	}

	for i := 0; ; i++ {
		proj := ctx.SelectProjection(i)
		if proj == nil {
			break
		}
		if concrete, ok := proj.(*gen.SelectProjectionContext); ok {
			v.VisitSelectProjection(concrete)
		}
	}

	if len(ctx.GetTokens(gen.SparqlParserSTAR)) > 0 {
		v.query.Projs = appendUnique(v.query.Projs, "*")
	}

	for i := 0; ; i++ {
		dc := ctx.DatasetClause(i)
		if dc == nil {
			break
		}
		if concrete, ok := dc.(*gen.DatasetClauseContext); ok {
			v.VisitDatasetClause(concrete)
		}
	}

	if wc := ctx.WhereClause(); wc != nil {
		wc.Accept(v)
	}
	if sm := ctx.SolutionModifier(); sm != nil {
		sm.Accept(v)
	}

	return nil
}

// VisitAskQuery implements SparqlParserVisitor interface for ASK queries.
func (v *SparqlVisitor) VisitAskQuery(ctx *gen.AskQueryContext) interface{} {
	if ctx == nil {
		return nil
	}

	v.query.Qtype = "ASK"

	if wc := ctx.WhereClause(); wc != nil {
		wc.Accept(v)
	}

	return nil
}

// VisitDatasetClause handles FROM/FROM NAMED.
func (v *SparqlVisitor) VisitDatasetClause(ctx *gen.DatasetClauseContext) interface{} {
	if ctx == nil {
		return nil
	}
	if def := ctx.DefaultGraphClause(); def != nil {
		if sel := def.SourceSelector(); sel != nil {
			iri := strings.Trim(sel.GetText(), "<>")
			if iri != "" {
				v.query.From = append(v.query.From, iri)
			}
		}
	}
	if named := ctx.NamedGraphClause(); named != nil {
		if sel := named.SourceSelector(); sel != nil {
			iri := strings.Trim(sel.GetText(), "<>")
			if iri != "" {
				v.query.FromNamed = append(v.query.FromNamed, iri)
			}
		}
	}
	return nil
}

func (v *SparqlVisitor) VisitDefaultGraphClause(ctx *gen.DefaultGraphClauseContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitNamedGraphClause(ctx *gen.NamedGraphClauseContext) interface{} {
	return nil
}

func (v *SparqlVisitor) VisitWhereClause(ctx *gen.WhereClauseContext) interface{} {
	if ctx == nil || ctx.GroupGraphPattern() == nil {
		return nil
	}
	if ggp, ok := ctx.GroupGraphPattern().(*gen.GroupGraphPatternContext); ok {
		patterns := v.collectPatternsFromGroup(ggp, "")
		if len(patterns) > 0 {
			v.query.Patterns = append(v.query.Patterns, patterns...)
		}
	}
	return nil
}

func (v *SparqlVisitor) VisitSolutionModifier(ctx *gen.SolutionModifierContext) interface{} {
	if ctx == nil {
		return nil
	}
	if hc := ctx.HavingClause(); hc != nil {
		v.VisitHavingClause(hc.(*gen.HavingClauseContext))
	}
	if oc := ctx.OrderClause(); oc != nil {
		v.VisitOrderClause(oc.(*gen.OrderClauseContext))
	}
	if loc := ctx.LimitOffsetClauses(); loc != nil {
		v.VisitLimitOffsetClauses(loc.(*gen.LimitOffsetClausesContext))
	}
	return nil
}

func (v *SparqlVisitor) VisitHavingClause(ctx *gen.HavingClauseContext) interface{} {
	if ctx == nil {
		return nil
	}

	// Get the full text of the HAVING clause, which includes the HAVING keyword
	text := ctx.GetText()
	if text == "" {
		return nil
	}

	// Remove the "HAVING" keyword at the beginning
	// The GetText() returns something like "HAVING(expression)" or "HAVING (expression)"
	text = strings.TrimSpace(text)

	// Case-insensitive removal of "HAVING"
	if len(text) > 6 && strings.ToUpper(text[:6]) == "HAVING" {
		text = strings.TrimSpace(text[6:])
	}

	if text == "" {
		return nil
	}

	v.query.Having = &HavingClause{Expression: text}
	return nil
}

func (v *SparqlVisitor) VisitLimitOffsetClauses(ctx *gen.LimitOffsetClausesContext) interface{} {
	if ctx == nil {
		return nil
	}
	if lc := ctx.LimitClause(); lc != nil {
		text := strings.TrimSpace(lc.GetText())
		text = strings.TrimPrefix(strings.ToUpper(text), "LIMIT")
		text = strings.TrimSpace(text)
		if n, err := strconv.Atoi(text); err == nil {
			v.query.Limit = n
		}
	}
	if oc := ctx.OffsetClause(); oc != nil {
		text := strings.TrimSpace(oc.GetText())
		text = strings.TrimPrefix(strings.ToUpper(text), "OFFSET")
		text = strings.TrimSpace(text)
		if n, err := strconv.Atoi(text); err == nil {
			v.query.Offset = n
		}
	}
	return nil
}

func (v *SparqlVisitor) VisitOrderClause(ctx *gen.OrderClauseContext) interface{} {
	if ctx == nil {
		return nil
	}
	for _, oc := range ctx.AllOrderCondition() {
		if c, ok := oc.(*gen.OrderConditionContext); ok {
			text := c.GetText()
			for _, part := range strings.FieldsFunc(text, func(r rune) bool { return r == '(' || r == ')' || r == ',' }) {
				part = strings.TrimSpace(part)
				if strings.HasPrefix(part, "?") || strings.HasPrefix(part, "$") {
					v.query.OrderBy = append(v.query.OrderBy, part)
					break
				}
			}
		}
	}
	return nil
}

func (v *SparqlVisitor) VisitGroupGraphPattern(ctx *gen.GroupGraphPatternContext) interface{} {
	patterns := v.collectPatternsFromGroup(ctx, "")
	if len(patterns) > 0 {
		v.query.Patterns = append(v.query.Patterns, patterns...)
	}
	return nil
}
func (v *SparqlVisitor) VisitTriplesBlock(ctx *gen.TriplesBlockContext) interface{} { return nil }
func (v *SparqlVisitor) VisitGraphPatternNotTriples(ctx *gen.GraphPatternNotTriplesContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitBind(ctx *gen.BindContext) interface{} {
	if ctx == nil {
		return nil
	}
	expr := ""
	if e := ctx.Expression(); e != nil {
		expr = e.GetText()
	}
	varName := ""
	if vctx := ctx.Var_(); vctx != nil {
		varName = vctx.GetText()
	}
	if varName != "" {
		v.query.Binds = append(v.query.Binds, &BindExpression{Expression: expr, Variable: varName})
	}
	return nil
}
func (v *SparqlVisitor) VisitOptionalGraphPattern(ctx *gen.OptionalGraphPatternContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitGraphGraphPattern(ctx *gen.GraphGraphPatternContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitGroupOrUnionGraphPattern(ctx *gen.GroupOrUnionGraphPatternContext) interface{} {
	return nil
}

// --- Helpers built from parse tree ---

// buildAggregateFromCtx converts an AggregateProjectionContext into Aggregate.
func (v *SparqlVisitor) buildAggregateFromCtx(ctx *gen.AggregateProjectionContext) *Aggregate {
	if ctx == nil || ctx.AggregateFunction() == nil || ctx.Var_() == nil {
		return nil
	}

	fnCtx, ok := ctx.AggregateFunction().(*gen.AggregateFunctionContext)
	if !ok {
		return nil
	}

	var funcName string
	switch {
	case fnCtx.COUNT() != nil:
		funcName = "count"
	case fnCtx.SUM() != nil:
		funcName = "sum"
	case fnCtx.AVG() != nil:
		funcName = "avg"
	case fnCtx.MIN() != nil:
		funcName = "min"
	case fnCtx.MAX() != nil:
		funcName = "max"
	default:
		return nil
	}

	variable := fnCtx.Var_().GetText()
	alias := ctx.Var_().GetText()
	distinct := fnCtx.DISTINCT() != nil
	return &Aggregate{Function: funcName, Variable: variable, Alias: alias, Distinct: distinct}
}

// collectPatternsFromGroup walks a groupGraphPattern and returns patterns (also populating BGPs).
func (v *SparqlVisitor) collectPatternsFromGroup(ctx *gen.GroupGraphPatternContext, graph string) []GraphPattern {
	var patterns []GraphPattern
	if ctx == nil {
		return patterns
	}

	// Triples blocks become BGPs
	for _, tb := range ctx.AllTriplesBlock() {
		if block, ok := tb.(*gen.TriplesBlockContext); ok {
			triples := v.collectTriplesBlock(block, graph)
			if len(triples) > 0 {
				bgp := &BGP{Triples: triples}
				patterns = append(patterns, bgp)
				v.query.Bgps = append(v.query.Bgps, bgp)
			}
		}
	}

	// Non-triple patterns
	for _, gpn := range ctx.AllGraphPatternNotTriples() {
		gpCtx, ok := gpn.(*gen.GraphPatternNotTriplesContext)
		if !ok {
			continue
		}

		if opt := gpCtx.OptionalGraphPattern(); opt != nil {
			if ogp, ok := opt.(*gen.OptionalGraphPatternContext); ok {
				sub := v.collectPatternsFromGroup(ogp.GroupGraphPattern().(*gen.GroupGraphPatternContext), graph)
				if len(sub) > 0 {
					patterns = append(patterns, &OptionalPattern{Patterns: sub})
				}
			}
			continue
		}

		if union := gpCtx.GroupOrUnionGraphPattern(); union != nil {
			if uctx, ok := union.(*gen.GroupOrUnionGraphPatternContext); ok {
				var alts [][]GraphPattern
				for _, grp := range uctx.AllGroupGraphPattern() {
					if g, ok := grp.(*gen.GroupGraphPatternContext); ok {
						alts = append(alts, v.collectPatternsFromGroup(g, graph))
					}
				}
				if len(alts) > 0 {
					patterns = append(patterns, &UnionPattern{Alternatives: alts})
				}
			}
			continue
		}

		if gg := gpCtx.GraphGraphPattern(); gg != nil {
			if gctx, ok := gg.(*gen.GraphGraphPatternContext); ok {
				gName := ""
				if vref := gctx.VarOrIRIref(); vref != nil {
					gName = strings.Trim(vref.GetText(), "<>")
				}
				sub := v.collectPatternsFromGroup(gctx.GroupGraphPattern().(*gen.GroupGraphPatternContext), gName)
				if len(sub) > 0 {
					patterns = append(patterns, sub...)
				}
			}
			continue
		}

		if bind := gpCtx.Bind(); bind != nil {
			if bctx, ok := bind.(*gen.BindContext); ok {
				v.VisitBind(bctx)
			}
		}

	}

	// Collect FILTER clauses attached at the group level
	for _, fil := range ctx.AllFilter_() {
		if fctx, ok := fil.(*gen.Filter_Context); ok {
			patterns = append(patterns, &FilterPattern{Expression: fctx.GetText()})
		}
	}

	return patterns
}

// collectTriplesBlock collects triples from a TriplesBlock, respecting nested blocks.
func (v *SparqlVisitor) collectTriplesBlock(ctx *gen.TriplesBlockContext, graph string) []*Triple {
	var triples []*Triple
	if ctx == nil {
		return triples
	}

	if ts := ctx.TriplesSameSubject(); ts != nil {
		if tctx, ok := ts.(*gen.TriplesSameSubjectContext); ok {
			triples = append(triples, v.collectTriplesSameSubject(tctx, graph)...)
		}
	}
	if next := ctx.TriplesBlock(); next != nil {
		if nctx, ok := next.(*gen.TriplesBlockContext); ok {
			triples = append(triples, v.collectTriplesBlock(nctx, graph)...)
		}
	}
	return triples
}

// collectTriplesSameSubject builds triples for one subject with property lists.
func (v *SparqlVisitor) collectTriplesSameSubject(ctx *gen.TriplesSameSubjectContext, graph string) []*Triple {
	var triples []*Triple
	if ctx == nil {
		return triples
	}

	subject := ""
	if vt := ctx.VarOrTerm(); vt != nil {
		subject = vt.GetText()
	} else if tn := ctx.TriplesNode(); tn != nil {
		subject = tn.GetText()
	}
	if subject == "" {
		return triples
	}

	if pln := ctx.PropertyListNotEmpty(); pln != nil {
		if pctx, ok := pln.(*gen.PropertyListNotEmptyContext); ok {
			triples = append(triples, v.collectTriplesFromPropertyList(pctx, subject, graph)...)
		}
	}

	if pl := ctx.PropertyList(); pl != nil {
		if pctx, ok := pl.(*gen.PropertyListContext); ok {
			if pne := pctx.PropertyListNotEmpty(); pne != nil {
				if pnctx, ok := pne.(*gen.PropertyListNotEmptyContext); ok {
					triples = append(triples, v.collectTriplesFromPropertyList(pnctx, subject, graph)...)
				}
			}
		}
	}

	return triples
}

// collectTriplesFromPropertyList expands verb/object lists to triples.
func (v *SparqlVisitor) collectTriplesFromPropertyList(ctx *gen.PropertyListNotEmptyContext, subject, graph string) []*Triple {
	var triples []*Triple
	if ctx == nil {
		return triples
	}

	verbs := ctx.AllVerb()
	objs := ctx.AllObjectList()
	for i := 0; i < len(verbs) && i < len(objs); i++ {
		verb := verbs[i].GetText()
		olist, ok := objs[i].(*gen.ObjectListContext)
		if !ok {
			continue
		}
		for _, obj := range olist.AllObject_() {
			if octx, ok := obj.(*gen.Object_Context); ok {
				oval := octx.GetText()
				triples = append(triples, &Triple{
					Subject:     subject,
					Predicate:   verb,
					Object:      oval,
					ObjectIsVar: strings.HasPrefix(oval, "?") || strings.HasPrefix(oval, "$"),
					Graph:       graph,
				})
			}
		}
	}
	return triples
}

// extractContextText recursively extracts all text from a context.
func (v *SparqlVisitor) extractContextText(ctx antlr.RuleContext) string {
	if ctx == nil {
		return ""
	}

	// Try to cast to concrete type to get start/stop
	if baseCtx, ok := ctx.(*antlr.BaseParserRuleContext); ok && baseCtx != nil {
		start := baseCtx.GetStart()
		stop := baseCtx.GetStop()

		if start == nil || stop == nil {
			return ""
		}

		input := start.GetInputStream()
		if input == nil {
			return ""
		}

		return input.GetTextFromInterval(antlr.NewInterval(start.GetTokenIndex(), stop.GetTokenIndex()))
	}

	return ""
}

// extractInterfaceText extracts text from any ANTLR rule context interface.
func (v *SparqlVisitor) extractInterfaceText(ctx antlr.RuleContext) string {
	if ctx == nil {
		return ""
	}

	// Try to cast to concrete type to get start/stop
	if baseCtx, ok := ctx.(*antlr.BaseParserRuleContext); ok && baseCtx != nil {
		start := baseCtx.GetStart()
		stop := baseCtx.GetStop()

		if start == nil || stop == nil {
			return ""
		}

		input := start.GetInputStream()
		if input == nil {
			return ""
		}

		return input.GetTextFromInterval(antlr.NewInterval(start.GetTokenIndex(), stop.GetTokenIndex()))
	}

	return ""
}

// Helper functions

func extractIRI(ctx *gen.IriRefContext) string {
	if ctx == nil {
		return ""
	}
	text := ctx.GetText()
	// Remove angle brackets
	return strings.TrimPrefix(strings.TrimSuffix(text, ">"), "<")
}

func extractIRIFromDatasetClause(text string, isNamed bool) string {
	// Extract IRI from text like "FROM <http://example.org>" or "FROM NAMED <http://example.org>"
	openIdx := strings.Index(text, "<")
	closeIdx := strings.LastIndex(text, ">")
	if openIdx != -1 && closeIdx != -1 && closeIdx > openIdx {
		return text[openIdx+1 : closeIdx]
	}
	return ""
}

// Stub methods to implement SparqlParserVisitor interface (or simple projections)

func (v *SparqlVisitor) VisitBaseDecl(ctx *gen.BaseDeclContext) interface{} { return nil }

func (v *SparqlVisitor) VisitSelectProjection(ctx *gen.SelectProjectionContext) interface{} {
	if ctx == nil {
		return nil
	}
	if agg := ctx.AggregateProjection(); agg != nil {
		if a, ok := agg.(*gen.AggregateProjectionContext); ok {
			if built := v.buildAggregateFromCtx(a); built != nil {
				v.query.Aggregates = append(v.query.Aggregates, built)
				if built.Alias != "" {
					v.query.Projs = appendUnique(v.query.Projs, built.Alias)
				}
			}
		}
		return nil
	}

	if vctx := ctx.Var_(); vctx != nil {
		v.query.Projs = appendUnique(v.query.Projs, vctx.GetText())
	}
	return nil
}

func (v *SparqlVisitor) VisitAggregateProjection(ctx *gen.AggregateProjectionContext) interface{} {
	if ctx == nil {
		return nil
	}
	if built := v.buildAggregateFromCtx(ctx); built != nil {
		v.query.Aggregates = append(v.query.Aggregates, built)
		if built.Alias != "" {
			v.query.Projs = appendUnique(v.query.Projs, built.Alias)
		}
	}
	return nil
}

func (v *SparqlVisitor) VisitAggregateFunction(ctx *gen.AggregateFunctionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitConstructQuery(ctx *gen.ConstructQueryContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitDescribeQuery(ctx *gen.DescribeQueryContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitSourceSelector(ctx *gen.SourceSelectorContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitGroupClause(ctx *gen.GroupClauseContext) interface{}         { return nil }
func (v *SparqlVisitor) VisitGroupCondition(ctx *gen.GroupConditionContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitHavingCondition(ctx *gen.HavingConditionContext) interface{} { return nil }
func (v *SparqlVisitor) VisitOrderCondition(ctx *gen.OrderConditionContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitLimitClause(ctx *gen.LimitClauseContext) interface{}         { return nil }
func (v *SparqlVisitor) VisitOffsetClause(ctx *gen.OffsetClauseContext) interface{}       { return nil }
func (v *SparqlVisitor) VisitFilter_(ctx *gen.Filter_Context) interface{}                 { return nil }
func (v *SparqlVisitor) VisitConstraint(ctx *gen.ConstraintContext) interface{}           { return nil }
func (v *SparqlVisitor) VisitFunctionCall(ctx *gen.FunctionCallContext) interface{}       { return nil }
func (v *SparqlVisitor) VisitArgList(ctx *gen.ArgListContext) interface{}                 { return nil }
func (v *SparqlVisitor) VisitConstructTemplate(ctx *gen.ConstructTemplateContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitConstructTriples(ctx *gen.ConstructTriplesContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitTriplesSameSubject(ctx *gen.TriplesSameSubjectContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitPropertyListNotEmpty(ctx *gen.PropertyListNotEmptyContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitPropertyList(ctx *gen.PropertyListContext) interface{} { return nil }
func (v *SparqlVisitor) VisitObjectList(ctx *gen.ObjectListContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitObject_(ctx *gen.Object_Context) interface{}           { return nil }
func (v *SparqlVisitor) VisitVerb(ctx *gen.VerbContext) interface{}                 { return nil }
func (v *SparqlVisitor) VisitTriplesNode(ctx *gen.TriplesNodeContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitBlankNodePropertyList(ctx *gen.BlankNodePropertyListContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitCollection(ctx *gen.CollectionContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitGraphNode(ctx *gen.GraphNodeContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitVarOrTerm(ctx *gen.VarOrTermContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitVarOrIRIref(ctx *gen.VarOrIRIrefContext) interface{} { return nil }
func (v *SparqlVisitor) VisitVar_(ctx *gen.Var_Context) interface{}               { return nil }
func (v *SparqlVisitor) VisitGraphTerm(ctx *gen.GraphTermContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitExpression(ctx *gen.ExpressionContext) interface{}   { return nil }
func (v *SparqlVisitor) VisitConditionalOrExpression(ctx *gen.ConditionalOrExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitConditionalAndExpression(ctx *gen.ConditionalAndExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitValueLogical(ctx *gen.ValueLogicalContext) interface{} { return nil }
func (v *SparqlVisitor) VisitRelationalExpression(ctx *gen.RelationalExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitNumericExpression(ctx *gen.NumericExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitAdditiveExpression(ctx *gen.AdditiveExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitMultiplicativeExpression(ctx *gen.MultiplicativeExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitUnaryExpression(ctx *gen.UnaryExpressionContext) interface{} { return nil }
func (v *SparqlVisitor) VisitPrimaryExpression(ctx *gen.PrimaryExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitBrackettedExpression(ctx *gen.BrackettedExpressionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitBuiltInCall(ctx *gen.BuiltInCallContext) interface{}         { return nil }
func (v *SparqlVisitor) VisitRegexExpression(ctx *gen.RegexExpressionContext) interface{} { return nil }
func (v *SparqlVisitor) VisitIriRefOrFunction(ctx *gen.IriRefOrFunctionContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitRdfLiteral(ctx *gen.RdfLiteralContext) interface{}         { return nil }
func (v *SparqlVisitor) VisitNumericLiteral(ctx *gen.NumericLiteralContext) interface{} { return nil }
func (v *SparqlVisitor) VisitNumericLiteralUnsigned(ctx *gen.NumericLiteralUnsignedContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitNumericLiteralPositive(ctx *gen.NumericLiteralPositiveContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitNumericLiteralNegative(ctx *gen.NumericLiteralNegativeContext) interface{} {
	return nil
}
func (v *SparqlVisitor) VisitBooleanLiteral(ctx *gen.BooleanLiteralContext) interface{} { return nil }
func (v *SparqlVisitor) VisitString_(ctx *gen.String_Context) interface{}               { return nil }
func (v *SparqlVisitor) VisitIriRef(ctx *gen.IriRefContext) interface{}                 { return nil }
func (v *SparqlVisitor) VisitPrefixedName(ctx *gen.PrefixedNameContext) interface{}     { return nil }
func (v *SparqlVisitor) VisitBlankNode(ctx *gen.BlankNodeContext) interface{}           { return nil }
