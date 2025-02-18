package application

import (
	"context"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

type MuscleUseCase interface {
	CreateMuscle(context.Context, uuid.UUID, *muscle.Muscle) error
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

func (us *muscleUseCase) CreateMuscle(ctx context.Context, userID uuid.UUID, muscle *muscle.Muscle) error {
	err := us.muscleService.AddMuscle(ctx, userID, muscle)
	if err != nil {
		return err
	}
	return nil
}

func (us *muscleUseCase) ListMusclesForUser(ctx context.Context, userID uuid.UUID) (*muscle.Muscle, error) {
	return &muscle.Muscle{}, nil
}
