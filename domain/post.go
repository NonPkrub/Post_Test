package domain

import (
	"time"
)

type Post struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	ViewCount int       `json:"view_count"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type PostReq struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostUpdateReq struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Published bool   `json:"published"`
	ViewCount int    `json:"view_count"`
}

type PostRes struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Published bool      `json:"published"`
	CreatedAt time.Time `json:"created_at"`
}

type PostAllReq struct {
	Published bool      `json:"published"`
	Title     string    `json:"title"`
	CreatedAt time.Time `json:"created_at"`
}

type Pagination struct {
	Page      int `json:"page"`
	PageSize  int `json:"limit"`
	Count     int `json:"count"`
	TotalPage int `json:"total_page"`
}

type PostResponse struct {
	Posts     []PostRes `json:"posts"`
	Count     int       `json:"count"`
	Limit     int       `json:"limit"`
	Page      int       `json:"page"`
	TotalPage int       `json:"total_page"`
}

type PostUseCase interface {
	Create(post *PostReq) (*PostRes, error)
	GetAll(query *PostAllReq, pagination *Pagination) (*PostResponse, error)
	GetByID(id string) (*PostRes, error)
	UpdateByID(post *PostUpdateReq) (*PostRes, error)
	DeleteByID(id string) error
}

type PostRepository interface {
	Create(post *Post) (*Post, error)
	FindAllField(query *PostAllReq, pagination *Pagination) ([]*Post, int64, int64, error)
	FindOne(post *Post) (*Post, error)
	UpdateByID(post *Post) (*Post, error)
	DeleteByID(post *Post) error
}
