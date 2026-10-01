package poker

import "fmt"

type Player struct {
	id            int
	name          string
	chips         int64
	holeCards     [2]Card
	bestHandValue HandValue
	bestCards     [5]Card
	currentBet    int64
	hasFolded     bool
	hasChecked    bool
	hasActed      bool
}

func NewPlayer(id int, name string) Player {
	return Player{
		id:         id,
		name:       name,
		chips:      1_000_000,
		hasFolded:  false,
		hasChecked: false,
		hasActed:   false,
	}
}

func (player Player) String() string {
	return fmt.Sprintf("%v[%v] — Chips: $%v", player.name, player.id, player.chips)
}

func (player *Player) GetHoleCards() []Card {
	return player.holeCards[:]
}

func (player *Player) GetBestHandValue() HandValue {
	return player.bestHandValue
}

func (player *Player) GetBestCards() []Card {
	return player.bestCards[:]
}

func (player *Player) PlaceBet(betAmount int64) int64 {
	player.chips -= betAmount
	player.currentBet += betAmount

	return betAmount
}

func (player *Player) CollectWinnings(winnings int64) {
	player.chips += winnings
}

func (player *Player) ResetCurrentBet() {
	player.currentBet = 0
}

func (player *Player) GetCurrentBet() int64 {
	return player.currentBet
}

func (player Player) GetName() string {
	return player.name
}
