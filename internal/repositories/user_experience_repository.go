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

func (r *UserExperienceRepository) Update(ctx context.Context, exp *models.UserExperience) error {
	query := `
		UPDATE user_experience SET
			job_title = $1,
			company_name = $2,
			location = $3,
			job_description = $4,
			achievements = $5,
			start_date = $6,
			end_date = $7,
			updated_at = NOW()
		WHERE id = $8
	`
	_, err := r.DB.ExecContext(ctx, query,
		exp.JobTitle, exp.CompanyName, exp.Location,
		exp.JobDescription, exp.Achievements,
		exp.StartDate, exp.EndDate,
		exp.ID,
	)
	return err
}

func (r *UserExperienceRepository) Delete(ctx context.Context, id uuid.UUID) error {
	query := `DELETE FROM user_experience WHERE id = $1`
	_, err := r.DB.ExecContext(ctx, query, id)
	return err
}
