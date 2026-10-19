package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
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
		fmt.Printf(
			"[%s] Error creating listener: %v\n",
			my_timestamp(),
			err,
		)
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
			logClientError(
				nil,
				"Error creating listener",
				err,
			)
			continue
		}
		client := NewClient(server.nextClientID, connection)
		server.nextClientID++

		fmt.Printf(
			"[%s] Client %d connected: %s\n",
			my_timestamp(),
			client.id,
			connection.RemoteAddr(),
		)

		go handleClientConnection(client, server.actionChannel)
	}
}

func buildActionResponse(playerAction poker.PlayerAction, parseErr error) string {
	if parseErr != nil {
		return fmt.Sprintf("Invalid action: %s\n", parseErr)
	}

	if playerAction.Amount > 0 {
		return fmt.Sprintf(
			"Action received: %s $%d\n",
			playerAction.Action.String(),
			playerAction.Amount,
		)
	}

	return fmt.Sprintf(
		"Action received: %s\n",
		playerAction.Action.String(),
	)
}

func handleActions(actionChannel chan ClientMessage) {
	for clientMessage := range actionChannel {
		playerAction, parseErr := parseClientMessage(clientMessage)

		response := buildActionResponse(playerAction, parseErr)

		sendErr := sendClientResponse(clientMessage.client, response)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending client response:",
				sendErr,
			)
			continue
		}

		logPlayerAction(
			clientMessage.client,
			playerAction,
			parseErr,
		)
	}
}
