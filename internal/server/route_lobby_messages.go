package server

import (
	"strings"
)

func (server *Server) routeLobbyClientMessage(clientMessage ClientMessage) {
	clientID := clientMessage.client.id

	if !server.gameHost.ClientHasPlayer(clientID) {
		server.handlePlayerRegistration(clientMessage)
		return
	}

	player := server.gameHost.GetPlayer(clientID)
	if player == nil {
		return
	}

	if server.gameHost.IsLobbyOwner(player) &&
		server.gameHost.GetExpectedNumberOfPlayers() == 0 {
		server.handleExpectedNumberOfPlayers(clientMessage)
		return
	}

	if server.gameHost.AutomaticStartPossible() {
		server.handleStartGame(clientMessage)
	}

	if server.gameHost.IsLobbyOwner(player) {
		server.routeLobbyOwnerMessage(clientMessage)
	}
}

func (server *Server) routeLobbyOwnerMessage(clientMessage ClientMessage) {
	command := strings.ToUpper(strings.TrimSpace(clientMessage.message))

	/*
		Potential Cases:
		KICK, PLAYERS, SET <NUMBER>, CLOSE
	*/
	switch command {
	case "START":
		server.handleStartGame(clientMessage)
		return
	}
}
