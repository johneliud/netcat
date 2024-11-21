package utils

import (
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

// Used to set up a connection environment for a client.
func handleNewConnection(conn net.Conn) {
	// Enforce a 1 min timeout for joining the chat to save on resources.
	conn.SetDeadline(time.Now().Add(60 * time.Second))

	displayLogo(conn)

	conn.Write([]byte("[ENTER YOUR NAME]: "))

	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		log.Printf("Error reading client name: %v\n", err)
		conn.Close()
		return
	}

	// Reset timeout after getting client's name.
	conn.SetDeadline(time.Time{})

	clientName := strings.TrimSpace(string(buffer[:n]))
	if clientName == "" {
		conn.Write([]byte("Name cannot be an empty value.\n"))
		conn.Close()
		return
	}

	// Prevent having duplicate names in the chat before granting client access.
	mu.Lock()
	for _, client := range clientsList {
		if strings.EqualFold(client.name, clientName) {
			mu.Unlock()
			conn.Write([]byte("Name already taken. Try another name to join the chat.\n"))
			conn.Close()
			return
		}
	}
	mu.Unlock()

	newClient := &Client{
		name: clientName,
		conn: conn,
	}
	mu.Lock()
	clientsList = append(clientsList, newClient)
	mu.Unlock()

	conn.Write([]byte(fmt.Sprintf("Welcome to the chat, %s!\n", clientName)))

	// Display chat history to the joining client.
	displayChatHistory(newClient)

	// Notify others of the new client.
	broadcastInformation(fmt.Sprintf("\n%v has joined our chat...", clientName), conn)

	go handleClient(newClient)
}
