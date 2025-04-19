package domain

import (
	"strconv"
	"time"
)

// Comment представляет основную сущность комментария
type Comment struct {
	ID        int       `json:"id" db:"id"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	Author    string    `json:"author" db:"author"`
	Text      string    `json:"text" db:"text"`
	PostID    int       `json:"postId" db:"post_id"`
	ParentID  *int      `json:"parentId,omitempty" db:"parent_id"`
}

// CreateCommentRequest DTO для создания комментария
type CreateCommentRequest struct {
	Author   string `json:"author"`
	Text     string `json:"text"`
	PostID   int    `json:"postId"`
	ParentID *int   `json:"parentId,omitempty"`
}

// ToDomain преобразует DTO в доменную модель
func (r CreateCommentRequest) ToDomain() Comment {
	return Comment{
		Author:   r.Author,
		Text:     r.Text,
		PostID:   r.PostID,
		ParentID: r.ParentID,
	}
}

func (c Comment) GetID() string {
	return strconv.Itoa(c.ID)
}
