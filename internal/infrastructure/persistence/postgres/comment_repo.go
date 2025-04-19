package postgres

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"social-comments/internal/core/domain"
	apperrors "social-comments/internal/core/errors"
)

type CommentRepository struct {
	db *sqlx.DB
}

func NewCommentRepository(db *sqlx.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (r *CommentRepository) Create(ctx context.Context, comment domain.Comment) (domain.Comment, error) {
	const query = `
		INSERT INTO comments (text, author, post_id, parent_id) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	var createdComment domain.Comment

	err := r.db.QueryRowxContext(ctx, query,
		comment.Text,
		comment.Author,
		comment.PostID,
		comment.ParentID,
	).StructScan(&createdComment)

	if err != nil {
		return domain.Comment{}, apperrors.ErrCreatingComment
	}

	return createdComment, nil
}

func (r *CommentRepository) GetByPostID(ctx context.Context, postID int, limit, offset int) ([]domain.Comment, error) {
	query := `
		SELECT id, created_at, author, text, post_id, parent_id
		FROM comments
		WHERE post_id = $1 AND parent_id IS NULL
		ORDER BY created_at DESC
		OFFSET $2`

	args := []interface{}{postID, offset}

	if limit > 0 {
		query += " LIMIT $3"
		args = append(args, limit)
	}

	var comments []domain.Comment
	if err := r.db.SelectContext(ctx, &comments, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrCommentNotFound
		}
		return nil, apperrors.GetCommentsByPost
	}

	return comments, nil
}

func (r *CommentRepository) GetReplies(ctx context.Context, parentID int) ([]domain.Comment, error) {
	const query = `
		SELECT id, created_at, author, text, post_id, parent_id
		FROM comments
		WHERE parent_id = $1
		ORDER BY created_at ASC`

	var replies []domain.Comment
	if err := r.db.SelectContext(ctx, &replies, query, parentID); err != nil {
		if err == sql.ErrNoRows {
			return nil, apperrors.ErrCommentNotFound
		}
		return nil, apperrors.ErrRepliesNotFound
	}

	return replies, nil
}
