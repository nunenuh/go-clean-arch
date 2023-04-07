package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}

func ProviderUserRepository(DB *gorm.DB) UserRepository {
	return &userRepository{DB: DB}
}
