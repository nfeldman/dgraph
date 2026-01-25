#!/usr/bin/env bash
set -euo pipefail
cd "$(dirname "$0")"
ANTLR_VERSION=4.11.1
ANTLR_JAR="antlr-$ANTLR_VERSION-complete.jar"
ANTLR_URL="https://www.antlr.org/download/$ANTLR_JAR"
GRAMMAR_FILE="SPARQL.g4"

if [ ! -f "$ANTLR_JAR" ]; then
	echo "Downloading ANTLR jar..."
	curl -OL "$ANTLR_URL"
fi

if [ ! -f "$GRAMMAR_FILE" ]; then
	echo "Missing $GRAMMAR_FILE. Download it from https://github.com/antlr/grammars-v4/tree/master/sparql"
	exit 1
fi

echo "Generating Go parser from $GRAMMAR_FILE..."
rm -rf gen
java -cp "$ANTLR_JAR" org.antlr.v4.Tool -Dlanguage=Go -visitor -o gen "$GRAMMAR_FILE"
echo "Done."
