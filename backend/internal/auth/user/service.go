package user

import (
	"context"
)

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

// func (s *Service) Login(w http.ResponseWriter, r *http.Request) {
// 	s.repo.Login(w, r)
// }

// func (s *Service) Logout(w http.ResponseWriter, r *http.Request) {
// 	s.repo.Logout(w, r)
// }

func (s *Service) Me(ctx context.Context, id string) (*User, error) {
	return s.repo.FindByID(ctx, id)
}