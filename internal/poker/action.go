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
	Hole
	Menu
)

type PlayerAction struct {
	PlayerID int
	Action   ActionType
	Amount   int64
}

type ActionResult struct {
	Action            ActionType
	Player            *Player
	Amount            int64
	BettingRoundEnded bool
	Message           string
	HandEnded         bool
}

func (action ActionType) String() string {
	switch action {
	case Fold:
		return "FOLD"
	case Check:
		return "CHECK"
	case Call:
		return "CALL"
	case Bet:
		return "BET"
	case Raise:
		return "RAISE"
	case AllIn:
		return "ALL IN"
	case Chips:
		return "CHIPS"
	case Hole:
		return "HOLE"
	case Menu:
		return "MENU"
	default:
		return "INVALID"
	}
}
