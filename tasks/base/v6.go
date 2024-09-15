package main

import (
	"fmt"
	"time"
)

func TaskSix() {
	ch1 := make(chan string)
	run := true

	go func(run *bool) {
		for {
			select {
			case <-ch1:
				fmt.Println("Goroutine 1 done (<-ch1)!")
				return
			default:

			}
		}
	}(&run)

	time.Sleep(150 * time.Millisecond)
	run = false
	time.Sleep(350 * time.Millisecond)

	close(ch1)

	ch2 := make(chan int)
	go func() {
		for {
			num, more := <-ch2
			if !more {
				fmt.Println("Goroutine 2 done!")
				return
			}

			fmt.Printf("Goroutine 2 working says %d\n", num)
		}
	}()

	ch2 <- 1
	ch2 <- 2
	ch2 <- 3
	close(ch2)

	time.Sleep(350 * time.Millisecond)
}

func main() {
	TaskSix()
}
