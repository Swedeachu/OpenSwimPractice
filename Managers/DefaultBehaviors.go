package Managers

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/df-mc/dragonfly/server/world"
	"strconv"
	"time"
)

// everything in this file is what the default behaviors will be when playing, and in the hub

// SetDefaultBehaviors sets the default behavior functions to the parameter handler
func SetDefaultBehaviors(handler *PlayerHandler) {
	handler.AddChatBehavior("chat", defaultChatBehavior)
	handler.AddHurtBehavior("hub", defaultHurtBehavior)
	handler.AddItemDropBehavior("default", defaultItemDropBehavior)
	handler.AddFoodLossBehavior("default", defaultFoodLossBehavior)
	handler.AddQuitBehavior("default", defaultQuitBehavior)
	handler.AddBlockBreakBehavior("default", AntiBlockBreakBehavior) // defined in BlockBehaviors
	handler.AddBlockPlaceBehavior("default", AntiBlockPlaceBehavior) // defined in BlockBehaviors
	handler.AddPlayerPunchAirBehavior("default", ClickBehavior)      // defined in AntiCheat
	handler.AddItemUseBehavior("hub forms", defaultUseItemBehavior)
	handler.AddItemUseBehavior("pearl", defaultPearlUseBehavior) // defined in CombatBehaviors
	Swim.SessionDataManager.SetPlayerScoreboardBehavior(handler.plr, defaultScoreboard)
	Swim.SessionDataManager.SetPlayerScoreTagBehavior(handler.plr, defaultScoreTag)
}

// defaultScoreboard the default default scoreboard behavior
func defaultScoreboard(p *player.Player, s *scoreboard.Scoreboard) {
	s.Set(0, " =============\u00A0")
	s.Set(1, fmt.Sprintf(" §bOnline: §f%d§7 / §3%d", GetOnlineCount(), GetMaxPlayerCount()))
	s.Set(2, fmt.Sprintf(" §bPing: §3%d", p.Latency().Milliseconds()))
	s.Set(3, " §bdiscord.gg/§3swim")
	s.Set(4, "\u00A0=============")
	p.SendScoreboard(s)
}

// defaultScoreTag shows the players cps and ping under their name tag
func defaultScoreTag(p *player.Player) {
	cps := GetCPS(p)
	ping := int(p.Latency().Milliseconds())
	p.SetScoreTag("§3" + strconv.Itoa(cps) + " CPS §7| §3" + strconv.Itoa(ping) + " MS")
}

// defaultChatBehavior print the message in chat in a different format with player's rank
func defaultChatBehavior(plr *player.Player, ctx *event.Context, message *string) {
	// fmt.Println(*message)
}

// defaultUseItemBehavior controls the form ui pop-ups in the hub
func defaultUseItemBehavior(plr *player.Player, ctx *event.Context) {
	heldItem, _ := plr.HeldItems()
	if heldItem.CustomName() == "§bFFA §7[Right Click]" {
		plr.SendForm(FormFFA())
	}
}

// defaultHurtBehavior all damage is cancelled in the default
func defaultHurtBehavior(plr *player.Player, ctx *event.Context, damage *float64, duration *time.Duration, source world.DamageSource) {
	ctx.Cancel()
}

// defaultItemDropBehavior prevents the player from dropping items
func defaultItemDropBehavior(plr *player.Player, ctx *event.Context, e *entity.Item) {
	ctx.Cancel()
}

// defaultFoodLossBehavior cancel hunger loss
func defaultFoodLossBehavior(plr *player.Player, ctx *event.Context, from int, to *int) {
	ctx.Cancel()
}

// defaultQuitBehavior removes player from any data structures they are in
func defaultQuitBehavior(plr *player.Player) {
	Swim.SessionDataManager.RemovePlayer(plr)
}
