package main

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...")

	deck := poker.NewDeck()

	deck.Shuffle()

	deck.Print()

	player1 := poker.NewPlayer(1, "Skyy", 1_000_000)

	fmt.Println(player1.String())

	card1 := deck.Draw()

	fmt.Println(card1)
}
