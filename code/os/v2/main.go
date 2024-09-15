package main 

import (
	"fmt"
	"os"
)

func main() {
	err := os.Setenv("MY_ENV_VAR", "12345")
	if err != nil {
		fmt.Println("Ошибка при установке переменной окружения", err)
		return 
	}	

	fmt.Println("MY_ENV_VAR:", os.Getenv("MY_ENV_VAR"))

	err = os.Unsetenv("MY_ENV_VAR")
	if err != nil {
		fmt.Println("Error delete env", err)
		return 
	}

	fmt.Println("MY_ENV_VAR after remove", os.Getenv("MY_ENV_VAR"))
}
