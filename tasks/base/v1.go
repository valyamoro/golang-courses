package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

type Action struct {
	Human
}

func main() {
	action := Action{Human{
		Age:  13,
		Name: "John",
	}}

	fmt.Println(action)
}
