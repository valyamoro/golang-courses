package main

import (
	"fmt"
	"math/rand"
	"time"
)

func readFromChannel(channel chan int) {
	for data := range channel {
		fmt.Printf("Now data is %d\n", data)
	}
}

func writeToChannel(channel chan int) {
	for {
		data := rand.Intn(100)
		channel <- data
	}
}

func TaskFive() {
	N := 10

	channel := make(chan int)
	defer close(channel)
	go writeToChannel(channel)
	go readFromChannel(channel)

	time.Sleep(time.Duration(N) * time.Second)
}

func main() {
	TaskFive()
}
