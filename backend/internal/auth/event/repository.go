package authEvent

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthEventRepository struct {
	db *pgxpool.Pool
}

func NewAuthEventRepository(db *pgxpool.Pool) *AuthEventRepository {
	return &AuthEventRepository{db: db}
}

func (r *AuthEventRepository) Create(ctx context.Context, e *AuthEvent) error {
	query := `
		INSERT INTO auth_events (user_id, client_db_id, event_type, status, ip_address, user_agent, metadata)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`
	_, err := r.db.Exec(ctx, query, e.UserID, e.ClientID, e.EventType, e.Status, e.IPAddress, e.UserAgent, e.Metadata)
	return err
}
