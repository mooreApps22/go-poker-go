package poker

import (
	"fmt"
	"strings"
)

/* HAND
deck                Deck
players             []*Player
communityCards      [5]Card
communityCardsDealt int
phase               HandPhase
winners             []*Player
dealerIndex         int
smallBlindIndex     int
bigBlindIndex       int
utgIndex            int // utg -> Under The Gun
pot                 Pot
currentCall         int64
*/

/* PLAYER
id            int
name          string
chips         int64
holeCards     [2]Card
bestHandValue HandValue
bestCards     [5]Card
currentBet    int64
hasFolded     bool
hasChecked    bool
*/

func (hand *Hand) AcceptBets() {
	for {
		if hand.bettingRoundComplete() {
			break
		}
		if hand.currentPlayerIndex >= len(hand.players) {
			hand.currentPlayerIndex = 0
		}

		// get current player
		player := hand.players[hand.currentPlayerIndex]

		if player.hasFolded {
			hand.currentPlayerIndex++
			continue
		}

		hand.displayPlayerActions(player)
		hand.handPlayerInput(player)
		// perform action
		// advance to next player

		hand.currentPlayerIndex++
	}
	hand.currentCall = 0
}

func (hand *Hand) bettingRoundComplete() bool {
	for _, player := range hand.players {
		if !player.hasFolded &&
			(!player.hasActed || player.currentBet != hand.currentCall) {
			return false
		}
	}
	return true
}

func (hand *Hand) displayPlayerActions(player *Player) {
	fmt.Printf("%v's Turn:\n", player.name)
	fmt.Println("player.currentBet: ", player.currentBet)
	fmt.Println("hand.currentCall: ", hand.currentCall)
	fmt.Println("hand.voluntaryBets: ", hand.voluntaryBets)
	if player.currentBet < hand.currentCall && !hand.voluntaryBets {
		amountToCall := hand.currentCall - player.currentBet
		fmt.Println("amountToCall: ", amountToCall)

		fmt.Println("[C] Call $", amountToCall)
	} else if player.currentBet < hand.currentCall && hand.voluntaryBets {
		fmt.Println("[B] Bet")
	} else if player.currentBet == hand.currentCall {
		fmt.Println("[K] Check")
		fmt.Println("[B] Bet")
	}
	fmt.Println("[R] Raise")
	fmt.Println("[F] Fold")
	fmt.Println("[A] All In")
}

func (hand *Hand) handPlayerInput(player *Player) {
	for {
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
			player.hasActed = true
			// hand.fold(player)
			return
		case "A", "ALL", "ALLIN", "ALL IN":
			// hand.allIn(player)
			return
		default:
			fmt.Println("Invalid action")
			player.hasActed = false
		}
	}
}

func (hand *Hand) call(player *Player) bool {
	amountToCall := hand.currentCall - player.currentBet

	if amountToCall <= 0 {
		fmt.Println("There is nothing to call.")
		return false
	}

	hand.pot.BuildBet(amountToCall, player)
	player.hasActed = true
	return true
}

func (hand *Hand) bet(player *Player) bool {
	fmt.Print("Enter the amount of your bet: ")

	var betAmount int64
	_, err := fmt.Scan(&betAmount)
	if err != nil {
		fmt.Println("Invalid bet amount")
		return false
	}

	if hand.currentCall != 0 {
		fmt.Println("There is already a bet. You must call or raise.")
		return false
	}

	if betAmount <= 0 {
		fmt.Println("Your bet must be greater than zero.")
		return false
	}

	if betAmount < hand.minimumBet {
		fmt.Printf("Your bet must be greater than %v\n.", hand.minimumBet)
		return false
	}

	hand.pot.BuildBet(betAmount, player)
	hand.currentCall = betAmount
	player.hasActed = true
	return true
}

func (hand *Hand) raise(player *Player) bool {
	fmt.Print("Enter the amount of your raise: ")

	var raiseAmount int64
	_, err := fmt.Scan(&raiseAmount)
	if err != nil {
		fmt.Println("Invalid bet amount")
		return false
	}

	if raiseAmount < hand.currentCall {
		fmt.Println("Your raise must be greater than the current call.")
		return false
	}

	if raiseAmount-hand.currentCall < hand.minimumBet {
		fmt.Printf("Your raise must be greater than %v\n.", hand.minimumBet)
		return false
	}

	amountToRaise := raiseAmount - player.currentBet

	hand.pot.BuildBet(amountToRaise, player)
	hand.currentCall = raiseAmount
	player.hasActed = true
	return true
}
