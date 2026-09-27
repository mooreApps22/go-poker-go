package poker

type HandPhase uint8

const (
	SetUp HandPhase = iota
	PreFlop
	Flop
	Turn
	River
	Showdown
)

type Hand struct {
	deck                Deck
	players             []*Player
	communityCards      [5]Card
	communityCardsDealt int
	phase               HandPhase
}

func NewHand(players []*Player) Hand {
	deck := NewDeck()
	deck.Shuffle()

	return Hand{
		deck:    deck,
		players: players,
		phase:   SetUp,
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
	hand.phase = PreFlop
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
	hand.phase = Flop
	return nil
}

func (hand Hand) CommunityCards() []Card {
	return hand.communityCards[:hand.communityCardsDealt]
}

func (hand *Hand) DealTurnCard() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	turnCard, err := hand.deck.Draw()
	if err != nil {
		return err
	}
	hand.communityCards[3] = turnCard
	hand.communityCardsDealt += 1
	hand.phase = Turn
	return nil
}

func (hand *Hand) DealRiverCard() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	riverCard, err := hand.deck.Draw()
	if err != nil {
		return err
	}
	hand.communityCards[4] = riverCard
	hand.communityCardsDealt += 1
	hand.phase = River
	return nil
}
