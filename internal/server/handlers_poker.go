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
	result, err := hand.ApplyPlayerAction(playerAction)
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

	server.handleActionResult(clientMessage, result)

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

func (server *Server) handleActionResult(
	clientMessage ClientMessage,
	result poker.ActionResult,
) {
	var message string

	switch result.Action {
	case poker.Fold:
		message = "folded.\n"
	case poker.Check:
		message = "checked.\n"
	case poker.Call:
		message = fmt.Sprintf(
			"called $%d.\n",
			result.Amount,
		)
	case poker.Bet:
		message = fmt.Sprintf(
			"bet $%d.\n",
			result.Amount,
		)
	case poker.Raise:
		message = fmt.Sprintf(
			"raised $%d.\n",
			result.Amount,
		)
	case poker.AllIn:
		message = fmt.Sprintf(
			"goes all in with $%d.\n",
			result.Amount,
		)
	default:
		return
	}

	server.broadcastActionResult(result, message)
	logPlayerAction(clientMessage.client, result)
}
