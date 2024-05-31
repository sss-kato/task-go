package controller

import (
	"net/http"
	"task-go/domain"
	"task-go/service"
	"task-go/util"
	"time"

	"github.com/labstack/echo/v4"
)

type UserControllerIF interface {
	Signup(e echo.Context) error
	Login(e echo.Context) error
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

	user, userErr := domain.NewUser(userReq.Name, userReq.Password, userReq.MailAdress)
	if userErr != nil {
		errMsg := &domain.Message{Message: userErr.Error()}
		util.WriteErrLog(userErr)
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	userRes, signupErr := uc.us.Signup(user)
	if signupErr != nil {
		errMsg := &domain.Message{Message: domain.ErrorMsg08}
		util.WriteErrLog(signupErr)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.JSON(http.StatusCreated, userRes)

}

func (uc *userController) Login(c echo.Context) error {

	userReq := &UserRequest{}
	bindErr := c.Bind(userReq)
	if bindErr != nil {
		util.WriteErrLog(bindErr)
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	user, userErr := domain.NewUser(userReq.Name, userReq.Password, userReq.MailAdress)
	if userErr != nil {
		errMsg := &domain.Message{Message: userErr.Error()}
		util.WriteErrLog(userErr)
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	token, loginErr := uc.us.Login(user)
	if loginErr != nil {
		errMsg := &domain.Message{Message: domain.ErrorMsg09}
		util.WriteErrLog(loginErr)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	cookie := new(http.Cookie)
	cookie.Name = "token"
	cookie.Value = token
	cookie.Expires = time.Now().Add(24 * time.Hour)
	cookie.Path = "/"

	// cookie.Secure = true
	cookie.HttpOnly = true
	cookie.SameSite = http.SameSiteNoneMode
	c.SetCookie(cookie)

	return c.NoContent(http.StatusOK)
}
