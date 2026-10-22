package server

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
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

func (server *Server) sendPrivateHoleCards() {
	for _, client := range server.clients {
		player := server.gameHost.GetPlayer(client.id)
		holeCards := poker.FormatPrivateCards(player.GetHoleCards())
		holeCardsResponse := fmt.Sprintf("%s\n", holeCards)
		sendErr := sendClientResponse(client, holeCardsResponse)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}
