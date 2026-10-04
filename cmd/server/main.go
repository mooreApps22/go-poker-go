package main

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...")
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
		fmt.Println("New Hand [N]")
		fmt.Println("Quit [Q]")
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

		game.NewHand()
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
		/*
			for _, player := range players {
				fmt.Println(player)
				fmt.Println(poker.FormatCards(player.GetHoleCards()))
				fmt.Println()
			}
		*/
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

		/*
			for offset := 0; offset < len(hand.GetPlayers()); offset++ {
				playerIndex := (hand.GetSmallBlindIndex() + offset) % len(hand.GetPlayers())
				player := hand.GetPlayers()[playerIndex]

				fmt.Printf("%v's Best Hand: %v\n", player.GetName(), player.GetBestHandValue().Category.String())
				fmt.Println("Tiebreakers: ", player.GetBestHandValue().Tiebreakers)
				fmt.Println()
			}
		*/

		hand.CreateSidePots()
		hand.PickWinners()
		hand.AwardPot()

		//DISPLAY WINNERS
		for _, player := range hand.GetWinner() {
			fmt.Printf("Winner: %v\n", player.GetName())
			fmt.Println(poker.FormatCards(player.GetBestCards()))
		}

		fmt.Println()
		fmt.Println("END OF HAND")
		fmt.Println()
	}
}
