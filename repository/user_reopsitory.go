package repository

import (
	"task-go/domain"

	"gorm.io/gorm"
)

type UserRepositoryIF interface {
	RegistUser(user *domain.UserIF) error
	GetUser(user *domain.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepositoryIF {
	return &userRepository{db}
}

func (ur *userRepository) RegistUser(user *domain.UserIF) error {

	err := ur.db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

func (ur *userRepository) GetUser(user *domain.User) error {

	// err := ur.db.Where("name=? AND password=?", user.Name, user.Password).First(user).Error

	// if err != nil {
	// 	return err
	// }

	return nil
}
