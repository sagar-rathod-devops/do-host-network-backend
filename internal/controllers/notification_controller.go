package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type NotificationController struct {
	NotificationService *services.NotificationService
}

func NewNotificationController(service *services.NotificationService) *NotificationController {
	return &NotificationController{NotificationService: service}
}

func (c *NotificationController) CreateNotification(ctx *gin.Context) {
	var notif models.Notification

	if err := ctx.ShouldBindJSON(&notif); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if err := c.NotificationService.CreateNotification(ctx, &notif); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create notification"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Notification created successfully"})
}

func (c *NotificationController) GetNotifications(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID"})
		return
	}

	notifications, err := c.NotificationService.GetNotificationsForUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch notifications"})
		return
	}

	if len(notifications) == 0 {
		ctx.JSON(http.StatusOK, gin.H{"message": "No notifications found"})
		return
	}

	ctx.JSON(http.StatusOK, notifications)
}
