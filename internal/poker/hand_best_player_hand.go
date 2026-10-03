package poker

func (hand Hand) EvaluateEachPlayersBestHand() {
	for _, player := range hand.players {
		player.bestHandValue, player.bestCards = findBestHandValue(
			hand.GetCommunityCards(),
			player.holeCards,
		)
	}
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
