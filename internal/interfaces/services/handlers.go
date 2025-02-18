package services

import (
	"github.com/CP-Payne/exercise/internal/application"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handlers struct {
	muscle *MuscleHandler
	// More handlers to be added
}

func NewHandlers(useCases application.UseCases, logger *zap.SugaredLogger) *Handlers {
	responseHelper := NewResponseHelper(logger)
	return &Handlers{
		muscle: NewMuscleHandler(useCases.MuscleUseCase(), logger, responseHelper),
	}
}

func (h *Handlers) RegisterRoutes(router chi.Router) {
	h.muscle.RegisterRoutes(router)
}
