package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type VideoProfileController struct {
	VideoProfileService *services.VideoProfileService
}

// POST /api/video
func (vc *VideoProfileController) CreateVideoProfile(ctx *gin.Context) {
	var input struct {
		UserID   string `json:"user_id" binding:"required"`
		VideoURL string `json:"video_url" binding:"required"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	video := &models.VideoProfile{
		UserID:   userID,
		VideoURL: input.VideoURL,
	}

	if err := vc.VideoProfileService.Create(context.Background(), video); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create video profile"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Video profile created successfully"})
}

// GET /api/video/:user_id
func (vc *VideoProfileController) GetVideoProfilesByUser(ctx *gin.Context) {
	userIDParam := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	profiles, err := vc.VideoProfileService.GetByUserID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, profiles)
}
