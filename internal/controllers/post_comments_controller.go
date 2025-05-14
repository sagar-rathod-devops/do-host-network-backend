package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type PostCommentController struct {
	PostCommentService *services.PostCommentService
}

func (controller *PostCommentController) CommentOnPost(ctx *gin.Context) {
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	usr, ok := user.(models.User)
	if !ok {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user data"})
		return
	}

	userID, err := uuid.Parse(usr.ID)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid user ID"})
		return
	}

	postIDStr := ctx.Param("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	var input struct {
		Comment string `json:"comment" binding:"required,min=1,max=500"`
	}
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.PostCommentService.CommentOnPost(userID, postID, input.Comment)
	if err != nil {
		if err.Error() == "post not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func (controller *PostCommentController) GetPostComments(ctx *gin.Context) {
	postIDStr := ctx.Param("post_id")
	postID, err := uuid.Parse(postIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID"})
		return
	}

	exists, err := controller.PostCommentService.PostCommentRepository.PostExists(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error checking post existence"})
		return
	}
	if !exists {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Post not found"})
		return
	}

	comments, err := controller.PostCommentService.GetPostComments(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": comments})
}
