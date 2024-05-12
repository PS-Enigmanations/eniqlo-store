package errs

import (
	"errors"
	"fmt"
)

var (
	UserExist             = errors.New("User already exists")
	ErrInvalidPhoneNumber = errors.New("invalid phone number")
)

type UserError struct {
	Err error
}

func (e UserError) Error() error {
	return fmt.Errorf(e.Err.Error())
}
