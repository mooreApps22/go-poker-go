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
	result poker.ActionResult,
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

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handlePlayerAction(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {

	hand := server.gameHost.GetHand()

	result, err := hand.ApplyPlayerAction(&playerAction)
	if err != nil {
		sendErr := sendClientResponse(
			clientMessage.client,
			err.Error()+"\n",
		)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending player action error",
				sendErr,
			)
		}
		return
	}

	switch playerAction.Action {
	case poker.Chips:
		server.handleChips(clientMessage, result)
	case poker.Fold:
		server.handleFold(clientMessage, result)
	case poker.Check:
		server.handleCheck(clientMessage, result)
	case poker.Call:
		server.handleCall(clientMessage, result)
	case poker.Bet:
		server.handleBet(clientMessage, result)
	case poker.Raise:
		server.handleRaise(clientMessage, result)
	case poker.AllIn:
		server.handleAllIn(clientMessage, result)
	}

	if result.BettingRoundEnded {

		err := hand.AdvancePhase()
		if err != nil {
			// handle error
			return
		}

		potValue := server.gameHost.GetHand().GetCurrentPotValue()
		potValueResponse := fmt.Sprintf("Current Pot Value: $%d.\n", potValue)
		server.broadcast(potValueResponse)

		communityCards := hand.GetCommunityCardsDealt()
		formattedCommunityCards := poker.FormatPublicCards(communityCards)

		server.broadcast(formattedCommunityCards)

	} else {
		currentPlayer := hand.GetCurrentPlayer()

		server.broadcastTurn(currentPlayer)
	}

}

func (server *Server) handleChips(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	response := fmt.Sprintf(
		"You have $%d in chips.\n",
		result.Amount,
	)

	sendPlayerAction(clientMessage, response, result)

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleFold(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcast(fmt.Sprintf(
		"%s folded.\n",
		result.Player.GetName(),
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleCheck(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcast(fmt.Sprintf(
		"%s checked.\n",
		result.Player.GetName(),
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleCall(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcast(fmt.Sprintf(
		"%s called $%d.\n",
		result.Player.GetName(),
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleBet(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcast(fmt.Sprintf(
		"%s bet $%d.\n",
		result.Player.GetName(),
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleRaise(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcast(fmt.Sprintf(
		"%s raises $%d.\n",
		result.Player.GetName(),
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleAllIn(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {

	server.broadcast(fmt.Sprintf(
		"%s goes all in with $%d.\n",
		result.Player.GetName(),
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}
