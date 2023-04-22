package Commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

// CordsCommand contains our command parameters (none for this one)
type CordsCommand struct {
}

// Run will be called when the player runs the command
func (c CordsCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		pos := plr.Position()
		plr.Messagef("Your position is %d %d %d", pos.X(), pos.Y(), pos.Z())
	} else {
		output.Printf("This command can only be run by a player")
	}
}
