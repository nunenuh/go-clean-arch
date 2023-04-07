package usecase

import (
	"context"
	"time"

	"github.com/nunenuh/go-clean-arch/internal/app/domain/entity"
	"github.com/nunenuh/go-clean-arch/internal/app/domain/repository"
)

type userService struct {
	userRepository repository.UserRepository
}

// create new 'service' or 'use-case' for User Entity
func NewUserService(r UserRepository) UserService {
	return &userService{
		userRepository: r,
	}
}

// implementation of FetchUsers
func (s *userService) FetchUsers(ctx context.Context) (*[]entity.User, error) {
	return s.userRepository.GetUsers(ctx)
}

// implementation of FetchUser
func (s *userService) FetchUser(ctx context.Context, userID int) (*entity.User, error) {
	return s.userRepository.GetUser(ctx, userID)
}

// implementation of CreateUser
func (s *userService) CreateUser(ctx context.Context, user *entity.User) error {
	// Set default value of 'Created' and 'Modified'.
	user.Created = time.Now().Unix()
	user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.CreateUser(ctx, user)
}

// implementation of UpdateUser
func (s *userService) UpdateUser(ctx context.Context, userID int, user *entity.User) error {
	// Set default value of 'Modified'.
	user.Modified = time.Now().Unix()

	// Pass to the repository layer.
	return s.userRepository.UpdateUser(ctx, userID, user)
}

// implementation of DeleteUser
func (s *userService) DeleteUser(ctx context.Context, userID int) error {
	// Pass to the repository layer.
	return s.userRepository.DeleteUser(ctx, userID)
}
