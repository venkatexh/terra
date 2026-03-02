package authorization

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

func (r *Repository) Create(ctx context.Context, groupID, clientDBID string) error {
	ID := uuid.NewString()
	_, err := r.db.Exec(ctx, `
		INSERT INTO oauth_authorizations(id, group_id, client_db_id)
		VALUES ($1, $2, $3)
	`, ID, groupID, clientDBID)
	return err
}

func (r *Repository) Upsert(ctx context.Context, a *Authorization) error {

	query := `
		INSERT INTO oauth_authorizations(id, group_id, client_db_id, scopes)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (group_id, client_db_id)
		DO UPDATE SET scopes = EXCLUDED.scopes
	`

	_, err := r.db.Exec(
		ctx, query, a.ID, a.GroupID, a.ClientDBID, a.Scopes,
	)

	return err
}

func (r *Repository) CheckExists(ctx context.Context, groupID, clientDBID string) (bool, error) {
	query := `
		SELECT EXISTS (
			SELECT 1 FROM oauth_authorizations
			WHERE user_id = $1 AND client_db_id = $2
		)
	`

	var exists bool
	err := r.db.QueryRow(ctx, query, groupID, clientDBID).Scan(&exists)
	return exists, err
}
func (r *Repository) FindByGroupID(ctx context.Context, groupID string) ([]AuthorizationDTO, error) {

	query := `
		SELECT id, group_id, client_db_id, created_at
		FROM oauth_authorizations
		WHERE group_id = $1
	`

	rows, err := r.db.Query(ctx, query, groupID)
	if err != nil {
		return nil, err
	}

	authorizations := make([]AuthorizationDTO, 0)
	for rows.Next() {
		var a AuthorizationDTO
		query := `
			SELECT name
			FROM oauth_clients
			WHERE id = $1
		`
		err := rows.Scan(&a.ID, &a.GroupID, &a.ClientDBID, &a.CreatedAt)
		if err != nil {
			return nil, err
		}

		var name string
		clientNameErr := r.db.QueryRow(ctx, query, a.ClientDBID).Scan(&name)

		if clientNameErr != nil {
			return nil, clientNameErr
		}

		a.ClientName = name
		authorizations = append(authorizations, a)
	}

	return authorizations, nil
}

func (r *Repository) FindByUserIDandClientDBID(ctx context.Context, userID, clientID string) (*Authorization, error) {

	query := `
		SELECT id, user_id, client_id, scopes, created_at
		FROM oauth_authorizations
		WHERE user_id = $1 AND client_db_id = $2
	`

	var a Authorization
	err := r.db.QueryRow(ctx, query, userID, clientID).Scan(&a.ID, &a.GroupID, &a.ClientDBID, &a.Scopes, &a.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &a, nil
}
