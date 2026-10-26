package poker

import (
	"errors"
	"fmt"
)

func (hand *Hand) bettingRoundComplete() bool {
	for _, player := range hand.players {
		if !player.hasFolded &&
			!player.hasAllIn &&
			(!player.hasActed || player.currentBet != hand.currentCall) {
			return false
		}
	}

	return true
}

func (hand *Hand) GetPlayerActionOptions(player *Player) string {
	var options string

	if hand.currentCall == 0 {
		options += "[K] Check\n"
		options += "[B] Bet\n"
	} else if player.currentBet < hand.currentCall {
		amountToCall := hand.currentCall - player.currentBet
		options += fmt.Sprintf("[C] Call $%d\n", amountToCall)
	} else {
		options += "[K] Check\n"
	}

	if hand.currentCall > 0 {
		options += "[R] Raise\n"
	}
	options += "[F] Fold\n"
	options += "[A] All In\n"
	options += "[P] Chips\n"
	options += "[H] Hole\n"

	return options
}

func (hand *Hand) call(player *Player) error {
	amountToCall := hand.currentCall - player.currentBet

	if amountToCall <= 0 {
		return errors.New("there is nothing to call")
	}

	if amountToCall >= player.chips {
		player.PlaceBet(player.chips)
		player.hasAllIn = true
		player.hasActed = true
		return nil
	}

	player.PlaceBet(amountToCall)
	player.hasActed = true
	return nil
}

func (hand *Hand) bet(player *Player, betAmount int64) error {
	if hand.currentCall != 0 {
		return errors.New("there is already a bet; You must call or raise")
	}

	if betAmount <= 0 {
		return errors.New("Your bet must be greater than zero.")
	}

	if betAmount < hand.minimumBet {
		return fmt.Errorf(
			"your bet must be at least %d",
			hand.minimumBet,
		)
	}

	if betAmount > player.chips {
		return fmt.Errorf(
			"you cannot bet more than $%d",
			player.chips,
		)
	}

	player.PlaceBet(betAmount)
	hand.currentCall = betAmount

	for _, currentPlayer := range hand.players {
		currentPlayer.hasActed = false
	}

	player.hasActed = true

	return nil
}

func (hand *Hand) raise(player *Player, raiseAmount int64) error {
	if raiseAmount > player.chips {
		return fmt.Errorf(
			"you cannot raise more than $%d",
			player.chips,
		)
	}

	if raiseAmount <= hand.currentCall {
		return errors.New("your raise must be greater than the current call")
	}

	if raiseAmount-hand.currentCall < hand.minimumBet {
		return fmt.Errorf(
			"your raise must be at least %v greater than the current call",
			hand.minimumBet,
		)
	}

	amountToRaise := raiseAmount - player.currentBet

	player.PlaceBet(amountToRaise)
	hand.currentCall = raiseAmount

	for _, currentPlayer := range hand.players {
		currentPlayer.hasActed = false
	}

	player.hasActed = true

	return nil
}

func (hand *Hand) fold(player *Player) {
	player.hasFolded = true
	player.hasActed = true
	hand.haveFoldedCount++
}

func (hand *Hand) IsOver() bool {
	return len(hand.players)-hand.haveFoldedCount == 1
}

func (hand *Hand) CanContinueBetting() bool {
	activePlayerCount := 0

	for _, player := range hand.players {
		if !player.hasFolded && !player.hasAllIn {
			activePlayerCount++
		}
	}

	return activePlayerCount >= 2
}

func (hand *Hand) allIn(player *Player) {
	allInAmount := player.chips

	player.PlaceBet(allInAmount)

	if player.currentBet > hand.currentCall {
		hand.currentCall = player.currentBet

		hand.currentCall = player.currentBet

		for _, otherPlayer := range hand.players {
			otherPlayer.hasActed = false
		}
	}

	player.hasAllIn = true
	player.hasActed = true
}
