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

func (uc *userController) Signup(c echo.Context) error {

	// user := &model.User{}

	name := c.FormValue("name")
	pw := c.FormValue("password")
	mail := c.FormValue("mailadress")

	user := domain.NewUser(name, pw, mail)

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
