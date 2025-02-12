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

// Aggregates
type Exercise struct {
	id              uuid.UUID
	name            string
	description     string
	displayImage    url.URL
	splitIDs        []uuid.UUID
	targetMuscleIDs []uuid.UUID
	equipmentIDs    []uuid.UUID
	category        string
	createdAt       time.Time
	updatedAt       time.Time
}

func NewExercise(name, description, category string) (Exercise, error) {
	if name == "" {
		return Exercise{}, ErrInvalidExercise
	}

	return Exercise{
		id:              uuid.New(),
		name:            name,
		description:     description,
		displayImage:    url.URL{},
		splitIDs:        make([]uuid.UUID, 0),
		targetMuscleIDs: make([]uuid.UUID, 0),
		equipmentIDs:    make([]uuid.UUID, 0),
		category:        category,
		createdAt:       time.Now(),
		updatedAt:       time.Now(),
	}, nil
}
