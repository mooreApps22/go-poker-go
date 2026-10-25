package poker

type ActionResult struct {
	Action            ActionType
	Player            *Player
	Amount            int64
	BettingRoundEnded bool
	Message           string
}
