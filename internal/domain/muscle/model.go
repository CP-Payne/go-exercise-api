package muscle

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidMuscle = errors.New("a muscle must have a name")
)

type Muscle struct {
	id   uuid.UUID
	name string
}

func NewMuscle(name string) (Muscle, error) {
	if name == "" {
		return Muscle{}, ErrInvalidMuscle
	}

	return Muscle{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (m *Muscle) GetId() uuid.UUID {
	return m.id
}

func (m *Muscle) SetId(id uuid.UUID) {
	m.id = id
}

func (m *Muscle) GetName() string {
	return m.name
}

func (m *Muscle) SetName(name string) {
	m.name = name
}
