package main

import (
	"log"
	"os"
	"time"
)

func TaskOne() {
	logger := log.New(os.Stderr, "", 0)

	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}

	log.Printf("The current time is %s", currentTime)

	response, err = ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		logger.Fatal(err.Error())
	}

	timeWithOffset := time.Now().Add(response.ClockOffset)
	log.Printf("The time with metadata %s", timeWithOffset)
}
