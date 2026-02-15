package auth

import (
	"context"
	"encoding/hex"
	"log"
	"time"
	"crypto/sha256"
	"errors"

	"terra/internal/token"
	"terra/internal/user"

	"github.com/google/uuid"
)

type Service struct {
	users *user.Repository
	tokens *token.Repository
}

func NewService(u *user.Repository, t *token.Repository) *Service {
	return &Service{
		users: u,
		tokens: t,
	}
}

func (s *Service) RequestLoginLink(ctx context.Context, email string) error {
	u, err := s.users.FindByEmail(ctx, email)
	if err != nil {
		u = &user.User{ID: uuid.NewString(), Email: email}

		err = s.users.Create(ctx, u)
		if err != nil {
			return err
		}
	}

	raw, hash, err := token.GenerateToken()
	if err != nil {
		return err
	}

	loginToken := &token.LoginToken{
		ID: uuid.NewString(),
		UserID: u.ID,
		Hash: hash,
		ExpiresAt: time.Now().Add(15 * time.Minute),	
	}

	err = s.tokens.Create(ctx, loginToken)
	if err != nil {
		return err
	}

	loginLink := "http://localhost:3000/login?token=" + raw
	log.Println("Magic login link:", loginLink)

	return nil
}

func (s *Service) VerifyLoginToken(ctx context.Context, rawToken string) (string,error) {
	hashBytes := sha256.Sum256([]byte(rawToken))
	hash := hex.EncodeToString(hashBytes[:])

	t, err := s.tokens.FindByHash(ctx, hash)

	if err != nil {
		return "", err
	}

	if t.Used {
		return "", errors.New("token already used")
	}

	if time.Now().After(t.ExpiresAt){
		return "", errors.New("token expired")
	}

	err = s.tokens.MarkUsed(ctx, t.ID)
	if err != nil {
		return "", err
	}

	jwtToken, err := GenerateJWT(t.UserID)

	if err != nil {
		return "", err
	}

	return jwtToken, nil
}