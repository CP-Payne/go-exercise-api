package split

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSplit = errors.New("a split must have a name")
)

type Split struct {
	id   uuid.UUID
	name string
}

func NewSplit(name string) (Split, error) {
	if name == "" {
		return Split{}, ErrInvalidSplit
	}
	return Split{
		id:   uuid.New(),
		name: name,
	}, nil
}

func (m *Split) GetId() uuid.UUID {
	return m.id
}

func (m *Split) GetName() string {
	return m.name
}

func (m *Split) SetName(name string) {
	m.name = name
}
