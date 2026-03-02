package otp

import (
	"context"
	"database/sql"
	"errors"
	"terra/internal/auth/user"
	"terra/internal/email"
	"terra/internal/session"
	"time"

	"github.com/google/uuid"
)

type Service struct {
	users    *user.Repository
	otps     *Repository
	sessions *session.Repository
	mailer   *email.Service
}

func NewService(u *user.Repository, o *Repository, s *session.Repository, m *email.Service) *Service {
	return &Service{users: u, otps: o, sessions: s, mailer: m}
}

func (s *Service) RequestOTP(ctx context.Context, email string) error {
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

	rawOTP, err := GenerateOTP()

	if err != nil {
		return err
	}

	hash := HashOTP(rawOTP)

	otp := OTP{
		ID:        uuid.NewString(),
		UserID:    u.ID,
		Hash:      hash,
		ExpiresAt: time.Now().Add(15 * time.Minute),
		Used:      false,
	}

	if err := s.otps.Create(ctx, otp); err != nil {
		return err
	}

	return s.mailer.SendOTP(email, rawOTP)
}

func (s *Service) VerifyOTP(ctx context.Context, rawOTP string) (string, error) {
	hash := HashOTP(rawOTP)

	otp, err := s.otps.FindByHash(ctx, hash)
	if err != nil {
		return "", err
	}

	if otp == nil {
		return "", errors.New("Invalid OTP")
	}

	// 2️⃣ validate
	if otp.Used {
		return "", errors.New("OTP already used")
	}

	if time.Now().After(otp.ExpiresAt) {
		return "", errors.New("OTP expired")
	}

	// 3️⃣ mark used
	if err := s.otps.MarkUsed(ctx, otp.ID); err != nil {
		return "", err
	}

	session := &session.Session{
		ID:        uuid.NewString(),
		UserID:    otp.UserID,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
	}

	if err := s.sessions.Create(ctx, session); err != nil {
		return "", err
	}

	return session.ID, nil
}
