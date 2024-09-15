package main

import "fmt"

func newSet(slice []string) []string {
	resMap := make(map[string]bool)
	result := []string{}

	for _, key := range slice {
		resMap[key] = true
	}

	for key, _ := range resMap {
		result = append(result, key)
	}

	return result
}

func TaskTwelve() {
	slice := []string{"cat", "cat", "dog", "cat", "three"}
	set := newSet(slice)
	fmt.Println(set)
}
