package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	port := "8080" // Specify the port you want to monitor

	// Create a TCP listener on the specified port
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	fmt.Println("Listening on port", port)

	// Accept incoming connections in a loop
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection (e.g., log the connection details)
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Get remote address and log the connection
	remoteAddr := conn.RemoteAddr().String()
	fmt.Println("Incoming request from:", remoteAddr)

	// Optionally, you can read data from the connection or send a response
}
