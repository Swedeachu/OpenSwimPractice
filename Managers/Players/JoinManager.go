package Players

import (
	"SwimPractice/Managers"
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
	Managers.DefaultWorldSettings(w) // set default settings to the world
	// this works as the join listener for what happens when players first log on
	for srv.Accept(func(plr *player.Player) {
		Managers.GlobalMessage("§7[§a+§7] §7" + plr.Name())
		plr.Teleport(spawnPos)                                                // tp player to spawn
		plr.SetGameMode(world.GameModeSurvival)                               // set them to survival
		plr.Inventory().Clear()                                               // clear inventory
		plr.Message("§2 Welcome to SwimPractice Practice! §3discord.gg/swim") // greeting message
		plr.PlaySound(world.Sound(sound.LevelUp{}))                           // join sound for only the player

		// most important part is handler assigning, we are setting hub behaviors on joining

		// create a handler for the player with an empty PlayerEventBehavior struct
		handler := NewPlayerHandler(plr, &PlayerEventBehavior{})

		// add the player to the session manager
		Managers.Swim.SessionDataManager.AddPlayer(plr)

		// we then set the empty PlayerEventBehavior to have the Hub behaviors
		SetHubBehaviors(handler)

		// attach the handler
		plr.Handle(handler)
	}) {
	}
}
