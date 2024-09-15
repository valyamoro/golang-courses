package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)

	log.Println("Это простое сообщение")
	log.Printf("Это сообщение с форматированием %s", "привет")

	file, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)
	log.Println("Сообщение записано в app.log")

	file, err = os.OpenFile("app2.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(file, "CUSTOM: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("Сообщение от собственного логера")

	file, err = os.Open("nonexistent_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	log.Println("Этот код не выполнится так как программа завершится на log.Fatal")
	log.Panic("Это паника, программма будет завершена с трассировкой стека")
}
