package usecase

import (
	"context"

	"github.com/nunenuh/go-clean-arch/internal/app/domain/entity"
)

// UserService is the interface for user usecase
type UserService interface {
	FetchUsers(ctx context.Context) (*[]entity.User, error)
	FetchUser(ctx context.Context, userID int) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, userID int, user *entity.User) error
	DeleteUser(ctx context.Context, userID int) error
}
