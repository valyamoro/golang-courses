package main

import "fmt"

func TaskEleven() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{2, 3, 4, 1, 5, 6}
	result := []int{}

	for _, val1 := range set1 {
		for _, val2 := range set2 {
			if val1 == val2 {
				result = append(result, val1)
			}
		}
	}

	fmt.Println(result)
}
