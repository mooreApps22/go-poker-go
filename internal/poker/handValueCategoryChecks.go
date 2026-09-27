package poker

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

func checkForTwoPair(rankCounts [15]int) bool {
	pairCount := 0
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func checkForOnePair(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			return true
		}
	}
	return false
}

func checkForThreeOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 3 {
			return true
		}
	}
	return false
}

func checkForFourOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 4 {
			return true
		}
	}
	return false
}
