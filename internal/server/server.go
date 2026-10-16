package server

import (
	"fmt"
	"net"
)

func CreateListener(port int) (net.Listener, error) {
	portString := fmt.Sprintf(":%d", port)

	return net.Listen("tcp", portString)
}

func handleActions(actionChannel chan ClientMessage) {
	for action := range actionChannel {
		fmt.Printf(
			"Client %d: %s",
			action.clientID,
			action.message,
		)
	}
}

func Run(listener net.Listener, actionChannel chan ClientMessage) {

	nextClientID := 1

	go handleActions(actionChannel)

	for {
		connection, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		client := NewClient(nextClientID, connection)
		nextClientID++

		fmt.Println("Client connected:", connection.RemoteAddr())

		go handleClient(client, actionChannel)
	}
}
