package services

import (
	"github.com/CP-Payne/exercise/internal/application"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// Handlers holds all HTTP handlers for the application
type Handlers struct {
	muscle *MuscleHandler
	// More handlers to be added
}

// NewHandlers creates and initializes all handlers with their required dependencies.
func NewHandlers(useCases application.UseCases, logger *zap.SugaredLogger) *Handlers {
	responseHelper := NewResponseHelper(logger)
	return &Handlers{
		muscle: NewMuscleHandler(useCases.MuscleUseCase(), logger, responseHelper),
	}
}

// RegisterRoutes registers all handler routes with the provided router
func (h *Handlers) RegisterRoutes(router chi.Router) {
	h.muscle.RegisterRoutes(router)
}
