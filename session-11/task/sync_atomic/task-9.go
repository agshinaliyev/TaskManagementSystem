package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int32
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt32(&counter, 10)
		}()
	}

	wg.Wait()

	fmt.Println("Final atomic counter value:", counter)

}
