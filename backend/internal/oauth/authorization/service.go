package authorization

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) CreateAuthorization(ctx context.Context, groupID, clientDBID, userID string) error {
	return s.repo.Create(ctx, groupID, clientDBID, userID)
}

func (s *Service) HasConsent(ctx context.Context, userID, clientID string) (bool, error) {
	a, err := s.repo.FindByUserIDandClientDBID(ctx, userID, clientID)
	if err != nil {
		return false, nil
	}

	return a != nil, nil
}

func (s *Service) Grant(ctx context.Context, auth *Authorization) error {
	return s.repo.Upsert(ctx, auth)
}

func (s *Service) FindByGroupID(ctx context.Context, groupID string) ([]AuthorizationDTO, error) {
	return s.repo.FindByGroupID(ctx, groupID)
}
