package errs

import (
	"errors"
	"fmt"
)

var (
	InvalidPhoneNumber         = errors.New("Invalid phone number")
)

type CommonError struct {
	Err error
}

func (e CommonError) Error() error {
	return fmt.Errorf(e.Err.Error())
}