package authEvent

import "time"

type AuthEvent struct {
	ID         string    `json:"id"`
	UserID     string    `json:"userId"`
	ClientID   *string   `json:"clientId"`
	EventType  string    `json:"eventType"`
	Status     string    `json:"status"`
	IPAddress  string    `json:"ipAddress"`
	UserAgent  string    `json:"userAgent"`
	Metadata   string    `json:"metadata"`
	OccurredAt time.Time `json:"occurredAt"`
}
