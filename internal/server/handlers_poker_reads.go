package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func (server *Server) handleMenu(
	clientMessage ClientMessage,
) {
	player := server.gameHost.GetPlayer(clientMessage.client.id)
	if player == nil {
		return
	}

	hand := server.gameHost.GetHand()

	response := hand.GetPlayerActionOptions(player)

	err := sendClientResponse(
		clientMessage.client,
		response,
	)
	if err != nil {
		logClientError(
			clientMessage.client,
			"Error sending hole cards",
			err,
		)
	}

}

func (server *Server) handleHoleCards(
	clientMessage ClientMessage,
) {
	player := server.gameHost.GetPlayer(clientMessage.client.id)
	if player == nil {
		return
	}

	response := poker.FormatPrivateCards(
		player.GetHoleCards(),
	)

	err := sendClientResponse(
		clientMessage.client,
		response,
	)
	if err != nil {
		logClientError(
			clientMessage.client,
			"Error sending hole cards",
			err,
		)
	}
}

func (server *Server) handleChips(
	clientMessage ClientMessage,
) {
	player := server.gameHost.GetPlayer(clientMessage.client.id)
	if player == nil {
		return
	}

	response := fmt.Sprintf(
		"You have $%d in chips.\n",
		player.GetChips(),
	)

	err := sendClientResponse(
		clientMessage.client,
		response,
	)
	if err != nil {
		logClientError(
			clientMessage.client,
			"Error sending hole cards",
			err,
		)
	}
}
