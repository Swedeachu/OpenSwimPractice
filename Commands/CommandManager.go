package Commands

import (
	"github.com/df-mc/dragonfly/server/cmd"
)

// RegisterAllCommands this registers all the commands
func RegisterAllCommands() {
	cmd.Register(cmd.New("teleport", "Teleport to an X Y Z Position in the current World", []string{"tp"}, TeleportCommand{}))
	cmd.Register(cmd.New("coordinates", "Get your position in the World", []string{"cord", "cords", "c"}, CordsCommand{}))
	cmd.Register(cmd.New("fly", "Toggles flying on/off", []string{}, FlyCommand{}))
	cmd.Register(cmd.New("announcement", "Announce a message globally", []string{"announce", "shout"}, AnnouncementCommand{}))
	cmd.Register(cmd.New("hub", "Teleports you back to the hub", []string{"spawn"}, HubCommand{}))
	cmd.Register(cmd.New("world", "Teleports you to (0, 0) in a world, does nothing if the world does not exist", []string{}, WorldChangeCommand{}))
}
