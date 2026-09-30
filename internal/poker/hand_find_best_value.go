package poker

func findHighestHandValue(possibleHandValues [21]HandValue, possibleHands [21][5]Card) (
	HandValue, [5]Card) {
	maxHandValue := possibleHandValues[0]
	maxHandIndex := 0

	for index := 1; index < len(possibleHandValues); index++ {
		winningHandValue, isTied := FindWinningHandValue(
			maxHandValue,
			possibleHandValues[index],
		)

		if !isTied && winningHandValue == possibleHandValues[index] {
			maxHandValue = winningHandValue
			maxHandIndex = index
		}
	}

	return maxHandValue, possibleHands[maxHandIndex]
}

func FindBestHandValue(communityCards [5]Card, holeCards [2]Card) (
	HandValue, [5]Card) {
	var possibleHandValues [21]HandValue
	var possibleHands [21][5]Card
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

			possibleHands[possibleHandsIndex] = possibleHand
			possibleHandValues[possibleHandsIndex] = EvaluateBestFiveCardHand(possibleHand)
			possibleHandsIndex++
		}
	}

	return findHighestHandValue(possibleHandValues, possibleHands)
}
