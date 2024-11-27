package goroutines

import (
	"fmt"
	"time"
)

func PrintLetters() {

	arr := [5]string{"A", "B", "C", "D", "E"}

	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr[i])
		time.Sleep(200 * time.Millisecond)
	}

}

func PrintNums() {

	for i := 1; i < 6; i++ {

		fmt.Println(i)
		time.Sleep(300 * time.Millisecond)
	}
}
