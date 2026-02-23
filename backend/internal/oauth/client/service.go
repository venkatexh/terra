package client

import (
	"context"
	"errors"
	"terra/internal/oauth/project"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	clientRepo  *Repository
	projectRepo *project.Repository
}

func NewService(r *Repository, pr *project.Repository) *Service {
	return &Service{clientRepo: r, projectRepo: pr}
}

func (s *Service) RegisterClient(ctx context.Context, name string, uris []string, projectID, userID string) (*Client, error) {

	ok, err := s.projectRepo.ExistsByIDAndUser(
		ctx,
		projectID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("unauthorized project access")
	}

	client := &Client{
		ID:           uuid.NewString(),
		Name:         name,
		ClientID:     GenerateClientID(),
		ClientSecret: GenerateSecret(),
		ProjectID:    projectID,
		CreatedAt:    time.Now(),
	}

	err = s.clientRepo.Create(ctx, client, uris)

	if err != nil {
		return nil, err
	}

	return client, nil
}

func (s *Service) FindByUserID(ctx context.Context, userID, projectID string) ([]ClientResponse, error) {

	ok, err := s.projectRepo.ExistsByIDAndUser(
		ctx,
		projectID,
		userID,
	)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("unauthorized project access")
	}

	return s.clientRepo.FindByUserID(ctx, projectID)
}

func (s *Service) FindClientByClientID(ctx context.Context, clientID string) (*ClientResponse, error) {
	return s.clientRepo.FindByClientID(ctx, clientID)
}
