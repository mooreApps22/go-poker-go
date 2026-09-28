package poker

import "fmt"

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

func (category HandRank) String() string {
	switch category {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of A Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of A Kind"
	case StraightFlush:
		return "Straight Flush"
	default:
		return "?"
	}
}

func findRankCounts(cards [5]Card) [15]int {
	var rankCounts [15]int
	for _, card := range cards {
		rankCounts[card.Rank] += 1
	}
	return rankCounts
}

func findHighCards(rankCounts [15]int) [5]Rank {
	var highCards [5]Rank
	highCardsIndex := 0
	for rank := Ace; rank >= Two; rank-- {
		if rankCounts[rank] > 0 {
			for repeatIndex := 0; repeatIndex < rankCounts[rank]; repeatIndex++ {
				highCards[highCardsIndex] = rank
				highCardsIndex++

				if highCardsIndex == len(highCards) {
					return highCards
				}
			}
		}
	}
	return highCards
}

func EvaluateBestFiveCardHand(cards [5]Card) HandValue {
	var handValue HandValue
	isFlush := checkIfFlush(cards)
	rankCounts := findRankCounts(cards)
	handValue.Tiebreakers = findHighCards(rankCounts)
	isStraight, isWheel := checkIfStraight(rankCounts)
	//Special Wheel Case Fix
	if isWheel {
		specialWheelStraightTieBreakerRearrangement(&handValue.Tiebreakers)
	}

	fmt.Println("High Cards: ", handValue.Tiebreakers)

	if isFlush && isStraight {
		handValue.Category = StraightFlush
	} else if checkForFourOfAKind(rankCounts) {
		handValue.Category = FourOfAKind
	} else if checkForFullHouse(rankCounts) {
		handValue.Category = FullHouse
	} else if isFlush {
		handValue.Category = Flush
	} else if isStraight {
		handValue.Category = Straight
	} else if checkForThreeOfAKind(rankCounts) {
		handValue.Category = ThreeOfAKind
	} else if checkForTwoPair(rankCounts) {
		handValue.Category = TwoPair
	} else if checkForOnePair(rankCounts) {
		handValue.Category = OnePair
	} else {
		handValue.Category = HighCard
	}

	return handValue
}
