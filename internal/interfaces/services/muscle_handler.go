package services

import (
	"errors"
	"net/http"

	"github.com/CP-Payne/exercise/internal/application"
	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/CP-Payne/exercise/internal/interfaces/repositories"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

// MuscleHandler handles HTTP requests related to muscle resources.
type MuscleHandler struct {
	muscleUseCase  application.MuscleUseCase
	logger         *zap.SugaredLogger
	responseHelper *ResponseHelper
}

// NewMuscleHandler creates a new muscle handler with the specified dependencies.
func NewMuscleHandler(muscleUseCase application.MuscleUseCase, logger *zap.SugaredLogger, responseHelper *ResponseHelper) *MuscleHandler {
	return &MuscleHandler{
		muscleUseCase:  muscleUseCase,
		logger:         logger,
		responseHelper: responseHelper,
	}
}

// RegisterRoutes sets up all muscle-related routes on the provided router.
func (h *MuscleHandler) RegisterRoutes(router chi.Router) {
	router.Route("/muscles", func(r chi.Router) {
		r.Get("/", h.GetMuscles)
		r.Get("/{muscleID}", h.GetMuscleByID)
		r.Post("/", h.CreateMuscle)
		r.Delete("/{muscleID}", h.DeleteMuscle)
	})
}

// MuscleListResponse represents a collection of muscle responses
type MuscleListResponse []MuscleResponse

// CreateMuscleResponse defines the expected structure for muscle creation requests.
type CreateMuscleRequest struct {
	Name   string `json:"name" validate:"required,max=30"`
	UserID string `json:"userID"`
}

// CreateMuscleResponse defines teh response structure after successfull muscle creation.
type CreateMuscleResponse struct {
	ID string `json:"id"`
}

// MuscleResponse defines the standard response structure for muscle data.
type MuscleResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

var (
	// tempUserID is a placeholder user ID for testing
	// TODO: Replace with JWT authentication to extract user ID from context.
	tempUserID = "762c3349-0230-4094-932b-5d0685fafd4e" // only used for testing, will retrieve it from jwt in context later on
)

// CreateMuscle handles POST requests to create a new muscle.
func (h *MuscleHandler) CreateMuscle(w http.ResponseWriter, r *http.Request) {
	var payload CreateMuscleRequest
	if err := h.responseHelper.readJSON(w, r, &payload); err != nil {
		h.responseHelper.badRequestResponse(w, r, err)
		return
	}

	if validationErrors := h.responseHelper.ValidateStruct(payload); validationErrors != nil {
		h.responseHelper.WriteValidationErrorResponse(w, validationErrors)
		return
	}

	domainMuscle, err := muscle.NewMuscle(muscle.MuscleParams{Name: payload.Name})
	if err != nil {
		if errors.Is(err, muscle.ErrInvalidMuscle) {
			h.responseHelper.badRequestResponse(w, r, err)
			return
		}
		h.responseHelper.internalServerError(w, r, err)
		return
	}

	if err := h.muscleUseCase.CreateMuscle(r.Context(), uuid.MustParse(tempUserID), domainMuscle); err != nil {
		if errors.Is(repositories.ErrDuplicateMuscleName, err) {
			h.responseHelper.badRequestResponse(w, r, err)
			return
		}
		h.responseHelper.internalServerError(w, r, err)
		return
	}
	// Only return ID
	response := CreateMuscleResponse{
		ID: domainMuscle.ID().String(),
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusCreated, response); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}
}

// GetMuscles handles GET requests to retrieve all muscles for the current user.
func (h *MuscleHandler) GetMuscles(w http.ResponseWriter, r *http.Request) {
	domainMuscles, err := h.muscleUseCase.ListMusclesForUser(r.Context(), uuid.MustParse(tempUserID))
	if err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}

	responseBody := make(MuscleListResponse, 0, len(domainMuscles))

	for _, m := range domainMuscles {
		responseBody = append(responseBody, MuscleResponse{
			ID:   m.ID().String(),
			Name: m.Name(),
		})
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusOK, responseBody); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}
}

// GetMuscleByID handles GET requests to retrieve a muscle by ID for the current user.
func (h *MuscleHandler) GetMuscleByID(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "muscleID")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.responseHelper.badRequestResponse(w, r, err)
		return
	}

	domainMuscle, err := h.muscleUseCase.GetMuscleByID(r.Context(), uuid.MustParse(tempUserID), id)
	if err != nil {
		switch err {
		case repositories.ErrNotFound:
			h.responseHelper.notFoundResponse(w, r, err)
			return
		default:
			h.responseHelper.internalServerError(w, r, err)
			return
		}
	}

	response := MuscleResponse{
		ID:   domainMuscle.ID().String(),
		Name: domainMuscle.Name(),
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusOK, response); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}

}

// DeleteMuscle handles DELETE requests to delete a muscle for the current user by ID.
func (h *MuscleHandler) DeleteMuscle(w http.ResponseWriter, r *http.Request) {
	idParam := chi.URLParam(r, "muscleID")
	id, err := uuid.Parse(idParam)
	if err != nil {
		h.responseHelper.badRequestResponse(w, r, err)
		return
	}

	if err := h.muscleUseCase.DeleteMuscle(r.Context(), uuid.MustParse(tempUserID), id); err != nil {
		switch {
		case errors.Is(err, repositories.ErrNotFound):
			h.responseHelper.notFoundResponse(w, r, err)
		default:
			h.responseHelper.internalServerError(w, r, err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)

}
