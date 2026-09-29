package poker

type Game struct {
	players     []*Player
	dealerIndex int
	hand        Hand
}

func NewGame(players []*Player) Game {
	return Game{
		players: players,
	}
}

func (game *Game) NewHand() {
	game.hand = NewHand(game.players)
}

func (game *Game) GetHand() *Hand {
	return &game.hand
}

func (game *Game) RotateDealer() {
	game.dealerIndex = (game.dealerIndex + 1) % len(game.players)
}
