package server

func (server *Server) startNextHand() {
	err := server.gameHost.StartHand()
	if err != nil {
		return
	}
	//broadcast public hand state
	// privately send hole cards
	// tell everyone whose turn it is
}
