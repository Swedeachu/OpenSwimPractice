package Players

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
)

// AntiBlockBreakBehavior prevents the player from breaking any blocks
func AntiBlockBreakBehavior(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	ctx.Cancel()
}

// AntiBlockPlaceBehavior prevents the player from placing blocks
func AntiBlockPlaceBehavior(ctx *event.Context, pos cube.Pos, b world.Block) {
	ctx.Cancel()
}
