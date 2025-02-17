package repositories

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repositories struct {
	Muscles *TargetMuscleRepository
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
