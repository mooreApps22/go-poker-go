package poker

type Pot struct {
	Value int64
}

func NewPot() Pot {
	return Pot{
		Value: 0,
	}
}

func (pot *Pot) BuildBet(betValue int64, player *Player) {
	pot.Value += player.PlaceBet(betValue)
}

func (pot Pot) GetValue() int64 {
	return pot.Value
}
