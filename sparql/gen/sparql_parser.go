// Code generated from SparqlParser.g4 by ANTLR 4.13.1. DO NOT EDIT.

package parser // SparqlParser

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/antlr4-go/antlr/v4"
)

// Suppress unused import errors
var _ = fmt.Printf
var _ = strconv.Itoa
var _ = sync.Once{}

type SparqlParser struct {
	*antlr.BaseParser
}

var SparqlParserParserStaticData struct {
	once                   sync.Once
	serializedATN          []int32
	LiteralNames           []string
	SymbolicNames          []string
	RuleNames              []string
	PredictionContextCache *antlr.PredictionContextCache
	atn                    *antlr.ATN
	decisionToDFA          []*antlr.DFA
}

func sparqlparserParserInit() {
	staticData := &SparqlParserParserStaticData
	staticData.LiteralNames = []string{
		"", "", "'ASC'", "'ASK'", "'BASE'", "'BIND'", "'AS'", "'GROUP'", "'HAVING'",
		"'COUNT'", "'SUM'", "'AVG'", "'MIN'", "'MAX'", "'BOUND'", "'BY'", "'CONSTRUCT'",
		"'DATATYPE'", "'DESC'", "'DESCRIBE'", "'DISTINCT'", "'FILTER'", "'FROM'",
		"'GRAPH'", "'LANG'", "'LANGMATCHES'", "'LIMIT'", "'NAMED'", "'OFFSET'",
		"'OPTIONAL'", "'ORDER'", "'PREFIX'", "'REDUCED'", "'REGEX'", "'SELECT'",
		"'STR'", "'UNION'", "'WHERE'", "'true'", "'false'", "'isLITERAL'", "'isBLANK'",
		"'isURI'", "'isIRI'", "'sameTerm'", "','", "'.'", "'&&'", "'||'", "'^^'",
		"'='", "'!'", "'>'", "'>='", "'<'", "'<='", "'{'", "'('", "'['", "'-'",
		"'!='", "'+'", "'}'", "')'", "']'", "';'", "'/'", "'*'",
	}
	staticData.SymbolicNames = []string{
		"", "A", "ASC", "ASK", "BASE", "BIND", "AS", "GROUP", "HAVING", "COUNT",
		"SUM", "AVG", "MIN", "MAX", "BOUND", "BY", "CONSTRUCT", "DATATYPE",
		"DESC", "DESCRIBE", "DISTINCT", "FILTER", "FROM", "GRAPH", "LANG", "LANGMATCHES",
		"LIMIT", "NAMED", "OFFSET", "OPTIONAL", "ORDER", "PREFIX", "REDUCED",
		"REGEX", "SELECT", "STR", "UNION", "WHERE", "TRUE", "FALSE", "IS_LITERAL",
		"IS_BLANK", "IS_URI", "IS_IRI", "SAME_TERM", "COMMA", "DOT", "DOUBLE_AMP",
		"DOUBLE_BAR", "DOUBLE_CARET", "EQUAL", "EXCLAMATION", "GREATER", "GREATER_OR_EQUAL",
		"LESS", "LESS_OR_EQUAL", "L_CURLY", "L_PAREN", "L_SQUARE", "MINUS",
		"NOT_EQUAL", "PLUS", "R_CURLY", "R_PAREN", "R_SQUARE", "SEMICOLON",
		"SLASH", "STAR", "IRI_REF", "PNAME_NS", "PNAME_LN", "BLANK_NODE_LABEL",
		"VAR1", "VAR2", "LANGTAG", "INTEGER", "DECIMAL", "DOUBLE", "INTEGER_POSITIVE",
		"DECIMAL_POSITIVE", "DOUBLE_POSITIVE", "INTEGER_NEGATIVE", "DECIMAL_NEGATIVE",
		"DOUBLE_NEGATIVE", "EXPONENT", "STRING_LITERAL1", "STRING_LITERAL2",
		"STRING_LITERAL_LONG1", "STRING_LITERAL_LONG2", "ECHAR", "NIL", "ANON",
		"PN_CHARS_U", "VARNAME", "PN_PREFIX", "PN_LOCAL", "WS", "COMMENT",
	}
	staticData.RuleNames = []string{
		"query", "prologue", "baseDecl", "prefixDecl", "selectQuery", "selectProjection",
		"aggregateProjection", "aggregateFunction", "constructQuery", "describeQuery",
		"askQuery", "datasetClause", "defaultGraphClause", "namedGraphClause",
		"sourceSelector", "whereClause", "solutionModifier", "groupClause",
		"groupCondition", "havingClause", "havingCondition", "limitOffsetClauses",
		"orderClause", "orderCondition", "limitClause", "offsetClause", "groupGraphPattern",
		"triplesBlock", "graphPatternNotTriples", "bind", "optionalGraphPattern",
		"graphGraphPattern", "groupOrUnionGraphPattern", "filter_", "constraint",
		"functionCall", "argList", "constructTemplate", "constructTriples",
		"triplesSameSubject", "propertyListNotEmpty", "propertyList", "objectList",
		"object_", "verb", "triplesNode", "blankNodePropertyList", "collection",
		"graphNode", "varOrTerm", "varOrIRIref", "var_", "graphTerm", "expression",
		"conditionalOrExpression", "conditionalAndExpression", "valueLogical",
		"relationalExpression", "numericExpression", "additiveExpression", "multiplicativeExpression",
		"unaryExpression", "primaryExpression", "brackettedExpression", "builtInCall",
		"regexExpression", "iriRefOrFunction", "rdfLiteral", "numericLiteral",
		"numericLiteralUnsigned", "numericLiteralPositive", "numericLiteralNegative",
		"booleanLiteral", "string_", "iriRef", "prefixedName", "blankNode",
	}
	staticData.PredictionContextCache = antlr.NewPredictionContextCache()
	staticData.serializedATN = []int32{
		4, 1, 97, 704, 2, 0, 7, 0, 2, 1, 7, 1, 2, 2, 7, 2, 2, 3, 7, 3, 2, 4, 7,
		4, 2, 5, 7, 5, 2, 6, 7, 6, 2, 7, 7, 7, 2, 8, 7, 8, 2, 9, 7, 9, 2, 10, 7,
		10, 2, 11, 7, 11, 2, 12, 7, 12, 2, 13, 7, 13, 2, 14, 7, 14, 2, 15, 7, 15,
		2, 16, 7, 16, 2, 17, 7, 17, 2, 18, 7, 18, 2, 19, 7, 19, 2, 20, 7, 20, 2,
		21, 7, 21, 2, 22, 7, 22, 2, 23, 7, 23, 2, 24, 7, 24, 2, 25, 7, 25, 2, 26,
		7, 26, 2, 27, 7, 27, 2, 28, 7, 28, 2, 29, 7, 29, 2, 30, 7, 30, 2, 31, 7,
		31, 2, 32, 7, 32, 2, 33, 7, 33, 2, 34, 7, 34, 2, 35, 7, 35, 2, 36, 7, 36,
		2, 37, 7, 37, 2, 38, 7, 38, 2, 39, 7, 39, 2, 40, 7, 40, 2, 41, 7, 41, 2,
		42, 7, 42, 2, 43, 7, 43, 2, 44, 7, 44, 2, 45, 7, 45, 2, 46, 7, 46, 2, 47,
		7, 47, 2, 48, 7, 48, 2, 49, 7, 49, 2, 50, 7, 50, 2, 51, 7, 51, 2, 52, 7,
		52, 2, 53, 7, 53, 2, 54, 7, 54, 2, 55, 7, 55, 2, 56, 7, 56, 2, 57, 7, 57,
		2, 58, 7, 58, 2, 59, 7, 59, 2, 60, 7, 60, 2, 61, 7, 61, 2, 62, 7, 62, 2,
		63, 7, 63, 2, 64, 7, 64, 2, 65, 7, 65, 2, 66, 7, 66, 2, 67, 7, 67, 2, 68,
		7, 68, 2, 69, 7, 69, 2, 70, 7, 70, 2, 71, 7, 71, 2, 72, 7, 72, 2, 73, 7,
		73, 2, 74, 7, 74, 2, 75, 7, 75, 2, 76, 7, 76, 1, 0, 1, 0, 1, 0, 1, 0, 1,
		0, 3, 0, 160, 8, 0, 1, 0, 1, 0, 1, 1, 3, 1, 165, 8, 1, 1, 1, 5, 1, 168,
		8, 1, 10, 1, 12, 1, 171, 9, 1, 1, 2, 1, 2, 1, 2, 1, 3, 1, 3, 1, 3, 1, 3,
		1, 4, 1, 4, 3, 4, 182, 8, 4, 1, 4, 4, 4, 185, 8, 4, 11, 4, 12, 4, 186,
		1, 4, 3, 4, 190, 8, 4, 1, 4, 5, 4, 193, 8, 4, 10, 4, 12, 4, 196, 9, 4,
		1, 4, 1, 4, 1, 4, 1, 5, 1, 5, 3, 5, 203, 8, 5, 1, 6, 1, 6, 1, 6, 1, 6,
		1, 6, 1, 6, 1, 7, 1, 7, 1, 7, 3, 7, 214, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7,
		1, 7, 1, 7, 3, 7, 222, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7,
		230, 8, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 238, 8, 7, 1, 7, 1,
		7, 1, 7, 1, 7, 1, 7, 1, 7, 3, 7, 246, 8, 7, 1, 7, 1, 7, 1, 7, 3, 7, 251,
		8, 7, 1, 8, 1, 8, 1, 8, 5, 8, 256, 8, 8, 10, 8, 12, 8, 259, 9, 8, 1, 8,
		1, 8, 1, 8, 1, 9, 1, 9, 4, 9, 266, 8, 9, 11, 9, 12, 9, 267, 1, 9, 3, 9,
		271, 8, 9, 1, 9, 5, 9, 274, 8, 9, 10, 9, 12, 9, 277, 9, 9, 1, 9, 3, 9,
		280, 8, 9, 1, 9, 1, 9, 1, 10, 1, 10, 5, 10, 286, 8, 10, 10, 10, 12, 10,
		289, 9, 10, 1, 10, 1, 10, 1, 11, 1, 11, 1, 11, 3, 11, 296, 8, 11, 1, 12,
		1, 12, 1, 13, 1, 13, 1, 13, 1, 14, 1, 14, 1, 15, 3, 15, 306, 8, 15, 1,
		15, 1, 15, 1, 16, 3, 16, 311, 8, 16, 1, 16, 3, 16, 314, 8, 16, 1, 16, 3,
		16, 317, 8, 16, 1, 16, 3, 16, 320, 8, 16, 1, 17, 1, 17, 1, 17, 4, 17, 325,
		8, 17, 11, 17, 12, 17, 326, 1, 18, 1, 18, 3, 18, 331, 8, 18, 1, 19, 1,
		19, 4, 19, 335, 8, 19, 11, 19, 12, 19, 336, 1, 20, 1, 20, 1, 21, 1, 21,
		3, 21, 343, 8, 21, 1, 21, 1, 21, 3, 21, 347, 8, 21, 3, 21, 349, 8, 21,
		1, 22, 1, 22, 1, 22, 4, 22, 354, 8, 22, 11, 22, 12, 22, 355, 1, 23, 1,
		23, 1, 23, 1, 23, 3, 23, 362, 8, 23, 1, 24, 1, 24, 1, 24, 1, 25, 1, 25,
		1, 25, 1, 26, 1, 26, 3, 26, 372, 8, 26, 1, 26, 1, 26, 3, 26, 376, 8, 26,
		1, 26, 3, 26, 379, 8, 26, 1, 26, 3, 26, 382, 8, 26, 5, 26, 384, 8, 26,
		10, 26, 12, 26, 387, 9, 26, 1, 26, 1, 26, 1, 27, 1, 27, 1, 27, 3, 27, 394,
		8, 27, 3, 27, 396, 8, 27, 1, 28, 1, 28, 1, 28, 1, 28, 3, 28, 402, 8, 28,
		1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 29, 1, 30, 1, 30, 1, 30, 1,
		31, 1, 31, 1, 31, 1, 31, 1, 32, 1, 32, 1, 32, 5, 32, 421, 8, 32, 10, 32,
		12, 32, 424, 9, 32, 1, 33, 1, 33, 1, 33, 1, 34, 1, 34, 1, 34, 3, 34, 432,
		8, 34, 1, 35, 1, 35, 1, 35, 1, 36, 1, 36, 1, 36, 1, 36, 1, 36, 5, 36, 442,
		8, 36, 10, 36, 12, 36, 445, 9, 36, 1, 36, 1, 36, 3, 36, 449, 8, 36, 1,
		37, 1, 37, 3, 37, 453, 8, 37, 1, 37, 1, 37, 1, 38, 1, 38, 1, 38, 3, 38,
		460, 8, 38, 3, 38, 462, 8, 38, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39, 1, 39,
		3, 39, 470, 8, 39, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 1, 40, 3, 40, 478,
		8, 40, 5, 40, 480, 8, 40, 10, 40, 12, 40, 483, 9, 40, 1, 41, 3, 41, 486,
		8, 41, 1, 42, 1, 42, 1, 42, 5, 42, 491, 8, 42, 10, 42, 12, 42, 494, 9,
		42, 1, 43, 1, 43, 1, 44, 1, 44, 3, 44, 500, 8, 44, 1, 45, 1, 45, 3, 45,
		504, 8, 45, 1, 46, 1, 46, 1, 46, 1, 46, 1, 47, 1, 47, 4, 47, 512, 8, 47,
		11, 47, 12, 47, 513, 1, 47, 1, 47, 1, 48, 1, 48, 3, 48, 520, 8, 48, 1,
		49, 1, 49, 3, 49, 524, 8, 49, 1, 50, 1, 50, 3, 50, 528, 8, 50, 1, 51, 1,
		51, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 1, 52, 3, 52, 538, 8, 52, 1, 53,
		1, 53, 1, 54, 1, 54, 1, 54, 5, 54, 545, 8, 54, 10, 54, 12, 54, 548, 9,
		54, 1, 55, 1, 55, 1, 55, 5, 55, 553, 8, 55, 10, 55, 12, 55, 556, 9, 55,
		1, 56, 1, 56, 1, 57, 1, 57, 1, 57, 3, 57, 563, 8, 57, 1, 58, 1, 58, 1,
		59, 1, 59, 1, 59, 1, 59, 1, 59, 5, 59, 572, 8, 59, 10, 59, 12, 59, 575,
		9, 59, 1, 60, 1, 60, 1, 60, 5, 60, 580, 8, 60, 10, 60, 12, 60, 583, 9,
		60, 1, 61, 3, 61, 586, 8, 61, 1, 61, 1, 61, 1, 62, 1, 62, 1, 62, 1, 62,
		1, 62, 1, 62, 1, 62, 3, 62, 597, 8, 62, 1, 63, 1, 63, 1, 63, 1, 63, 1,
		64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64,
		1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1,
		64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64,
		1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1,
		64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64, 1, 64,
		1, 64, 1, 64, 3, 64, 658, 8, 64, 1, 65, 1, 65, 1, 65, 1, 65, 1, 65, 1,
		65, 1, 65, 3, 65, 667, 8, 65, 1, 65, 1, 65, 1, 66, 1, 66, 3, 66, 673, 8,
		66, 1, 67, 1, 67, 1, 67, 1, 67, 3, 67, 679, 8, 67, 1, 68, 1, 68, 1, 68,
		3, 68, 684, 8, 68, 1, 69, 1, 69, 1, 70, 1, 70, 1, 71, 1, 71, 1, 72, 1,
		72, 1, 73, 1, 73, 1, 74, 1, 74, 3, 74, 698, 8, 74, 1, 75, 1, 75, 1, 76,
		1, 76, 1, 76, 0, 0, 77, 0, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24,
		26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46, 48, 50, 52, 54, 56, 58, 60,
		62, 64, 66, 68, 70, 72, 74, 76, 78, 80, 82, 84, 86, 88, 90, 92, 94, 96,
		98, 100, 102, 104, 106, 108, 110, 112, 114, 116, 118, 120, 122, 124, 126,
		128, 130, 132, 134, 136, 138, 140, 142, 144, 146, 148, 150, 152, 0, 14,
		2, 0, 20, 20, 32, 32, 2, 0, 2, 2, 18, 18, 1, 0, 72, 73, 3, 0, 50, 50, 52,
		55, 60, 60, 2, 0, 59, 59, 61, 61, 1, 0, 66, 67, 3, 0, 51, 51, 59, 59, 61,
		61, 1, 0, 75, 77, 1, 0, 78, 80, 1, 0, 81, 83, 1, 0, 38, 39, 1, 0, 85, 86,
		1, 0, 69, 70, 2, 0, 71, 71, 91, 91, 731, 0, 154, 1, 0, 0, 0, 2, 164, 1,
		0, 0, 0, 4, 172, 1, 0, 0, 0, 6, 175, 1, 0, 0, 0, 8, 179, 1, 0, 0, 0, 10,
		202, 1, 0, 0, 0, 12, 204, 1, 0, 0, 0, 14, 250, 1, 0, 0, 0, 16, 252, 1,
		0, 0, 0, 18, 263, 1, 0, 0, 0, 20, 283, 1, 0, 0, 0, 22, 292, 1, 0, 0, 0,
		24, 297, 1, 0, 0, 0, 26, 299, 1, 0, 0, 0, 28, 302, 1, 0, 0, 0, 30, 305,
		1, 0, 0, 0, 32, 310, 1, 0, 0, 0, 34, 321, 1, 0, 0, 0, 36, 330, 1, 0, 0,
		0, 38, 332, 1, 0, 0, 0, 40, 338, 1, 0, 0, 0, 42, 348, 1, 0, 0, 0, 44, 350,
		1, 0, 0, 0, 46, 361, 1, 0, 0, 0, 48, 363, 1, 0, 0, 0, 50, 366, 1, 0, 0,
		0, 52, 369, 1, 0, 0, 0, 54, 390, 1, 0, 0, 0, 56, 401, 1, 0, 0, 0, 58, 403,
		1, 0, 0, 0, 60, 410, 1, 0, 0, 0, 62, 413, 1, 0, 0, 0, 64, 417, 1, 0, 0,
		0, 66, 425, 1, 0, 0, 0, 68, 431, 1, 0, 0, 0, 70, 433, 1, 0, 0, 0, 72, 448,
		1, 0, 0, 0, 74, 450, 1, 0, 0, 0, 76, 456, 1, 0, 0, 0, 78, 469, 1, 0, 0,
		0, 80, 471, 1, 0, 0, 0, 82, 485, 1, 0, 0, 0, 84, 487, 1, 0, 0, 0, 86, 495,
		1, 0, 0, 0, 88, 499, 1, 0, 0, 0, 90, 503, 1, 0, 0, 0, 92, 505, 1, 0, 0,
		0, 94, 509, 1, 0, 0, 0, 96, 519, 1, 0, 0, 0, 98, 523, 1, 0, 0, 0, 100,
		527, 1, 0, 0, 0, 102, 529, 1, 0, 0, 0, 104, 537, 1, 0, 0, 0, 106, 539,
		1, 0, 0, 0, 108, 541, 1, 0, 0, 0, 110, 549, 1, 0, 0, 0, 112, 557, 1, 0,
		0, 0, 114, 559, 1, 0, 0, 0, 116, 564, 1, 0, 0, 0, 118, 566, 1, 0, 0, 0,
		120, 576, 1, 0, 0, 0, 122, 585, 1, 0, 0, 0, 124, 596, 1, 0, 0, 0, 126,
		598, 1, 0, 0, 0, 128, 657, 1, 0, 0, 0, 130, 659, 1, 0, 0, 0, 132, 670,
		1, 0, 0, 0, 134, 674, 1, 0, 0, 0, 136, 683, 1, 0, 0, 0, 138, 685, 1, 0,
		0, 0, 140, 687, 1, 0, 0, 0, 142, 689, 1, 0, 0, 0, 144, 691, 1, 0, 0, 0,
		146, 693, 1, 0, 0, 0, 148, 697, 1, 0, 0, 0, 150, 699, 1, 0, 0, 0, 152,
		701, 1, 0, 0, 0, 154, 159, 3, 2, 1, 0, 155, 160, 3, 8, 4, 0, 156, 160,
		3, 16, 8, 0, 157, 160, 3, 18, 9, 0, 158, 160, 3, 20, 10, 0, 159, 155, 1,
		0, 0, 0, 159, 156, 1, 0, 0, 0, 159, 157, 1, 0, 0, 0, 159, 158, 1, 0, 0,
		0, 160, 161, 1, 0, 0, 0, 161, 162, 5, 0, 0, 1, 162, 1, 1, 0, 0, 0, 163,
		165, 3, 4, 2, 0, 164, 163, 1, 0, 0, 0, 164, 165, 1, 0, 0, 0, 165, 169,
		1, 0, 0, 0, 166, 168, 3, 6, 3, 0, 167, 166, 1, 0, 0, 0, 168, 171, 1, 0,
		0, 0, 169, 167, 1, 0, 0, 0, 169, 170, 1, 0, 0, 0, 170, 3, 1, 0, 0, 0, 171,
		169, 1, 0, 0, 0, 172, 173, 5, 4, 0, 0, 173, 174, 5, 68, 0, 0, 174, 5, 1,
		0, 0, 0, 175, 176, 5, 31, 0, 0, 176, 177, 5, 69, 0, 0, 177, 178, 5, 68,
		0, 0, 178, 7, 1, 0, 0, 0, 179, 181, 5, 34, 0, 0, 180, 182, 7, 0, 0, 0,
		181, 180, 1, 0, 0, 0, 181, 182, 1, 0, 0, 0, 182, 189, 1, 0, 0, 0, 183,
		185, 3, 10, 5, 0, 184, 183, 1, 0, 0, 0, 185, 186, 1, 0, 0, 0, 186, 184,
		1, 0, 0, 0, 186, 187, 1, 0, 0, 0, 187, 190, 1, 0, 0, 0, 188, 190, 5, 67,
		0, 0, 189, 184, 1, 0, 0, 0, 189, 188, 1, 0, 0, 0, 190, 194, 1, 0, 0, 0,
		191, 193, 3, 22, 11, 0, 192, 191, 1, 0, 0, 0, 193, 196, 1, 0, 0, 0, 194,
		192, 1, 0, 0, 0, 194, 195, 1, 0, 0, 0, 195, 197, 1, 0, 0, 0, 196, 194,
		1, 0, 0, 0, 197, 198, 3, 30, 15, 0, 198, 199, 3, 32, 16, 0, 199, 9, 1,
		0, 0, 0, 200, 203, 3, 102, 51, 0, 201, 203, 3, 12, 6, 0, 202, 200, 1, 0,
		0, 0, 202, 201, 1, 0, 0, 0, 203, 11, 1, 0, 0, 0, 204, 205, 5, 57, 0, 0,
		205, 206, 3, 14, 7, 0, 206, 207, 5, 6, 0, 0, 207, 208, 3, 102, 51, 0, 208,
		209, 5, 63, 0, 0, 209, 13, 1, 0, 0, 0, 210, 211, 5, 9, 0, 0, 211, 213,
		5, 57, 0, 0, 212, 214, 5, 20, 0, 0, 213, 212, 1, 0, 0, 0, 213, 214, 1,
		0, 0, 0, 214, 215, 1, 0, 0, 0, 215, 216, 3, 102, 51, 0, 216, 217, 5, 63,
		0, 0, 217, 251, 1, 0, 0, 0, 218, 219, 5, 10, 0, 0, 219, 221, 5, 57, 0,
		0, 220, 222, 5, 20, 0, 0, 221, 220, 1, 0, 0, 0, 221, 222, 1, 0, 0, 0, 222,
		223, 1, 0, 0, 0, 223, 224, 3, 102, 51, 0, 224, 225, 5, 63, 0, 0, 225, 251,
		1, 0, 0, 0, 226, 227, 5, 11, 0, 0, 227, 229, 5, 57, 0, 0, 228, 230, 5,
		20, 0, 0, 229, 228, 1, 0, 0, 0, 229, 230, 1, 0, 0, 0, 230, 231, 1, 0, 0,
		0, 231, 232, 3, 102, 51, 0, 232, 233, 5, 63, 0, 0, 233, 251, 1, 0, 0, 0,
		234, 235, 5, 12, 0, 0, 235, 237, 5, 57, 0, 0, 236, 238, 5, 20, 0, 0, 237,
		236, 1, 0, 0, 0, 237, 238, 1, 0, 0, 0, 238, 239, 1, 0, 0, 0, 239, 240,
		3, 102, 51, 0, 240, 241, 5, 63, 0, 0, 241, 251, 1, 0, 0, 0, 242, 243, 5,
		13, 0, 0, 243, 245, 5, 57, 0, 0, 244, 246, 5, 20, 0, 0, 245, 244, 1, 0,
		0, 0, 245, 246, 1, 0, 0, 0, 246, 247, 1, 0, 0, 0, 247, 248, 3, 102, 51,
		0, 248, 249, 5, 63, 0, 0, 249, 251, 1, 0, 0, 0, 250, 210, 1, 0, 0, 0, 250,
		218, 1, 0, 0, 0, 250, 226, 1, 0, 0, 0, 250, 234, 1, 0, 0, 0, 250, 242,
		1, 0, 0, 0, 251, 15, 1, 0, 0, 0, 252, 253, 5, 16, 0, 0, 253, 257, 3, 74,
		37, 0, 254, 256, 3, 22, 11, 0, 255, 254, 1, 0, 0, 0, 256, 259, 1, 0, 0,
		0, 257, 255, 1, 0, 0, 0, 257, 258, 1, 0, 0, 0, 258, 260, 1, 0, 0, 0, 259,
		257, 1, 0, 0, 0, 260, 261, 3, 30, 15, 0, 261, 262, 3, 32, 16, 0, 262, 17,
		1, 0, 0, 0, 263, 270, 5, 19, 0, 0, 264, 266, 3, 100, 50, 0, 265, 264, 1,
		0, 0, 0, 266, 267, 1, 0, 0, 0, 267, 265, 1, 0, 0, 0, 267, 268, 1, 0, 0,
		0, 268, 271, 1, 0, 0, 0, 269, 271, 5, 67, 0, 0, 270, 265, 1, 0, 0, 0, 270,
		269, 1, 0, 0, 0, 271, 275, 1, 0, 0, 0, 272, 274, 3, 22, 11, 0, 273, 272,
		1, 0, 0, 0, 274, 277, 1, 0, 0, 0, 275, 273, 1, 0, 0, 0, 275, 276, 1, 0,
		0, 0, 276, 279, 1, 0, 0, 0, 277, 275, 1, 0, 0, 0, 278, 280, 3, 30, 15,
		0, 279, 278, 1, 0, 0, 0, 279, 280, 1, 0, 0, 0, 280, 281, 1, 0, 0, 0, 281,
		282, 3, 32, 16, 0, 282, 19, 1, 0, 0, 0, 283, 287, 5, 3, 0, 0, 284, 286,
		3, 22, 11, 0, 285, 284, 1, 0, 0, 0, 286, 289, 1, 0, 0, 0, 287, 285, 1,
		0, 0, 0, 287, 288, 1, 0, 0, 0, 288, 290, 1, 0, 0, 0, 289, 287, 1, 0, 0,
		0, 290, 291, 3, 30, 15, 0, 291, 21, 1, 0, 0, 0, 292, 295, 5, 22, 0, 0,
		293, 296, 3, 24, 12, 0, 294, 296, 3, 26, 13, 0, 295, 293, 1, 0, 0, 0, 295,
		294, 1, 0, 0, 0, 296, 23, 1, 0, 0, 0, 297, 298, 3, 28, 14, 0, 298, 25,
		1, 0, 0, 0, 299, 300, 5, 27, 0, 0, 300, 301, 3, 28, 14, 0, 301, 27, 1,
		0, 0, 0, 302, 303, 3, 148, 74, 0, 303, 29, 1, 0, 0, 0, 304, 306, 5, 37,
		0, 0, 305, 304, 1, 0, 0, 0, 305, 306, 1, 0, 0, 0, 306, 307, 1, 0, 0, 0,
		307, 308, 3, 52, 26, 0, 308, 31, 1, 0, 0, 0, 309, 311, 3, 34, 17, 0, 310,
		309, 1, 0, 0, 0, 310, 311, 1, 0, 0, 0, 311, 313, 1, 0, 0, 0, 312, 314,
		3, 38, 19, 0, 313, 312, 1, 0, 0, 0, 313, 314, 1, 0, 0, 0, 314, 316, 1,
		0, 0, 0, 315, 317, 3, 44, 22, 0, 316, 315, 1, 0, 0, 0, 316, 317, 1, 0,
		0, 0, 317, 319, 1, 0, 0, 0, 318, 320, 3, 42, 21, 0, 319, 318, 1, 0, 0,
		0, 319, 320, 1, 0, 0, 0, 320, 33, 1, 0, 0, 0, 321, 322, 5, 7, 0, 0, 322,
		324, 5, 15, 0, 0, 323, 325, 3, 36, 18, 0, 324, 323, 1, 0, 0, 0, 325, 326,
		1, 0, 0, 0, 326, 324, 1, 0, 0, 0, 326, 327, 1, 0, 0, 0, 327, 35, 1, 0,
		0, 0, 328, 331, 3, 106, 53, 0, 329, 331, 3, 102, 51, 0, 330, 328, 1, 0,
		0, 0, 330, 329, 1, 0, 0, 0, 331, 37, 1, 0, 0, 0, 332, 334, 5, 8, 0, 0,
		333, 335, 3, 40, 20, 0, 334, 333, 1, 0, 0, 0, 335, 336, 1, 0, 0, 0, 336,
		334, 1, 0, 0, 0, 336, 337, 1, 0, 0, 0, 337, 39, 1, 0, 0, 0, 338, 339, 3,
		68, 34, 0, 339, 41, 1, 0, 0, 0, 340, 342, 3, 48, 24, 0, 341, 343, 3, 50,
		25, 0, 342, 341, 1, 0, 0, 0, 342, 343, 1, 0, 0, 0, 343, 349, 1, 0, 0, 0,
		344, 346, 3, 50, 25, 0, 345, 347, 3, 48, 24, 0, 346, 345, 1, 0, 0, 0, 346,
		347, 1, 0, 0, 0, 347, 349, 1, 0, 0, 0, 348, 340, 1, 0, 0, 0, 348, 344,
		1, 0, 0, 0, 349, 43, 1, 0, 0, 0, 350, 351, 5, 30, 0, 0, 351, 353, 5, 15,
		0, 0, 352, 354, 3, 46, 23, 0, 353, 352, 1, 0, 0, 0, 354, 355, 1, 0, 0,
		0, 355, 353, 1, 0, 0, 0, 355, 356, 1, 0, 0, 0, 356, 45, 1, 0, 0, 0, 357,
		358, 7, 1, 0, 0, 358, 362, 3, 126, 63, 0, 359, 362, 3, 68, 34, 0, 360,
		362, 3, 102, 51, 0, 361, 357, 1, 0, 0, 0, 361, 359, 1, 0, 0, 0, 361, 360,
		1, 0, 0, 0, 362, 47, 1, 0, 0, 0, 363, 364, 5, 26, 0, 0, 364, 365, 5, 75,
		0, 0, 365, 49, 1, 0, 0, 0, 366, 367, 5, 28, 0, 0, 367, 368, 5, 75, 0, 0,
		368, 51, 1, 0, 0, 0, 369, 371, 5, 56, 0, 0, 370, 372, 3, 54, 27, 0, 371,
		370, 1, 0, 0, 0, 371, 372, 1, 0, 0, 0, 372, 385, 1, 0, 0, 0, 373, 376,
		3, 56, 28, 0, 374, 376, 3, 66, 33, 0, 375, 373, 1, 0, 0, 0, 375, 374, 1,
		0, 0, 0, 376, 378, 1, 0, 0, 0, 377, 379, 5, 46, 0, 0, 378, 377, 1, 0, 0,
		0, 378, 379, 1, 0, 0, 0, 379, 381, 1, 0, 0, 0, 380, 382, 3, 54, 27, 0,
		381, 380, 1, 0, 0, 0, 381, 382, 1, 0, 0, 0, 382, 384, 1, 0, 0, 0, 383,
		375, 1, 0, 0, 0, 384, 387, 1, 0, 0, 0, 385, 383, 1, 0, 0, 0, 385, 386,
		1, 0, 0, 0, 386, 388, 1, 0, 0, 0, 387, 385, 1, 0, 0, 0, 388, 389, 5, 62,
		0, 0, 389, 53, 1, 0, 0, 0, 390, 395, 3, 78, 39, 0, 391, 393, 5, 46, 0,
		0, 392, 394, 3, 54, 27, 0, 393, 392, 1, 0, 0, 0, 393, 394, 1, 0, 0, 0,
		394, 396, 1, 0, 0, 0, 395, 391, 1, 0, 0, 0, 395, 396, 1, 0, 0, 0, 396,
		55, 1, 0, 0, 0, 397, 402, 3, 60, 30, 0, 398, 402, 3, 64, 32, 0, 399, 402,
		3, 62, 31, 0, 400, 402, 3, 58, 29, 0, 401, 397, 1, 0, 0, 0, 401, 398, 1,
		0, 0, 0, 401, 399, 1, 0, 0, 0, 401, 400, 1, 0, 0, 0, 402, 57, 1, 0, 0,
		0, 403, 404, 5, 5, 0, 0, 404, 405, 5, 57, 0, 0, 405, 406, 3, 106, 53, 0,
		406, 407, 5, 6, 0, 0, 407, 408, 3, 102, 51, 0, 408, 409, 5, 63, 0, 0, 409,
		59, 1, 0, 0, 0, 410, 411, 5, 29, 0, 0, 411, 412, 3, 52, 26, 0, 412, 61,
		1, 0, 0, 0, 413, 414, 5, 23, 0, 0, 414, 415, 3, 100, 50, 0, 415, 416, 3,
		52, 26, 0, 416, 63, 1, 0, 0, 0, 417, 422, 3, 52, 26, 0, 418, 419, 5, 36,
		0, 0, 419, 421, 3, 52, 26, 0, 420, 418, 1, 0, 0, 0, 421, 424, 1, 0, 0,
		0, 422, 420, 1, 0, 0, 0, 422, 423, 1, 0, 0, 0, 423, 65, 1, 0, 0, 0, 424,
		422, 1, 0, 0, 0, 425, 426, 5, 21, 0, 0, 426, 427, 3, 68, 34, 0, 427, 67,
		1, 0, 0, 0, 428, 432, 3, 126, 63, 0, 429, 432, 3, 128, 64, 0, 430, 432,
		3, 70, 35, 0, 431, 428, 1, 0, 0, 0, 431, 429, 1, 0, 0, 0, 431, 430, 1,
		0, 0, 0, 432, 69, 1, 0, 0, 0, 433, 434, 3, 148, 74, 0, 434, 435, 3, 72,
		36, 0, 435, 71, 1, 0, 0, 0, 436, 449, 5, 90, 0, 0, 437, 438, 5, 57, 0,
		0, 438, 443, 3, 106, 53, 0, 439, 440, 5, 45, 0, 0, 440, 442, 3, 106, 53,
		0, 441, 439, 1, 0, 0, 0, 442, 445, 1, 0, 0, 0, 443, 441, 1, 0, 0, 0, 443,
		444, 1, 0, 0, 0, 444, 446, 1, 0, 0, 0, 445, 443, 1, 0, 0, 0, 446, 447,
		5, 63, 0, 0, 447, 449, 1, 0, 0, 0, 448, 436, 1, 0, 0, 0, 448, 437, 1, 0,
		0, 0, 449, 73, 1, 0, 0, 0, 450, 452, 5, 56, 0, 0, 451, 453, 3, 76, 38,
		0, 452, 451, 1, 0, 0, 0, 452, 453, 1, 0, 0, 0, 453, 454, 1, 0, 0, 0, 454,
		455, 5, 62, 0, 0, 455, 75, 1, 0, 0, 0, 456, 461, 3, 78, 39, 0, 457, 459,
		5, 46, 0, 0, 458, 460, 3, 76, 38, 0, 459, 458, 1, 0, 0, 0, 459, 460, 1,
		0, 0, 0, 460, 462, 1, 0, 0, 0, 461, 457, 1, 0, 0, 0, 461, 462, 1, 0, 0,
		0, 462, 77, 1, 0, 0, 0, 463, 464, 3, 98, 49, 0, 464, 465, 3, 80, 40, 0,
		465, 470, 1, 0, 0, 0, 466, 467, 3, 90, 45, 0, 467, 468, 3, 82, 41, 0, 468,
		470, 1, 0, 0, 0, 469, 463, 1, 0, 0, 0, 469, 466, 1, 0, 0, 0, 470, 79, 1,
		0, 0, 0, 471, 472, 3, 88, 44, 0, 472, 481, 3, 84, 42, 0, 473, 477, 5, 65,
		0, 0, 474, 475, 3, 88, 44, 0, 475, 476, 3, 84, 42, 0, 476, 478, 1, 0, 0,
		0, 477, 474, 1, 0, 0, 0, 477, 478, 1, 0, 0, 0, 478, 480, 1, 0, 0, 0, 479,
		473, 1, 0, 0, 0, 480, 483, 1, 0, 0, 0, 481, 479, 1, 0, 0, 0, 481, 482,
		1, 0, 0, 0, 482, 81, 1, 0, 0, 0, 483, 481, 1, 0, 0, 0, 484, 486, 3, 80,
		40, 0, 485, 484, 1, 0, 0, 0, 485, 486, 1, 0, 0, 0, 486, 83, 1, 0, 0, 0,
		487, 492, 3, 86, 43, 0, 488, 489, 5, 45, 0, 0, 489, 491, 3, 86, 43, 0,
		490, 488, 1, 0, 0, 0, 491, 494, 1, 0, 0, 0, 492, 490, 1, 0, 0, 0, 492,
		493, 1, 0, 0, 0, 493, 85, 1, 0, 0, 0, 494, 492, 1, 0, 0, 0, 495, 496, 3,
		96, 48, 0, 496, 87, 1, 0, 0, 0, 497, 500, 3, 100, 50, 0, 498, 500, 5, 1,
		0, 0, 499, 497, 1, 0, 0, 0, 499, 498, 1, 0, 0, 0, 500, 89, 1, 0, 0, 0,
		501, 504, 3, 94, 47, 0, 502, 504, 3, 92, 46, 0, 503, 501, 1, 0, 0, 0, 503,
		502, 1, 0, 0, 0, 504, 91, 1, 0, 0, 0, 505, 506, 5, 58, 0, 0, 506, 507,
		3, 80, 40, 0, 507, 508, 5, 64, 0, 0, 508, 93, 1, 0, 0, 0, 509, 511, 5,
		57, 0, 0, 510, 512, 3, 96, 48, 0, 511, 510, 1, 0, 0, 0, 512, 513, 1, 0,
		0, 0, 513, 511, 1, 0, 0, 0, 513, 514, 1, 0, 0, 0, 514, 515, 1, 0, 0, 0,
		515, 516, 5, 63, 0, 0, 516, 95, 1, 0, 0, 0, 517, 520, 3, 98, 49, 0, 518,
		520, 3, 90, 45, 0, 519, 517, 1, 0, 0, 0, 519, 518, 1, 0, 0, 0, 520, 97,
		1, 0, 0, 0, 521, 524, 3, 102, 51, 0, 522, 524, 3, 104, 52, 0, 523, 521,
		1, 0, 0, 0, 523, 522, 1, 0, 0, 0, 524, 99, 1, 0, 0, 0, 525, 528, 3, 102,
		51, 0, 526, 528, 3, 148, 74, 0, 527, 525, 1, 0, 0, 0, 527, 526, 1, 0, 0,
		0, 528, 101, 1, 0, 0, 0, 529, 530, 7, 2, 0, 0, 530, 103, 1, 0, 0, 0, 531,
		538, 3, 148, 74, 0, 532, 538, 3, 134, 67, 0, 533, 538, 3, 136, 68, 0, 534,
		538, 3, 144, 72, 0, 535, 538, 3, 152, 76, 0, 536, 538, 5, 90, 0, 0, 537,
		531, 1, 0, 0, 0, 537, 532, 1, 0, 0, 0, 537, 533, 1, 0, 0, 0, 537, 534,
		1, 0, 0, 0, 537, 535, 1, 0, 0, 0, 537, 536, 1, 0, 0, 0, 538, 105, 1, 0,
		0, 0, 539, 540, 3, 108, 54, 0, 540, 107, 1, 0, 0, 0, 541, 546, 3, 110,
		55, 0, 542, 543, 5, 48, 0, 0, 543, 545, 3, 110, 55, 0, 544, 542, 1, 0,
		0, 0, 545, 548, 1, 0, 0, 0, 546, 544, 1, 0, 0, 0, 546, 547, 1, 0, 0, 0,
		547, 109, 1, 0, 0, 0, 548, 546, 1, 0, 0, 0, 549, 554, 3, 112, 56, 0, 550,
		551, 5, 47, 0, 0, 551, 553, 3, 112, 56, 0, 552, 550, 1, 0, 0, 0, 553, 556,
		1, 0, 0, 0, 554, 552, 1, 0, 0, 0, 554, 555, 1, 0, 0, 0, 555, 111, 1, 0,
		0, 0, 556, 554, 1, 0, 0, 0, 557, 558, 3, 114, 57, 0, 558, 113, 1, 0, 0,
		0, 559, 562, 3, 116, 58, 0, 560, 561, 7, 3, 0, 0, 561, 563, 3, 116, 58,
		0, 562, 560, 1, 0, 0, 0, 562, 563, 1, 0, 0, 0, 563, 115, 1, 0, 0, 0, 564,
		565, 3, 118, 59, 0, 565, 117, 1, 0, 0, 0, 566, 573, 3, 120, 60, 0, 567,
		568, 7, 4, 0, 0, 568, 572, 3, 120, 60, 0, 569, 572, 3, 140, 70, 0, 570,
		572, 3, 142, 71, 0, 571, 567, 1, 0, 0, 0, 571, 569, 1, 0, 0, 0, 571, 570,
		1, 0, 0, 0, 572, 575, 1, 0, 0, 0, 573, 571, 1, 0, 0, 0, 573, 574, 1, 0,
		0, 0, 574, 119, 1, 0, 0, 0, 575, 573, 1, 0, 0, 0, 576, 581, 3, 122, 61,
		0, 577, 578, 7, 5, 0, 0, 578, 580, 3, 122, 61, 0, 579, 577, 1, 0, 0, 0,
		580, 583, 1, 0, 0, 0, 581, 579, 1, 0, 0, 0, 581, 582, 1, 0, 0, 0, 582,
		121, 1, 0, 0, 0, 583, 581, 1, 0, 0, 0, 584, 586, 7, 6, 0, 0, 585, 584,
		1, 0, 0, 0, 585, 586, 1, 0, 0, 0, 586, 587, 1, 0, 0, 0, 587, 588, 3, 124,
		62, 0, 588, 123, 1, 0, 0, 0, 589, 597, 3, 126, 63, 0, 590, 597, 3, 128,
		64, 0, 591, 597, 3, 132, 66, 0, 592, 597, 3, 134, 67, 0, 593, 597, 3, 136,
		68, 0, 594, 597, 3, 144, 72, 0, 595, 597, 3, 102, 51, 0, 596, 589, 1, 0,
		0, 0, 596, 590, 1, 0, 0, 0, 596, 591, 1, 0, 0, 0, 596, 592, 1, 0, 0, 0,
		596, 593, 1, 0, 0, 0, 596, 594, 1, 0, 0, 0, 596, 595, 1, 0, 0, 0, 597,
		125, 1, 0, 0, 0, 598, 599, 5, 57, 0, 0, 599, 600, 3, 106, 53, 0, 600, 601,
		5, 63, 0, 0, 601, 127, 1, 0, 0, 0, 602, 603, 5, 35, 0, 0, 603, 604, 5,
		57, 0, 0, 604, 605, 3, 106, 53, 0, 605, 606, 5, 63, 0, 0, 606, 658, 1,
		0, 0, 0, 607, 608, 5, 24, 0, 0, 608, 609, 5, 57, 0, 0, 609, 610, 3, 106,
		53, 0, 610, 611, 5, 63, 0, 0, 611, 658, 1, 0, 0, 0, 612, 613, 5, 25, 0,
		0, 613, 614, 5, 57, 0, 0, 614, 615, 3, 106, 53, 0, 615, 616, 5, 45, 0,
		0, 616, 617, 3, 106, 53, 0, 617, 618, 5, 63, 0, 0, 618, 658, 1, 0, 0, 0,
		619, 620, 5, 17, 0, 0, 620, 621, 5, 57, 0, 0, 621, 622, 3, 106, 53, 0,
		622, 623, 5, 63, 0, 0, 623, 658, 1, 0, 0, 0, 624, 625, 5, 14, 0, 0, 625,
		626, 5, 57, 0, 0, 626, 627, 3, 102, 51, 0, 627, 628, 5, 63, 0, 0, 628,
		658, 1, 0, 0, 0, 629, 630, 5, 44, 0, 0, 630, 631, 5, 57, 0, 0, 631, 632,
		3, 106, 53, 0, 632, 633, 5, 45, 0, 0, 633, 634, 3, 106, 53, 0, 634, 635,
		5, 63, 0, 0, 635, 658, 1, 0, 0, 0, 636, 637, 5, 43, 0, 0, 637, 638, 5,
		57, 0, 0, 638, 639, 3, 106, 53, 0, 639, 640, 5, 63, 0, 0, 640, 658, 1,
		0, 0, 0, 641, 642, 5, 42, 0, 0, 642, 643, 5, 57, 0, 0, 643, 644, 3, 106,
		53, 0, 644, 645, 5, 63, 0, 0, 645, 658, 1, 0, 0, 0, 646, 647, 5, 41, 0,
		0, 647, 648, 5, 57, 0, 0, 648, 649, 3, 106, 53, 0, 649, 650, 5, 63, 0,
		0, 650, 658, 1, 0, 0, 0, 651, 652, 5, 40, 0, 0, 652, 653, 5, 57, 0, 0,
		653, 654, 3, 106, 53, 0, 654, 655, 5, 63, 0, 0, 655, 658, 1, 0, 0, 0, 656,
		658, 3, 130, 65, 0, 657, 602, 1, 0, 0, 0, 657, 607, 1, 0, 0, 0, 657, 612,
		1, 0, 0, 0, 657, 619, 1, 0, 0, 0, 657, 624, 1, 0, 0, 0, 657, 629, 1, 0,
		0, 0, 657, 636, 1, 0, 0, 0, 657, 641, 1, 0, 0, 0, 657, 646, 1, 0, 0, 0,
		657, 651, 1, 0, 0, 0, 657, 656, 1, 0, 0, 0, 658, 129, 1, 0, 0, 0, 659,
		660, 5, 33, 0, 0, 660, 661, 5, 57, 0, 0, 661, 662, 3, 106, 53, 0, 662,
		663, 5, 45, 0, 0, 663, 666, 3, 106, 53, 0, 664, 665, 5, 45, 0, 0, 665,
		667, 3, 106, 53, 0, 666, 664, 1, 0, 0, 0, 666, 667, 1, 0, 0, 0, 667, 668,
		1, 0, 0, 0, 668, 669, 5, 63, 0, 0, 669, 131, 1, 0, 0, 0, 670, 672, 3, 148,
		74, 0, 671, 673, 3, 72, 36, 0, 672, 671, 1, 0, 0, 0, 672, 673, 1, 0, 0,
		0, 673, 133, 1, 0, 0, 0, 674, 678, 3, 146, 73, 0, 675, 679, 5, 74, 0, 0,
		676, 677, 5, 49, 0, 0, 677, 679, 3, 148, 74, 0, 678, 675, 1, 0, 0, 0, 678,
		676, 1, 0, 0, 0, 678, 679, 1, 0, 0, 0, 679, 135, 1, 0, 0, 0, 680, 684,
		3, 138, 69, 0, 681, 684, 3, 140, 70, 0, 682, 684, 3, 142, 71, 0, 683, 680,
		1, 0, 0, 0, 683, 681, 1, 0, 0, 0, 683, 682, 1, 0, 0, 0, 684, 137, 1, 0,
		0, 0, 685, 686, 7, 7, 0, 0, 686, 139, 1, 0, 0, 0, 687, 688, 7, 8, 0, 0,
		688, 141, 1, 0, 0, 0, 689, 690, 7, 9, 0, 0, 690, 143, 1, 0, 0, 0, 691,
		692, 7, 10, 0, 0, 692, 145, 1, 0, 0, 0, 693, 694, 7, 11, 0, 0, 694, 147,
		1, 0, 0, 0, 695, 698, 5, 68, 0, 0, 696, 698, 3, 150, 75, 0, 697, 695, 1,
		0, 0, 0, 697, 696, 1, 0, 0, 0, 698, 149, 1, 0, 0, 0, 699, 700, 7, 12, 0,
		0, 700, 151, 1, 0, 0, 0, 701, 702, 7, 13, 0, 0, 702, 153, 1, 0, 0, 0, 75,
		159, 164, 169, 181, 186, 189, 194, 202, 213, 221, 229, 237, 245, 250, 257,
		267, 270, 275, 279, 287, 295, 305, 310, 313, 316, 319, 326, 330, 336, 342,
		346, 348, 355, 361, 371, 375, 378, 381, 385, 393, 395, 401, 422, 431, 443,
		448, 452, 459, 461, 469, 477, 481, 485, 492, 499, 503, 513, 519, 523, 527,
		537, 546, 554, 562, 571, 573, 581, 585, 596, 657, 666, 672, 678, 683, 697,
	}
	deserializer := antlr.NewATNDeserializer(nil)
	staticData.atn = deserializer.Deserialize(staticData.serializedATN)
	atn := staticData.atn
	staticData.decisionToDFA = make([]*antlr.DFA, len(atn.DecisionToState))
	decisionToDFA := staticData.decisionToDFA
	for index, state := range atn.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(state, index)
	}
}

// SparqlParserInit initializes any static state used to implement SparqlParser. By default the
// static state used to implement the parser is lazily initialized during the first call to
// NewSparqlParser(). You can call this function if you wish to initialize the static state ahead
// of time.
func SparqlParserInit() {
	staticData := &SparqlParserParserStaticData
	staticData.once.Do(sparqlparserParserInit)
}

// NewSparqlParser produces a new parser instance for the optional input antlr.TokenStream.
func NewSparqlParser(input antlr.TokenStream) *SparqlParser {
	SparqlParserInit()
	this := new(SparqlParser)
	this.BaseParser = antlr.NewBaseParser(input)
	staticData := &SparqlParserParserStaticData
	this.Interpreter = antlr.NewParserATNSimulator(this, staticData.atn, staticData.decisionToDFA, staticData.PredictionContextCache)
	this.RuleNames = staticData.RuleNames
	this.LiteralNames = staticData.LiteralNames
	this.SymbolicNames = staticData.SymbolicNames
	this.GrammarFileName = "SparqlParser.g4"

	return this
}

// SparqlParser tokens.
const (
	SparqlParserEOF                  = antlr.TokenEOF
	SparqlParserA                    = 1
	SparqlParserASC                  = 2
	SparqlParserASK                  = 3
	SparqlParserBASE                 = 4
	SparqlParserBIND                 = 5
	SparqlParserAS                   = 6
	SparqlParserGROUP                = 7
	SparqlParserHAVING               = 8
	SparqlParserCOUNT                = 9
	SparqlParserSUM                  = 10
	SparqlParserAVG                  = 11
	SparqlParserMIN                  = 12
	SparqlParserMAX                  = 13
	SparqlParserBOUND                = 14
	SparqlParserBY                   = 15
	SparqlParserCONSTRUCT            = 16
	SparqlParserDATATYPE             = 17
	SparqlParserDESC                 = 18
	SparqlParserDESCRIBE             = 19
	SparqlParserDISTINCT             = 20
	SparqlParserFILTER               = 21
	SparqlParserFROM                 = 22
	SparqlParserGRAPH                = 23
	SparqlParserLANG                 = 24
	SparqlParserLANGMATCHES          = 25
	SparqlParserLIMIT                = 26
	SparqlParserNAMED                = 27
	SparqlParserOFFSET               = 28
	SparqlParserOPTIONAL             = 29
	SparqlParserORDER                = 30
	SparqlParserPREFIX               = 31
	SparqlParserREDUCED              = 32
	SparqlParserREGEX                = 33
	SparqlParserSELECT               = 34
	SparqlParserSTR                  = 35
	SparqlParserUNION                = 36
	SparqlParserWHERE                = 37
	SparqlParserTRUE                 = 38
	SparqlParserFALSE                = 39
	SparqlParserIS_LITERAL           = 40
	SparqlParserIS_BLANK             = 41
	SparqlParserIS_URI               = 42
	SparqlParserIS_IRI               = 43
	SparqlParserSAME_TERM            = 44
	SparqlParserCOMMA                = 45
	SparqlParserDOT                  = 46
	SparqlParserDOUBLE_AMP           = 47
	SparqlParserDOUBLE_BAR           = 48
	SparqlParserDOUBLE_CARET         = 49
	SparqlParserEQUAL                = 50
	SparqlParserEXCLAMATION          = 51
	SparqlParserGREATER              = 52
	SparqlParserGREATER_OR_EQUAL     = 53
	SparqlParserLESS                 = 54
	SparqlParserLESS_OR_EQUAL        = 55
	SparqlParserL_CURLY              = 56
	SparqlParserL_PAREN              = 57
	SparqlParserL_SQUARE             = 58
	SparqlParserMINUS                = 59
	SparqlParserNOT_EQUAL            = 60
	SparqlParserPLUS                 = 61
	SparqlParserR_CURLY              = 62
	SparqlParserR_PAREN              = 63
	SparqlParserR_SQUARE             = 64
	SparqlParserSEMICOLON            = 65
	SparqlParserSLASH                = 66
	SparqlParserSTAR                 = 67
	SparqlParserIRI_REF              = 68
	SparqlParserPNAME_NS             = 69
	SparqlParserPNAME_LN             = 70
	SparqlParserBLANK_NODE_LABEL     = 71
	SparqlParserVAR1                 = 72
	SparqlParserVAR2                 = 73
	SparqlParserLANGTAG              = 74
	SparqlParserINTEGER              = 75
	SparqlParserDECIMAL              = 76
	SparqlParserDOUBLE               = 77
	SparqlParserINTEGER_POSITIVE     = 78
	SparqlParserDECIMAL_POSITIVE     = 79
	SparqlParserDOUBLE_POSITIVE      = 80
	SparqlParserINTEGER_NEGATIVE     = 81
	SparqlParserDECIMAL_NEGATIVE     = 82
	SparqlParserDOUBLE_NEGATIVE      = 83
	SparqlParserEXPONENT             = 84
	SparqlParserSTRING_LITERAL1      = 85
	SparqlParserSTRING_LITERAL2      = 86
	SparqlParserSTRING_LITERAL_LONG1 = 87
	SparqlParserSTRING_LITERAL_LONG2 = 88
	SparqlParserECHAR                = 89
	SparqlParserNIL                  = 90
	SparqlParserANON                 = 91
	SparqlParserPN_CHARS_U           = 92
	SparqlParserVARNAME              = 93
	SparqlParserPN_PREFIX            = 94
	SparqlParserPN_LOCAL             = 95
	SparqlParserWS                   = 96
	SparqlParserCOMMENT              = 97
)

// SparqlParser rules.
const (
	SparqlParserRULE_query                    = 0
	SparqlParserRULE_prologue                 = 1
	SparqlParserRULE_baseDecl                 = 2
	SparqlParserRULE_prefixDecl               = 3
	SparqlParserRULE_selectQuery              = 4
	SparqlParserRULE_selectProjection         = 5
	SparqlParserRULE_aggregateProjection      = 6
	SparqlParserRULE_aggregateFunction        = 7
	SparqlParserRULE_constructQuery           = 8
	SparqlParserRULE_describeQuery            = 9
	SparqlParserRULE_askQuery                 = 10
	SparqlParserRULE_datasetClause            = 11
	SparqlParserRULE_defaultGraphClause       = 12
	SparqlParserRULE_namedGraphClause         = 13
	SparqlParserRULE_sourceSelector           = 14
	SparqlParserRULE_whereClause              = 15
	SparqlParserRULE_solutionModifier         = 16
	SparqlParserRULE_groupClause              = 17
	SparqlParserRULE_groupCondition           = 18
	SparqlParserRULE_havingClause             = 19
	SparqlParserRULE_havingCondition          = 20
	SparqlParserRULE_limitOffsetClauses       = 21
	SparqlParserRULE_orderClause              = 22
	SparqlParserRULE_orderCondition           = 23
	SparqlParserRULE_limitClause              = 24
	SparqlParserRULE_offsetClause             = 25
	SparqlParserRULE_groupGraphPattern        = 26
	SparqlParserRULE_triplesBlock             = 27
	SparqlParserRULE_graphPatternNotTriples   = 28
	SparqlParserRULE_bind                     = 29
	SparqlParserRULE_optionalGraphPattern     = 30
	SparqlParserRULE_graphGraphPattern        = 31
	SparqlParserRULE_groupOrUnionGraphPattern = 32
	SparqlParserRULE_filter_                  = 33
	SparqlParserRULE_constraint               = 34
	SparqlParserRULE_functionCall             = 35
	SparqlParserRULE_argList                  = 36
	SparqlParserRULE_constructTemplate        = 37
	SparqlParserRULE_constructTriples         = 38
	SparqlParserRULE_triplesSameSubject       = 39
	SparqlParserRULE_propertyListNotEmpty     = 40
	SparqlParserRULE_propertyList             = 41
	SparqlParserRULE_objectList               = 42
	SparqlParserRULE_object_                  = 43
	SparqlParserRULE_verb                     = 44
	SparqlParserRULE_triplesNode              = 45
	SparqlParserRULE_blankNodePropertyList    = 46
	SparqlParserRULE_collection               = 47
	SparqlParserRULE_graphNode                = 48
	SparqlParserRULE_varOrTerm                = 49
	SparqlParserRULE_varOrIRIref              = 50
	SparqlParserRULE_var_                     = 51
	SparqlParserRULE_graphTerm                = 52
	SparqlParserRULE_expression               = 53
	SparqlParserRULE_conditionalOrExpression  = 54
	SparqlParserRULE_conditionalAndExpression = 55
	SparqlParserRULE_valueLogical             = 56
	SparqlParserRULE_relationalExpression     = 57
	SparqlParserRULE_numericExpression        = 58
	SparqlParserRULE_additiveExpression       = 59
	SparqlParserRULE_multiplicativeExpression = 60
	SparqlParserRULE_unaryExpression          = 61
	SparqlParserRULE_primaryExpression        = 62
	SparqlParserRULE_brackettedExpression     = 63
	SparqlParserRULE_builtInCall              = 64
	SparqlParserRULE_regexExpression          = 65
	SparqlParserRULE_iriRefOrFunction         = 66
	SparqlParserRULE_rdfLiteral               = 67
	SparqlParserRULE_numericLiteral           = 68
	SparqlParserRULE_numericLiteralUnsigned   = 69
	SparqlParserRULE_numericLiteralPositive   = 70
	SparqlParserRULE_numericLiteralNegative   = 71
	SparqlParserRULE_booleanLiteral           = 72
	SparqlParserRULE_string_                  = 73
	SparqlParserRULE_iriRef                   = 74
	SparqlParserRULE_prefixedName             = 75
	SparqlParserRULE_blankNode                = 76
)

// IQueryContext is an interface to support dynamic dispatch.
type IQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Prologue() IPrologueContext
	EOF() antlr.TerminalNode
	SelectQuery() ISelectQueryContext
	ConstructQuery() IConstructQueryContext
	DescribeQuery() IDescribeQueryContext
	AskQuery() IAskQueryContext

	// IsQueryContext differentiates from other interfaces.
	IsQueryContext()
}

type QueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyQueryContext() *QueryContext {
	var p = new(QueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_query
	return p
}

func InitEmptyQueryContext(p *QueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_query
}

func (*QueryContext) IsQueryContext() {}

func NewQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *QueryContext {
	var p = new(QueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_query

	return p
}

func (s *QueryContext) GetParser() antlr.Parser { return s.parser }

func (s *QueryContext) Prologue() IPrologueContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrologueContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrologueContext)
}

func (s *QueryContext) EOF() antlr.TerminalNode {
	return s.GetToken(SparqlParserEOF, 0)
}

func (s *QueryContext) SelectQuery() ISelectQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectQueryContext)
}

func (s *QueryContext) ConstructQuery() IConstructQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructQueryContext)
}

func (s *QueryContext) DescribeQuery() IDescribeQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDescribeQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDescribeQueryContext)
}

func (s *QueryContext) AskQuery() IAskQueryContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAskQueryContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAskQueryContext)
}

func (s *QueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *QueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *QueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterQuery(s)
	}
}

func (s *QueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitQuery(s)
	}
}

func (s *QueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Query() (localctx IQueryContext) {
	localctx = NewQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, SparqlParserRULE_query)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(154)
		p.Prologue()
	}
	p.SetState(159)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserSELECT:
		{
			p.SetState(155)
			p.SelectQuery()
		}

	case SparqlParserCONSTRUCT:
		{
			p.SetState(156)
			p.ConstructQuery()
		}

	case SparqlParserDESCRIBE:
		{
			p.SetState(157)
			p.DescribeQuery()
		}

	case SparqlParserASK:
		{
			p.SetState(158)
			p.AskQuery()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	{
		p.SetState(161)
		p.Match(SparqlParserEOF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrologueContext is an interface to support dynamic dispatch.
type IPrologueContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BaseDecl() IBaseDeclContext
	AllPrefixDecl() []IPrefixDeclContext
	PrefixDecl(i int) IPrefixDeclContext

	// IsPrologueContext differentiates from other interfaces.
	IsPrologueContext()
}

type PrologueContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrologueContext() *PrologueContext {
	var p = new(PrologueContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prologue
	return p
}

func InitEmptyPrologueContext(p *PrologueContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prologue
}

func (*PrologueContext) IsPrologueContext() {}

func NewPrologueContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrologueContext {
	var p = new(PrologueContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prologue

	return p
}

func (s *PrologueContext) GetParser() antlr.Parser { return s.parser }

func (s *PrologueContext) BaseDecl() IBaseDeclContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBaseDeclContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBaseDeclContext)
}

func (s *PrologueContext) AllPrefixDecl() []IPrefixDeclContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IPrefixDeclContext); ok {
			len++
		}
	}

	tst := make([]IPrefixDeclContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IPrefixDeclContext); ok {
			tst[i] = t.(IPrefixDeclContext)
			i++
		}
	}

	return tst
}

func (s *PrologueContext) PrefixDecl(i int) IPrefixDeclContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixDeclContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixDeclContext)
}

func (s *PrologueContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrologueContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrologueContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPrologue(s)
	}
}

func (s *PrologueContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPrologue(s)
	}
}

func (s *PrologueContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPrologue(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Prologue() (localctx IPrologueContext) {
	localctx = NewPrologueContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, SparqlParserRULE_prologue)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(164)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserBASE {
		{
			p.SetState(163)
			p.BaseDecl()
		}

	}
	p.SetState(169)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserPREFIX {
		{
			p.SetState(166)
			p.PrefixDecl()
		}

		p.SetState(171)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBaseDeclContext is an interface to support dynamic dispatch.
type IBaseDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BASE() antlr.TerminalNode
	IRI_REF() antlr.TerminalNode

	// IsBaseDeclContext differentiates from other interfaces.
	IsBaseDeclContext()
}

type BaseDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBaseDeclContext() *BaseDeclContext {
	var p = new(BaseDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_baseDecl
	return p
}

func InitEmptyBaseDeclContext(p *BaseDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_baseDecl
}

func (*BaseDeclContext) IsBaseDeclContext() {}

func NewBaseDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BaseDeclContext {
	var p = new(BaseDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_baseDecl

	return p
}

func (s *BaseDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *BaseDeclContext) BASE() antlr.TerminalNode {
	return s.GetToken(SparqlParserBASE, 0)
}

func (s *BaseDeclContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *BaseDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BaseDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BaseDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBaseDecl(s)
	}
}

func (s *BaseDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBaseDecl(s)
	}
}

func (s *BaseDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBaseDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BaseDecl() (localctx IBaseDeclContext) {
	localctx = NewBaseDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, SparqlParserRULE_baseDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(172)
		p.Match(SparqlParserBASE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(173)
		p.Match(SparqlParserIRI_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixDeclContext is an interface to support dynamic dispatch.
type IPrefixDeclContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PREFIX() antlr.TerminalNode
	PNAME_NS() antlr.TerminalNode
	IRI_REF() antlr.TerminalNode

	// IsPrefixDeclContext differentiates from other interfaces.
	IsPrefixDeclContext()
}

type PrefixDeclContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixDeclContext() *PrefixDeclContext {
	var p = new(PrefixDeclContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixDecl
	return p
}

func InitEmptyPrefixDeclContext(p *PrefixDeclContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixDecl
}

func (*PrefixDeclContext) IsPrefixDeclContext() {}

func NewPrefixDeclContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixDeclContext {
	var p = new(PrefixDeclContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prefixDecl

	return p
}

func (s *PrefixDeclContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixDeclContext) PREFIX() antlr.TerminalNode {
	return s.GetToken(SparqlParserPREFIX, 0)
}

func (s *PrefixDeclContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_NS, 0)
}

func (s *PrefixDeclContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *PrefixDeclContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixDeclContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixDeclContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPrefixDecl(s)
	}
}

func (s *PrefixDeclContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPrefixDecl(s)
	}
}

func (s *PrefixDeclContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPrefixDecl(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) PrefixDecl() (localctx IPrefixDeclContext) {
	localctx = NewPrefixDeclContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, SparqlParserRULE_prefixDecl)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(175)
		p.Match(SparqlParserPREFIX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(176)
		p.Match(SparqlParserPNAME_NS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(177)
		p.Match(SparqlParserIRI_REF)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectQueryContext is an interface to support dynamic dispatch.
type ISelectQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SELECT() antlr.TerminalNode
	WhereClause() IWhereClauseContext
	SolutionModifier() ISolutionModifierContext
	STAR() antlr.TerminalNode
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext
	DISTINCT() antlr.TerminalNode
	REDUCED() antlr.TerminalNode
	AllSelectProjection() []ISelectProjectionContext
	SelectProjection(i int) ISelectProjectionContext

	// IsSelectQueryContext differentiates from other interfaces.
	IsSelectQueryContext()
}

type SelectQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectQueryContext() *SelectQueryContext {
	var p = new(SelectQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectQuery
	return p
}

func InitEmptySelectQueryContext(p *SelectQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectQuery
}

func (*SelectQueryContext) IsSelectQueryContext() {}

func NewSelectQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectQueryContext {
	var p = new(SelectQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_selectQuery

	return p
}

func (s *SelectQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectQueryContext) SELECT() antlr.TerminalNode {
	return s.GetToken(SparqlParserSELECT, 0)
}

func (s *SelectQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *SelectQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *SelectQueryContext) STAR() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTAR, 0)
}

func (s *SelectQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *SelectQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *SelectQueryContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SparqlParserDISTINCT, 0)
}

func (s *SelectQueryContext) REDUCED() antlr.TerminalNode {
	return s.GetToken(SparqlParserREDUCED, 0)
}

func (s *SelectQueryContext) AllSelectProjection() []ISelectProjectionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ISelectProjectionContext); ok {
			len++
		}
	}

	tst := make([]ISelectProjectionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ISelectProjectionContext); ok {
			tst[i] = t.(ISelectProjectionContext)
			i++
		}
	}

	return tst
}

func (s *SelectQueryContext) SelectProjection(i int) ISelectProjectionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISelectProjectionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISelectProjectionContext)
}

func (s *SelectQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterSelectQuery(s)
	}
}

func (s *SelectQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitSelectQuery(s)
	}
}

func (s *SelectQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitSelectQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) SelectQuery() (localctx ISelectQueryContext) {
	localctx = NewSelectQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, SparqlParserRULE_selectQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(179)
		p.Match(SparqlParserSELECT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(181)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserDISTINCT || _la == SparqlParserREDUCED {
		{
			p.SetState(180)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SparqlParserDISTINCT || _la == SparqlParserREDUCED) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	p.SetState(189)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserL_PAREN, SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(184)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-57)) & ^0x3f) == 0 && ((int64(1)<<(_la-57))&98305) != 0) {
			{
				p.SetState(183)
				p.SelectProjection()
			}

			p.SetState(186)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserSTAR:
		{
			p.SetState(188)
			p.Match(SparqlParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(194)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserFROM {
		{
			p.SetState(191)
			p.DatasetClause()
		}

		p.SetState(196)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(197)
		p.WhereClause()
	}
	{
		p.SetState(198)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISelectProjectionContext is an interface to support dynamic dispatch.
type ISelectProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVar_Context
	AggregateProjection() IAggregateProjectionContext

	// IsSelectProjectionContext differentiates from other interfaces.
	IsSelectProjectionContext()
}

type SelectProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySelectProjectionContext() *SelectProjectionContext {
	var p = new(SelectProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectProjection
	return p
}

func InitEmptySelectProjectionContext(p *SelectProjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_selectProjection
}

func (*SelectProjectionContext) IsSelectProjectionContext() {}

func NewSelectProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SelectProjectionContext {
	var p = new(SelectProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_selectProjection

	return p
}

func (s *SelectProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *SelectProjectionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *SelectProjectionContext) AggregateProjection() IAggregateProjectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateProjectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateProjectionContext)
}

func (s *SelectProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SelectProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SelectProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterSelectProjection(s)
	}
}

func (s *SelectProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitSelectProjection(s)
	}
}

func (s *SelectProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitSelectProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) SelectProjection() (localctx ISelectProjectionContext) {
	localctx = NewSelectProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, SparqlParserRULE_selectProjection)
	p.SetState(202)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(200)
			p.Var_()
		}

	case SparqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(201)
			p.AggregateProjection()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateProjectionContext is an interface to support dynamic dispatch.
type IAggregateProjectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	AggregateFunction() IAggregateFunctionContext
	AS() antlr.TerminalNode
	Var_() IVar_Context
	R_PAREN() antlr.TerminalNode

	// IsAggregateProjectionContext differentiates from other interfaces.
	IsAggregateProjectionContext()
}

type AggregateProjectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateProjectionContext() *AggregateProjectionContext {
	var p = new(AggregateProjectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_aggregateProjection
	return p
}

func InitEmptyAggregateProjectionContext(p *AggregateProjectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_aggregateProjection
}

func (*AggregateProjectionContext) IsAggregateProjectionContext() {}

func NewAggregateProjectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateProjectionContext {
	var p = new(AggregateProjectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_aggregateProjection

	return p
}

func (s *AggregateProjectionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateProjectionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *AggregateProjectionContext) AggregateFunction() IAggregateFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAggregateFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAggregateFunctionContext)
}

func (s *AggregateProjectionContext) AS() antlr.TerminalNode {
	return s.GetToken(SparqlParserAS, 0)
}

func (s *AggregateProjectionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *AggregateProjectionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *AggregateProjectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateProjectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateProjectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterAggregateProjection(s)
	}
}

func (s *AggregateProjectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitAggregateProjection(s)
	}
}

func (s *AggregateProjectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitAggregateProjection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) AggregateProjection() (localctx IAggregateProjectionContext) {
	localctx = NewAggregateProjectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, SparqlParserRULE_aggregateProjection)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(204)
		p.Match(SparqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(205)
		p.AggregateFunction()
	}
	{
		p.SetState(206)
		p.Match(SparqlParserAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(207)
		p.Var_()
	}
	{
		p.SetState(208)
		p.Match(SparqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAggregateFunctionContext is an interface to support dynamic dispatch.
type IAggregateFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	COUNT() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Var_() IVar_Context
	R_PAREN() antlr.TerminalNode
	DISTINCT() antlr.TerminalNode
	SUM() antlr.TerminalNode
	AVG() antlr.TerminalNode
	MIN() antlr.TerminalNode
	MAX() antlr.TerminalNode

	// IsAggregateFunctionContext differentiates from other interfaces.
	IsAggregateFunctionContext()
}

type AggregateFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAggregateFunctionContext() *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_aggregateFunction
	return p
}

func InitEmptyAggregateFunctionContext(p *AggregateFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_aggregateFunction
}

func (*AggregateFunctionContext) IsAggregateFunctionContext() {}

func NewAggregateFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AggregateFunctionContext {
	var p = new(AggregateFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_aggregateFunction

	return p
}

func (s *AggregateFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *AggregateFunctionContext) COUNT() antlr.TerminalNode {
	return s.GetToken(SparqlParserCOUNT, 0)
}

func (s *AggregateFunctionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *AggregateFunctionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *AggregateFunctionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *AggregateFunctionContext) DISTINCT() antlr.TerminalNode {
	return s.GetToken(SparqlParserDISTINCT, 0)
}

func (s *AggregateFunctionContext) SUM() antlr.TerminalNode {
	return s.GetToken(SparqlParserSUM, 0)
}

func (s *AggregateFunctionContext) AVG() antlr.TerminalNode {
	return s.GetToken(SparqlParserAVG, 0)
}

func (s *AggregateFunctionContext) MIN() antlr.TerminalNode {
	return s.GetToken(SparqlParserMIN, 0)
}

func (s *AggregateFunctionContext) MAX() antlr.TerminalNode {
	return s.GetToken(SparqlParserMAX, 0)
}

func (s *AggregateFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AggregateFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AggregateFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitAggregateFunction(s)
	}
}

func (s *AggregateFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitAggregateFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) AggregateFunction() (localctx IAggregateFunctionContext) {
	localctx = NewAggregateFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, SparqlParserRULE_aggregateFunction)
	var _la int

	p.SetState(250)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserCOUNT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(210)
			p.Match(SparqlParserCOUNT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(211)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(213)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDISTINCT {
			{
				p.SetState(212)
				p.Match(SparqlParserDISTINCT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(215)
			p.Var_()
		}
		{
			p.SetState(216)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserSUM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(218)
			p.Match(SparqlParserSUM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(219)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(221)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDISTINCT {
			{
				p.SetState(220)
				p.Match(SparqlParserDISTINCT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(223)
			p.Var_()
		}
		{
			p.SetState(224)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserAVG:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(226)
			p.Match(SparqlParserAVG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(227)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(229)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDISTINCT {
			{
				p.SetState(228)
				p.Match(SparqlParserDISTINCT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(231)
			p.Var_()
		}
		{
			p.SetState(232)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserMIN:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(234)
			p.Match(SparqlParserMIN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(235)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(237)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDISTINCT {
			{
				p.SetState(236)
				p.Match(SparqlParserDISTINCT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(239)
			p.Var_()
		}
		{
			p.SetState(240)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserMAX:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(242)
			p.Match(SparqlParserMAX)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(243)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(245)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDISTINCT {
			{
				p.SetState(244)
				p.Match(SparqlParserDISTINCT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		{
			p.SetState(247)
			p.Var_()
		}
		{
			p.SetState(248)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructQueryContext is an interface to support dynamic dispatch.
type IConstructQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	CONSTRUCT() antlr.TerminalNode
	ConstructTemplate() IConstructTemplateContext
	WhereClause() IWhereClauseContext
	SolutionModifier() ISolutionModifierContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext

	// IsConstructQueryContext differentiates from other interfaces.
	IsConstructQueryContext()
}

type ConstructQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructQueryContext() *ConstructQueryContext {
	var p = new(ConstructQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructQuery
	return p
}

func InitEmptyConstructQueryContext(p *ConstructQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructQuery
}

func (*ConstructQueryContext) IsConstructQueryContext() {}

func NewConstructQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructQueryContext {
	var p = new(ConstructQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructQuery

	return p
}

func (s *ConstructQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructQueryContext) CONSTRUCT() antlr.TerminalNode {
	return s.GetToken(SparqlParserCONSTRUCT, 0)
}

func (s *ConstructQueryContext) ConstructTemplate() IConstructTemplateContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTemplateContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTemplateContext)
}

func (s *ConstructQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *ConstructQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *ConstructQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *ConstructQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *ConstructQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConstructQuery(s)
	}
}

func (s *ConstructQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConstructQuery(s)
	}
}

func (s *ConstructQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConstructQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ConstructQuery() (localctx IConstructQueryContext) {
	localctx = NewConstructQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, SparqlParserRULE_constructQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(252)
		p.Match(SparqlParserCONSTRUCT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(253)
		p.ConstructTemplate()
	}
	p.SetState(257)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserFROM {
		{
			p.SetState(254)
			p.DatasetClause()
		}

		p.SetState(259)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(260)
		p.WhereClause()
	}
	{
		p.SetState(261)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDescribeQueryContext is an interface to support dynamic dispatch.
type IDescribeQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	DESCRIBE() antlr.TerminalNode
	SolutionModifier() ISolutionModifierContext
	STAR() antlr.TerminalNode
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext
	WhereClause() IWhereClauseContext
	AllVarOrIRIref() []IVarOrIRIrefContext
	VarOrIRIref(i int) IVarOrIRIrefContext

	// IsDescribeQueryContext differentiates from other interfaces.
	IsDescribeQueryContext()
}

type DescribeQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDescribeQueryContext() *DescribeQueryContext {
	var p = new(DescribeQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_describeQuery
	return p
}

func InitEmptyDescribeQueryContext(p *DescribeQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_describeQuery
}

func (*DescribeQueryContext) IsDescribeQueryContext() {}

func NewDescribeQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DescribeQueryContext {
	var p = new(DescribeQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_describeQuery

	return p
}

func (s *DescribeQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *DescribeQueryContext) DESCRIBE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDESCRIBE, 0)
}

func (s *DescribeQueryContext) SolutionModifier() ISolutionModifierContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISolutionModifierContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISolutionModifierContext)
}

func (s *DescribeQueryContext) STAR() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTAR, 0)
}

func (s *DescribeQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *DescribeQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *DescribeQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *DescribeQueryContext) AllVarOrIRIref() []IVarOrIRIrefContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			len++
		}
	}

	tst := make([]IVarOrIRIrefContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVarOrIRIrefContext); ok {
			tst[i] = t.(IVarOrIRIrefContext)
			i++
		}
	}

	return tst
}

func (s *DescribeQueryContext) VarOrIRIref(i int) IVarOrIRIrefContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *DescribeQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DescribeQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DescribeQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterDescribeQuery(s)
	}
}

func (s *DescribeQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitDescribeQuery(s)
	}
}

func (s *DescribeQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitDescribeQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) DescribeQuery() (localctx IDescribeQueryContext) {
	localctx = NewDescribeQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, SparqlParserRULE_describeQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(263)
		p.Match(SparqlParserDESCRIBE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(270)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.SetState(265)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for ok := true; ok; ok = ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&55) != 0) {
			{
				p.SetState(264)
				p.VarOrIRIref()
			}

			p.SetState(267)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}

	case SparqlParserSTAR:
		{
			p.SetState(269)
			p.Match(SparqlParserSTAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}
	p.SetState(275)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserFROM {
		{
			p.SetState(272)
			p.DatasetClause()
		}

		p.SetState(277)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	p.SetState(279)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserWHERE || _la == SparqlParserL_CURLY {
		{
			p.SetState(278)
			p.WhereClause()
		}

	}
	{
		p.SetState(281)
		p.SolutionModifier()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAskQueryContext is an interface to support dynamic dispatch.
type IAskQueryContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ASK() antlr.TerminalNode
	WhereClause() IWhereClauseContext
	AllDatasetClause() []IDatasetClauseContext
	DatasetClause(i int) IDatasetClauseContext

	// IsAskQueryContext differentiates from other interfaces.
	IsAskQueryContext()
}

type AskQueryContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAskQueryContext() *AskQueryContext {
	var p = new(AskQueryContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_askQuery
	return p
}

func InitEmptyAskQueryContext(p *AskQueryContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_askQuery
}

func (*AskQueryContext) IsAskQueryContext() {}

func NewAskQueryContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AskQueryContext {
	var p = new(AskQueryContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_askQuery

	return p
}

func (s *AskQueryContext) GetParser() antlr.Parser { return s.parser }

func (s *AskQueryContext) ASK() antlr.TerminalNode {
	return s.GetToken(SparqlParserASK, 0)
}

func (s *AskQueryContext) WhereClause() IWhereClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IWhereClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IWhereClauseContext)
}

func (s *AskQueryContext) AllDatasetClause() []IDatasetClauseContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			len++
		}
	}

	tst := make([]IDatasetClauseContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IDatasetClauseContext); ok {
			tst[i] = t.(IDatasetClauseContext)
			i++
		}
	}

	return tst
}

func (s *AskQueryContext) DatasetClause(i int) IDatasetClauseContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDatasetClauseContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDatasetClauseContext)
}

func (s *AskQueryContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AskQueryContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AskQueryContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterAskQuery(s)
	}
}

func (s *AskQueryContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitAskQuery(s)
	}
}

func (s *AskQueryContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitAskQuery(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) AskQuery() (localctx IAskQueryContext) {
	localctx = NewAskQueryContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, SparqlParserRULE_askQuery)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(283)
		p.Match(SparqlParserASK)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(287)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserFROM {
		{
			p.SetState(284)
			p.DatasetClause()
		}

		p.SetState(289)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(290)
		p.WhereClause()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDatasetClauseContext is an interface to support dynamic dispatch.
type IDatasetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FROM() antlr.TerminalNode
	DefaultGraphClause() IDefaultGraphClauseContext
	NamedGraphClause() INamedGraphClauseContext

	// IsDatasetClauseContext differentiates from other interfaces.
	IsDatasetClauseContext()
}

type DatasetClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDatasetClauseContext() *DatasetClauseContext {
	var p = new(DatasetClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_datasetClause
	return p
}

func InitEmptyDatasetClauseContext(p *DatasetClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_datasetClause
}

func (*DatasetClauseContext) IsDatasetClauseContext() {}

func NewDatasetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DatasetClauseContext {
	var p = new(DatasetClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_datasetClause

	return p
}

func (s *DatasetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DatasetClauseContext) FROM() antlr.TerminalNode {
	return s.GetToken(SparqlParserFROM, 0)
}

func (s *DatasetClauseContext) DefaultGraphClause() IDefaultGraphClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IDefaultGraphClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IDefaultGraphClauseContext)
}

func (s *DatasetClauseContext) NamedGraphClause() INamedGraphClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INamedGraphClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INamedGraphClauseContext)
}

func (s *DatasetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DatasetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DatasetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterDatasetClause(s)
	}
}

func (s *DatasetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitDatasetClause(s)
	}
}

func (s *DatasetClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitDatasetClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) DatasetClause() (localctx IDatasetClauseContext) {
	localctx = NewDatasetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, SparqlParserRULE_datasetClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(292)
		p.Match(SparqlParserFROM)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(295)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		{
			p.SetState(293)
			p.DefaultGraphClause()
		}

	case SparqlParserNAMED:
		{
			p.SetState(294)
			p.NamedGraphClause()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IDefaultGraphClauseContext is an interface to support dynamic dispatch.
type IDefaultGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	SourceSelector() ISourceSelectorContext

	// IsDefaultGraphClauseContext differentiates from other interfaces.
	IsDefaultGraphClauseContext()
}

type DefaultGraphClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyDefaultGraphClauseContext() *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_defaultGraphClause
	return p
}

func InitEmptyDefaultGraphClauseContext(p *DefaultGraphClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_defaultGraphClause
}

func (*DefaultGraphClauseContext) IsDefaultGraphClauseContext() {}

func NewDefaultGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultGraphClauseContext {
	var p = new(DefaultGraphClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_defaultGraphClause

	return p
}

func (s *DefaultGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceSelectorContext)
}

func (s *DefaultGraphClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultGraphClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultGraphClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterDefaultGraphClause(s)
	}
}

func (s *DefaultGraphClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitDefaultGraphClause(s)
	}
}

func (s *DefaultGraphClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitDefaultGraphClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) DefaultGraphClause() (localctx IDefaultGraphClauseContext) {
	localctx = NewDefaultGraphClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, SparqlParserRULE_defaultGraphClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(297)
		p.SourceSelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INamedGraphClauseContext is an interface to support dynamic dispatch.
type INamedGraphClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NAMED() antlr.TerminalNode
	SourceSelector() ISourceSelectorContext

	// IsNamedGraphClauseContext differentiates from other interfaces.
	IsNamedGraphClauseContext()
}

type NamedGraphClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNamedGraphClauseContext() *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_namedGraphClause
	return p
}

func InitEmptyNamedGraphClauseContext(p *NamedGraphClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_namedGraphClause
}

func (*NamedGraphClauseContext) IsNamedGraphClauseContext() {}

func NewNamedGraphClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NamedGraphClauseContext {
	var p = new(NamedGraphClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_namedGraphClause

	return p
}

func (s *NamedGraphClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *NamedGraphClauseContext) NAMED() antlr.TerminalNode {
	return s.GetToken(SparqlParserNAMED, 0)
}

func (s *NamedGraphClauseContext) SourceSelector() ISourceSelectorContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ISourceSelectorContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ISourceSelectorContext)
}

func (s *NamedGraphClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NamedGraphClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NamedGraphClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNamedGraphClause(s)
	}
}

func (s *NamedGraphClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNamedGraphClause(s)
	}
}

func (s *NamedGraphClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNamedGraphClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NamedGraphClause() (localctx INamedGraphClauseContext) {
	localctx = NewNamedGraphClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, SparqlParserRULE_namedGraphClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(299)
		p.Match(SparqlParserNAMED)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(300)
		p.SourceSelector()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISourceSelectorContext is an interface to support dynamic dispatch.
type ISourceSelectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext

	// IsSourceSelectorContext differentiates from other interfaces.
	IsSourceSelectorContext()
}

type SourceSelectorContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySourceSelectorContext() *SourceSelectorContext {
	var p = new(SourceSelectorContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_sourceSelector
	return p
}

func InitEmptySourceSelectorContext(p *SourceSelectorContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_sourceSelector
}

func (*SourceSelectorContext) IsSourceSelectorContext() {}

func NewSourceSelectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SourceSelectorContext {
	var p = new(SourceSelectorContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_sourceSelector

	return p
}

func (s *SourceSelectorContext) GetParser() antlr.Parser { return s.parser }

func (s *SourceSelectorContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *SourceSelectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SourceSelectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SourceSelectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterSourceSelector(s)
	}
}

func (s *SourceSelectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitSourceSelector(s)
	}
}

func (s *SourceSelectorContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitSourceSelector(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) SourceSelector() (localctx ISourceSelectorContext) {
	localctx = NewSourceSelectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, SparqlParserRULE_sourceSelector)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(302)
		p.IriRef()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IWhereClauseContext is an interface to support dynamic dispatch.
type IWhereClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GroupGraphPattern() IGroupGraphPatternContext
	WHERE() antlr.TerminalNode

	// IsWhereClauseContext differentiates from other interfaces.
	IsWhereClauseContext()
}

type WhereClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyWhereClauseContext() *WhereClauseContext {
	var p = new(WhereClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_whereClause
	return p
}

func InitEmptyWhereClauseContext(p *WhereClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_whereClause
}

func (*WhereClauseContext) IsWhereClauseContext() {}

func NewWhereClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *WhereClauseContext {
	var p = new(WhereClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_whereClause

	return p
}

func (s *WhereClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *WhereClauseContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *WhereClauseContext) WHERE() antlr.TerminalNode {
	return s.GetToken(SparqlParserWHERE, 0)
}

func (s *WhereClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *WhereClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *WhereClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterWhereClause(s)
	}
}

func (s *WhereClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitWhereClause(s)
	}
}

func (s *WhereClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitWhereClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) WhereClause() (localctx IWhereClauseContext) {
	localctx = NewWhereClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 30, SparqlParserRULE_whereClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(305)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserWHERE {
		{
			p.SetState(304)
			p.Match(SparqlParserWHERE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	}
	{
		p.SetState(307)
		p.GroupGraphPattern()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ISolutionModifierContext is an interface to support dynamic dispatch.
type ISolutionModifierContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GroupClause() IGroupClauseContext
	HavingClause() IHavingClauseContext
	OrderClause() IOrderClauseContext
	LimitOffsetClauses() ILimitOffsetClausesContext

	// IsSolutionModifierContext differentiates from other interfaces.
	IsSolutionModifierContext()
}

type SolutionModifierContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptySolutionModifierContext() *SolutionModifierContext {
	var p = new(SolutionModifierContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_solutionModifier
	return p
}

func InitEmptySolutionModifierContext(p *SolutionModifierContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_solutionModifier
}

func (*SolutionModifierContext) IsSolutionModifierContext() {}

func NewSolutionModifierContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *SolutionModifierContext {
	var p = new(SolutionModifierContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_solutionModifier

	return p
}

func (s *SolutionModifierContext) GetParser() antlr.Parser { return s.parser }

func (s *SolutionModifierContext) GroupClause() IGroupClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupClauseContext)
}

func (s *SolutionModifierContext) HavingClause() IHavingClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingClauseContext)
}

func (s *SolutionModifierContext) OrderClause() IOrderClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderClauseContext)
}

func (s *SolutionModifierContext) LimitOffsetClauses() ILimitOffsetClausesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitOffsetClausesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitOffsetClausesContext)
}

func (s *SolutionModifierContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *SolutionModifierContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *SolutionModifierContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterSolutionModifier(s)
	}
}

func (s *SolutionModifierContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitSolutionModifier(s)
	}
}

func (s *SolutionModifierContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitSolutionModifier(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) SolutionModifier() (localctx ISolutionModifierContext) {
	localctx = NewSolutionModifierContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, SparqlParserRULE_solutionModifier)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(310)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserGROUP {
		{
			p.SetState(309)
			p.GroupClause()
		}

	}
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserHAVING {
		{
			p.SetState(312)
			p.HavingClause()
		}

	}
	p.SetState(316)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserORDER {
		{
			p.SetState(315)
			p.OrderClause()
		}

	}
	p.SetState(319)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserLIMIT || _la == SparqlParserOFFSET {
		{
			p.SetState(318)
			p.LimitOffsetClauses()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupClauseContext is an interface to support dynamic dispatch.
type IGroupClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GROUP() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllGroupCondition() []IGroupConditionContext
	GroupCondition(i int) IGroupConditionContext

	// IsGroupClauseContext differentiates from other interfaces.
	IsGroupClauseContext()
}

type GroupClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupClauseContext() *GroupClauseContext {
	var p = new(GroupClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupClause
	return p
}

func InitEmptyGroupClauseContext(p *GroupClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupClause
}

func (*GroupClauseContext) IsGroupClauseContext() {}

func NewGroupClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupClauseContext {
	var p = new(GroupClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupClause

	return p
}

func (s *GroupClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupClauseContext) GROUP() antlr.TerminalNode {
	return s.GetToken(SparqlParserGROUP, 0)
}

func (s *GroupClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SparqlParserBY, 0)
}

func (s *GroupClauseContext) AllGroupCondition() []IGroupConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupConditionContext); ok {
			len++
		}
	}

	tst := make([]IGroupConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupConditionContext); ok {
			tst[i] = t.(IGroupConditionContext)
			i++
		}
	}

	return tst
}

func (s *GroupClauseContext) GroupCondition(i int) IGroupConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupConditionContext)
}

func (s *GroupClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGroupClause(s)
	}
}

func (s *GroupClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGroupClause(s)
	}
}

func (s *GroupClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGroupClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GroupClause() (localctx IGroupClauseContext) {
	localctx = NewGroupClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 34, SparqlParserRULE_groupClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(321)
		p.Match(SparqlParserGROUP)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(322)
		p.Match(SparqlParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(324)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&3028705701900992512) != 0) || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&458679) != 0) {
		{
			p.SetState(323)
			p.GroupCondition()
		}

		p.SetState(326)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupConditionContext is an interface to support dynamic dispatch.
type IGroupConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Expression() IExpressionContext
	Var_() IVar_Context

	// IsGroupConditionContext differentiates from other interfaces.
	IsGroupConditionContext()
}

type GroupConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupConditionContext() *GroupConditionContext {
	var p = new(GroupConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupCondition
	return p
}

func InitEmptyGroupConditionContext(p *GroupConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupCondition
}

func (*GroupConditionContext) IsGroupConditionContext() {}

func NewGroupConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupConditionContext {
	var p = new(GroupConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupCondition

	return p
}

func (s *GroupConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupConditionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *GroupConditionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *GroupConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGroupCondition(s)
	}
}

func (s *GroupConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGroupCondition(s)
	}
}

func (s *GroupConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGroupCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GroupCondition() (localctx IGroupConditionContext) {
	localctx = NewGroupConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 36, SparqlParserRULE_groupCondition)
	p.SetState(330)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 27, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(328)
			p.Expression()
		}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(329)
			p.Var_()
		}

	case antlr.ATNInvalidAltNumber:
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHavingClauseContext is an interface to support dynamic dispatch.
type IHavingClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	HAVING() antlr.TerminalNode
	AllHavingCondition() []IHavingConditionContext
	HavingCondition(i int) IHavingConditionContext

	// IsHavingClauseContext differentiates from other interfaces.
	IsHavingClauseContext()
}

type HavingClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingClauseContext() *HavingClauseContext {
	var p = new(HavingClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_havingClause
	return p
}

func InitEmptyHavingClauseContext(p *HavingClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_havingClause
}

func (*HavingClauseContext) IsHavingClauseContext() {}

func NewHavingClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingClauseContext {
	var p = new(HavingClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_havingClause

	return p
}

func (s *HavingClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingClauseContext) HAVING() antlr.TerminalNode {
	return s.GetToken(SparqlParserHAVING, 0)
}

func (s *HavingClauseContext) AllHavingCondition() []IHavingConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IHavingConditionContext); ok {
			len++
		}
	}

	tst := make([]IHavingConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IHavingConditionContext); ok {
			tst[i] = t.(IHavingConditionContext)
			i++
		}
	}

	return tst
}

func (s *HavingClauseContext) HavingCondition(i int) IHavingConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IHavingConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IHavingConditionContext)
}

func (s *HavingClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterHavingClause(s)
	}
}

func (s *HavingClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitHavingClause(s)
	}
}

func (s *HavingClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitHavingClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) HavingClause() (localctx IHavingClauseContext) {
	localctx = NewHavingClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, SparqlParserRULE_havingClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(332)
		p.Match(SparqlParserHAVING)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(334)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-14)) & ^0x3f) == 0 && ((int64(1)<<(_la-14))&126109587742395401) != 0) {
		{
			p.SetState(333)
			p.HavingCondition()
		}

		p.SetState(336)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IHavingConditionContext is an interface to support dynamic dispatch.
type IHavingConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Constraint() IConstraintContext

	// IsHavingConditionContext differentiates from other interfaces.
	IsHavingConditionContext()
}

type HavingConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyHavingConditionContext() *HavingConditionContext {
	var p = new(HavingConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_havingCondition
	return p
}

func InitEmptyHavingConditionContext(p *HavingConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_havingCondition
}

func (*HavingConditionContext) IsHavingConditionContext() {}

func NewHavingConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *HavingConditionContext {
	var p = new(HavingConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_havingCondition

	return p
}

func (s *HavingConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *HavingConditionContext) Constraint() IConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *HavingConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *HavingConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *HavingConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterHavingCondition(s)
	}
}

func (s *HavingConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitHavingCondition(s)
	}
}

func (s *HavingConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitHavingCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) HavingCondition() (localctx IHavingConditionContext) {
	localctx = NewHavingConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, SparqlParserRULE_havingCondition)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(338)
		p.Constraint()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitOffsetClausesContext is an interface to support dynamic dispatch.
type ILimitOffsetClausesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LimitClause() ILimitClauseContext
	OffsetClause() IOffsetClauseContext

	// IsLimitOffsetClausesContext differentiates from other interfaces.
	IsLimitOffsetClausesContext()
}

type LimitOffsetClausesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitOffsetClausesContext() *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses
	return p
}

func InitEmptyLimitOffsetClausesContext(p *LimitOffsetClausesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses
}

func (*LimitOffsetClausesContext) IsLimitOffsetClausesContext() {}

func NewLimitOffsetClausesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitOffsetClausesContext {
	var p = new(LimitOffsetClausesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_limitOffsetClauses

	return p
}

func (s *LimitOffsetClausesContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitOffsetClausesContext) LimitClause() ILimitClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ILimitClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ILimitClauseContext)
}

func (s *LimitOffsetClausesContext) OffsetClause() IOffsetClauseContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOffsetClauseContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOffsetClauseContext)
}

func (s *LimitOffsetClausesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitOffsetClausesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitOffsetClausesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterLimitOffsetClauses(s)
	}
}

func (s *LimitOffsetClausesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitLimitOffsetClauses(s)
	}
}

func (s *LimitOffsetClausesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitLimitOffsetClauses(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) LimitOffsetClauses() (localctx ILimitOffsetClausesContext) {
	localctx = NewLimitOffsetClausesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, SparqlParserRULE_limitOffsetClauses)
	var _la int

	p.SetState(348)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserLIMIT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(340)
			p.LimitClause()
		}
		p.SetState(342)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserOFFSET {
			{
				p.SetState(341)
				p.OffsetClause()
			}

		}

	case SparqlParserOFFSET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(344)
			p.OffsetClause()
		}
		p.SetState(346)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserLIMIT {
			{
				p.SetState(345)
				p.LimitClause()
			}

		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderClauseContext is an interface to support dynamic dispatch.
type IOrderClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ORDER() antlr.TerminalNode
	BY() antlr.TerminalNode
	AllOrderCondition() []IOrderConditionContext
	OrderCondition(i int) IOrderConditionContext

	// IsOrderClauseContext differentiates from other interfaces.
	IsOrderClauseContext()
}

type OrderClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderClauseContext() *OrderClauseContext {
	var p = new(OrderClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderClause
	return p
}

func InitEmptyOrderClauseContext(p *OrderClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderClause
}

func (*OrderClauseContext) IsOrderClauseContext() {}

func NewOrderClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderClauseContext {
	var p = new(OrderClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderClause

	return p
}

func (s *OrderClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderClauseContext) ORDER() antlr.TerminalNode {
	return s.GetToken(SparqlParserORDER, 0)
}

func (s *OrderClauseContext) BY() antlr.TerminalNode {
	return s.GetToken(SparqlParserBY, 0)
}

func (s *OrderClauseContext) AllOrderCondition() []IOrderConditionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IOrderConditionContext); ok {
			len++
		}
	}

	tst := make([]IOrderConditionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IOrderConditionContext); ok {
			tst[i] = t.(IOrderConditionContext)
			i++
		}
	}

	return tst
}

func (s *OrderClauseContext) OrderCondition(i int) IOrderConditionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOrderConditionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOrderConditionContext)
}

func (s *OrderClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterOrderClause(s)
	}
}

func (s *OrderClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitOrderClause(s)
	}
}

func (s *OrderClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitOrderClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) OrderClause() (localctx IOrderClauseContext) {
	localctx = NewOrderClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 44, SparqlParserRULE_orderClause)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(350)
		p.Match(SparqlParserORDER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(351)
		p.Match(SparqlParserBY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(353)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&144149315936731140) != 0) || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&55) != 0) {
		{
			p.SetState(352)
			p.OrderCondition()
		}

		p.SetState(355)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOrderConditionContext is an interface to support dynamic dispatch.
type IOrderConditionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	ASC() antlr.TerminalNode
	DESC() antlr.TerminalNode
	Constraint() IConstraintContext
	Var_() IVar_Context

	// IsOrderConditionContext differentiates from other interfaces.
	IsOrderConditionContext()
}

type OrderConditionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOrderConditionContext() *OrderConditionContext {
	var p = new(OrderConditionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderCondition
	return p
}

func InitEmptyOrderConditionContext(p *OrderConditionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_orderCondition
}

func (*OrderConditionContext) IsOrderConditionContext() {}

func NewOrderConditionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OrderConditionContext {
	var p = new(OrderConditionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_orderCondition

	return p
}

func (s *OrderConditionContext) GetParser() antlr.Parser { return s.parser }

func (s *OrderConditionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *OrderConditionContext) ASC() antlr.TerminalNode {
	return s.GetToken(SparqlParserASC, 0)
}

func (s *OrderConditionContext) DESC() antlr.TerminalNode {
	return s.GetToken(SparqlParserDESC, 0)
}

func (s *OrderConditionContext) Constraint() IConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *OrderConditionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *OrderConditionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OrderConditionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OrderConditionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterOrderCondition(s)
	}
}

func (s *OrderConditionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitOrderCondition(s)
	}
}

func (s *OrderConditionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitOrderCondition(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) OrderCondition() (localctx IOrderConditionContext) {
	localctx = NewOrderConditionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, SparqlParserRULE_orderCondition)
	var _la int

	p.SetState(361)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserASC, SparqlParserDESC:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(357)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SparqlParserASC || _la == SparqlParserDESC) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(358)
			p.BrackettedExpression()
		}

	case SparqlParserBOUND, SparqlParserDATATYPE, SparqlParserLANG, SparqlParserLANGMATCHES, SparqlParserREGEX, SparqlParserSTR, SparqlParserIS_LITERAL, SparqlParserIS_BLANK, SparqlParserIS_URI, SparqlParserIS_IRI, SparqlParserSAME_TERM, SparqlParserL_PAREN, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(359)
			p.Constraint()
		}

	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(360)
			p.Var_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ILimitClauseContext is an interface to support dynamic dispatch.
type ILimitClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	LIMIT() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsLimitClauseContext differentiates from other interfaces.
	IsLimitClauseContext()
}

type LimitClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyLimitClauseContext() *LimitClauseContext {
	var p = new(LimitClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitClause
	return p
}

func InitEmptyLimitClauseContext(p *LimitClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_limitClause
}

func (*LimitClauseContext) IsLimitClauseContext() {}

func NewLimitClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LimitClauseContext {
	var p = new(LimitClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_limitClause

	return p
}

func (s *LimitClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *LimitClauseContext) LIMIT() antlr.TerminalNode {
	return s.GetToken(SparqlParserLIMIT, 0)
}

func (s *LimitClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *LimitClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LimitClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LimitClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterLimitClause(s)
	}
}

func (s *LimitClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitLimitClause(s)
	}
}

func (s *LimitClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitLimitClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) LimitClause() (localctx ILimitClauseContext) {
	localctx = NewLimitClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 48, SparqlParserRULE_limitClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(363)
		p.Match(SparqlParserLIMIT)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(364)
		p.Match(SparqlParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOffsetClauseContext is an interface to support dynamic dispatch.
type IOffsetClauseContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OFFSET() antlr.TerminalNode
	INTEGER() antlr.TerminalNode

	// IsOffsetClauseContext differentiates from other interfaces.
	IsOffsetClauseContext()
}

type OffsetClauseContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOffsetClauseContext() *OffsetClauseContext {
	var p = new(OffsetClauseContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_offsetClause
	return p
}

func InitEmptyOffsetClauseContext(p *OffsetClauseContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_offsetClause
}

func (*OffsetClauseContext) IsOffsetClauseContext() {}

func NewOffsetClauseContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OffsetClauseContext {
	var p = new(OffsetClauseContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_offsetClause

	return p
}

func (s *OffsetClauseContext) GetParser() antlr.Parser { return s.parser }

func (s *OffsetClauseContext) OFFSET() antlr.TerminalNode {
	return s.GetToken(SparqlParserOFFSET, 0)
}

func (s *OffsetClauseContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *OffsetClauseContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OffsetClauseContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OffsetClauseContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterOffsetClause(s)
	}
}

func (s *OffsetClauseContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitOffsetClause(s)
	}
}

func (s *OffsetClauseContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitOffsetClause(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) OffsetClause() (localctx IOffsetClauseContext) {
	localctx = NewOffsetClauseContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, SparqlParserRULE_offsetClause)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(366)
		p.Match(SparqlParserOFFSET)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(367)
		p.Match(SparqlParserINTEGER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupGraphPatternContext is an interface to support dynamic dispatch.
type IGroupGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	AllTriplesBlock() []ITriplesBlockContext
	TriplesBlock(i int) ITriplesBlockContext
	AllGraphPatternNotTriples() []IGraphPatternNotTriplesContext
	GraphPatternNotTriples(i int) IGraphPatternNotTriplesContext
	AllFilter_() []IFilter_Context
	Filter_(i int) IFilter_Context
	AllDOT() []antlr.TerminalNode
	DOT(i int) antlr.TerminalNode

	// IsGroupGraphPatternContext differentiates from other interfaces.
	IsGroupGraphPatternContext()
}

type GroupGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupGraphPatternContext() *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupGraphPattern
	return p
}

func InitEmptyGroupGraphPatternContext(p *GroupGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupGraphPattern
}

func (*GroupGraphPatternContext) IsGroupGraphPatternContext() {}

func NewGroupGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupGraphPatternContext {
	var p = new(GroupGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupGraphPattern

	return p
}

func (s *GroupGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupGraphPatternContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_CURLY, 0)
}

func (s *GroupGraphPatternContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_CURLY, 0)
}

func (s *GroupGraphPatternContext) AllTriplesBlock() []ITriplesBlockContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(ITriplesBlockContext); ok {
			len++
		}
	}

	tst := make([]ITriplesBlockContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(ITriplesBlockContext); ok {
			tst[i] = t.(ITriplesBlockContext)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) TriplesBlock(i int) ITriplesBlockContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesBlockContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesBlockContext)
}

func (s *GroupGraphPatternContext) AllGraphPatternNotTriples() []IGraphPatternNotTriplesContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphPatternNotTriplesContext); ok {
			len++
		}
	}

	tst := make([]IGraphPatternNotTriplesContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphPatternNotTriplesContext); ok {
			tst[i] = t.(IGraphPatternNotTriplesContext)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) GraphPatternNotTriples(i int) IGraphPatternNotTriplesContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphPatternNotTriplesContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphPatternNotTriplesContext)
}

func (s *GroupGraphPatternContext) AllFilter_() []IFilter_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IFilter_Context); ok {
			len++
		}
	}

	tst := make([]IFilter_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IFilter_Context); ok {
			tst[i] = t.(IFilter_Context)
			i++
		}
	}

	return tst
}

func (s *GroupGraphPatternContext) Filter_(i int) IFilter_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFilter_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFilter_Context)
}

func (s *GroupGraphPatternContext) AllDOT() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserDOT)
}

func (s *GroupGraphPatternContext) DOT(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserDOT, i)
}

func (s *GroupGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGroupGraphPattern(s)
	}
}

func (s *GroupGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGroupGraphPattern(s)
	}
}

func (s *GroupGraphPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGroupGraphPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GroupGraphPattern() (localctx IGroupGraphPatternContext) {
	localctx = NewGroupGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, SparqlParserRULE_groupGraphPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(369)
		p.Match(SparqlParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(371)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0 {
		{
			p.SetState(370)
			p.TriplesBlock()
		}

	}
	p.SetState(385)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&72057594585284640) != 0 {
		p.SetState(375)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}

		switch p.GetTokenStream().LA(1) {
		case SparqlParserBIND, SparqlParserGRAPH, SparqlParserOPTIONAL, SparqlParserL_CURLY:
			{
				p.SetState(373)
				p.GraphPatternNotTriples()
			}

		case SparqlParserFILTER:
			{
				p.SetState(374)
				p.Filter_()
			}

		default:
			p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
			goto errorExit
		}
		p.SetState(378)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserDOT {
			{
				p.SetState(377)
				p.Match(SparqlParserDOT)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}

		}
		p.SetState(381)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0 {
			{
				p.SetState(380)
				p.TriplesBlock()
			}

		}

		p.SetState(387)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(388)
		p.Match(SparqlParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesBlockContext is an interface to support dynamic dispatch.
type ITriplesBlockContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TriplesSameSubject() ITriplesSameSubjectContext
	DOT() antlr.TerminalNode
	TriplesBlock() ITriplesBlockContext

	// IsTriplesBlockContext differentiates from other interfaces.
	IsTriplesBlockContext()
}

type TriplesBlockContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesBlockContext() *TriplesBlockContext {
	var p = new(TriplesBlockContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesBlock
	return p
}

func InitEmptyTriplesBlockContext(p *TriplesBlockContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesBlock
}

func (*TriplesBlockContext) IsTriplesBlockContext() {}

func NewTriplesBlockContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesBlockContext {
	var p = new(TriplesBlockContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesBlock

	return p
}

func (s *TriplesBlockContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesBlockContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesSameSubjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *TriplesBlockContext) DOT() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOT, 0)
}

func (s *TriplesBlockContext) TriplesBlock() ITriplesBlockContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesBlockContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesBlockContext)
}

func (s *TriplesBlockContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesBlockContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesBlockContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterTriplesBlock(s)
	}
}

func (s *TriplesBlockContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitTriplesBlock(s)
	}
}

func (s *TriplesBlockContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitTriplesBlock(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) TriplesBlock() (localctx ITriplesBlockContext) {
	localctx = NewTriplesBlockContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, SparqlParserRULE_triplesBlock)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(390)
		p.TriplesSameSubject()
	}
	p.SetState(395)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserDOT {
		{
			p.SetState(391)
			p.Match(SparqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(393)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0 {
			{
				p.SetState(392)
				p.TriplesBlock()
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphPatternNotTriplesContext is an interface to support dynamic dispatch.
type IGraphPatternNotTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OptionalGraphPattern() IOptionalGraphPatternContext
	GroupOrUnionGraphPattern() IGroupOrUnionGraphPatternContext
	GraphGraphPattern() IGraphGraphPatternContext
	Bind() IBindContext

	// IsGraphPatternNotTriplesContext differentiates from other interfaces.
	IsGraphPatternNotTriplesContext()
}

type GraphPatternNotTriplesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphPatternNotTriplesContext() *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples
	return p
}

func InitEmptyGraphPatternNotTriplesContext(p *GraphPatternNotTriplesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples
}

func (*GraphPatternNotTriplesContext) IsGraphPatternNotTriplesContext() {}

func NewGraphPatternNotTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphPatternNotTriplesContext {
	var p = new(GraphPatternNotTriplesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphPatternNotTriples

	return p
}

func (s *GraphPatternNotTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphPatternNotTriplesContext) OptionalGraphPattern() IOptionalGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IOptionalGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IOptionalGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GroupOrUnionGraphPattern() IGroupOrUnionGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupOrUnionGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupOrUnionGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) GraphGraphPattern() IGraphGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphGraphPatternContext)
}

func (s *GraphPatternNotTriplesContext) Bind() IBindContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBindContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBindContext)
}

func (s *GraphPatternNotTriplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphPatternNotTriplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphPatternNotTriplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGraphPatternNotTriples(s)
	}
}

func (s *GraphPatternNotTriplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGraphPatternNotTriples(s)
	}
}

func (s *GraphPatternNotTriplesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGraphPatternNotTriples(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GraphPatternNotTriples() (localctx IGraphPatternNotTriplesContext) {
	localctx = NewGraphPatternNotTriplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, SparqlParserRULE_graphPatternNotTriples)
	p.SetState(401)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserOPTIONAL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(397)
			p.OptionalGraphPattern()
		}

	case SparqlParserL_CURLY:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(398)
			p.GroupOrUnionGraphPattern()
		}

	case SparqlParserGRAPH:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(399)
			p.GraphGraphPattern()
		}

	case SparqlParserBIND:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(400)
			p.Bind()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBindContext is an interface to support dynamic dispatch.
type IBindContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BIND() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	AS() antlr.TerminalNode
	Var_() IVar_Context
	R_PAREN() antlr.TerminalNode

	// IsBindContext differentiates from other interfaces.
	IsBindContext()
}

type BindContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBindContext() *BindContext {
	var p = new(BindContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_bind
	return p
}

func InitEmptyBindContext(p *BindContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_bind
}

func (*BindContext) IsBindContext() {}

func NewBindContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BindContext {
	var p = new(BindContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_bind

	return p
}

func (s *BindContext) GetParser() antlr.Parser { return s.parser }

func (s *BindContext) BIND() antlr.TerminalNode {
	return s.GetToken(SparqlParserBIND, 0)
}

func (s *BindContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *BindContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BindContext) AS() antlr.TerminalNode {
	return s.GetToken(SparqlParserAS, 0)
}

func (s *BindContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *BindContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *BindContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BindContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BindContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBind(s)
	}
}

func (s *BindContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBind(s)
	}
}

func (s *BindContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBind(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Bind() (localctx IBindContext) {
	localctx = NewBindContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, SparqlParserRULE_bind)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(403)
		p.Match(SparqlParserBIND)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(404)
		p.Match(SparqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(405)
		p.Expression()
	}
	{
		p.SetState(406)
		p.Match(SparqlParserAS)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(407)
		p.Var_()
	}
	{
		p.SetState(408)
		p.Match(SparqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IOptionalGraphPatternContext is an interface to support dynamic dispatch.
type IOptionalGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	OPTIONAL() antlr.TerminalNode
	GroupGraphPattern() IGroupGraphPatternContext

	// IsOptionalGraphPatternContext differentiates from other interfaces.
	IsOptionalGraphPatternContext()
}

type OptionalGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyOptionalGraphPatternContext() *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern
	return p
}

func InitEmptyOptionalGraphPatternContext(p *OptionalGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern
}

func (*OptionalGraphPatternContext) IsOptionalGraphPatternContext() {}

func NewOptionalGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *OptionalGraphPatternContext {
	var p = new(OptionalGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_optionalGraphPattern

	return p
}

func (s *OptionalGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *OptionalGraphPatternContext) OPTIONAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserOPTIONAL, 0)
}

func (s *OptionalGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *OptionalGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *OptionalGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *OptionalGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterOptionalGraphPattern(s)
	}
}

func (s *OptionalGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitOptionalGraphPattern(s)
	}
}

func (s *OptionalGraphPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitOptionalGraphPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) OptionalGraphPattern() (localctx IOptionalGraphPatternContext) {
	localctx = NewOptionalGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 60, SparqlParserRULE_optionalGraphPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(410)
		p.Match(SparqlParserOPTIONAL)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(411)
		p.GroupGraphPattern()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphGraphPatternContext is an interface to support dynamic dispatch.
type IGraphGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GRAPH() antlr.TerminalNode
	VarOrIRIref() IVarOrIRIrefContext
	GroupGraphPattern() IGroupGraphPatternContext

	// IsGraphGraphPatternContext differentiates from other interfaces.
	IsGraphGraphPatternContext()
}

type GraphGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphGraphPatternContext() *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphGraphPattern
	return p
}

func InitEmptyGraphGraphPatternContext(p *GraphGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphGraphPattern
}

func (*GraphGraphPatternContext) IsGraphGraphPatternContext() {}

func NewGraphGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphGraphPatternContext {
	var p = new(GraphGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphGraphPattern

	return p
}

func (s *GraphGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphGraphPatternContext) GRAPH() antlr.TerminalNode {
	return s.GetToken(SparqlParserGRAPH, 0)
}

func (s *GraphGraphPatternContext) VarOrIRIref() IVarOrIRIrefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *GraphGraphPatternContext) GroupGraphPattern() IGroupGraphPatternContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *GraphGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGraphGraphPattern(s)
	}
}

func (s *GraphGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGraphGraphPattern(s)
	}
}

func (s *GraphGraphPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGraphGraphPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GraphGraphPattern() (localctx IGraphGraphPatternContext) {
	localctx = NewGraphGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 62, SparqlParserRULE_graphGraphPattern)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(413)
		p.Match(SparqlParserGRAPH)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(414)
		p.VarOrIRIref()
	}
	{
		p.SetState(415)
		p.GroupGraphPattern()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGroupOrUnionGraphPatternContext is an interface to support dynamic dispatch.
type IGroupOrUnionGraphPatternContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllGroupGraphPattern() []IGroupGraphPatternContext
	GroupGraphPattern(i int) IGroupGraphPatternContext
	AllUNION() []antlr.TerminalNode
	UNION(i int) antlr.TerminalNode

	// IsGroupOrUnionGraphPatternContext differentiates from other interfaces.
	IsGroupOrUnionGraphPatternContext()
}

type GroupOrUnionGraphPatternContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGroupOrUnionGraphPatternContext() *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern
	return p
}

func InitEmptyGroupOrUnionGraphPatternContext(p *GroupOrUnionGraphPatternContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern
}

func (*GroupOrUnionGraphPatternContext) IsGroupOrUnionGraphPatternContext() {}

func NewGroupOrUnionGraphPatternContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GroupOrUnionGraphPatternContext {
	var p = new(GroupOrUnionGraphPatternContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_groupOrUnionGraphPattern

	return p
}

func (s *GroupOrUnionGraphPatternContext) GetParser() antlr.Parser { return s.parser }

func (s *GroupOrUnionGraphPatternContext) AllGroupGraphPattern() []IGroupGraphPatternContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			len++
		}
	}

	tst := make([]IGroupGraphPatternContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGroupGraphPatternContext); ok {
			tst[i] = t.(IGroupGraphPatternContext)
			i++
		}
	}

	return tst
}

func (s *GroupOrUnionGraphPatternContext) GroupGraphPattern(i int) IGroupGraphPatternContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGroupGraphPatternContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGroupGraphPatternContext)
}

func (s *GroupOrUnionGraphPatternContext) AllUNION() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserUNION)
}

func (s *GroupOrUnionGraphPatternContext) UNION(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserUNION, i)
}

func (s *GroupOrUnionGraphPatternContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GroupOrUnionGraphPatternContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GroupOrUnionGraphPatternContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGroupOrUnionGraphPattern(s)
	}
}

func (s *GroupOrUnionGraphPatternContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGroupOrUnionGraphPattern(s)
	}
}

func (s *GroupOrUnionGraphPatternContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGroupOrUnionGraphPattern(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GroupOrUnionGraphPattern() (localctx IGroupOrUnionGraphPatternContext) {
	localctx = NewGroupOrUnionGraphPatternContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 64, SparqlParserRULE_groupOrUnionGraphPattern)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(417)
		p.GroupGraphPattern()
	}
	p.SetState(422)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserUNION {
		{
			p.SetState(418)
			p.Match(SparqlParserUNION)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(419)
			p.GroupGraphPattern()
		}

		p.SetState(424)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFilter_Context is an interface to support dynamic dispatch.
type IFilter_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	FILTER() antlr.TerminalNode
	Constraint() IConstraintContext

	// IsFilter_Context differentiates from other interfaces.
	IsFilter_Context()
}

type Filter_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFilter_Context() *Filter_Context {
	var p = new(Filter_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_filter_
	return p
}

func InitEmptyFilter_Context(p *Filter_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_filter_
}

func (*Filter_Context) IsFilter_Context() {}

func NewFilter_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Filter_Context {
	var p = new(Filter_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_filter_

	return p
}

func (s *Filter_Context) GetParser() antlr.Parser { return s.parser }

func (s *Filter_Context) FILTER() antlr.TerminalNode {
	return s.GetToken(SparqlParserFILTER, 0)
}

func (s *Filter_Context) Constraint() IConstraintContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstraintContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstraintContext)
}

func (s *Filter_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Filter_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Filter_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterFilter_(s)
	}
}

func (s *Filter_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitFilter_(s)
	}
}

func (s *Filter_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitFilter_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Filter_() (localctx IFilter_Context) {
	localctx = NewFilter_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 66, SparqlParserRULE_filter_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(425)
		p.Match(SparqlParserFILTER)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(426)
		p.Constraint()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstraintContext is an interface to support dynamic dispatch.
type IConstraintContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	BuiltInCall() IBuiltInCallContext
	FunctionCall() IFunctionCallContext

	// IsConstraintContext differentiates from other interfaces.
	IsConstraintContext()
}

type ConstraintContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstraintContext() *ConstraintContext {
	var p = new(ConstraintContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constraint
	return p
}

func InitEmptyConstraintContext(p *ConstraintContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constraint
}

func (*ConstraintContext) IsConstraintContext() {}

func NewConstraintContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstraintContext {
	var p = new(ConstraintContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constraint

	return p
}

func (s *ConstraintContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstraintContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *ConstraintContext) BuiltInCall() IBuiltInCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltInCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *ConstraintContext) FunctionCall() IFunctionCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IFunctionCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IFunctionCallContext)
}

func (s *ConstraintContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstraintContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstraintContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConstraint(s)
	}
}

func (s *ConstraintContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConstraint(s)
	}
}

func (s *ConstraintContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConstraint(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Constraint() (localctx IConstraintContext) {
	localctx = NewConstraintContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, SparqlParserRULE_constraint)
	p.SetState(431)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(428)
			p.BrackettedExpression()
		}

	case SparqlParserBOUND, SparqlParserDATATYPE, SparqlParserLANG, SparqlParserLANGMATCHES, SparqlParserREGEX, SparqlParserSTR, SparqlParserIS_LITERAL, SparqlParserIS_BLANK, SparqlParserIS_URI, SparqlParserIS_IRI, SparqlParserSAME_TERM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(429)
			p.BuiltInCall()
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(430)
			p.FunctionCall()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IFunctionCallContext is an interface to support dynamic dispatch.
type IFunctionCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext
	ArgList() IArgListContext

	// IsFunctionCallContext differentiates from other interfaces.
	IsFunctionCallContext()
}

type FunctionCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyFunctionCallContext() *FunctionCallContext {
	var p = new(FunctionCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_functionCall
	return p
}

func InitEmptyFunctionCallContext(p *FunctionCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_functionCall
}

func (*FunctionCallContext) IsFunctionCallContext() {}

func NewFunctionCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FunctionCallContext {
	var p = new(FunctionCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_functionCall

	return p
}

func (s *FunctionCallContext) GetParser() antlr.Parser { return s.parser }

func (s *FunctionCallContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *FunctionCallContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *FunctionCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FunctionCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FunctionCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterFunctionCall(s)
	}
}

func (s *FunctionCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitFunctionCall(s)
	}
}

func (s *FunctionCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitFunctionCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) FunctionCall() (localctx IFunctionCallContext) {
	localctx = NewFunctionCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 70, SparqlParserRULE_functionCall)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(433)
		p.IriRef()
	}
	{
		p.SetState(434)
		p.ArgList()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IArgListContext is an interface to support dynamic dispatch.
type IArgListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NIL() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_PAREN() antlr.TerminalNode
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsArgListContext differentiates from other interfaces.
	IsArgListContext()
}

type ArgListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyArgListContext() *ArgListContext {
	var p = new(ArgListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_argList
	return p
}

func InitEmptyArgListContext(p *ArgListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_argList
}

func (*ArgListContext) IsArgListContext() {}

func NewArgListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArgListContext {
	var p = new(ArgListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_argList

	return p
}

func (s *ArgListContext) GetParser() antlr.Parser { return s.parser }

func (s *ArgListContext) NIL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNIL, 0)
}

func (s *ArgListContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *ArgListContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ArgListContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ArgListContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *ArgListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserCOMMA)
}

func (s *ArgListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserCOMMA, i)
}

func (s *ArgListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArgListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArgListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterArgList(s)
	}
}

func (s *ArgListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitArgList(s)
	}
}

func (s *ArgListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitArgList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ArgList() (localctx IArgListContext) {
	localctx = NewArgListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, SparqlParserRULE_argList)
	var _la int

	p.SetState(448)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserNIL:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(436)
			p.Match(SparqlParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(437)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(438)
			p.Expression()
		}
		p.SetState(443)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		for _la == SparqlParserCOMMA {
			{
				p.SetState(439)
				p.Match(SparqlParserCOMMA)
				if p.HasError() {
					// Recognition error - abort rule
					goto errorExit
				}
			}
			{
				p.SetState(440)
				p.Expression()
			}

			p.SetState(445)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}
			_la = p.GetTokenStream().LA(1)
		}
		{
			p.SetState(446)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructTemplateContext is an interface to support dynamic dispatch.
type IConstructTemplateContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_CURLY() antlr.TerminalNode
	R_CURLY() antlr.TerminalNode
	ConstructTriples() IConstructTriplesContext

	// IsConstructTemplateContext differentiates from other interfaces.
	IsConstructTemplateContext()
}

type ConstructTemplateContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTemplateContext() *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTemplate
	return p
}

func InitEmptyConstructTemplateContext(p *ConstructTemplateContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTemplate
}

func (*ConstructTemplateContext) IsConstructTemplateContext() {}

func NewConstructTemplateContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTemplateContext {
	var p = new(ConstructTemplateContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTemplate

	return p
}

func (s *ConstructTemplateContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTemplateContext) L_CURLY() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_CURLY, 0)
}

func (s *ConstructTemplateContext) R_CURLY() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_CURLY, 0)
}

func (s *ConstructTemplateContext) ConstructTriples() IConstructTriplesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTriplesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTriplesContext)
}

func (s *ConstructTemplateContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructTemplateContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructTemplateContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConstructTemplate(s)
	}
}

func (s *ConstructTemplateContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConstructTemplate(s)
	}
}

func (s *ConstructTemplateContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConstructTemplate(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ConstructTemplate() (localctx IConstructTemplateContext) {
	localctx = NewConstructTemplateContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 74, SparqlParserRULE_constructTemplate)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(450)
		p.Match(SparqlParserL_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(452)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0 {
		{
			p.SetState(451)
			p.ConstructTriples()
		}

	}
	{
		p.SetState(454)
		p.Match(SparqlParserR_CURLY)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConstructTriplesContext is an interface to support dynamic dispatch.
type IConstructTriplesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TriplesSameSubject() ITriplesSameSubjectContext
	DOT() antlr.TerminalNode
	ConstructTriples() IConstructTriplesContext

	// IsConstructTriplesContext differentiates from other interfaces.
	IsConstructTriplesContext()
}

type ConstructTriplesContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConstructTriplesContext() *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTriples
	return p
}

func InitEmptyConstructTriplesContext(p *ConstructTriplesContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_constructTriples
}

func (*ConstructTriplesContext) IsConstructTriplesContext() {}

func NewConstructTriplesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConstructTriplesContext {
	var p = new(ConstructTriplesContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_constructTriples

	return p
}

func (s *ConstructTriplesContext) GetParser() antlr.Parser { return s.parser }

func (s *ConstructTriplesContext) TriplesSameSubject() ITriplesSameSubjectContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesSameSubjectContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesSameSubjectContext)
}

func (s *ConstructTriplesContext) DOT() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOT, 0)
}

func (s *ConstructTriplesContext) ConstructTriples() IConstructTriplesContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConstructTriplesContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConstructTriplesContext)
}

func (s *ConstructTriplesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConstructTriplesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConstructTriplesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConstructTriples(s)
	}
}

func (s *ConstructTriplesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConstructTriples(s)
	}
}

func (s *ConstructTriplesContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConstructTriples(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ConstructTriples() (localctx IConstructTriplesContext) {
	localctx = NewConstructTriplesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 76, SparqlParserRULE_constructTriples)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(456)
		p.TriplesSameSubject()
	}
	p.SetState(461)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserDOT {
		{
			p.SetState(457)
			p.Match(SparqlParserDOT)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(459)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if (int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0 {
			{
				p.SetState(458)
				p.ConstructTriples()
			}

		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesSameSubjectContext is an interface to support dynamic dispatch.
type ITriplesSameSubjectContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarOrTerm() IVarOrTermContext
	PropertyListNotEmpty() IPropertyListNotEmptyContext
	TriplesNode() ITriplesNodeContext
	PropertyList() IPropertyListContext

	// IsTriplesSameSubjectContext differentiates from other interfaces.
	IsTriplesSameSubjectContext()
}

type TriplesSameSubjectContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesSameSubjectContext() *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesSameSubject
	return p
}

func InitEmptyTriplesSameSubjectContext(p *TriplesSameSubjectContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesSameSubject
}

func (*TriplesSameSubjectContext) IsTriplesSameSubjectContext() {}

func NewTriplesSameSubjectContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesSameSubjectContext {
	var p = new(TriplesSameSubjectContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesSameSubject

	return p
}

func (s *TriplesSameSubjectContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesSameSubjectContext) VarOrTerm() IVarOrTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *TriplesSameSubjectContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *TriplesSameSubjectContext) TriplesNode() ITriplesNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesNodeContext)
}

func (s *TriplesSameSubjectContext) PropertyList() IPropertyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListContext)
}

func (s *TriplesSameSubjectContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesSameSubjectContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesSameSubjectContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterTriplesSameSubject(s)
	}
}

func (s *TriplesSameSubjectContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitTriplesSameSubject(s)
	}
}

func (s *TriplesSameSubjectContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitTriplesSameSubject(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) TriplesSameSubject() (localctx ITriplesSameSubjectContext) {
	localctx = NewTriplesSameSubjectContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 78, SparqlParserRULE_triplesSameSubject)
	p.SetState(469)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserTRUE, SparqlParserFALSE, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(463)
			p.VarOrTerm()
		}
		{
			p.SetState(464)
			p.PropertyListNotEmpty()
		}

	case SparqlParserL_PAREN, SparqlParserL_SQUARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(466)
			p.TriplesNode()
		}
		{
			p.SetState(467)
			p.PropertyList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyListNotEmptyContext is an interface to support dynamic dispatch.
type IPropertyListNotEmptyContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllVerb() []IVerbContext
	Verb(i int) IVerbContext
	AllObjectList() []IObjectListContext
	ObjectList(i int) IObjectListContext
	AllSEMICOLON() []antlr.TerminalNode
	SEMICOLON(i int) antlr.TerminalNode

	// IsPropertyListNotEmptyContext differentiates from other interfaces.
	IsPropertyListNotEmptyContext()
}

type PropertyListNotEmptyContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyListNotEmptyContext() *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty
	return p
}

func InitEmptyPropertyListNotEmptyContext(p *PropertyListNotEmptyContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty
}

func (*PropertyListNotEmptyContext) IsPropertyListNotEmptyContext() {}

func NewPropertyListNotEmptyContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListNotEmptyContext {
	var p = new(PropertyListNotEmptyContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_propertyListNotEmpty

	return p
}

func (s *PropertyListNotEmptyContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListNotEmptyContext) AllVerb() []IVerbContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IVerbContext); ok {
			len++
		}
	}

	tst := make([]IVerbContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IVerbContext); ok {
			tst[i] = t.(IVerbContext)
			i++
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) Verb(i int) IVerbContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVerbContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVerbContext)
}

func (s *PropertyListNotEmptyContext) AllObjectList() []IObjectListContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObjectListContext); ok {
			len++
		}
	}

	tst := make([]IObjectListContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObjectListContext); ok {
			tst[i] = t.(IObjectListContext)
			i++
		}
	}

	return tst
}

func (s *PropertyListNotEmptyContext) ObjectList(i int) IObjectListContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObjectListContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObjectListContext)
}

func (s *PropertyListNotEmptyContext) AllSEMICOLON() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserSEMICOLON)
}

func (s *PropertyListNotEmptyContext) SEMICOLON(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserSEMICOLON, i)
}

func (s *PropertyListNotEmptyContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyListNotEmptyContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyListNotEmptyContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPropertyListNotEmpty(s)
	}
}

func (s *PropertyListNotEmptyContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPropertyListNotEmpty(s)
	}
}

func (s *PropertyListNotEmptyContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPropertyListNotEmpty(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) PropertyListNotEmpty() (localctx IPropertyListNotEmptyContext) {
	localctx = NewPropertyListNotEmptyContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 80, SparqlParserRULE_propertyListNotEmpty)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(471)
		p.Verb()
	}
	{
		p.SetState(472)
		p.ObjectList()
	}
	p.SetState(481)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserSEMICOLON {
		{
			p.SetState(473)
			p.Match(SparqlParserSEMICOLON)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		p.SetState(477)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)

		if _la == SparqlParserA || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&55) != 0) {
			{
				p.SetState(474)
				p.Verb()
			}
			{
				p.SetState(475)
				p.ObjectList()
			}

		}

		p.SetState(483)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPropertyListContext is an interface to support dynamic dispatch.
type IPropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PropertyListNotEmpty() IPropertyListNotEmptyContext

	// IsPropertyListContext differentiates from other interfaces.
	IsPropertyListContext()
}

type PropertyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPropertyListContext() *PropertyListContext {
	var p = new(PropertyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyList
	return p
}

func InitEmptyPropertyListContext(p *PropertyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_propertyList
}

func (*PropertyListContext) IsPropertyListContext() {}

func NewPropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PropertyListContext {
	var p = new(PropertyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_propertyList

	return p
}

func (s *PropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *PropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *PropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPropertyList(s)
	}
}

func (s *PropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPropertyList(s)
	}
}

func (s *PropertyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPropertyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) PropertyList() (localctx IPropertyListContext) {
	localctx = NewPropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 82, SparqlParserRULE_propertyList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(485)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserA || ((int64((_la-68)) & ^0x3f) == 0 && ((int64(1)<<(_la-68))&55) != 0) {
		{
			p.SetState(484)
			p.PropertyListNotEmpty()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObjectListContext is an interface to support dynamic dispatch.
type IObjectListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllObject_() []IObject_Context
	Object_(i int) IObject_Context
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode

	// IsObjectListContext differentiates from other interfaces.
	IsObjectListContext()
}

type ObjectListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObjectListContext() *ObjectListContext {
	var p = new(ObjectListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_objectList
	return p
}

func InitEmptyObjectListContext(p *ObjectListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_objectList
}

func (*ObjectListContext) IsObjectListContext() {}

func NewObjectListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ObjectListContext {
	var p = new(ObjectListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_objectList

	return p
}

func (s *ObjectListContext) GetParser() antlr.Parser { return s.parser }

func (s *ObjectListContext) AllObject_() []IObject_Context {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IObject_Context); ok {
			len++
		}
	}

	tst := make([]IObject_Context, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IObject_Context); ok {
			tst[i] = t.(IObject_Context)
			i++
		}
	}

	return tst
}

func (s *ObjectListContext) Object_(i int) IObject_Context {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IObject_Context); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IObject_Context)
}

func (s *ObjectListContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserCOMMA)
}

func (s *ObjectListContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserCOMMA, i)
}

func (s *ObjectListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ObjectListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ObjectListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterObjectList(s)
	}
}

func (s *ObjectListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitObjectList(s)
	}
}

func (s *ObjectListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitObjectList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ObjectList() (localctx IObjectListContext) {
	localctx = NewObjectListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 84, SparqlParserRULE_objectList)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(487)
		p.Object_()
	}
	p.SetState(492)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserCOMMA {
		{
			p.SetState(488)
			p.Match(SparqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(489)
			p.Object_()
		}

		p.SetState(494)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IObject_Context is an interface to support dynamic dispatch.
type IObject_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	GraphNode() IGraphNodeContext

	// IsObject_Context differentiates from other interfaces.
	IsObject_Context()
}

type Object_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyObject_Context() *Object_Context {
	var p = new(Object_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_object_
	return p
}

func InitEmptyObject_Context(p *Object_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_object_
}

func (*Object_Context) IsObject_Context() {}

func NewObject_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Object_Context {
	var p = new(Object_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_object_

	return p
}

func (s *Object_Context) GetParser() antlr.Parser { return s.parser }

func (s *Object_Context) GraphNode() IGraphNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphNodeContext)
}

func (s *Object_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Object_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Object_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterObject_(s)
	}
}

func (s *Object_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitObject_(s)
	}
}

func (s *Object_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitObject_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Object_() (localctx IObject_Context) {
	localctx = NewObject_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 86, SparqlParserRULE_object_)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(495)
		p.GraphNode()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVerbContext is an interface to support dynamic dispatch.
type IVerbContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarOrIRIref() IVarOrIRIrefContext
	A() antlr.TerminalNode

	// IsVerbContext differentiates from other interfaces.
	IsVerbContext()
}

type VerbContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVerbContext() *VerbContext {
	var p = new(VerbContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_verb
	return p
}

func InitEmptyVerbContext(p *VerbContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_verb
}

func (*VerbContext) IsVerbContext() {}

func NewVerbContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VerbContext {
	var p = new(VerbContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_verb

	return p
}

func (s *VerbContext) GetParser() antlr.Parser { return s.parser }

func (s *VerbContext) VarOrIRIref() IVarOrIRIrefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrIRIrefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrIRIrefContext)
}

func (s *VerbContext) A() antlr.TerminalNode {
	return s.GetToken(SparqlParserA, 0)
}

func (s *VerbContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VerbContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VerbContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterVerb(s)
	}
}

func (s *VerbContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitVerb(s)
	}
}

func (s *VerbContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitVerb(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Verb() (localctx IVerbContext) {
	localctx = NewVerbContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 88, SparqlParserRULE_verb)
	p.SetState(499)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(497)
			p.VarOrIRIref()
		}

	case SparqlParserA:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(498)
			p.Match(SparqlParserA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ITriplesNodeContext is an interface to support dynamic dispatch.
type ITriplesNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Collection() ICollectionContext
	BlankNodePropertyList() IBlankNodePropertyListContext

	// IsTriplesNodeContext differentiates from other interfaces.
	IsTriplesNodeContext()
}

type TriplesNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyTriplesNodeContext() *TriplesNodeContext {
	var p = new(TriplesNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesNode
	return p
}

func InitEmptyTriplesNodeContext(p *TriplesNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_triplesNode
}

func (*TriplesNodeContext) IsTriplesNodeContext() {}

func NewTriplesNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TriplesNodeContext {
	var p = new(TriplesNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_triplesNode

	return p
}

func (s *TriplesNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *TriplesNodeContext) Collection() ICollectionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ICollectionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ICollectionContext)
}

func (s *TriplesNodeContext) BlankNodePropertyList() IBlankNodePropertyListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlankNodePropertyListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlankNodePropertyListContext)
}

func (s *TriplesNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TriplesNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TriplesNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterTriplesNode(s)
	}
}

func (s *TriplesNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitTriplesNode(s)
	}
}

func (s *TriplesNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitTriplesNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) TriplesNode() (localctx ITriplesNodeContext) {
	localctx = NewTriplesNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 90, SparqlParserRULE_triplesNode)
	p.SetState(503)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(501)
			p.Collection()
		}

	case SparqlParserL_SQUARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(502)
			p.BlankNodePropertyList()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlankNodePropertyListContext is an interface to support dynamic dispatch.
type IBlankNodePropertyListContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_SQUARE() antlr.TerminalNode
	PropertyListNotEmpty() IPropertyListNotEmptyContext
	R_SQUARE() antlr.TerminalNode

	// IsBlankNodePropertyListContext differentiates from other interfaces.
	IsBlankNodePropertyListContext()
}

type BlankNodePropertyListContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodePropertyListContext() *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList
	return p
}

func InitEmptyBlankNodePropertyListContext(p *BlankNodePropertyListContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList
}

func (*BlankNodePropertyListContext) IsBlankNodePropertyListContext() {}

func NewBlankNodePropertyListContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodePropertyListContext {
	var p = new(BlankNodePropertyListContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_blankNodePropertyList

	return p
}

func (s *BlankNodePropertyListContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodePropertyListContext) L_SQUARE() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_SQUARE, 0)
}

func (s *BlankNodePropertyListContext) PropertyListNotEmpty() IPropertyListNotEmptyContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPropertyListNotEmptyContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPropertyListNotEmptyContext)
}

func (s *BlankNodePropertyListContext) R_SQUARE() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_SQUARE, 0)
}

func (s *BlankNodePropertyListContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodePropertyListContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankNodePropertyListContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBlankNodePropertyList(s)
	}
}

func (s *BlankNodePropertyListContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBlankNodePropertyList(s)
	}
}

func (s *BlankNodePropertyListContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBlankNodePropertyList(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BlankNodePropertyList() (localctx IBlankNodePropertyListContext) {
	localctx = NewBlankNodePropertyListContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 92, SparqlParserRULE_blankNodePropertyList)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(505)
		p.Match(SparqlParserL_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(506)
		p.PropertyListNotEmpty()
	}
	{
		p.SetState(507)
		p.Match(SparqlParserR_SQUARE)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// ICollectionContext is an interface to support dynamic dispatch.
type ICollectionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	R_PAREN() antlr.TerminalNode
	AllGraphNode() []IGraphNodeContext
	GraphNode(i int) IGraphNodeContext

	// IsCollectionContext differentiates from other interfaces.
	IsCollectionContext()
}

type CollectionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyCollectionContext() *CollectionContext {
	var p = new(CollectionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_collection
	return p
}

func InitEmptyCollectionContext(p *CollectionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_collection
}

func (*CollectionContext) IsCollectionContext() {}

func NewCollectionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CollectionContext {
	var p = new(CollectionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_collection

	return p
}

func (s *CollectionContext) GetParser() antlr.Parser { return s.parser }

func (s *CollectionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *CollectionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *CollectionContext) AllGraphNode() []IGraphNodeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IGraphNodeContext); ok {
			len++
		}
	}

	tst := make([]IGraphNodeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IGraphNodeContext); ok {
			tst[i] = t.(IGraphNodeContext)
			i++
		}
	}

	return tst
}

func (s *CollectionContext) GraphNode(i int) IGraphNodeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphNodeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphNodeContext)
}

func (s *CollectionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CollectionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CollectionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterCollection(s)
	}
}

func (s *CollectionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitCollection(s)
	}
}

func (s *CollectionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitCollection(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Collection() (localctx ICollectionContext) {
	localctx = NewCollectionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 94, SparqlParserRULE_collection)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(509)
		p.Match(SparqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	p.SetState(511)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for ok := true; ok; ok = ((int64((_la-38)) & ^0x3f) == 0 && ((int64(1)<<(_la-38))&14003310299709443) != 0) {
		{
			p.SetState(510)
			p.GraphNode()
		}

		p.SetState(513)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}
	{
		p.SetState(515)
		p.Match(SparqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphNodeContext is an interface to support dynamic dispatch.
type IGraphNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VarOrTerm() IVarOrTermContext
	TriplesNode() ITriplesNodeContext

	// IsGraphNodeContext differentiates from other interfaces.
	IsGraphNodeContext()
}

type GraphNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphNodeContext() *GraphNodeContext {
	var p = new(GraphNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphNode
	return p
}

func InitEmptyGraphNodeContext(p *GraphNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphNode
}

func (*GraphNodeContext) IsGraphNodeContext() {}

func NewGraphNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphNodeContext {
	var p = new(GraphNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphNode

	return p
}

func (s *GraphNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphNodeContext) VarOrTerm() IVarOrTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVarOrTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVarOrTermContext)
}

func (s *GraphNodeContext) TriplesNode() ITriplesNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(ITriplesNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(ITriplesNodeContext)
}

func (s *GraphNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGraphNode(s)
	}
}

func (s *GraphNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGraphNode(s)
	}
}

func (s *GraphNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGraphNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GraphNode() (localctx IGraphNodeContext) {
	localctx = NewGraphNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 96, SparqlParserRULE_graphNode)
	p.SetState(519)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserTRUE, SparqlParserFALSE, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(517)
			p.VarOrTerm()
		}

	case SparqlParserL_PAREN, SparqlParserL_SQUARE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(518)
			p.TriplesNode()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarOrTermContext is an interface to support dynamic dispatch.
type IVarOrTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVar_Context
	GraphTerm() IGraphTermContext

	// IsVarOrTermContext differentiates from other interfaces.
	IsVarOrTermContext()
}

type VarOrTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrTermContext() *VarOrTermContext {
	var p = new(VarOrTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrTerm
	return p
}

func InitEmptyVarOrTermContext(p *VarOrTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrTerm
}

func (*VarOrTermContext) IsVarOrTermContext() {}

func NewVarOrTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrTermContext {
	var p = new(VarOrTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_varOrTerm

	return p
}

func (s *VarOrTermContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrTermContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *VarOrTermContext) GraphTerm() IGraphTermContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IGraphTermContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IGraphTermContext)
}

func (s *VarOrTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterVarOrTerm(s)
	}
}

func (s *VarOrTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitVarOrTerm(s)
	}
}

func (s *VarOrTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitVarOrTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) VarOrTerm() (localctx IVarOrTermContext) {
	localctx = NewVarOrTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 98, SparqlParserRULE_varOrTerm)
	p.SetState(523)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(521)
			p.Var_()
		}

	case SparqlParserTRUE, SparqlParserFALSE, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(522)
			p.GraphTerm()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVarOrIRIrefContext is an interface to support dynamic dispatch.
type IVarOrIRIrefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	Var_() IVar_Context
	IriRef() IIriRefContext

	// IsVarOrIRIrefContext differentiates from other interfaces.
	IsVarOrIRIrefContext()
}

type VarOrIRIrefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVarOrIRIrefContext() *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrIRIref
	return p
}

func InitEmptyVarOrIRIrefContext(p *VarOrIRIrefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_varOrIRIref
}

func (*VarOrIRIrefContext) IsVarOrIRIrefContext() {}

func NewVarOrIRIrefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VarOrIRIrefContext {
	var p = new(VarOrIRIrefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_varOrIRIref

	return p
}

func (s *VarOrIRIrefContext) GetParser() antlr.Parser { return s.parser }

func (s *VarOrIRIrefContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *VarOrIRIrefContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *VarOrIRIrefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VarOrIRIrefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VarOrIRIrefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterVarOrIRIref(s)
	}
}

func (s *VarOrIRIrefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitVarOrIRIref(s)
	}
}

func (s *VarOrIRIrefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitVarOrIRIref(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) VarOrIRIref() (localctx IVarOrIRIrefContext) {
	localctx = NewVarOrIRIrefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 100, SparqlParserRULE_varOrIRIref)
	p.SetState(527)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(525)
			p.Var_()
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(526)
			p.IriRef()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IVar_Context is an interface to support dynamic dispatch.
type IVar_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	VAR1() antlr.TerminalNode
	VAR2() antlr.TerminalNode

	// IsVar_Context differentiates from other interfaces.
	IsVar_Context()
}

type Var_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyVar_Context() *Var_Context {
	var p = new(Var_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_var_
	return p
}

func InitEmptyVar_Context(p *Var_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_var_
}

func (*Var_Context) IsVar_Context() {}

func NewVar_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Var_Context {
	var p = new(Var_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_var_

	return p
}

func (s *Var_Context) GetParser() antlr.Parser { return s.parser }

func (s *Var_Context) VAR1() antlr.TerminalNode {
	return s.GetToken(SparqlParserVAR1, 0)
}

func (s *Var_Context) VAR2() antlr.TerminalNode {
	return s.GetToken(SparqlParserVAR2, 0)
}

func (s *Var_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Var_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Var_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterVar_(s)
	}
}

func (s *Var_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitVar_(s)
	}
}

func (s *Var_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitVar_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Var_() (localctx IVar_Context) {
	localctx = NewVar_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 102, SparqlParserRULE_var_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(529)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserVAR1 || _la == SparqlParserVAR2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IGraphTermContext is an interface to support dynamic dispatch.
type IGraphTermContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext
	RdfLiteral() IRdfLiteralContext
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	BlankNode() IBlankNodeContext
	NIL() antlr.TerminalNode

	// IsGraphTermContext differentiates from other interfaces.
	IsGraphTermContext()
}

type GraphTermContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyGraphTermContext() *GraphTermContext {
	var p = new(GraphTermContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphTerm
	return p
}

func InitEmptyGraphTermContext(p *GraphTermContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_graphTerm
}

func (*GraphTermContext) IsGraphTermContext() {}

func NewGraphTermContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *GraphTermContext {
	var p = new(GraphTermContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_graphTerm

	return p
}

func (s *GraphTermContext) GetParser() antlr.Parser { return s.parser }

func (s *GraphTermContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *GraphTermContext) RdfLiteral() IRdfLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRdfLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *GraphTermContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *GraphTermContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *GraphTermContext) BlankNode() IBlankNodeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBlankNodeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBlankNodeContext)
}

func (s *GraphTermContext) NIL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNIL, 0)
}

func (s *GraphTermContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *GraphTermContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *GraphTermContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterGraphTerm(s)
	}
}

func (s *GraphTermContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitGraphTerm(s)
	}
}

func (s *GraphTermContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitGraphTerm(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) GraphTerm() (localctx IGraphTermContext) {
	localctx = NewGraphTermContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 104, SparqlParserRULE_graphTerm)
	p.SetState(537)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(531)
			p.IriRef()
		}

	case SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(532)
			p.RdfLiteral()
		}

	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(533)
			p.NumericLiteral()
		}

	case SparqlParserTRUE, SparqlParserFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(534)
			p.BooleanLiteral()
		}

	case SparqlParserBLANK_NODE_LABEL, SparqlParserANON:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(535)
			p.BlankNode()
		}

	case SparqlParserNIL:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(536)
			p.Match(SparqlParserNIL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	ConditionalOrExpression() IConditionalOrExpressionContext

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_expression
	return p
}

func InitEmptyExpressionContext(p *ExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_expression
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) ConditionalOrExpression() IConditionalOrExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalOrExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalOrExpressionContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (s *ExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) Expression() (localctx IExpressionContext) {
	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 106, SparqlParserRULE_expression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(539)
		p.ConditionalOrExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalOrExpressionContext is an interface to support dynamic dispatch.
type IConditionalOrExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllConditionalAndExpression() []IConditionalAndExpressionContext
	ConditionalAndExpression(i int) IConditionalAndExpressionContext
	AllDOUBLE_BAR() []antlr.TerminalNode
	DOUBLE_BAR(i int) antlr.TerminalNode

	// IsConditionalOrExpressionContext differentiates from other interfaces.
	IsConditionalOrExpressionContext()
}

type ConditionalOrExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalOrExpressionContext() *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression
	return p
}

func InitEmptyConditionalOrExpressionContext(p *ConditionalOrExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression
}

func (*ConditionalOrExpressionContext) IsConditionalOrExpressionContext() {}

func NewConditionalOrExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalOrExpressionContext {
	var p = new(ConditionalOrExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalOrExpression

	return p
}

func (s *ConditionalOrExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalOrExpressionContext) AllConditionalAndExpression() []IConditionalAndExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
			len++
		}
	}

	tst := make([]IConditionalAndExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IConditionalAndExpressionContext); ok {
			tst[i] = t.(IConditionalAndExpressionContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalOrExpressionContext) ConditionalAndExpression(i int) IConditionalAndExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IConditionalAndExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IConditionalAndExpressionContext)
}

func (s *ConditionalOrExpressionContext) AllDOUBLE_BAR() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserDOUBLE_BAR)
}

func (s *ConditionalOrExpressionContext) DOUBLE_BAR(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_BAR, i)
}

func (s *ConditionalOrExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalOrExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalOrExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConditionalOrExpression(s)
	}
}

func (s *ConditionalOrExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConditionalOrExpression(s)
	}
}

func (s *ConditionalOrExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConditionalOrExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ConditionalOrExpression() (localctx IConditionalOrExpressionContext) {
	localctx = NewConditionalOrExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 108, SparqlParserRULE_conditionalOrExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(541)
		p.ConditionalAndExpression()
	}
	p.SetState(546)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserDOUBLE_BAR {
		{
			p.SetState(542)
			p.Match(SparqlParserDOUBLE_BAR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(543)
			p.ConditionalAndExpression()
		}

		p.SetState(548)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IConditionalAndExpressionContext is an interface to support dynamic dispatch.
type IConditionalAndExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllValueLogical() []IValueLogicalContext
	ValueLogical(i int) IValueLogicalContext
	AllDOUBLE_AMP() []antlr.TerminalNode
	DOUBLE_AMP(i int) antlr.TerminalNode

	// IsConditionalAndExpressionContext differentiates from other interfaces.
	IsConditionalAndExpressionContext()
}

type ConditionalAndExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyConditionalAndExpressionContext() *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression
	return p
}

func InitEmptyConditionalAndExpressionContext(p *ConditionalAndExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression
}

func (*ConditionalAndExpressionContext) IsConditionalAndExpressionContext() {}

func NewConditionalAndExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionalAndExpressionContext {
	var p = new(ConditionalAndExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_conditionalAndExpression

	return p
}

func (s *ConditionalAndExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionalAndExpressionContext) AllValueLogical() []IValueLogicalContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IValueLogicalContext); ok {
			len++
		}
	}

	tst := make([]IValueLogicalContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IValueLogicalContext); ok {
			tst[i] = t.(IValueLogicalContext)
			i++
		}
	}

	return tst
}

func (s *ConditionalAndExpressionContext) ValueLogical(i int) IValueLogicalContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IValueLogicalContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IValueLogicalContext)
}

func (s *ConditionalAndExpressionContext) AllDOUBLE_AMP() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserDOUBLE_AMP)
}

func (s *ConditionalAndExpressionContext) DOUBLE_AMP(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_AMP, i)
}

func (s *ConditionalAndExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionalAndExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionalAndExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterConditionalAndExpression(s)
	}
}

func (s *ConditionalAndExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitConditionalAndExpression(s)
	}
}

func (s *ConditionalAndExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitConditionalAndExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ConditionalAndExpression() (localctx IConditionalAndExpressionContext) {
	localctx = NewConditionalAndExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 110, SparqlParserRULE_conditionalAndExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(549)
		p.ValueLogical()
	}
	p.SetState(554)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserDOUBLE_AMP {
		{
			p.SetState(550)
			p.Match(SparqlParserDOUBLE_AMP)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(551)
			p.ValueLogical()
		}

		p.SetState(556)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IValueLogicalContext is an interface to support dynamic dispatch.
type IValueLogicalContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	RelationalExpression() IRelationalExpressionContext

	// IsValueLogicalContext differentiates from other interfaces.
	IsValueLogicalContext()
}

type ValueLogicalContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyValueLogicalContext() *ValueLogicalContext {
	var p = new(ValueLogicalContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_valueLogical
	return p
}

func InitEmptyValueLogicalContext(p *ValueLogicalContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_valueLogical
}

func (*ValueLogicalContext) IsValueLogicalContext() {}

func NewValueLogicalContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ValueLogicalContext {
	var p = new(ValueLogicalContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_valueLogical

	return p
}

func (s *ValueLogicalContext) GetParser() antlr.Parser { return s.parser }

func (s *ValueLogicalContext) RelationalExpression() IRelationalExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRelationalExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRelationalExpressionContext)
}

func (s *ValueLogicalContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ValueLogicalContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ValueLogicalContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterValueLogical(s)
	}
}

func (s *ValueLogicalContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitValueLogical(s)
	}
}

func (s *ValueLogicalContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitValueLogical(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) ValueLogical() (localctx IValueLogicalContext) {
	localctx = NewValueLogicalContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 112, SparqlParserRULE_valueLogical)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(557)
		p.RelationalExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRelationalExpressionContext is an interface to support dynamic dispatch.
type IRelationalExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllNumericExpression() []INumericExpressionContext
	NumericExpression(i int) INumericExpressionContext
	EQUAL() antlr.TerminalNode
	NOT_EQUAL() antlr.TerminalNode
	LESS() antlr.TerminalNode
	GREATER() antlr.TerminalNode
	LESS_OR_EQUAL() antlr.TerminalNode
	GREATER_OR_EQUAL() antlr.TerminalNode

	// IsRelationalExpressionContext differentiates from other interfaces.
	IsRelationalExpressionContext()
}

type RelationalExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRelationalExpressionContext() *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_relationalExpression
	return p
}

func InitEmptyRelationalExpressionContext(p *RelationalExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_relationalExpression
}

func (*RelationalExpressionContext) IsRelationalExpressionContext() {}

func NewRelationalExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RelationalExpressionContext {
	var p = new(RelationalExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_relationalExpression

	return p
}

func (s *RelationalExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RelationalExpressionContext) AllNumericExpression() []INumericExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericExpressionContext); ok {
			len++
		}
	}

	tst := make([]INumericExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericExpressionContext); ok {
			tst[i] = t.(INumericExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RelationalExpressionContext) NumericExpression(i int) INumericExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericExpressionContext)
}

func (s *RelationalExpressionContext) EQUAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserEQUAL, 0)
}

func (s *RelationalExpressionContext) NOT_EQUAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserNOT_EQUAL, 0)
}

func (s *RelationalExpressionContext) LESS() antlr.TerminalNode {
	return s.GetToken(SparqlParserLESS, 0)
}

func (s *RelationalExpressionContext) GREATER() antlr.TerminalNode {
	return s.GetToken(SparqlParserGREATER, 0)
}

func (s *RelationalExpressionContext) LESS_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserLESS_OR_EQUAL, 0)
}

func (s *RelationalExpressionContext) GREATER_OR_EQUAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserGREATER_OR_EQUAL, 0)
}

func (s *RelationalExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RelationalExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RelationalExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitRelationalExpression(s)
	}
}

func (s *RelationalExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitRelationalExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) RelationalExpression() (localctx IRelationalExpressionContext) {
	localctx = NewRelationalExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 114, SparqlParserRULE_relationalExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(559)
		p.NumericExpression()
	}
	p.SetState(562)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1221601398924247040) != 0 {
		{
			p.SetState(560)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&1221601398924247040) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(561)
			p.NumericExpression()
		}

	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericExpressionContext is an interface to support dynamic dispatch.
type INumericExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AdditiveExpression() IAdditiveExpressionContext

	// IsNumericExpressionContext differentiates from other interfaces.
	IsNumericExpressionContext()
}

type NumericExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericExpressionContext() *NumericExpressionContext {
	var p = new(NumericExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericExpression
	return p
}

func InitEmptyNumericExpressionContext(p *NumericExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericExpression
}

func (*NumericExpressionContext) IsNumericExpressionContext() {}

func NewNumericExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericExpressionContext {
	var p = new(NumericExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericExpression

	return p
}

func (s *NumericExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericExpressionContext) AdditiveExpression() IAdditiveExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IAdditiveExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IAdditiveExpressionContext)
}

func (s *NumericExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNumericExpression(s)
	}
}

func (s *NumericExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNumericExpression(s)
	}
}

func (s *NumericExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNumericExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NumericExpression() (localctx INumericExpressionContext) {
	localctx = NewNumericExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 116, SparqlParserRULE_numericExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(564)
		p.AdditiveExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IAdditiveExpressionContext is an interface to support dynamic dispatch.
type IAdditiveExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllMultiplicativeExpression() []IMultiplicativeExpressionContext
	MultiplicativeExpression(i int) IMultiplicativeExpressionContext
	AllNumericLiteralPositive() []INumericLiteralPositiveContext
	NumericLiteralPositive(i int) INumericLiteralPositiveContext
	AllNumericLiteralNegative() []INumericLiteralNegativeContext
	NumericLiteralNegative(i int) INumericLiteralNegativeContext
	AllPLUS() []antlr.TerminalNode
	PLUS(i int) antlr.TerminalNode
	AllMINUS() []antlr.TerminalNode
	MINUS(i int) antlr.TerminalNode

	// IsAdditiveExpressionContext differentiates from other interfaces.
	IsAdditiveExpressionContext()
}

type AdditiveExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyAdditiveExpressionContext() *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_additiveExpression
	return p
}

func InitEmptyAdditiveExpressionContext(p *AdditiveExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_additiveExpression
}

func (*AdditiveExpressionContext) IsAdditiveExpressionContext() {}

func NewAdditiveExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *AdditiveExpressionContext {
	var p = new(AdditiveExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_additiveExpression

	return p
}

func (s *AdditiveExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *AdditiveExpressionContext) AllMultiplicativeExpression() []IMultiplicativeExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			len++
		}
	}

	tst := make([]IMultiplicativeExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IMultiplicativeExpressionContext); ok {
			tst[i] = t.(IMultiplicativeExpressionContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) MultiplicativeExpression(i int) IMultiplicativeExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IMultiplicativeExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IMultiplicativeExpressionContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralPositive() []INumericLiteralPositiveContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
			len++
		}
	}

	tst := make([]INumericLiteralPositiveContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericLiteralPositiveContext); ok {
			tst[i] = t.(INumericLiteralPositiveContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralPositive(i int) INumericLiteralPositiveContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralPositiveContext)
}

func (s *AdditiveExpressionContext) AllNumericLiteralNegative() []INumericLiteralNegativeContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
			len++
		}
	}

	tst := make([]INumericLiteralNegativeContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(INumericLiteralNegativeContext); ok {
			tst[i] = t.(INumericLiteralNegativeContext)
			i++
		}
	}

	return tst
}

func (s *AdditiveExpressionContext) NumericLiteralNegative(i int) INumericLiteralNegativeContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralNegativeContext)
}

func (s *AdditiveExpressionContext) AllPLUS() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserPLUS)
}

func (s *AdditiveExpressionContext) PLUS(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserPLUS, i)
}

func (s *AdditiveExpressionContext) AllMINUS() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserMINUS)
}

func (s *AdditiveExpressionContext) MINUS(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserMINUS, i)
}

func (s *AdditiveExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *AdditiveExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *AdditiveExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitAdditiveExpression(s)
	}
}

func (s *AdditiveExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitAdditiveExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) AdditiveExpression() (localctx IAdditiveExpressionContext) {
	localctx = NewAdditiveExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 118, SparqlParserRULE_additiveExpression)
	var _la int

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(566)
		p.MultiplicativeExpression()
	}
	p.SetState(573)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
	if p.HasError() {
		goto errorExit
	}
	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			p.SetState(571)
			p.GetErrorHandler().Sync(p)
			if p.HasError() {
				goto errorExit
			}

			switch p.GetTokenStream().LA(1) {
			case SparqlParserMINUS, SparqlParserPLUS:
				{
					p.SetState(567)
					_la = p.GetTokenStream().LA(1)

					if !(_la == SparqlParserMINUS || _la == SparqlParserPLUS) {
						p.GetErrorHandler().RecoverInline(p)
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(568)
					p.MultiplicativeExpression()
				}

			case SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE:
				{
					p.SetState(569)
					p.NumericLiteralPositive()
				}

			case SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
				{
					p.SetState(570)
					p.NumericLiteralNegative()
				}

			default:
				p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
				goto errorExit
			}

		}
		p.SetState(575)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_alt = p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 65, p.GetParserRuleContext())
		if p.HasError() {
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IMultiplicativeExpressionContext is an interface to support dynamic dispatch.
type IMultiplicativeExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	AllUnaryExpression() []IUnaryExpressionContext
	UnaryExpression(i int) IUnaryExpressionContext
	AllSTAR() []antlr.TerminalNode
	STAR(i int) antlr.TerminalNode
	AllSLASH() []antlr.TerminalNode
	SLASH(i int) antlr.TerminalNode

	// IsMultiplicativeExpressionContext differentiates from other interfaces.
	IsMultiplicativeExpressionContext()
}

type MultiplicativeExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyMultiplicativeExpressionContext() *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression
	return p
}

func InitEmptyMultiplicativeExpressionContext(p *MultiplicativeExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression
}

func (*MultiplicativeExpressionContext) IsMultiplicativeExpressionContext() {}

func NewMultiplicativeExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MultiplicativeExpressionContext {
	var p = new(MultiplicativeExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_multiplicativeExpression

	return p
}

func (s *MultiplicativeExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *MultiplicativeExpressionContext) AllUnaryExpression() []IUnaryExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			len++
		}
	}

	tst := make([]IUnaryExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IUnaryExpressionContext); ok {
			tst[i] = t.(IUnaryExpressionContext)
			i++
		}
	}

	return tst
}

func (s *MultiplicativeExpressionContext) UnaryExpression(i int) IUnaryExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IUnaryExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IUnaryExpressionContext)
}

func (s *MultiplicativeExpressionContext) AllSTAR() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserSTAR)
}

func (s *MultiplicativeExpressionContext) STAR(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserSTAR, i)
}

func (s *MultiplicativeExpressionContext) AllSLASH() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserSLASH)
}

func (s *MultiplicativeExpressionContext) SLASH(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserSLASH, i)
}

func (s *MultiplicativeExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MultiplicativeExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MultiplicativeExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitMultiplicativeExpression(s)
	}
}

func (s *MultiplicativeExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitMultiplicativeExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) MultiplicativeExpression() (localctx IMultiplicativeExpressionContext) {
	localctx = NewMultiplicativeExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 120, SparqlParserRULE_multiplicativeExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(576)
		p.UnaryExpression()
	}
	p.SetState(581)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	for _la == SparqlParserSLASH || _la == SparqlParserSTAR {
		{
			p.SetState(577)
			_la = p.GetTokenStream().LA(1)

			if !(_la == SparqlParserSLASH || _la == SparqlParserSTAR) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}
		{
			p.SetState(578)
			p.UnaryExpression()
		}

		p.SetState(583)
		p.GetErrorHandler().Sync(p)
		if p.HasError() {
			goto errorExit
		}
		_la = p.GetTokenStream().LA(1)
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IUnaryExpressionContext is an interface to support dynamic dispatch.
type IUnaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PrimaryExpression() IPrimaryExpressionContext
	EXCLAMATION() antlr.TerminalNode
	PLUS() antlr.TerminalNode
	MINUS() antlr.TerminalNode

	// IsUnaryExpressionContext differentiates from other interfaces.
	IsUnaryExpressionContext()
}

type UnaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyUnaryExpressionContext() *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_unaryExpression
	return p
}

func InitEmptyUnaryExpressionContext(p *UnaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_unaryExpression
}

func (*UnaryExpressionContext) IsUnaryExpressionContext() {}

func NewUnaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *UnaryExpressionContext {
	var p = new(UnaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_unaryExpression

	return p
}

func (s *UnaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *UnaryExpressionContext) PrimaryExpression() IPrimaryExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrimaryExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrimaryExpressionContext)
}

func (s *UnaryExpressionContext) EXCLAMATION() antlr.TerminalNode {
	return s.GetToken(SparqlParserEXCLAMATION, 0)
}

func (s *UnaryExpressionContext) PLUS() antlr.TerminalNode {
	return s.GetToken(SparqlParserPLUS, 0)
}

func (s *UnaryExpressionContext) MINUS() antlr.TerminalNode {
	return s.GetToken(SparqlParserMINUS, 0)
}

func (s *UnaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *UnaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *UnaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitUnaryExpression(s)
	}
}

func (s *UnaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitUnaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) UnaryExpression() (localctx IUnaryExpressionContext) {
	localctx = NewUnaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 122, SparqlParserRULE_unaryExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(585)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if (int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2884555561330802688) != 0 {
		{
			p.SetState(584)
			_la = p.GetTokenStream().LA(1)

			if !((int64(_la) & ^0x3f) == 0 && ((int64(1)<<_la)&2884555561330802688) != 0) {
				p.GetErrorHandler().RecoverInline(p)
			} else {
				p.GetErrorHandler().ReportMatch(p)
				p.Consume()
			}
		}

	}
	{
		p.SetState(587)
		p.PrimaryExpression()
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrimaryExpressionContext is an interface to support dynamic dispatch.
type IPrimaryExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BrackettedExpression() IBrackettedExpressionContext
	BuiltInCall() IBuiltInCallContext
	IriRefOrFunction() IIriRefOrFunctionContext
	RdfLiteral() IRdfLiteralContext
	NumericLiteral() INumericLiteralContext
	BooleanLiteral() IBooleanLiteralContext
	Var_() IVar_Context

	// IsPrimaryExpressionContext differentiates from other interfaces.
	IsPrimaryExpressionContext()
}

type PrimaryExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrimaryExpressionContext() *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_primaryExpression
	return p
}

func InitEmptyPrimaryExpressionContext(p *PrimaryExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_primaryExpression
}

func (*PrimaryExpressionContext) IsPrimaryExpressionContext() {}

func NewPrimaryExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimaryExpressionContext {
	var p = new(PrimaryExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_primaryExpression

	return p
}

func (s *PrimaryExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimaryExpressionContext) BrackettedExpression() IBrackettedExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBrackettedExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBrackettedExpressionContext)
}

func (s *PrimaryExpressionContext) BuiltInCall() IBuiltInCallContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBuiltInCallContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBuiltInCallContext)
}

func (s *PrimaryExpressionContext) IriRefOrFunction() IIriRefOrFunctionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefOrFunctionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefOrFunctionContext)
}

func (s *PrimaryExpressionContext) RdfLiteral() IRdfLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRdfLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRdfLiteralContext)
}

func (s *PrimaryExpressionContext) NumericLiteral() INumericLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralContext)
}

func (s *PrimaryExpressionContext) BooleanLiteral() IBooleanLiteralContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IBooleanLiteralContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IBooleanLiteralContext)
}

func (s *PrimaryExpressionContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *PrimaryExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimaryExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimaryExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPrimaryExpression(s)
	}
}

func (s *PrimaryExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPrimaryExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) PrimaryExpression() (localctx IPrimaryExpressionContext) {
	localctx = NewPrimaryExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 124, SparqlParserRULE_primaryExpression)
	p.SetState(596)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserL_PAREN:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(589)
			p.BrackettedExpression()
		}

	case SparqlParserBOUND, SparqlParserDATATYPE, SparqlParserLANG, SparqlParserLANGMATCHES, SparqlParserREGEX, SparqlParserSTR, SparqlParserIS_LITERAL, SparqlParserIS_BLANK, SparqlParserIS_URI, SparqlParserIS_IRI, SparqlParserSAME_TERM:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(590)
			p.BuiltInCall()
		}

	case SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(591)
			p.IriRefOrFunction()
		}

	case SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(592)
			p.RdfLiteral()
		}

	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(593)
			p.NumericLiteral()
		}

	case SparqlParserTRUE, SparqlParserFALSE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(594)
			p.BooleanLiteral()
		}

	case SparqlParserVAR1, SparqlParserVAR2:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(595)
			p.Var_()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBrackettedExpressionContext is an interface to support dynamic dispatch.
type IBrackettedExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	L_PAREN() antlr.TerminalNode
	Expression() IExpressionContext
	R_PAREN() antlr.TerminalNode

	// IsBrackettedExpressionContext differentiates from other interfaces.
	IsBrackettedExpressionContext()
}

type BrackettedExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBrackettedExpressionContext() *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_brackettedExpression
	return p
}

func InitEmptyBrackettedExpressionContext(p *BrackettedExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_brackettedExpression
}

func (*BrackettedExpressionContext) IsBrackettedExpressionContext() {}

func NewBrackettedExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BrackettedExpressionContext {
	var p = new(BrackettedExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_brackettedExpression

	return p
}

func (s *BrackettedExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *BrackettedExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *BrackettedExpressionContext) Expression() IExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BrackettedExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *BrackettedExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BrackettedExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BrackettedExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBrackettedExpression(s)
	}
}

func (s *BrackettedExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBrackettedExpression(s)
	}
}

func (s *BrackettedExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBrackettedExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BrackettedExpression() (localctx IBrackettedExpressionContext) {
	localctx = NewBrackettedExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 126, SparqlParserRULE_brackettedExpression)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(598)
		p.Match(SparqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(599)
		p.Expression()
	}
	{
		p.SetState(600)
		p.Match(SparqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBuiltInCallContext is an interface to support dynamic dispatch.
type IBuiltInCallContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STR() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	R_PAREN() antlr.TerminalNode
	LANG() antlr.TerminalNode
	LANGMATCHES() antlr.TerminalNode
	COMMA() antlr.TerminalNode
	DATATYPE() antlr.TerminalNode
	BOUND() antlr.TerminalNode
	Var_() IVar_Context
	SAME_TERM() antlr.TerminalNode
	IS_IRI() antlr.TerminalNode
	IS_URI() antlr.TerminalNode
	IS_BLANK() antlr.TerminalNode
	IS_LITERAL() antlr.TerminalNode
	RegexExpression() IRegexExpressionContext

	// IsBuiltInCallContext differentiates from other interfaces.
	IsBuiltInCallContext()
}

type BuiltInCallContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBuiltInCallContext() *BuiltInCallContext {
	var p = new(BuiltInCallContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_builtInCall
	return p
}

func InitEmptyBuiltInCallContext(p *BuiltInCallContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_builtInCall
}

func (*BuiltInCallContext) IsBuiltInCallContext() {}

func NewBuiltInCallContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BuiltInCallContext {
	var p = new(BuiltInCallContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_builtInCall

	return p
}

func (s *BuiltInCallContext) GetParser() antlr.Parser { return s.parser }

func (s *BuiltInCallContext) STR() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTR, 0)
}

func (s *BuiltInCallContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *BuiltInCallContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *BuiltInCallContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BuiltInCallContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *BuiltInCallContext) LANG() antlr.TerminalNode {
	return s.GetToken(SparqlParserLANG, 0)
}

func (s *BuiltInCallContext) LANGMATCHES() antlr.TerminalNode {
	return s.GetToken(SparqlParserLANGMATCHES, 0)
}

func (s *BuiltInCallContext) COMMA() antlr.TerminalNode {
	return s.GetToken(SparqlParserCOMMA, 0)
}

func (s *BuiltInCallContext) DATATYPE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDATATYPE, 0)
}

func (s *BuiltInCallContext) BOUND() antlr.TerminalNode {
	return s.GetToken(SparqlParserBOUND, 0)
}

func (s *BuiltInCallContext) Var_() IVar_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IVar_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IVar_Context)
}

func (s *BuiltInCallContext) SAME_TERM() antlr.TerminalNode {
	return s.GetToken(SparqlParserSAME_TERM, 0)
}

func (s *BuiltInCallContext) IS_IRI() antlr.TerminalNode {
	return s.GetToken(SparqlParserIS_IRI, 0)
}

func (s *BuiltInCallContext) IS_URI() antlr.TerminalNode {
	return s.GetToken(SparqlParserIS_URI, 0)
}

func (s *BuiltInCallContext) IS_BLANK() antlr.TerminalNode {
	return s.GetToken(SparqlParserIS_BLANK, 0)
}

func (s *BuiltInCallContext) IS_LITERAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserIS_LITERAL, 0)
}

func (s *BuiltInCallContext) RegexExpression() IRegexExpressionContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IRegexExpressionContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IRegexExpressionContext)
}

func (s *BuiltInCallContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BuiltInCallContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BuiltInCallContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBuiltInCall(s)
	}
}

func (s *BuiltInCallContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBuiltInCall(s)
	}
}

func (s *BuiltInCallContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBuiltInCall(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BuiltInCall() (localctx IBuiltInCallContext) {
	localctx = NewBuiltInCallContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 128, SparqlParserRULE_builtInCall)
	p.SetState(657)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserSTR:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(602)
			p.Match(SparqlParserSTR)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(603)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(604)
			p.Expression()
		}
		{
			p.SetState(605)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserLANG:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(607)
			p.Match(SparqlParserLANG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(608)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(609)
			p.Expression()
		}
		{
			p.SetState(610)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserLANGMATCHES:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(612)
			p.Match(SparqlParserLANGMATCHES)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(613)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(614)
			p.Expression()
		}
		{
			p.SetState(615)
			p.Match(SparqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(616)
			p.Expression()
		}
		{
			p.SetState(617)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserDATATYPE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(619)
			p.Match(SparqlParserDATATYPE)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(620)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(621)
			p.Expression()
		}
		{
			p.SetState(622)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserBOUND:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(624)
			p.Match(SparqlParserBOUND)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(625)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(626)
			p.Var_()
		}
		{
			p.SetState(627)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserSAME_TERM:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(629)
			p.Match(SparqlParserSAME_TERM)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(630)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(631)
			p.Expression()
		}
		{
			p.SetState(632)
			p.Match(SparqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(633)
			p.Expression()
		}
		{
			p.SetState(634)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserIS_IRI:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(636)
			p.Match(SparqlParserIS_IRI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(637)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(638)
			p.Expression()
		}
		{
			p.SetState(639)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserIS_URI:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(641)
			p.Match(SparqlParserIS_URI)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(642)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(643)
			p.Expression()
		}
		{
			p.SetState(644)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserIS_BLANK:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(646)
			p.Match(SparqlParserIS_BLANK)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(647)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(648)
			p.Expression()
		}
		{
			p.SetState(649)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserIS_LITERAL:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(651)
			p.Match(SparqlParserIS_LITERAL)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(652)
			p.Match(SparqlParserL_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(653)
			p.Expression()
		}
		{
			p.SetState(654)
			p.Match(SparqlParserR_PAREN)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserREGEX:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(656)
			p.RegexExpression()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRegexExpressionContext is an interface to support dynamic dispatch.
type IRegexExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	REGEX() antlr.TerminalNode
	L_PAREN() antlr.TerminalNode
	AllExpression() []IExpressionContext
	Expression(i int) IExpressionContext
	AllCOMMA() []antlr.TerminalNode
	COMMA(i int) antlr.TerminalNode
	R_PAREN() antlr.TerminalNode

	// IsRegexExpressionContext differentiates from other interfaces.
	IsRegexExpressionContext()
}

type RegexExpressionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRegexExpressionContext() *RegexExpressionContext {
	var p = new(RegexExpressionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_regexExpression
	return p
}

func InitEmptyRegexExpressionContext(p *RegexExpressionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_regexExpression
}

func (*RegexExpressionContext) IsRegexExpressionContext() {}

func NewRegexExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RegexExpressionContext {
	var p = new(RegexExpressionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_regexExpression

	return p
}

func (s *RegexExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *RegexExpressionContext) REGEX() antlr.TerminalNode {
	return s.GetToken(SparqlParserREGEX, 0)
}

func (s *RegexExpressionContext) L_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserL_PAREN, 0)
}

func (s *RegexExpressionContext) AllExpression() []IExpressionContext {
	children := s.GetChildren()
	len := 0
	for _, ctx := range children {
		if _, ok := ctx.(IExpressionContext); ok {
			len++
		}
	}

	tst := make([]IExpressionContext, len)
	i := 0
	for _, ctx := range children {
		if t, ok := ctx.(IExpressionContext); ok {
			tst[i] = t.(IExpressionContext)
			i++
		}
	}

	return tst
}

func (s *RegexExpressionContext) Expression(i int) IExpressionContext {
	var t antlr.RuleContext
	j := 0
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IExpressionContext); ok {
			if j == i {
				t = ctx.(antlr.RuleContext)
				break
			}
			j++
		}
	}

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *RegexExpressionContext) AllCOMMA() []antlr.TerminalNode {
	return s.GetTokens(SparqlParserCOMMA)
}

func (s *RegexExpressionContext) COMMA(i int) antlr.TerminalNode {
	return s.GetToken(SparqlParserCOMMA, i)
}

func (s *RegexExpressionContext) R_PAREN() antlr.TerminalNode {
	return s.GetToken(SparqlParserR_PAREN, 0)
}

func (s *RegexExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RegexExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RegexExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterRegexExpression(s)
	}
}

func (s *RegexExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitRegexExpression(s)
	}
}

func (s *RegexExpressionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitRegexExpression(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) RegexExpression() (localctx IRegexExpressionContext) {
	localctx = NewRegexExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 130, SparqlParserRULE_regexExpression)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(659)
		p.Match(SparqlParserREGEX)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(660)
		p.Match(SparqlParserL_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(661)
		p.Expression()
	}
	{
		p.SetState(662)
		p.Match(SparqlParserCOMMA)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}
	{
		p.SetState(663)
		p.Expression()
	}
	p.SetState(666)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	_la = p.GetTokenStream().LA(1)

	if _la == SparqlParserCOMMA {
		{
			p.SetState(664)
			p.Match(SparqlParserCOMMA)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(665)
			p.Expression()
		}

	}
	{
		p.SetState(668)
		p.Match(SparqlParserR_PAREN)
		if p.HasError() {
			// Recognition error - abort rule
			goto errorExit
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIriRefOrFunctionContext is an interface to support dynamic dispatch.
type IIriRefOrFunctionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IriRef() IIriRefContext
	ArgList() IArgListContext

	// IsIriRefOrFunctionContext differentiates from other interfaces.
	IsIriRefOrFunctionContext()
}

type IriRefOrFunctionContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriRefOrFunctionContext() *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction
	return p
}

func InitEmptyIriRefOrFunctionContext(p *IriRefOrFunctionContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction
}

func (*IriRefOrFunctionContext) IsIriRefOrFunctionContext() {}

func NewIriRefOrFunctionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefOrFunctionContext {
	var p = new(IriRefOrFunctionContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_iriRefOrFunction

	return p
}

func (s *IriRefOrFunctionContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRefOrFunctionContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *IriRefOrFunctionContext) ArgList() IArgListContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IArgListContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IArgListContext)
}

func (s *IriRefOrFunctionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriRefOrFunctionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriRefOrFunctionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterIriRefOrFunction(s)
	}
}

func (s *IriRefOrFunctionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitIriRefOrFunction(s)
	}
}

func (s *IriRefOrFunctionContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitIriRefOrFunction(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) IriRefOrFunction() (localctx IIriRefOrFunctionContext) {
	localctx = NewIriRefOrFunctionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 132, SparqlParserRULE_iriRefOrFunction)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(670)
		p.IriRef()
	}
	p.SetState(672)
	p.GetErrorHandler().Sync(p)

	if p.GetInterpreter().AdaptivePredict(p.BaseParser, p.GetTokenStream(), 71, p.GetParserRuleContext()) == 1 {
		{
			p.SetState(671)
			p.ArgList()
		}

	} else if p.HasError() { // JIM
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IRdfLiteralContext is an interface to support dynamic dispatch.
type IRdfLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	String_() IString_Context
	LANGTAG() antlr.TerminalNode
	DOUBLE_CARET() antlr.TerminalNode
	IriRef() IIriRefContext

	// IsRdfLiteralContext differentiates from other interfaces.
	IsRdfLiteralContext()
}

type RdfLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyRdfLiteralContext() *RdfLiteralContext {
	var p = new(RdfLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfLiteral
	return p
}

func InitEmptyRdfLiteralContext(p *RdfLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_rdfLiteral
}

func (*RdfLiteralContext) IsRdfLiteralContext() {}

func NewRdfLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RdfLiteralContext {
	var p = new(RdfLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_rdfLiteral

	return p
}

func (s *RdfLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *RdfLiteralContext) String_() IString_Context {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IString_Context); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IString_Context)
}

func (s *RdfLiteralContext) LANGTAG() antlr.TerminalNode {
	return s.GetToken(SparqlParserLANGTAG, 0)
}

func (s *RdfLiteralContext) DOUBLE_CARET() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_CARET, 0)
}

func (s *RdfLiteralContext) IriRef() IIriRefContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IIriRefContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IIriRefContext)
}

func (s *RdfLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RdfLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RdfLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitRdfLiteral(s)
	}
}

func (s *RdfLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitRdfLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) RdfLiteral() (localctx IRdfLiteralContext) {
	localctx = NewRdfLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 134, SparqlParserRULE_rdfLiteral)
	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(674)
		p.String_()
	}
	p.SetState(678)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}
	switch p.GetTokenStream().LA(1) {
	case SparqlParserLANGTAG:
		{
			p.SetState(675)
			p.Match(SparqlParserLANGTAG)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserDOUBLE_CARET:
		{
			p.SetState(676)
			p.Match(SparqlParserDOUBLE_CARET)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}
		{
			p.SetState(677)
			p.IriRef()
		}

	case SparqlParserEOF, SparqlParserA, SparqlParserBIND, SparqlParserAS, SparqlParserHAVING, SparqlParserBOUND, SparqlParserDATATYPE, SparqlParserFILTER, SparqlParserGRAPH, SparqlParserLANG, SparqlParserLANGMATCHES, SparqlParserLIMIT, SparqlParserOFFSET, SparqlParserOPTIONAL, SparqlParserORDER, SparqlParserREGEX, SparqlParserSTR, SparqlParserTRUE, SparqlParserFALSE, SparqlParserIS_LITERAL, SparqlParserIS_BLANK, SparqlParserIS_URI, SparqlParserIS_IRI, SparqlParserSAME_TERM, SparqlParserCOMMA, SparqlParserDOT, SparqlParserDOUBLE_AMP, SparqlParserDOUBLE_BAR, SparqlParserEQUAL, SparqlParserEXCLAMATION, SparqlParserGREATER, SparqlParserGREATER_OR_EQUAL, SparqlParserLESS, SparqlParserLESS_OR_EQUAL, SparqlParserL_CURLY, SparqlParserL_PAREN, SparqlParserL_SQUARE, SparqlParserMINUS, SparqlParserNOT_EQUAL, SparqlParserPLUS, SparqlParserR_CURLY, SparqlParserR_PAREN, SparqlParserR_SQUARE, SparqlParserSEMICOLON, SparqlParserSLASH, SparqlParserSTAR, SparqlParserIRI_REF, SparqlParserPNAME_NS, SparqlParserPNAME_LN, SparqlParserBLANK_NODE_LABEL, SparqlParserVAR1, SparqlParserVAR2, SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE, SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE, SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE, SparqlParserSTRING_LITERAL1, SparqlParserSTRING_LITERAL2, SparqlParserNIL, SparqlParserANON:

	default:
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralContext is an interface to support dynamic dispatch.
type INumericLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	NumericLiteralUnsigned() INumericLiteralUnsignedContext
	NumericLiteralPositive() INumericLiteralPositiveContext
	NumericLiteralNegative() INumericLiteralNegativeContext

	// IsNumericLiteralContext differentiates from other interfaces.
	IsNumericLiteralContext()
}

type NumericLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralContext() *NumericLiteralContext {
	var p = new(NumericLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteral
	return p
}

func InitEmptyNumericLiteralContext(p *NumericLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteral
}

func (*NumericLiteralContext) IsNumericLiteralContext() {}

func NewNumericLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralContext {
	var p = new(NumericLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteral

	return p
}

func (s *NumericLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralContext) NumericLiteralUnsigned() INumericLiteralUnsignedContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralUnsignedContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralUnsignedContext)
}

func (s *NumericLiteralContext) NumericLiteralPositive() INumericLiteralPositiveContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralPositiveContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralPositiveContext)
}

func (s *NumericLiteralContext) NumericLiteralNegative() INumericLiteralNegativeContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(INumericLiteralNegativeContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(INumericLiteralNegativeContext)
}

func (s *NumericLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNumericLiteral(s)
	}
}

func (s *NumericLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNumericLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NumericLiteral() (localctx INumericLiteralContext) {
	localctx = NewNumericLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 136, SparqlParserRULE_numericLiteral)
	p.SetState(683)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserINTEGER, SparqlParserDECIMAL, SparqlParserDOUBLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(680)
			p.NumericLiteralUnsigned()
		}

	case SparqlParserINTEGER_POSITIVE, SparqlParserDECIMAL_POSITIVE, SparqlParserDOUBLE_POSITIVE:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(681)
			p.NumericLiteralPositive()
		}

	case SparqlParserINTEGER_NEGATIVE, SparqlParserDECIMAL_NEGATIVE, SparqlParserDOUBLE_NEGATIVE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(682)
			p.NumericLiteralNegative()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralUnsignedContext is an interface to support dynamic dispatch.
type INumericLiteralUnsignedContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER() antlr.TerminalNode
	DECIMAL() antlr.TerminalNode
	DOUBLE() antlr.TerminalNode

	// IsNumericLiteralUnsignedContext differentiates from other interfaces.
	IsNumericLiteralUnsignedContext()
}

type NumericLiteralUnsignedContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralUnsignedContext() *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned
	return p
}

func InitEmptyNumericLiteralUnsignedContext(p *NumericLiteralUnsignedContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned
}

func (*NumericLiteralUnsignedContext) IsNumericLiteralUnsignedContext() {}

func NewNumericLiteralUnsignedContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralUnsignedContext {
	var p = new(NumericLiteralUnsignedContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralUnsigned

	return p
}

func (s *NumericLiteralUnsignedContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralUnsignedContext) INTEGER() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER, 0)
}

func (s *NumericLiteralUnsignedContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL, 0)
}

func (s *NumericLiteralUnsignedContext) DOUBLE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE, 0)
}

func (s *NumericLiteralUnsignedContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralUnsignedContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralUnsignedContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNumericLiteralUnsigned(s)
	}
}

func (s *NumericLiteralUnsignedContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNumericLiteralUnsigned(s)
	}
}

func (s *NumericLiteralUnsignedContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNumericLiteralUnsigned(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NumericLiteralUnsigned() (localctx INumericLiteralUnsignedContext) {
	localctx = NewNumericLiteralUnsignedContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 138, SparqlParserRULE_numericLiteralUnsigned)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(685)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-75)) & ^0x3f) == 0 && ((int64(1)<<(_la-75))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralPositiveContext is an interface to support dynamic dispatch.
type INumericLiteralPositiveContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_POSITIVE() antlr.TerminalNode
	DECIMAL_POSITIVE() antlr.TerminalNode
	DOUBLE_POSITIVE() antlr.TerminalNode

	// IsNumericLiteralPositiveContext differentiates from other interfaces.
	IsNumericLiteralPositiveContext()
}

type NumericLiteralPositiveContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralPositiveContext() *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive
	return p
}

func InitEmptyNumericLiteralPositiveContext(p *NumericLiteralPositiveContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive
}

func (*NumericLiteralPositiveContext) IsNumericLiteralPositiveContext() {}

func NewNumericLiteralPositiveContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralPositiveContext {
	var p = new(NumericLiteralPositiveContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralPositive

	return p
}

func (s *NumericLiteralPositiveContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralPositiveContext) INTEGER_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) DECIMAL_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) DOUBLE_POSITIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_POSITIVE, 0)
}

func (s *NumericLiteralPositiveContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralPositiveContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralPositiveContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNumericLiteralPositive(s)
	}
}

func (s *NumericLiteralPositiveContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNumericLiteralPositive(s)
	}
}

func (s *NumericLiteralPositiveContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNumericLiteralPositive(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NumericLiteralPositive() (localctx INumericLiteralPositiveContext) {
	localctx = NewNumericLiteralPositiveContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 140, SparqlParserRULE_numericLiteralPositive)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(687)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-78)) & ^0x3f) == 0 && ((int64(1)<<(_la-78))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// INumericLiteralNegativeContext is an interface to support dynamic dispatch.
type INumericLiteralNegativeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	INTEGER_NEGATIVE() antlr.TerminalNode
	DECIMAL_NEGATIVE() antlr.TerminalNode
	DOUBLE_NEGATIVE() antlr.TerminalNode

	// IsNumericLiteralNegativeContext differentiates from other interfaces.
	IsNumericLiteralNegativeContext()
}

type NumericLiteralNegativeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyNumericLiteralNegativeContext() *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative
	return p
}

func InitEmptyNumericLiteralNegativeContext(p *NumericLiteralNegativeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative
}

func (*NumericLiteralNegativeContext) IsNumericLiteralNegativeContext() {}

func NewNumericLiteralNegativeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *NumericLiteralNegativeContext {
	var p = new(NumericLiteralNegativeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_numericLiteralNegative

	return p
}

func (s *NumericLiteralNegativeContext) GetParser() antlr.Parser { return s.parser }

func (s *NumericLiteralNegativeContext) INTEGER_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserINTEGER_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) DECIMAL_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDECIMAL_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) DOUBLE_NEGATIVE() antlr.TerminalNode {
	return s.GetToken(SparqlParserDOUBLE_NEGATIVE, 0)
}

func (s *NumericLiteralNegativeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *NumericLiteralNegativeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *NumericLiteralNegativeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterNumericLiteralNegative(s)
	}
}

func (s *NumericLiteralNegativeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitNumericLiteralNegative(s)
	}
}

func (s *NumericLiteralNegativeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitNumericLiteralNegative(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) NumericLiteralNegative() (localctx INumericLiteralNegativeContext) {
	localctx = NewNumericLiteralNegativeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 142, SparqlParserRULE_numericLiteralNegative)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(689)
		_la = p.GetTokenStream().LA(1)

		if !((int64((_la-81)) & ^0x3f) == 0 && ((int64(1)<<(_la-81))&7) != 0) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBooleanLiteralContext is an interface to support dynamic dispatch.
type IBooleanLiteralContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	TRUE() antlr.TerminalNode
	FALSE() antlr.TerminalNode

	// IsBooleanLiteralContext differentiates from other interfaces.
	IsBooleanLiteralContext()
}

type BooleanLiteralContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBooleanLiteralContext() *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_booleanLiteral
	return p
}

func InitEmptyBooleanLiteralContext(p *BooleanLiteralContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_booleanLiteral
}

func (*BooleanLiteralContext) IsBooleanLiteralContext() {}

func NewBooleanLiteralContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BooleanLiteralContext {
	var p = new(BooleanLiteralContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_booleanLiteral

	return p
}

func (s *BooleanLiteralContext) GetParser() antlr.Parser { return s.parser }

func (s *BooleanLiteralContext) TRUE() antlr.TerminalNode {
	return s.GetToken(SparqlParserTRUE, 0)
}

func (s *BooleanLiteralContext) FALSE() antlr.TerminalNode {
	return s.GetToken(SparqlParserFALSE, 0)
}

func (s *BooleanLiteralContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BooleanLiteralContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BooleanLiteralContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBooleanLiteral(s)
	}
}

func (s *BooleanLiteralContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBooleanLiteral(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BooleanLiteral() (localctx IBooleanLiteralContext) {
	localctx = NewBooleanLiteralContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 144, SparqlParserRULE_booleanLiteral)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(691)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserTRUE || _la == SparqlParserFALSE) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IString_Context is an interface to support dynamic dispatch.
type IString_Context interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	STRING_LITERAL1() antlr.TerminalNode
	STRING_LITERAL2() antlr.TerminalNode

	// IsString_Context differentiates from other interfaces.
	IsString_Context()
}

type String_Context struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyString_Context() *String_Context {
	var p = new(String_Context)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_string_
	return p
}

func InitEmptyString_Context(p *String_Context) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_string_
}

func (*String_Context) IsString_Context() {}

func NewString_Context(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *String_Context {
	var p = new(String_Context)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_string_

	return p
}

func (s *String_Context) GetParser() antlr.Parser { return s.parser }

func (s *String_Context) STRING_LITERAL1() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTRING_LITERAL1, 0)
}

func (s *String_Context) STRING_LITERAL2() antlr.TerminalNode {
	return s.GetToken(SparqlParserSTRING_LITERAL2, 0)
}

func (s *String_Context) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *String_Context) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *String_Context) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterString_(s)
	}
}

func (s *String_Context) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitString_(s)
	}
}

func (s *String_Context) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitString_(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) String_() (localctx IString_Context) {
	localctx = NewString_Context(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 146, SparqlParserRULE_string_)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(693)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserSTRING_LITERAL1 || _la == SparqlParserSTRING_LITERAL2) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IIriRefContext is an interface to support dynamic dispatch.
type IIriRefContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	IRI_REF() antlr.TerminalNode
	PrefixedName() IPrefixedNameContext

	// IsIriRefContext differentiates from other interfaces.
	IsIriRefContext()
}

type IriRefContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyIriRefContext() *IriRefContext {
	var p = new(IriRefContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRef
	return p
}

func InitEmptyIriRefContext(p *IriRefContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_iriRef
}

func (*IriRefContext) IsIriRefContext() {}

func NewIriRefContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IriRefContext {
	var p = new(IriRefContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_iriRef

	return p
}

func (s *IriRefContext) GetParser() antlr.Parser { return s.parser }

func (s *IriRefContext) IRI_REF() antlr.TerminalNode {
	return s.GetToken(SparqlParserIRI_REF, 0)
}

func (s *IriRefContext) PrefixedName() IPrefixedNameContext {
	var t antlr.RuleContext
	for _, ctx := range s.GetChildren() {
		if _, ok := ctx.(IPrefixedNameContext); ok {
			t = ctx.(antlr.RuleContext)
			break
		}
	}

	if t == nil {
		return nil
	}

	return t.(IPrefixedNameContext)
}

func (s *IriRefContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IriRefContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IriRefContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterIriRef(s)
	}
}

func (s *IriRefContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitIriRef(s)
	}
}

func (s *IriRefContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitIriRef(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) IriRef() (localctx IIriRefContext) {
	localctx = NewIriRefContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 148, SparqlParserRULE_iriRef)
	p.SetState(697)
	p.GetErrorHandler().Sync(p)
	if p.HasError() {
		goto errorExit
	}

	switch p.GetTokenStream().LA(1) {
	case SparqlParserIRI_REF:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(695)
			p.Match(SparqlParserIRI_REF)
			if p.HasError() {
				// Recognition error - abort rule
				goto errorExit
			}
		}

	case SparqlParserPNAME_NS, SparqlParserPNAME_LN:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(696)
			p.PrefixedName()
		}

	default:
		p.SetError(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
		goto errorExit
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IPrefixedNameContext is an interface to support dynamic dispatch.
type IPrefixedNameContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	PNAME_LN() antlr.TerminalNode
	PNAME_NS() antlr.TerminalNode

	// IsPrefixedNameContext differentiates from other interfaces.
	IsPrefixedNameContext()
}

type PrefixedNameContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyPrefixedNameContext() *PrefixedNameContext {
	var p = new(PrefixedNameContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixedName
	return p
}

func InitEmptyPrefixedNameContext(p *PrefixedNameContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_prefixedName
}

func (*PrefixedNameContext) IsPrefixedNameContext() {}

func NewPrefixedNameContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrefixedNameContext {
	var p = new(PrefixedNameContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_prefixedName

	return p
}

func (s *PrefixedNameContext) GetParser() antlr.Parser { return s.parser }

func (s *PrefixedNameContext) PNAME_LN() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_LN, 0)
}

func (s *PrefixedNameContext) PNAME_NS() antlr.TerminalNode {
	return s.GetToken(SparqlParserPNAME_NS, 0)
}

func (s *PrefixedNameContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrefixedNameContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrefixedNameContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterPrefixedName(s)
	}
}

func (s *PrefixedNameContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitPrefixedName(s)
	}
}

func (s *PrefixedNameContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitPrefixedName(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) PrefixedName() (localctx IPrefixedNameContext) {
	localctx = NewPrefixedNameContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 150, SparqlParserRULE_prefixedName)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(699)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserPNAME_NS || _la == SparqlParserPNAME_LN) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}

// IBlankNodeContext is an interface to support dynamic dispatch.
type IBlankNodeContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Getter signatures
	BLANK_NODE_LABEL() antlr.TerminalNode
	ANON() antlr.TerminalNode

	// IsBlankNodeContext differentiates from other interfaces.
	IsBlankNodeContext()
}

type BlankNodeContext struct {
	antlr.BaseParserRuleContext
	parser antlr.Parser
}

func NewEmptyBlankNodeContext() *BlankNodeContext {
	var p = new(BlankNodeContext)
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNode
	return p
}

func InitEmptyBlankNodeContext(p *BlankNodeContext) {
	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, nil, -1)
	p.RuleIndex = SparqlParserRULE_blankNode
}

func (*BlankNodeContext) IsBlankNodeContext() {}

func NewBlankNodeContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BlankNodeContext {
	var p = new(BlankNodeContext)

	antlr.InitBaseParserRuleContext(&p.BaseParserRuleContext, parent, invokingState)

	p.parser = parser
	p.RuleIndex = SparqlParserRULE_blankNode

	return p
}

func (s *BlankNodeContext) GetParser() antlr.Parser { return s.parser }

func (s *BlankNodeContext) BLANK_NODE_LABEL() antlr.TerminalNode {
	return s.GetToken(SparqlParserBLANK_NODE_LABEL, 0)
}

func (s *BlankNodeContext) ANON() antlr.TerminalNode {
	return s.GetToken(SparqlParserANON, 0)
}

func (s *BlankNodeContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BlankNodeContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BlankNodeContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.EnterBlankNode(s)
	}
}

func (s *BlankNodeContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(SparqlParserListener); ok {
		listenerT.ExitBlankNode(s)
	}
}

func (s *BlankNodeContext) Accept(visitor antlr.ParseTreeVisitor) interface{} {
	switch t := visitor.(type) {
	case SparqlParserVisitor:
		return t.VisitBlankNode(s)

	default:
		return t.VisitChildren(s)
	}
}

func (p *SparqlParser) BlankNode() (localctx IBlankNodeContext) {
	localctx = NewBlankNodeContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 152, SparqlParserRULE_blankNode)
	var _la int

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(701)
		_la = p.GetTokenStream().LA(1)

		if !(_la == SparqlParserBLANK_NODE_LABEL || _la == SparqlParserANON) {
			p.GetErrorHandler().RecoverInline(p)
		} else {
			p.GetErrorHandler().ReportMatch(p)
			p.Consume()
		}
	}

errorExit:
	if p.HasError() {
		v := p.GetError()
		localctx.SetException(v)
		p.GetErrorHandler().ReportError(p, v)
		p.GetErrorHandler().Recover(p, v)
		p.SetError(nil)
	}
	p.ExitRule()
	return localctx
	goto errorExit // Trick to prevent compiler error if the label is not used
}
