package poker

func (hand Hand) EvaluateEachPlayersBestHand() {
	for _, player := range hand.players {
		player.BestHandValue = FindBestHandValue(hand.GetCommunityCards(), player.HoleCards)
	}
}
