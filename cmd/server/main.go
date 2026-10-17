package main

import (
	"fmt"

	//	"github.com/mooreApps22/go-poker-go/internal/poker"
	"github.com/mooreApps22/go-poker-go/internal/server"
)

func main() {

	pokerServer, err := server.NewServer(8080)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}
	fmt.Println("Go Poker Go Server is listening...")

	pokerServer.Run()

	/*
		fmt.Println()

		player1 := poker.NewPlayer(1, "Adam")
		player2 := poker.NewPlayer(2, "Bill")
		player3 := poker.NewPlayer(3, "Cathy")
		player4 := poker.NewPlayer(4, "Debra")

		players := []*poker.Player{
			&player1,
			&player2,
			&player3,
			&player4,
		}

		game := poker.NewGame(players)

		for {
			fmt.Println("[N] New Hand")
			fmt.Println("[Q] Quit")
			fmt.Print("> ")

			game.RotateDealer()
			var input string
			fmt.Scan(&input)

			if input == "q" || input == "Q" {
				break
			} else if input == "n" || input == "N" {
				fmt.Println("Starting new hand...")
			} else {
				fmt.Println("Invalid command: ", input)
				continue
			}

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
