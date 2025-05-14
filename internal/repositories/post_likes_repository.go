package repositories

import (
	"database/sql"

	"github.com/google/uuid"
)

type PostLikeRepository struct {
	DB *sql.DB
}

// CreateLike adds a new like for a post by a user
func (repo *PostLikeRepository) CreateLike(userID, postID uuid.UUID) error {
	_, err := repo.DB.Exec(`
		INSERT INTO post_likes (user_id, post_id) 
		VALUES ($1, $2) 
		ON CONFLICT (user_id, post_id) DO NOTHING`, userID, postID)
	return err
}

// RemoveLike removes a like from a post by a user
func (repo *PostLikeRepository) RemoveLike(userID, postID uuid.UUID) error {
	_, err := repo.DB.Exec(`
		DELETE FROM post_likes 
		WHERE user_id = $1 AND post_id = $2`, userID, postID)
	return err
}

// GetLikes retrieves all likes for a post
func (repo *PostLikeRepository) GetLikes(postID uuid.UUID) ([]uuid.UUID, error) {
	rows, err := repo.DB.Query(`
		SELECT user_id 
		FROM post_likes 
		WHERE post_id = $1`, postID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var likes []uuid.UUID
	for rows.Next() {
		var userID uuid.UUID
		if err := rows.Scan(&userID); err != nil {
			return nil, err
		}
		likes = append(likes, userID)
	}
	return likes, nil
}
