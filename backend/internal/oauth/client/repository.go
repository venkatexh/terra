package client

import (
	"context"
	"terra/internal/middleware"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Repository struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(ctx context.Context, c *Client, redirectUris []string) error {

	tx, err := r.db.Begin(ctx)

	if err != nil {
		return err
	}

	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx,
		`INSERT INTO oauth_clients(id, name, client_id, client_secret, project_id)
		 VALUES ($1, $2, $3, $4, $5)`,
		c.ID,
		c.Name,
		c.ClientID,
		c.ClientSecret,
		c.ProjectID,
	)

	if err != nil {
		return err
	}

	for _, uri := range redirectUris {
		_, err = tx.Exec(ctx,
			`INSERT INTO oauth_client_redirect_uris(id, client_db_id, uri)
			 VALUES ($1, $2, $3)`,
			uuid.New().String(),
			c.ID,
			uri,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *Repository) FindByUserID(ctx context.Context, projectID string) ([]ClientResponse, error) {

	query := `
		SELECT id, name, client_id, client_secret, project_id
		FROM oauth_clients
		WHERE project_id = $1
	`

	rows, err := r.db.Query(ctx, query, projectID)

	var clients []ClientResponse
	for rows.Next() {
		var c ClientResponse
		err = rows.Scan(&c.ID, &c.Name, &c.ClientID, &c.ClientSecret, &c.ProjectID)

		c.RedirectURIs, err = r.FindRedirectURIsByClientID(ctx, c.ID)

		clients = append(clients, c)
	}

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *Repository) FindByClientID(ctx context.Context, clientID string) (*ClientResponse, error) {

	query := `
		SELECT c.id, c.name, c.client_id, c.client_secret, c.project_id
		FROM oauth_clients c
		JOIN oauth_projects p on c.project_id = p.id
		WHERE c.client_id = $1
		AND p.user_id = $2
	`
	sessionUserID := ctx.Value(middleware.UserKey).(string)
	row := r.db.QueryRow(ctx, query, clientID, sessionUserID)

	var c ClientResponse
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.ClientID,
		&c.ClientSecret,
		&c.ProjectID,
	)

	if err != nil {
		return nil, err
	}

	c.RedirectURIs, err = r.FindRedirectURIsByClientID(ctx, c.ID)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) FindRedirectURIsByClientID(ctx context.Context, clientID string) ([]RedirectURI, error) {

	query := `
		SELECT id, uri, client_db_id
		FROM oauth_client_redirect_uris
		WHERE client_db_id = $1
	`
	rows, err := r.db.Query(ctx, query, clientID)

	var uris []RedirectURI
	for rows.Next() {
		var uri RedirectURI
		err = rows.Scan(&uri.ID, &uri.URI, &uri.ClientDBID)
		if err != nil {
			return nil, err
		}
		uris = append(uris, uri)
	}

	return uris, nil
}
