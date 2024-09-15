package main

import "fmt"

func subset(val float32) int {
	return int(val) / 10 * 10
}

func TaskTen() {
	temperature := []float32{-25.4, -27.0, 13.0, 19.0, 23.3, 21.3}
	result := make(map[int][]float32)

	for _, val := range temperature {
		key := subset(val)
		result[key] = append(result[key], val)
	}

	fmt.Println(result)
}
