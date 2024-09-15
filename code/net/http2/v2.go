package main

import (
	"fmt"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Router struct {
	routes map[string]HandlerFunc
}

func NewRouter() *Router {
	return &Router{make(map[string]HandlerFunc)}
}

func (r *Router) Handle(path string, handler HandlerFunc) {
	r.routes[path] = handler
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := req.URL.Path

	if handler, ok := r.routes[path]; ok {
		handler(w, req)
	} else {
		http.NotFound(w, req)
	}
}

func main() {
	router := NewRouter()

	router.Handle("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Welcome to homepage!")
	})
}

//func main() {
//	ln, err := net.Listen("tcp", ":8080")
//	if err != nil {
//		fmt.Println("Ошибка при запуске сервера:", err)
//		return
//	}
//	defer ln.Close()
//
//	fmt.Println("Сервер запущен на порту 8080...")
//
//	for {
//		conn, err := ln.Accept()
//		if err != nil {
//			fmt.Println("Ошибка при подключении:", err)
//			continue
//		}
//
//		go handleConnection(conn)
//	}
//}
//
//func handleConnection(conn net.Conn) {
//	defer conn.Close()
//
//	buf := make([]byte, 1024)
//	n, err := conn.Read(buf)
//	if err != nil {
//		fmt.Println("Ошибка чтения:", err)
//		return
//	}
//
//	request := string(buf[:n])
//	fmt.Println("Получен запрос:\n", request)
//
//	lines := strings.Split(request, "\r\n")
//	if len(lines) > 0 {
//		firstLine := strings.Split(lines[0], " ")
//		if len(firstLine) >= 2 {
//			method := firstLine[0]
//			path := firstLine[1]
//
//			if method == "GET" {
//				switch path {
//				case "/":
//					handleRoot(conn)
//				case "/about":
//					handleAbout(conn)
//				default:
//					handleNotFound(conn)
//				}
//			} else {
//				response := "HTTP/1.1 405 Method Not Allowed\r\n" +
//					"Content-Type: text/plain\r\n" +
//					"Content-Length: 19\r\n" +
//					"\r\n" +
//					"Method Not Allowed"
//				conn.Write([]byte(response))
//			}
//		}
//	}
//}
//
//func handleRoot(conn net.Conn) {
//	response := "HTTP/1.1 200 OK\r\n" +
//		"Content-Type: text/plain\r\n" +
//		"Content-Length: 11\r\n" +
//		"\r\n" +
//		"Hello Root!"
//	conn.Write([]byte(response))
//}
//
//func handleAbout(conn net.Conn) {
//	response := "HTTP/1.1 200 OK\r\n" +
//		"Content-Type: text/plain\r\n" +
//		"Content-Length: 12\r\n" +
//		"\r\n" +
//		"About Page!"
//	conn.Write([]byte(response))
//}
//
//func handleNotFound(conn net.Conn) {
//	response := "HTTP/1.1 404 Not Found\r\n" +
//		"Content-Type: text/plain\r\n" +
//		"Content-Length: 9\r\n" +
//		"\r\n" +
//		"Not Found"
//
//	conn.Write([]byte(response))
//}
