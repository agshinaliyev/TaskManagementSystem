package main

import (
	"errors"
)

func ConvertTemperature(value float64, scale string) (float64, error) {
	if scale == "CtoF" {
		return (value * 9 / 5) + 32, nil
	} else if scale == "FtoC" {
		return (value - 32) * 5 / 9, nil
	} else {
		return 0, errors.New("invalid scale: use 'CtoF' or 'FtoC'")
	}
}
