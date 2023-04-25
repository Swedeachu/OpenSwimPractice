package Managers

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/sirupsen/logrus"
	"log"
	"math/rand"
	"os"
	"path/filepath"
)

// SwimPractice our special single instance "base classes" go here
type SwimPractice struct {
	WorldManager       *WorldManager
	SessionDataManager *PlayerSessionDataManager
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
	// first reset behaviors back to default
	handler := Swim.SessionDataManager.Players[plr].Handler
	handler.ClearPlayerBehaviors()
	SetDefaultBehaviors(handler)
	// teleport back to hub
	spawnVec3 := mgl64.Vec3{-31, 100, 0}
	TeleportToWorld(plr, "flat", spawnVec3)
	// give hub kit
	HubKit(plr)
	// set levels back to 0
	plr.SetExperienceProgress(0)
	plr.SetExperienceLevel(0)
	// get rid of any cool downs
	data := Swim.SessionDataManager.Players[plr]
	data.PlayerDataLock()
	defer data.PlayerDataUnlock()
	data.PearlCoolDown = 0
	// message
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

// VectorRandOffset3D offsets a 3D vectors x and z values by a random set range
func VectorRandOffset3D(x int, y int, z int, offsetX int, offsetZ int) mgl64.Vec3 {
	offSetX := rand.Intn(offsetX+1) - offsetX/2
	offSetZ := rand.Intn(offsetZ+1) - offsetZ/2
	return mgl64.Vec3{float64(x + offSetX), float64(y), float64(z + offSetZ)}
}

// GetOnlineCount get the amount of players online
func GetOnlineCount() int {
	return len(Swim.ServerInstance.Players())
}

// GetMaxPlayerCount get the max amount of players allowed to join the server
func GetMaxPlayerCount() int {
	return Swim.ServerInstance.MaxPlayerCount()
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
