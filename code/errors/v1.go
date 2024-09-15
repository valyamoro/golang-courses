package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("nonexistent_file.txt")
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			fmt.Println("Файл не существует.")
		} else {
			fmt.Println("Произошла ошибка:", err)
		}
	}
}
