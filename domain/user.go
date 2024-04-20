package domain

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"net/mail"
	"time"

	"golang.org/x/crypto/pbkdf2"
)

type UserIF interface {
}

type User struct {
	ID         uint
	Name       string
	Password   string
	Mailadress string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

func NewUser(name string, pw string, mail string) (UserIF, error) {

	nameErr := validateName(name)
	if nameErr != nil {
		return nil, nameErr
	}

	pwErr := validatePassword(pw)
	if pwErr != nil {
		return nil, pwErr
	}

	mailErr := validateMailAdress(mail)
	if mailErr != nil {
		return nil, mailErr
	}
	hashedPw := hashedPassword(pw)

	return &User{Name: name, Password: hashedPw, Mailadress: mail}, nil
}

func validateName(name string) error {

	length := len(name)

	if length > 15 {
		return errors.New("username must be fifteen characters or fewer.")
	} else if length < 5 {

		return errors.New("username must be at least five characters long.")
	}

	return nil
}

func validatePassword(pw string) error {

	length := len(pw)

	if length > 15 {

		return errors.New("password must be fifteen characters or fewer.")

	} else if length < 5 {

		return errors.New("password must be at least five characters long.")
	}
	return nil
}

func validateMailAdress(ma string) error {

	lengh := len(ma)

	if lengh > 30 {

		return errors.New("password must be thirty characters or fewer.")

	} else if lengh < 5 {

		return errors.New("password must be at least five characters long.")
	}
	//  RFC 5322の観点でチェック
	_, err := mail.ParseAddress(ma)

	if err != nil {
		return err
	}

	return nil
}

func hashedPassword(pw string) string {
	salt := base64.StdEncoding.EncodeToString([]byte(pw))
	key := pbkdf2.Key([]byte(pw), []byte(salt), 1000, 5, sha256.New)
	return hex.EncodeToString(key[:])
}
