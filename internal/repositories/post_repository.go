package repositories

import (
	"context"
	"database/sql"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(db *sql.DB) *PostRepository {
	return &PostRepository{DB: db}
}

func (r *PostRepository) CreatePost(ctx context.Context, post *models.ContentPost) (*models.ContentPost, error) {
	query := `
        INSERT INTO content_post (user_id, post_content, media_url)
        VALUES ($1, $2, $3)
        RETURNING id, user_id, post_content, media_url, created_at
    `

	row := r.DB.QueryRowContext(ctx, query, post.UserID, post.PostContent, post.MediaURL)
	var created models.ContentPost
	err := row.Scan(&created.ID, &created.UserID, &created.PostContent, &created.MediaURL, &created.CreatedAt)
	if err != nil {
		return nil, err
	}

	return &created, nil
}

func (r *PostRepository) GetAll(ctx context.Context) ([]models.ContentPost, error) {
	query := `SELECT id, user_id, post_content, media_url, created_at FROM content_post ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.ContentPost
	for rows.Next() {
		var post models.ContentPost
		err := rows.Scan(&post.ID, &post.UserID, &post.PostContent, &post.MediaURL, &post.CreatedAt)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}

	return posts, nil
}

func (r *PostRepository) GetPostsByUserID(userID uuid.UUID) ([]models.ContentPost, error) {
	rows, err := r.DB.Query(`
		SELECT 
			cp.id, cp.user_id, cp.post_content, cp.media_url, cp.created_at,
			u.id, u.username, u.email
			FROM content_post cp
			INNER JOIN users u ON cp.user_id = u.id
			WHERE cp.user_id = $1
		ORDER BY cp.created_at DESC
	`, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var posts []models.ContentPost
	for rows.Next() {
		var post models.ContentPost
		var user models.User

		err := rows.Scan(
			&post.ID, &post.UserID, &post.PostContent, &post.MediaURL, &post.CreatedAt,
			&user.ID, &user.Username, &user.Email,
		)
		if err != nil {
			return nil, err
		}

		post.User = &user
		posts = append(posts, post)
	}
	return posts, nil
}
