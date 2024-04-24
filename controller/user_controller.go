package controller

import (
	"net/http"
	"task-go/domain"
	"task-go/service"

	"github.com/labstack/echo/v4"
)

type UserControllerIF interface {
	Signup(e echo.Context) error
}

type userController struct {
	us service.UserServiceIF
}

func NewUserController(us service.UserServiceIF) UserControllerIF {
	return &userController{us}
}

type UserRequest struct {
	Name       string `json:"name"`
	Password   string `json:"password"`
	MailAdress string `json:"mailadress"`
}

func (uc *userController) Signup(c echo.Context) error {

	// userTest := struct {
	// 	Name       string
	// 	Password   string
	// 	MailAdress string
	// }{}

	userReq := new(UserRequest)
	if err := c.Bind(userReq); err != nil {
		return err
	}
	user := domain.NewUser(userReq.Name, userReq.Password, userReq.Password)

	nameErr := user.ValidateName()
	if nameErr != nil {
		return nameErr
	}

	pwErr := user.ValidatePassword()
	if pwErr != nil {
		return pwErr
	}

	mailErr := user.ValidateMailAdress()
	if mailErr != nil {
		return mailErr
	}

	userRes, err := uc.us.Signup(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userRes)

}
