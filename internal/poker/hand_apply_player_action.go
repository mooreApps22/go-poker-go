package poker

import (
	"errors"
)

func (hand *Hand) ApplyPlayerAction(action PlayerAction) (ActionResult, error) {
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
	case Fold:
		hand.fold(player)
	case Check:
		if player.GetCurrentBet() != hand.currentCall {
			err = errors.New("you must meet the current call to check")
		} else {
			player.hasActed = true
		}
	case Call:
		result.Amount = hand.currentCall - player.currentBet
		err = hand.call(player)
	case Bet:
		result.Amount = action.Amount
		hand.bet(player, action.Amount)
	case Raise:
		result.Amount = action.Amount
		hand.raise(player, action.Amount)
	case AllIn:
		action.Amount = player.GetChips()
		hand.allIn(player)
	default:
		return result, errors.New("action not implemented")
	}

	if err != nil {
		return result, err
	}

	result.Action = action.Action
	result.Player = player

	if hand.bettingRoundComplete() {
		result.BettingRoundEnded = true
		hand.resetBettingRound()

		return result, nil
	}

	hand.advanceCurrentPlayer()

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

func (hand *Hand) SetFirstPostFlopPlayer() {
	hand.currentPlayerIndex = hand.dealerIndex
	hand.advanceCurrentPlayer()
}

func (hand *Hand) resetBettingRound() {
	hand.currentCall = 0

	for _, player := range hand.players {
		player.currentBet = 0
		player.hasActed = false
	}
}
