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

func (server *Server) Run() {

	server.nextClientID = 1

	go handleActions(server.actionChannel)

	for {
		connection, err := server.listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		client := NewClient(server.nextClientID, connection)
		server.nextClientID++

		fmt.Println("Client connected:", connection.RemoteAddr())

		go handleClientConnection(client, server.actionChannel)
	}
}

func handleActions(actionChannel chan ClientMessage) {
	for clientMessage := range actionChannel {
		playerAction, err := parseClientMessage(clientMessage)

		var response string

		if err != nil {
			response = fmt.Sprintf("Invalid action: %s\n", err)
		} else if playerAction.Amount > 0 {
			response = fmt.Sprintf(
				"Action received: %s $%d\n",
				playerAction.Action.String(),
				playerAction.Amount,
			)
		} else {
			response = fmt.Sprintf(
				"Action received: %s\n",
				playerAction.Action.String(),
			)
		}

		err = sendClientResponse(clientMessage.client, response)
		if err != nil {
			fmt.Println("Error sending client response:", err)
		}
	}
}
