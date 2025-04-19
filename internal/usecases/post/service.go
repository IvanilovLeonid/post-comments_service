package post

import (
	"context"
	"errors"
	"fmt"
	"log"
	"social-comments/internal/core/repository"

	"social-comments/internal/core/domain"
	"social-comments/internal/core/errors"
	"social-comments/pkg/logging"
	"social-comments/pkg/utils/pagination"
)

type Service struct {
	repo   repository.PostRepository
	logger logger.Logger
}

func NewService(repo repository.PostRepository, logger logger.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) CreatePost(ctx context.Context, req domain.CreatePostRequest) (*domain.PostResponse, error) {
	if err := validatePostRequest(req); err != nil {
		s.logger.Error("validation error: %v", err)
		return &domain.PostResponse{}, err
	}

	post := domain.Post{
		Title:         req.Title,
		Author:        req.Author,
		Content:       req.Content,
		AllowComments: req.AllowComments,
	}

	createdPost, err := s.repo.Create(ctx, post) // ОШИБКА Cannot use 'post' (type domain. Post) as the type domain. CreatePostRequest
	if err != nil {
		s.logger.Error("failed to create post: %v", err)
		return &domain.PostResponse{}, apperrors.ErrCreatingPost
	}

	return &domain.PostResponse{
		ID:        createdPost.ID,
		Title:     createdPost.Title,
		Author:    createdPost.Author,
		Content:   createdPost.Content,
		CreatedAt: createdPost.CreatedAt,
	}, nil
}

func (s *Service) GetPostByID(ctx context.Context, id int) (*domain.Post, error) {
	if id <= 0 {
		return &domain.Post{}, errors.New("invalid post ID")
	}

	post, err := s.repo.GetByID(ctx, id)
	if err != nil {
		s.logger.Error("failed to get post by ID: %v", err)
		return &domain.Post{}, apperrors.ErrPostNotFoundByID
	}

	return &post, nil
}

func (s *Service) GetAllPosts(ctx context.Context, page, pageSize int) ([]domain.Post, error) {
	paginator := pagination.NewPaginator(page, pageSize)
	posts, err := s.repo.GetAll(ctx, paginator.Limit(), paginator.Offset())

	log.Printf("Service GetAllPosts: limit=%d, offset=%d", paginator.Limit(), paginator.Offset())

	if err != nil {
		s.logger.Error("failed to get posts: %v", err)
		return nil, apperrors.ErrPostNotFound
	}

	log.Printf("Repository returned %d posts", len(posts))

	return posts, nil
}

func validatePostRequest(req domain.CreatePostRequest) error {
	if req.Title == "" {
		return errors.New("post title cannot be empty")
	}
	if req.Author == "" {
		return errors.New("post author cannot be empty")
	}
	if len(req.Content) > apperrors.MaxContentLength {
		return fmt.Errorf("content exceeds maximum length of %d", apperrors.MaxContentLength)
	}
	return nil
}
