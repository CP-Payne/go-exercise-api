package muscle

import (
	"context"

	"github.com/google/uuid"
)

type MuscleRepository interface {
	Add(ctx context.Context, userId uuid.UUID, muscle *Muscle) error
	GetById(ctx context.Context, id uuid.UUID) (*Muscle, error)
	GetAll(ctx context.Context, userId uuid.UUID) ([]*Muscle, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
