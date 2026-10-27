package server

import (
	"strings"
)

func (server *Server) routeClientMessages(clientMessageChannel chan ClientMessage) {
	for clientMessage := range clientMessageChannel {
		if server.gameHost.Started() {
			server.routeGameClientMessage(clientMessage)
		} else {
			server.routeLobbyClientMessage(clientMessage)
		}
	}
}

func (server *Server) routeGameClientMessage(
	clientMessage ClientMessage,
) {
	command := strings.ToUpper(
		strings.TrimSpace(clientMessage.message),
	)

	switch command {
	case "P", "CHIP", "CHIPS":
		server.handleChips(clientMessage)
		return
	case "H", "HOLE":
		server.handleHoleCards(clientMessage)
		return
	case "M", "MENU":
		server.handleMenu(clientMessage)
		return
	}

	playerAction, parseErr := parsePokerClientMessage(clientMessage)
	if parseErr != nil {
		server.handleInvalidPokerAction(clientMessage, parseErr)
		return
	}

	server.handlePlayerAction(clientMessage, playerAction)
}
