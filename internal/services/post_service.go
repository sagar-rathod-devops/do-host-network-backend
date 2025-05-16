package services

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/models"
	"github.com/sagar-rathod-devops/do-host-network-backend/internal/repositories"
)

type PostService struct {
	Repo *repositories.PostRepository
}

func NewPostService(repo *repositories.PostRepository) *PostService {
	return &PostService{Repo: repo}
}

func (s *PostService) CreatePost(ctx context.Context, p *models.ContentPost) (*models.ContentPost, error) {
	createdPost, err := s.Repo.CreatePost(ctx, p)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func (s *PostService) GetAllContentPosts(ctx context.Context) ([]models.PostWithDetails, error) {
	return s.Repo.GetAllWithDetails(ctx)
}

func (s *PostService) GetPostsByUserID(userID uuid.UUID) ([]models.ContentPost, error) {
	posts, err := s.Repo.GetPostsByUserID(userID)
	if err != nil {
		return nil, err
	}
	if len(posts) == 0 {
		return nil, errors.New("no posts found for this user")
	}
	return posts, nil
}
