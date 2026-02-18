package token

import "time"

type AccessToken struct {
	Token string
	UserID string
	ClientID string
	ExpiresAt time.Time
}