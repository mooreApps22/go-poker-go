package host

import (
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

type LobbyActionType uint8

const (
	InvalidLobbyAction LobbyActionType = iota
	SetExpectedPlayers
	StartGame
)

type LobbyAction struct {
	ClientID int
	Action   LobbyActionType
	Amount   int
}

type Host struct {
	expectedNumberOfPlayers int
	started                 bool
	players                 []*poker.Player
	lobbyOwner              *poker.Player
	game                    *poker.Game
}

func NewHost() Host {
	return Host{}
}

func (host *Host) ClientJoined(clientID int) string {
	return "Enter your name:\n"
}

func (host *Host) Started() bool {
	return host.started
}

func (host *Host) ClientHasPlayer(id int) bool {
	for _, player := range host.players {
		if player.GetID() == id {
			return true
		}
	}
	return false
}

func (host *Host) RegisterPlayer(clientID int, name string) string {
	player := poker.NewPlayer(clientID, name)

	if len(host.players) == 0 {
		host.lobbyOwner = &player
		host.players = append(host.players, &player)
		return "How many players should join? (4-10)\n"
	}

	host.players = append(host.players, &player)
	return "Waiting for the lobby owner to start the game.\n"
}

func (host *Host) IsLobbyOwner(player *poker.Player) bool {
	if player == nil || host.lobbyOwner == nil {
		return false
	}
	return player == host.lobbyOwner
}

func (host *Host) GetPlayer(clientID int) *poker.Player {
	for _, player := range host.players {
		if player.GetID() == clientID {
			return player
		}
	}
	return nil
}

func (host *Host) GetExpectedNumberOfPlayers() int {
	return host.expectedNumberOfPlayers
}

func (host *Host) SetExpectedNumberOfPlayers(number int) bool {
	if number < 4 || number > 10 {
		return false
	}
	host.expectedNumberOfPlayers = number
	return true
}

func (host *Host) StartGame() bool {
	if len(host.players) < 4 {
		return false
	}

	game := poker.NewGame(host.players)

	host.game = &game
	host.started = true
	return true
}

func (host *Host) AutomaticStartPossible() bool {
	return len(host.players) == host.expectedNumberOfPlayers
}
