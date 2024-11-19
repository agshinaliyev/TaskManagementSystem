package main

import "testing"

func TestNewUserProfile(t *testing.T) {
	tests := []struct {
		name     string
		age      int
		expected UserProfile
	}{
		{"Agshin", 25, UserProfile{Name: "Agshin", Age: 25, IsMinor: false}},
		{"Ali", 15, UserProfile{Name: "Ali", Age: 15, IsMinor: true}},
	}

	for _, test := range tests {
		result := NewUserProfile(test.name, test.age)
		if result != test.expected {
			t.Errorf("For input (%s, %d), expected %+v but got %+v", test.name, test.age, test.expected, result)

		}

	}
}
