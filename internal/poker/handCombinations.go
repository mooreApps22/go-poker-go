package poker

type HandRank uint8

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type HandValue struct {
	Rank        HandRank
	Tiebreakers [5]Rank
}

//func EvaluateHand(holeCards [2]Card, communityCards []Card) HandValue {
