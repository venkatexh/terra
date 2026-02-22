package project

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProject(ctx context.Context, p *Project) error {
	return s.repo.Create(ctx, p)
}

func (s *Service) FindByUserID(ctx context.Context, userID string) ([]Project, error) {
	return s.repo.FindByUserID(ctx, userID)
}
