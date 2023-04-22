package Players

// this file is mainly auto generated boilerplate code, do not modify

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/event"
	"github.com/df-mc/dragonfly/server/item"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/skin"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"net"
	"time"
)

// The way this works is with Behavior setting functions, we will have functions in a PlayerBehavior struct,
// where we can easily change what will happen when each of these events is called on the player object at run time.
// Virtual functions with some gopher styling

// PlayerEventBehavior this struct is functions for each behavior to be overwritten virtually at run time
type PlayerEventBehavior struct {
	PlayerMoveBehavior             func(*event.Context, mgl64.Vec3, float64, float64)
	PlayerJumpBehavior             func()
	PlayerTeleportBehavior         func(*event.Context, mgl64.Vec3)
	PlayerChangeWorldBehavior      func(*world.World, *world.World)
	PlayerToggleSprintBehavior     func(*event.Context, bool)
	PlayerToggleSneakBehavior      func(*event.Context, bool)
	PlayerChatBehavior             func(*event.Context, *string)
	PlayerFoodLossBehavior         func(*event.Context, int, *int)
	PlayerHealBehavior             func(*event.Context, *float64, world.HealingSource)
	PlayerHurtBehavior             func(*event.Context, *float64, *time.Duration, world.DamageSource)
	PlayerDeathBehavior            func(world.DamageSource, *bool)
	PlayerRespawnBehavior          func(*mgl64.Vec3, **world.World)
	PlayerSkinChangeBehavior       func(*event.Context, *skin.Skin)
	PlayerStartBreakBehavior       func(*event.Context, cube.Pos)
	PlayerBlockBreakBehavior       func(*event.Context, cube.Pos, *[]item.Stack, *int)
	PlayerBlockPlaceBehavior       func(*event.Context, cube.Pos, world.Block)
	PlayerBlockPickBehavior        func(*event.Context, cube.Pos, world.Block)
	PlayerItemUseBehavior          func(*event.Context)
	PlayerItemUseOnBlockBehavior   func(*event.Context, cube.Pos, cube.Face, mgl64.Vec3)
	PlayerItemUseOnEntityBehavior  func(*event.Context, world.Entity)
	PlayerItemConsumeBehavior      func(*event.Context, item.Stack)
	PlayerAttackEntityBehavior     func(*event.Context, world.Entity, *float64, *float64, *bool)
	PlayerExperienceGainBehavior   func(*event.Context, *int)
	PlayerPunchAirBehavior         func(*event.Context)
	PlayerSignEditBehavior         func(*event.Context, string, string)
	PlayerItemDamageBehavior       func(*event.Context, item.Stack, int)
	PlayerItemPickupBehavior       func(*event.Context, item.Stack)
	PlayerItemDropBehavior         func(*event.Context, *entity.Item)
	PlayerTransferBehavior         func(*event.Context, *net.UDPAddr)
	PlayerCommandExecutionBehavior func(*event.Context, cmd.Command, []string)
	PlayerQuitBehavior             func(*player.Player)
}

// PlayerHandler this struct defines the handler to have optional overwrites for the event interfaces with our behaviors
type PlayerHandler struct {
	player.NopHandler
	plr            *player.Player
	eventBehaviors *PlayerEventBehavior
}

// NewPlayerHandler returns a new PlayerHandler for our event behaviors
func NewPlayerHandler(plr *player.Player, behaviors *PlayerEventBehavior) *PlayerHandler {
	return &PlayerHandler{plr: plr, eventBehaviors: behaviors}
}

// ClearPlayerBehaviors sets all event behavior functions in the PlayerHandler's PlayerEventBehavior to nil
func ClearPlayerBehaviors(handler *PlayerHandler) {
	handler.eventBehaviors = &PlayerEventBehavior{}
}

/*@@@@@@@@@@@@@@@@@@@@@@@@@
     BEHAVIOR SETTING!!!
@@@@@@@@@@@@@@@@@@@@@@@@@@@*/

// SetMoveBehavior sets a new move behavior for the player handler
func (handler *PlayerHandler) SetMoveBehavior(moveBehavior func(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64)) {
	handler.eventBehaviors.PlayerMoveBehavior = moveBehavior
}

// SetJumpBehavior sets a new jump behavior for the player handler
func (handler *PlayerHandler) SetJumpBehavior(jumpBehavior func()) {
	handler.eventBehaviors.PlayerJumpBehavior = jumpBehavior
}

// SetTeleportBehavior sets a new teleport behavior for the player handler
func (handler *PlayerHandler) SetTeleportBehavior(teleportBehavior func(ctx *event.Context, pos mgl64.Vec3)) {
	handler.eventBehaviors.PlayerTeleportBehavior = teleportBehavior
}

// SetChangeWorldBehavior sets a new change world behavior for the player handler
func (handler *PlayerHandler) SetChangeWorldBehavior(changeWorldBehavior func(before, after *world.World)) {
	handler.eventBehaviors.PlayerChangeWorldBehavior = changeWorldBehavior
}

// SetToggleSprintBehavior sets a new toggle sprint behavior for the player handler
func (handler *PlayerHandler) SetToggleSprintBehavior(toggleSprintBehavior func(ctx *event.Context, after bool)) {
	handler.eventBehaviors.PlayerToggleSprintBehavior = toggleSprintBehavior
}

// SetToggleSneakBehavior sets a new toggle sneak behavior for the player handler
func (handler *PlayerHandler) SetToggleSneakBehavior(toggleSneakBehavior func(ctx *event.Context, after bool)) {
	handler.eventBehaviors.PlayerToggleSneakBehavior = toggleSneakBehavior
}

// SetChatBehavior sets a new chat behavior for the player handler
func (handler *PlayerHandler) SetChatBehavior(chatBehavior func(ctx *event.Context, message *string)) {
	handler.eventBehaviors.PlayerChatBehavior = chatBehavior
}

// SetFoodLossBehavior sets a new food loss behavior for the player handler
func (handler *PlayerHandler) SetFoodLossBehavior(foodLossBehavior func(ctx *event.Context, from int, to *int)) {
	handler.eventBehaviors.PlayerFoodLossBehavior = foodLossBehavior
}

// SetHealBehavior sets a new heal behavior for the player handler
func (handler *PlayerHandler) SetHealBehavior(healBehavior func(ctx *event.Context, health *float64, src world.HealingSource)) {
	handler.eventBehaviors.PlayerHealBehavior = healBehavior
}

// SetHurtBehavior sets a new hurt behavior for the player handler
func (handler *PlayerHandler) SetHurtBehavior(hurtBehavior func(ctx *event.Context, damage *float64, duration *time.Duration, source world.DamageSource)) {
	handler.eventBehaviors.PlayerHurtBehavior = hurtBehavior
}

// SetDeathBehavior sets a new death behavior for the player handler
func (handler *PlayerHandler) SetDeathBehavior(deathBehavior func(src world.DamageSource, keepInv *bool)) {
	handler.eventBehaviors.PlayerDeathBehavior = deathBehavior
}

// SetRespawnBehavior sets a new respawn behavior for the player handler
func (handler *PlayerHandler) SetRespawnBehavior(respawnBehavior func(pos *mgl64.Vec3, w **world.World)) {
	handler.eventBehaviors.PlayerRespawnBehavior = respawnBehavior
}

// SetSkinChangeBehavior sets a new skin change behavior for the player handler
func (handler *PlayerHandler) SetSkinChangeBehavior(skinChangeBehavior func(ctx *event.Context, skin *skin.Skin)) {
	handler.eventBehaviors.PlayerSkinChangeBehavior = skinChangeBehavior
}

// SetStartBreakBehavior sets a new start break behavior for the player handler
func (handler *PlayerHandler) SetStartBreakBehavior(startBreakBehavior func(ctx *event.Context, pos cube.Pos)) {
	handler.eventBehaviors.PlayerStartBreakBehavior = startBreakBehavior
}

// SetBlockBreakBehavior sets a new block break behavior for the player handler
func (handler *PlayerHandler) SetBlockBreakBehavior(blockBreakBehavior func(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int)) {
	handler.eventBehaviors.PlayerBlockBreakBehavior = blockBreakBehavior
}

// SetBlockPlaceBehavior sets a new block place behavior for the player handler
func (handler *PlayerHandler) SetBlockPlaceBehavior(blockPlaceBehavior func(ctx *event.Context, pos cube.Pos, b world.Block)) {
	handler.eventBehaviors.PlayerBlockPlaceBehavior = blockPlaceBehavior
}

// SetBlockPickBehavior sets a new block pick behavior for the player handler
func (handler *PlayerHandler) SetBlockPickBehavior(blockPickBehavior func(ctx *event.Context, pos cube.Pos, b world.Block)) {
	handler.eventBehaviors.PlayerBlockPickBehavior = blockPickBehavior
}

// SetItemUseBehavior sets a new item use behavior for the player handler
func (handler *PlayerHandler) SetItemUseBehavior(itemUseBehavior func(ctx *event.Context)) {
	handler.eventBehaviors.PlayerItemUseBehavior = itemUseBehavior
}

// SetItemUseOnBlockBehavior sets a new item use on block behavior for the player handler
func (handler *PlayerHandler) SetItemUseOnBlockBehavior(itemUseOnBlockBehavior func(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3)) {
	handler.eventBehaviors.PlayerItemUseOnBlockBehavior = itemUseOnBlockBehavior
}

// SetItemUseOnEntityBehavior sets a new item use on entity behavior for the player handler
func (handler *PlayerHandler) SetItemUseOnEntityBehavior(itemUseOnEntityBehavior func(ctx *event.Context, e world.Entity)) {
	handler.eventBehaviors.PlayerItemUseOnEntityBehavior = itemUseOnEntityBehavior
}

// SetItemConsumeBehavior sets a new item consume behavior for the player handler
func (handler *PlayerHandler) SetItemConsumeBehavior(itemConsumeBehavior func(ctx *event.Context, item item.Stack)) {
	handler.eventBehaviors.PlayerItemConsumeBehavior = itemConsumeBehavior
}

// SetAttackEntityBehavior sets a new attack entity behavior for the player handler
func (handler *PlayerHandler) SetAttackEntityBehavior(attackEntityBehavior func(ctx *event.Context, e world.Entity, force, height *float64, critical *bool)) {
	handler.eventBehaviors.PlayerAttackEntityBehavior = attackEntityBehavior
}

// SetExperienceGainBehavior sets a new experience gain behavior for the player handler
func (handler *PlayerHandler) SetExperienceGainBehavior(experienceGainBehavior func(ctx *event.Context, amount *int)) {
	handler.eventBehaviors.PlayerExperienceGainBehavior = experienceGainBehavior
}

// SetPunchAirBehavior sets a new punch air behavior for the player handler
func (handler *PlayerHandler) SetPunchAirBehavior(punchAirBehavior func(ctx *event.Context)) {
	handler.eventBehaviors.PlayerPunchAirBehavior = punchAirBehavior
}

// SetSignEditBehavior sets a new sign edit behavior for the player handler
func (handler *PlayerHandler) SetSignEditBehavior(signEditBehavior func(ctx *event.Context, oldText, newText string)) {
	handler.eventBehaviors.PlayerSignEditBehavior = signEditBehavior
}

// SetItemDamageBehavior sets a new item damage behavior for the player handler
func (handler *PlayerHandler) SetItemDamageBehavior(itemDamageBehavior func(ctx *event.Context, i item.Stack, damage int)) {
	handler.eventBehaviors.PlayerItemDamageBehavior = itemDamageBehavior
}

// SetItemPickupBehavior sets a new item pickup behavior for the player handler
func (handler *PlayerHandler) SetItemPickupBehavior(itemPickupBehavior func(ctx *event.Context, i item.Stack)) {
	handler.eventBehaviors.PlayerItemPickupBehavior = itemPickupBehavior
}

// SetItemDropBehavior sets a new item drop behavior for the player handler
func (handler *PlayerHandler) SetItemDropBehavior(itemDropBehavior func(ctx *event.Context, e *entity.Item)) {
	handler.eventBehaviors.PlayerItemDropBehavior = itemDropBehavior
}

// SetTransferBehavior sets a new transfer behavior for the player handler
func (handler *PlayerHandler) SetTransferBehavior(transferBehavior func(ctx *event.Context, addr *net.UDPAddr)) {
	handler.eventBehaviors.PlayerTransferBehavior = transferBehavior
}

// SetCommandExecutionBehavior sets a new command execution behavior for the player handler
func (handler *PlayerHandler) SetCommandExecutionBehavior(commandExecutionBehavior func(ctx *event.Context, command cmd.Command, args []string)) {
	handler.eventBehaviors.PlayerCommandExecutionBehavior = commandExecutionBehavior
}

// SetQuitBehavior sets a new quit behavior for the player handler
func (handler *PlayerHandler) SetQuitBehavior(quitBehavior func(plr *player.Player)) {
	handler.eventBehaviors.PlayerQuitBehavior = quitBehavior
}

/*@@@@@@@@@@@@@@@@@@@@@@@@@
  CALLBACK BOILER PLATING!!
@@@@@@@@@@@@@@@@@@@@@@@@@@@*/

// HandleMove virtual override to the player's move behavior
func (handler *PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	if handler.eventBehaviors.PlayerMoveBehavior != nil {
		handler.eventBehaviors.PlayerMoveBehavior(ctx, newPos, newYaw, newPitch)
	}
}

// HandleJump virtual override to the player's jump behavior
func (handler *PlayerHandler) HandleJump() {
	if handler.eventBehaviors.PlayerJumpBehavior != nil {
		handler.eventBehaviors.PlayerJumpBehavior()
	}
}

// HandleTeleport virtual override to the player's teleport behavior
func (handler *PlayerHandler) HandleTeleport(ctx *event.Context, pos mgl64.Vec3) {
	if handler.eventBehaviors.PlayerTeleportBehavior != nil {
		handler.eventBehaviors.PlayerTeleportBehavior(ctx, pos)
	}
}

// HandleChangeWorld virtual override to the player's change world behavior
func (handler *PlayerHandler) HandleChangeWorld(before, after *world.World) {
	if handler.eventBehaviors.PlayerChangeWorldBehavior != nil {
		handler.eventBehaviors.PlayerChangeWorldBehavior(before, after)
	}
}

// HandleToggleSprint virtual override to the player's toggle sprint behavior
func (handler *PlayerHandler) HandleToggleSprint(ctx *event.Context, after bool) {
	if handler.eventBehaviors.PlayerToggleSprintBehavior != nil {
		handler.eventBehaviors.PlayerToggleSprintBehavior(ctx, after)
	}
}

// HandleToggleSneak virtual override to the player's toggle sneak behavior
func (handler *PlayerHandler) HandleToggleSneak(ctx *event.Context, after bool) {
	if handler.eventBehaviors.PlayerToggleSneakBehavior != nil {
		handler.eventBehaviors.PlayerToggleSneakBehavior(ctx, after)
	}
}

// HandleChat virtual override to the player's chat behavior
func (handler *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	if handler.eventBehaviors.PlayerChatBehavior != nil {
		handler.eventBehaviors.PlayerChatBehavior(ctx, message)
	}
}

// HandleFoodLoss virtual override to the player's food loss behavior
func (handler *PlayerHandler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	if handler.eventBehaviors.PlayerFoodLossBehavior != nil {
		handler.eventBehaviors.PlayerFoodLossBehavior(ctx, from, to)
	}
}

// HandleHeal virtual override to the player's heal behavior
func (handler *PlayerHandler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	if handler.eventBehaviors.PlayerHealBehavior != nil {
		handler.eventBehaviors.PlayerHealBehavior(ctx, health, src)
	}
}

// HandleHurt virtual override to the player's hurt behavior
func (handler *PlayerHandler) HandleHurt(ctx *event.Context, damage *float64, duration *time.Duration, source world.DamageSource) {
	if handler.eventBehaviors.PlayerHurtBehavior != nil {
		handler.eventBehaviors.PlayerHurtBehavior(ctx, damage, duration, source)
	}
}

// HandleDeath virtual override to the player's death behavior
func (handler *PlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	if handler.eventBehaviors.PlayerDeathBehavior != nil {
		handler.eventBehaviors.PlayerDeathBehavior(src, keepInv)
	}
}

// HandleRespawn virtual override to the player's respawn behavior
func (handler *PlayerHandler) HandleRespawn(pos *mgl64.Vec3, w **world.World) {
	if handler.eventBehaviors.PlayerRespawnBehavior != nil {
		handler.eventBehaviors.PlayerRespawnBehavior(pos, w)
	}
}

// HandleSkinChange virtual override to the player's skin change behavior
func (handler *PlayerHandler) HandleSkinChange(ctx *event.Context, skin *skin.Skin) {
	if handler.eventBehaviors.PlayerSkinChangeBehavior != nil {
		handler.eventBehaviors.PlayerSkinChangeBehavior(ctx, skin)
	}
}

// HandleStartBreak virtual override to the player's start break behavior
func (handler *PlayerHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	if handler.eventBehaviors.PlayerStartBreakBehavior != nil {
		handler.eventBehaviors.PlayerStartBreakBehavior(ctx, pos)
	}
}

// HandleBlockBreak virtual override to the player's block break behavior
func (handler *PlayerHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	if handler.eventBehaviors.PlayerBlockBreakBehavior != nil {
		handler.eventBehaviors.PlayerBlockBreakBehavior(ctx, pos, drops, xp)
	}
}

// HandleBlockPlace virtual override to the player's block place behavior
func (handler *PlayerHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, b world.Block) {
	if handler.eventBehaviors.PlayerBlockPlaceBehavior != nil {
		handler.eventBehaviors.PlayerBlockPlaceBehavior(ctx, pos, b)
	}
}

// HandleBlockPick virtual override to the player's block pick behavior
func (handler *PlayerHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, b world.Block) {
	if handler.eventBehaviors.PlayerBlockPickBehavior != nil {
		handler.eventBehaviors.PlayerBlockPickBehavior(ctx, pos, b)
	}
}

// HandleItemUse virtual override to the player's item use behavior
func (handler *PlayerHandler) HandleItemUse(ctx *event.Context) {
	if handler.eventBehaviors.PlayerItemUseBehavior != nil {
		handler.eventBehaviors.PlayerItemUseBehavior(ctx)
	}
}

// HandleItemUseOnBlock virtual override to the player's item use on block behavior
func (handler *PlayerHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	if handler.eventBehaviors.PlayerItemUseOnBlockBehavior != nil {
		handler.eventBehaviors.PlayerItemUseOnBlockBehavior(ctx, pos, face, clickPos)
	}
}

// HandleItemUseOnEntity virtual override to the player's item use on entity behavior
func (handler *PlayerHandler) HandleItemUseOnEntity(ctx *event.Context, e world.Entity) {
	if handler.eventBehaviors.PlayerItemUseOnEntityBehavior != nil {
		handler.eventBehaviors.PlayerItemUseOnEntityBehavior(ctx, e)
	}
}

// HandleItemConsume virtual override to the player's item consume behavior
func (handler *PlayerHandler) HandleItemConsume(ctx *event.Context, item item.Stack) {
	if handler.eventBehaviors.PlayerItemConsumeBehavior != nil {
		handler.eventBehaviors.PlayerItemConsumeBehavior(ctx, item)
	}
}

// HandleAttackEntity virtual override to the player's attack entity behavior
func (handler *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force, height *float64, critical *bool) {
	if handler.eventBehaviors.PlayerAttackEntityBehavior != nil {
		handler.eventBehaviors.PlayerAttackEntityBehavior(ctx, e, force, height, critical)
	}
}

// HandleExperienceGain virtual override to the player's experience gain behavior
func (handler *PlayerHandler) HandleExperienceGain(ctx *event.Context, amount *int) {
	if handler.eventBehaviors.PlayerExperienceGainBehavior != nil {
		handler.eventBehaviors.PlayerExperienceGainBehavior(ctx, amount)
	}
}

// HandlePunchAir virtual override to the player's punch air behavior
func (handler *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	if handler.eventBehaviors.PlayerPunchAirBehavior != nil {
		handler.eventBehaviors.PlayerPunchAirBehavior(ctx)
	}
}

// HandleSignEdit virtual override to the player's sign edit behavior
func (handler *PlayerHandler) HandleSignEdit(ctx *event.Context, oldText, newText string) {
	if handler.eventBehaviors.PlayerSignEditBehavior != nil {
		handler.eventBehaviors.PlayerSignEditBehavior(ctx, oldText, newText)
	}
}

// HandleItemDamage virtual override to the player's item damage behavior
func (handler *PlayerHandler) HandleItemDamage(ctx *event.Context, i item.Stack, damage int) {
	if handler.eventBehaviors.PlayerItemDamageBehavior != nil {
		handler.eventBehaviors.PlayerItemDamageBehavior(ctx, i, damage)
	}
}

// HandleItemPickup virtual override to the player's item pickup behavior
func (handler *PlayerHandler) HandleItemPickup(ctx *event.Context, i item.Stack) {
	if handler.eventBehaviors.PlayerItemPickupBehavior != nil {
		handler.eventBehaviors.PlayerItemPickupBehavior(ctx, i)
	}
}

// HandleItemDrop virtual override to the player's item drop behavior
func (handler *PlayerHandler) HandleItemDrop(ctx *event.Context, e *entity.Item) {
	if handler.eventBehaviors.PlayerItemDropBehavior != nil {
		handler.eventBehaviors.PlayerItemDropBehavior(ctx, e)
	}
}

// HandleTransfer virtual override to the player's transfer behavior
func (handler *PlayerHandler) HandleTransfer(ctx *event.Context, addr *net.UDPAddr) {
	if handler.eventBehaviors.PlayerTransferBehavior != nil {
		handler.eventBehaviors.PlayerTransferBehavior(ctx, addr)
	}
}

// HandleCommandExecution virtual override to the player's command execution behavior
func (handler *PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	if handler.eventBehaviors.PlayerCommandExecutionBehavior != nil {
		handler.eventBehaviors.PlayerCommandExecutionBehavior(ctx, command, args)
	}
}

// HandleQuit virtual override to the player's quit behavior
func (handler *PlayerHandler) HandleQuit() {
	if handler.eventBehaviors.PlayerQuitBehavior != nil {
		handler.eventBehaviors.PlayerQuitBehavior(handler.plr)
	}
}
