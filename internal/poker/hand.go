package poker

type Hand struct {
	deck                Deck
	players             []*Player
	communityCards      [5]Card
	communityCardsDealt int
	phase               HandPhase
	winners             []*Player
	dealerIndex         int
	smallBlindIndex     int
	bigBlindIndex       int
	utgIndex            int // utg -> Under The Gun
	pot                 Pot
	currentCall         int64
}

type HandPhase uint8

const (
	SetUp HandPhase = iota
	PreFlop
	Flop
	Turn
	River
	Showdown
)

type HandRank uint8

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type HandValue struct {
	Category    HandRank
	Tiebreakers [5]Rank
}

func NewHand(players []*Player) Hand {
	deck := NewDeck()
	deck.Shuffle()

	return Hand{
		deck:    deck,
		players: players,
		phase:   SetUp,
		pot:     NewPot(),
	}
}
