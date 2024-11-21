package utils

import "net"

// Will remove a client from the chat upon termination by the client.
func removeClient(conn net.Conn) {
	mu.Lock()
	defer mu.Unlock()

	for i, client := range clientsList {
		if client.conn == conn {
			client.conn.Close()
			clientsList = append(clientsList[:i], clientsList[i+1:]...)
			break
		}
	}
}
