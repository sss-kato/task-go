package domain

import "github.com/cockroachdb/errors"

type ProjectIF interface {
	GetProjectName() string
}

type project struct {
	name string
}

func NewProject(pnm string) (ProjectIF, error) {

	err := validateProjectName(pnm)
	if err != nil {
		return nil, err
	}

	return &project{name: pnm}, nil
}

func validateProjectName(pnm string) error {

	length := len(pnm)
	if length < 30 {
		return errors.New(ErrorMsg11)
	}

	return nil
}

func (p *project) GetProjectName() string {

	return p.name
}
