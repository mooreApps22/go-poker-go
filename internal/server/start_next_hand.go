package server

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
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

	// privately send hole cards
	server.broadcast("These are your hole cards:\n")
	server.sendPrivateHoleCards()

	// tell everyone whose turn it is
	server.broadcast(fmt.Sprintf(
		"It's %s's turn.\n",
		currentPlayer.GetName(),
	))
}

func (server *Server) broadcastPhase() {
	hand := server.gameHost.GetHand()

	switch hand.GetPhase() {
	case poker.PreFlop:
		server.broadcast("Phase: PreFlop\n")
	case poker.Flop:
		server.broadcast("Phase: Flop\n")
	case poker.Turn:
		server.broadcast("Phase: Turn\n")
	case poker.River:
		server.broadcast("Phase: River\n")
	case poker.Showdown:
		server.broadcast("Phase: Showdown\n")
	}
}
