package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/config"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
	"github.com/sagar-rathod-devops/do-host-network-backend/utils"
)

type PostController struct {
	PostService *services.PostService
}

func NewPostController(service *services.PostService) *PostController {
	return &PostController{
		PostService: service,
	}
}

func (pc *PostController) CreatePost(ctx *gin.Context) {
	fmt.Println("🚀 Received request to create post")

	// 1. Parse form-data fields
	userIDStr := ctx.PostForm("user_id")
	postContent := ctx.PostForm("post_content")

	if strings.TrimSpace(userIDStr) == "" || strings.TrimSpace(postContent) == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id and post_content are required"})
		return
	}

	// 2. Parse user ID to UUID
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id format"})
		return
	}

	// 3. Get uploaded media file (optional)
	fileHeader, err := ctx.FormFile("media_url")
	mediaURL := "" // string, default empty

	if err == nil && fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open media file"})
			return
		}
		defer file.Close()

		fmt.Println("📁 Media file received:", fileHeader.Filename)

		// 4. Load config & initialize uploader
		cfg, err := config.LoadConfig(".")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
			return
		}

		uploader, err := utils.NewS3Uploader(cfg)
		if err != nil || uploader == nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize S3 uploader"})
			return
		}

		// 5. Generate S3 key and upload
		key := fmt.Sprintf("post-media/%s_%d_%s", userID, time.Now().Unix(), fileHeader.Filename)
		url, err := uploader.UploadFile(file, fileHeader, key)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload media to S3"})
			return
		}

		fmt.Println("✅ Media uploaded successfully:", url)
		mediaURL = url // assign string directly
	} else {
		fmt.Println("⚠️ No media uploaded or error reading media:", err)
	}

	// 6. Create post model
	post := &models.ContentPost{
		UserID:      userID,
		PostContent: strings.TrimSpace(postContent),
		MediaURL:    mediaURL, // string, no pointer
	}

	// 7. Save post to DB
	createdPost, err := pc.PostService.CreatePost(context.Background(), post)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create post"})
		return
	}

	// 8. Prepare response
	if createdPost.MediaURL != "" {
		ctx.JSON(http.StatusCreated, gin.H{
			"message":      "Post created successfully",
			"post_content": createdPost.PostContent,
			"media_url":    createdPost.MediaURL,
		})
	} else {
		ctx.JSON(http.StatusCreated, gin.H{
			"message":      "Post created successfully",
			"post_content": createdPost.PostContent,
		})
	}
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
