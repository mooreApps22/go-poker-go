package poker

import (
	"errors"
)

func (hand *Hand) ApplyPlayerAction(action PlayerAction) error {
	if hand.IsOver() {
		return errors.New("hand is over")
	}

	player := hand.players[hand.currentPlayerIndex]

	if player.id != action.PlayerID {
		return errors.New("not your turn")
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
	default:
		return errors.New("action not implemented")
	}

	if err != nil {
		return err
	}

	hand.advanceCurrentPlayer()

	return nil
}

func (hand *Hand) advanceCurrentPlayer() {
	hand.currentPlayerIndex++

	if hand.currentPlayerIndex >= len(hand.players) {
		hand.currentPlayerIndex = 0
	}
}
