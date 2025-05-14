package models

import (
	"time"

	"github.com/google/uuid"
)

type Follower struct {
	FollowerID uuid.UUID `json:"follower_id"`
	FollowedID uuid.UUID `json:"followed_id"`
	CreatedAt  time.Time `json:"created_at"`
}
