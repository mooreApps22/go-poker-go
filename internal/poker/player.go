package poker

import "fmt"

type Player struct {
	ID            int
	Name          string
	Chips         int64
	HoleCards     [2]Card
	BestHandValue HandValue
	BestCards     [5]Card
	CurrentBet    int64
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

func (player *Player) PlaceBet(betAmount int64) int64 {
	player.Chips -= betAmount
	player.CurrentBet += betAmount
	return betAmount
}

func (player *Player) CollectWinnings(winnings int64) {
	player.Chips += winnings
}

func (player *Player) ResetCurrentBet() {
	player.CurrentBet = 0
}

func (player *Player) GetCurrentBet() int64 {
	return player.CurrentBet
}

func (player Player) GetName() string {
	return player.Name
}
