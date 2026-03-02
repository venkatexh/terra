package otp

import (
	"context"
	// "time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db  *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, otp OTP) error {
	query := `
		INSERT INTO otps (id, user_id, otp_hash, expires_at)
		VALUES ($1, $2, $3, $4)
	`
	_, err := r.db.Exec(ctx, query, otp.ID, otp.UserID, otp.Hash, otp.ExpiresAt)
	return err
}

func (r *Repository) FindByHash(ctx context.Context, hash string) (*OTP, error) {
	query := `
		SELECT id, user_id, otp_hash, expires_at, used
		FROM otps
		WHERE otp_hash = $1
	`

	row := r.db.QueryRow(ctx, query, hash)

	var otp OTP
	err := row.Scan(
		&otp.ID,
		&otp.UserID,
		&otp.Hash,
		&otp.ExpiresAt,
		&otp.Used,
	)

	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *Repository) MarkUsed(ctx context.Context, id string) error {
	return nil
}

