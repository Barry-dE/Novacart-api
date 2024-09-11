// The utils package parses and writes JSON responses for HTTP requests.
package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// JSONParser parses the JSON body of an HTTP request into the provided payload structure.
func JSONParser(r *http.Request, payload any) error {
	if r.Body == nil {
		return fmt.Errorf("no request body was returned")
	}
	return json.NewDecoder(r.Body).Decode(payload)
}

// JSONWriter writes the provided data as a JSON response with the given status code.
func JSONWriter(w http.ResponseWriter, statusCode int, v any) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	return json.NewEncoder(w).Encode(v)
}

// ErrorWriter generates a standard JSON error response with the given status code and error message.
func ErrorWriter(w http.ResponseWriter, statusCode int, err error) {
	JSONWriter(w, statusCode, map[string]string{"error": err.Error()})
}
