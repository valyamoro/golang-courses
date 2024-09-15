package main

import (
	"fmt"
	"math"
	"math/big"
)

func TaskTwentyTwo() {
	firstValue := math.Pow(2, 20)
	secondValue := math.Pow(2, 30)
	firstStr := fmt.Sprintf("%f", firstValue)
	secondStr := fmt.Sprintf("%f", secondValue)
	a := new(big.Int)
	a.SetString(firstStr, 10)
	b := new(big.Int)
	b.SetString(secondStr, 10)

	div := new(big.Int)
	div.Div(a, b)
	fmt.Printf("Результат деления: %v\n", div)

	mul := new(big.Int)
	mul.Mul(a, b)
	fmt.Printf("Результат умножения: %v\n", mul)

	sum := new(big.Int)
	sum.Add(a, b)
	fmt.Printf("Результат сложения: %v\n", sum)

	dif := new(big.Int)
	dif.Sub(a, b)
	fmt.Printf("Результат вычитания %v\n", dif)
}
