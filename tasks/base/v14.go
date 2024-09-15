package main

import "fmt"

type Type interface{}

func getType(t Type) string {
	return fmt.Sprintf("Type is %T", t)
}

func TaskFourteen() {
	var t Type

	t = 0
	fmt.Println(getType(t))

	t = "abc"
	fmt.Println(getType(t))

	t = true
	fmt.Println(getType(t))

	t = make(chan int)
	fmt.Println(getType(t))
}
