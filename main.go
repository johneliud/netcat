package main

import (
	"log"
	"net"
	"os"
	"strconv"
	"sync"

	"github.com/johneliud/netcat/utils"
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

const defaultPort = ":8989"

func main() {
	switch len(os.Args) {
	case 1:
		utils.StartServer(defaultPort)
	case 2:
		port := os.Args[1]
		portNum, err := strconv.Atoi(port)
		if err != nil {
			log.Printf("Error converting %q to an int: %v\n", port, err)
			return
		}

		if portNum < 1024 || portNum > 65535 {
			log.Println("Invalid specified port. Must be between 1024 and 65535.")
			return
		}
		utils.StartServer(":" + strconv.Itoa(portNum))
	default:
		log.Println("Usage: go run . or go run . [PORT]")
		return
	}
}
