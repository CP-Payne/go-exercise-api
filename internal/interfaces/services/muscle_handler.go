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

type MuscleHandler struct {
	muscleUseCase  application.MuscleUseCase
	logger         *zap.SugaredLogger
	responseHelper *ResponseHelper
}

func NewMuscleHandler(muscleUseCase application.MuscleUseCase, logger *zap.SugaredLogger, responseHelper *ResponseHelper) *MuscleHandler {
	return &MuscleHandler{
		muscleUseCase:  muscleUseCase,
		logger:         logger,
		responseHelper: responseHelper,
	}
}

func (h *MuscleHandler) RegisterRoutes(router chi.Router) {
	router.Route("/muscles", func(r chi.Router) {
		r.Get("/", h.GetMuscles)
		r.Get("/{muscleID}", h.GetMuscleByID)
		r.Post("/", h.CreateMuscle)
		r.Delete("/{muscleID}", h.DeleteMuscle)
	})
}

type CreateMusclePayload struct {
	ID     string `json:"id"`
	Name   string `json:"name" validate:"required,max=30"`
	UserID string `json:"userID"`
}

var (
	tempUserID = "762c3349-0230-4094-932b-5d0685fafd4e" // only used for testing, will retrieve it from jwt in context later on
)

func (h *MuscleHandler) CreateMuscle(w http.ResponseWriter, r *http.Request) {
	var payload CreateMusclePayload
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
	response := struct {
		ID string `json:"id"`
	}{
		ID: domainMuscle.ID().String(),
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusCreated, response); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}
}

func (h *MuscleHandler) GetMuscles(w http.ResponseWriter, r *http.Request) {
	domainMuscles, err := h.muscleUseCase.ListMusclesForUser(r.Context(), uuid.MustParse(tempUserID))
	if err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}

	type ResponseMuscle struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	responseBody := []ResponseMuscle{}

	for _, m := range domainMuscles {
		rm := ResponseMuscle{
			ID:   m.ID().String(),
			Name: m.Name(),
		}

		responseBody = append(responseBody, rm)
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusOK, responseBody); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}
}

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
	type ResponseMuscle struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}

	responseMuscle := ResponseMuscle{
		ID:   domainMuscle.ID().String(),
		Name: domainMuscle.Name(),
	}

	if err := h.responseHelper.jsonResponse(w, http.StatusOK, responseMuscle); err != nil {
		h.responseHelper.internalServerError(w, r, err)
		return
	}

}

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
