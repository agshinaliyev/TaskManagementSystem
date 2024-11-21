package main

import "fmt"

func main() {
	var x int = 10
	var ptr *int
	ptr = &x
	fmt.Printf("Value of x: %d \n", x)
	fmt.Printf("Address of x: %d \n", &x)
	fmt.Printf("Value at pointer: %d \n", *ptr)

}
