package main

import (
	"fmt"
	"time"
)

func worker(workerName int, in <-chan int) {
	for {
		num := <-in
		fmt.Printf("Goroutine #%d: value: %ds\n", workerName, num)
	}
}

func TaskFour() {
	var N int
	fmt.Println("Количество гоурутин:")
	fmt.Scanf("%d\n", &N)

	workerInput := make(chan int)

	for i := 0; i < N; i++ {
		go worker(i, workerInput)
	}

	for {
		workerInput <- time.Now().Second()
		time.Sleep(time.Second)
	}
}

func main() {
	TaskFour()
}
