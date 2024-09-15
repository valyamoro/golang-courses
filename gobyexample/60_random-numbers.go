package main

import (
	"fmt"
	"math/rand"
)

func main() {
	p := fmt.Print

	p(rand.Intn(100), ",")
	p(rand.Intn(100))

	fmt.Println()

	fmt.Print((rand.Float64()*5)+5, ",")
	fmt.Print((rand.Float64() * 5) + 5)
	fmt.Println()
}
