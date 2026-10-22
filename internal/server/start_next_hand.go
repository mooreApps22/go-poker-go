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
	currentPlayer := hand.GetUTGPlayer()

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

	server.broadcast(fmt.Sprintf(
		"It's %s's turn.\n",
		currentPlayer.GetName(),
	))

	// privately send hole cards
	// tell everyone whose turn it is
}
