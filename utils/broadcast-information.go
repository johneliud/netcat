package utils

import (
	"fmt"
	"log"
	"net"
	"time"
)

// Used to broadcast a client's message to other clients in the chat room.
func broadcastInformation(message string, senderConn net.Conn) {
	chatHistoryMu.Lock()
	chatHistory = append(chatHistory, message)
	chatHistoryMu.Unlock()

	mu.Lock()
	defer mu.Unlock()

	for _, client := range clientsList {
		// Exclude sender from seeing their broadcast event.
		if client.conn != senderConn {
			_, err := client.conn.Write([]byte(message))
			client.conn.Write([]byte(fmt.Sprintf("\n[%s][%s]:", time.Now().Format(time.DateTime), client.name)))
			if err != nil {
				log.Printf("Error broadcasting to %v: %v\n", &client.name, err)
				go removeClient(client.conn)
			}
		}
	}
}
