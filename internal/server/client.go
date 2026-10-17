package server

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	id         int
	connection net.Conn
}

type ClientMessage struct {
	clientID int
	message  string
}

func NewClient(id int, connection net.Conn) *Client {
	return &Client{
		id:         id,
		connection: connection,
	}
}

func handleClient(client *Client, actionChannel chan ClientMessage) {
	defer client.connection.Close()

	reader := bufio.NewReader(client.connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Client disconnected:", client.id)
			return
		}

		actionChannel <- ClientMessage{
			clientID: client.id,
			message:  message,
		}
	}
}
