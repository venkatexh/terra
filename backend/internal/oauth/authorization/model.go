package authorization

import "time"

type Authorization struct {
	ID string
	GroupID string
	ClientDBID string
	Scopes []string
	CreatedAt time.Time
}

type AuthorizationDTO struct {
	ID        string `json:"id"`
	GroupID    string `json:"userId"`
	ClientDBID  string `json:"clientId"`
	ClientName string `json:"clientName"`
	CreatedAt time.Time `json:"createdAt"`
}

type AuthorizationRequest struct {
	GroupID string `json:"groupId"`
	ClientDBID string `json:"clientId"`
	// Scopes []string `json:"scopes"`
}
