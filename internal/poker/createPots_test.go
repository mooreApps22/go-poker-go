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
