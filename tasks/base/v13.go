package main

import "fmt"

func TaskThirteen() {
	x := 1
	y := 2
	fmt.Printf("x=%d, y=%d\n", x, y)

	x, y = y, x
	fmt.Printf("x=%d, y=%d\n", x, y)
}
