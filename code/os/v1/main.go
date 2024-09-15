package main 

import (
	"fmt"
	"os"
)

func main() {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Ошибка:", err)
		return 
	}

	fmt.Println("Текущий рабочий каталог: ", dir)

	path := "example.txt"
	if _, err := os.Stat(path); os.IsNotExist(err) {
		fmt.Printf("Файл или каталог %s does not exists", path)
	} else {
		fmt.Printf("File or directory doesn`t exists %s", path)
	}
}
