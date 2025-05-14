package models

import (
	"time"

	"github.com/google/uuid"
)

type Following struct {
	FollowedID uuid.UUID `json:"followed_id"`
	FollowerID uuid.UUID `json:"follower_id"`
	CreatedAt  time.Time `json:"created_at"`
}
