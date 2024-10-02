package main

import "fmt"

const total = 100

func main() {
	fmt.Println("Welcome to the Task Management System")
	fmt.Println("Project start date is: 2024-09-18 00:00:00")
	fmt.Println("Task Management System")
	defer func() {
		fmt.Println(check(false))
		if err := recover(); err != nil {
			fmt.Println("Recovered from panic \n", err)
		}
	}()

	fmt.Println(doneCount(total, 5))
	taskListing(95)

	fmt.Println(predicate(92))
	fmt.Println(remainingTasks(15))
	printTaskDetails("task 1")
}

func check(status bool) string {

	if status == false {
		return "IN PROGRESS"
	}
	return "All the tasks completed"
}

func doneCount(total, remain int) int {
	return total - remain

}

func remainCount(total, done int) int {
	return total - done

}
func taskListing(done int) {
	if done > total {

		panic("Completed tasks exceeds the total count")

	}
	for i := 1; i <= total-done; i++ {
		fmt.Println("Task - ", i)
	}

}
func predicate(taskCount int) (int, bool) {

	if taskCount > 90 {
		return remainCount(total, 5), true

	}
	return -1, false
}

func remainingTasks(count int) int {
	if count < 1 {
		return 0

	}
	fmt.Println("Tasks remaining : ", count)

	return remainingTasks(count - 1)

}
func printTaskDetails(taskNames ...string) {
	fmt.Println("Task Details:")
	for _, taskName := range taskNames {
		fmt.Println("- " + taskName)
	}
}
