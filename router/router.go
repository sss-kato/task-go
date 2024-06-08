package router

import (
	"task-go/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.UserControllerIF) *echo.Echo {

	e := echo.New()
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	return e
}
