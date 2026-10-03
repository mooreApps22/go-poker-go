package poker

func (hand *Hand) ResetFoldedHands() {
	for _, player := range hand.players {
		player.hasFolded = false
	}
}

func (hand *Hand) PostBlinds() {
	hand.pot.BuildBet(SmallBlindBet, hand.players[hand.smallBlindIndex])
	hand.pot.BuildBet(BigBlindBet, hand.players[hand.bigBlindIndex])
}

func (hand Hand) GetPot() Pot {
	return hand.pot
}
