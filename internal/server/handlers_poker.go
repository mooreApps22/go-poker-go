package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func (server *Server) handleInvalidPokerAction(
	clientMessage ClientMessage,
	parseErr error,
) {
	response := fmt.Sprintf("Invalid action: %s\n", parseErr)

	sendErr := sendClientResponse(clientMessage.client, response)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending client response:",
			sendErr,
		)
	}
}

func sendPlayerAction(
	clientMessage ClientMessage,
	response string,
	playerAction poker.PlayerAction,
) {
	sendErr := sendClientResponse(clientMessage.client, response)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending client response:",
			sendErr,
		)
		return
	}

	logPlayerAction(clientMessage.client, playerAction)
}

func (server *Server) handleChipRequest(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	response := fmt.Sprintf(
		"You have $%d in chips.\n",
		server.gameHost.GetPlayer(clientMessage.client.id).GetChips(),
	)

	sendPlayerAction(
		clientMessage,
		response,
		playerAction,
	)
}

func (server *Server) handlePlayerAction(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	switch playerAction.Action {
	case poker.Fold:
		server.handleFold(clientMessage, playerAction)
	case poker.Check:
		server.handleCheck(clientMessage, playerAction)
	case poker.Call:
		server.handleCall(clientMessage, playerAction)
	case poker.Bet:
		server.handleBet(clientMessage, playerAction)
	case poker.Raise:
		server.handleRaise(clientMessage, playerAction)
	case poker.AllIn:
		server.handleAllIn(clientMessage, playerAction)
	}

}

func (server *Server) handleFold(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to Fold.\n",
		playerAction,
	)
}

func (server *Server) handleCheck(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to Check.\n",
		playerAction,
	)
}

func (server *Server) handleCall(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to Call.\n",
		playerAction,
	)
}

func (server *Server) handleBet(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to Bet.\n",
		playerAction,
	)
}

func (server *Server) handleRaise(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to Raise.\n",
		playerAction,
	)
}

func (server *Server) handleAllIn(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic
	sendPlayerAction(
		clientMessage,
		"You have selected to go All In.\n",
		playerAction,
	)
}
