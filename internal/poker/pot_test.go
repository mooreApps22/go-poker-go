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
