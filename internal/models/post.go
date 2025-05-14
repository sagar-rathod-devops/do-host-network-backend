package models

import (
	"time"

	"github.com/google/uuid"
)

type ContentPost struct {
	ID          uuid.UUID `json:"id"`
	UserID      uuid.UUID `json:"user_id"`
	PostContent string    `json:"post_content"`
	MediaURL    string    `json:"media_url"`
	CreatedAt   time.Time `json:"created_at"`
	User        *User     `json:"user,omitempty"`
}

type JobPost struct {
	ID              uuid.UUID `json:"id"`
	UserID          uuid.UUID `json:"user_id"`
	JobTitle        string    `json:"job_title"`
	CompanyName     string    `json:"company_name"`
	JobDescription  string    `json:"job_description"`
	JobApplyURL     string    `json:"job_apply_url"`
	Location        string    `json:"location"`
	PostDate        time.Time `json:"post_date"`
	LastDateToApply time.Time `json:"last_date_to_apply"`
	CreatedAt       time.Time `json:"created_at"`
}
