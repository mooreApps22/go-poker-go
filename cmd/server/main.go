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

	card1, err := deck.Draw()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(card1.PrettyCardString())
}
