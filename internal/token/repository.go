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

func (r *Repository) Create(ctx context.Context, t *LoginToken) error {
	query := `
		INSERT INTO login_tokens (id, user_id, token_hash, expires_at)
		VALUES ($1, $2, $3, $4)
	`

	_, err := r.db.Exec(ctx, query,
		t.ID,
		t.UserID,
		t.Hash,
		t.ExpiresAt,
)

return err
}

func (r *Repository) FindByHash(ctx context.Context, hash string) (*LoginToken, error) {
	query := `
		SELECT id, user_id, token_hash, expires_at, used
		FROM login_tokens
		WHERE token_hash = $1
	`

	row := r.db.QueryRow(ctx, query, hash)

	var t LoginToken
	err := row.Scan(
		&t.ID,
		&t.UserID,
		&t.Hash,
		&t.ExpiresAt,
		&t.Used,
	)

	if err != nil {
		return nil, err
	}

	return &t, nil
}

func (r *Repository) MarkUsed(ctx context.Context, id string) error {
	query := `
		UPDATE login_tokens
		SET used = true
		WHERE id = $1
	`

	_, err := r.db.Exec(ctx, query, id)
	return err
}