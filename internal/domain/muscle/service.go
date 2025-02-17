package muscle

import (
	"context"

	"github.com/google/uuid"
)

type IMuscleService interface {
	AddMuscle(context.Context, uuid.UUID, *Muscle) error
	RemoveMuscle(context.Context, uuid.UUID) error
	GetMuscles(context.Context, uuid.UUID) ([]*Muscle, error)
	GetMuscleById(context.Context, uuid.UUID) (*Muscle, error)
}

type MuscleService struct {
	muscleRepository MuscleRepository
}

func NewMuscleService(repo MuscleRepository) *MuscleService {
	return &MuscleService{
		muscleRepository: repo,
	}
}

func (s *MuscleService) AddMuscle(ctx context.Context, userId uuid.UUID, muscle *Muscle) error {
	return s.muscleRepository.Add(ctx, userId, muscle)
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
