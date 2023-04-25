package Items

import (
	"SwimPractice/Entity"
	_ "embed"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
	"time"
)

// SwimPearl basically the same as the default pearl just custom air travel settings for its drag and terminal velocity
type SwimPearl struct{}

// Use create a swim pearl entity when using the swim pearl item
func (SwimPearl) Use(w *world.World, user item.User, ctx *item.UseContext) bool {
	// spawn a pearl at the eye position of the user, and set the pearl's velocity to a scaled direction vector of the user
	pos := entity.EyePosition(user)
	e := Entity.NewSwimPearl(pos, user)
	w.AddEntity(e)
	e.SetVelocity(user.Rotation().Vec3().Mul(2))
	// play the sound and consume the item
	w.PlaySound(pos, sound.ItemThrow{})
	ctx.SubtractFromCount(1)
	return true
}

// Cooldown vanilla
func (SwimPearl) Cooldown() time.Duration {
	return time.Second
}

// MaxCount vanilla
func (SwimPearl) MaxCount() int {
	return 16
}

// EncodeItem encode it
func (SwimPearl) EncodeItem() (name string, meta int16) {
	return "minecraft:ender_pearl", 0
}
