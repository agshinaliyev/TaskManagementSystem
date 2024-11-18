package main

import (
	"context"
	"fmt"
	"time"
)

func loop(ctx context.Context) {

	for i := 1; i < 11; i++ {

		select {
		case <-ctx.Done():
			fmt.Println("timeout")
			return
		default:
			fmt.Println(i)
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())

	go loop(ctx)
	time.Sleep(3 * time.Second)

	cancelFunc()
	time.Sleep(1 * time.Second)

}
