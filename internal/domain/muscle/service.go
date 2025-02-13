package muscle

import "github.com/google/uuid"

type MuscleService struct {
	muscleRepository MuscleRepository
}

func (s *MuscleService) AddMuscle(muscle Muscle) error {
	return s.muscleRepository.Add(muscle)
}

func (s *MuscleService) RemoveMuscle(id uuid.UUID) error {
	return s.muscleRepository.Delete(id)
}

func (s *MuscleService) GetMuscles() ([]Muscle, error) {
	return s.muscleRepository.GetAll()
}

func (s *MuscleService) GetMuscleById(id uuid.UUID) (Muscle, error) {
	return s.muscleRepository.GetById(id)
}
