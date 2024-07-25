package service

import (
	"task-go/dto"
	"task-go/repository"
)

type ProjectServiceIF interface {
	CreateProject(pd dto.ProjectDto) error
}

type projectService struct {
	pr repository.ProjectRepositoryIF
}
