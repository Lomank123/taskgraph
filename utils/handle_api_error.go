package utils

import (
	"encoding/json"
	"net/http"
)

// HandleAPIError handles API errors and writes them to the response writer
func HandleAPIError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(map[string]string{"error": message})
}
