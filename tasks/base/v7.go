package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Maps struct {
	key   int
	value int
}

func insertData(ch chan Maps, data map[int]int, sn *sync.Mutex) {
	c := <-ch
	sn.Lock()
	data[c.key] = c.value
	sn.Unlock()
}

func toChanMap(ch chan Maps) {
	key := rand.Intn(100)
	value := rand.Intn(100)
	ch <- Maps{key, value}
}

func TaskSeven() {
	ch := make(chan Maps)
	data := make(map[int]int)
	sn := sync.Mutex{}
	defer close(ch)

	for i := 0; i < 10; i++ {
		go toChanMap(ch)
		go insertData(ch, data, &sn)
	}

	time.Sleep(time.Second)
	fmt.Println(data)
}
