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
	hand.winners = nil

	for _, player := range hand.players {
		if !player.hasFolded {
			hand.winners = []*Player{player}
			break
		}
	}

	if len(hand.winners) == 0 {
		return
	}

	for _, player := range hand.players {
		if player.hasFolded || player == hand.winners[0] {
			continue
		}

		winningPlayer, isTied := comparePlayersHandValues(
			hand.winners[0],
			player,
		)

		if isTied {
			hand.winners = append(hand.winners, player)
		} else if winningPlayer == player {
			hand.winners = []*Player{player}
		}
	}
}
