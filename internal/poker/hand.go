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
	currentCall         int64
	currentPlayerIndex  int
	voluntaryBets       bool
	minimumBet          int64
	haveFoldedCount     int
	contributionLevels  []int64
	pots                []Pot
}

const BigBlindBet = 4_000
const SmallBlindBet = 2_000

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

	// only add players with chips > 0

	return Hand{
		deck:               deck,
		players:            players,
		phase:              SetUp,
		currentCall:        BigBlindBet,
		voluntaryBets:      false,
		minimumBet:         BigBlindBet,
		haveFoldedCount:    0,
		contributionLevels: nil,
		pots:               nil,
	}
}

func (hand *Hand) ResetPlayersForNewHand() {
	for _, player := range hand.players {
		player.ResetForNewHand()
	}
}
