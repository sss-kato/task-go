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
	Name       string
	Password   string
	MailAdress string
}

func (uc *userController) Signup(c echo.Context) error {

	userReq := &UserRequest{}
	if err := c.Bind(userReq); err != nil {
		return err
	}

	user := domain.NewUser(userReq.Name, userReq.Password, userReq.MailAdress)

	nameErr := user.ValidateName()
	if nameErr != nil {
		errMsg := &domain.Message{Message: nameErr.Error()}
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	pwErr := user.ValidatePassword()
	if pwErr != nil {
		errMsg := &domain.Message{Message: pwErr.Error()}
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	mailErr := user.ValidateMailAdress()
	if mailErr != nil {
		errMsg := &domain.Message{Message: mailErr.Error()}
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	userRes, err := uc.us.Signup(user)
	if err != nil {
		errMsg := &domain.Message{Message: "signup failed."}
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, userRes)

}
