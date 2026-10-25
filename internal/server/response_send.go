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

func (server *Server) broadcastActionResult(result poker.ActionResult, handleMessage string) {

	for _, client := range server.clients {
		var sendErr error
		var response string
		if client.id == result.Player.GetID() {
			response = fmt.Sprintf("You %s", handleMessage)
		} else {
			response = fmt.Sprintf("%s %s", result.Player.GetName(), handleMessage)
		}
		sendErr = sendClientResponse(client, response)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}

func (server *Server) broadcastTurn(currentPlayer *poker.Player) {
	messageToOther := fmt.Sprintf("It's %s's turn.\n", currentPlayer.GetName())

	for _, client := range server.clients {
		var sendErr error
		if client.id == currentPlayer.GetID() {
			sendErr = sendClientResponse(client, "It's your turn.\n")
		} else {
			sendErr = sendClientResponse(client, messageToOther)
		}
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
