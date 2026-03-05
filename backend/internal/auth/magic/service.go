package magic

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	authEvent "terra/internal/auth/event"
	"terra/internal/auth/token"
	"terra/internal/auth/user"
	"terra/internal/session"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	users    *user.Repository
	tokens   *token.Repository
	sessions *session.Repository
	authSvc  *authEvent.AuthEventService
	mailer   Mailer
}

func NewService(u *user.Repository, t *token.Repository, s *session.Repository, a *authEvent.AuthEventService, m Mailer) *Service {
	return &Service{users: u, tokens: t, sessions: s, mailer: m}
}

func (s *Service) RequestMagicLink(ctx context.Context, email string) error {
	u, err := s.users.FindByEmail(ctx, email)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {

			u, err = s.users.Create(email)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	rawToken, err := GenerateToken()

	if err != nil {
		return err
	}

	hash := HashToken(rawToken)

	loginToken := &token.LoginToken{
		ID:        uuid.NewString(),
		UserID:    u.ID,
		Hash:      hash,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Used:      false,
	}

	if err := s.tokens.Create(ctx, loginToken); err != nil {
		return err
	}

	link := fmt.Sprintf("http://localhost:8080/auth/magic-link/verify?token=%s", rawToken)

	// return s.mailer.SendMagicLink(email, link)
	log.Println("Login link", link)
	return nil
}

func (s *Service) VerifyMagicToken(ctx context.Context, rawToken string) (string, error) {
	hash := HashToken(rawToken)

	token, err := s.tokens.FindByHash(ctx, hash)
	if err != nil {
		return "", err
	}

	if token == nil {
		return "", errors.New("invalid token")
	}

	// 2️⃣ validate
	if token.Used {
		return "", errors.New("token already used")
	}

	if time.Now().After(token.ExpiresAt) {
		return "", errors.New("token expired")
	}

	// 3️⃣ mark used
	if err := s.tokens.MarkUsed(ctx, token.ID); err != nil {
		return "", err
	}

	session := &session.Session{
		ID:        uuid.NewString(),
		UserID:    token.UserID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.sessions.Create(ctx, session); err != nil {
		return "", err
	}

	return session.ID, nil
}
