package utils

import (
	"fmt"
	"log"
	"strings"
	"time"
)

// Will handle individual client message processing.
func handleClient(client *Client) {
	defer removeClientFromChat(client.conn)
	buffer := make([]byte, 4096)

	for {
		client.conn.Write([]byte(fmt.Sprintf("[%s][%s]:", time.Now().Format(time.DateTime), client.name)))
		n, err := client.conn.Read(buffer)
		if err != nil {
			if err.Error() == "EOF" {
				// Notify other clients in the chat.
				broadcastInformation(fmt.Sprintf("\n%s has left our chat...", client.name), client.conn)
				return
			}
			log.Printf("Error reading from %v: %v\n", client.name, err)
			return
		}

		clientMessage := strings.TrimSpace(string(buffer[:n]))
		if clientMessage == "" {
			log.Println("Cannot send an empty message to the chat.")
			break
		}
		// Broadcast the message to other clients.
		formattedMessage := fmt.Sprintf("\n[%s][%s]: %s", time.Now().Format(time.DateTime), client.name, clientMessage)
		broadcastInformation(formattedMessage, client.conn)
	}
}
