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

type UserEducationController struct {
	Service *services.UserEducationService
}

// POST /api/education
func (c *UserEducationController) Create(ctx *gin.Context) {
	var input struct {
		UserID          uuid.UUID `json:"user_id"`
		Degree          string    `json:"degree"`
		InstitutionName string    `json:"institution_name"`
		FieldOfStudy    string    `json:"field_of_study"`
		Grade           string    `json:"grade"`
		Year            string    `json:"year"` // e.g., "2023"
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	if input.UserID == uuid.Nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "user_id is required"})
		return
	}

	education := &models.UserEducation{
		ID:              uuid.New(),
		UserID:          input.UserID,
		Degree:          input.Degree,
		InstitutionName: input.InstitutionName,
		FieldOfStudy:    input.FieldOfStudy,
		Grade:           input.Grade,
		Year:            input.Year,
	}

	if err := c.Service.Create(context.Background(), education); err != nil {
		fmt.Printf("DEBUG: Failed to create education entry: %v\n", err) // <--- add this line
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create education entry", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "User education created successfully"})
}

// GET /api/education/:user_id
func (c *UserEducationController) GetByUser(ctx *gin.Context) {
	userIDStr := ctx.Param("user_id")
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user_id"})
		return
	}

	educations, err := c.Service.GetByUserID(context.Background(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, educations)
}

// PUT /api/education/:id
func (c *UserEducationController) Update(ctx *gin.Context) {
	idStr := ctx.Param("id")
	eduID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid education ID"})
		return
	}

	var input struct {
		Degree          string `json:"degree"`
		InstitutionName string `json:"institution_name"`
		FieldOfStudy    string `json:"field_of_study"`
		Grade           string `json:"grade"`
		Year            string `json:"year"`
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	updatedEdu := &models.UserEducation{
		ID:              eduID,
		Degree:          input.Degree,
		InstitutionName: input.InstitutionName,
		FieldOfStudy:    input.FieldOfStudy,
		Grade:           input.Grade,
		Year:            input.Year,
	}

	if err := c.Service.Update(context.Background(), updatedEdu); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update education", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User education updated successfully"})
}

// DELETE /api/education/:id
func (c *UserEducationController) Delete(ctx *gin.Context) {
	idStr := ctx.Param("id")
	eduID, err := uuid.Parse(idStr)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid education ID"})
		return
	}

	if err := c.Service.Delete(context.Background(), eduID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete education", "details": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User education deleted successfully"})
}
