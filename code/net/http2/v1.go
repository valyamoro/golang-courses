package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		name := r.FormValue("name")
		fmt.Fprintf(w, "Hello, %s!", name)
	} else {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/", helloHandler)
	http.HandleFunc("/form", formHandler)
	fmt.Println("Сервер запущен на http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func main2() {
	resp, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Err reading", err)
		return
	}

	fmt.Println("Response:", string(body))
}

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func main3() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", helloHandler)

	wrappedMux := LoggingMiddleware(mux)

	fmt.Println("Server started at http://localhost:8080")
	http.ListenAndServe(":8080", wrappedMux)
}
