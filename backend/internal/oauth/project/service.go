package project

import (
	"context"
	"errors"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateProject(ctx context.Context, p *Project) (*Project,error) {
	return s.repo.Create(ctx, p)
}

func (s *Service) FindByUserID(ctx context.Context, userID string) ([]Project, error) {
	return s.repo.FindByUserID(ctx, userID)
}

func (s *Service) FindProjectByID(ctx context.Context, id, userID string) (*Project, error) {
	ok, err := s.repo.ExistsByIDAndUser(
		ctx,
		id,
		userID,
	)

	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, errors.New("unauthorized project access")
	}
	
	return s.repo.FindByID(ctx, id)
}
