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
		`INSERT INTO oauth_clients(id, name, client_id, client_secret)
		 VALUES ($1, $2, $3, $4)`,
		c.ID,
		c.Name,
		c.ClientID,
		c.ClientSecret,
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
