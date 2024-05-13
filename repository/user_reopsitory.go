package repository

import (
	"task-go/dto"

	"github.com/cockroachdb/errors"

	"gorm.io/gorm"
)

type UserRepositoryIF interface {
	RegistUser(ud *dto.UserDto) error
	GetUser(user *dto.UserDto) (int, error)
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
		return errors.New(err.Error())
	}

	return nil
}

func (ur *userRepository) GetUser(user *dto.UserDto) (int, error) {

	result := ur.db.Where("name=? AND password=?", user.Name, user.Password).First(user)
	err := result.Error
	userCnt := result.RowsAffected

	if err != nil {
		return int(userCnt), errors.New(err.Error())
	}

	return int(userCnt), nil
}
