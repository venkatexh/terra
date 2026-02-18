package token

import "time"

type LoginToken struct {
	ID string
	UserID string
	Hash string
	ExpiresAt time.Time
	Used bool
}