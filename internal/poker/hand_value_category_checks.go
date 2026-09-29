package poker

func checkForOnePair(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			return true
		}
	}
	return false
}

func checkForTwoPair(rankCounts [15]int) bool {
	pairCount := 0
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func checkForThreeOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 3 {
			return true
		}
	}
	return false
}

func checkForFullHouse(rankCounts [15]int) bool {
	pairCount := 0
	threeOfAKindCount := 0
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			pairCount++
		}
		if rankCounts[rank] == 3 {
			threeOfAKindCount++
		}
	}
	return pairCount == 1 && threeOfAKindCount == 1
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

func checkIfFlush(cards [5]Card) bool {
	suitSample := cards[0].Suit

	for _, card := range cards[1:] {
		if suitSample != card.Suit {
			return false
		}
	}
	return true
}

func checkForFourOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 4 {
			return true
		}
	}
	return false
}
