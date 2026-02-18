package session

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

func (r *Repository) Create(ctx context.Context, s *Session) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO user_sessions (id, user_id, expires_at)
		VALUES ($1, $2, $3)
	`, s.ID, s.UserID, s.ExpiresAt)
	return err
}

func (r *Repository) FindByID(ctx context.Context, id string) (*Session, error) {
	row := r.db.QueryRow(ctx, `
		SELECT id, user_id, expires_at
		FROM user_sessions
		WHERE id = $1
	`, id)

	var s Session
	err := row.Scan(
		&s.ID,
		&s.UserID,
		&s.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &s, nil
}