package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	repo := NewInMemoryRepository()
	service := NewUserService(repo)

	http.HandleFunc("/login", service.Login)
	http.HandleFunc("/users", service.HandleUsers)
	http.HandleFunc("/users/", service.HandleUserByID)
	http.HandleFunc("/users/items", service.GetItemsByUser)

	server := &http.Server{
		Addr:         ":8080",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(server.ListenAndServe())
}
