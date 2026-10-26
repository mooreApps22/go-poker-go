package server

import (
	"fmt"

	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func sendClientResponse(client *Client, response string) error {
	_, err := client.connection.Write([]byte(response))
	if err != nil {
		fmt.Printf("Error writing to connection: %v\n", err)
		return err
	}
	return nil
}

func (server *Server) broadcast(message string) {
	for _, client := range server.clients {
		sendErr := sendClientResponse(client, message)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}

func (server *Server) broadcastActionResult(result poker.ActionResult, handleMessage string) {

	for _, client := range server.clients {
		var sendErr error
		var response string
		if client.id == result.Player.GetID() {
			response = fmt.Sprintf("You %s", handleMessage)
		} else {
			response = fmt.Sprintf("%s %s", result.Player.GetName(), handleMessage)
		}
		sendErr = sendClientResponse(client, response)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}

func (server *Server) broadcastTurn(currentPlayer *poker.Player) {
	messageToOther := fmt.Sprintf("It's %s's turn.\n", currentPlayer.GetName())

	for _, client := range server.clients {
		var sendErr error
		if client.id == currentPlayer.GetID() {
			sendErr = sendClientResponse(client, "It's your turn.\n")
		} else {
			sendErr = sendClientResponse(client, messageToOther)
		}
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}

func (server *Server) sendPrivateHoleCards() {
	for _, client := range server.clients {
		player := server.gameHost.GetPlayer(client.id)
		holeCards := poker.FormatPrivateCards(player.GetHoleCards())
		holeCardsResponse := fmt.Sprintf("%s\n", holeCards)
		sendErr := sendClientResponse(client, holeCardsResponse)
		if sendErr != nil {
			logClientError(
				client,
				"Error broadcasting response",
				sendErr,
			)
		}
	}
}

func (server *Server) broadcastShowdown(hand *poker.Hand) {
	server.broadcast("SHOWDOWN\n")
	hand.CreatePots()
	hand.PickWinners()
	hand.AwardPots()

	server.broadcastPots(hand)
	server.broadcastWinners(hand)
}

func (server *Server) broadcastPots(hand *poker.Hand) {
	for potIndex, pot := range hand.GetPots() {
		var response string

		if potIndex == 0 {
			response = fmt.Sprintf(
				"Main Pot: $%d\n",
				pot.GetValue(),
			)
		} else {
			response = fmt.Sprintf(
				"Side Pot %c: $%d\n",
				'A'+rune(potIndex-1),
				pot.GetValue(),
			)
		}

		server.broadcast(response)
		/*
			eligiblePlayers := pot.GetEligiblePlayers()

			response = "Eligible: "

			for playerIndex, player := range eligiblePlayers {
				if playerIndex > 0 {
					response += ", "
				}

				response += player.GetName()
			}

			response += "\n"

			server.broadcast(response)
		*/
	}
}

func (server *Server) broadcastWinners(hand *poker.Hand) {
	for potIndex, pot := range hand.GetPots() {
		winners := pot.GetWinners()

		if len(winners) == 0 {
			continue
		}

		winnings := pot.GetValue() / int64(len(winners))

		var potName string

		if potIndex == 0 {
			potName = "Main Pot"
		} else {
			potName = fmt.Sprintf(
				"Side Pot %c",
				'A'+rune(potIndex-1),
			)
		}

		for _, winner := range winners {
			response := fmt.Sprintf(
				"%s's %s wins $%d from the %s.\n",
				winner.GetName(),
				winner.GetBestHandValue().Category.String(),
				winnings,
				potName,
			)

			server.broadcast(response)
			cardResponse := fmt.Sprintf(
				"%s\n",
				poker.FormatPrivateCards(winner.GetBestCards()))
			server.broadcast(cardResponse)
		}
	}
}
