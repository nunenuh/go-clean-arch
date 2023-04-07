package post

import (
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
