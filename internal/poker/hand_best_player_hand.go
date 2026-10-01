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

func (hand *Hand) PickWinners() {
	hand.winners = []*Player{hand.players[0]}
	for index := 1; index < len(hand.players); index++ {
		winningPlayer, isTied := comparePlayersHandValues(
			hand.winners[0],
			hand.players[index],
		)

		if isTied {
			hand.winners = append(hand.winners, hand.players[index])
		} else if winningPlayer == hand.players[index] {
			hand.winners = []*Player{hand.players[index]}
		}
	}
}
