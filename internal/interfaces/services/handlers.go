package services

import (
	"github.com/CP-Payne/exercise/internal/application"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

type Handlers struct {
	muscle *MuscleHandler
}

func NewHandlers(applicationUseCases application.ApplicationUseCases, logger *zap.SugaredLogger) *Handlers {
	return &Handlers{
		muscle: NewMuscleHandler(applicationUseCases.Muscle, logger),
	}
}

func (h *Handlers) RegisterRoutes(router chi.Router) {
	h.muscle.RegisterRoutes(router)
}
