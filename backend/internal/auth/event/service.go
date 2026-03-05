package authEvent

import (
	"context"
	"encoding/json"
)

type AuthEventService struct {
	authEventRepo *AuthEventRepository
}

func NewAuthEventService(authEventRepo *AuthEventRepository) *AuthEventService {
	return &AuthEventService{authEventRepo: authEventRepo}
}

func (s *AuthEventService) RecordEvent(ctx context.Context, UserID string, ClientID *string, EventType, Status, IPAddress, UserAgent string, Metadata map[string]interface{}) error {

	metaJSON, err := json.Marshal(Metadata)
	if err != nil {
		return err
	}

	event := &AuthEvent{
		UserID:    UserID,
		ClientID:  ClientID,
		EventType: EventType,
		Status:    Status,
		IPAddress: IPAddress,
		UserAgent: UserAgent,
		Metadata:  string(metaJSON),
	}

	err = s.authEventRepo.Create(ctx, event)

	if err != nil {
		return err
	}

	return nil
}
