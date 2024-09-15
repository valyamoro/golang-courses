package main

import "fmt"

func TaskSeventeen(arr []int, target int) int {
	idx := -1
	start := 0
	end := len(arr) - 1

	for start <= end {
		mid := (start + end) / 2
		if arr[mid] == target {
			idx = mid
			break
		} else if arr[mid] < target {
			start = mid + 1
		} else if arr[mid] > target {
			end = mid - 1
		}
	}

	fmt.Println(idx)

	return idx
}
