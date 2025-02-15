package muscle

import (
	"context"

	"github.com/google/uuid"
)

type MuscleService struct {
	muscleRepository MuscleRepository
}

func (s *MuscleService) AddMuscle(ctx context.Context, muscle *Muscle) error {
	return s.muscleRepository.Add(ctx, muscle)
}

func (s *MuscleService) RemoveMuscle(ctx context.Context, id uuid.UUID) error {
	return s.muscleRepository.Delete(ctx, id)
}

func (s *MuscleService) GetMuscles(ctx context.Context, userId uuid.UUID) ([]*Muscle, error) {
	return s.muscleRepository.GetAll(ctx, userId)
}

func (s *MuscleService) GetMuscleById(ctx context.Context, id uuid.UUID) (*Muscle, error) {
	return s.muscleRepository.GetById(ctx, id)
}
