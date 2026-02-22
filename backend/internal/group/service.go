package group

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateGroup(ctx context.Context, userID, name string) error {
	return s.repo.Create(ctx, userID, name)
}

func (s *Service) FindGroupByID(ctx context.Context, id string) (*Group, error) {
	return s.repo.FindByID(ctx, id)
}

func (s *Service) FindByUserID(ctx context.Context, userID string) ([]Group, error) {
	return s.repo.FindByUserID(ctx, userID)
}