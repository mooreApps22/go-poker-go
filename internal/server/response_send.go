package server

import (
	"fmt"
)

func sendClientResponse(client *Client, response string) error {
	_, err := client.connection.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error writing to connection: %v\n", err)
		return err
	}
	return nil
}

func (server *Server) broadcast(message string) {
	for _, client := range server.clients {
		sendErr := sendClientResponse(client, message)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}
