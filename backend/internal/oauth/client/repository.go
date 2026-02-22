package client

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

func (r *Repository) Create(ctx context.Context, c *Client) error {

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

	for _, uri := range c.RedirectURIs {
		_, err = tx.Exec(ctx,
			`INSERT INTO oauth_client_redirect_uris(client_db_id, uri)
			 VALUES ($1, $2)`,
			c.ID,
			uri,
		)
		if err != nil {
			return err
		}
	}

	return tx.Commit(ctx)
}

func (r *Repository) FindByUserID(ctx context.Context, projectID string) ([]Client, error) {

	query := `
		SELECT id, name, client_id, client_secret, project_id
		FROM oauth_clients
		WHERE project_id = $1
	`

	rows, err := r.db.Query(ctx, query, projectID)

	var clients []Client
	for rows.Next() {
		var c Client
		err = rows.Scan(&c.ID, &c.Name, &c.ClientID, &c.ClientSecret, &c.ProjectID)

		c.RedirectURIs, err = r.FindRedirectURIsByClientID(ctx, c.ID)

		clients = append(clients, c)
	}

	if err != nil {
		return nil, err
	}

	return clients, nil
}

func (r *Repository) GetByClientID(ctx context.Context, clientID string) (*Client, error) {
	row := r.db.QueryRow(ctx, `
	SELECT id, name, client_id, client_secret
	FROM oauth_clients
	WHERE client_id = $1
	`, clientID)

	var c Client
	err := row.Scan(
		&c.ID,
		&c.Name,
		&c.ClientID,
		&c.ClientSecret,
	)

	if err != nil {
		return nil, err
	}

	return &c, nil
}

func (r *Repository) FindRedirectURIsByClientID(ctx context.Context, clientID string) ([]string, error) {

	query := `
		SELECT uri
		FROM oauth_client_redirect_uris
		WHERE client_db_id = $1
	`
	rows, err := r.db.Query(ctx, query, clientID)

	var uris []string
	for rows.Next() {
		var uri string
		err = rows.Scan(&uri)
		if err != nil {
			return nil, err
		}
		uris = append(uris, uri)
	}

	return uris, nil
}
