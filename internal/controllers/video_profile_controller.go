package controllers

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/config"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
	"github.com/sagar-rathod-devops/do-host-network-backend/utils"
)

type VideoProfileController struct {
	VideoProfileService *services.VideoProfileService
}

// POST /api/video
// POST /api/video/upload
func (vc *VideoProfileController) UploadVideo(ctx *gin.Context) {
	userIDStr := ctx.PostForm("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	fileHeader, err := ctx.FormFile("video")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get video file"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open video file"})
		return
	}
	defer file.Close()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
		return
	}

	uploader, err := utils.NewS3Uploader(cfg)
	if err != nil || uploader == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create S3 uploader"})
		return
	}

	key := fmt.Sprintf("videos/%s_%d_%s", userID, time.Now().Unix(), fileHeader.Filename)
	videoURL, err := uploader.UploadFile(file, fileHeader, key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload video"})
		return
	}

	video := &models.VideoProfile{
		UserID:   userID,
		VideoURL: videoURL,
	}

	if err := vc.VideoProfileService.Create(ctx.Request.Context(), video); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save video profile"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message":   "Video uploaded successfully",
		"video_url": videoURL,
	})
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

func (vc *VideoProfileController) StreamVideo(ctx *gin.Context) {
	videoURL := ctx.Query("url")
	if videoURL == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Video URL required"})
		return
	}

	resp, err := http.Get(videoURL)
	if err != nil || resp.StatusCode != 200 {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch video"})
		return
	}
	defer resp.Body.Close()

	ctx.Header("Content-Type", resp.Header.Get("Content-Type"))
	ctx.Status(resp.StatusCode)
	io.Copy(ctx.Writer, resp.Body)
}

func (vc *VideoProfileController) UpdateVideo(ctx *gin.Context) {
	videoIDStr := ctx.Param("id")
	videoID, err := uuid.Parse(videoIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	fileHeader, err := ctx.FormFile("video")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to get new video file"})
		return
	}

	file, err := fileHeader.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open video file"})
		return
	}
	defer file.Close()

	cfg, err := config.LoadConfig(".")
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
		return
	}

	uploader, err := utils.NewS3Uploader(cfg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create uploader"})
		return
	}

	key := fmt.Sprintf("videos/%s_%d_%s", videoID, time.Now().Unix(), fileHeader.Filename)
	videoURL, err := uploader.UploadFile(file, fileHeader, key)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload new video"})
		return
	}

	video := &models.VideoProfile{
		ID:       videoID,
		VideoURL: videoURL,
	}

	if err := vc.VideoProfileService.Update(ctx.Request.Context(), video); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update video"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message":   "Video updated successfully",
		"video_url": videoURL,
	})
}

func (vc *VideoProfileController) DeleteVideo(ctx *gin.Context) {
	videoIDStr := ctx.Param("id")
	videoID, err := uuid.Parse(videoIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid video ID"})
		return
	}

	if err := vc.VideoProfileService.Delete(ctx.Request.Context(), videoID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete video"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Video deleted successfully"})
}
