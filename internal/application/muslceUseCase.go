package application

import (
	"context"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

type MuscleUseCase interface {
	AddMuscle(context.Context, uuid.UUID, *muscle.Muscle) error
	ListMusclesForUser(context.Context, uuid.UUID) (*muscle.Muscle, error)
}

type muscleUseCase struct {
	muscleService muscle.MuscleService
}

func NewMuscleUseCase(muscleService muscle.MuscleService) *muscleUseCase {
	return &muscleUseCase{
		muscleService: muscleService,
	}
}

func (us *muscleUseCase) AddMuscle(ctx context.Context, userID uuid.UUID, muscle *muscle.Muscle) error {
	return nil
}

func (us *muscleUseCase) ListMusclesForUser(ctx context.Context, userID uuid.UUID) (*muscle.Muscle, error) {
	return &muscle.Muscle{}, nil
}
