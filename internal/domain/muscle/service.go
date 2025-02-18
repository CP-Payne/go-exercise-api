package muscle

import (
	"context"

	"github.com/google/uuid"
)

type MuscleService interface {
	AddMuscle(context.Context, uuid.UUID, *Muscle) error
	RemoveMuscle(context.Context, uuid.UUID) error
	ListMuscles(context.Context, uuid.UUID) ([]*Muscle, error)
	GetMuscleByID(context.Context, uuid.UUID) (*Muscle, error)
}

type muscleService struct {
	repo MuscleRepository
}

func NewMuscleService(repo MuscleRepository) MuscleService {
	return &muscleService{
		repo: repo,
	}
}

func (s *muscleService) AddMuscle(ctx context.Context, userID uuid.UUID, muscle *Muscle) error {
	return s.repo.Add(ctx, userID, muscle)
}

func (s *muscleService) RemoveMuscle(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}

func (s *muscleService) ListMuscles(ctx context.Context, userID uuid.UUID) ([]*Muscle, error) {
	return s.repo.List(ctx, userID)
}

func (s *muscleService) GetMuscleByID(ctx context.Context, id uuid.UUID) (*Muscle, error) {
	return s.repo.GetByID(ctx, id)
}
