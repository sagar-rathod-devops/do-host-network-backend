package models

import (
	"time"

	"github.com/google/uuid"
)

type UserProfile struct {
	ID                  uuid.UUID `json:"id"`
	UserID              uuid.UUID `json:"user_id"`
	FullName            string    `json:"full_name"`
	Designation         string    `json:"designation"`
	Organization        string    `json:"organization"`
	ProfessionalSummary string    `json:"professional_summary"`
	Location            string    `json:"location"`
	Email               string    `json:"email"`
	ContactNumber       string    `json:"contact_number"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
}
