package sparql

import (
	"net/http"
)

// TODO: Basic HTTP integration stub. Finish after translator.
func handleSPARQL(w http.ResponseWriter, r *http.Request) {
	// Read query from POST or GET
	// Parse/translate/execute (call ParseAndTranslate, send response)
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("SPARQL integration endpoint not implemented yet"))
}
