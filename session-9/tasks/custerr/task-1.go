package custerr

import "errors"

func DivisionError() error {

	return errors.New("division by zero is not allowed")
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, DivisionError()
	}
	return a / b, nil

}
