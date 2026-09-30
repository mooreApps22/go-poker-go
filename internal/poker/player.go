package poker

import "fmt"

type Player struct {
	ID            int
	Name          string
	Chips         int64
	HoleCards     [2]Card
	BestHandValue HandValue
	BestCards     [5]Card
}

func NewPlayer(id int, name string) Player {
	return Player{
		ID:    id,
		Name:  name,
		Chips: 1_000_000,
	}
}

func (player Player) String() string {
	return fmt.Sprintf("%v[%v] — Chips: $%v", player.Name, player.ID, player.Chips)
}

func (player *Player) GetBestCards() []Card {
	return player.BestCards[:]
}
