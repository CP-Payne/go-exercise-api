package application

import (
	"context"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

type MuscleApplication interface {
	AddTargetMuscle(context.Context, uuid.UUID, muscle.Muscle) error
	GetTargetMusclesForUser(context.Context, uuid.UUID) (*muscle.Muscle, error)
}

type MuscleUseCase struct {
	muscleService muscle.IMuscleService
}

func NewMuscleUseCase(muscleService muscle.IMuscleService) *MuscleUseCase {
	return &MuscleUseCase{
		muscleService: muscleService,
	}
}

func (us *MuscleUseCase) AddTargetMuscle(ctx context.Context, userId uuid.UUID, muscle muscle.Muscle) error {
	return nil
}

func (us *MuscleUseCase) GetTargetMusclesForUser(ctx context.Context, userId uuid.UUID) (*muscle.Muscle, error) {
	return &muscle.Muscle{}, nil
}
