package controllers

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type UserExperienceController struct {
	UserExperienceService *services.UserExperienceService
}

func NewUserExperienceController(service *services.UserExperienceService) *UserExperienceController {
	return &UserExperienceController{UserExperienceService: service}
}

func (c *UserExperienceController) Create(ctx *gin.Context) {
	var input models.UserExperience

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if input.UserID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	if err := c.UserExperienceService.Create(context.Background(), &input); err != nil {
		fmt.Printf("DEBUG: Failed to create user experience: %v\n", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user experience", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User experience created successfully"})
}

func (c *UserExperienceController) GetByUserID(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid user_id format"})
		return
	}

	experiences, err := c.UserExperienceService.GetByUserID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if len(experiences) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "No user experience found"})
		return
	}

	ctx.JSON(http.StatusOK, experiences)
}

// PUT /api/experience/:id
func (c *UserExperienceController) Update(ctx *gin.Context) {
	experienceID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experience ID"})
		return
	}

	var input models.UserExperience
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	input.ID = experienceID
	if err := c.UserExperienceService.Update(context.Background(), &input); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user experience", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User experience updated successfully"})
}

// DELETE /api/experience/:id
func (c *UserExperienceController) Delete(ctx *gin.Context) {
	experienceID, err := uuid.Parse(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid experience ID"})
		return
	}

	if err := c.UserExperienceService.Delete(context.Background(), experienceID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user experience", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User experience deleted successfully"})
}
