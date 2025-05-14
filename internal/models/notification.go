package models

import (
	"time"

	"github.com/google/uuid"
)

type Notification struct {
	ID              uuid.UUID `json:"id"`
	RecipientUserID uuid.UUID `json:"recipient_user_id"` // corrected tag name
	SenderUserID    uuid.UUID `json:"sender_user_id"`
	Type            string    `json:"type"`
	EntityID        uuid.UUID `json:"entity_id"`
	EntityType      string    `json:"entity_type"`
	Message         string    `json:"message"`
	IsRead          bool      `json:"is_read"`
	CreatedAt       time.Time `json:"created_at"`
}
