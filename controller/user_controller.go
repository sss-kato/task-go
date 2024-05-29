package controller

import (
	"html"
	"net/http"
	"task-go/domain"
	"task-go/service"
	"task-go/util"

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
	Name       string
	Password   string
	MailAdress string
}

func (uc *userController) Signup(c echo.Context) error {

	userReq := &UserRequest{}

	if bindErr := c.Bind(userReq); bindErr != nil {
		util.WriteErrLog(bindErr)
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	user, useErr := domain.NewUser(html.EscapeString(userReq.Name), userReq.Password, html.EscapeString(userReq.MailAdress))

	if useErr != nil {

		errMsg := &domain.Message{Message: useErr.Error()}

		util.WriteErrLog(useErr)

		return c.JSON(http.StatusBadRequest, errMsg)
	}

	userRes, signupErr := uc.us.Signup(user)
	if signupErr != nil {
		errMsg := &domain.Message{Message: "signup failed."}

		util.WriteErrLog(signupErr)

		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, userRes)

}
