package poker

import (
	"fmt"
	"math/rand"
	"time"
)

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

func (deck Deck) Print() {
	const cardsPerGroup = 7
	for start := 0; start < len(deck.cards); start += cardsPerGroup {
		end := start + cardsPerGroup

		if end > len(deck.cards) {
			end = len(deck.cards)
		}
		fmt.Println(PrettyCardsString((deck.cards[start:end])))
	}
}

func (deck *Deck) Shuffle() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(deck.cards) - 1; i > 0; i-- {
		randomIndex := random.Intn(i + 1)

		deck.cards[i], deck.cards[randomIndex] =
			deck.cards[randomIndex], deck.cards[i]
	}
}
