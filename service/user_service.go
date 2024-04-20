package service

import (
	"task-go/domain"
	"task-go/repository"
)

type UserServiceIF interface {
	Signup(user *domain.User) (domain.UserResponse, error)
	Login(user *domain.User) (string, error)
}

type userService struct {
	ur repository.UserRepositoryIF
}

func NewUserService(ur repository.UserRepositoryIF) UserServiceIF {
	return &userService{ur}
}

func (us *userService) Signup(user *domain.User) (domain.UserResponse, error) {

	// newUser := model.User{Name: user.Name, Mailadress: user.Mailadress, Password: hashed(user.Password)}
	err := us.ur.RegistUser(user)
	if err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{ID: user.ID, Name: user.Name}
	return resUser, nil
}

func (us *userService) Login(user *domain.User) (string, error) {
	// user.Password = hashed(user.Password)

	// err := us.ur.GetUser(user)
	// if err != nil {
	// 	return "", err
	// }
	token := "test"
	return token, nil
}

// func hashed(p string) string {
// 	salt := base64.StdEncoding.EncodeToString([]byte(p))

// 	key := pbkdf2.Key([]byte(p), []byte(salt), 10, 10, sha256.New)

// 	return hex.EncodeToString(key[:])
// }
