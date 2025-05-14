package controllers

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{PostService: service}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	var input struct {
		UserID      string `json:"user_id" binding:"required"`
		PostContent string `json:"post_content" binding:"required"`
		MediaURL    string `json:"media_url"`
	}

	// Bind the request body to the input struct
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Convert the UserID to UUID
	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UserID format"})
		return
	}

	// Create the post object
	post := &models.ContentPost{
		UserID:      userID,
		PostContent: input.PostContent,
		MediaURL:    input.MediaURL,
	}

	// Call the service to create the post
	if _, err := pc.PostService.CreatePost(context.Background(), post); err != nil {
		log.Printf("CreatePost error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// Return only a success message
	ctx.JSON(http.StatusCreated, gin.H{"message": "Post created successfully"})
}

func (c *PostController) GetAllContentPosts(ctx *gin.Context) {
	posts, err := c.PostService.GetAllContentPosts(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func (c *PostController) GetPostsByUserID(ctx *gin.Context) {
	userIDParam := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDParam)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	posts, err := c.PostService.GetPostsByUserID(userID)
	if err != nil {
		if err.Error() == "no posts found for this user" {
			ctx.JSON(http.StatusNotFound, gin.H{"message": "No posts found for this user"})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user posts"})
		return
	}

	ctx.JSON(http.StatusOK, posts)
}
