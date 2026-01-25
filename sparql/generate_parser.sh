#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
ANTLR_VERSION=4.13.1
ANTLR_JAR="antlr-${ANTLR_VERSION}-complete.jar"
ANTLR_URL="https://www.antlr.org/download/${ANTLR_JAR}"
LEXER_GRAMMAR_FILE="SparqlLexer.g4"
PARSER_GRAMMAR_FILE="SparqlParser.g4"

if [[ ! -f ${ANTLR_JAR} ]]; then
	echo "Downloading ANTLR jar..."
	curl -OL "${ANTLR_URL}"
fi

if [[ ! -f ${LEXER_GRAMMAR_FILE} ]] || [[ ! -f ${PARSER_GRAMMAR_FILE} ]]; then
	echo "Missing grammar files. Download from https://github.com/antlr/grammars-v4/tree/master/sparql"
	exit 1
fi

echo "Generating Go parser from grammar files..."
rm -rf gen
java -cp "${ANTLR_JAR}" org.antlr.v4.Tool -Dlanguage=Go -visitor -o gen "${LEXER_GRAMMAR_FILE}" "${PARSER_GRAMMAR_FILE}"
echo "Done."
