package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sagar-rathod-devops/do-host-network-backend/config"
	"github.com/sagar-rathod-devops/do-host-network-backend/utils"
)

type UploadController struct {
	Uploader *utils.S3Uploader
}

func NewUploadController(cfg config.Config) (*UploadController, error) {
	uploader, err := utils.NewS3Uploader(cfg)
	if err != nil {
		return nil, err
	}
	return &UploadController{Uploader: uploader}, nil
}

func (ctrl *UploadController) UploadFile(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "File is required"})
		return
	}

	// You can customize the S3 key (file path in the bucket) as needed
	key := fileHeader.Filename

	url, err := ctrl.Uploader.UploadFile(file, fileHeader, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file", "details": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"url": url})
}
