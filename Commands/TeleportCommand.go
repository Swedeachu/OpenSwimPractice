package Commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
)

// TeleportCommand contains our command parameters.
type TeleportCommand struct {
	X int `cmd:"X"`
	Y int `cmd:"Y"`
	Z int `cmd:"Z"`
}

// Run will be called when the player runs the command
func (c TeleportCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		newPos := mgl64.Vec3{float64(c.X), float64(c.Y), float64(c.Z)} // make 3D vector from args
		plr.Teleport(newPos)                                           // teleport player to that vector
		plr.Messagef("Teleported to %d %d %d", c.X, c.Y, c.Z)
	} else {
		output.Printf("This command can only be run by a player")
	}
}
