package repositories

import (
	"database/sql"

	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type UserProfileRepository struct {
	DB *sql.DB
}

func NewUserProfileRepository(db *sql.DB) *UserProfileRepository {
	return &UserProfileRepository{DB: db}
}

func (r *UserProfileRepository) Create(profile *models.UserProfile) error {
	query := `
		INSERT INTO user_profile (
			id, user_id, full_name, designation, organization, 
			professional_summary, location, email, contact_number, created_at, updated_at
		) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
	`
	_, err := r.DB.Exec(query,
		profile.ID,
		profile.UserID,
		profile.FullName,
		profile.Designation,
		profile.Organization,
		profile.ProfessionalSummary,
		profile.Location,
		profile.Email,
		profile.ContactNumber,
		profile.CreatedAt,
		profile.UpdatedAt,
	)
	if err != nil {
		println("DB Insert Error:", err.Error()) // ← helpful
	}
	return err
}

func (r *UserProfileRepository) GetByUserID(userID string) (*models.UserProfile, error) {
	query := `SELECT id, user_id, full_name, designation, organization,
			  professional_summary, location, email, contact_number,
			  created_at, updated_at FROM user_profile WHERE user_id = $1`

	row := r.DB.QueryRow(query, userID)

	var profile models.UserProfile
	err := row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.FullName,
		&profile.Designation,
		&profile.Organization,
		&profile.ProfessionalSummary,
		&profile.Location,
		&profile.Email,
		&profile.ContactNumber,
		&profile.CreatedAt,
		&profile.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &profile, nil
}
