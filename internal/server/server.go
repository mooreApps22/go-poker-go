package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/host"
	"net"
)

type Server struct {
	listener             net.Listener
	clientMessageChannel chan ClientMessage
	nextClientID         int
	gameHost             *host.Host
	clients              []*Client
}

func NewServer(port int, gameHost *host.Host) (Server, error) {
	listener, err := createListener(port)
	if err != nil {
		fmt.Printf(
			"[%s] Error creating listener: %v\n",
			my_timestamp(),
			err,
		)
		return Server{}, err
	}

	return Server{
		listener:             listener,
		clientMessageChannel: make(chan ClientMessage),
		nextClientID:         1,
		gameHost:             gameHost,
		clients:              nil,
	}, nil
}

func createListener(port int) (net.Listener, error) {
	portString := fmt.Sprintf(":%d", port)

	return net.Listen("tcp", portString)
}

func (server *Server) Run() {

	go server.routeClientMessages(server.clientMessageChannel)

	for {
		connection, err := server.listener.Accept()
		if err != nil {
			logClientError(
				nil,
				"Error accepting connection",
				err,
			)
			continue
		}
		client := NewClient(server.nextClientID, connection)
		server.nextClientID++

		logClientConnected(client)
		server.clients = append(server.clients, client)

		response := server.gameHost.ClientJoined(client.id)
		sendErr := sendClientResponse(client, response)
		if sendErr != nil {
			logClientError(
				client,
				"Error at host prompting player name",
				sendErr,
			)
			client.connection.Close()
			continue
		}

		go server.receiveClientMessages(client, server.clientMessageChannel)
	}
}
