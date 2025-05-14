package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type VideoProfileService struct {
	Repo *repositories.VideoProfileRepository
}

func (s *VideoProfileService) Create(ctx context.Context, video *models.VideoProfile) error {
	return s.Repo.Create(video)
}

func (s *VideoProfileService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*models.VideoProfile, error) {
	return s.Repo.GetByUserID(userID)
}
