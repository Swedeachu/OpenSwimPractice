package Managers

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/form"
)

type ffaForm struct{}

func (f ffaForm) Submit(submitter form.Submitter, pressed form.Button) {
	plr := submitter.(*player.Player)
	if pressed.Text == "§4Pot FFA" {
		// teleport player to a random place in the pot ffa world
		pos := VectorRandOffset3D(700, 67, 400, 100, 100)
		ClearPlayer(plr)
		TeleportToWorld(plr, "PotFFA", pos)
		// assign and remove old and new behaviors for pot ffa onto the players behavior tree
		handler := Swim.SessionDataManager.Players[plr].Handler
		handler.RemoveHurtBehavior("hub")
		handler.AddHurtBehavior("ffa", combatHurtBehavior)
		handler.AddPlayerAttackEntityBehavior("ffa", combatAttackBehavior)
		// give the pot kit
		PotKit(plr)
	}
}

func FormFFA() form.Menu {
	var buttons []form.Button
	buttons = append(buttons, form.NewButton("§4Pot FFA", "textures/items/potion_bottle_heal"))
	return form.NewMenu(ffaForm{}, "Free For All Modes").WithButtons(buttons...)
}
