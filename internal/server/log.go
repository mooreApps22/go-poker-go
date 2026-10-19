package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
	"time"
)

func my_timestamp() string {
	return time.Now().Format(time.TimeOnly)
}

func logClientConnected(client *Client) {
	fmt.Printf(
		"[%s] Client %d connected: %s\n",
		my_timestamp(),
		client.id,
		client.connection.RemoteAddr(),
	)
}

func logClientDisconnected(client *Client) {
	fmt.Printf(
		"[%s] Client %d disconnected\n",
		time.Now().Format(time.TimeOnly),
		client.id,
	)
}

func logClientError(client *Client, message string, err error) {
	fmt.Printf(
		"[%s] Client %d - %s: %v\n",
		my_timestamp(),
		client.id,
		message,
		err,
	)
}

func logPlayerAction(
	client *Client,
	playerAction poker.PlayerAction,
	parseErr error,
) {
	if parseErr != nil {
		logClientError(client, "parse action", parseErr)
		return
	}

	if playerAction.Amount > 0 {
		fmt.Printf(
			"Client %d: %s $%d\n",
			client.id,
			playerAction.Action.String(),
			playerAction.Amount,
		)
		return
	}

	fmt.Printf(
		"Client %d: %s\n",
		client.id,
		playerAction.Action.String(),
	)
}
