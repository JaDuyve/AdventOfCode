package utils

import (
	"fmt"
	"os"
)

func WriteFile(path string, data string) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.WriteString(data)
	if err != nil {
		return err
	}

	return nil
}
