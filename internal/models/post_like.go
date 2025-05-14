package models

import (
	"time"

	"github.com/google/uuid"
)

type PostLike struct {
	ID        uuid.UUID `json:"id"`
	UserID    uuid.UUID `json:"user_id"`
	PostID    uuid.UUID `json:"post_id"`
	CreatedAt time.Time `json:"created_at"`
}
