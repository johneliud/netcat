package utils

import (
	"fmt"
	"log"
	"net"
	"sync"
)

// The Client struct is used to represent a connected client in the chat.
type Client struct {
	name string
	conn net.Conn
}

var (
	maximumConnections = 10
	clientsList        = make([]*Client, 0, maximumConnections)
	mu                 sync.Mutex // Used to synchronize access to all shared resources in the chat.
	chatHistory        []string
	chatHistoryMu      sync.Mutex
)

// Used for establishing a TCP connection with the netcat server (nc).
func StartServer(port string) {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Printf("Error starting server on port %q: %v\n", port, err)
		return
	}
	defer listener.Close()

	fmt.Printf("Successfully listening on port %q\n", port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Error accepting connection: %v\n", err)
			continue
		}

		mu.Lock()
		// Limit access to the chat room if the number of clients exceed the maximum connections.
		if len(clientsList) >= maximumConnections {
			mu.Unlock()
			conn.Write([]byte("Chat room currently full. Try again later.\n"))
			conn.Close()
			continue
		}
		mu.Unlock()

		go handleNewConnection(conn)
	}
}
