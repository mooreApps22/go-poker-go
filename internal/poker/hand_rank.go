package poker

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
