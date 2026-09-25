package poker

import "fmt"

type Deck struct {
	cards [52]Card
}

func NewDeck() Deck {
	var deck Deck
	index := 0

	for suit := Spades; suit <= Hearts; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			card := Card{
				Rank: rank,
				Suit: suit,
			}
			deck.cards[index] = card
			index++
		}
	}
	return deck
}

func (deck Deck) PrintDeck() {
	for index := 0; index < 52; index++ {
		fmt.Println(deck.cards[index])
	}
}
