package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(email string) (*User, error) {
	user := &User{
		ID:    uuid.NewString(),
		Email: email,
	}

	query := `
		INSERT INTO users (id, email)
		VALUES ($1, $2)
		RETURNING created_at
	`

	err := r.db.QueryRow(context.Background(), query,
		user.ID,
		user.Email,
	).Scan(&user.CreatedAt)

	return user, err
}

func (r *Repository) FindByEmail(ctx context.Context, email string) (*User, error) {
	query := `
		SELECT id, email, email_verified, created_at
		FROM users
		WHERE email = $1
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

func (r *Repository) FindByID(ctx context.Context, id string) (*User, error) {
	query := `
		SELECT id, email, name, email_verified, created_at
		FROM users
		WHERE id = $1
	`
	row := r.db.QueryRow(ctx, query, id)

	var u User
	err := row.Scan(
		&u.ID,
		&u.Email,
		&u.Name,
		&u.EmailVerified,
		&u.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &u, nil
}
