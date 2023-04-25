package Managers

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
)

// don't have a reason to make this scalable yet, since there isn't that much to even override right now

// WorldHandler any world events we want to override, such as sounds
type WorldHandler struct {
	world.NopHandler
}

// HandleSound cancels the weird swing sound bedrock has
func (h WorldHandler) HandleSound(ctx *event.Context, s world.Sound, _ mgl64.Vec3) {
	if _, ok := s.(sound.Attack); ok {
		ctx.Cancel()
	}
}

// HandleFireSpread cancels the fire spread event
func (h WorldHandler) HandleFireSpread(ctx *event.Context, from, to cube.Pos) {
	ctx.Cancel()
}

// HandleBlockBurn cancels the block burn event
func (h WorldHandler) HandleBlockBurn(ctx *event.Context, pos cube.Pos) {
	ctx.Cancel()
}
