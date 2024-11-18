package main

import (
	"context"
	"fmt"
	"time"
)

func simulateTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed.")
	case <-ctx.Done():
		fmt.Println("Task cancelled due to timeout ")
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	simulateTask(ctx)

}
