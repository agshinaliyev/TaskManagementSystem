package main

import "fmt"

func main() {

	var x int = 5
	fmt.Println("Value of x before incrementing by value ", x)
	incrementByValue(x)
	fmt.Println("Value of x after incrementing by value ", x)

	fmt.Println("Value of x before incrementing by value ", x)

	ptr2 := incrementByPointer(&x)

	fmt.Println("Value of x after incrementing by pointer", *ptr2)

}

func incrementByValue(val int) int {

	z := val + 1
	return z
}

func incrementByPointer(ptr *int) *int {

	z := *ptr + 1
	return &z
}
