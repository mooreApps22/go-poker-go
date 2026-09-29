package poker

func higherCategory(handValueA HandValue, handValueB HandValue) HandValue {

	if handValueA.Category > handValueB.Category {
		return handValueA
	}
	return handValueB
}

func compareHandValueTiebreakers(handValueA HandValue, handValueB HandValue) (
	HandValue, bool) {
	tiebreakerA := handValueA.Tiebreakers
	tiebreakerB := handValueB.Tiebreakers

	for index := 0; index < len(tiebreakerA); index++ {
		if tiebreakerA[index] != tiebreakerB[index] {
			if tiebreakerA[index] > tiebreakerB[index] {
				return handValueA, false
			}
			return handValueB, false
		}
	}
	return handValueA, true
}

func FindWinningHandValue(handValueA HandValue, handValueB HandValue) (
	HandValue, bool) {
	if handValueA.Category != handValueB.Category {
		return higherCategory(handValueA, handValueB), false
	}

	return compareHandValueTiebreakers(handValueA, handValueB)
}

func EvaluateBestFiveCardHand(cards [5]Card) HandValue {
	var handValue HandValue
	isFlush := checkIfFlush(cards)
	rankCounts := findRankCounts(cards)
	handValue.Tiebreakers = findHighCards(rankCounts)
	isStraight, isWheel := checkIfStraight(rankCounts)
	// Check if Straight is a Wheel Straight --> A, 2, 3, 4, 5
	if isWheel {
		specialWheelStraightTieBreakerRearrangement(&handValue.Tiebreakers)
	}

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

	arrangeTieBreakers(&handValue, rankCounts)

	return handValue
}
