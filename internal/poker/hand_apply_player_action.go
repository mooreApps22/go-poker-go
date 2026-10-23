package poker

import (
	"errors"
)

func (hand *Hand) ApplyPlayerAction(action PlayerAction) (bool, error) {
	if hand.IsOver() {
		return false, errors.New("hand is over")
	}

	player := hand.players[hand.currentPlayerIndex]

	if player.id != action.PlayerID {
		return false, errors.New("not your turn")
	}

	var err error

	switch action.Action {
	case Call:
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
	default:
		return false, errors.New("action not implemented")
	}

	if err != nil {
		return false, err
	}

	if hand.bettingRoundComplete() {
		hand.currentCall = 0

		for _, player := range hand.players {
			player.currentBet = 0
			player.hasActed = false
		}
		return true, nil
	}

	hand.advanceCurrentPlayer()

	return false, nil
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
