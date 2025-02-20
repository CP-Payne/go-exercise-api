package services

import (
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

func init() {
	Validate = validator.New(validator.WithRequiredStructEnabled())
}

func (rh *ResponseHelper) readJSON(w http.ResponseWriter, r *http.Request, data any) error {
	maxBytes := 1_048_578

	r.Body = http.MaxBytesReader(w, r.Body, int64(maxBytes))

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	return decoder.Decode(data)
}

func (rh *ResponseHelper) writeJSONError(w http.ResponseWriter, status int, message string) error {
	type envelope struct {
		Error string `json:"error"`
	}
	return rh.writeJSON(w, status, &envelope{Error: message})
}

func (rh *ResponseHelper) writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(data)
}

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
