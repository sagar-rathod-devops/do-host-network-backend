package models

import (
	"time"

	"github.com/google/uuid"
)

type VideoProfile struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	VideoURL  string    `json:"video_url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
