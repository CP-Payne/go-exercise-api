package muscle

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrInvalidMuscle is returned when attempting to create a muscle without a name
	ErrInvalidMuscle = errors.New("a muscle must have a name")
)

// MuscleParams contains the parameters needed to create a new Muscle
type MuscleParams struct {
	ID   uuid.UUID
	Name string
}

// Muscle represents a muscle in the exercise system
type Muscle struct {
	id   uuid.UUID
	name string
}

// NewMuscle creates a new Muscle entity with validation
func NewMuscle(params MuscleParams) (*Muscle, error) {
	if params.Name == "" {
		return &Muscle{}, ErrInvalidMuscle
	}

	if params.ID == uuid.Nil {
		params.ID = uuid.New()
	}

	return &Muscle{
		id:   params.ID,
		name: params.Name,
	}, nil
}

func (m *Muscle) ID() uuid.UUID { return m.id }
func (m *Muscle) Name() string  { return m.name }
