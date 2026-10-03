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

func (hand *Hand) addEligiblePlayersToPot(level int64, potIndex int) {
	for _, player := range hand.players {
		if player.totalContribution >= level && !player.hasFolded {
			hand.pots[potIndex].eligiblePlayers = append(
				hand.pots[potIndex].eligiblePlayers,
				player,
			)
		}
	}
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

func (hand *Hand) CreateSidePots() {
	hand.collectContributionLevels()
	var sidePot int64
	var previousLevel int64

	for potIndex, level := range hand.contributionLevels {
		numberOfPlayerWithLevel := hand.countPlayersWithContributionLevel(level)
		sidePot, previousLevel = calculateSidePot(
			level,
			previousLevel,
			int(numberOfPlayerWithLevel),
		)
		hand.pots = append(hand.pots, Pot{
			value: sidePot,
		})
		hand.addEligiblePlayersToPot(level, potIndex)
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

func (hand *Hand) AwardPot() {
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
