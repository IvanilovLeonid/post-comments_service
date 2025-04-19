package memory

import (
	"context"
	"social-comments/internal/core/domain"
	"social-comments/internal/core/errors"
	"sync"
	"time"
)

type CommentRepository struct {
	mu       sync.RWMutex
	comments []domain.Comment
	lastID   int
}

func NewCommentRepository(initialCapacity int) *CommentRepository {
	return &CommentRepository{
		comments: make([]domain.Comment, 0, initialCapacity),
		lastID:   0,
	}
}

func (r *CommentRepository) Create(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(comment.Text) > apperrors.MaxContentLength {
		return domain.Comment{}, apperrors.ErrContentTooLong
	}

	r.lastID++
	newComment := domain.Comment{
		ID:        r.lastID,
		CreatedAt: time.Now(),
		Author:    comment.Author,
		Text:      comment.Text,
		PostID:    comment.PostID,
		ParentID:  comment.ParentID,
	}

	r.comments = append(r.comments, newComment)
	return newComment, nil
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID, limit, offset int) ([]domain.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if offset < 0 || limit < 0 {
		return nil, apperrors.ErrInvalidPagination
	}

	var result []domain.Comment
	for _, c := range r.comments {
		if c.ParentID == nil && c.PostID == postID {
			result = append(result, c)
		}
	}

	if offset >= len(result) {
		return nil, nil
	}

	end := offset + limit
	if end > len(result) || limit == 0 {
		end = len(result)
	}

	return result[offset:end], nil
}

func (r *CommentRepository) GetReplies(ctx context.Context, parentID int) ([]domain.Comment, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var replies []domain.Comment
	for _, c := range r.comments {
		if c.ParentID != nil && *c.ParentID == parentID {
			replies = append(replies, c)
		}
	}

	return replies, nil
}
