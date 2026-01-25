# HAVING Clause Parsing Issue

## Status

Two test cases are currently skipped due to an ANTLR parsing issue with HAVING clauses:

- `TestE2EParseAndTranslate/SELECT_with_HAVING_clause`
- `TestE2EParseAndTranslate/HAVING_with_complex_expression`

## Problem

When parsing SPARQL queries with HAVING clauses, the ANTLR parser's `GetText()` method returns
incomplete text for the HAVING expression. For example:

- Expected: `(COUNT(?o) > 5)`
- Actual: `COUNT(?o` (truncated)

## Root Cause

The issue appears to be in the ANTLR grammar or parser generation for `HavingClauseContext`. When
calling `GetText()` on the context, it doesn't properly capture the full constraint expression.

## Investigation Notes

1. The grammar defines: `havingClause : HAVING havingCondition+`
2. Each `havingCondition` is a `constraint`
3. A `constraint` can be `brackettedExpression`, `builtInCall`, or `functionCall`
4. The expression `(COUNT(?o) > 5)` should be parsed as a `brackettedExpression` containing a
   comparison
5. However, `GetText()` on the HavingConditionContext's Constraint returns only the function call
   part

## Possible Fixes

1. **Grammar Issue**: The SparqlParser.g4 grammar may need to be reviewed and regenerated
   - Check if HAVING grammar matches SPARQL 1.1 spec exactly
   - Regenerate the parser using the ANTLR 4.11.1 or 4.13.1 jar file

2. **Parser Generation**: The existing generated parser in `sparql/gen/` may be stale
   - Run `./generate_parser.sh` to regenerate the parser
   - Compare against the official SPARQL grammar from https://github.com/antlr/grammars-v4

3. **Visitor Implementation**: May need a different approach
   - Instead of relying on `GetText()`, manually build the expression by walking the parse tree
   - Extract individual parts of the HAVING expression and reconstruct it

## Next Steps

1. Review the SPARQL ANTLR grammar files in this directory
2. Verify they match the official spec or latest community version
3. Regenerate the parser and retest
4. If still broken, implement manual tree walking in the visitor
5. Re-enable the skipped tests once fixed

## References

- SPARQL 1.1 Query Language: https://www.w3.org/TR/sparql11-query/#rAggregate
- ANTLR SPARQL Grammar: https://github.com/antlr/grammars-v4/tree/master/sparql
- Current Grammar Files: SparqlLexer.g4, SparqlParser.g4
