package server

import (
	"fmt"
	"net"
)

type Server struct {
	listener      net.Listener
	actionChannel chan ClientMessage
	nextClientID  int
}

func NewServer(port int) (Server, error) {
	listener, err := createListener(8080)
	if err != nil {
		fmt.Println("Error creating listener:", err)
		return Server{}, err
	}

	return Server{
		listener:      listener,
		actionChannel: make(chan ClientMessage),
		nextClientID:  1,
	}, nil
}

func createListener(port int) (net.Listener, error) {
	portString := fmt.Sprintf(":%d", port)

	return net.Listen("tcp", portString)
}

func handleActions(actionChannel chan ClientMessage) {
	for action := range actionChannel {
		fmt.Printf(
			"Client %d: %s",
			action.clientID,
			action.message,
		)
	}
}

func (server *Server) Run() {

	nextClientID := 1

	go handleActions(server.actionChannel)

	for {
		connection, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		client := NewClient(nextClientID, connection)
		nextClientID++

		fmt.Println("Client connected:", connection.RemoteAddr())

		go handleClient(client, server.actionChannel)
	}
}
