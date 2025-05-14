package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type FollowController struct {
	FollowService *services.FollowService
}

// FollowUser handles the request for a user to follow another user
func (controller *FollowController) FollowUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	followerID, err := uuid.Parse(user.ID) // Convert string to uuid.UUID
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followedIDStr := ctx.Param("followed_id")
	followedID, err := uuid.Parse(followedIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followed user ID"})
		return
	}

	err = controller.FollowService.FollowUser(followerID, followedID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User followed successfully"})
}

func (controller *FollowController) UnfollowUser(ctx *gin.Context) {
	user := ctx.MustGet("user").(models.User)

	followerID, err := uuid.Parse(user.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followedIDStr := ctx.Param("followed_id")
	followedID, err := uuid.Parse(followedIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid followed user ID"})
		return
	}

	err = controller.FollowService.UnfollowUser(followerID, followedID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User unfollowed successfully"})
}

// GetFollowers handles the request to get all followers of a user
func (controller *FollowController) GetFollowers(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followers, err := controller.FollowService.GetFollowers(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"followers": followers})
}

// GetFollowings handles the request to get all users a user is following
func (controller *FollowController) GetFollowings(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	followings, err := controller.FollowService.GetFollowings(userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"followings": followings})
}
