package main

import (
	"fmt"
)

func main() {
	//task1
	fmt.Println("Main function started")

	//go goroutines.PrintNumbers()
	//time.Sleep(1 * time.Second)

	//task2
	//go goroutines.PrintLetters()
	//go goroutines.PrintNums()
	//time.Sleep(2 * time.Second)

	//task3
	//ch := make(chan int)
	//go func() {
	//	time.Sleep(500 * time.Millisecond)
	//	ch <- 42
	//
	//}()
	//value := <-ch
	//fmt.Println(value)

	//task4
	//ch := make(chan int)
	//go func() {
	//	for i := 1; i < 6; i++ {
	//		ch <- i
	//	}
	//	close(ch)
	//}()
	//for c := range ch {
	//	fmt.Println(c)
	//}

	//task5
	//ch := make(chan int, 3)
	//ch <- 10
	//ch <- 20
	//ch <- 30
	//fmt.Println("Sent values into buffered channel")
	//go func() {
	//	for c := range ch {
	//
	//		fmt.Println(c)
	//
	//	}
	//
	//}()

	//task6
	//ch := make(chan string)
	//ch <- "Hello"
	//?? just running the code below showed that all goroutines are asleep - deadlock!
	//am I missing something?
	fmt.Println("Main function ended")

}
