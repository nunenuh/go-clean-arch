package user

import (
	"context"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Created  int64  `json:"created"`
	Modified int64  `json:"modified"`
	Deleted  int64  `json:"deleted"`
}

type UserRepository interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}

type UserService interface {
	GetUsers(ctx context.Context) (*[]User, error)
	GetUser(ctx context.Context, userID int) (*User, error)
	CreateUser(ctx context.Context, user *User) error
	UpdateUser(ctx context.Context, userID int, user *User) error
	DeleteUser(ctx context.Context, userID int) error
}
