package poker

func (hand *Hand) PostBlinds() {
	hand.pot.BuildBet(SmallBlindBet, hand.players[hand.smallBlindIndex])
	hand.pot.BuildBet(BigBlindBet, hand.players[hand.bigBlindIndex])

	for _, player := range hand.players {
		player.hasFolded = false
	}
}

func (hand *Hand) AwardPot() {
	winnings := hand.pot.value / int64(len(hand.winners))

	for _, winner := range hand.winners {
		winner.CollectWinnings(winnings)
	}

	hand.pot.value = 0
	for _, player := range hand.players {
		player.ResetCurrentBet()
	}
}

func (hand Hand) GetPot() Pot {
	return hand.pot
}
