package poker

func findHighestHandValue(possibleHandValues [21]HandValue, possibleHands [21][5]Card) (
	HandValue, [5]Card) {
	maxHandValue := possibleHandValues[0]
	maxHandIndex := 0

	for index := 1; index < len(possibleHandValues); index++ {
		winningHandValue, isTied := findWinningHandValue(
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
