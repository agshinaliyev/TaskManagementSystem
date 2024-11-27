package wrapper

import (
	"fmt"
	"os"
)

func FileOpener(filename string) error {

	file, err := os.Open(filename)
	err = ReadFile(file)
	if err != nil {
		err := fmt.Errorf("error : failed to open file: %v", err)
		fmt.Println(err)
	}
	return nil

}

func ReadFile(f *os.File) error {

	bytes := make([]byte, 1024)

	file, err := f.Read(bytes)

	if err != nil {
		return fmt.Errorf("failed to read file  %v", err)
	}
	fmt.Println(file)
	err = f.Close()
	return nil
}
