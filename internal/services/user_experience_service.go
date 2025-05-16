package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type UserExperienceService struct {
	UserExperienceRepository *repositories.UserExperienceRepository
}

func NewUserExperienceService(repo *repositories.UserExperienceRepository) *UserExperienceService {
	return &UserExperienceService{UserExperienceRepository: repo}
}

func (s *UserExperienceService) Create(ctx context.Context, exp *models.UserExperience) error {
	// Call the repository to save the user experience
	return s.UserExperienceRepository.Create(ctx, exp)
}

func (s *UserExperienceService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]models.UserExperience, error) {
	// Fetch user experiences by user ID from the repository
	return s.UserExperienceRepository.GetByUserID(ctx, userID)
}

func (s *UserExperienceService) Update(ctx context.Context, exp *models.UserExperience) error {
	return s.UserExperienceRepository.Update(ctx, exp)
}

func (s *UserExperienceService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.UserExperienceRepository.Delete(ctx, id)
}
