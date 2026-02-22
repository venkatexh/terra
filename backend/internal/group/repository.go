package group

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, userID, name string) error {

	query := `
		INSERT INTO app_groups (user_id, name, created_at, updated_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, userID, name, time.Now(), time.Now())

	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Group, error) {

	query := `
		SELECT id, name, created_at
		FROM app_groups
		WHERE id = $1
	`
	row := r.db.QueryRow(ctx, query, id)

	var group Group
	err := row.Scan(&group.ID, &group.Name, &group.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &group, nil
}

func (r *Repository) FindByUserID(ctx context.Context, userID string) ([]Group, error) {
	query := `
		SELECT id, name, created_at
		FROM app_groups
		WHERE user_id = $1
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var groups []Group
	for rows.Next() {
		var g Group
		err := rows.Scan(&g.ID, &g.Name, &g.CreatedAt)
		if err != nil {
			return nil, err
		}
		groups = append(groups, g)
	}

	return groups, nil
}
