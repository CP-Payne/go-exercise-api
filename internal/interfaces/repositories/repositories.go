package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	_ "github.com/lib/pq"
)

// Errors used by all repository implementations
var (
	// ErrNotFound indicates the requested resource does not exist
	ErrNotFound = errors.New("resource not found")

	// ErrConflict indicates a conflict with an existing resource
	// (e.g., unique constraint violation)
	ErrConflict = errors.New("resource already exists")
)

var (
	// QueryTimeoutDuration defines the standard timeout for database operations
	QueryTimeoutDuration = time.Second * 5
)

// Repositories provides access to all repository implementations
// in a central location for dependecy injection
type Repositories struct {
	Muscles muscle.MuscleRepository
}

// NewRepositories creates and initializes all repository implementations
func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Muscles: NewTargetMuscleRepository(db),
	}
}

// withTx executes the provided function within a database transaction
// and handles commit/rollback automatically based on function result
func withTx(db *sql.DB, ctx context.Context, fn func(*sql.Tx) error) error {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
