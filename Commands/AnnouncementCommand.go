package Commands

import (
	"SwimPractice/Managers"
	"github.com/df-mc/dragonfly/server/cmd"
	"strings"
)

// AnnouncementCommand contains our command parameters
type AnnouncementCommand struct {
	Message cmd.Varargs `cmd:"message"`
}

// Run will be called when the player runs the command
func (a AnnouncementCommand) Run(source cmd.Source, output *cmd.Output) {
	msg := strings.TrimSpace(string(a.Message))
	if len(msg) == 0 {
		// empty
		return
	}
	Managers.GlobalMessage("§3§lANNOUNCEMENT:§r§b \n" + strings.ReplaceAll(msg, "\\n", "\n"))
}
