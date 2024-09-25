package main

import (
	"fmt"
	"strconv"
)

func main() {

	var status string = "IN PROGRESS"
	var createdTasks int = 25
	var isCompleted bool = false
	const availableTasks int = 100
	const (
		low = iota
		medium
		high
	)
	formatBool := strconv.FormatBool(isCompleted)

	fmt.Println("Welcome to the Task Management System")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Task Management System")
	fmt.Printf("Current project status: %s \n", status)
	fmt.Printf("Tasks completed: %d out of %d \n", createdTasks, availableTasks)
	fmt.Printf("Task priorities: Low - %d,Medium - %d,High - %d \n", low, medium, high)
	fmt.Println("Is the project completed?", formatBool)

}
