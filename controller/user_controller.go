package controller

import (
	"net/http"
	"task-go/model"
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

	user := &model.User{}

	if err := c.Bind(user); err != nil {
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
