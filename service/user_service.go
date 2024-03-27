package service

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"task-go/model"
	"task-go/repository"

	"golang.org/x/crypto/pbkdf2"
)

type UserServiceIF interface {
	Signup(user *model.User) (model.UserResponse, error)
	Login(user *model.User) (string, error)
}

type userService struct {
	ur repository.UserRepositoryIF
}

func NewUserService(ur repository.UserRepositoryIF) UserServiceIF {
	return &userService{ur}
}

func (us *userService) Signup(user *model.User) (model.UserResponse, error) {

	newUser := model.User{Name: user.Name, Mailadress: user.Mailadress, Password: hashed(user.Password)}
	err := us.ur.RegistUser(&newUser)
	if err != nil {
		return model.UserResponse{}, err
	}

	resUser := model.UserResponse{ID: newUser.ID, Name: newUser.Name}
	return resUser, nil
}

func (us *userService) Login(user *model.User) (string, error) {
	user.Password = hashed(user.Password)

	err := us.ur.GetUser(user)
	if err != nil {
		return "", err
	}
	token := "test"
	return token, nil
}

func hashed(p string) string {
	salt := base64.StdEncoding.EncodeToString([]byte(p))

	key := pbkdf2.Key([]byte(p), []byte(salt), 10, 10, sha256.New)

	return hex.EncodeToString(key[:])
}
