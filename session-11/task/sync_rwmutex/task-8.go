package main

import (
	"fmt"
	"sync"
	"time"
)

type counter struct {
	counter int
	mutex   sync.RWMutex
}

func (c *counter) reader() int {
	c.mutex.RLock()
	fmt.Println("Reader accessed", c.counter)
	defer c.mutex.RUnlock()
	return c.counter
}

func (c *counter) writer() int {
	c.mutex.Lock()
	c.counter++
	fmt.Print(c.counter)
	defer c.mutex.Unlock()
	return c.counter
}

func main() {
	counter := counter{}
	var wg sync.WaitGroup

	wg.Add(4)
	go func() {
		defer wg.Done()
		counter.reader()
		time.Sleep(100 * time.Millisecond) // Simulate work
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Writer accessed ", counter.writer())
		time.Sleep(150 * time.Millisecond) // Simulate work
	}()
	go func() {
		defer wg.Done()
		counter.reader()
	}()
	go func() {
		defer wg.Done()
		fmt.Println("Writer accessed ", counter.writer())
	}()

	wg.Wait()
}
