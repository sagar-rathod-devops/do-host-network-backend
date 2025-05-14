package controllers

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/services"
)

type JobController struct {
	JobService *services.JobService
}

func (jc *JobController) CreateJobPost(ctx *gin.Context) {
	var input struct {
		UserID          string `json:"user_id" binding:"required"`
		JobTitle        string `json:"job_title" binding:"required"`
		CompanyName     string `json:"company_name" binding:"required"`
		JobDescription  string `json:"job_description" binding:"required"`
		JobApplyURL     string `json:"job_apply_url"`
		Location        string `json:"location"`
		LastDateToApply string `json:"last_date_to_apply"` // Expecting format: YYYY-MM-DD
	}

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input: " + err.Error()})
		return
	}

	userID, err := uuid.Parse(input.UserID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var lastDate time.Time
	if input.LastDateToApply != "" {
		lastDate, err = time.Parse("2006-01-02", input.LastDateToApply)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD"})
			return
		}
	} else {
		lastDate = time.Now().AddDate(0, 1, 0) // Default 1 month
	}

	now := time.Now()

	jobPost := &models.JobPost{
		UserID:          userID,
		JobTitle:        input.JobTitle,
		CompanyName:     input.CompanyName,
		JobDescription:  input.JobDescription,
		JobApplyURL:     input.JobApplyURL,
		Location:        input.Location,
		LastDateToApply: lastDate,
		PostDate:        now,
		CreatedAt:       now,
	}

	_, err = jc.JobService.CreateJobPost(context.Background(), jobPost)
	if err != nil {
		log.Printf("CreateJobPost error: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create job post"})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Job post created successfully"})
}

func (jc *JobController) GetAllJobPosts(ctx *gin.Context) {
	jobPosts, err := jc.JobService.GetAllJobPosts(context.Background())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve job posts"})
		return
	}

	ctx.JSON(http.StatusOK, jobPosts)
}
