package main

import (
	"fmt"
	"sync"
)

func main() {
	var mutex sync.Mutex
	var counter int
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			counter++
			mutex.Unlock()

		}()

	}
	wg.Wait()
	fmt.Println("Counter:", counter)

}
