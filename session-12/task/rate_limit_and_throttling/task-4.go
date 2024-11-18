package main

import (
	"fmt"
	"time"
)

func worker(jobs chan int) {
	for job := range jobs {
		fmt.Printf("Job %d started \n", job)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	jobs := make(chan int, 2)
	ticker := time.NewTicker(500 * time.Millisecond)
	defer ticker.Stop()

	go worker(jobs)
	for i := 1; i < 6; i++ {
		<-ticker.C
		jobs <- i
	}
	close(jobs)
	time.Sleep(7 * time.Second)

}
