package muscle

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidMuscle = errors.New("a muscle must have a name")
)

type MuscleParams struct {
	ID   uuid.UUID
	Name string
}

type Muscle struct {
	id   uuid.UUID
	name string
}

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

func (m *Muscle) Name() string { return m.name }
