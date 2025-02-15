package services

import (
	"github.com/CP-Payne/exercise/internal/application"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type MuscleHandler struct {
	muscleUseCase *application.MuscleUseCase
	logger        *zap.Logger
}

func NewMuscleHandler(muscleUseCase *application.MuscleUseCase, logger *zap.Logger) *MuscleHandler {
	return &MuscleHandler{
		muscleUseCase: muscleUseCase,
		logger:        logger,
	}
}

func (h *MuscleHandler) RegisterRoutes(router chi.Router) {

}
