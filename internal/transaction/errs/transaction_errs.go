package errs

import (
	"errors"
	"fmt"
)

var (
	PaidIsNotEnough         = errors.New("Paid is not enough")
)

type TransactionError struct {
	Err error
}

func (e TransactionError) Error() error {
	return fmt.Errorf(e.Err.Error())
}