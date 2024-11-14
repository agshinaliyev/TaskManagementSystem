package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
}
type Employee struct {
	Person
	EmployeeId int
	Position   string
}

func main() {
	employee := Employee{
		Person: Person{"Ali",
			"Aliyev",
		},
		EmployeeId: 12345,
		Position:   "Software Engineer",
	}
	fmt.Printf("Name : %s %s \nEmployeeId: %d \nPosition: %s",
		employee.FirstName, employee.LastName, employee.EmployeeId, employee.Position)
}
