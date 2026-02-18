package token

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

func (r *Repository) Create(ctx context.Context, t *AccessToken) error {
	_, err := r.db.Exec(ctx, `
	INSERT INTO oauth_access_tokens
	(token, user_id, client_id, expires_at)
	VALUES ($1, $2, $3, $4)
	`,
		t.Token,
		t.UserID,
		t.ClientID,
		t.ExpiresAt,
	)
	return err
}