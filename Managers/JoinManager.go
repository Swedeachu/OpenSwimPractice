package Managers

import (
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// JoinManager Handles what happens when a player joins the server, is constantly listening for connections in a goroutine
func JoinManager(srv *server.Server) {
	// also this is where we set world's spawn point and properties
	spawnVec3 := mgl64.Vec3{-31, 100, 0}
	spawnPos := cube.PosFromVec3(spawnVec3).Vec3Centre()
	w := srv.World()
	w.SetSpawn(cube.PosFromVec3(spawnPos))
	DefaultWorldSettings(w) // set default settings to the world
	// this works as the join listener for what happens when players first log on
	for srv.Accept(func(plr *player.Player) {
		GlobalMessage("§7[§a+§7] §7" + plr.Name())
		plr.Teleport(spawnPos)                                        // tp player to spawn
		plr.Message("§2 Welcome to Swim Practice! §3discord.gg/swim") // greeting message
		plr.PlaySound(world.Sound(sound.LevelUp{}))                   // join sound for only the player
		HubKit(plr)                                                   // give hub kit
		// clear levels
		plr.SetExperienceProgress(0)
		plr.SetExperienceLevel(0)

		// most important part is handler assigning, we are setting hub behaviors on joining

		// create a handler for the player
		handler := NewPlayerHandler(plr)

		// add the player to the session manager
		Swim.SessionDataManager.AddPlayer(plr)

		// we then set the empty PlayerEventBehavior to have the Hub behaviors
		SetDefaultBehaviors(handler)

		// attach the handler
		plr.Handle(handler)
		Swim.SessionDataManager.SetPlayerHandler(plr, handler)
	}) {
	}
}
