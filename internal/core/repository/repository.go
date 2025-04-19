package repository

import (
	"context"
	"social-comments/internal/core/domain"
)

type Gateways struct {
	PostRepository
	CommentRepository
}

func NewGateways(posts PostRepository, comments CommentRepository) *Gateways {
	return &Gateways{
		PostRepository:    posts,
		CommentRepository: comments,
	}
}

type PostRepository interface {
	Create(ctx context.Context, post domain.Post) (domain.Post, error)
	GetByID(ctx context.Context, id int) (domain.Post, error)
	GetAll(ctx context.Context, limit, offset int) ([]domain.Post, error)
}

type CommentRepository interface {
	Create(ctx context.Context, comment domain.Comment) (domain.Comment, error)
	GetByPostID(ctx context.Context, postID, limit, offset int) ([]domain.Comment, error)
	GetReplies(ctx context.Context, parentID int) ([]domain.Comment, error)
}
