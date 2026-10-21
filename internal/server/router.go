package server

import ()

func (server *Server) routeClientMessages(clientMessageChannel chan ClientMessage) {
	for clientMessage := range clientMessageChannel {
		if server.gameHost.Started() {
			server.routePokerClientMessage(clientMessage)
		} else {
			server.routeLobbyClientMessage(clientMessage)
		}
	}
}

func (server *Server) routePokerClientMessage(clientMessage ClientMessage) {
	playerAction, parseErr := parsePokerClientMessage(clientMessage)

	response := buildActionResponse(playerAction, parseErr)

	sendErr := sendClientResponse(clientMessage.client, response)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending client response:",
			sendErr,
		)
		return
	}

	logPlayerAction(
		clientMessage.client,
		playerAction,
		parseErr,
	)
}

func (server *Server) routeLobbyClientMessage(clientMessage ClientMessage) {
	if !server.gameHost.ClientHasPlayer(clientMessage.client.id) {
		server.handlePlayerRegistration(clientMessage)
		return
	}

	currentPlayer := server.gameHost.GetPlayer(clientMessage.client.id)
	if currentPlayer == nil {
		return
	}

	if server.gameHost.IsLobbyOwner(currentPlayer) &&
		server.gameHost.GetExpectedNumberOfPlayers() == 0 {
		server.handleExpectedNumberOfPlayers(clientMessage)
		return
	}

}
