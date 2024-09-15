package main

import (
	"fmt"
	"sync"
)

var counter2 int = 0

func main() {
	ch := make(chan bool)
	var mutex sync.Mutex

	for i := 1; i < 5; i++ {
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++ {
		<-ch
	}

	fmt.Println("The end")
}

func work(number int, ch chan bool, mutex *sync.Mutex) {
	mutex.Lock()
	counter2 = 0
	for k := 1; k <= 5; k++ {
		counter2++
		fmt.Println("Goroutine", number, "-", counter2)
	}
	mutex.Unlock()
	ch <- true
}
