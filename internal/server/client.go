package server

import (
	"bufio"
	"net"
)

type Client struct {
	id         int
	connection net.Conn
}

type ClientMessage struct {
	client  *Client
	message string
}

func NewClient(id int, connection net.Conn) *Client {
	return &Client{
		id:         id,
		connection: connection,
	}
}

func (server *Server) receiveClientMessages(client *Client, clientMessageChannel chan ClientMessage) {
	defer client.connection.Close()

	reader := bufio.NewReader(client.connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logClientDisconnected(client)
			return
		}

		clientMessageChannel <- ClientMessage{
			client:  client,
			message: message,
		}
	}
}
