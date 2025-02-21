package muscle

import (
	"context"

	"github.com/google/uuid"
)

// MuscleService defines the business operations available for muscles
type MuscleService interface {
	AddMuscle(ctx context.Context, userID uuid.UUID, muscle *Muscle) error
	RemoveMuscle(ctx context.Context, userID, muscleID uuid.UUID) error
	ListMuscles(ctx context.Context, userID uuid.UUID) ([]*Muscle, error)
	GetMuscleByID(ctx context.Context, userID, muscleID uuid.UUID) (*Muscle, error)
}

type muscleService struct {
	repo MuscleRepository
}

// NewMuscleService create a new service with the provided repository
func NewMuscleService(repo MuscleRepository) MuscleService {
	return &muscleService{
		repo: repo,
	}
}

func (s *muscleService) AddMuscle(ctx context.Context, userID uuid.UUID, muscle *Muscle) error {
	return s.repo.Add(ctx, userID, muscle)
}

func (s *muscleService) RemoveMuscle(ctx context.Context, userID, muscleID uuid.UUID) error {
	return s.repo.Delete(ctx, userID, muscleID)
}

func (s *muscleService) ListMuscles(ctx context.Context, userID uuid.UUID) ([]*Muscle, error) {
	return s.repo.List(ctx, userID)
}

func (s *muscleService) GetMuscleByID(ctx context.Context, userID, muscleID uuid.UUID) (*Muscle, error) {
	return s.repo.GetByID(ctx, userID, muscleID)
}
