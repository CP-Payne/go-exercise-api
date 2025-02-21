package domain

import (
	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/CP-Payne/exercise/internal/interfaces/repositories"
)

// DomainServices provides access to all domain services
// from a centralized location
type DomainServices struct {
	Muscle muscle.MuscleService
}

// NewDomainServices creates and initializes all domain service implementations
func NewDomainServices(r *repositories.Repositories) *DomainServices {
	return &DomainServices{
		Muscle: muscle.NewMuscleService(r.Muscles),
	}
}
