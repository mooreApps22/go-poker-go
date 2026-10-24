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

	hand := server.gameHost.GetHand()

	bettingRoundComplete, err := hand.ApplyPlayerAction(&playerAction)
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

	if bettingRoundComplete {

		err := hand.AdvancePhase()
		if err != nil {
			// handle error
			return
		}

		server.broadcastPhase()

		communityCards := hand.GetCommunityCardsDealt()
		formattedCommunityCards := poker.FormatPublicCards(communityCards)

		server.broadcast(formattedCommunityCards)

	} else {
		currentPlayer := hand.GetCurrentPlayer()

		server.broadcastTurn(currentPlayer)
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
	player := server.gameHost.GetPlayer(playerAction.PlayerID)

	server.broadcast(fmt.Sprintf(
		"%s called $%d.\n",
		player.GetName(),
		playerAction.Amount,
	))

	logPlayerAction(clientMessage.client, playerAction)
}

func (server *Server) handleBet(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	//placeholder logic

	player := server.gameHost.GetPlayer(playerAction.PlayerID)

	server.broadcast(fmt.Sprintf(
		"%s bet $%d.\n",
		player.GetName(),
		playerAction.Amount,
	))

	logPlayerAction(clientMessage.client, playerAction)
}

func (server *Server) handleRaise(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	player := server.gameHost.GetPlayer(playerAction.PlayerID)

	server.broadcast(fmt.Sprintf(
		"%s raises $%d.\n",
		player.GetName(),
		playerAction.Amount,
	))

	logPlayerAction(clientMessage.client, playerAction)
}

func (server *Server) handleAllIn(
	clientMessage ClientMessage,
	playerAction poker.PlayerAction,
) {
	player := server.gameHost.GetPlayer(playerAction.PlayerID)

	server.broadcast(fmt.Sprintf(
		"%s goes all in with $%d.\n",
		player.GetName(),
		playerAction.Amount,
	))

	logPlayerAction(clientMessage.client, playerAction)
}
