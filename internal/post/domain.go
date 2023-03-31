package post

import (
	"context"
	"go-post-clean-arch/internal/user"
)

type Post struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Content  string `json:"content"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
	Deleted  int64  `json:"deleted"`
	User     int    `json:"user"`
}

type PostAndUser struct {
	ID       int       `json:"id"`
	Title    string    `json:"title"`
	Content  string    `json:"content"`
	Created  int64     `json:"created"`
	Modified int64     `json:"modified"`
	Deleted  int64     `json:"deleted"`
	User     user.User `json:"user"`
}

type PostRepository interface {
	GetPosts(ctx context.Context) (*[]PostAndUser, error)
	GetPost(ctx context.Context, postID int) (*PostAndUser, error)
	CreatePost(ctx context.Context, post *Post) error
	UpdatePost(ctx context.Context, postID int, post *Post) error
	DeletePost(ctx context.Context, postID int) error
}

type PostService interface {
	FetchPosts(ctx context.Context) (*[]PostAndUser, error)
	FetchPost(ctx context.Context, postID int) (*PostAndUser, error)
	BuildPost(ctx context.Context, post *Post, userID int) error
	ModifyPost(ctx context.Context, postID int, post *Post, userID int) error
	DestroyPost(ctc context.Context, postID int) error
}
