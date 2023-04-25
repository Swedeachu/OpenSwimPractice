package Commands

import (
	"SwimPractice/Managers"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
)

// KnockBackCommand contains our command parameters.
type KnockBackCommand struct {
	Force  float64 `cmd:"Force"`
	Height float64 `cmd:"Height"`
}

// Run will be called when the player runs the command
func (c KnockBackCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		Managers.Force = c.Force
		Managers.Height = c.Height
		plr.Messagef("Set Force to %f and Height to %f", c.Force, c.Height)
	}
}
