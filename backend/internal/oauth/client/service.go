package client

import (
	"context"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	repo *Repository
}

func NewService(r *Repository) *Service {
	return &Service{repo: r}
}

func (s *Service) RegisterClient(ctx context.Context, name string, uris []string) (*Client, error) {

	client := &Client{
		ID: uuid.NewString(),
		Name: name,
		ClientID: GenerateClientID(),
		ClientSecret: GenerateSecret(),
		RedirectURIs: uris,
		CreatedAt: time.Now(),
	}

	err := s.repo.Create(ctx, client)
	if err != nil {
		return nil, err
	}

	return client, nil
}