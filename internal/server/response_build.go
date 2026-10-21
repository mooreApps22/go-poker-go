package server

import (
	"fmt"
	"github.com/mooreApps22/go-poker-go/internal/poker"
)

func buildActionResponse(playerAction poker.PlayerAction) string {
	if playerAction.Amount > 0 {
		return fmt.Sprintf(
			"Action received: %s $%d\n",
			playerAction.Action.String(),
			playerAction.Amount,
		)
	}

	return fmt.Sprintf(
		"Action received: %s\n",
		playerAction.Action.String(),
	)
}
