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

	go processClientMessages(server.actionChannel)

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

		logClientConnected(client)

		go receiveClientMessages(client, server.actionChannel)
	}
}
