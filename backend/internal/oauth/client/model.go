package client

import "time"

type Client struct {
	ID string `json:"id"`
	Name string `json:"name"`
	ClientID string `json:"clientId"`
	ClientSecret string `json:"clientSecret"`
	RedirectURIs []string `json:"redirectUris"`
	ProjectID string `json:"projectId"`
	CreatedAt time.Time `json:"createdAt"`
}