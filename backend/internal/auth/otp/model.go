package otp

import (
	"time"
)
type OTP struct {
	ID        string `json:"id"`
	UserID    string `json:"user_id"`
	Hash      string    `json:"hash"`
	ExpiresAt time.Time `json:"expires_at"`
	Used      bool      `json:"used"`
}