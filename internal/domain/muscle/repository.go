package muscle

import (
	"context"

	"github.com/google/uuid"
)

// MuscleRepository defines the storage operations for Muscle entities
type MuscleRepository interface {
	Add(ctx context.Context, userId uuid.UUID, muscle *Muscle) error
	GetByID(ctx context.Context, userID, muscleID uuid.UUID) (*Muscle, error)
	List(ctx context.Context, userId uuid.UUID) ([]*Muscle, error)
	Delete(ctx context.Context, userID, muscleID uuid.UUID) error
}
