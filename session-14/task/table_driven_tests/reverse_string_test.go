package main

import (
	"fmt"
	"testing"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"level", "unexpected"}, // Palindrome, so expect "unexpected"
		{"", "unexpected"},      // Empty string, expect "unexpected"
		{"abc123", "321cba"},    // Alphanumeric string
	}

	for _, test := range tests {
		t.Run(fmt.Sprintf("Input: %s", test.input), func(t *testing.T) {
			result := ReverseString(test.input)
			if result != test.expected {
				t.Errorf("For input %s, expected %s, but got %s", test.input, test.expected, result)
			}
		})
	}
}
