package poker

type Hand struct {
	deck                Deck
	players             []*Player
	communityCards      [5]Card
	communityCardsDealt int
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

func (hand *Hand) DealFlopCards() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	for flopCardIndex := 0; flopCardIndex < 3; flopCardIndex++ {
		card, err := hand.deck.Draw()
		if err != nil {
			return err
		}
		hand.communityCards[flopCardIndex] = card
	}
	hand.communityCardsDealt += 3
	return nil
}

func (hand Hand) CommunityCards() []Card {
	return hand.communityCards[:hand.communityCardsDealt]
}
