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
		if numberOfPlayerWithLevel == 1 {
			player := hand.getPlayerAtOrAboveContributionLevel(level)
			if player != nil {
				player.CollectWinnings(sidePot)
			}

			continue
		}

		sidePot, previousLevel = calculateSidePot(
			level,
			previousLevel,
			int(numberOfPlayerWithLevel),
		)

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
