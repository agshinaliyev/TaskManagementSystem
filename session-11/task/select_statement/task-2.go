package main

//breakpoint

import (
	"fmt"
	"time"
)

func waitForChannel(ch chan int) string {
	select {
	case data := <-ch:
		return fmt.Sprintf("Received value: %d", data)
	case <-time.After(3 * time.Second):
		return "Timeout: no data received"
	}
}

func main() {
	ch := make(chan int)

	result := waitForChannel(ch)
	fmt.Println(result)
}
