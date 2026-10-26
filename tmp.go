
==================== internal/poker/action_result.go ====================
package poker

type ActionResult struct {
	Action            ActionType
	Player            *Player
	Amount            int64
	BettingRoundEnded bool
	Message           string
}

==================== internal/poker/card.go ====================
package poker

import "fmt"

type Rank uint8

const (
	Two Rank = iota + 2
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

func (rank Rank) String() string {
	switch rank {
	case Two:
		return "2"
	case Three:
		return "3"
	case Four:
		return "4"
	case Five:
		return "5"
	case Six:
		return "6"
	case Seven:
		return "7"
	case Eight:
		return "8"
	case Nine:
		return "9"
	case Ten:
		return "10"
	case Jack:
		return "J"
	case Queen:
		return "Q"
	case King:
		return "K"
	case Ace:
		return "A"
	default:
		return "?"
	}
}

type Suit uint8

const (
	Spades Suit = iota
	Clubs
	Diamonds
	Hearts
)

func (s Suit) String() string {
	switch s {
	case Spades:
		return "♠\uFE0E"
	case Clubs:
		return "♣\uFE0E"
	case Diamonds:
		return "♦\uFE0E"
	case Hearts:
		return "❤\uFE0E"
	default:
		return "?"
	}
}

type Card struct {
	Rank Rank
	Suit Suit
}

func (card Card) String() string {
	return "[" + card.Rank.String() + card.Suit.String() + "]"
}

func (card Card) PrettyRow(row int) string {
	switch row {
	case 0:
		return "┌─────┐"
	case 1:
		return fmt.Sprintf("┊ %-2v%v ┊", card.Rank, card.Suit)
	case 2:
		return fmt.Sprintf("┊ %v%2v ┊", card.Suit, card.Rank)
	case 3:
		return "└─────┘"
	default:
		return ""
	}
}

func (card Card) PrettyCardString() string {
	result := ""

	for row := 0; row < 4; row++ {
		result += card.PrettyRow(row)

		if row < 3 {
			result += "\n"
		}
	}
	return result
}

// Packet Function because no receiver
func FormatPrivateCards(cards []Card) string {
	result := ""
	for row := 0; row < 4; row++ {
		for _, card := range cards {
			result += card.PrettyRow(row)
			result += " "
		}

		if row < 3 {
			result += "\n"
		}
	}
	return result
}

func FormatPublicCards(cards []Card) string {
	result := ""
	for row := 0; row < 4; row++ {
		for _, card := range cards {
			result += card.PrettyRow(row)
			result += " "
		}

		result += "\n"
	}
	return result
}

==================== internal/poker/createPots_test.go ====================
package poker

import "testing"

func TestCreatePots_EqualContributions(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	for _, player := range players {
		player.totalContribution = 100_000
	}

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 1 {
		t.Fatalf(
			"expected 1 pot, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 400_000 {
		t.Errorf(
			"expected pot value 400_000, got %d",
			hand.pots[0].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 4 {
		t.Errorf(
			"expected 4 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}
}

func TestCreatePots_OneShort(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	adam.totalContribution = 50_000
	bill.totalContribution = 100_000
	cathy.totalContribution = 100_000
	debra.totalContribution = 100_000

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 2 {
		t.Fatalf(
			"expected 2 pots, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 200_000 {
		t.Errorf(
			"expected Pot A value 200_000, got %d",
			hand.pots[0].value,
		)
	}

	if hand.pots[1].value != 150_000 {
		t.Errorf(
			"expected Pot B value 150_000, got %d",
			hand.pots[1].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 4 {
		t.Errorf(
			"expected 4 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}

	if len(hand.pots[1].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[1].eligiblePlayers),
		)
	}
}

func TestCreatePots_TwoAllIn(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	adam.totalContribution = 50_000
	bill.totalContribution = 100_000
	cathy.totalContribution = 200_000
	debra.totalContribution = 200_000

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 3 {
		t.Fatalf(
			"expected 3 pots, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 200_000 {
		t.Errorf(
			"expected pot A value 200_000, got %d",
			hand.pots[0].value,
		)
	}

	if hand.pots[1].value != 150_000 {
		t.Errorf(
			"expected pot B value 150_000, got %d",
			hand.pots[1].value,
		)
	}

	if hand.pots[2].value != 200_000 {
		t.Errorf(
			"expected pot C value 200_000, got %d",
			hand.pots[2].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 4 {
		t.Errorf(
			"expected 4 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}

	if len(hand.pots[1].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[1].eligiblePlayers),
		)
	}

	if len(hand.pots[2].eligiblePlayers) != 2 {
		t.Errorf(
			"expected 2 eligible players, got %d",
			len(hand.pots[2].eligiblePlayers),
		)
	}
}

func TestCreatePots_FoldedPlayerDeadMoney(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	for _, player := range players {
		player.totalContribution = 100_000
	}

	adam.hasFolded = true

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 1 {
		t.Fatalf(
			"expected 1 pot, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 400_000 {
		t.Errorf(
			"expected pot A value 400_000, got %d",
			hand.pots[0].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}
}

func TestCreatePots_FoldedPlayerLowerContributionLevel(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	adam.totalContribution = 25_000
	bill.totalContribution = 100_000
	cathy.totalContribution = 100_000
	debra.totalContribution = 100_000

	adam.hasFolded = true

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 1 {
		t.Fatalf(
			"expected 1 pot, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 325_000 {
		t.Errorf(
			"expected pot A value 325_000, got %d",
			hand.pots[0].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}
}

func TestCreatePots_FoldedPlayerThreePots(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")
	debra := NewPlayer(4, "Debra")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
		&debra,
	}

	adam.totalContribution = 50_000
	bill.totalContribution = 100_000
	cathy.totalContribution = 200_000
	debra.totalContribution = 200_000

	bill.hasFolded = true

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 2 {
		t.Fatalf(
			"expected 2 pots, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 200_000 {
		t.Errorf(
			"expected pot A value 200_000, got %d",
			hand.pots[0].value,
		)
	}

	if hand.pots[1].value != 350_000 {
		t.Errorf(
			"expected pot B value 350_000, got %d",
			hand.pots[1].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}

	if len(hand.pots[1].eligiblePlayers) != 2 {
		t.Errorf(
			"expected 2 eligible players, got %d",
			len(hand.pots[1].eligiblePlayers),
		)
	}
}

func TestCreatePots_UncalledExcess(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
	}

	cathyStartingChips := cathy.chips

	adam.totalContribution = 50_000
	bill.totalContribution = 100_000
	cathy.totalContribution = 300_000

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 2 {
		t.Fatalf(
			"expected 2 pot, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 150_000 {
		t.Errorf(
			"expected pot A value 150_000, got %d",
			hand.pots[0].value,
		)
	}

	if hand.pots[1].value != 100_000 {
		t.Errorf(
			"expected pot B value 100_000, got %d",
			hand.pots[1].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 3 {
		t.Errorf(
			"expected 3 eligible players, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}

	if len(hand.pots[1].eligiblePlayers) != 2 {
		t.Errorf(
			"expected 2 eligible players, got %d",
			len(hand.pots[1].eligiblePlayers),
		)
	}

	if cathy.chips != cathyStartingChips+200_000 {
		t.Errorf(
			"expected Cathy to receive 200_000 uncalled chips, got %d",
			cathy.chips-cathyStartingChips,
		)
	}
}

func TestCreatePots_FoldedUncalled(t *testing.T) {

	adam := NewPlayer(1, "Adam")
	bill := NewPlayer(2, "Bill")
	cathy := NewPlayer(3, "Cathy")

	players := []*Player{
		&adam,
		&bill,
		&cathy,
	}

	cathyStartingChips := cathy.chips

	adam.totalContribution = 100_000
	bill.totalContribution = 300_000
	cathy.totalContribution = 500_000

	/*
		100k * 3 = 300k
		fold
		200k * 2 = 400k + 300k = 700k
		fold

	*/

	adam.hasFolded = true
	bill.hasFolded = true

	hand := NewHand(players)

	hand.CreatePots()

	if len(hand.pots) != 1 {
		t.Fatalf(
			"expected 1 pot, got %d",
			len(hand.pots),
		)
	}

	if hand.pots[0].value != 700_000 {
		t.Errorf(
			"expected pot A value 700_000, got %d",
			hand.pots[0].value,
		)
	}

	if len(hand.pots[0].eligiblePlayers) != 1 {
		t.Errorf(
			"expected 1 eligible player, got %d",
			len(hand.pots[0].eligiblePlayers),
		)
	}

	if cathy.chips != cathyStartingChips+200_000 {
		t.Errorf(
			"expected Cathy to receive 200_000 uncalled chips, got %d",
			cathy.chips-cathyStartingChips,
		)
	}
}

==================== internal/poker/deck.go ====================
package poker

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	cards [52]Card
	next  int
}

func NewDeck() Deck {
	var deck Deck
	index := 0

	for suit := Spades; suit <= Hearts; suit++ {
		for rank := Two; rank <= Ace; rank++ {
			card := Card{
				Rank: rank,
				Suit: suit,
			}
			deck.cards[index] = card
			index++
		}
	}
	return deck
}

func (deck Deck) Print() {
	const cardsPerGroup = 7
	for start := 0; start < len(deck.cards); start += cardsPerGroup {
		end := start + cardsPerGroup

		if end > len(deck.cards) {
			end = len(deck.cards)
		}
		fmt.Println(FormatPublicCards((deck.cards[start:end])))
	}
}

func (deck *Deck) Shuffle() {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := len(deck.cards) - 1; i > 0; i-- {
		randomIndex := random.Intn(i + 1)

		deck.cards[i], deck.cards[randomIndex] =
			deck.cards[randomIndex], deck.cards[i]
	}
}

func (deck *Deck) Draw() (Card, error) {
	if deck.next >= len(deck.cards) {
		return Card{}, errors.New("cannot draw from empty deck")
	}

	nextCard := deck.cards[deck.next]
	deck.next++
	return nextCard, nil
}

==================== internal/poker/game.go ====================
package poker

import "fmt"

type Game struct {
	players     []*Player
	dealerIndex int
	hand        Hand
}

func NewGame(players []*Player) Game {
	return Game{
		players:     players,
		dealerIndex: -1,
	}
}

func (game *Game) removePlayersWithNoChips() {
	activePlayers := game.players[:0]

	for _, player := range game.players {
		if player.chips > 0 {
			activePlayers = append(activePlayers, player)
		} else {
			fmt.Printf("%v has left the poker table, due to losing all chips.\n", player.name)
		}
	}

	game.players = activePlayers
}

func (game *Game) SetUpNewHand() {
	game.removePlayersWithNoChips()

	game.dealerIndex = (game.dealerIndex + 1) % len(game.players)

	game.hand = NewHand(game.players)

	game.hand.dealerIndex = game.dealerIndex
	game.hand.smallBlindIndex = (game.dealerIndex + 1) % len(game.players)
	game.hand.bigBlindIndex = (game.dealerIndex + 2) % len(game.players)
	game.hand.utgIndex = (game.dealerIndex + 3) % len(game.players)
	game.hand.currentPlayerIndex = game.hand.utgIndex
}

func (game *Game) GetHand() *Hand {
	return &game.hand
}

func (game *Game) RotateDealer() {
	game.dealerIndex = (game.dealerIndex + 1) % len(game.players)
}

func (game *Game) GetDealerIndex() int {
	return game.dealerIndex
}

==================== internal/poker/hand_apply_player_action.go ====================
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

func (hand *Hand) SetFirstPostFlopPlayer() {
	hand.currentPlayerIndex = hand.dealerIndex
	hand.advanceCurrentPlayer()
}

==================== internal/poker/hand_best_player_hand.go ====================
package poker

import "fmt"

func (hand Hand) ShowDown() {
	fmt.Println("SHOWDOWN PHASE: ")

	for _, player := range hand.players {
		player.bestHandValue, player.bestCards = findBestHandValue(
			hand.GetCommunityCards(),
			player.holeCards,
		)
	}
}

func findBestHandValue(communityCards [5]Card, holeCards [2]Card) (
	HandValue, [5]Card) {
	var possibleHandValues [21]HandValue
	var possibleHands [21][5]Card
	possibleHandsIndex := 0

	var availableCards [7]Card
	availableCards[0] = holeCards[0]
	availableCards[1] = holeCards[1]

	for index := 0; index < len(communityCards); index++ {
		availableCards[index+2] = communityCards[index]
	}

	for firstExcluded := 0; firstExcluded < len(availableCards)-1; firstExcluded++ {
		for secondExcluded := firstExcluded + 1; secondExcluded < 7; secondExcluded++ {
			var possibleHand [5]Card
			possibleHandIndex := 0

			for cardIndex := 0; cardIndex < len(availableCards); cardIndex++ {
				if cardIndex != firstExcluded && cardIndex != secondExcluded {
					possibleHand[possibleHandIndex] = availableCards[cardIndex]
					possibleHandIndex++
				}
			}

			possibleHands[possibleHandsIndex] = possibleHand
			possibleHandValues[possibleHandsIndex] = EvaluateBestFiveCardHand(possibleHand)
			possibleHandsIndex++
		}
	}

	return findHighestHandValue(possibleHandValues, possibleHands)
}

func comparePlayersHandValues(player1 *Player, player2 *Player) (*Player, bool) {
	winningHandValue, isTied := findWinningHandValue(
		player1.bestHandValue,
		player2.bestHandValue,
	)

	if isTied {
		return player1, true
	}

	if winningHandValue == player1.bestHandValue {
		return player1, false
	}

	return player2, false
}

==================== internal/poker/hand_betting.go ====================
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

==================== internal/poker/hand_blinds.go ====================
package poker

func (hand *Hand) ResetFoldedHands() {
	for _, player := range hand.players {
		player.hasFolded = false
	}
}

func (hand *Hand) PostBlinds() {
	hand.players[hand.smallBlindIndex].PlaceBet(SmallBlindBet)
	hand.players[hand.bigBlindIndex].PlaceBet(BigBlindBet)
}

==================== internal/poker/hand_deal.go ====================
package poker

import "errors"

func (hand *Hand) DealHoleCards() error {
	for holeCardIndex := 0; holeCardIndex < 2; holeCardIndex++ {
		for offset := 0; offset < len(hand.players); offset++ {
			playerIndex := (hand.smallBlindIndex + offset) % len(hand.players)
			card, err := hand.deck.Draw()
			if err != nil {
				return err
			}

			hand.players[playerIndex].holeCards[holeCardIndex] = card
		}
	}
	hand.phase = PreFlop
	return nil
}

func (hand *Hand) DealFlopCards() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	for flopCardIndex := 0; flopCardIndex < 3; flopCardIndex++ {
		card, err := hand.deck.Draw()
		if err != nil {
			return err
		}
		hand.communityCards[flopCardIndex] = card
	}
	hand.communityCardsDealt += 3
	hand.phase = Flop
	return nil
}

func (hand Hand) GetCommunityCardsDealt() []Card {
	return hand.communityCards[:hand.communityCardsDealt]
}

func (hand *Hand) DealTurnCard() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	turnCard, err := hand.deck.Draw()
	if err != nil {
		return err
	}
	hand.communityCards[3] = turnCard
	hand.communityCardsDealt += 1
	hand.phase = Turn
	return nil
}

func (hand *Hand) DealRiverCard() error {
	_, err := hand.deck.Draw()
	if err != nil {
		return err
	}

	riverCard, err := hand.deck.Draw()
	if err != nil {
		return err
	}
	hand.communityCards[4] = riverCard
	hand.communityCardsDealt += 1
	hand.phase = River
	return nil
}

func (hand *Hand) AdvancePhase() error {
	switch hand.phase {
	case PreFlop:
		return hand.DealFlopCards()
	case Flop:
		return hand.DealTurnCard()
	case Turn:
		return hand.DealRiverCard()
	case River:
		hand.phase = Showdown
		hand.ShowDown()
		return nil
	default:
		return errors.New("hand cannot advance from current phase")
	}
}

==================== internal/poker/hand_find_best_value.go ====================
package poker

func findHighestHandValue(possibleHandValues [21]HandValue, possibleHands [21][5]Card) (
	HandValue, [5]Card) {
	maxHandValue := possibleHandValues[0]
	maxHandIndex := 0

	for index := 1; index < len(possibleHandValues); index++ {
		winningHandValue, isTied := findWinningHandValue(
			maxHandValue,
			possibleHandValues[index],
		)

		if !isTied && winningHandValue == possibleHandValues[index] {
			maxHandValue = winningHandValue
			maxHandIndex = index
		}
	}

	return maxHandValue, possibleHands[maxHandIndex]
}

==================== internal/poker/hand_getters.go ====================
package poker

func (hand Hand) GetDeck() Deck {
	return hand.deck
}

func (hand Hand) GetPlayers() []*Player {
	return hand.players
}

func (hand Hand) GetCommunityCards() [5]Card {
	return hand.communityCards
}

func (hand Hand) GetWinner() []*Player {
	return hand.winners
}

func (hand Hand) GetSmallBlindIndex() int {
	return hand.smallBlindIndex
}

func (hand Hand) GetDealerIndex() int {
	return hand.dealerIndex
}

func (hand Hand) GetBigBlindIndex() int {
	return hand.bigBlindIndex
}

func (hand Hand) GetUtgIndex() int {
	return hand.utgIndex
}

func (hand *Hand) GetCurrentPotValue() int64 {
	var currentPotValue int64

	for _, player := range hand.players {
		currentPotValue += player.totalContribution
	}
	return currentPotValue
}

==================== internal/poker/hand.go ====================
package poker

type Hand struct {
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
	currentCall         int64
	currentPlayerIndex  int
	voluntaryBets       bool
	minimumBet          int64
	haveFoldedCount     int
	contributionLevels  []int64
	pots                []Pot
}

const BigBlindBet = 4_000
const SmallBlindBet = 2_000

type HandPhase uint8

const (
	SetUp HandPhase = iota
	PreFlop
	Flop
	Turn
	River
	Showdown
)

type HandRank uint8

const (
	HighCard HandRank = iota
	OnePair
	TwoPair
	ThreeOfAKind
	Straight
	Flush
	FullHouse
	FourOfAKind
	StraightFlush
)

type HandValue struct {
	Category    HandRank
	Tiebreakers [5]Rank
}

func NewHand(players []*Player) Hand {
	deck := NewDeck()
	deck.Shuffle()

	return Hand{
		deck:               deck,
		players:            players,
		phase:              SetUp,
		currentCall:        BigBlindBet,
		voluntaryBets:      false,
		minimumBet:         BigBlindBet,
		haveFoldedCount:    0,
		contributionLevels: nil,
		pots:               nil,
	}
}

func (hand *Hand) ResetPlayersForNewHand() {
	for _, player := range hand.players {
		player.ResetForNewHand()
	}
}

func (hand *Hand) GetDealer() *Player {
	return hand.players[hand.dealerIndex]
}

func (hand *Hand) GetSmallBlindPlayer() *Player {
	return hand.players[hand.smallBlindIndex]
}

func (hand *Hand) GetBigBlindPlayer() *Player {
	return hand.players[hand.bigBlindIndex]
}

func (hand *Hand) GetUTGPlayer() *Player {
	return hand.players[hand.utgIndex]
}

func (hand *Hand) GetCurrentPlayer() *Player {
	return hand.players[hand.currentPlayerIndex]
}

func (hand *Hand) GetPhase() HandPhase {
	return hand.phase
}

func (phase HandPhase) String() string {
	switch phase {
	case PreFlop:
		return "The Pre-Flop:\n"
	case Flop:
		return "The Flop:\n"
	case Turn:
		return "The Turn:\n"
	case River:
		return "The River:\n"
	case Showdown:
		return "The Showdown:\n"
	default:
		return "?"
	}
}

func (hand *Hand) GetPots() []Pot {
	return hand.pots
}

func (hand *Hand) GetWinners() []*Player {
	return hand.winners
}

==================== internal/poker/hand_rank.go ====================
package poker

func (category HandRank) String() string {
	switch category {
	case HighCard:
		return "High Card"
	case OnePair:
		return "One Pair"
	case TwoPair:
		return "Two Pair"
	case ThreeOfAKind:
		return "Three of A Kind"
	case Straight:
		return "Straight"
	case Flush:
		return "Flush"
	case FullHouse:
		return "Full House"
	case FourOfAKind:
		return "Four of A Kind"
	case StraightFlush:
		return "Straight Flush"
	default:
		return "?"
	}
}

func findRankCounts(cards [5]Card) [15]int {
	var rankCounts [15]int
	for _, card := range cards {
		rankCounts[card.Rank] += 1
	}
	return rankCounts
}

func findHighCards(rankCounts [15]int) [5]Rank {
	var highCards [5]Rank
	highCardsIndex := 0
	for rank := Ace; rank >= Two; rank-- {
		if rankCounts[rank] > 0 {
			for repeatIndex := 0; repeatIndex < rankCounts[rank]; repeatIndex++ {
				highCards[highCardsIndex] = rank
				highCardsIndex++

				if highCardsIndex == len(highCards) {
					return highCards
				}
			}
		}
	}
	return highCards
}

==================== internal/poker/hand_sidebets.go ====================
package poker

import (
	"fmt"
	"slices"
)

func calculateSidePot(
	currentContributionLevel int64,
	previousContributionLevel int64,
	numberOfPlayersAtOrAboveCurrentLevel int,
) (int64, int64) {
	return (currentContributionLevel - previousContributionLevel) *
		int64(numberOfPlayersAtOrAboveCurrentLevel), currentContributionLevel
}

func (hand *Hand) collectContributionLevels() {
	for _, player := range hand.players {
		if player.totalContribution > 0 {
			hand.contributionLevels = append(
				hand.contributionLevels,
				player.totalContribution,
			)
		}
	}
	slices.Sort(hand.contributionLevels)
	hand.contributionLevels = slices.Compact(hand.contributionLevels)
}

func (hand *Hand) countPlayersWithContributionLevel(contributeLevel int64) int {
	count := 0
	for _, player := range hand.players {
		if player.totalContribution >= contributeLevel {
			count++
		}
	}
	return count
}

func (hand *Hand) getEligiblePlayers(level int64) []*Player {
	var eligiblePlayers []*Player

	for _, player := range hand.players {
		if player.totalContribution >= level && !player.hasFolded {
			eligiblePlayers = append(eligiblePlayers, player)
		}
	}

	return eligiblePlayers
}

func (hand *Hand) outputPotsData() {
	for potIndex, pot := range hand.pots {
		if potIndex == 0 {
			fmt.Print("Main Pot : $")
		} else {
			fmt.Print("Side Pot ")
		}

		switch potIndex {
		case 1:
			fmt.Print("A: $")
		case 2:
			fmt.Print("B: $")
		case 3:
			fmt.Print("C: $")
		}

		fmt.Print(pot.value)
		fmt.Println()
		fmt.Print("Eligible: ", pot.eligiblePlayers)
		fmt.Println()
	}
}

func (hand *Hand) getPlayerAtOrAboveContributionLevel(level int64) *Player {
	for _, player := range hand.players {
		if player.totalContribution >= level {
			return player
		}
	}
	return nil
}

func (hand *Hand) CreatePots() {
	hand.collectContributionLevels()
	var sidePot int64
	var previousLevel int64

	for _, level := range hand.contributionLevels {

		numberOfPlayerWithLevel := hand.countPlayersWithContributionLevel(level)

		sidePot, previousLevel = calculateSidePot(
			level,
			previousLevel,
			int(numberOfPlayerWithLevel),
		)

		if numberOfPlayerWithLevel == 1 {
			player := hand.getPlayerAtOrAboveContributionLevel(level)
			if player != nil {
				player.CollectWinnings(sidePot) // return chips
			}
			continue
		}

		eligiblePlayers := hand.getEligiblePlayers(level)

		pot := Pot{
			value:           sidePot,
			eligiblePlayers: eligiblePlayers,
		}

		if len(hand.pots) > 0 &&
			slices.Equal(
				hand.pots[len(hand.pots)-1].eligiblePlayers,
				pot.eligiblePlayers,
			) {
			hand.pots[len(hand.pots)-1].value += pot.value
		} else {
			hand.pots = append(hand.pots, pot)
		}
	}

	hand.outputPotsData()
}

func (hand *Hand) PickWinners() {
	//hand.winners = nil

	for potIndex := range hand.pots {
		pot := &hand.pots[potIndex]

		for _, player := range pot.eligiblePlayers {
			if !player.hasFolded {
				pot.winners = []*Player{player}
				break
			}
		}

		if len(pot.winners) == 0 {
			continue
		}

		for _, player := range pot.eligiblePlayers {
			if player.hasFolded || player == pot.winners[0] {
				continue
			}

			winningPlayer, isTied := comparePlayersHandValues(
				pot.winners[0],
				player,
			)

			if isTied {
				pot.winners = append(pot.winners, player)
			} else if winningPlayer == player {
				pot.winners = []*Player{player}
			}
		}
	}
}

func (hand *Hand) AwardPots() {
	for potIndex := range hand.pots {
		pot := &hand.pots[potIndex]

		if len(pot.winners) == 0 {
			continue
		}

		winnings := pot.value / int64(len(pot.winners))

		for _, winner := range pot.winners {
			winner.CollectWinnings(winnings)
		}
	}
}

==================== internal/poker/hand_tiebreaker.go ====================
package poker

func findOnePairRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			return rank, true
		}
	}
	return 0, false
}

// Returns the high pair and lower pair ranks in order
func findTwoPairsRank(rankCounts [15]int) (Rank, Rank, bool) {
	var pair1 Rank
	var pair2 Rank
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 && pair1 == 0 {
			pair1 = rank
		} else if rankCounts[rank] == 2 && pair2 == 0 {
			pair2 = rank
		}

		if pair1 != 0 && pair2 != 0 {
			if pair1 < pair2 {
				pair1, pair2 = pair2, pair1
			}
			return pair1, pair2, true
		}
	}

	return 0, 0, false
}

func findThreeOfAKindRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 3 {
			return rank, true
		}
	}
	return 0, false
}

func findFourOfAKindRank(rankCounts [15]int) (Rank, bool) {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 4 {
			return rank, true
		}
	}
	return 0, false
}

func arrangeTieBreakers(handValue *HandValue, rankCounts [15]int) {
	switch handValue.Category {
	case HighCard:
		//FREE
		return
	case OnePair:
		var arranged [5]Rank

		pairRank, isFound := findOnePairRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = pairRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != pairRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case TwoPair:
		// pair1, pair2, kicker, doesn't matter, doesn't matter
		var arranged [5]Rank

		pair1Rank, pair2Rank, isFound := findTwoPairsRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = pair1Rank
		arranged[1] = pair2Rank
		arrangedIndex := 2

		for _, rank := range handValue.Tiebreakers {
			if rank != pair1Rank && rank != pair2Rank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case ThreeOfAKind:
		// trips, kicker1, kicker2, doesn't matter, doesn't matter
		var arranged [5]Rank

		tripsRank, isFound := findThreeOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = tripsRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != tripsRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case Straight:
		//FREE
		return
	case Flush:
		//FREE
		return
	case FullHouse:
		// trips, pair, doesn't matter, doesn't matter, doesn't matter
		var arranged [5]Rank

		tripsRank, isFound := findThreeOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		pairRank, isFound := findOnePairRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = tripsRank
		arranged[1] = pairRank

		handValue.Tiebreakers = arranged
		return
	case FourOfAKind:
		// quad, kicker, doesn't matter, doesn't matter, doesn't matter
		var arranged [5]Rank

		quadsRank, isFound := findFourOfAKindRank(rankCounts)
		if !isFound {
			return
		}

		arranged[0] = quadsRank
		arrangedIndex := 1

		for _, rank := range handValue.Tiebreakers {
			if rank != quadsRank {
				arranged[arrangedIndex] = rank
				arrangedIndex++
			}
		}

		handValue.Tiebreakers = arranged
		return
	case StraightFlush:
		//FREE
		return
	default:
		return
	}
}

==================== internal/poker/hand_value_category_checks.go ====================
package poker

func checkForOnePair(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			return true
		}
	}
	return false
}

func checkForTwoPair(rankCounts [15]int) bool {
	pairCount := 0
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			pairCount++
		}
	}
	return pairCount == 2
}

func checkForThreeOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 3 {
			return true
		}
	}
	return false
}

func checkForFullHouse(rankCounts [15]int) bool {
	pairCount := 0
	threeOfAKindCount := 0
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 2 {
			pairCount++
		}
		if rankCounts[rank] == 3 {
			threeOfAKindCount++
		}
	}
	return pairCount == 1 && threeOfAKindCount == 1
}

// A Wheel is when a Straight == A, 5, 4, 3, 2, but 5 is treated as the High Card
func specialWheelStraightTieBreakerRearrangement(ranks *[5]Rank) {
	ranks[0] = Five
	ranks[1] = Four
	ranks[2] = Three
	ranks[3] = Two
	ranks[4] = Ace
}

func checkIfStraight(rankCounts [15]int) (bool, bool) {

	if rankCounts[Ace] == 1 &&
		rankCounts[Two] == 1 &&
		rankCounts[Three] == 1 &&
		rankCounts[Four] == 1 &&
		rankCounts[Five] == 1 {
		return true, true
	}
	straightTally := 0
	straightTallyBegan := false
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 0 {
			if straightTallyBegan == false {
				continue
			} else {
				return false, false
			}
		} else if rankCounts[rank] > 1 {
			return false, false
		} else {
			straightTallyBegan = true
			straightTally += 1
		}
	}
	return straightTally == 5, false
}

func checkIfFlush(cards [5]Card) bool {
	suitSample := cards[0].Suit

	for _, card := range cards[1:] {
		if suitSample != card.Suit {
			return false
		}
	}
	return true
}

func checkForFourOfAKind(rankCounts [15]int) bool {
	for rank := Two; rank <= Ace; rank++ {
		if rankCounts[rank] == 4 {
			return true
		}
	}
	return false
}

==================== internal/poker/hand_value.go ====================
package poker

func higherCategory(handValueA HandValue, handValueB HandValue) HandValue {

	if handValueA.Category > handValueB.Category {
		return handValueA
	}
	return handValueB
}

func compareHandValueTiebreakers(handValueA HandValue, handValueB HandValue) (
	HandValue, bool) {
	tiebreakerA := handValueA.Tiebreakers
	tiebreakerB := handValueB.Tiebreakers

	for index := 0; index < len(tiebreakerA); index++ {
		if tiebreakerA[index] != tiebreakerB[index] {
			if tiebreakerA[index] > tiebreakerB[index] {
				return handValueA, false
			}
			return handValueB, false
		}
	}
	return handValueA, true
}

func findWinningHandValue(handValueA HandValue, handValueB HandValue) (
	HandValue, bool) {
	if handValueA.Category != handValueB.Category {
		return higherCategory(handValueA, handValueB), false
	}

	return compareHandValueTiebreakers(handValueA, handValueB)
}

func EvaluateBestFiveCardHand(cards [5]Card) HandValue {
	var handValue HandValue
	isFlush := checkIfFlush(cards)
	rankCounts := findRankCounts(cards)
	handValue.Tiebreakers = findHighCards(rankCounts)
	isStraight, isWheel := checkIfStraight(rankCounts)
	// Check if Straight is a Wheel Straight --> A, 2, 3, 4, 5
	if isWheel {
		specialWheelStraightTieBreakerRearrangement(&handValue.Tiebreakers)
	}

	if isFlush && isStraight {
		handValue.Category = StraightFlush
	} else if checkForFourOfAKind(rankCounts) {
		handValue.Category = FourOfAKind
	} else if checkForFullHouse(rankCounts) {
		handValue.Category = FullHouse
	} else if isFlush {
		handValue.Category = Flush
	} else if isStraight {
		handValue.Category = Straight
	} else if checkForThreeOfAKind(rankCounts) {
		handValue.Category = ThreeOfAKind
	} else if checkForTwoPair(rankCounts) {
		handValue.Category = TwoPair
	} else if checkForOnePair(rankCounts) {
		handValue.Category = OnePair
	} else {
		handValue.Category = HighCard
	}

	arrangeTieBreakers(&handValue, rankCounts)

	return handValue
}

==================== internal/poker/player_action.go ====================
package poker

type ActionType uint8

const (
	Invalid ActionType = iota
	Fold
	Check
	Call
	Bet
	Raise
	AllIn
	Chips
	Hole
	Menu
)

type PlayerAction struct {
	PlayerID int
	Action   ActionType
	Amount   int64
}

func (action ActionType) String() string {
	switch action {
	case Fold:
		return "FOLD"
	case Check:
		return "CHECK"
	case Call:
		return "CALL"
	case Bet:
		return "BET"
	case Raise:
		return "RAISE"
	case AllIn:
		return "ALL IN"
	case Chips:
		return "CHIPS"
	case Hole:
		return "HOLE"
	case Menu:
		return "MENU"
	default:
		return "INVALID"
	}
}

==================== internal/poker/player.go ====================
package poker

import "fmt"

type Player struct {
	id                int
	name              string
	chips             int64
	holeCards         [2]Card
	bestHandValue     HandValue
	bestCards         [5]Card
	currentBet        int64
	hasFolded         bool
	hasChecked        bool
	hasActed          bool
	hasAllIn          bool
	totalContribution int64
}

func NewPlayer(id int, name string) Player {
	return Player{
		id:         id,
		name:       name,
		chips:      1_000_000,
		hasFolded:  false,
		hasChecked: false,
		hasActed:   false,
		hasAllIn:   false,
	}
}

func (player Player) String() string {
	return fmt.Sprintf("%v[%v] — Chips: $%v", player.name, player.id, player.chips)
}

func (player *Player) GetHoleCards() []Card {
	return player.holeCards[:]
}

func (player *Player) GetBestHandValue() HandValue {
	return player.bestHandValue
}

func (player *Player) GetBestCards() []Card {
	return player.bestCards[:]
}

func (player *Player) PlaceBet(betAmount int64) int64 {
	player.chips -= betAmount
	player.currentBet += betAmount
	player.totalContribution += betAmount

	return betAmount
}

func (player *Player) CollectWinnings(winnings int64) {
	player.chips += winnings
	fmt.Printf("Winner: %v collects $%v\n", player.name, winnings)
}

func (player *Player) ResetCurrentBet() {
	player.currentBet = 0
}

func (player *Player) GetCurrentBet() int64 {
	return player.currentBet
}

func (player Player) GetName() string {
	return player.name
}

func (player *Player) ResetForNewHand() {
	player.currentBet = 0
	player.totalContribution = 0
	player.hasActed = false
	player.hasFolded = false
	player.hasAllIn = false
}

func (player Player) GetID() int {
	return player.id
}

func (player Player) GetChips() int64 {
	return player.chips
}

==================== internal/poker/pot.go ====================
package poker

type Pot struct {
	value           int64
	eligiblePlayers []*Player
	winners         []*Player
}

func NewPot() Pot {
	return Pot{
		value:           0,
		eligiblePlayers: nil,
	}
}

func (pot *Pot) BuildBet(betValue int64, player *Player) {
	pot.value += player.PlaceBet(betValue)
}

func (pot Pot) GetValue() int64 {
	return pot.value
}

func (pot Pot) GetEligiblePlayers() []*Player {
	return pot.eligiblePlayers
}

func (pot *Pot) GetWinners() []*Player {
	return pot.winners
}

==================== internal/poker/pot_test.go ====================
package poker

import "testing"

// Parameters: currentContributionLevel, previous, numberOfPlayersAtOrAbove
func TestCalculateSidePot_1(t *testing.T) {
	sidePot, _ := calculateSidePot(
		100_000,
		0,
		4,
	)

	if sidePot != 400_000 {
		t.Errorf(
			"expected 400_000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_2(t *testing.T) {
	sidePot, _ := calculateSidePot(
		200_000,
		100_000,
		3,
	)

	if sidePot != 300_000 {
		t.Errorf(
			"expected 300_000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_3(t *testing.T) {
	sidePot, _ := calculateSidePot(
		500_000,
		200_000,
		2,
	)

	if sidePot != 600_000 {
		t.Errorf(
			"expected 600_000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_4(t *testing.T) {
	sidePot, _ := calculateSidePot(
		1_000_000,
		500_000,
		1,
	)

	if sidePot != 500_000 {
		t.Errorf(
			"expected 500_000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_5(t *testing.T) {
	sidePot, _ := calculateSidePot(
		2000,
		0,
		4,
	)

	if sidePot != 8000 {
		t.Errorf(
			"expected 8000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_6(t *testing.T) {
	sidePot, _ := calculateSidePot(
		4000,
		2000,
		3,
	)

	if sidePot != 6000 {
		t.Errorf(
			"expected 6000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_7(t *testing.T) {
	sidePot, _ := calculateSidePot(
		96000,
		4000,
		4,
	)

	if sidePot != 368_000 {
		t.Errorf(
			"expected 368000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_8(t *testing.T) {
	sidePot, _ := calculateSidePot(
		2_000_000,
		500_000,
		2,
	)

	if sidePot != 3_000_000 {
		t.Errorf(
			"expected 3_000_000; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_9(t *testing.T) {
	sidePot, _ := calculateSidePot(
		500_000,
		500_000,
		3,
	)

	if sidePot != 0 {
		t.Errorf(
			"expected 0; got %d",
			sidePot,
		)
	}
}

func TestCalculateSidePot_10(t *testing.T) {
	sidePot, _ := calculateSidePot(
		500_000,
		100_000,
		0,
	)

	if sidePot != 0 {
		t.Errorf(
			"expected 0; got %d",
			sidePot,
		)
	}
}
