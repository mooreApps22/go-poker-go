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

func (server *Server) processClientMessages(clientMessageChannel chan ClientMessage) {
	for clientMessage := range clientMessageChannel {
		if server.gameHost.Started() {
			server.processPokerClientMessage(clientMessage)
		} else {
			server.processLobbyClientMessage(clientMessage)
		}
	}
}

func (server *Server) processPokerClientMessage(clientMessage ClientMessage) {
	playerAction, parseErr := parsePokerClientMessage(clientMessage)

	response := buildActionResponse(playerAction, parseErr)

	sendErr := sendClientResponse(clientMessage.client, response)
	if sendErr != nil {
		logClientError(
			clientMessage.client,
			"Error sending client response:",
			sendErr,
		)
		return
	}

	logPlayerAction(
		clientMessage.client,
		playerAction,
		parseErr,
	)
}

func (server *Server) processLobbyClientMessage(clientMessage ClientMessage) {
}
