package main

import (
	"fmt"
	"time"
)

func Sleep(s int) {
	<-time.After(time.Duration(s) * time.Second)
}

func TaskTwentyFive() {
	fmt.Println("Start...")
	Sleep(3)
	fmt.Println("Finish!")
}
