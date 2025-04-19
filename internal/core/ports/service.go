package ports

import (
	"context"
	"social-comments/internal/core/domain"
	"social-comments/internal/core/repository"
	"social-comments/internal/usecases/comment"
	"social-comments/internal/usecases/post"
	"social-comments/pkg/logging"
)

type Services struct {
	PostService
	CommentService
}

func NewServices(gateways *repository.Gateways, logger *logger.Logger) *Services {
	return &Services{
		PostService:    post.NewService(gateways.PostRepository, *logger),
		CommentService: comment.NewService(gateways.CommentRepository, gateways.PostRepository, *logger),
	}
}

type PostService interface {
	CreatePost(ctx context.Context, req domain.CreatePostRequest) (*domain.PostResponse, error)
	GetPostByID(ctx context.Context, id int) (*domain.Post, error)
	GetAllPosts(ctx context.Context, page, pageSize int) ([]domain.Post, error)
}

type CommentService interface {
	CreateComment(ctx context.Context, comment domain.CreateCommentRequest) (*domain.Comment, error)
	GetCommentsByPostID(ctx context.Context, postID, page, pageSize int) ([]domain.Comment, error)
	GetCommentReplies(ctx context.Context, commentID int) ([]*domain.Comment, error)
}
