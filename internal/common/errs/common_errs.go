package errs

import (
	"errors"
	"fmt"
)

var (
	InvalidPhoneNumber = errors.New("Invalid phone number")
	ErrProductNotFound = errors.New("product not found")
	ErrImageUrlInvalid = errors.New("invalid image url")
)

type CommonError struct {
	Err error
}

func (e CommonError) Error() error {
	return fmt.Errorf(e.Err.Error())
}
