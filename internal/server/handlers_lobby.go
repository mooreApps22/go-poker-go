package server

import (
	"fmt"
	"strconv"
	"strings"
)

func (server *Server) handlePlayerRegistration(clientMessage ClientMessage) {
	name := strings.TrimSpace(clientMessage.message)

	response := server.gameHost.RegisterPlayer(
		clientMessage.client.id,
		name,
	)
	sendErr := sendClientResponse(clientMessage.client, response)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending lobby response",
			sendErr,
		)
	}
}

func (server *Server) handleExpectedNumberOfPlayers(clientMessage ClientMessage) {
	number, err := strconv.Atoi(strings.TrimSpace(clientMessage.message))
	if err != nil {
		sendErr := sendClientResponse(
			clientMessage.client,
			"Expected a number.\n",
		)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending lobby response",
				sendErr,
			)
		}
		return
	}

	if !server.gameHost.SetExpectedNumberOfPlayers(number) {
		sendErr := sendClientResponse(
			clientMessage.client,
			"Number can't be under 4 or over 10.\n",
		)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending lobby response",
				sendErr,
			)
		}
		return
	}

	sendErr := sendClientResponse(
		clientMessage.client,
		fmt.Sprintf("Waiting for %d players.\n", number),
	)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending lobby response",
			sendErr,
		)
	}
}

func (server *Server) handleStartGame(clientMessage ClientMessage) {
	if !server.gameHost.CanStartGame() {
		// send message: not enough players yet
		sendErr := sendClientResponse(
			clientMessage.client,
			"Waiting for at least four players.\n",
		)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending lobby response",
				sendErr,
			)
		}
		return
	}

	server.gameHost.SetStartGame()
	sendErr := sendClientResponse(
		clientMessage.client,
		"Start Game.\n",
	)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending lobby response",
			sendErr,
		)
	}
}
