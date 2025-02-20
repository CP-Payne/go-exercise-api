package application

import (
	"context"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

type MuscleUseCase interface {
	CreateMuscle(ctx context.Context, userID uuid.UUID, muscle *muscle.Muscle) error
	ListMusclesForUser(ctx context.Context, userID uuid.UUID) ([]*muscle.Muscle, error)
	DeleteMuscle(ctx context.Context, userID, muscleID uuid.UUID) error
	GetMuscleByID(ctx context.Context, userID, muscleID uuid.UUID) (*muscle.Muscle, error)
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

func (us *muscleUseCase) ListMusclesForUser(ctx context.Context, userID uuid.UUID) ([]*muscle.Muscle, error) {
	return us.muscleService.ListMuscles(ctx, userID)
}

func (us *muscleUseCase) GetMuscleByID(ctx context.Context, userID, muscleID uuid.UUID) (*muscle.Muscle, error) {
	return us.muscleService.GetMuscleByID(ctx, userID, muscleID)
}

func (us *muscleUseCase) DeleteMuscle(ctx context.Context, userID, muscleID uuid.UUID) error {
	return us.muscleService.RemoveMuscle(ctx, userID, muscleID)
}
