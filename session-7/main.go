package main

import (
	"errors"
	"fmt"
	"strings"
)

const (
	low = iota
	medium
	high
)

func main() {
	fmt.Println("Welcome to the Task Management System")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Task Management System")

	const totalTasks = 100
	done := 101
	remains := totalTasks - done
	if done > 90 && done < 100 {
		fmt.Printf("Tasks remaining %d out of %d \n", remains, totalTasks)

	}
	status := "in progress"
	if done == 100 {
		fmt.Println("Project is done")

	} else {

		fmt.Printf("Current project status: %s \n", strings.ToUpper(status))
	}
	switch {
	case done > 0 && done <= 30:
		fmt.Println("Project is in the starting phase")
	case done > 30 && done <= 60:
		fmt.Println("Project is in the midway")
	case done > 60 && done <= 100:
		fmt.Println("Project is in the final phase")
	}

	fmt.Printf("Task priorities: Low - %d,Medium - %d,High - %d \n", low, medium, high)
	for i := 1; i <= totalTasks-done; i++ {
		fmt.Println("Task -", i)
	}

	if done > totalTasks {
		errName := errors.New("task count exceeds the total task number")
		fmt.Println("Error:", errName)
		done = totalTasks
		fmt.Println("Completed tasks reset to maximum")
	}
}
