package controllers

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/config"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
	"github.com/sagar-rathod-devops/do-host-network-backend/utils"
)

type UserProfileController struct {
	UserProfileService *services.UserProfileService
}

func NewUserProfileController(profileService *services.UserProfileService) *UserProfileController {
	return &UserProfileController{
		UserProfileService: profileService,
	}
}

func (ctrl *UserProfileController) Create(ctx *gin.Context) {
	fmt.Println("🚀 Received request to create user profile")

	// 1. Parse form-data fields
	var input struct {
		UserID              string `form:"user_id" binding:"required"`
		FullName            string `form:"full_name" binding:"required"`
		Designation         string `form:"designation"`
		Organization        string `form:"organization"`
		ProfessionalSummary string `form:"professional_summary"`
		Location            string `form:"location"`
		Email               string `form:"email" binding:"required"`
		ContactNumber       string `form:"contact_number"`
	}

	if err := ctx.ShouldBind(&input); err != nil {
		fmt.Println("❌ Failed to bind form fields:", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	fmt.Println("✅ Form fields parsed successfully")

	// 2. Parse UUID
	uid, err := uuid.Parse(input.UserID)
	if err != nil {
		fmt.Println("❌ Invalid UUID format:", err.Error())
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}
	fmt.Println("✅ UserID UUID parsed:", uid.String())

	// 3. Get file
	fileHeader, err := ctx.FormFile("profile_image")
	var profileImageURL *string

	if err == nil && fileHeader != nil {
		file, err := fileHeader.Open()
		if err != nil {
			fmt.Println("❌ Failed to open profile image file:", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to open profile image"})
			return
		}
		defer file.Close()

		fmt.Println("📁 File received:", fileHeader.Filename)

		// 4. Load config & uploader
		cfg, err := config.LoadConfig(".")
		if err != nil {
			fmt.Println("❌ Failed to load config:", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to load config"})
			return
		}

		uploader, err := utils.NewS3Uploader(cfg)
		if err != nil || uploader == nil {
			fmt.Println("❌ Failed to create S3 uploader:", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create S3 uploader"})
			return
		}

		// 5. Upload to S3
		key := fmt.Sprintf("profile-images/%s_%d_%s", uid, time.Now().Unix(), fileHeader.Filename)
		url, err := uploader.UploadFile(file, fileHeader, key)
		if err != nil {
			fmt.Println("❌ Failed to upload image to S3:", err.Error())
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload profile image"})
			return
		}

		fmt.Println("✅ Image uploaded to S3 successfully. URL:", url)
		profileImageURL = &url
	} else {
		fmt.Println("⚠️ No image uploaded or error reading image:", err)
	}

	// 6. Create user profile model
	profile := &models.UserProfile{
		UserID:              uid,
		ProfileImage:        profileImageURL,
		FullName:            cleanString(input.FullName),
		Designation:         stringPtr(input.Designation),
		Organization:        stringPtr(input.Organization),
		ProfessionalSummary: stringPtr(input.ProfessionalSummary),
		Location:            stringPtr(input.Location),
		Email:               cleanString(input.Email),
		ContactNumber:       stringPtr(input.ContactNumber),
	}

	// 7. Save to DB
	fmt.Println("💾 Saving user profile to database")
	if _, err := ctrl.UserProfileService.Create(context.Background(), profile); err != nil {
		fmt.Println("❌ Failed to create user profile:", err.Error())
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile"})
		return
	}

	fmt.Println("✅ User profile created successfully:", profile.ID.String())
	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User profile created successfully",
		"profile": profile,
	})
}

func cleanString(s string) string {
	s = strings.TrimSpace(s)
	unquoted, err := strconv.Unquote(s)
	if err != nil {
		return s
	}
	return unquoted
}

func stringPtr(s string) *string {
	cleaned := cleanString(s)
	if cleaned == "" {
		return nil
	}
	return &cleaned
}

func (ctrl *UserProfileController) GetByUserID(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	profile, err := ctrl.UserProfileService.GetByUserID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	ctx.JSON(http.StatusOK, profile)
}

func (ctrl *UserProfileController) GetAll(ctx *gin.Context) {
	profiles, err := ctrl.UserProfileService.GetAll(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch profiles"})
		return
	}
	ctx.JSON(http.StatusOK, profiles)
}

func (ctrl *UserProfileController) Update(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	var input models.UserProfile
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedProfile, err := ctrl.UserProfileService.Update(context.Background(), userID, &input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, updatedProfile)
}

func (ctrl *UserProfileController) Delete(ctx *gin.Context) {
	userID := ctx.Param("user_id")

	if err := ctrl.UserProfileService.Delete(context.Background(), userID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete profile", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User profile deleted successfully"})
}
