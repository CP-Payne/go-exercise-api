package exercise

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// Errors
var (
	ErrInvalidExercise = errors.New("an exercise must have a name")
)

// Entities
type Muscle struct {
	Id   uuid.UUID
	Name string
}

type Equipment struct {
	Id   uuid.UUID
	Name string
}

type Split struct {
	Id   uuid.UUID
	Name string
}

// Aggregates
type Exercise struct {
	id            uuid.UUID
	name          string
	description   string
	displayImage  url.URL
	split         []*Split
	targetMuscles []*Muscle
	equipment     []*Equipment
	category      string
	createdAt     time.Time
	updatedAt     time.Time
}

func NewExercise(name, description, category string) (Exercise, error) {
	if name == "" {
		return Exercise{}, ErrInvalidExercise
	}

	return Exercise{
		id:            uuid.New(),
		name:          name,
		description:   description,
		displayImage:  url.URL{},
		split:         make([]*Split, 0),
		targetMuscles: make([]*Muscle, 0),
		equipment:     make([]*Equipment, 0),
		category:      category,
		createdAt:     time.Now(),
		updatedAt:     time.Now(),
	}, nil
}
