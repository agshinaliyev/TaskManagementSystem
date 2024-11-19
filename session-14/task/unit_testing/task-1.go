package main

import (
	"errors"
)

func Calculator(a, b float64, operation string) (float64, error) {

	switch operation {
	case "add":
		return a + b, nil
	case "subtract":
		return a - b, nil
	case "multiply":
		return a * b, nil
	case "divide":
		if b == 0 {
			return -1, errors.New("cannot divide by zero")
		}
		return a / b, nil
	default:
		return -1, errors.New("invalid operation")
	}

}
