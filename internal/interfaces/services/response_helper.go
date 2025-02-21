package services

import (
	"net/http"

	"go.uber.org/zap"
)

// ResponseHelper provides methods to standardize HTTP responses.
// It encapsulates response formatting and error handling logic.
type ResponseHelper struct {
	logger *zap.SugaredLogger
}

// NewResponseHelper creates a new response helper with the specified logger.
func NewResponseHelper(logger *zap.SugaredLogger) *ResponseHelper {
	return &ResponseHelper{logger: logger}
}

// internalServerError logs and sends a 500 internal Server Error response.
func (rh *ResponseHelper) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}

// forbiddenResponse logs and sends a 403 Forbidden response.
func (rh *ResponseHelper) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	rh.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")
	rh.writeJSONError(w, http.StatusForbidden, "forbidden")
}

// badRequestResponse logs and sends a 400 Bad Request response with the specified error message.
func (rh *ResponseHelper) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusBadRequest, err.Error())
}

// notFoundResponse logs and sends a 404 Not Found response.
func (rh *ResponseHelper) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusNotFound, "not found")
}

// conflictResponse logs and sends a 409 Conflict response.
func (rh *ResponseHelper) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Errorw("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusConflict, "conflict")
}

// unauthorizedErrorResponse logs and sends a 401 Unauthorized response.
func (rh *ResponseHelper) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

// unauthorizedBasicErrorResponse logs and sends a 401 Unauthorized response with Basic authentication header.
func (rh *ResponseHelper) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	rh.writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

// rateLimitExceededResponse logs and sends a 429 Too Many Requests response with retry information.
func (rh *ResponseHelper) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	rh.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	rh.writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
