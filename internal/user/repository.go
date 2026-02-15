package user

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

func (r *Repository) Create(ctx context.Context, u *User) error {
	query := `
		INSERT INTO users (id, email) 
		VALUES ($1, $2)
	`

	_, err := r.db.Exec(ctx, query, u.ID, u.Email)
	return err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, email_verified, created_at
		FROM user
		WHERE email = &1
	`

	row := r.db.QueryRow(ctx, query, email)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.EmailVerified,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}