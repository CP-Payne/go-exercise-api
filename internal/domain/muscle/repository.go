package muscle

import "github.com/google/uuid"

type MuscleRepository interface {
	Save(muscle Muscle) error
	FindById(id uuid.UUID) (Muscle, error)
	FindAll() ([]Muscle, error)
	Delete(id uuid.UUID) error
}
