package poker

func findHighestHandValue(possibleHands [21]HandValue) HandValue {
	maxHandValue := possibleHands[0]

	for index := 1; index < len(possibleHands); index++ {
		maxHandValue, _ = FindWinningHandValue(maxHandValue, possibleHands[index])
	}

	return maxHandValue
}

func FindBestHandValue(communityCards [5]Card, holeCards [2]Card) HandValue {
	var possibleHands [21]HandValue
	possibleHandsIndex := 0

	var availableCards [7]Card
	availableCards[0] = holeCards[0]
	availableCards[1] = holeCards[1]

	for index := 0; index < len(communityCards); index++ {
		availableCards[index+2] = communityCards[index]
	}

	for firstExcluded := 0; firstExcluded < len(availableCards)-1; firstExcluded++ {
		for secondExcluded := firstExcluded + 1; secondExcluded < 7; secondExcluded++ {
			var possibleHand [5]Card
			possibleHandIndex := 0

			for cardIndex := 0; cardIndex < len(availableCards); cardIndex++ {
				if cardIndex != firstExcluded && cardIndex != secondExcluded {
					possibleHand[possibleHandIndex] = availableCards[cardIndex]
					possibleHandIndex++
				}
			}

			possibleHands[possibleHandsIndex] = EvaluateBestFiveCardHand(possibleHand)
			possibleHandsIndex++
		}
	}

	return findHighestHandValue(possibleHands)
}
