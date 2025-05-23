package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type UserEducationService struct {
	Repo *repositories.UserEducationRepository
}

func (s *UserEducationService) Create(ctx context.Context, edu *models.UserEducation) error {
	return s.Repo.Create(edu)
}

func (s *UserEducationService) GetByUserID(ctx context.Context, userID uuid.UUID) ([]*models.UserEducation, error) {
	return s.Repo.GetByUserID(userID)
}

func (s *UserEducationService) Update(ctx context.Context, edu *models.UserEducation) error {
	return s.Repo.Update(edu)
}

func (s *UserEducationService) Delete(ctx context.Context, eduID uuid.UUID) error {
	return s.Repo.Delete(eduID)
}
