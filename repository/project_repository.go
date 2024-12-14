package repository

import (
	"log"
	"task-go/dto"

	"github.com/cockroachdb/errors"

	"gorm.io/gorm"
)

// import (
// 	"task-go/dto"

// 	"gorm.io/gorm"
// )

type ProjectRepositoryIF interface {
	GetProjects(ud dto.UserDto) ([]dto.ProjectDto, error)
	CreateProject(pd *dto.ProjectDto) error
	// 	DeleteProject(pd dto.ProjectDto) error
}

type projectRepository struct {
	db *gorm.DB
}

func NewProjectRepository(db *gorm.DB) ProjectRepositoryIF {
	return &projectRepository{db}
}

// func (pr *projectRepository) GetProjects(pd dto.ProjectDto, ud dto.UserDto) ([]dto.ProjectDto, error) {

// }

func (pr *projectRepository) CreateProject(pd *dto.ProjectDto) error {

	tx := pr.db.Begin()
	projectErr := tx.Create(pd).Error
	if projectErr != nil {
		tx.Rollback()
		return errors.New(projectErr.Error())
	}

	pumd := &dto.Project_User_MappingDto{ProjectID: pd.ID, UserID: pd.UserID}
	mappingErr := tx.Create(pumd).Error
	if mappingErr != nil {
		tx.Rollback()
		return errors.New(mappingErr.Error())
	}

	tx.Commit()

	return nil
}

func (pr *projectRepository) GetProjects(ud dto.UserDto) ([]dto.ProjectDto, error) {

	var projects []dto.ProjectDto

	if err := pr.db.Table("projects").
		Joins("JOIN project_user_mappings ON projects.id = project_user_mappings.project_id AND  project_user_mappings.user_id = ?", ud.ID).
		Find(&projects).Error; err != nil {

		log.Fatal(err)

		return nil, errors.New(err.Error())

	}

	return projects, nil
}
