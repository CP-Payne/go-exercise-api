package exercise

import "github.com/google/uuid"

type ExerciseRepository interface {
	Get(uuid.UUID) (Exercise, error)
	Add(Exercise) error
	Update(Exercise) error
}

