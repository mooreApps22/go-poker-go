package main

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func main() {
	fmt.Println("Poker server starting...")

	deck := poker.NewDeck()

	deck.PrintDeck()
}
