package util

import (
	"github.com/cockroachdb/errors"
)

type AppError struct {
}

const blank = ""

func SetErrorInfo(err error) error {

	return errors.Wrap(err, blank)

	// fmt.Printf("%+v", test)
}

func SetErrorMsg(errorMsg string) error {
	return errors.New(errorMsg)
}

// func SetErrorInfo(err Error) {

// 	t := errors.New(err.Error())
// }
