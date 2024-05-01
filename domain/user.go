package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/mail"

	"golang.org/x/crypto/pbkdf2"
)

type UserIF interface {
	ValidateName() error
	ValidatePassword() error
	ValidateMailAdress() error
	HashedPassword()

	GetName() string
	GetMailAdress() string
	GetPassWord() string
}

type user struct {
	name       string
	password   string
	mailadress string
}

func NewUser(nm string, pw string, ma string) UserIF {

	return &user{name: nm, password: pw, mailadress: ma}
}

func (u *user) ValidateName() error {

	length := len(u.name)

	if length > 15 {
		return errors.New("username must be fifteen characters or fewer.")
	} else if length < 5 {

		return errors.New("username must be at least five characters long.")
	}

	return nil
}

func (u *user) ValidatePassword() error {

	length := len(u.password)

	if length > 15 {

		return errors.New("password must be fifteen characters or fewer.")

	} else if length < 5 {

		return errors.New("password must be at least five characters long.")
	}
	return nil
}

func (u *user) ValidateMailAdress() error {

	lengh := len(u.mailadress)

	if lengh > 30 {

		return errors.New("mailadress must be thirty characters or fewer.")

	} else if lengh < 5 {

		return errors.New("mailadress must be at least five characters long.")
	}
	//  RFC 5322の観点でチェック
	_, err := mail.ParseAddress(u.mailadress)

	if err != nil {

		return errors.New("mailadress is invalid.")
	}

	return nil
}

func (u *user) HashedPassword() {
	salt := base64.StdEncoding.EncodeToString([]byte(u.password))
	key := pbkdf2.Key([]byte(u.password), []byte(salt), 1000, 5, sha256.New)
	u.password = hex.EncodeToString(key[:])
}

func (u *user) GetName() string {
	return u.name
}

func (u *user) GetMailAdress() string {
	return u.mailadress
}

func (u *user) GetPassWord() string {
	return u.password
}
