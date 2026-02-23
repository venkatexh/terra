package client

import "time"

type Client struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	ClientID     string    `json:"clientId"`
	ClientSecret string    `json:"clientSecret"`
	ProjectID    string    `json:"projectId"`
	CreatedAt    time.Time `json:"createdAt"`
}

type ClientResponse struct {
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	ClientID     string        `json:"clientId"`
	ClientSecret string        `json:"clientSecret"`
	RedirectURIs []RedirectURI `json:"redirectUris"`
	CreatedAt    time.Time     `json:"createdAt"`
	ProjectID    string        `json:"projectId"`
}

type ClientRequest struct {
	Name         string   `json:"name"`
	RedirectURIs []string `json:"redirectUris"`
}

type RedirectURI struct {
	ID         string    `json:"id"`
	ClientDBID string    `json:"clientId"`
	URI        string    `json:"uri"`
	CreatedAt  time.Time `json:"createdAt"`
}
