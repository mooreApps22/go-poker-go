package main

import (
	"fmt"

	//	"github.com/mooreApps22/go-poker-go/internal/poker"
	"github.com/mooreApps22/go-poker-go/internal/host"
	"github.com/mooreApps22/go-poker-go/internal/server"
)

func main() {

	gameHost := host.NewHost()

	pokerServer, err := server.NewServer(8080, &gameHost)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}
	fmt.Println("Go Poker Go Server is listening...")

	pokerServer.Run()

	/*

		for {
			game.RotateDealer()
			game.SetUpNewHand()
			hand := game.GetHand()

			// PRE-FLOP PHASE
			fmt.Println("Blind bets posted:")
			hand.ResetPlayersForNewHand()
			hand.PostBlinds()
			for _, player := range hand.GetPlayers() {
				fmt.Printf("%v's Current Bet: %v\n", player.GetName(), player.GetCurrentBet())
			}

			err := hand.DealHoleCards()
			if err != nil {
				fmt.Println("Error:", err)
				return
			}
			hand.AcceptBets()

			//FLOP PHASE
			hand.DealFlopCards()
			hand.AcceptBets()
			fmt.Println("Community Flop Cards:")
			fmt.Println(poker.FormatCards(hand.GetCommunityCardsDealt()))

			//TURN PHASE
			hand.DealTurnCard()
			hand.AcceptBets()
			fmt.Println("Community Turn Card:")
			fmt.Println(poker.FormatCards(hand.GetCommunityCardsDealt()))

			//RIVER PHASE
			hand.DealRiverCard()
			hand.AcceptBets()
			fmt.Println("Community River Card:")
			fmt.Println(poker.FormatCards(hand.GetCommunityCardsDealt()))

			hand.ShowDown()

			hand.CreatePots()
			hand.PickWinners()
			hand.AwardPots()

			//DISPLAY WINNERS
			for _, player := range hand.GetWinner() {
				fmt.Printf("Winner: %v\n", player.GetName())
				fmt.Println(poker.FormatCards(player.GetBestCards()))
			}

			fmt.Println()
			fmt.Println("END OF HAND")
			fmt.Println()
		}
	*/
}
