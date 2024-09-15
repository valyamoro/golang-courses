package main

import "fmt"

func IndexOf[T comparable](slice []T, item T) int {
	for i, v := range slice {
		if v == item {
			return i
		}
	}

	return -1
}

func main() {
	nums := []int{10, 20, 30, 40}
	strs := []string{"apple", "banana"}

	fmt.Println(IndexOf(nums, 30))
	fmt.Println(IndexOf(strs, "banana"))
}
