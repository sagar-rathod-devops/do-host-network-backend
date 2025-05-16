package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID                  uuid.UUID
	UserID              uuid.UUID
	ProfileImage        *string // nullable
	FullName            string  // required
	Designation         *string // nullable
	Organization        *string // nullable
	ProfessionalSummary *string // nullable
	Location            *string // nullable
	Email               string  // required
	ContactNumber       *string // nullable
	CreatedAt           time.Time
	UpdatedAt           time.Time
}
