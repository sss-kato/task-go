package service

import (
	"task-go/domain"
	"task-go/dto"
	"task-go/repository"
	"time"

	"github.com/cockroachdb/errors"
	"github.com/golang-jwt/jwt"
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

	ud := &dto.UserDto{Name: user.GetName(), Mailadress: user.GetMailAdress(), Password: user.GetPassWord()}

	err := us.ur.RegistUser(ud)
	if err != nil {
		return domain.UserResponse{}, err
	}

	resUser := domain.UserResponse{ID: ud.ID, Name: ud.Name}
	return resUser, nil
}

func (us *userService) Login(user domain.UserIF) (string, error) {

	ud := &dto.UserDto{Name: user.GetName(), Password: user.GetPassWord()}
	userCnt, userId, userErr := us.ur.GetUser(ud)

	if userCnt == 0 {
		return "", errors.New(domain.ErrorMsg09)
	}

	if userErr != nil {
		return "", userErr
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})

	const key = "4fe269f707e7ffdf0c772994046a4242449de81d1acef7bc2dc6588099fabec2"
	tokenString, tokenErr := token.SignedString([]byte(key))
	if tokenErr != nil {
		return "", tokenErr
	}

	return tokenString, nil
}
