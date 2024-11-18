package main

import (
	"fmt"
	"sync"
)

func main() {

	grades := make(map[string]int)
	var mutex sync.Mutex
	var wg sync.WaitGroup

	update := func(name string, grade int) {
		defer wg.Done()
		mutex.Lock()
		grades[name] = grade
		mutex.Unlock()
	}
	wg.Add(3)
	go update("Garay", 90)
	go update("Ali", 85)
	go update("Medina", 78)
	wg.Wait()

	fmt.Println("Final grades :", grades)
}
