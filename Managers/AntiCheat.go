package Managers

import (
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"strconv"
	"time"
)

// ClickBehavior append the current unix timestamp to the player's clicks array
func ClickBehavior(plr *player.Player, ctx *event.Context) {
	session := Swim.SessionDataManager
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	now := time.Now().UnixMilli() // get the unix timestamp
	*session.Players[plr].ClicksArray = append(*session.Players[plr].ClicksArray, now)
	// set the cpsTag based on the CPS value
	cps := GetCPS(plr)
	var cpsTag string
	if cps >= 13 && cps < 15 {
		cpsTag = "§e" + strconv.Itoa(cps)
	} else if cps >= 15 && cps < 18 {
		cpsTag = "§6" + strconv.Itoa(cps)
	} else if cps >= 18 {
		cpsTag = "§c" + strconv.Itoa(cps)
	} else {
		cpsTag = "§a" + strconv.Itoa(cps)
	}
	plr.SendPopup("§b"+cpsTag+" §3CPS", "")
}

// UpdateCPS removes any clicks older than a second from a player's clicks array
func UpdateCPS(plr *player.Player) {
	clicks := *Swim.SessionDataManager.Players[plr].ClicksArray // get the clicks array they have
	now := time.Now().UnixMilli()
	for len(clicks) > 0 && now >= clicks[0]+1000 {
		clicks = clicks[1:] // remove clicks older than 1 second
	}
	*Swim.SessionDataManager.Players[plr].ClicksArray = clicks
}

// GetCPS gets the clicks from the player's clicks array
func GetCPS(plr *player.Player) int {
	return len(*Swim.SessionDataManager.Players[plr].ClicksArray)
}
