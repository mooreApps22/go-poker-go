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
			player.hasActed = true
			hand.call(player)
			return
		case "B", "BET":
			player.hasActed = true
			hand.bet(player)
			return
		case "R", "RAISE":
			player.hasActed = true
			// hand.raise(player)
			return
		case "F", "FOLD":
			player.hasActed = true
			// hand.fold(player)
			return
		case "A", "ALL", "ALLIN", "ALL IN":
			player.hasActed = true
			// hand.allIn(player)
			return
		default:
			fmt.Println("Invalid action")
			player.hasActed = false
		}
	}
}

func (hand *Hand) call(player *Player) {
	amountToCall := hand.currentCall - player.currentBet

	hand.pot.BuildBet(amountToCall, player)
}

func (hand *Hand) bet(player *Player) {
	fmt.Print("Enter the amount of your bet: ")

	var betAmount int64
	_, err := fmt.Scan(&betAmount)
	if err != nil {
		fmt.Println("Invalid bet amount")
		return
	}

	if hand.currentCall != 0 {
		fmt.Println("There is already a bet. You must call or raise.")
		return
	}

	if betAmount <= 0 {
		fmt.Println("Your bet must be greater than zero.")
		return
	}

	hand.pot.BuildBet(betAmount, player)
	hand.currentCall = betAmount
}
