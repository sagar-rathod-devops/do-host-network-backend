package models

import (
	"time"

	"github.com/google/uuid"
)

type UserExperience struct {
	ID             uuid.UUID `json:"id"`
	UserID         uuid.UUID `json:"user_id"`
	JobTitle       string    `json:"job_title"`
	CompanyName    string    `json:"company_name"`
	Location       string    `json:"location"`
	JobDescription string    `json:"job_description"`
	Achievements   string    `json:"achievements"`
	StartDate      string    `json:"start_date"`
	EndDate        string    `json:"end_date"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
