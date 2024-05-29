package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"net/mail"

	"github.com/cockroachdb/errors"
	"golang.org/x/crypto/pbkdf2"
)

type UserIF interface {
	GetName() string
	GetMailAdress() string
	GetPassWord() string
}

type user struct {
	name       string
	password   string
	mailadress string
}

func NewUser(nm string, pw string, ma string) (UserIF, error) {

	nmErr := validateName(nm)
	if nmErr != nil {
		return nil, nmErr
	}

	maErr := validateMailAdress(ma)
	if maErr != nil {
		return nil, maErr
	}

	pwErr := validatePassword(pw)
	if pwErr != nil {
		return nil, pwErr
	}
	hashedPw := hashedPassword(pw)

	return &user{name: nm, password: hashedPw, mailadress: ma}, nil
}

func validateName(nm string) error {

	length := len(nm)

	if length > 15 {

		return errors.New(ErrorMsg01)

	} else if length < 5 {

		return errors.New(ErrorMsg02)
	}

	return nil
}

func validatePassword(pw string) error {

	length := len(pw)

	if length > 15 {

		return errors.New(ErrorMsg03)

	} else if length < 5 {

		return errors.New(ErrorMsg04)
	}
	return nil
}

func validateMailAdress(ma string) error {

	length := len(ma)

	if length > 30 {

		return errors.New(ErrorMsg05)

	} else if length < 5 {

		return errors.New(ErrorMsg06)
	}

	//  RFC 5322の観点でチェック
	_, err := mail.ParseAddress(ma)

	if err != nil {

		return errors.New(ErrorMsg07)
	}

	return nil

}

func hashedPassword(pw string) string {
	salt := base64.StdEncoding.EncodeToString([]byte(pw))
	key := pbkdf2.Key([]byte(pw), []byte(salt), 1000, 5, sha256.New)
	return hex.EncodeToString(key[:])
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
