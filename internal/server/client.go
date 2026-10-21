package server

import (
	"bufio"
	"fmt"
	"net"
	"strconv"
	"strings"
	// "github.com/mooreApps22/go-poker-go/internal/poker"
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
	if !server.gameHost.ClientHasPlayer(clientMessage.client.id) {
		name := strings.TrimSpace(clientMessage.message)

		response := server.gameHost.RegisterPlayer(
			clientMessage.client.id,
			name,
		)
		sendErr := sendClientResponse(clientMessage.client, response)
		if sendErr != nil {
			logClientError(
				clientMessage.client,
				"Error sending lobby response",
				sendErr,
			)
		}
		return
	}

	currentPlayer := server.gameHost.GetPlayer(clientMessage.client.id)
	if currentPlayer == nil {

		return
	}

	if server.gameHost.IsLobbyOwner(currentPlayer) {
		if server.gameHost.GetExpectedNumberOfPlayers() == 0 {
			number, err := strconv.Atoi(strings.TrimSpace(clientMessage.message))
			if err != nil {
				sendErr := sendClientResponse(
					clientMessage.client,
					"Expected a number.\n",
				)
				if sendErr != nil {
					logClientError(
						clientMessage.client,
						"Error sending lobby response",
						sendErr,
					)
				}
				return
			}

			if !server.gameHost.SetExpectedNumberOfPlayers(number) {
				sendErr := sendClientResponse(
					clientMessage.client,
					"Number can't be under 4 or over 10.\n",
				)
				if sendErr != nil {
					logClientError(
						clientMessage.client,
						"Error sending lobby response",
						sendErr,
					)
				}
				return
			}

			sendErr := sendClientResponse(
				clientMessage.client,
				fmt.Sprintf("Waiting for %d players.\n", number),
			)
			if sendErr != nil {
				logClientError(
					clientMessage.client,
					"Error sending lobby response",
					sendErr,
				)
			}
			return
		}
	}

}
