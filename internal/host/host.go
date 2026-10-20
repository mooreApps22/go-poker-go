package host

type Host struct {
	expectedNumberOfPlayers int
	started                 bool
	// players
	// lobbyOwner
	// game
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
