package controllers

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type UserProfileController struct {
	UserProfileService *services.UserProfileService
}

func NewUserProfileController(service *services.UserProfileService) *UserProfileController {
	return &UserProfileController{UserProfileService: service}
}

func (ctrl *UserProfileController) Create(ctx *gin.Context) {
	var input struct {
		UserID              string `json:"user_id" binding:"required"`
		FullName            string `json:"full_name" binding:"required"`
		Designation         string `json:"designation"`
		Organization        string `json:"organization"`
		ProfessionalSummary string `json:"professional_summary"`
		Location            string `json:"location"`
		Email               string `json:"email" binding:"required"`
		ContactNumber       string `json:"contact_number"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	uid, err := uuid.Parse(input.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid UUID format"})
		return
	}

	profile := &models.UserProfile{
		UserID:              uid,
		FullName:            input.FullName,
		Designation:         input.Designation,
		Organization:        input.Organization,
		ProfessionalSummary: input.ProfessionalSummary,
		Location:            input.Location,
		Email:               input.Email,
		ContactNumber:       input.ContactNumber,
	}

	if _, err := ctrl.UserProfileService.Create(context.Background(), profile); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user profile", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"message": "User profile created successfully",
	})

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
