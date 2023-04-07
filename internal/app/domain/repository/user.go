package repository

import (
	"context"

	"github.com/nunenuh/go-clean-arch/internal/app/domain/entity"
)

type UserRepository interface {
	GetUsers(ctx context.Context) (*[]entity.User, error)
	GetUser(ctx context.Context, userID int) (*entity.User, error)
	CreateUser(ctx context.Context, user *entity.User) error
	UpdateUser(ctx context.Context, userID int, user *entity.User) error
	DeleteUser(ctx context.Context, userID int) error
}
