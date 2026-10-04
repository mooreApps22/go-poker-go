package poker

func (hand *Hand) ResetFoldedHands() {
	for _, player := range hand.players {
		player.hasFolded = false
	}
}

func (hand *Hand) PostBlinds() {
	hand.players[hand.smallBlindIndex].PlaceBet(SmallBlindBet)
	hand.players[hand.bigBlindIndex].PlaceBet(BigBlindBet)
}
