package Items

import (
	"SwimPractice/Entity"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/sound"
)

// SwimPot nearly the same as the default potion just with adjusted physics and splash radius
type SwimPot struct {
	Type potion.Potion
}

// Use creates a Swim pot entity that has adjusted velocity and physics and splash radius
func (s SwimPot) Use(w *world.World, user item.User, ctx *item.UseContext) bool {
	// spawn a pot at the eye position of the user, and set the pot's velocity to a scaled direction vector of the user
	pos := entity.EyePosition(user)
	e := Entity.NewSwimPot(pos, user, s.Type)
	w.AddEntity(e)
	e.SetVelocity(user.Rotation().Vec3().Mul(0.5))
	// play the sound and consume the item
	w.PlaySound(pos, sound.ItemThrow{})
	ctx.SubtractFromCount(1)
	return true
}

// MaxCount vanilla
func (SwimPot) MaxCount() int {
	return 1
}

// EncodeItem encode it with proper meta
func (s SwimPot) EncodeItem() (name string, meta int16) {
	return "minecraft:splash_potion", int16(s.Type.Uint8())
}
