package authcode

import "time"

type AuthorizationCode struct {
	Code string
	ClientID string
	UserID string
	RedirectURI string
	ExpiresAt time.Time
}