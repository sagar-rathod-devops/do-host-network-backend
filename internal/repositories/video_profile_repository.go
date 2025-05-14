package repositories

import (
	"database/sql"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type VideoProfileRepository struct {
	DB *sql.DB
}

func (r *VideoProfileRepository) Create(video *models.VideoProfile) error {
	query := `
        INSERT INTO video_profile (id, user_id, video_url)
        VALUES ($1, $2, $3)
        RETURNING created_at, updated_at
    `
	video.ID = uuid.New()
	return r.DB.QueryRow(query, video.ID, video.UserID, video.VideoURL).
		Scan(&video.CreatedAt, &video.UpdatedAt)
}

func (r *VideoProfileRepository) GetByUserID(userID uuid.UUID) ([]*models.VideoProfile, error) {
	query := `SELECT id, user_id, video_url, created_at, updated_at FROM video_profile WHERE user_id = $1`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var profiles []*models.VideoProfile
	for rows.Next() {
		var v models.VideoProfile
		err := rows.Scan(&v.ID, &v.UserID, &v.VideoURL, &v.CreatedAt, &v.UpdatedAt)
		if err != nil {
			return nil, err
		}
		profiles = append(profiles, &v)
	}
	return profiles, nil
}
