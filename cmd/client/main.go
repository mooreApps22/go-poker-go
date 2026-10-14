package main

import (
	"fmt"

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
}
