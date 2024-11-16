package custerr

import "fmt"

type ErrorDivision struct {
	A       int
	B       int
	Message string
}

func (err ErrorDivision) Error() string {

	return fmt.Sprintf("Error: Cannot divide %d by %d", err.A, err.B)
}

func Divides(a, b int) (float64, error) {

	if b == 0 {

		return 0, ErrorDivision{a, b, "Cannot divide "}

	}
	return float64(a / b), nil

}
