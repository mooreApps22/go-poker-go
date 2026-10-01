package poker

type Pot struct {
	value int64
}

func NewPot() Pot {
	return Pot{
		value: 0,
	}
}

func (pot *Pot) BuildBet(betValue int64, player *Player) {
	pot.value += player.PlaceBet(betValue)
}

func (pot Pot) GetValue() int64 {
	return pot.value
}
