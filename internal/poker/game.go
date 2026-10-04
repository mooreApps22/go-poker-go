package poker

import "fmt"

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

func (game *Game) removePlayersWithNoChips() {
	activePlayers := game.players[:0]

	for _, player := range game.players {
		if player.chips > 0 {
			activePlayers = append(activePlayers, player)
		} else {
			fmt.Printf("%v has left the poker table, due to losing all chips.\n", player.name)
		}
	}

	game.players = activePlayers
}

func (game *Game) SetUpNewHand() {
	game.removePlayersWithNoChips()

	game.hand = NewHand(game.players)

	game.hand.dealerIndex = (game.dealerIndex + 1) % len(game.players)
	game.hand.smallBlindIndex = (game.dealerIndex + 2) % len(game.players)
	game.hand.bigBlindIndex = (game.dealerIndex + 3) % len(game.players)
	game.hand.utgIndex = (game.dealerIndex + 4) % len(game.players)
	game.hand.currentPlayerIndex = game.hand.utgIndex
}

func (game *Game) GetHand() *Hand {
	return &game.hand
}

func (game *Game) RotateDealer() {
	game.dealerIndex = (game.dealerIndex + 1) % len(game.players)
}
