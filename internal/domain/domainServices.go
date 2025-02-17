package domain

import (
	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/CP-Payne/exercise/internal/interfaces/repositories"
)

type DomainServices struct {
	Muscle muscle.IMuscleService
}

func NewDomainServices(r *repositories.Repositories) *DomainServices {
	return &DomainServices{
		Muscle: muscle.NewMuscleService(r.Muscles),
	}
}
