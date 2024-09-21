package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validator is a new instance of the validator
var Validator = validator.New()

/*
The ParseJSON function is a utility function that decodes the request body into the payload.
*/
func ParseJSON(r *http.Request, payload any) error {
	// Check if the request body is nil
	if r.Body == nil {
		return nil
	}
	// Decode the request body into the payload
	return json.NewDecoder(r.Body).Decode(&payload)
}

/*
The WriteJSON function is a utility function that writes the data to the response writer in JSON format.
*/
func WriteJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

/*
The WriteError function is a utility function that writes an error message to the response writer in JSON format.
*/
func WriteError(w http.ResponseWriter, status int, err error) error {
	return WriteJSON(w, status, map[string]string{"error": err.Error()})
}

/*
The ErrorHandler function is a utility function that logs the error and returns an internal server error to the client.
*/
func ErrorHandler(w http.ResponseWriter, status int, err error) {
	log.Println(err)
	w.WriteHeader(status)
	http.Error(w, "Internal server error", http.StatusInternalServerError)
}
