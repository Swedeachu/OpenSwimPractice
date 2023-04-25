package Entity

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/block/cube/trace"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/particle"
	"github.com/df-mc/dragonfly/server/world/sound"
	"github.com/go-gl/mathgl/mgl64"
	"time"
)

type SplashPotionType struct {
	Type potion.Potion
}

func (SplashPotionType) EncodeEntity() string { return "minecraft:splash_potion" }
func (SplashPotionType) BBox(world.Entity) cube.BBox {
	return cube.Box(-0.125, 0, -0.125, 0.125, 0.25, 0.125)
}

// NewSwimPot create the swim pot entity
func NewSwimPot(pos mgl64.Vec3, owner world.Entity, potType potion.Potion) *entity.Ent {
	potionConfig := splashPotConfig
	potionConfig.Potion = potType
	potionConfig.Hit = splashPotionHit(0.75, potType)
	return entity.Config{Behaviour: potionConfig.New(owner)}.New(SplashPotionType{Type: potType}, pos)
}

var splashPotConfig = entity.ProjectileBehaviourConfig{
	Gravity: 0.080,            // play around with this value
	Drag:    0.0025,           // play around with this value
	Damage:  -1,               // negative value so we can't get knocked and damaged by it
	Potion:  potion.Awkward(), // default effect type is awkward
	Hit:     nil,              // we set this in NewSwimPot
}

// splashPotionHit small trick to make us get more info than is normally available when the potion hits
func splashPotionHit(radius float64, pot potion.Potion) func(e *entity.Ent, res trace.Result) {
	return func(e *entity.Ent, res trace.Result) {
		// define a quick ignoreFilter filter function
		ignoreFilter := func(ent world.Entity) bool {
			_, living := ent.(entity.Living)
			return living == false || ent == e // can't be a non-living entity
		}
		// make an expanded bounding box
		expansion := mgl64.Vec3{radius, radius, radius}
		// apply the expansion to the bounding box
		splashRadius := res.BBox().Extend(expansion)
		pos := res.Position()
		w := e.Owner().World()
		// iterate all the entities within this expanded bounding box
		for _, e := range w.EntitiesWithin(splashRadius.GrowVec3(expansion.Mul(2)), ignoreFilter) {
			// check if we hit a player
			if plr, ok := e.(*player.Player); ok {
				// go through the effects of the potion
				for _, eff := range pot.Effects() {
					// check if our potion is potent, meaning it applies an instant effect such as instant health, damage, etc
					if potentEffect, ok := eff.Type().(effect.PotentType); ok {
						// if distance of player's eye position to the pot is greater than 2 block the potency is halved
						distance := entity.EyePosition(plr).Sub(pos).Len()
						potency := 1.0 // potency is how much of the potions full effect applies (includes level)
						if distance >= 3.0 {
							potency = 0.5
						}
						plr.AddEffect(effect.NewInstant(potentEffect.WithPotency(potency), 2))
					} else { // else wise it is a lasting effect (poison, regeneration, etc)
						plr.AddEffect(effect.New(eff.Type().(effect.LastingType), 2, 30*time.Second))
					}
				}
			}
		}
		// play the glass break sound in the world at the potion splash
		w.PlaySound(pos, sound.GlassBreak{})
		// spawn particles with color correlating to the potion type
		color, _ := effect.ResultingColour(pot.Effects())
		w.AddParticle(pos, particle.Splash{Colour: color})
	}
}
