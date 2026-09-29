package main

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...\n")

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

	hand := poker.NewHand(players)
	err := hand.DealHoleCards()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, player := range players {
		fmt.Println(player)
		fmt.Println(poker.PrettyCardsString(player.HoleCards[:]))
		fmt.Println()
	}

	hand.DealFlopCards()

	//	fmt.Println("Community Flop Cards:")
	//	fmt.Println(poker.PrettyCardsString(hand.CommunityCards()))

	hand.DealTurnCard()

	//	fmt.Println("Community Turn Card:")
	//	fmt.Println(poker.PrettyCardsString(hand.CommunityCards()))

	hand.DealRiverCard()

	fmt.Println("Community River Card:")
	fmt.Println(poker.PrettyCardsString(hand.CommunityCards()))

	//Test
	hand.EvaluateEachPlayersBestHand()

	fmt.Println("Adam's Best Hand: ", player1.BestHandValue.Category.String())
	fmt.Println("Tiebreakers: ", player1.BestHandValue.Tiebreakers)
	fmt.Println()

	fmt.Println("Bill's Best Hand: ", player2.BestHandValue.Category.String())
	fmt.Println("Tiebreakers: ", player2.BestHandValue.Tiebreakers)
	fmt.Println()

	fmt.Println("Cathy's Best Hand: ", player3.BestHandValue.Category.String())
	fmt.Println("Tiebreakers: ", player3.BestHandValue.Tiebreakers)
	fmt.Println()

	fmt.Println("Debra's Best Hand: ", player4.BestHandValue.Category.String())
	fmt.Println("Tiebreakers: ", player4.BestHandValue.Tiebreakers)
	fmt.Println()
}
