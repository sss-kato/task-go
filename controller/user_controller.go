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

	user, err := domain.NewUser(name, pw, mail)

	if err != nil {
		return err
	}

	// if err := user.ValidateName(); err != nil {
	// 	return err
	// }

	// if err := user.ValidatePassword(); err != nil {
	// 	return err
	// }

	userRes, err := uc.us.Signup(user)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, userRes)

}
