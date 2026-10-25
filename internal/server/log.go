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
	result poker.ActionResult,
) {
	if result.Amount > 0 {
		fmt.Printf(
			"[%s] Client %d: %s $%d\n",
			my_timestamp(),
			client.id,
			result.Action.String(),
			result.Amount,
		)
		return
	}

	fmt.Printf(
		"[%s] Client %d: %s\n",
		my_timestamp(),
		client.id,
		result.Action.String(),
	)
}
