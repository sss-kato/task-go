package model

import (
	"errors"
	"time"
)

type User struct {
	ID         uint
	Name       string
	Password   string
	Mailadress string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func (u *User) ValidateName() error {
	length := len(u.Name)

	if length > 15 {
		return errors.New("ee")
	}

	return nil
}
