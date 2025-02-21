package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// Validate is a global validator instance used for request validation
var Validate *validator.Validate

// init initializes the validator with required structural validation enabled
func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

// readJSON decodes a JSON request body into the provided struct.
// It enforces a max request size and disallows unknown fields.
func (rh *ResponseHelper) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

// writeJSONError writes a standardized JSON error response.
// It wraps the error message in a consistent envelope structure.
func (rh *ResponseHelper) writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return rh.writeJSON(w, status, &envelope{Error: message})
}

// writeJSON marshals data to JSON and writes it to the response.
// It sets appropriate Content-Type headers and status code.
func (rh *ResponseHelper) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

// jsonResponse sends a successful JSON response with the provided data.
// It standardises responses by wrapping data in a common envelope structure.
func (rh *ResponseHelper) jsonResponse(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	if status == http.StatusNoContent || data == nil || data == "" {
		w.WriteHeader(status)
		return nil
	}

	type envelope struct {
		Data any `json:"data"`
	}
	return rh.writeJSON(w, status, &envelope{Data: data})
}
