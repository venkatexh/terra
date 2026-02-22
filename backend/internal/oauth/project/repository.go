package project

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, p *Project) error {
	query := `
		INSERT INTO oauth_projects
		(id, user_id, name, description)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query, p.ID, p.UserID, p.Name, p.Description)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) FindByUserID(ctx context.Context, userID string) ([]Project, error) {
	query := `
		SELECT id, user_id, name, description
		FROM oauth_projects
		WHERE user_id = $1
	`

	rows, err := r.db.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}

	var projects []Project
	for rows.Next() {
		var p Project
		err := rows.Scan(&p.ID, &p.UserID, &p.Name, &p.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, p)
	}

	return projects, nil
}

func (r *Repository) ExistsByIDAndUser(
	ctx context.Context,
	projectID, userID string,
) (bool, error) {

	query := `
    SELECT EXISTS (
    SELECT 1 FROM oauth_projects
    WHERE id=$1 AND user_id=$2
  )`
	var exists bool

	err := r.db.QueryRow(ctx, query, projectID, userID).Scan(&exists)
	return exists, err
}
