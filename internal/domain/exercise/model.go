package exercise

import (
	"errors"
	"net/url"
	"time"

	"github.com/google/uuid"
)

// Errors
var (
	ErrInvalidExerciseName = errors.New("an exercise must have a name")
	ErrInvalidUpdateDate   = errors.New("date cannot be in the past")
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
		return Exercise{}, ErrInvalidExerciseName
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

func (e *Exercise) SetName(name string) {
	e.name = name
}

func (e *Exercise) GetName() string {
	return e.name
}

func (e *Exercise) SetDescription(description string) {
	e.description = description
}

func (e *Exercise) GetDescription() string {
	return e.description
}

func (e *Exercise) SetDisplayImage(displayImage url.URL) {
	e.displayImage = displayImage
}

func (e *Exercise) GetDisplayImage() url.URL {
	return e.displayImage
}

func (e *Exercise) AddSplit(splitID uuid.UUID) {
	e.splitIDs = append(e.splitIDs, splitID)
}

func (e *Exercise) GetSplits() []uuid.UUID {
	return e.splitIDs
}

func (e *Exercise) AddTargetMuscle(targetMuscleID uuid.UUID) {
	e.targetMuscleIDs = append(e.targetMuscleIDs, targetMuscleID)
}

func (e *Exercise) GetTargetMuscles() []uuid.UUID {
	return e.targetMuscleIDs
}

func (e *Exercise) AddEquipment(equipmentID uuid.UUID) {
	e.equipmentIDs = append(e.equipmentIDs, equipmentID)
}

func (e *Exercise) GetEquipments() []uuid.UUID {
	return e.equipmentIDs
}

func (e *Exercise) SetCategory(category string) {
	e.category = category
}

func (e *Exercise) GetCategory() string {
	return e.category
}

func (e *Exercise) GetUpdatedAt() time.Time {
	return e.updatedAt
}
func (e *Exercise) SetUpdatedAt(t time.Time) error {

	if t.Before(time.Now()) {
		return errors.New("invalid date")
	}
	e.updatedAt = t
	return nil
}
