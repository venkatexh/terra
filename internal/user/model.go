package user

import "time"

type User struct {
	ID string
	Email string
	EmailVerified bool
	CreatedAt time.Time
}