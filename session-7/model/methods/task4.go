package main

import "fmt"

type Student struct {
	Name   string
	Grades []int
}

func (s *Student) Average() float64 {

	sum := 0
	for i := 0; i <= len(s.Grades)-1; i++ {

		sum = sum + s.Grades[i]
	}
	return float64(sum) / float64(len(s.Grades))

}
func main() {

	student1 := Student{
		Name:   "Ali",
		Grades: []int{72, 86, 98, 56},
	}

	student2 := Student{
		Name:   "Vali",
		Grades: []int{55, 77, 65, 99},
	}
	average1 := student1.Average()
	average2 := student2.Average()

	fmt.Printf("Student1 : %s , Average grade : %f \n", student1.Name, average1)
	fmt.Printf("Student2 : %s , Average grade : %f \n", student2.Name, average2)

	if average1 > average2 {
		fmt.Println("Ali has higher grade")

	} else if average1 < average2 {

		fmt.Println("Vali has higher grade")

	} else {
		fmt.Println("Grades are equal")
	}
}
