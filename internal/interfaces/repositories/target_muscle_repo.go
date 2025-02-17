package repositories

import (
	"context"
	"database/sql"
	"time"

	"github.com/CP-Payne/exercise/internal/domain/muscle"
	"github.com/google/uuid"
)

var (
	QueryTimeoutDuration = time.Second * 5
)

type TargetMuscleRepository struct {
	db *sql.DB
}

func NewTargetMuscleRepository(db *sql.DB) *TargetMuscleRepository {
	return &TargetMuscleRepository{db: db}
}

type PostgresMuscle struct {
	Id        uuid.UUID
	Name      string
	UserId    uuid.UUID
	CreatedAt time.Time
}

func (r *TargetMuscleRepository) Add(ctx context.Context, userId uuid.UUID, muscle *muscle.Muscle) error {
	query := `
		INSERT INTO target_muscles (id, muscle_name, user_id, created_at)
		VALUES($1, $2, $3, $4)
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	_, err := r.db.ExecContext(ctx,
		query,
		muscle.GetId(),
		muscle.GetName(),
		userId,
		time.Now(),
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *TargetMuscleRepository) GetById(ctx context.Context, id uuid.UUID) (*muscle.Muscle, error) {
	return &muscle.Muscle{}, nil
}

func (r *TargetMuscleRepository) GetAll(ctx context.Context, userId uuid.UUID) ([]*muscle.Muscle, error) {
	query := `
		SELECT * FROM target_muscles
		WHERE user_id = $1
	`

	ctx, cancel := context.WithTimeout(ctx, QueryTimeoutDuration)
	defer cancel()

	rows, err := r.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	muscles := []*muscle.Muscle{}

	for rows.Next() {
		var pm PostgresMuscle
		err := rows.Scan(&pm.Id, &pm.Name, &pm.UserId, &pm.CreatedAt)
		if err != nil {
			return nil, err
		}
		muscles = append(muscles, PostgresMuscleToMuscle(pm))
	}

	return []*muscle.Muscle{}, nil
}

func (r *TargetMuscleRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return nil
}

func PostgresMuscleToMuscle(pm PostgresMuscle) *muscle.Muscle {
	m := muscle.Muscle{}
	m.SetId(pm.Id)
	m.SetName(pm.Name)
	return &m
}
