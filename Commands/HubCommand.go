package Commands

import (
	"SwimPractice/Managers"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

// HubCommand contains our command parameters (none for this one)
type HubCommand struct {
}

// Run will be called when the player runs the command
func (c HubCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		Managers.Hub(plr)
	}
}
