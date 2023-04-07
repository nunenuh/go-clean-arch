package repository

import (
	"internal/app/adapter/database"
	"internal/app/adapter/database/model"
	"internal/app/domain/entity"
)

type User struct{}

func (u User) GetUser() entity.User {
	db, err := database.Connection()
	var user model.User

	result := db.First(&user, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Created:  user.Created,
		Modified: user.Modified,
		Deleted:  user.Deleted,
	}
}

func (u User) GetUsers() *[]entity.User {
	db, err := database.Connection()
	var user model.User

	result := db.First(&user, 1)
	if result.Error != nil {
		panic(result.Error)
	}
	return entity.User{
		ID:       user.ID,
		Name:     user.Name,
		Username: user.Username,
		Password: user.Password,
		Email:    user.Email,
		Phone:    user.Phone,
		Created:  user.Created,
		Modified: user.Modified,
		Deleted:  user.Deleted,
	}
}
