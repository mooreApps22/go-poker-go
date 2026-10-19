package server

import (
	"fmt"
	"net"
	"time"
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

		fmt.Printf(
			"[%s] Client %d connected: %s\n",
			time.Now().Format(time.TimeOnly),
			client.id,
			connection.RemoteAddr(),
		)

		go handleClientConnection(client, server.actionChannel)
	}
}

func handleActions(actionChannel chan ClientMessage) {
	for clientMessage := range actionChannel {
		playerAction, parseErr := parseClientMessage(clientMessage)

		var response string

		if parseErr != nil {
			response = fmt.Sprintf("Invalid action: %s\n", parseErr)
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

		sendErr := sendClientResponse(clientMessage.client, response)
		if sendErr != nil {
			fmt.Println("Error sending client response:", sendErr)
		}
		fmt.Printf(
			"Client %d: %s",
			clientMessage.client.id,
			playerAction.Action.String(),
		)
		if playerAction.Amount > 0 {
			fmt.Printf(
				" %d ",
				playerAction.Amount,
			)
		}
		if parseErr != nil {
			fmt.Printf(
				" Error: %s",
				parseErr,
			)
		}
		fmt.Printf("\n")
	}
}
