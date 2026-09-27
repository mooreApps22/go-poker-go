package main

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...\n")

	player1 := poker.NewPlayer(1, "Skyy", 1_000_000)
	player2 := poker.NewPlayer(2, "Emily", 1_000_000)
	player3 := poker.NewPlayer(3, "Sheeba", 1_000_000)
	player4 := poker.NewPlayer(4, "Cleo", 1_000_000)

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

	fmt.Println("Community Flop Cards:")
	fmt.Println(poker.PrettyCardsString(hand.CommunityCards()))
}
