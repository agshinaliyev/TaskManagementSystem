package main

import (
	"fmt"
	"time"
)

func printer() {
	fmt.Println("task executed")
}

func main() {
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	for i := 0; i < 5; i++ {
		printer()
		<-ticker.C
	}
	fmt.Println("ticker stopped")
}
