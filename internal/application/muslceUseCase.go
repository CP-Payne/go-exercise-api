package application

import (
	"context"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

type MuscleUseCase struct {
	muscleService muscle.MuscleService
}

func (us *MuscleUseCase) AddTargetMuscle(ctx context.Context, userId uuid.UUID, muscle muscle.Muscle) error {
	return nil
}

func (us *MuscleUseCase) GetTargetMusclesForUser(ctx context.Context, userId uuid.UUID) (*muscle.Muscle, error) {
	return &muscle.Muscle{}, nil
}
