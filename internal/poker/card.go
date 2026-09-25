package poker

type Rank uint8

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (rank Rank) String() string {
	switch rank {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "?"
	}
}

type Suit uint8

const (
	Spades Suit = iota
	Clubs
	Diamonds
	Hearts
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠"
	case Hearts:
		return "♥"
	case Clubs:
		return "♣"
	case Diamonds:
		return "♦"
	default:
		return "?"
	}
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return "[" + card.Rank.String() + card.Suit.String() + "]"
}
