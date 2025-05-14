package repositories

import (
	"database/sql"
	"log"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type PostCommentRepository struct {
	DB *sql.DB
}

func (repo *PostCommentRepository) CreateComment(userID, postID uuid.UUID, comment string) error {
	_, err := repo.DB.Exec(`
        INSERT INTO post_comments (user_id, post_id, comment) 
        VALUES ($1, $2, $3)`, userID, postID, comment)
	if err != nil {
		log.Printf("Error inserting comment for user %v on post %v: %v", userID, postID, err)
		return err
	}
	return nil
}

func (repo *PostCommentRepository) GetComments(postID uuid.UUID) ([]models.PostComment, error) {
	rows, err := repo.DB.Query(`
		SELECT id, user_id, post_id, comment, created_at
		FROM post_comments 
		WHERE post_id = $1`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var comments []models.PostComment
	for rows.Next() {
		var comment models.PostComment
		if err := rows.Scan(&comment.ID, &comment.UserID, &comment.PostID, &comment.Comment, &comment.CreatedAt); err != nil {
			return nil, err
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repo *PostCommentRepository) PostExists(postID uuid.UUID) (bool, error) {
	var exists bool
	err := repo.DB.QueryRow("SELECT EXISTS (SELECT 1 FROM content_post WHERE id = $1)", postID).Scan(&exists)
	return exists, err
}
