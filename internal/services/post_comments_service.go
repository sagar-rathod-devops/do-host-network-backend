package services

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type PostCommentService struct {
	PostCommentRepository *repositories.PostCommentRepository
}

func (service *PostCommentService) CommentOnPost(userID, postID uuid.UUID, comment string) error {
	exists, err := service.PostCommentRepository.PostExists(postID)
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("post not found")
	}
	return service.PostCommentRepository.CreateComment(userID, postID, comment)
}

func (service *PostCommentService) GetPostComments(postID uuid.UUID) ([]models.PostComment, error) {
	return service.PostCommentRepository.GetComments(postID)
}
