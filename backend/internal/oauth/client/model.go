package client

import "time"

type Client struct {
	ID string
	Name string
	ClientID string
	ClientSecret string
	RedirectURIs []string
	CreatedAt time.Time
}