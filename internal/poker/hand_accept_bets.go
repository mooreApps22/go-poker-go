package poker

import (
	"errors"
	"fmt"
)

/*
func (hand *Hand) AcceptBets() {
	if hand.IsOver() {
		return
	}
	for {
		if hand.bettingRoundComplete() {
			break
		}

		if hand.currentPlayerIndex >= len(hand.players) {
			hand.currentPlayerIndex = 0
		}

		player := hand.players[hand.currentPlayerIndex]

		if player.hasFolded || player.hasAllIn {
			hand.currentPlayerIndex++
			continue
		}

		hand.handlePlayerInput(player)

		if hand.IsOver() {
			return
		}

		hand.currentPlayerIndex++
	}

	// Reset for next Betting Round
	hand.currentCall = 0

	for _, player := range hand.players {
		player.currentBet = 0
		player.hasActed = false
	}
}
*/

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

func (hand *Hand) displayPlayerActions(player *Player) {
	if len(hand.GetCommunityCardsDealt()) > 0 {
		fmt.Println("Community Cards Dealt:")
		fmt.Println(FormatPublicCards(hand.GetCommunityCardsDealt()))
	}
	if len(player.holeCards) > 0 {
		fmt.Printf("%v's Hole Cards:\n", player.name)
		fmt.Println(FormatPublicCards(player.GetHoleCards()))
	}
	fmt.Println("Current Pot: $", hand.GetCurrentPotValue())
	fmt.Printf("%v's Turn:\n", player.name)
	if hand.currentCall == 0 {
		fmt.Println("[K] Check")
		fmt.Println("[B] Bet")
	} else if player.currentBet < hand.currentCall {
		amountToCall := hand.currentCall - player.currentBet
		fmt.Println("[C] Call $", amountToCall)
	} else {
		fmt.Println("[K] Check")
	}

	if hand.currentCall > 0 {
		fmt.Println("[R] Raise")
	}
	fmt.Println("[F] Fold")
	fmt.Println("[A] All In")
	fmt.Println("[P] Chips")
}

/*
func (hand *Hand) handlePlayerInput(player *Player) {
	for {
		hand.displayPlayerActions(player)

		fmt.Print("> ")

		var input string
		fmt.Scan(&input)
		upperInput := strings.ToUpper(input)

		switch upperInput {
		case "K", "CHECK":
			if player.currentBet != hand.currentCall {
				fmt.Println("You cannot check.")
				continue
			}
			player.hasActed = true
			return
		case "C", "CALL":
			if hand.call(player) {
				return
			}
		case "B", "BET":
			if hand.bet(player) {
				return
			}
		case "R", "RAISE":
			if hand.raise(player) {
				return
			}
		case "F", "FOLD":
			hand.fold(player)
			return
		case "A", "ALL", "ALLIN", "ALL IN":
			hand.allIn(player)
			return
		case "P", "CHIP", "CHIPS":
			fmt.Println("Chips: $", player.chips)
		default:
			fmt.Println("Invalid action")
			player.hasActed = false
		}
	}
}
*/

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
