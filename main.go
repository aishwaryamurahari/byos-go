package main

import (
	"byos-go/routes"
	"byos-go/server"
	"byos-go/utils"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	// Start listening on port 8080
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Error starting TCP server: %v", err)
	}
	defer listener.Close()

	fmt.Println("Server listening on http://localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v", err)
			continue
		}

		// Handle each connection in a new goroutine (like a lightweight thread)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// buffer := make([]byte, 4096)
	// n, err := conn.Read(buffer)
	// if err != nil {
	// 	log.Printf("Error reading data: %v", err)
	// 	return
	// }

	// request := string(buffer[:n])

	var requestData strings.Builder
	tempBuf := make([]byte, 1024)

	for {
		n, err := conn.Read(tempBuf)
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("Read error: %v", err)
			return
		}
		requestData.Write(tempBuf[:n])

		// Optional: stop if end of headers received
		if strings.Contains(requestData.String(), "\r\n\r\n") {
			break
		}
	}

	request := requestData.String()

	parsed, err := utils.ParseRequest(request)
	if err != nil {
		log.Printf("Error parsing request: %v", err)
		return
	}

	// Debug: print incoming request
	fmt.Printf("Method: %s, Path: %s\n", parsed.Method, parsed.Path)
	fmt.Println("==== Incoming Request ====")
	fmt.Println(request)

	if strings.HasPrefix(parsed.Path, "/static/") {
		contentType, body, status := server.ServeStaticFile(parsed.Path)
		statusLine := map[int]string{
			200: "200 OK",
			400: "400 Bad Request",
			403: "403 Forbidden",
			404: "404 Not Found",
			500: "500 Internal Server Error",
		}[status]

		response := fmt.Sprintf("HTTP/1.1 %s\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s",
			statusLine, contentType, len(body), body)
		if _, err := conn.Write([]byte(response)); err != nil {
			log.Printf("Error writing response: %v", err)
		}
		return
	}

	handler, exists := routes.MatchRoute(parsed.Method, parsed.Path)
	if exists {
		contentType, body := handler()
		response := fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: %s\r\nContent-Length: %d\r\n\r\n%s",
			contentType, len(body), body)
		conn.Write([]byte(response))
	} else {
		body := "404 Not Found"
		response := fmt.Sprintf("HTTP/1.1 404 Not Found\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(body), body)
		conn.Write([]byte(response))
	}
}
