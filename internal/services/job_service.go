package services

import (
	"context"

	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type JobService struct {
	Repo *repositories.JobRepository
}

func (s *JobService) CreateJobPost(ctx context.Context, post *models.JobPost) (*models.JobPost, error) {
	if err := s.Repo.CreateJobPost(post); err != nil {
		return nil, err
	}
	return post, nil
}

func (s *JobService) GetAllJobPosts(ctx context.Context) ([]models.JobPost, error) {
	return s.Repo.GetAll(ctx)
}
