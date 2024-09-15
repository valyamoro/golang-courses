package main 

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error on create file", err)
		return 
	}

	fmt.Println("File created", file.Name())
	file.Close()

	err = os.Remove("example.txt")
	if err != nil {
		fmt.Println("Error on remove file", err)
		return 
	}

	fmt.Println("File successful removed")
}
