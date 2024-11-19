package main

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

func isValidInput(s string) bool {
	if s == "" {
		return false
	}
	for _, char := range s {
		if !unicode.IsLetter(char) && !unicode.IsNumber(char) {
			return false
		}
	}
	return true
}

func palindrome(s string) bool {
	lower := strings.ToLower(s)
	length := len(lower)

	for i := 0; i < length/2; i++ {
		if lower[i] != lower[length-1-i] {
			return false
		}
	}
	return true
}

func ReverseString(s string) string {
	var reversed string
	if !isValidInput(s) || palindrome(s) {
		err := errors.New("unexpected")
		fmt.Println(err)

	} else {
		var reversed string
		for i := len(s) - 1; i >= 0; i-- {
			reversed += string(s[i])
		}
	}

	return reversed
}
func main() {
	fmt.Println(ReverseString("level"))
}
