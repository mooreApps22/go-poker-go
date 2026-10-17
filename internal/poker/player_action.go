package poker

type ActionType uint8

const (
	Invalid ActionType = iota
	Fold
	Check
	Call
	Bet
	Raise
	AllIn
	Chips
)

type PlayerAction struct {
	PlayerID int
	Action   ActionType
	Amount   int64
}
