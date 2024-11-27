package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := sync.WaitGroup{}
	wg.Add(3)

	go func() {

		fmt.Println("Start of func1")
		time.Sleep(1 * time.Second)
		fmt.Println("Func 1 finished")
		defer wg.Done()

	}()
	go func() {
		fmt.Println("Start of func2")
		time.Sleep(1 * time.Second)
		fmt.Println("Func 2 finished")
		defer wg.Done()

	}()
	go func() {

		fmt.Println("Start of func3")
		time.Sleep(1 * time.Second)
		fmt.Println("Func 3 finished")
		defer wg.Done()

	}()
	wg.Wait()
	fmt.Println("Main ended")

}
