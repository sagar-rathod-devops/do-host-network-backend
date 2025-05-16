package repositories

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
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
                        id, user_id, profile_image, full_name, designation, organization,
                        professional_summary, location, email, contact_number, created_at, updated_at
                ) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)
        `
	_, err := r.DB.Exec(query,
		profile.ID,
		profile.UserID,
		profile.ProfileImage,
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
	query := `SELECT id, user_id, profile_image, full_name, designation, organization,
                          professional_summary, location, email, contact_number,
                          created_at, updated_at FROM user_profile WHERE user_id = $1`

	row := r.DB.QueryRow(query, userID)

	var profile models.UserProfile
	err := row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.ProfileImage,
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

func (r *UserProfileRepository) GetAll() ([]*models.UserProfile, error) {
	query := `SELECT id, user_id, profile_image, full_name, designation, organization,
                     professional_summary, location, email, contact_number,
                     created_at, updated_at FROM user_profile`

	rows, err := r.DB.Query(query)
	if err != nil {
		log.Printf("DB query error: %v", err)
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.UserProfile

	for rows.Next() {
		var (
			idStr     string
			userIDStr string

			profileImage        sql.NullString
			fullName            sql.NullString
			designation         sql.NullString
			organization        sql.NullString
			professionalSummary sql.NullString
			location            sql.NullString
			email               sql.NullString
			contactNumber       sql.NullString

			createdAt time.Time
			updatedAt time.Time
		)

		err := rows.Scan(
			&idStr,
			&userIDStr,
			&profileImage,
			&fullName,
			&designation,
			&organization,
			&professionalSummary,
			&location,
			&email,
			&contactNumber,
			&createdAt,
			&updatedAt,
		)
		if err != nil {
			log.Printf("Row scan error: %v", err)
			return nil, err
		}

		id, err := uuid.Parse(idStr)
		if err != nil {
			log.Printf("UUID parse error for ID: %v", err)
			return nil, err
		}

		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			log.Printf("UUID parse error for UserID: %v", err)
			return nil, err
		}

		profile := &models.UserProfile{
			ID:                  id,
			UserID:              userID,
			ProfileImage:        nullToString(profileImage),          // *string
			FullName:            derefString(nullToString(fullName)), // string
			Designation:         nullToString(designation),           // *string
			Organization:        nullToString(organization),          // *string
			ProfessionalSummary: nullToString(professionalSummary),   // *string
			Location:            nullToString(location),              // *string
			Email:               derefString(nullToString(email)),    // string
			ContactNumber:       nullToString(contactNumber),         // *string
			CreatedAt:           createdAt,
			UpdatedAt:           updatedAt,
		}

		profiles = append(profiles, profile)
	}

	if err := rows.Err(); err != nil {
		log.Printf("Rows iteration error: %v", err)
		return nil, err
	}

	return profiles, nil
}

// Helper function to convert sql.NullString to plain string
func nullToString(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

func derefString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func (r *UserProfileRepository) Update(userID string, updated *models.UserProfile) (*models.UserProfile, error) {
	query := `
		UPDATE user_profile SET
			profile_image = $1,
			full_name = $2,
			designation = $3,
			organization = $4,
			professional_summary = $5,
			location = $6,
			email = $7,
			contact_number = $8,
			updated_at = $9
		WHERE user_id = $10
		RETURNING id, user_id, profile_image, full_name, designation, organization,
				  professional_summary, location, email, contact_number, created_at, updated_at`

	row := r.DB.QueryRow(query,
		updated.ProfileImage,
		updated.FullName,
		updated.Designation,
		updated.Organization,
		updated.ProfessionalSummary,
		updated.Location,
		updated.Email,
		updated.ContactNumber,
		updated.UpdatedAt,
		userID,
	)

	var profile models.UserProfile
	err := row.Scan(
		&profile.ID,
		&profile.UserID,
		&profile.ProfileImage,
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

func (r *UserProfileRepository) Delete(userID string) error {
	query := `DELETE FROM user_profile WHERE user_id = $1`
	_, err := r.DB.Exec(query, userID)
	return err
}
