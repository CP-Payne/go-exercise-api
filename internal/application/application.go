package application

import "github.com/CP-Payne/exercise/internal/domain"

type UseCases interface {
	MuscleUseCase() MuscleUseCase
}

type useCases struct {
	Muscle MuscleUseCase
}

func NewUseCases(domainServices domain.DomainServices) UseCases {
	return &useCases{
		Muscle: NewMuscleUseCase(domainServices.Muscle),
	}
}

func (u *useCases) MuscleUseCase() MuscleUseCase {
	return u.Muscle
}
