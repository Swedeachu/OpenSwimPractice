package PlayerManager

import "github.com/df-mc/dragonfly/server/player"

// PlayerData holds all the session data and behaviors we care about, such as scoreboard behavior callback and ranks
type PlayerData struct {
	ScoreboardBehavior func(*player.Player)
}

// SetPlayerScoreboardBehavior sets the scoreboard behavior function for the given player
func (session *PlayerSessionDataManager) SetPlayerScoreboardBehavior(p *player.Player, behavior func(*player.Player)) {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	// only works if the player is in the session manager
	if data, ok := session.players[p]; ok {
		data.ScoreboardBehavior = behavior
		session.players[p] = data
	}
}
