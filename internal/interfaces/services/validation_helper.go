package services

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
)

// ValidationError represents a structured validation error
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// validationMessages maps validation tags to user-friendly error messages.
var validationMessages = map[string]string{
	"required": "This field is required.",
	"min":      "Must be at least %s characters long.",
	"max":      "Must be at most %s characters long.",
	"email":    "Must be a valid email address.",
	"uuid":     "Must be a valid UUID.",
}

// ValidateStruct validates a struct and returns a list of detailed validation error messages
func (rh *ResponseHelper) ValidateStruct(data interface{}) []ValidationError {
	validate := validator.New()
	err := validate.Struct(data)
	if err == nil {
		return nil
	}

	var errors []ValidationError
	for _, err := range err.(validator.ValidationErrors) {
		// Lookup validation message, default to generic message
		msg, exists := validationMessages[err.Tag()]
		if exists {
			if err.Param() != "" {
				msg = fmt.Sprintf(msg, err.Param()) // Handle parameters like min=3, max=10
			}
		} else {
			msg = fmt.Sprintf("Invalid value for %s.", err.Field()) // Default message
		}

		errors = append(errors, ValidationError{
			Field:   err.Field(),
			Message: msg,
		})
	}

	return errors
}

// WriteValidationErrorResponse sends a structured validation error response
func (rh *ResponseHelper) WriteValidationErrorResponse(w http.ResponseWriter, errors []ValidationError) {
	response := map[string]interface{}{
		"errors": errors,
	}

	rh.jsonResponse(w, http.StatusBadRequest, response)
}
