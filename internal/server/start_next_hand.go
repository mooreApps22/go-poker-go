package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
	"strings"
)

func (server *Server) startNextHand() {
	err := server.gameHost.StartHand()
	if err != nil {
		return
	}

	hand := server.gameHost.GetHand()
	dealer := hand.GetDealer()
	smallBlind := hand.GetSmallBlindPlayer()
	bigBlind := hand.GetBigBlindPlayer()
	currentPlayer := hand.GetCurrentPlayer()

	server.broadcast("The table is ready for the NEXT HAND.\n")
	server.broadcast("The deck has been shuffled.\n")

	server.broadcast(fmt.Sprintf("%s is the dealer.\n", dealer.GetName()))

	server.broadcast(fmt.Sprintf(
		"%s posted the small blind: $%d.\n",
		smallBlind.GetName(),
		poker.SmallBlindBet,
	))
	server.broadcast(fmt.Sprintf(
		"%s posted the big blind: $%d.\n",
		bigBlind.GetName(),
		poker.BigBlindBet,
	))

	server.broadcast("The hole card have been dealt.\n")

	server.broadcast("These are your hole cards:\n")
	server.sendPrivateHoleCards()

	server.broadcastTurn(currentPlayer)
}

//New

func (server *Server) waitForNextDeal() {
	nextDealer := server.gameHost.GetNextDealer()

	server.gameHost.SetWaitingForDeal(true)

	server.broadcastWaitingForDealer(nextDealer)

}

func (server *Server) handleDealCommand(clientMessage ClientMessage) {
	nextDealer := server.gameHost.GetNextDealer()

	if clientMessage.client.id != nextDealer.GetID() {
		sendClientResponse(
			clientMessage.client,
			"Waiting for the dealer to start the next hand.\n",
		)
		return
	}

	command := strings.ToUpper(
		strings.TrimSpace(clientMessage.message),
	)

	switch command {
	case "", "D", "DEAL":
		server.gameHost.SetWaitingForDeal(false)
		server.startNextHand()
	default:
		sendClientResponse(
			clientMessage.client,
			"Type 'deal' to deal.\n",
		)
	}

}
