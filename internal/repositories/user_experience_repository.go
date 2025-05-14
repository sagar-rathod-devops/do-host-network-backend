package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type UserExperienceRepository struct {
	DB *sql.DB
}

func NewUserExperienceRepository(db *sql.DB) *UserExperienceRepository {
	return &UserExperienceRepository{DB: db}
}

func (r *UserExperienceRepository) Create(ctx context.Context, exp *models.UserExperience) error {
	query := `
		INSERT INTO user_experience (
			id, user_id, job_title, company_name, location,
			job_description, achievements, start_date, end_date
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9
		) RETURNING created_at, updated_at
	`

	exp.ID = uuid.New()

	return r.DB.QueryRowContext(ctx, query,
		exp.ID, exp.UserID, exp.JobTitle, exp.CompanyName, exp.Location,
		exp.JobDescription, exp.Achievements, exp.StartDate, exp.EndDate,
	).Scan(&exp.CreatedAt, &exp.UpdatedAt)
}

func (r *UserExperienceRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]models.UserExperience, error) {
	query := `
		SELECT id, user_id, job_title, company_name, location,
			   job_description, achievements, start_date, end_date, created_at, updated_at
		FROM user_experience
		WHERE user_id = $1
	`

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var experiences []models.UserExperience
	for rows.Next() {
		var exp models.UserExperience
		err := rows.Scan(
			&exp.ID, &exp.UserID, &exp.JobTitle, &exp.CompanyName, &exp.Location,
			&exp.JobDescription, &exp.Achievements, &exp.StartDate, &exp.EndDate,
			&exp.CreatedAt, &exp.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		experiences = append(experiences, exp)
	}

	return experiences, nil
}
