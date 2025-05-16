package services

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type UserProfileService struct {
	Repo *repositories.UserProfileRepository
}

func NewUserProfileService(repo *repositories.UserProfileRepository) *UserProfileService {
	return &UserProfileService{Repo: repo}
}

func (s *UserProfileService) Create(ctx context.Context, profile *models.UserProfile) (*models.UserProfile, error) {
	profile.ID = uuid.New()
	profile.CreatedAt = time.Now()
	profile.UpdatedAt = time.Now()

	err := s.Repo.Create(profile)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

func (s *UserProfileService) GetByUserID(ctx context.Context, userID string) (*models.UserProfile, error) {
	return s.Repo.GetByUserID(userID)
}

func (s *UserProfileService) GetAll(ctx context.Context) ([]*models.UserProfile, error) {
	return s.Repo.GetAll()
}

func (s *UserProfileService) Update(ctx context.Context, userID string, updated *models.UserProfile) (*models.UserProfile, error) {
	updated.UpdatedAt = time.Now()
	return s.Repo.Update(userID, updated)
}

func (s *UserProfileService) Delete(ctx context.Context, userID string) error {
	return s.Repo.Delete(userID)
}
