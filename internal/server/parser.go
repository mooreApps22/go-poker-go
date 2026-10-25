package server

import (
	"errors"
	"github.com/mooreApps22/go-poker-go/internal/poker"
	"strconv"
	"strings"
)

func parsePokerClientMessage(clientMessage ClientMessage) (poker.PlayerAction, error) {
	fields := strings.Fields(clientMessage.message)
	if len(fields) == 0 {
		return poker.PlayerAction{}, errors.New("empty player action")
	}

	var action poker.ActionType
	var amount int64

	switch strings.ToUpper(fields[0]) {
	case "F", "FOLD":
		action = poker.Fold
	case "K", "CHECK":
		action = poker.Check
	case "C", "CALL":
		action = poker.Call
	case "B", "BET":
		action = poker.Bet

		if len(fields) < 2 {
			return poker.PlayerAction{}, errors.New("bet requires an amount")
		}

		parsedAmount, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return poker.PlayerAction{}, errors.New("invalid bet amount")
		}

		amount = parsedAmount
		if amount <= 0 {
			return poker.PlayerAction{}, errors.New("bet amount must be positive")
		}
	case "R", "RAISE":
		action = poker.Raise
		if len(fields) < 2 {
			return poker.PlayerAction{}, errors.New("raise requires an amount")
		}

		parsedAmount, err := strconv.ParseInt(fields[1], 10, 64)
		if err != nil {
			return poker.PlayerAction{}, errors.New("invalid raise amount")
		}

		amount = parsedAmount
		if amount <= 0 {
			return poker.PlayerAction{}, errors.New("raise amount must be positive")
		}
	case "A", "ALL", "ALLIN":
		action = poker.AllIn
	case "P", "CHIP", "CHIPS":
		action = poker.Chips
	case "H", "HOLE":
		action = poker.Hole
	case "M", "MENU":
		action = poker.Menu
	default:
		action = poker.Invalid
	}

	if action == poker.Invalid {
		return poker.PlayerAction{}, errors.New("invalid player action")
	}

	return poker.PlayerAction{
		PlayerID: clientMessage.client.id,
		Action:   action,
		Amount:   amount,
	}, nil
}
