package server

import (
	"errors"
	"fmt"
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
	var err error

	switch strings.ToUpper(fields[0]) {
	case "F", "FOLD":
		action = poker.Fold
	case "K", "CHECK":
		action = poker.Check
	case "C", "CALL":
		action = poker.Call
	case "B", "BET":
		action = poker.Bet
		amount, err = parseActionAmount(fields, "bet")
	case "R", "RAISE":
		action = poker.Raise
		amount, err = parseActionAmount(fields, "raise")
	case "A", "ALL", "ALLIN":
		action = poker.AllIn
	default:
		return poker.PlayerAction{},
			errors.New("invalid player action")
	}

	if err != nil {
		return poker.PlayerAction{}, err
	}

	return poker.PlayerAction{
		PlayerID: clientMessage.client.id,
		Action:   action,
		Amount:   amount,
	}, nil
}

func parseActionAmount(fields []string, actionName string) (int64, error) {
	if len(fields) < 2 {
		return 0, fmt.Errorf("%s requires an amount", actionName)
	}

	amount, err := strconv.ParseInt(fields[1], 10, 64)
	if err != nil {
		return 0, fmt.Errorf("invalid %s amount", actionName)
	}

	if amount <= 0 {
		return 0, fmt.Errorf("%s amount must be positive", actionName)
	}

	return amount, nil
}
