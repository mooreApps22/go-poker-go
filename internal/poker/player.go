package poker

import "fmt"

type Player struct {
	ID        int
	Name      string
	Chips     int64
	HoleCards [2]Card
}

func NewPlayer(id int, name string, chips int64) Player {
	return Player{
		ID:    id,
		Name:  name,
		Chips: chips,
	}
}

func (player Player) String() string {
	return fmt.Sprintf("%v[%v] — Chips: $%v", player.Name, player.ID, player.Chips)
}
