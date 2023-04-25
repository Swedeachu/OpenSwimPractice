package Managers

import (
	"SwimPractice/Items"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/entity/effect"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/item/enchantment"
	"github.com/df-mc/dragonfly/server/item/potion"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"time"
)

// ClearEffects removes all the active effects on the player
func ClearEffects(p *player.Player) {
	for _, e := range p.Effects() {
		p.RemoveEffect(e.Type())
	}
}

// ClearPlayer removes effects and clears inventory
func ClearPlayer(p *player.Player) {
	p.Inventory().Clear()
	p.Armour().Clear()
	p.Heal(p.MaxHealth(), entity.FoodHealingSource{})
	ClearEffects(p)
}

// HubKit the default hub kit for the player
func HubKit(p *player.Player) {
	ClearPlayer(p)
	p.SetGameMode(world.GameModeSurvival)
	p.Inventory().SetItem(0, item.NewStack(item.Sword{Tier: item.ToolTierDiamond}, 1).WithCustomName("§bFFA §7[Right Click]"))
}

// PotKit the default pot kit for the player
func PotKit(p *player.Player) {
	ClearPlayer(p)
	p.SetGameMode(world.GameModeSurvival)
	p.AddEffect(effect.New(effect.Speed{}, 1, 24*time.Hour).WithoutParticles())
	diamondArmor(p) // give diamond armor
	p.Inventory().SetItem(0, item.NewStack(item.Sword{Tier: item.ToolTierDiamond}, 1).WithEnchantments(item.NewEnchantment(enchantment.Unbreaking{}, 10)))
	p.Inventory().SetItem(1, item.NewStack(Items.SwimPearl{}, 16).WithCustomName("§3Pearl"))
	// p.Inventory().AddItem(item.NewStack(item.SplashPotion{Type: potion.Healing()}, 34))
	p.Inventory().AddItem(item.NewStack(Items.SwimPot{Type: potion.Healing()}, 34))
}

// diamondArmor gives a player full diamond armor
func diamondArmor(p *player.Player) {
	unbreaking := item.NewEnchantment(enchantment.Unbreaking{}, 10)
	p.Armour().SetHelmet(item.NewStack(item.Helmet{Tier: item.ArmourTierDiamond{}}, 1).WithEnchantments(unbreaking))
	p.Armour().SetChestplate(item.NewStack(item.Chestplate{Tier: item.ArmourTierDiamond{}}, 1).WithEnchantments(unbreaking))
	p.Armour().SetLeggings(item.NewStack(item.Leggings{Tier: item.ArmourTierDiamond{}}, 1).WithEnchantments(unbreaking))
	p.Armour().SetBoots(item.NewStack(item.Boots{Tier: item.ArmourTierDiamond{}}, 1).WithEnchantments(unbreaking))
}
