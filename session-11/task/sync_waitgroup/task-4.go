package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}
	var sum1, sum2 int

	wg.Add(2)

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	go func() {

		value := 0
		for k, v := range nums {
			if k == len(nums)/2 {
				break
			}
			value = value + v

		}
		sum1 = value
		fmt.Println("Partial sum1", sum1)
		defer wg.Done()
	}()
	go func() {

		value := 0
		for i := len(nums); i > 4; i-- {

			value = value + i

		}
		sum2 = value
		fmt.Println("Partial sum2", sum2)
		defer wg.Done()

	}()
	wg.Wait()

	fmt.Println("Total sum :", sum1+sum2)

}
