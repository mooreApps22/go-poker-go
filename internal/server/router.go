package server

import (
	"strings"
)

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
	if parseErr != nil {
		server.handleInvalidPokerAction(clientMessage, parseErr)
		return
	}

	server.handlePlayerAction(clientMessage, playerAction)
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

	if server.gameHost.AutomaticStartPossible() {
		server.handleStartGame(clientMessage)
	}

	if server.gameHost.IsLobbyOwner(currentPlayer) {
		command := strings.ToUpper(strings.TrimSpace(clientMessage.message))

		switch command {
		case "START":
			server.handleStartGame(clientMessage)
			return
		}
	}

}
