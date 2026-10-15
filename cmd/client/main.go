package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/mooreApps22/go-poker-go/internal/client"
)

func main() {
	connection, err := client.CreateDial("localhost", 8080)
	if err != nil {
		fmt.Println("Client connection error", err)
		return
	}

	defer connection.Close()
	fmt.Println("Connected to Go Poker Go Server...")

	inputReader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
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
