package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type NotificationService struct {
	NotificationRepository *repositories.NotificationRepository
}

func NewNotificationService(repo *repositories.NotificationRepository) *NotificationService {
	return &NotificationService{NotificationRepository: repo}
}

func (s *NotificationService) CreateNotification(ctx context.Context, n *models.Notification) error {
	if n.ID == uuid.Nil {
		n.ID = uuid.New()
	}
	return s.NotificationRepository.Create(ctx, n)
}

func (s *NotificationService) GetNotificationsForUser(ctx context.Context, userID uuid.UUID) ([]models.Notification, error) {
	return s.NotificationRepository.GetByUserID(ctx, userID)
}
