package Commands

import (
	"SwimPractice/Managers"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"strings"
)

// WorldChangeCommand contains our command parameters
type WorldChangeCommand struct {
	Message cmd.Varargs `cmd:"world name"`
}

// Run will be called when the player runs the command
func (a WorldChangeCommand) Run(source cmd.Source, output *cmd.Output) {
	if plr, ok := source.(*player.Player); ok {
		msg := strings.TrimSpace(string(a.Message))
		if len(msg) == 0 {
			// empty
			return
		}
		Managers.TeleportToWorld(plr, msg, mgl64.Vec3{0, 100, 0})
	}
}
