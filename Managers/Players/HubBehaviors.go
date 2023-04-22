package Players

import (
	"SwimPractice/Managers"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"github.com/df-mc/dragonfly/server/world"
	"math/rand"
	"strconv"
	"time"
)

// SetHubBehaviors sets the hub behavior functions to the parameter handler
func SetHubBehaviors(handler *PlayerHandler) {
	handler.SetChatBehavior(HubChatBehavior)
	handler.SetHurtBehavior(HubHurtBehavior)
	handler.SetItemDropBehavior(HubItemDropBehavior)
	handler.SetFoodLossBehavior(HubFoodLossBehavior)
	handler.SetQuitBehavior(HubQuitBehavior)
	handler.SetBlockBreakBehavior(AntiBlockBreakBehavior) // defined in BlockBehaviors
	handler.SetBlockPlaceBehavior(AntiBlockPlaceBehavior) // defined in BlockBehaviors
	Managers.Swim.SessionDataManager.SetPlayerScoreboardBehavior(handler.plr, HubScoreboard)
}

func HubScoreboard(p *player.Player) {
	s := scoreboard.New("Title")
	s.Set(0, "Hello, "+strconv.Itoa(rand.Intn(100)))
	s.Set(1, "World!")
	p.SendScoreboard(s)
}

// HubChatBehavior print the message in chat in a different format with player's rank
func HubChatBehavior(ctx *event.Context, message *string) {
	// fmt.Println(*message)
}

// HubHurtBehavior all damage is cancelled in the hub
func HubHurtBehavior(ctx *event.Context, damage *float64, duration *time.Duration, source world.DamageSource) {
	ctx.Cancel()
}

// HubItemDropBehavior prevents the player from dropping items
func HubItemDropBehavior(ctx *event.Context, e *entity.Item) {
	ctx.Cancel()
}

// HubFoodLossBehavior cancel hunger loss
func HubFoodLossBehavior(ctx *event.Context, from int, to *int) {
	ctx.Cancel()
}

// HubQuitBehavior removes player from any data structures they are in
func HubQuitBehavior(plr *player.Player) {
	Managers.Swim.SessionDataManager.RemovePlayer(plr)
}
