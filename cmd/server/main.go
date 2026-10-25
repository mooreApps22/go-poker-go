package main

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/host"
	"github.com/mooreApps22/go-poker-go/internal/server"
)

func main() {

	gameHost := host.NewHost()

	pokerServer, err := server.NewServer(8080, &gameHost)
	if err != nil {
		fmt.Println("Error creating server:", err)
		return
	}
	fmt.Println("Go Poker Go Server is listening...")

	pokerServer.Run()

	/*

		for {
			hand.ShowDown()

			hand.CreatePots()
			hand.PickWinners()
			hand.AwardPots()

			//DISPLAY WINNERS
			for _, player := range hand.GetWinner() {
				fmt.Printf("Winner: %v\n", player.GetName())
				fmt.Println(poker.FormatCards(player.GetBestCards()))
			}

			fmt.Println()
			fmt.Println("END OF HAND")
			fmt.Println()
		}
	*/
}
