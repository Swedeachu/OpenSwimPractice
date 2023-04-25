package Entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
	_ "unsafe"
)

type PearlType struct {
}

func (PearlType) EncodeEntity() string { return "minecraft:ender_pearl" }
func (PearlType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

// NewSwimPearl creates the Swim Pearl entity
func NewSwimPearl(pos mgl64.Vec3, owner world.Entity) *entity.Ent {
	return entity.Config{Behaviour: pearlConfig.New(owner)}.New(PearlType{}, pos)
}

var pearlConfig = entity.ProjectileBehaviourConfig{
	Gravity:               0.03,   // play around with this value
	Drag:                  0.0025, // play around with this value
	KnockBackForceAddend:  0.25,   // knock back force on entity hit
	KnockBackHeightAddend: 0.25,   // height knock back on entity hit
	Particle:              particle.EndermanTeleport{},
	Sound:                 sound.Teleport{},
	Hit:                   pearlHit,
}

// pearlHit what happens when the projectile hits another entity or block
func pearlHit(pearl *entity.Ent, target trace.Result) {
	// get the owning entity of the pearl to teleport them to the trace result
	if plr, ok := pearl.Owner().(*player.Player); ok {
		plr.Teleport(target.Position())
	}
}
