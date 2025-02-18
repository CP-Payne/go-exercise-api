package services

import (
	"net/http"

	"github.com/CP-Payne/exercise/internal/application"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type MuscleHandler struct {
	muscleUseCase application.MuscleUseCase
	logger        *zap.SugaredLogger
}

func NewMuscleHandler(muscleUseCase application.MuscleUseCase, logger *zap.SugaredLogger) *MuscleHandler {
	return &MuscleHandler{
		muscleUseCase: muscleUseCase,
		logger:        logger,
	}
}

func (h *MuscleHandler) RegisterRoutes(router chi.Router) {
	router.Route("/muscles", func(r chi.Router) {
		r.Get("/", h.AddMuscles)
	})
}

func (h *MuscleHandler) AddMuscles(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("This is working")
}
