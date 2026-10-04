package poker

import "fmt"

func (hand Hand) ShowDown() {
	fmt.Println("SHOWDOWN PHASE: ")

	for _, player := range hand.players {
		player.bestHandValue, player.bestCards = findBestHandValue(
			hand.GetCommunityCards(),
			player.holeCards,
		)
	}
}

func findBestHandValue(communityCards [5]Card, holeCards [2]Card) (
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

func comparePlayersHandValues(player1 *Player, player2 *Player) (*Player, bool) {
	winningHandValue, isTied := findWinningHandValue(
		player1.bestHandValue,
		player2.bestHandValue,
	)

	if isTied {
		return player1, true
	}

	if winningHandValue == player1.bestHandValue {
		return player1, false
	}

	return player2, false
}
