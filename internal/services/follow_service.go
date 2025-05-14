package services

import (
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type FollowService struct {
	FollowRepository *repositories.FollowRepository
}

// FollowUser allows a user to follow another user
func (service *FollowService) FollowUser(followerID, followedID uuid.UUID) error {
	return service.FollowRepository.FollowUser(followerID, followedID)
}

// UnfollowUser allows a user to unfollow another user
func (service *FollowService) UnfollowUser(followerID, followedID uuid.UUID) error {
	return service.FollowRepository.UnfollowUser(followerID, followedID)
}

// GetFollowers retrieves a list of followers for a user
func (service *FollowService) GetFollowers(userID uuid.UUID) ([]uuid.UUID, error) {
	return service.FollowRepository.GetFollowers(userID)
}

// GetFollowings retrieves a list of users that a user is following
func (service *FollowService) GetFollowings(userID uuid.UUID) ([]uuid.UUID, error) {
	return service.FollowRepository.GetFollowings(userID)
}
