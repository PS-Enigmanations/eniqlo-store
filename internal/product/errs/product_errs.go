package errs

import (
	"errors"
	"fmt"
)

var (
	ProductIsNotExists         = errors.New("Product is not exists")
)

type ProductError struct {
	Err error
}

func (e ProductError) Error() error {
	return fmt.Errorf(e.Err.Error())
}