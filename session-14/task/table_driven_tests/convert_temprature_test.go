package main

import (
	"fmt"
	"testing"
)

func TestConvertTemperature(t *testing.T) {
	tests := []struct {
		inputValue float64
		scale      string
		expected   float64
		shouldErr  bool
	}{
		{0, "CtoF", 32, false},    // 0°C to 32°F
		{100, "CtoF", 212, false}, // 100°C to 212°F
		{-40, "CtoF", -40, false}, // -40°C to -40°F
		{32, "FtoC", 0, false},    // 32°F to 0°C
		{212, "FtoC", 100, false}, // 212°F to 100°C
		{-40, "FtoC", -40, false}, // -40°F to -40°C
		{100, "Invalid", 0, true}, // Invalid scale
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("%v_%s", test.inputValue, test.scale), func(t *testing.T) {
			result, err := ConvertTemperature(test.inputValue, test.scale)
			if test.shouldErr && err == nil {
				t.Errorf("expected error but got none")
			}
			if !test.shouldErr && err != nil {
				t.Errorf("did not expect error but got: %v", err)
			}
			if !test.shouldErr && result != test.expected {
				t.Errorf("expected %v but got %v", test.expected, result)
			}
		})
	}
}
