package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	_ "github.com/lib/pq"
)

var (
	QueryTimeoutDuration = time.Second * 5
)

type Repositories struct {
	Muscles muscle.MuscleRepository
}

func NewRepositories(db *sql.DB) *Repositories {
	return &Repositories{
		Muscles: NewTargetMuscleRepository(db),
	}
}

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
