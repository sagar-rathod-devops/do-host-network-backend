package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type JobRepository struct {
	DB *sql.DB
}

func (r *JobRepository) CreateJobPost(post *models.JobPost) error {
	query := `
		INSERT INTO job_post (
			id, user_id, job_title, company_name, job_description,
			job_apply_url, location, post_date, last_date_to_apply, created_at
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10
		)
	`

	post.ID = uuid.New()

	_, err := r.DB.Exec(
		query,
		post.ID, post.UserID, post.JobTitle, post.CompanyName, post.JobDescription,
		post.JobApplyURL, post.Location, post.PostDate, post.LastDateToApply, post.CreatedAt,
	)

	return err
}

func (r *JobRepository) GetAll(ctx context.Context) ([]models.JobPost, error) {
	query := `
		SELECT id, user_id, job_title, company_name, job_description,
		       job_apply_url, location, post_date, last_date_to_apply, created_at
		FROM job_post
		ORDER BY created_at DESC
	`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var jobPosts []models.JobPost
	for rows.Next() {
		var jp models.JobPost
		err := rows.Scan(
			&jp.ID, &jp.UserID, &jp.JobTitle, &jp.CompanyName, &jp.JobDescription,
			&jp.JobApplyURL, &jp.Location, &jp.PostDate, &jp.LastDateToApply, &jp.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		jobPosts = append(jobPosts, jp)
	}

	return jobPosts, nil
}
