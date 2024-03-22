package repository

import (
	"task-go/model"

	"gorm.io/gorm"
)

type UserRepositoryIF interface {
	CreateUser(user *model.User) error
	GetUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryIF {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUser(user *model.User) error {
	if err := ur.db.Where("name=? AND password=?", user.Name, user.Password).First(user).Error; err != nil {

		return err
	}

	return nil
}
