package Managers

import (
	"SwimPractice/Managers/PlayerManager"
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"path/filepath"
)

// SwimPractice our special single instance "base classes" go here
type SwimPractice struct {
	WorldManager       *WorldManager
	SessionDataManager *PlayerManager.PlayerSessionDataManager
	ServerInstance     *server.Server
}

var Swim *SwimPractice

// GlobalMessage send a message in chat globally to each player
func GlobalMessage(message string) {
	for _, player := range Swim.ServerInstance.Players() {
		player.Message(message)
	}
}

// Hub teleports a player to hub and defaults their inventory and game mode etc
func Hub(plr *player.Player) {
	spawnVec3 := mgl64.Vec3{-31, 100, 0}
	TeleportToWorld(plr, "flat", spawnVec3)
	plr.SetGameMode(world.GameModeSurvival) // set them to survival
	plr.Inventory().Clear()                 // clear inventory
	plr.Messagef("§aSent you to the Hub!")
}

// TeleportToWorld Teleport a player to a world by name at a certain position
func TeleportToWorld(plr *player.Player, worldName string, pos mgl64.Vec3) {
	newWorld, ok := Swim.WorldManager.World(worldName)
	if ok {
		newWorld.AddEntity(plr)
		plr.Teleport(cube.PosFromVec3(pos).Vec3Centre())
	}
}

// LoadWorlds load all the worlds on the disk using WorldManager
func LoadWorlds() {
	worldsPath := filepath.Join(filepath.Dir(os.Args[0]), "Worlds")
	// Create the world WorldManager
	Swim.WorldManager = NewWorldManager(Swim.ServerInstance, worldsPath, &logrus.Logger{})
	// List of world folder names
	worldFolders := []string{"PotFFA"}
	// Load each world in the list
	for _, folder := range worldFolders {
		worldFolderPath := filepath.Join(worldsPath, folder)
		worldName := folder
		err := Swim.WorldManager.LoadWorld(worldFolderPath, worldName)
		if err != nil {
			log.Fatalf("Failed to load world %s: %v", worldName, err)
		} else {
			fmt.Println("Loaded world " + worldFolderPath)
		}
	}
}
