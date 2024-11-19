package main

import "errors"

type UserProfile struct {
	Name    string
	Age     int
	IsMinor bool
}

func NewUserProfile(name string, age int) UserProfile {

	user := UserProfile{
		Name:    name,
		Age:     age,
		IsMinor: false,
	}

	if user.Age < 18 {
		user.IsMinor = true
		errors.New("age must be greater than 18")
		return user
	}
	return user

}
