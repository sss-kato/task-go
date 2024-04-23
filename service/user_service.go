package service

import (
	"task-go/domain"
	"task-go/dto"
	"task-go/repository"
)

type UserServiceIF interface {
	Signup(user domain.UserIF) (domain.UserResponse, error)
	Login(user domain.UserIF) (string, error)
}

type userService struct {
	ur repository.UserRepositoryIF
}

func NewUserService(ur repository.UserRepositoryIF) UserServiceIF {
	return &userService{ur}
}

func (us *userService) Signup(user domain.UserIF) (domain.UserResponse, error) {

	user.HashedPassword()
	ud := &dto.UserDto{Name: user.GetName(), Mailadress: user.GetMailAdress(), Password: user.GetPassWord()}

	err := us.ur.RegistUser(ud)
	if err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{ID: ud.ID, Name: ud.Name}
	return resUser, nil
}

func (us *userService) Login(user domain.UserIF) (string, error) {
	// user.Password = hashed(user.Password)

	// err := us.ur.GetUser(user)
	// if err != nil {
	// 	return "", err
	// }
	token := "test"
	return token, nil
}
