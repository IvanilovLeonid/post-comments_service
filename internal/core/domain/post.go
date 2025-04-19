package domain

import (
	"strconv"
	"time"
)

// Post представляет основную сущность поста
type Post struct {
	ID            int       `json:"id" db:"id"`
	CreatedAt     time.Time `json:"createdAt" db:"created_at"`
	Title         string    `json:"title" db:"title"`
	Author        string    `json:"author" db:"author"`
	Content       string    `json:"content" db:"content"`
	AllowComments bool      `json:"allowComments" db:"allow_comments"`
}

// CreatePostRequest DTO для создания поста
type CreatePostRequest struct {
	Title         string `json:"title"`
	Author        string `json:"author"`
	Content       string `json:"content"`
	AllowComments bool   `json:"allowComments"`
}

// ToDomain преобразует DTO в доменную модель
func (r CreatePostRequest) ToDomain() Post {
	return Post{
		Title:         r.Title,
		Author:        r.Author,
		Content:       r.Content,
		AllowComments: r.AllowComments,
	}
}

// PostResponse DTO для ответа с постом
type PostResponse struct {
	ID        int       `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Title     string    `json:"title"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
}

// ToResponse преобразует доменную модель в DTO ответа
func (p Post) ToResponse() PostResponse {
	return PostResponse{
		ID:        p.ID,
		CreatedAt: p.CreatedAt,
		Title:     p.Title,
		Author:    p.Author,
		Content:   p.Content,
	}
}

func (p Post) GetID() string {
	return strconv.Itoa(p.ID)
}
