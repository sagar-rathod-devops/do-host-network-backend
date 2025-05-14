package repositories

import (
	"database/sql"

	"github.com/google/uuid"
)

type FollowRepository struct {
	DB *sql.DB
}

// FollowUser allows a user to follow another user
func (repo *FollowRepository) FollowUser(followerID, followedID uuid.UUID) error {
	_, err := repo.DB.Exec(`
		INSERT INTO followers (follower_id, followed_id) 
		VALUES ($1, $2) 
		ON CONFLICT (follower_id, followed_id) DO NOTHING`, followerID, followedID)

	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(`
		INSERT INTO followings (follower_id, followed_id) 
		VALUES ($1, $2) 
		ON CONFLICT (follower_id, followed_id) DO NOTHING`, followerID, followedID)

	return err
}

// UnfollowUser allows a user to unfollow another user
func (repo *FollowRepository) UnfollowUser(followerID, followedID uuid.UUID) error {
	_, err := repo.DB.Exec(`
		DELETE FROM followers 
		WHERE follower_id = $1 AND followed_id = $2`, followerID, followedID)

	if err != nil {
		return err
	}

	_, err = repo.DB.Exec(`
		DELETE FROM followings 
		WHERE follower_id = $1 AND followed_id = $2`, followerID, followedID)

	return err
}

// GetFollowers retrieves a list of followers for a user
func (repo *FollowRepository) GetFollowers(userID uuid.UUID) ([]uuid.UUID, error) {
	rows, err := repo.DB.Query(`
		SELECT follower_id 
		FROM followers 
		WHERE followed_id = $1`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followers []uuid.UUID
	for rows.Next() {
		var followerID uuid.UUID
		if err := rows.Scan(&followerID); err != nil {
			return nil, err
		}
		followers = append(followers, followerID)
	}
	return followers, nil
}

// GetFollowings retrieves a list of users that a user is following
func (repo *FollowRepository) GetFollowings(userID uuid.UUID) ([]uuid.UUID, error) {
	rows, err := repo.DB.Query(`
		SELECT followed_id 
		FROM followings 
		WHERE follower_id = $1`, userID)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var followings []uuid.UUID
	for rows.Next() {
		var followedID uuid.UUID
		if err := rows.Scan(&followedID); err != nil {
			return nil, err
		}
		followings = append(followings, followedID)
	}
	return followings, nil
}
