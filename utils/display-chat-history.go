package utils

// Used to show every chat event to a new client after joining the chat.
func displayChatHistory(client *Client) {
	chatHistoryMu.Lock()
	defer chatHistoryMu.Unlock()

	if len(chatHistory) > 0 {
		client.conn.Write([]byte("\n --Chat History-- \n"))
		for _, history := range chatHistory {
			client.conn.Write([]byte(history))
		}
		client.conn.Write([]byte("\n --End Of Chat History-- \n"))
	}
}
