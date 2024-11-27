package main

import (
	"fmt"
	"session-9/tasks/authsystem"
	"session-9/tasks/custerr"
	"session-9/tasks/wrapper"
)

func main() {
	//task1
	divide, err := custerr.Divide(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result :", divide)
	}

	fmt.Println("-----------------------")

	//task2
	divide2, err := custerr.Divides(5, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result :", divide2)
	}
	fmt.Println("-----------------------")

	//task3
	err = wrapper.OpenFile("nonexistent.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("-----------------------")

	//task4
	err = wrapper.FileOpener("nonexistent.txt")
	// ask: couldn't specify filename in console
	fmt.Println("-----------------------")

	//task5
	authenticate, err := authsystem.Authenticate("user1", "password1")

	if err == nil {
		fmt.Printf("Login succesfull %v ", authenticate)
	}

}
