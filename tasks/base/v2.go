package main

import (
	"fmt"
	"sync"
	"time"
)

func double(i int) int {
	return i * i
}

func SquareSync(arr []int) []int {
	res := make([]int, 0)
	wg := &sync.WaitGroup{}

	for i := range arr {
		wg.Add(1)
		go func() {
			defer wg.Done()
			res = append(res, double(arr[i]))
		}()

		wg.Wait()
	}

	return res
}

func SquareSleep(arr []int) []int {
	res := make([]int, 0)

	for i := range arr {
		go func() {
			res = append(res, double(arr[i]))
		}()

		time.Sleep(1 * time.Second)
	}

	return res
}

func SquareCh(arr []int) []int {
	ch := make(chan int)
	res := make([]int, 0)

	go func([]int, chan int) {
		for i := range arr {
			ch <- double(arr[i])
		}

		close(ch)
	}(arr, ch)

	for {
		square, ok := <-ch
		if !ok {
			break
		}

		res = append(res, square)
	}

	return res
}

func TaskTwo(arr []int) {
	fmt.Printf("Wait Group: %v\n", SquareSync(arr))
	fmt.Printf("Sleep: %v\n", SquareSleep(arr))
	fmt.Printf("With Channel: %v\n", SquareCh(arr))
}
