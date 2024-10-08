package service

import (
	"task-go/domain"
	"task-go/dto"
	"task-go/repository"
)

type ProjectServiceIF interface {
	CreateProject(pr domain.ProjectIF) error
}

type projectService struct {
	pr repository.ProjectRepositoryIF
}

func NewProjectService(pr repository.ProjectRepositoryIF) ProjectServiceIF {
	return &projectService{pr}
}

func (ps *projectService) CreateProject(pr domain.ProjectIF) error {
	pd := &dto.ProjectDto{Name: pr.GetProjectName(), UserID: uint(pr.GetUserID())}
	err := ps.pr.CreateProject(pd)
	if err != nil {
		return err
	}
	return nil
}
