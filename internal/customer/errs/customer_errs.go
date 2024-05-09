package errs

import (
	"errors"
	"fmt"
)

var (
	CustomerExist         = errors.New("Phone number already exists")
)

type CustomerError struct {
	Err error
}

func (e CustomerError) Error() error {
	return fmt.Errorf(e.Err.Error())
}