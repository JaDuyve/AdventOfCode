package utils

import (
	"fmt"
	"os"
)

func WriteFile(path string, data []byte) error {
	f, err := os.Create(path)
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		return err
	}

	return nil
}
