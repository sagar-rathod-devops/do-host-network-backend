package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type PostLikeController struct {
	PostLikeService *services.PostLikeService
}

// LikePost handles POST request for liking a post
func (controller *PostLikeController) LikePost(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	userID := user.ID
	postIDStr := ctx.Param("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = controller.PostLikeService.LikePost(uuidUserID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post liked successfully"})
}

// UnlikePost handles POST request for unliking a post
func (controller *PostLikeController) UnlikePost(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)
	userID := user.ID
	postIDStr := ctx.Param("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	uuidUserID, err := uuid.Parse(userID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	err = controller.PostLikeService.UnlikePost(uuidUserID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Post unliked successfully"})
}

// GetPostLikes handles GET request to retrieve likes for a post
func (controller *PostLikeController) GetPostLikes(ctx *gin.Context) {
	postIDStr := ctx.Param("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	likes, err := controller.PostLikeService.GetPostLikes(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
