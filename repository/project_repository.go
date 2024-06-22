package repository

import (
	"task-go/dto"

	"gorm.io/gorm"
)

type ProjectRepositoryIF interface {
	GetProjects(pd *dto.ProjectDto) error
	CreateProject(pd *dto.ProjectDto) error
	DeleteProject(pd *dto.ProjectDto) error
}

type projectRepository struct {
	db *gorm.DB
}

// func NewProjectRepository(db *gorm.DB) ProjectRepositoryIF {
// 	return &projectRepository{db}
// }

// func (pr *projectRepository) GetProjects(pd *dto.ProjectDto) error {

// }
