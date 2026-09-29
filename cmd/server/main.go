package main

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...")
	fmt.Println()

	player1 := poker.NewPlayer(1, "Adam", 1_000_000)
	player2 := poker.NewPlayer(2, "Bill", 1_000_000)
	player3 := poker.NewPlayer(3, "Cathy", 1_000_000)
	player4 := poker.NewPlayer(4, "Debra", 1_000_000)

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

		err := hand.DealHoleCards()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		for _, player := range players {
			fmt.Println(player)
			fmt.Println(poker.FormatCards(player.HoleCards[:]))
			fmt.Println()
		}

		hand.DealFlopCards()

		fmt.Println("Community Flop Cards:")
		fmt.Println(poker.FormatCards(hand.CommunityCards()))

		hand.DealTurnCard()

		fmt.Println("Community Turn Card:")
		fmt.Println(poker.FormatCards(hand.CommunityCards()))

		hand.DealRiverCard()

		fmt.Println("Community River Card:")
		fmt.Println(poker.FormatCards(hand.CommunityCards()))

		//Test
		hand.EvaluateEachPlayersBestHand()

		for offset := 0; offset < len(hand.GetPlayers()); offset++ {
			playerIndex := (hand.GetSmallBlindIndex() + offset) % len(hand.GetPlayers())
			player := hand.GetPlayers()[playerIndex]

			fmt.Printf("%v's Best Hand: %v\n", player.Name, player.BestHandValue.Category.String())
			fmt.Println("Tiebreakers: ", player.BestHandValue.Tiebreakers)
			fmt.Println()
		}

		hand.PickWinners()
		fmt.Println("Winner: ", hand.GetWinner())
		game.RotateDealer()
	}
}
