package goroutines

import (
	"fmt"
	"time"
)

func PrintNumbers() {

	for i := 0; i < 6; i++ {
		fmt.Println(i)
		time.Sleep(100 * time.Millisecond)
	}

}
