package poker

import "fmt"

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
		return "♠\uFE0E"
	case Clubs:
		return "♣\uFE0E"
	case Diamonds:
		return "♦\uFE0E"
	case Hearts:
		return "❤\uFE0E"
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

func (card Card) PrettyRow(row int) string {
	switch row {
	case 0:
		return "┌─────┐"
	case 1:
		return fmt.Sprintf("┊ %-2v%v ┊", card.Rank, card.Suit)
	case 2:
		return fmt.Sprintf("┊ %v%2v ┊", card.Suit, card.Rank)
	case 3:
		return "└─────┘"
	default:
		return ""
	}
}

func (card Card) PrettyCardString() string {
	result := ""

	for row := 0; row < 4; row++ {
		result += card.PrettyRow(row)

		if row < 3 {
			result += "\n"
		}
	}
	return result
}

func PrettyCardsString(cards []Card) string {
	result := ""
	for row := 0; row < 4; row++ {
		for _, card := range cards {
			result += card.PrettyRow(row)
			result += " "
		}

		if row < 3 {
			result += "\n"
		}
	}
	return result
}
