package poker

func (hand Hand) GetDeck() Deck {
	return hand.deck
}

func (hand Hand) GetPlayers() []*Player {
	return hand.players
}

func (hand Hand) GetCommunityCards() [5]Card {
	return hand.communityCards
}

func (hand Hand) GetCommunityCardsDealt() int {
	return hand.communityCardsDealt
}

func (hand Hand) GetWinner() []*Player {
	return hand.winners
}

func (hand Hand) GetSmallBlindIndex() int {
	return hand.smallBlindIndex
}

func (hand Hand) GetDealerIndex() int {
	return hand.dealerIndex
}

func (hand Hand) GetBigBlindIndex() int {
	return hand.bigBlindIndex
}

func (hand Hand) GetUtgIndex() int {
	return hand.utgIndex
}
