package authcode

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

func (r *Repository) Create(ctx context.Context, c *AuthorizationCode) error {
	_, err := r.db.Exec(ctx, `
		INSERT INTO oauth_authorization_codes(code, client_id, user_id, redirect_uri, expires_at)
		VALUES ($1, $2, $3, $4, $5)
	`,
		c.Code,
		c.ClientID,
		c.UserID,
		c.RedirectURI,
		c.ExpiresAt,
	)
	return err
}

func (r *Repository) Get(ctx context.Context, code string) (*AuthorizationCode, error) {
	row := r.db.QueryRow(ctx, `
	SELECT code, client_id, user_id, redirect_uri, expires_at
	FROM oauth_authorization_codes
	WHERE code = $1
	`, code)

	var c AuthorizationCode
	err := row.Scan(
		&c.Code,
		&c.ClientID,
		&c.UserID,
		&c.RedirectURI,
		&c.ExpiresAt,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) Delete(ctx context.Context, code string) error {
	_, err := r.db.Exec(ctx, `
	DELETE FROM oauth_authorization_codes
	WHERE code = $1
	`, code)
	return err
}