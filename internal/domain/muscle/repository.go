package muscle

import "github.com/google/uuid"

type MuscleRepository interface {
	Add(muscle Muscle) error
	GetById(id uuid.UUID) (Muscle, error)
	GetAll() ([]Muscle, error)
	Delete(id uuid.UUID) error
}
