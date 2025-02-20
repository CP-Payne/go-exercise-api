package repositories

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

var (
	ErrDuplicateMuscleName = errors.New("a muscle with that name already exists")
)

type TargetMuscleRepository struct {
	db *sql.DB
}

func NewTargetMuscleRepository(db *sql.DB) *TargetMuscleRepository {
	return &TargetMuscleRepository{db: db}
}

type PostgresMuscle struct {
	ID        uuid.UUID
	Name      string
	UserID    uuid.UUID
	CreatedAt time.Time
}

func (r *TargetMuscleRepository) Add(ctx context.Context, userID uuid.UUID, muscle *muscle.Muscle) error {
	query := `
		INSERT INTO target_muscles (id, muscle_name, user_id, created_at)
		VALUES($1, $2, $3, $4)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := r.db.ExecContext(ctx,
		query,
		muscle.ID(),
		muscle.Name(),
		userID,
		time.Now(),
	)
	if err != nil {
		switch {
		case err.Error() == `pq: duplicate key value violates unique constraint "target_muscles_muscle_name_key"`:
			return ErrDuplicateMuscleName
		default:
			return err
		}
	}
	return nil
}

func (r *TargetMuscleRepository) GetByID(ctx context.Context, userID, muscleID uuid.UUID) (*muscle.Muscle, error) {
	query := `
		SELECT id, muscle_name, user_id, created_at FROM target_muscles
		WHERE user_id = $1 AND id = $2	
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	var pm PostgresMuscle

	err := r.db.QueryRowContext(ctx,
		query,
		userID,
		muscleID,
	).Scan(
		&pm.ID,
		&pm.Name,
		&pm.UserID,
		&pm.CreatedAt,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrNotFound
		default:
			return nil, err
		}
	}

	return PostgresMuscleToMuscle(pm)
}

func (r *TargetMuscleRepository) List(ctx context.Context, userID uuid.UUID) ([]*muscle.Muscle, error) {
	query := `
		SELECT id, muscle_name, user_id, created_at FROM target_muscles
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	muscles := []*muscle.Muscle{}

	for rows.Next() {
		var pm PostgresMuscle
		err := rows.Scan(&pm.ID, &pm.Name, &pm.UserID, &pm.CreatedAt)
		if err != nil {
			return nil, err
		}

		m, err := PostgresMuscleToMuscle(pm)
		if err != nil {
			return nil, err
		}

		muscles = append(muscles, m)
	}

	return muscles, nil
}

func (r *TargetMuscleRepository) Delete(ctx context.Context, userID, muscleID uuid.UUID) error {
	query := `
		DELETE FROM target_muscles
		WHERE user_id = $1 AND id = $2	
	`
	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	res, err := r.db.ExecContext(ctx, query, userID, muscleID)
	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows == 0 {
		return ErrNotFound
	}

	return nil
}

func PostgresMuscleToMuscle(pm PostgresMuscle) (*muscle.Muscle, error) {

	return muscle.NewMuscle(muscle.MuscleParams{
		ID:   pm.ID,
		Name: pm.Name,
	})

}
