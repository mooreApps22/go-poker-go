package main

import (
	"bufio"
	"fmt"
	"net"
	"os"

	"github.com/mooreApps22/go-poker-go/internal/client"
)

func readUserInput(connection net.Conn) {
	inputReader := bufio.NewReader(os.Stdin)

	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		_, err = connection.Write([]byte(input))
		if err != nil {
			fmt.Println("Error sending message:", err)
			return
		}
	}
}

func readServerMessages(connection net.Conn) {
	serverReader := bufio.NewReader(connection)

	for {
		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading from server:", err)
			return
		}

		fmt.Print("Server - ", response)
	}
}

func main() {
	connection, err := client.CreateDial("localhost", 8080)
	if err != nil {
		fmt.Println("Client connection error", err)
		return
	}

	defer connection.Close()
	fmt.Println("Connected to Go Poker Go Server...")

	go readServerMessages(connection)

	readUserInput(connection)
}
