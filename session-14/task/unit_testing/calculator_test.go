package main

import "testing"

func TestCalculator(t *testing.T) {
	tests := []struct {
		a, b      float64
		operation string
		expected  float64
		fails     bool
	}{
		{10, 5, "add", 15, false},
		{10, 5, "subtract", 5, false},
		{10, 5, "multiply", 50, false},
		{10, 5, "divide", 2, false},
		{10, 0, "divide", -1, true},
		{10, 0, "x", -1, true},
	}

	for _, test := range tests {
		result, err := Calculator(test.a, test.b, test.operation)
		if test.fails {
			if err == nil {
				t.Errorf("Expected an error for operation %s with inputs %f and %f", test.operation, test.a, test.b)
			}
		} else {
			if err != nil {
				t.Errorf("Did not expect an error but got: %v", err)
			}
			if result != test.expected {
				t.Errorf("For operation %s with inputs %f and %f, expected %f but got %f",
					test.operation, test.a, test.b, test.expected, result)
			}
		}
	}
}
