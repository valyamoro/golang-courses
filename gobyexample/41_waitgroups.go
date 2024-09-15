package main

import (
	"fmt"
	"sync"
	"time"
)

func worker3(id int) {
	fmt.Printf("worker %d starting\n", id)
	time.Sleep(time.Second)
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			worker3(i)
		}(i)
	}

	wg.Wait()
}
