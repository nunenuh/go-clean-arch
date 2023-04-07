package repository

import (
	"context"

	"github.com/nunenuh/go-clean-arch/internal/app/domain/entity"
)

type PostRepository interface {
	GetPosts(ctx context.Context) (*[]entity.PostAndUser, error)
	GetPost(ctx context.Context, postID int) (*entity.PostAndUser, error)
	CreatePost(ctx context.Context, post *entity.Post) error
	UpdatePost(ctx context.Context, postID int, post *entity.Post) error
	DeletePost(ctx context.Context, postID int) error
}
