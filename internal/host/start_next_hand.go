package host

func (host *Host) StartHand() error {
	host.game.SetUpNewHand()

	hand := host.game.GetHand()

	hand.ResetPlayersForNewHand()
	hand.PostBlinds()

	err := hand.DealHoleCards()
	if err != nil {
		return err
	}

	return nil
}

func (host *Host) SetWaitingForDeal(waiting bool) {
	host.waitingForDeal = waiting
}

func (host *Host) IsWaitingForDeal() bool {
	return host.waitingForDeal
}
