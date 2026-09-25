package main

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...")
	card := poker.Card{
		Rank: poker.Ace,
		Suit: poker.Spades,
	}

	fmt.Println(card)
}
