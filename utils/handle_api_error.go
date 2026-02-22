package utils

import (
	"encoding/json"
	"net/http"
)

// HandleAPIError handles API errors and writes them to the response writer
func HandleAPIError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	// It is better than w.Write([]byte(`{"error": "` + message + `"}`))
	// because you don't have to validate yourself with properly setting all the quotes.
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
