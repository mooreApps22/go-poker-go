package server

import (
	"bufio"
	"fmt"
	"net"
	"time"
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

func handleClientConnection(client *Client, actionChannel chan ClientMessage) {
	defer client.connection.Close()

	reader := bufio.NewReader(client.connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf(
				"[%s] Client %d disconnected\n",
				time.Now().Format(time.TimeOnly),
				client.id,
			)
			return
		}

		actionChannel <- ClientMessage{
			client:  client,
			message: message,
		}
	}
}

func sendClientResponse(client *Client, response string) error {
	_, err := client.connection.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error writing to connection: %v\n", err)
		return err
	}
	return nil
}
