package domain

import (
	"reflect"

	"github.com/cockroachdb/errors"
)

type ProjectIF interface {
	GetProjectName() string
	GetUserID() int
}

type project struct {
	name string
	uid  int
}

func NewProject(pnm string, uid int) (ProjectIF, error) {

	err := validateProjectName(pnm)
	if err != nil {
		return nil, err
	}

	return &project{name: pnm, uid: uid}, nil
}

func validateProjectName(pnm string) error {

	length := len(pnm)
	if length < 30 {
		return errors.New(ErrorMsg11)
	}

	return nil
}

func validateUserID(uid interface{}) error {

	kind := reflect.TypeOf(uid).Kind()

	if kind != reflect.Int {

		return errors.New(ErrorMsg12)
	}

	return nil
}

func (p *project) GetProjectName() string {

	return p.name
}

func (p *project) GetUserID() int {

	return p.uid
}
