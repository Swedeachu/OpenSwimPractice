package Commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
)

type flyMode struct {
	world.GameMode
}

func (flyMode) AllowsFlying() bool { return true }

// FlyCommand contains our command parameters (none for this one)
type FlyCommand struct {
}

// Run will be called when the player runs the command
func (c FlyCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		if plr.Flying() {
			plr.SetGameMode(world.GameModeSurvival)
			plr.StopFlying()
			plr.Messagef("Stopped flying")
		} else {
			plr.SetGameMode(flyMode{GameMode: plr.GameMode()})
			plr.StartFlying()
			plr.Messagef("Started flying")
		}
	} else {
		output.Printf("This command can only be run by a player")
	}
}
