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

func receiveClientMessages(client *Client, actionChannel chan ClientMessage) {
	defer client.connection.Close()

	reader := bufio.NewReader(client.connection)

	for {
		message, err := reader.ReadString('\n')
		if err != nil {
			logClientDisconnected(client)
			return
		}

		actionChannel <- ClientMessage{
			client:  client,
			message: message,
		}
	}
}

func processClientMessages(actionChannel chan ClientMessage) {
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
