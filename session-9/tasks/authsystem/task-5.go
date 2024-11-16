package authsystem

import (
	"fmt"
)

var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}

func Authenticate(username, password string) (string, error) {
	passValue, ok := users[username]
	if !ok {

		err := fmt.Errorf("user not found %s %v \n", username, UserNotFound)
		fmt.Println(err)
		return "", err
	}
	if passValue != password {
		err := fmt.Errorf("error: invalid password for user: %s %v \n", username, InvalidPassword)
		fmt.Println(err)
		return "", err
	}
	return "", nil
}

type ErrorResponse struct {
	Code    int
	Message string
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

var (
	UserNotFound = ErrorResponse{
		Code:    404,
		Message: "User not found",
	}
	InvalidPassword = ErrorResponse{
		Code:    420,
		Message: "Invalid password",
	}
)
