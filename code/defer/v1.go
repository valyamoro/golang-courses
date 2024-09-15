package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	fmt.Println("Start work")
	panic("Произошла критическая ошибка!")
	fmt.Println("Конец работы!")
}
