package repository

import (
	"task-go/dto"

	"gorm.io/gorm"
)

type UserRepositoryIF interface {
	RegistUser(ud *dto.UserDto) error
	GetUser(user *dto.UserDto) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryIF {
	return &userRepository{db}
}

func (ur *userRepository) RegistUser(ud *dto.UserDto) error {

	err := ur.db.Create(ud).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUser(user *dto.UserDto) error {

	// err := ur.db.Where("name=? AND password=?", user.Name, user.Password).First(user).Error

	// if err != nil {
	// 	return err
	// }

	return nil
}
