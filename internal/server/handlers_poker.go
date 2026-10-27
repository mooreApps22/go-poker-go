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

	// APPLY PLAYER ACTION
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

	switch result.Action {
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

		if hand.GetPhase() != poker.Showdown {
			communityCards := hand.GetCommunityCardsDealt()
			formattedCommunityCards := poker.FormatPublicCards(communityCards)

			server.broadcast(hand.GetPhase().String())

			server.broadcast(formattedCommunityCards)

			hand.SetFirstPostFlopPlayer()
			currentPlayer := hand.GetCurrentPlayer()

			server.broadcastTurn(currentPlayer)
		} else {
			server.broadcastShowdown(hand)
			server.startNextHand()
		}
	} else {
		currentPlayer := hand.GetCurrentPlayer()

		server.broadcastTurn(currentPlayer)
	}
}

func (server *Server) handleFold(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcastActionResult(result, "folded.\n")

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleCheck(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcastActionResult(result, "checked.\n")

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleCall(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcastActionResult(result, fmt.Sprintf(
		"called $%d.\n",
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleBet(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcastActionResult(result, fmt.Sprintf(
		"bet $%d.\n",
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleRaise(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	server.broadcastActionResult(result, fmt.Sprintf(
		"raised $%d.\n",
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}

func (server *Server) handleAllIn(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {

	server.broadcastActionResult(result, fmt.Sprintf(
		"goes all in with $%d.\n",
		result.Amount,
	))

	logPlayerAction(clientMessage.client, result)
}
