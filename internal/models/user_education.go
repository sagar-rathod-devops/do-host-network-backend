package models

import (
	"time"

	"github.com/google/uuid"
)

type UserEducation struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	Degree          string    `json:"degree"`
	InstitutionName string    `json:"institution_name"`
	FieldOfStudy    string    `json:"field_of_study"`
	Grade           string    `json:"grade"`
	Year            string    `json:"year"` // Keep as string
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
