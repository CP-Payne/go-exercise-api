package application

import "github.com/CP-Payne/exercise/internal/domain"

type ApplicationUseCases struct {
	Muscle MuscleApplication
}

func NewApplicationUseCases(domainServices domain.DomainServices) *ApplicationUseCases {
	return &ApplicationUseCases{
		Muscle: NewMuscleUseCase(domainServices.Muscle),
	}
}
