package Managers

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

// The way this works is with Behavior adding and removing functions to each trigger-able player event
// MultiHandler behavior system but for a single handler

// PlayerEventBehavior this struct is maps of functions with named identifiers for each behavior to be overwritten virtually at run time
type PlayerEventBehavior struct {
	PlayerMoveBehaviors             map[string]func(*player.Player, *event.Context, mgl64.Vec3, float64, float64)
	PlayerJumpBehaviors             map[string]func(*player.Player)
	PlayerTeleportBehaviors         map[string]func(*player.Player, *event.Context, mgl64.Vec3)
	PlayerChangeWorldBehaviors      map[string]func(*player.Player, *world.World, *world.World)
	PlayerToggleSprintBehaviors     map[string]func(*player.Player, *event.Context, bool)
	PlayerToggleSneakBehaviors      map[string]func(*player.Player, *event.Context, bool)
	PlayerChatBehaviors             map[string]func(*player.Player, *event.Context, *string)
	PlayerFoodLossBehaviors         map[string]func(*player.Player, *event.Context, int, *int)
	PlayerHealBehaviors             map[string]func(*player.Player, *event.Context, *float64, world.HealingSource)
	PlayerHurtBehaviors             map[string]func(*player.Player, *event.Context, *float64, *time.Duration, world.DamageSource)
	PlayerDeathBehaviors            map[string]func(*player.Player, world.DamageSource, *bool)
	PlayerRespawnBehaviors          map[string]func(*player.Player, *mgl64.Vec3, **world.World)
	PlayerSkinChangeBehaviors       map[string]func(*player.Player, *event.Context, *skin.Skin)
	PlayerStartBreakBehaviors       map[string]func(*player.Player, *event.Context, cube.Pos)
	PlayerBlockBreakBehaviors       map[string]func(*player.Player, *event.Context, cube.Pos, *[]item.Stack, *int)
	PlayerBlockPlaceBehaviors       map[string]func(*player.Player, *event.Context, cube.Pos, world.Block)
	PlayerBlockPickBehaviors        map[string]func(*player.Player, *event.Context, cube.Pos, world.Block)
	PlayerItemUseBehaviors          map[string]func(*player.Player, *event.Context)
	PlayerItemUseOnBlockBehaviors   map[string]func(*player.Player, *event.Context, cube.Pos, cube.Face, mgl64.Vec3)
	PlayerItemUseOnEntityBehaviors  map[string]func(*player.Player, *event.Context, world.Entity)
	PlayerItemConsumeBehaviors      map[string]func(*player.Player, *event.Context, item.Stack)
	PlayerAttackEntityBehaviors     map[string]func(*player.Player, *event.Context, world.Entity, *float64, *float64, *bool)
	PlayerExperienceGainBehaviors   map[string]func(*player.Player, *event.Context, *int)
	PlayerPunchAirBehaviors         map[string]func(*player.Player, *event.Context)
	PlayerSignEditBehaviors         map[string]func(*player.Player, *event.Context, string, string)
	PlayerItemDamageBehaviors       map[string]func(*player.Player, *event.Context, item.Stack, int)
	PlayerItemPickupBehaviors       map[string]func(*player.Player, *event.Context, item.Stack)
	PlayerItemDropBehaviors         map[string]func(*player.Player, *event.Context, *entity.Item)
	PlayerTransferBehaviors         map[string]func(*player.Player, *event.Context, *net.UDPAddr)
	PlayerCommandExecutionBehaviors map[string]func(*player.Player, *event.Context, cmd.Command, []string)
	PlayerQuitBehaviors             map[string]func(*player.Player)
}

// PlayerHandler this struct defines the handler to have optional overwrites for the event interfaces with our behaviors
type PlayerHandler struct {
	player.NopHandler
	plr            *player.Player
	eventBehaviors *PlayerEventBehavior
}

// NewPlayerHandler returns a new PlayerHandler for our event behaviors
func NewPlayerHandler(plr *player.Player) *PlayerHandler {
	return &PlayerHandler{
		plr:            plr,
		eventBehaviors: newPlayerEventBehavior(),
	}
}

func newPlayerEventBehavior() *PlayerEventBehavior {
	return &PlayerEventBehavior{
		PlayerMoveBehaviors:             make(map[string]func(*player.Player, *event.Context, mgl64.Vec3, float64, float64)),
		PlayerJumpBehaviors:             make(map[string]func(*player.Player)),
		PlayerTeleportBehaviors:         make(map[string]func(*player.Player, *event.Context, mgl64.Vec3)),
		PlayerChangeWorldBehaviors:      make(map[string]func(*player.Player, *world.World, *world.World)),
		PlayerToggleSprintBehaviors:     make(map[string]func(*player.Player, *event.Context, bool)),
		PlayerToggleSneakBehaviors:      make(map[string]func(*player.Player, *event.Context, bool)),
		PlayerChatBehaviors:             make(map[string]func(*player.Player, *event.Context, *string)),
		PlayerFoodLossBehaviors:         make(map[string]func(*player.Player, *event.Context, int, *int)),
		PlayerHealBehaviors:             make(map[string]func(*player.Player, *event.Context, *float64, world.HealingSource)),
		PlayerHurtBehaviors:             make(map[string]func(*player.Player, *event.Context, *float64, *time.Duration, world.DamageSource)),
		PlayerDeathBehaviors:            make(map[string]func(*player.Player, world.DamageSource, *bool)),
		PlayerRespawnBehaviors:          make(map[string]func(*player.Player, *mgl64.Vec3, **world.World)),
		PlayerSkinChangeBehaviors:       make(map[string]func(*player.Player, *event.Context, *skin.Skin)),
		PlayerStartBreakBehaviors:       make(map[string]func(*player.Player, *event.Context, cube.Pos)),
		PlayerBlockBreakBehaviors:       make(map[string]func(*player.Player, *event.Context, cube.Pos, *[]item.Stack, *int)),
		PlayerBlockPlaceBehaviors:       make(map[string]func(*player.Player, *event.Context, cube.Pos, world.Block)),
		PlayerBlockPickBehaviors:        make(map[string]func(*player.Player, *event.Context, cube.Pos, world.Block)),
		PlayerItemUseBehaviors:          make(map[string]func(*player.Player, *event.Context)),
		PlayerItemUseOnBlockBehaviors:   make(map[string]func(*player.Player, *event.Context, cube.Pos, cube.Face, mgl64.Vec3)),
		PlayerItemUseOnEntityBehaviors:  make(map[string]func(*player.Player, *event.Context, world.Entity)),
		PlayerItemConsumeBehaviors:      make(map[string]func(*player.Player, *event.Context, item.Stack)),
		PlayerAttackEntityBehaviors:     make(map[string]func(*player.Player, *event.Context, world.Entity, *float64, *float64, *bool)),
		PlayerExperienceGainBehaviors:   make(map[string]func(*player.Player, *event.Context, *int)),
		PlayerPunchAirBehaviors:         make(map[string]func(*player.Player, *event.Context)),
		PlayerSignEditBehaviors:         make(map[string]func(*player.Player, *event.Context, string, string)),
		PlayerItemDamageBehaviors:       make(map[string]func(*player.Player, *event.Context, item.Stack, int)),
		PlayerItemPickupBehaviors:       make(map[string]func(*player.Player, *event.Context, item.Stack)),
		PlayerItemDropBehaviors:         make(map[string]func(*player.Player, *event.Context, *entity.Item)),
		PlayerTransferBehaviors:         make(map[string]func(*player.Player, *event.Context, *net.UDPAddr)),
		PlayerCommandExecutionBehaviors: make(map[string]func(*player.Player, *event.Context, cmd.Command, []string)),
		PlayerQuitBehaviors:             make(map[string]func(*player.Player)),
	}
}

// ClearPlayerBehaviors sets all event behavior functions in the PlayerHandler's PlayerEventBehavior to nil
func (handler *PlayerHandler) ClearPlayerBehaviors() {
	handler.eventBehaviors = newPlayerEventBehavior()
}

/*@@@@@@@@@@@@@@@@@@@@@@@@@
     BEHAVIOR SETTING!!!
@@@@@@@@@@@@@@@@@@@@@@@@@@@*/

// AddMoveBehavior adds a move behavior to the player with a named identifier string
func (handler *PlayerHandler) AddMoveBehavior(identifier string, behavior func(*player.Player, *event.Context, mgl64.Vec3, float64, float64)) {
	handler.eventBehaviors.PlayerMoveBehaviors[identifier] = behavior
}

// RemoveMoveBehavior removes a move behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveMoveBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerMoveBehaviors, identifier)
}

// AddJumpBehavior adds a jump behavior to the player with a named identifier string
func (handler *PlayerHandler) AddJumpBehavior(identifier string, behavior func(*player.Player)) {
	handler.eventBehaviors.PlayerJumpBehaviors[identifier] = behavior
}

// RemoveJumpBehavior removes a jump behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveJumpBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerJumpBehaviors, identifier)
}

// AddTeleportBehavior adds a teleport behavior to the player with a named identifier string
func (handler *PlayerHandler) AddTeleportBehavior(identifier string, behavior func(*player.Player, *event.Context, mgl64.Vec3)) {
	handler.eventBehaviors.PlayerTeleportBehaviors[identifier] = behavior
}

// RemoveTeleportBehavior removes a teleport behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveTeleportBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerTeleportBehaviors, identifier)
}

// AddChangeWorldBehavior adds a change world behavior to the player with a named identifier string
func (handler *PlayerHandler) AddChangeWorldBehavior(identifier string, behavior func(*player.Player, *world.World, *world.World)) {
	handler.eventBehaviors.PlayerChangeWorldBehaviors[identifier] = behavior
}

// RemoveChangeWorldBehavior removes a change world behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveChangeWorldBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerChangeWorldBehaviors, identifier)
}

// AddToggleSprintBehavior adds a toggle sprint behavior to the player with a named identifier string
func (handler *PlayerHandler) AddToggleSprintBehavior(identifier string, behavior func(*player.Player, *event.Context, bool)) {
	handler.eventBehaviors.PlayerToggleSprintBehaviors[identifier] = behavior
}

// RemoveToggleSprintBehavior removes a toggle sprint behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveToggleSprintBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerToggleSprintBehaviors, identifier)
}

// AddToggleSneakBehavior adds a toggle sneak behavior to the player with a named identifier string
func (handler *PlayerHandler) AddToggleSneakBehavior(identifier string, behavior func(*player.Player, *event.Context, bool)) {
	handler.eventBehaviors.PlayerToggleSneakBehaviors[identifier] = behavior
}

// RemoveToggleSneakBehavior removes a toggle sneak behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveToggleSneakBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerToggleSneakBehaviors, identifier)
}

// AddChatBehavior adds a chat behavior to the player with a named identifier string
func (handler *PlayerHandler) AddChatBehavior(identifier string, behavior func(*player.Player, *event.Context, *string)) {
	handler.eventBehaviors.PlayerChatBehaviors[identifier] = behavior
}

// RemoveChatBehavior removes a chat behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveChatBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerChatBehaviors, identifier)
}

// AddFoodLossBehavior adds a food loss behavior to the player with a named identifier string
func (handler *PlayerHandler) AddFoodLossBehavior(identifier string, behavior func(*player.Player, *event.Context, int, *int)) {
	handler.eventBehaviors.PlayerFoodLossBehaviors[identifier] = behavior
}

// RemoveFoodLossBehavior removes a food loss behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveFoodLossBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerFoodLossBehaviors, identifier)
}

// AddHealBehavior adds a heal behavior to the player with a named identifier string
func (handler *PlayerHandler) AddHealBehavior(identifier string, behavior func(*player.Player, *event.Context, *float64, world.HealingSource)) {
	handler.eventBehaviors.PlayerHealBehaviors[identifier] = behavior
}

// RemoveHealBehavior removes a heal behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveHealBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerHealBehaviors, identifier)
}

// AddHurtBehavior adds a hurt behavior to the player with a named identifier string
func (handler *PlayerHandler) AddHurtBehavior(identifier string, behavior func(*player.Player, *event.Context, *float64, *time.Duration, world.DamageSource)) {
	handler.eventBehaviors.PlayerHurtBehaviors[identifier] = behavior
}

// RemoveHurtBehavior removes a hurt behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveHurtBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerHurtBehaviors, identifier)
}

// AddDeathBehavior adds a death behavior to the player with a named identifier string
func (handler *PlayerHandler) AddDeathBehavior(identifier string, behavior func(*player.Player, world.DamageSource, *bool)) {
	handler.eventBehaviors.PlayerDeathBehaviors[identifier] = behavior
}

// RemoveDeathBehavior removes a death behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveDeathBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerDeathBehaviors, identifier)
}

// AddRespawnBehavior adds a respawn behavior to the player with a named identifier string
func (handler *PlayerHandler) AddRespawnBehavior(identifier string, behavior func(*player.Player, *mgl64.Vec3, **world.World)) {
	handler.eventBehaviors.PlayerRespawnBehaviors[identifier] = behavior
}

// RemoveRespawnBehavior removes a respawn behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveRespawnBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerRespawnBehaviors, identifier)
}

// AddSkinChangeBehavior adds a skin change behavior to the player with a named identifier string
func (handler *PlayerHandler) AddSkinChangeBehavior(identifier string, behavior func(*player.Player, *event.Context, *skin.Skin)) {
	handler.eventBehaviors.PlayerSkinChangeBehaviors[identifier] = behavior
}

// RemoveSkinChangeBehavior removes a skin change behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveSkinChangeBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerSkinChangeBehaviors, identifier)
}

// AddStartBreakBehavior adds a start break behavior to the player with a named identifier string
func (handler *PlayerHandler) AddStartBreakBehavior(identifier string, behavior func(*player.Player, *event.Context, cube.Pos)) {
	handler.eventBehaviors.PlayerStartBreakBehaviors[identifier] = behavior
}

// RemoveStartBreakBehavior removes a start break behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveStartBreakBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerStartBreakBehaviors, identifier)
}

// AddBlockBreakBehavior adds a block break behavior to the player with a named identifier string
func (handler *PlayerHandler) AddBlockBreakBehavior(identifier string, behavior func(*player.Player, *event.Context, cube.Pos, *[]item.Stack, *int)) {
	handler.eventBehaviors.PlayerBlockBreakBehaviors[identifier] = behavior
}

// RemoveBlockBreakBehavior removes a block break behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveBlockBreakBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerBlockBreakBehaviors, identifier)
}

// AddBlockPlaceBehavior adds a block place behavior to the player with a named identifier string
func (handler *PlayerHandler) AddBlockPlaceBehavior(identifier string, behavior func(*player.Player, *event.Context, cube.Pos, world.Block)) {
	handler.eventBehaviors.PlayerBlockPlaceBehaviors[identifier] = behavior
}

// RemoveBlockPlaceBehavior removes a block place behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveBlockPlaceBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerBlockPlaceBehaviors, identifier)
}

// AddBlockPickBehavior adds a block pick behavior to the player with a named identifier string
func (handler *PlayerHandler) AddBlockPickBehavior(identifier string, behavior func(*player.Player, *event.Context, cube.Pos, world.Block)) {
	handler.eventBehaviors.PlayerBlockPickBehaviors[identifier] = behavior
}

// RemoveBlockPickBehavior removes a block pick behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveBlockPickBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerBlockPickBehaviors, identifier)
}

// AddItemUseBehavior adds an item use behavior to the player with a named identifier string
func (handler *PlayerHandler) AddItemUseBehavior(identifier string, behavior func(*player.Player, *event.Context)) {
	handler.eventBehaviors.PlayerItemUseBehaviors[identifier] = behavior
}

// RemoveItemUseBehavior removes an item use behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveItemUseBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemUseBehaviors, identifier)
}

// AddItemUseOnBlockBehavior adds an item use on block behavior to the player with a named identifier string
func (handler *PlayerHandler) AddItemUseOnBlockBehavior(identifier string, behavior func(*player.Player, *event.Context, cube.Pos, cube.Face, mgl64.Vec3)) {
	handler.eventBehaviors.PlayerItemUseOnBlockBehaviors[identifier] = behavior
}

// RemoveItemUseOnBlockBehavior removes an item use on block behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveItemUseOnBlockBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemUseOnBlockBehaviors, identifier)
}

// AddItemUseOnEntityBehavior adds an item use on entity behavior to the player with a named identifier string
func (handler *PlayerHandler) AddItemUseOnEntityBehavior(identifier string, behavior func(*player.Player, *event.Context, world.Entity)) {
	handler.eventBehaviors.PlayerItemUseOnEntityBehaviors[identifier] = behavior
}

// RemoveItemUseOnEntityBehavior removes an item use on entity behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveItemUseOnEntityBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemUseOnEntityBehaviors, identifier)
}

// AddPlayerItemConsumeBehavior adds an item consume behavior to the player with a named identifier string
func (handler *PlayerHandler) AddPlayerItemConsumeBehavior(identifier string, behavior func(*player.Player, *event.Context, item.Stack)) {
	handler.eventBehaviors.PlayerItemConsumeBehaviors[identifier] = behavior
}

// RemovePlayerItemConsumeBehavior removes an item consume behavior from the player with a named identifier string
func (handler *PlayerHandler) RemovePlayerItemConsumeBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemConsumeBehaviors, identifier)
}

// AddPlayerAttackEntityBehavior adds an attack entity behavior to the player with a named identifier string
func (handler *PlayerHandler) AddPlayerAttackEntityBehavior(identifier string, behavior func(*player.Player, *event.Context, world.Entity, *float64, *float64, *bool)) {
	handler.eventBehaviors.PlayerAttackEntityBehaviors[identifier] = behavior
}

// RemovePlayerAttackEntityBehavior removes an attack entity behavior from the player with a named identifier string
func (handler *PlayerHandler) RemovePlayerAttackEntityBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerAttackEntityBehaviors, identifier)
}

// AddPlayerExperienceGainBehavior adds an experience gain behavior to the player with a named identifier string
func (handler *PlayerHandler) AddPlayerExperienceGainBehavior(identifier string, behavior func(*player.Player, *event.Context, *int)) {
	handler.eventBehaviors.PlayerExperienceGainBehaviors[identifier] = behavior
}

// RemovePlayerExperienceGainBehavior removes an experience gain behavior from the player with a named identifier string
func (handler *PlayerHandler) RemovePlayerExperienceGainBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerExperienceGainBehaviors, identifier)
}

// AddPlayerPunchAirBehavior adds a punch air behavior to the player with a named identifier string
func (handler *PlayerHandler) AddPlayerPunchAirBehavior(identifier string, behavior func(*player.Player, *event.Context)) {
	handler.eventBehaviors.PlayerPunchAirBehaviors[identifier] = behavior
}

// RemovePlayerPunchAirBehavior removes a punch air behavior from the player with a named identifier string
func (handler *PlayerHandler) RemovePlayerPunchAirBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerPunchAirBehaviors, identifier)
}

// AddPlayerSignEditBehavior adds a sign edit behavior to the player with a named identifier string
func (handler *PlayerHandler) AddPlayerSignEditBehavior(identifier string, behavior func(*player.Player, *event.Context, string, string)) {
	handler.eventBehaviors.PlayerSignEditBehaviors[identifier] = behavior
}

// RemovePlayerSignEditBehavior removes a sign edit behavior from the player with a named identifier string
func (handler *PlayerHandler) RemovePlayerSignEditBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerSignEditBehaviors, identifier)
}

// AddItemPickupBehavior adds an item pickup behavior to the player with a named identifier string
func (handler *PlayerHandler) AddItemPickupBehavior(identifier string, behavior func(*player.Player, *event.Context, item.Stack)) {
	handler.eventBehaviors.PlayerItemPickupBehaviors[identifier] = behavior
}

// RemoveItemPickupBehavior removes an item pickup behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveItemPickupBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemPickupBehaviors, identifier)
}

// AddItemDropBehavior adds an item drop behavior to the player with a named identifier string
func (handler *PlayerHandler) AddItemDropBehavior(identifier string, behavior func(*player.Player, *event.Context, *entity.Item)) {
	handler.eventBehaviors.PlayerItemDropBehaviors[identifier] = behavior
}

// RemoveItemDropBehavior removes an item drop behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveItemDropBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerItemDropBehaviors, identifier)
}

// AddTransferBehavior adds a transfer behavior to the player with a named identifier string
func (handler *PlayerHandler) AddTransferBehavior(identifier string, behavior func(*player.Player, *event.Context, *net.UDPAddr)) {
	handler.eventBehaviors.PlayerTransferBehaviors[identifier] = behavior
}

// RemoveTransferBehavior removes a transfer behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveTransferBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerTransferBehaviors, identifier)
}

// AddCommandExecutionBehavior adds a command execution behavior to the player with a named identifier string
func (handler *PlayerHandler) AddCommandExecutionBehavior(identifier string, behavior func(*player.Player, *event.Context, cmd.Command, []string)) {
	handler.eventBehaviors.PlayerCommandExecutionBehaviors[identifier] = behavior
}

// RemoveCommandExecutionBehavior removes a command execution behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveCommandExecutionBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerCommandExecutionBehaviors, identifier)
}

// AddQuitBehavior adds a quit behavior to the player with a named identifier string
func (handler *PlayerHandler) AddQuitBehavior(identifier string, behavior func(*player.Player)) {
	handler.eventBehaviors.PlayerQuitBehaviors[identifier] = behavior
}

// RemoveQuitBehavior removes a quit behavior from the player with a named identifier string
func (handler *PlayerHandler) RemoveQuitBehavior(identifier string) {
	delete(handler.eventBehaviors.PlayerQuitBehaviors, identifier)
}

/*@@@@@@@@@@@@@@@@@@@@@@@@@
  CALLBACK BOILER PLATING!!
@@@@@@@@@@@@@@@@@@@@@@@@@@@*/

// HandleMove virtual override to the player's move behavior
func (handler *PlayerHandler) HandleMove(ctx *event.Context, newPos mgl64.Vec3, newYaw, newPitch float64) {
	for _, behavior := range handler.eventBehaviors.PlayerMoveBehaviors {
		behavior(handler.plr, ctx, newPos, newYaw, newPitch)
	}
}

// HandleJump virtual override to the player's jump behavior
func (handler *PlayerHandler) HandleJump() {
	for _, behavior := range handler.eventBehaviors.PlayerJumpBehaviors {
		behavior(handler.plr)
	}
}

// HandleTeleport virtual override to the player's teleport behavior
func (handler *PlayerHandler) HandleTeleport(ctx *event.Context, newPos mgl64.Vec3) {
	for _, behavior := range handler.eventBehaviors.PlayerTeleportBehaviors {
		behavior(handler.plr, ctx, newPos)
	}
}

// HandleChangeWorld virtual override to the player's change world behavior
func (handler *PlayerHandler) HandleChangeWorld(oldWorld, newWorld *world.World) {
	for _, behavior := range handler.eventBehaviors.PlayerChangeWorldBehaviors {
		behavior(handler.plr, oldWorld, newWorld)
	}
}

// HandleToggleSprint virtual override to the player's toggle sprint behavior
func (handler *PlayerHandler) HandleToggleSprint(ctx *event.Context, sprinting bool) {
	for _, behavior := range handler.eventBehaviors.PlayerToggleSprintBehaviors {
		behavior(handler.plr, ctx, sprinting)
	}
}

// HandleToggleSneak virtual override to the player's toggle sneak behavior
func (handler *PlayerHandler) HandleToggleSneak(ctx *event.Context, sneaking bool) {
	for _, behavior := range handler.eventBehaviors.PlayerToggleSneakBehaviors {
		behavior(handler.plr, ctx, sneaking)
	}
}

// HandleChat virtual override to the player's chat behavior
func (handler *PlayerHandler) HandleChat(ctx *event.Context, message *string) {
	for _, behavior := range handler.eventBehaviors.PlayerChatBehaviors {
		behavior(handler.plr, ctx, message)
	}
}

// HandleFoodLoss virtual override to the player's food loss behavior
func (handler *PlayerHandler) HandleFoodLoss(ctx *event.Context, from int, to *int) {
	for _, behavior := range handler.eventBehaviors.PlayerFoodLossBehaviors {
		behavior(handler.plr, ctx, from, to)
	}
}

// HandleHeal virtual override to the player's heal behavior
func (handler *PlayerHandler) HandleHeal(ctx *event.Context, health *float64, src world.HealingSource) {
	for _, behavior := range handler.eventBehaviors.PlayerHealBehaviors {
		behavior(handler.plr, ctx, health, src)
	}
}

// HandleHurt virtual override to the player's hurt behavior
func (handler *PlayerHandler) HandleHurt(ctx *event.Context, damage *float64, attackImmunity *time.Duration, src world.DamageSource) {
	for _, behavior := range handler.eventBehaviors.PlayerHurtBehaviors {
		behavior(handler.plr, ctx, damage, attackImmunity, src)
	}
}

// HandleDeath virtual override to the player's death behavior
func (handler *PlayerHandler) HandleDeath(src world.DamageSource, keepInv *bool) {
	for _, behavior := range handler.eventBehaviors.PlayerDeathBehaviors {
		behavior(handler.plr, src, keepInv)
	}
}

// HandleRespawn virtual override to the player's respawn behavior
func (handler *PlayerHandler) HandleRespawn(pos *mgl64.Vec3, world **world.World) {
	for _, behavior := range handler.eventBehaviors.PlayerRespawnBehaviors {
		behavior(handler.plr, pos, world)
	}
}

// HandleSkinChange virtual override to the player's skin change behavior
func (handler *PlayerHandler) HandleSkinChange(ctx *event.Context, newSkin *skin.Skin) {
	for _, behavior := range handler.eventBehaviors.PlayerSkinChangeBehaviors {
		behavior(handler.plr, ctx, newSkin)
	}
}

// HandleStartBreak virtual override to the player's start break behavior
func (handler *PlayerHandler) HandleStartBreak(ctx *event.Context, pos cube.Pos) {
	for _, behavior := range handler.eventBehaviors.PlayerStartBreakBehaviors {
		behavior(handler.plr, ctx, pos)
	}
}

// HandleBlockBreak virtual override to the player's block break behavior
func (handler *PlayerHandler) HandleBlockBreak(ctx *event.Context, pos cube.Pos, drops *[]item.Stack, xp *int) {
	for _, behavior := range handler.eventBehaviors.PlayerBlockBreakBehaviors {
		behavior(handler.plr, ctx, pos, drops, xp)
	}
}

// HandleBlockPlace virtual override to the player's block place behavior
func (handler *PlayerHandler) HandleBlockPlace(ctx *event.Context, pos cube.Pos, block world.Block) {
	for _, behavior := range handler.eventBehaviors.PlayerBlockPlaceBehaviors {
		behavior(handler.plr, ctx, pos, block)
	}
}

// HandleBlockPick virtual override to the player's block pick behavior
func (handler *PlayerHandler) HandleBlockPick(ctx *event.Context, pos cube.Pos, block world.Block) {
	for _, behavior := range handler.eventBehaviors.PlayerBlockPickBehaviors {
		behavior(handler.plr, ctx, pos, block)
	}
}

// HandleItemUse virtual override to the player's item use behavior
func (handler *PlayerHandler) HandleItemUse(ctx *event.Context) {
	for _, behavior := range handler.eventBehaviors.PlayerItemUseBehaviors {
		behavior(handler.plr, ctx)
	}
}

// HandleItemUseOnBlock virtual override to the player's item use on block behavior
func (handler *PlayerHandler) HandleItemUseOnBlock(ctx *event.Context, pos cube.Pos, face cube.Face, clickPos mgl64.Vec3) {
	for _, behavior := range handler.eventBehaviors.PlayerItemUseOnBlockBehaviors {
		behavior(handler.plr, ctx, pos, face, clickPos)
	}
}

// HandleItemUseOnEntity virtual override to the player's item use on entity behavior
func (handler *PlayerHandler) HandleItemUseOnEntity(ctx *event.Context, entity world.Entity) {
	for _, behavior := range handler.eventBehaviors.PlayerItemUseOnEntityBehaviors {
		behavior(handler.plr, ctx, entity)
	}
}

// HandleItemConsume virtual override to the player's item consume behavior
func (handler *PlayerHandler) HandleItemConsume(ctx *event.Context, stack item.Stack) {
	for _, behavior := range handler.eventBehaviors.PlayerItemConsumeBehaviors {
		behavior(handler.plr, ctx, stack)
	}
}

// HandleAttackEntity virtual override to the player's attack entity behavior
func (handler *PlayerHandler) HandleAttackEntity(ctx *event.Context, e world.Entity, force *float64, height *float64, critical *bool) {
	for _, behavior := range handler.eventBehaviors.PlayerAttackEntityBehaviors {
		behavior(handler.plr, ctx, e, force, height, critical)
	}
}

// HandleExperienceGain virtual override to the player's experience gain behavior
func (handler *PlayerHandler) HandleExperienceGain(ctx *event.Context, amount *int) {
	for _, behavior := range handler.eventBehaviors.PlayerExperienceGainBehaviors {
		behavior(handler.plr, ctx, amount)
	}
}

// HandlePunchAir virtual override to the player's punch air behavior
func (handler *PlayerHandler) HandlePunchAir(ctx *event.Context) {
	for _, behavior := range handler.eventBehaviors.PlayerPunchAirBehaviors {
		behavior(handler.plr, ctx)
	}
}

// HandleSignEdit virtual override to the player's sign edit behavior
func (handler *PlayerHandler) HandleSignEdit(ctx *event.Context, line1, line2 string) {
	for _, behavior := range handler.eventBehaviors.PlayerSignEditBehaviors {
		behavior(handler.plr, ctx, line1, line2)
	}
}

// HandleItemDamage virtual override to the player's item damage behavior
func (handler *PlayerHandler) HandleItemDamage(ctx *event.Context, stack item.Stack, damage int) {
	for _, behavior := range handler.eventBehaviors.PlayerItemDamageBehaviors {
		behavior(handler.plr, ctx, stack, damage)
	}
}

// HandleItemPickup virtual override to the player's item pickup behavior
func (handler *PlayerHandler) HandleItemPickup(ctx *event.Context, stack item.Stack) {
	for _, behavior := range handler.eventBehaviors.PlayerItemPickupBehaviors {
		behavior(handler.plr, ctx, stack)
	}
}

// HandleItemDrop virtual override to the player's item drop behavior
func (handler *PlayerHandler) HandleItemDrop(ctx *event.Context, droppedItem *entity.Item) {
	for _, behavior := range handler.eventBehaviors.PlayerItemDropBehaviors {
		behavior(handler.plr, ctx, droppedItem)
	}
}

// HandleTransfer virtual override to the player's transfer behavior
func (handler *PlayerHandler) HandleTransfer(ctx *event.Context, address *net.UDPAddr) {
	for _, behavior := range handler.eventBehaviors.PlayerTransferBehaviors {
		behavior(handler.plr, ctx, address)
	}
}

// HandleCommandExecution virtual override to the player's command execution behavior
func (handler *PlayerHandler) HandleCommandExecution(ctx *event.Context, command cmd.Command, args []string) {
	for _, behavior := range handler.eventBehaviors.PlayerCommandExecutionBehaviors {
		behavior(handler.plr, ctx, command, args)
	}
}

// HandleQuit virtual override to the player's quit behavior
func (handler *PlayerHandler) HandleQuit() {
	for _, behavior := range handler.eventBehaviors.PlayerQuitBehaviors {
		behavior(handler.plr)
	}
}
