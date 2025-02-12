package equipment

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidEquipment = errors.New("an equipment must have a name")
)

type Equipment struct {
	id   uuid.UUID
	name string
}

func NewSplit(name string) (Equipment, error) {
	if name == "" {
		return Equipment{}, ErrInvalidEquipment
	}
	return Equipment{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (m *Equipment) GetId() uuid.UUID {
	return m.id
}

func (m *Equipment) GetName() string {
	return m.name
}

func (m *Equipment) SetName(name string) {
	m.name = name
}
