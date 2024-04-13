package model

import (
	"errors"
	"net/mail"
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
		return errors.New("username must be fifteen characters or fewer.")
	} else if length < 5 {

		return errors.New("username must be at least five characters long.")
	}

	return nil
}

func (u *User) ValidatePassword() error {

	length := len(u.Password)

	if length > 15 {

		return errors.New("password must be fifteen characters or fewer.")

	} else if length < 5 {

		return errors.New("password must be at least five characters long.")
	}
	return nil
}

func (u *User) ValidateMailAdress() error {

	lengh := len(u.Mailadress)

	if lengh > 30 {

		return errors.New("password must be thirty characters or fewer.")

	} else if lengh < 5 {

		return errors.New("password must be at least five characters long.")
	}

	_, err := mail.ParseAddress(u.Mailadress)

	if err != nil {
		return err
	}

	return nil
}
