package services

import (
	"net/http"

	"go.uber.org/zap"
)

type ResponseHelper struct {
	logger *zap.SugaredLogger
}

func NewResponseHelper(logger *zap.SugaredLogger) *ResponseHelper {
	return &ResponseHelper{logger: logger}
}

func (rh *ResponseHelper) internalServerError(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Errorw("internal error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusInternalServerError, "the server encountered a problem")
}
func (rh *ResponseHelper) forbiddenResponse(w http.ResponseWriter, r *http.Request) {
	rh.logger.Warnw("forbidden", "method", r.Method, "path", r.URL.Path, "error")
	rh.writeJSONError(w, http.StatusForbidden, "forbidden")
}

func (rh *ResponseHelper) badRequestResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("bad request", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusBadRequest, err.Error())
}

func (rh *ResponseHelper) notFoundResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("not found error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusNotFound, "not found")
}

func (rh *ResponseHelper) conflictResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Errorw("conflict response", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusConflict, "conflict")
}

func (rh *ResponseHelper) unauthorizedErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("unauthorized error", "method", r.Method, "path", r.URL.Path, "error", err.Error())
	rh.writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (rh *ResponseHelper) unauthorizedBasicErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	rh.logger.Warnw("unauthorized basic error", "method", r.Method, "path", r.URL.Path, "error", err.Error())

	w.Header().Set("WWW-Authenticate", `Basic realm="restricted", charset="UTF-8"`)
	rh.writeJSONError(w, http.StatusUnauthorized, "unauthorized")
}

func (rh *ResponseHelper) rateLimitExceededResponse(w http.ResponseWriter, r *http.Request, retryAfter string) {
	rh.logger.Warnw("rate limit exceeded", "method", r.Method, "path", r.URL.Path)

	w.Header().Set("Retry-After", retryAfter)

	rh.writeJSONError(w, http.StatusTooManyRequests, "rate limit exceeded, retry after: "+retryAfter)
}
