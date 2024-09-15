package main

import "fmt"

func TaskEight() {
	var number int64 = 100
	var numberOfBite int64 = 5

	switchToOne := false

	switch switchToOne {
	case true:
		number |= 1 << numberOfBite
	default:
		number &^= 1 << numberOfBite
	}

	fmt.Println("Result is -", number)
}
