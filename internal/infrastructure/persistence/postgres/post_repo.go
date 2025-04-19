package postgres

import (
	"context"
	"database/sql"
	apperrors "social-comments/internal/core/errors"

	"github.com/jmoiron/sqlx"
	"social-comments/internal/core/domain"
)

type PostRepository struct {
	db *sqlx.DB
}

func NewPostRepository(db *sqlx.DB) *PostRepository {
	return &PostRepository{db: db}
}

func (r *PostRepository) Create(ctx context.Context, post domain.Post) (domain.Post, error) {
	const query = `
		INSERT INTO posts (title, content, author, allow_comments) 
		VALUES ($1, $2, $3, $4)
		RETURNING id, created_at`

	var createdPost domain.Post

	err := r.db.QueryRowxContext(ctx, query,
		post.Title,
		post.Content,
		post.Author,
		post.AllowComments,
	).StructScan(&createdPost)

	if err != nil {
		return domain.Post{}, apperrors.ErrCreatingPost
	}

	return createdPost, nil
}

func (r *PostRepository) GetByID(ctx context.Context, id int) (domain.Post, error) {
	const query = `
		SELECT id, created_at, title, content, author, allow_comments
		FROM posts
		WHERE id = $1`

	var post domain.Post
	if err := r.db.GetContext(ctx, &post, query, id); err != nil {
		if err == sql.ErrNoRows {
			return domain.Post{}, apperrors.ErrPostNotFound
		}
		return domain.Post{}, apperrors.ErrPostNotFoundByID
	}

	return post, nil
}

func (r *PostRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Post, error) {
	query := `
		SELECT id, created_at, title, content, author, allow_comments
		FROM posts
		ORDER BY created_at DESC
		OFFSET $1`

	args := []interface{}{offset}

	if limit > 0 {
		query += " LIMIT $2"
		args = append(args, limit)
	}

	var posts []domain.Post
	if err := r.db.SelectContext(ctx, &posts, query, args...); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, apperrors.ErrGetAllPosts
	}

	return posts, nil
}
