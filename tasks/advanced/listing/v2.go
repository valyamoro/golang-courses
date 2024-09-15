package main

import "fmt"

func test() (x int) {
	defer func() {
		x++
	}()

	x = 1
	return
}

func anotherTest() int {
	var x int
	defer func() {
		x++
	}()

	return x + 1
}

func main() {
	fmt.Println(test())
	fmt.Println(anotherTest())
}
