package poker

func findOnePairRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			return rank, true
		}
	}
	return 0, false
}

// Returns the high pair and lower pair ranks in order
func findTwoPairsRank(rankCounts [15]int) (Rank, Rank, bool) {
	var pair1 Rank
	var pair2 Rank
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 && pair1 == 0 {
			pair1 = rank
		} else if rankCounts[rank] == 2 && pair2 == 0 {
			pair2 = rank
		}

		if pair1 != 0 && pair2 != 0 {
			if pair1 < pair2 {
				pair1, pair2 = pair2, pair1
			}
			return pair1, pair2, true
		}
	}

	return 0, 0, false
}

func findThreeOfAKindRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 3 {
			return rank, true
		}
	}
	return 0, false
}

func findFourOfAKindRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 4 {
			return rank, true
		}
	}
	return 0, false
}

func arrangeTieBreakers(handValue *HandValue, rankCounts [15]int) {
	switch handValue.Category {
	case HighCard:
		//FREE
		return
	case OnePair:
		var arranged [5]Rank

		pairRank, isFound := findOnePairRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = pairRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != pairRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case TwoPair:
		// pair1, pair2, kicker, doesn't matter, doesn't matter
		var arranged [5]Rank

		pair1Rank, pair2Rank, isFound := findTwoPairsRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = pair1Rank
		arranged[1] = pair2Rank
		arrangedIndex := 2

		for _, rank := range handValue.Tiebreakers {
			if rank != pair1Rank && rank != pair2Rank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case ThreeOfAKind:
		// trips, kicker1, kicker2, doesn't matter, doesn't matter
		var arranged [5]Rank

		tripsRank, isFound := findThreeOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = tripsRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != tripsRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case Straight:
		//FREE
		return
	case Flush:
		//FREE
		return
	case FullHouse:
		// trips, pair, doesn't matter, doesn't matter, doesn't matter
		var arranged [5]Rank

		tripsRank, isFound := findThreeOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		pairRank, isFound := findOnePairRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = tripsRank
		arranged[1] = pairRank

		handValue.Tiebreakers = arranged
		return
	case FourOfAKind:
		// quad, kicker, doesn't matter, doesn't matter, doesn't matter
		var arranged [5]Rank

		quadsRank, isFound := findFourOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = quadsRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != quadsRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case StraightFlush:
		//FREE
		return
	default:
		return
	}
}
