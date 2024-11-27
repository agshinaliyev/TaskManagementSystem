package wrapper

import (
	"fmt"
	"os"
)

func OpenFile(filename string) error {
	file, err := os.Open(filename)

	if err != nil {
		return fmt.Errorf("wrapped error : file not found %v ", err)
	}
	fmt.Println(file)
	err = file.Close()

	return nil
}
