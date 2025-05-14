package repositories

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
)

type NotificationRepository struct {
	DB *sql.DB
}

func NewNotificationRepository(db *sql.DB) *NotificationRepository {
	return &NotificationRepository{DB: db}
}

func (r *NotificationRepository) Create(ctx context.Context, n *models.Notification) error {
	n.CreatedAt = time.Now()

	if n.SenderUserID == uuid.Nil || n.RecipientUserID == uuid.Nil {
		return fmt.Errorf("sender_user_id and recipient_user_id cannot be empty")
	}

	query := `INSERT INTO notifications (id, recipient_user_id, sender_user_id, type, entity_id, entity_type, message, is_read, created_at)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := r.DB.ExecContext(ctx, query,
		n.ID, n.RecipientUserID, n.SenderUserID, n.Type, n.EntityID, n.EntityType, n.Message, n.IsRead, n.CreatedAt)

	if err != nil {
		log.Printf("Failed to insert notification: %v", err)
		return fmt.Errorf("failed to create notification: %w", err)
	}

	return nil
}

func (r *NotificationRepository) GetByUserID(ctx context.Context, userID uuid.UUID) ([]models.Notification, error) {
	query := `SELECT id, recipient_user_id, sender_user_id, type, entity_id, entity_type, message, is_read, created_at
			  FROM notifications WHERE recipient_user_id = $1 ORDER BY created_at DESC`

	rows, err := r.DB.QueryContext(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var n models.Notification
		if err := rows.Scan(
			&n.ID, &n.RecipientUserID, &n.SenderUserID,
			&n.Type, &n.EntityID, &n.EntityType,
			&n.Message, &n.IsRead, &n.CreatedAt,
		); err != nil {
			return nil, err
		}
		notifications = append(notifications, n)
	}
	return notifications, nil
}
