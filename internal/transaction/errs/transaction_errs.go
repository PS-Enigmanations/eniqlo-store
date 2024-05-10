package errs

import (
	"errors"
	"fmt"
)

var (
	PaidIsNotEnough         = errors.New("Paid is not enough")
	ChangeIsNotRight		= errors.New("Change is not right")
)

type TransactionError struct {
	Err error
}

func (e TransactionError) Error() error {
	return fmt.Errorf(e.Err.Error())
}