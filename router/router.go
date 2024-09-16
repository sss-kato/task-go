package router

import (
	"task-go/controller"

	"github.com/labstack/echo/v4"
)

func NewRouter(uc controller.UserControllerIF, pc controller.ProjectControllerIF) *echo.Echo {

	e := echo.New()
	e.POST("/signup", uc.Signup)
	e.POST("/login", uc.Login)
	e.POST("/logout", uc.Logout)

	p := e.Group("/project")
	p.POST("/create", pc.CreateProject)
	return e
}
