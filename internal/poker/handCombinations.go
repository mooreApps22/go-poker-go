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

func findRankCounts(cards [5]Card) [15]int {
	var rankCounts [15]int
	for _, card := range cards {
		rankCounts[card.Rank] += 1
	}
	return rankCounts
}

func checkIfFlush(cards [5]Card) bool {
	suitSample := cards[0].Suit

	for _, card := range cards[1:] {
		if suitSample != card.Suit {
			return false
		}
	}
	return true
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

// A Wheel is when a Straight == A, 5, 4, 3, 2, but 5 is treated as the High Card
func specialWheelStraightTieBreakerRearrangement(ranks *[5]Rank) {
	ranks[0] = Five
	ranks[1] = Four
	ranks[2] = Three
	ranks[3] = Two
	ranks[4] = Ace
}

func checkIfStraight(rankCounts [15]int) (bool, bool) {

	if rankCounts[Ace] == 1 &&
		rankCounts[Two] == 1 &&
		rankCounts[Three] == 1 &&
		rankCounts[Four] == 1 &&
		rankCounts[Five] == 1 {
		return true, true
	}
	straightTally := 0
	straightTallyBegan := false
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 0 {
			if straightTallyBegan == false {
				continue
			} else {
				return false, false
			}
		} else if rankCounts[rank] > 1 {
			return false, false
		} else {
			straightTallyBegan = true
			straightTally += 1
		}
	}
	return straightTally == 5, false
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

	for rank := Two; rank <= Ace; rank++ {
		fmt.Println("Rank: ", rank, "Count: ", rankCounts[rank])
	}

	//Just for testing for now
	if isFlush {
		fmt.Println("This hand is FLUSH!!")
	} else {
		fmt.Println("This hand is NOT FLUSH!!")
	}
	if isStraight {
		fmt.Println("This hand is STRAIGHT!!")
	} else {
		fmt.Println("This hand is NOT STRAIGHT!!")
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
