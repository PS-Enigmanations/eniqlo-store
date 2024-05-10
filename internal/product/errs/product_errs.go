package errs

import (
	"errors"
	"fmt"
)

var (
	ProductIsNotExists       = errors.New("Product is not exists")
	StockIsNotEnough         = errors.New("Stock is not enough")
)

type ProductError struct {
	Err error
}

func (e ProductError) Error() error {
	return fmt.Errorf(e.Err.Error())
}