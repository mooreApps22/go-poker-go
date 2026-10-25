package poker

import (
	"errors"
)

func (hand *Hand) ApplyPlayerAction(action *PlayerAction) (ActionResult, error) {
	var result ActionResult

	if hand.IsOver() {
		return result, errors.New("hand is over")
	}

	player := hand.players[hand.currentPlayerIndex]

	if player.id != action.PlayerID {
		return result, errors.New("not your turn")
	}

	var err error

	switch action.Action {
	case Call:
		amountToCall := hand.currentCall - player.currentBet
		action.Amount = amountToCall

		err = hand.call(player)
	case Check:
		if player.GetCurrentBet() != hand.currentCall {
			err = errors.New("you must meet the current call to check")
		} else {
			player.hasActed = true
		}
	case Fold:
		hand.fold(player)
	case Bet:
		hand.bet(player, action.Amount)
	case Raise:
		hand.raise(player, action.Amount)
	case AllIn:
		hand.allIn(player)
		action.Amount = player.GetChips()
	case Chips:
		action.Amount = player.GetChips()
	case Hole:
		result.Message = FormatPrivateCards(player.GetHoleCards())
	case Menu:
		result.Message = hand.GetPlayerActionOptions(player)
	default:
		return result, errors.New("action not implemented")
	}

	if err != nil {
		return result, err
	}

	result.Action = action.Action
	result.Player = player
	result.Amount = action.Amount

	if hand.bettingRoundComplete() {
		result.BettingRoundEnded = true

		hand.currentCall = 0

		for _, player := range hand.players {
			player.currentBet = 0
			player.hasActed = false
		}

		return result, nil
	}

	if hand.turnAdvancingAction(result) {
		hand.advanceCurrentPlayer()
	}

	return result, nil
}

func (hand *Hand) advanceCurrentPlayer() {
	for {
		hand.currentPlayerIndex++

		if hand.currentPlayerIndex >= len(hand.players) {
			hand.currentPlayerIndex = 0
		}

		player := hand.players[hand.currentPlayerIndex]

		if !player.hasFolded && !player.hasAllIn {
			return
		}
	}
}

func (hand *Hand) turnAdvancingAction(result ActionResult) bool {
	return result.Action != Chips &&
		result.Action != Hole &&
		result.Action != Menu
}
