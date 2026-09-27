package poker

type Hand struct {
	deck      Deck
	players   []*Player
	community [5]Card
}

func NewHand(players []*Player) Hand {
	deck := NewDeck()
	deck.Shuffle()

	return Hand{
		deck:    deck,
		players: players,
	}
}

func (hand *Hand) DealHoleCards() error {
	for holeCardIndex := 0; holeCardIndex < 2; holeCardIndex++ {
		for _, player := range hand.players {
			card, err := hand.deck.Draw()
			if err != nil {
				return err
			}

			player.HoleCards[holeCardIndex] = card
		}
	}
	return nil
}
