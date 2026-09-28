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

	/*
		for _, player := range players {
			fmt.Println(player)
			fmt.Println(poker.PrettyCardsString(player.HoleCards[:]))
			fmt.Println()
		}
	*/

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

	communityCards := hand.CommunityCards()
	var communityHand [5]poker.Card
	copy(communityHand[:], communityCards)

	communityHandValue := poker.EvaluateBestFiveCardHand(communityHand)
	fmt.Println(communityHandValue.Category.String())
	fmt.Println()

	testCards1 := [5]poker.Card{
		{Rank: poker.Five, Suit: poker.Clubs},
		{Rank: poker.Five, Suit: poker.Hearts},
		{Rank: poker.Three, Suit: poker.Clubs},
		{Rank: poker.Three, Suit: poker.Diamonds},
		{Rank: poker.Three, Suit: poker.Spades},
	}

	testCards2 := [5]poker.Card{
		{Rank: poker.Five, Suit: poker.Clubs},
		{Rank: poker.Ace, Suit: poker.Hearts},
		{Rank: poker.Queen, Suit: poker.Clubs},
		{Rank: poker.Two, Suit: poker.Clubs},
		{Rank: poker.Four, Suit: poker.Clubs},
	}

	fmt.Println("Test Cards:")

	fmt.Println(poker.PrettyCardsString(testCards1[:]))
	handValue1 := poker.EvaluateBestFiveCardHand(testCards1)
	fmt.Println(handValue1.Category.String())
	fmt.Println()

	fmt.Println(poker.PrettyCardsString(testCards2[:]))
	handValue2 := poker.EvaluateBestFiveCardHand(testCards2)
	fmt.Println(handValue2.Category.String())
	fmt.Println()

}
