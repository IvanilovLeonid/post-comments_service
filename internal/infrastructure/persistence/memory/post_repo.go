package memory

import (
	"context"
	"fmt"
	"social-comments/internal/core/domain"
	apperrors "social-comments/internal/core/errors"
	"sort"
	"sync"
	"time"
)

type PostRepository struct {
	mu     sync.RWMutex
	posts  []domain.Post
	lastID int
}

func NewPostRepository(initialCapacity int) *PostRepository {
	return &PostRepository{
		posts:  make([]domain.Post, 0, initialCapacity),
		lastID: 0,
	}
}

func (r *PostRepository) Create(ctx context.Context, post domain.Post) (domain.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.lastID++
	newPost := domain.Post{
		ID:            r.lastID,
		CreatedAt:     time.Now(),
		Title:         post.Title,
		Author:        post.Author,
		Content:       post.Content,
		AllowComments: post.AllowComments,
	}

	r.posts = append(r.posts, newPost)
	return newPost, nil
}

func (r *PostRepository) GetByID(ctx context.Context, id int) (domain.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if id <= 0 || id > r.lastID {
		return domain.Post{}, apperrors.ErrPostNotFound
	}

	return r.posts[id-1], nil
}

func (r *PostRepository) GetAll(ctx context.Context, limit, offset int) ([]domain.Post, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	// Всегда делаем полный дамп для отладки
	fmt.Println("\n=== Repository State ===")
	fmt.Printf("Total posts: %d\n", len(r.posts))
	fmt.Printf("Requested limit: %d, offset: %d\n", limit, offset)

	if offset < 0 || limit < 0 {
		return nil, apperrors.ErrInvalidPagination
	}

	// Копируем и сортируем посты (новые сначала)
	sortedPosts := make([]domain.Post, len(r.posts))
	copy(sortedPosts, r.posts)
	sort.Slice(sortedPosts, func(i, j int) bool {
		return sortedPosts[i].CreatedAt.After(sortedPosts[j].CreatedAt)
	})

	// Корректируем offset
	if offset >= len(sortedPosts) {
		return []domain.Post{}, nil
	}

	end := offset + limit
	if end > len(sortedPosts) {
		end = len(sortedPosts)
	}

	result := sortedPosts[offset:end]
	fmt.Printf("Returning %d posts (IDs: %v)\n", len(result), getPostIDs(result))
	return result, nil
}

func getPostIDs(posts []domain.Post) []int {
	ids := make([]int, len(posts))
	for i, p := range posts {
		ids[i] = p.ID
	}
	return ids
}

func (r *PostRepository) DebugDumpPosts() {
	r.mu.RLock()
	defer r.mu.RUnlock()

	fmt.Println("\n=== Current Posts Dump ===")
	fmt.Printf("Total posts: %d\n", len(r.posts))
	for i, post := range r.posts {
		fmt.Printf("[%d] ID: %d, Title: %s, CreatedAt: %v\n",
			i, post.ID, post.Title, post.CreatedAt)
	}
}
