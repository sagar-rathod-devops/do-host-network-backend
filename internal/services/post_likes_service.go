package services

import (
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type PostLikeService struct {
	PostLikeRepository *repositories.PostLikeRepository
}

// LikePost allows a user to like a post
func (service *PostLikeService) LikePost(userID, postID uuid.UUID) error {
	return service.PostLikeRepository.CreateLike(userID, postID)
}

// UnlikePost allows a user to remove a like from a post
func (service *PostLikeService) UnlikePost(userID, postID uuid.UUID) error {
	return service.PostLikeRepository.RemoveLike(userID, postID)
}

// GetPostLikes retrieves all likes for a specific post
func (service *PostLikeService) GetPostLikes(postID uuid.UUID) ([]uuid.UUID, error) {
	return service.PostLikeRepository.GetLikes(postID)
}
