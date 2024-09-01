package controller

import (
	"net/http"
	"task-go/domain"
	"task-go/service"
	"task-go/util"

	"github.com/labstack/echo/v4"
)

type ProjectControllerIF interface {
	CreateProject(c echo.Context) error
}

type projectController struct {
	ps service.ProjectServiceIF
}

func NewProjectController(ps service.ProjectServiceIF) ProjectControllerIF {
	return &projectController{ps}
}

type ProjectRequest struct {
	Name   string
	UserID int
}

func (pc *projectController) CreateProject(c echo.Context) error {

	projectReq := &ProjectRequest{}

	if bindErr := c.Bind(projectReq); bindErr != nil {
		util.WriteErrLog(bindErr)
		return c.JSON(http.StatusBadRequest, bindErr.Error())
	}

	project, projectErr := domain.NewProject(projectReq.Name, projectReq.UserID)
	if projectErr != nil {
		errMsg := &domain.Message{Message: projectErr.Error()}
		util.WriteErrLog(projectErr)
		return c.JSON(http.StatusBadRequest, errMsg)
	}

	createPjErr := pc.ps.CreateProject(project)
	if createPjErr != nil {
		errMsg := &domain.Message{Message: domain.ErrorMsg13}
		util.WriteErrLog(createPjErr)
		return c.JSON(http.StatusInternalServerError, errMsg)
	}

	return c.NoContent(http.StatusOK)
}
