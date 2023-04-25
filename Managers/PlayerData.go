package Managers

import (
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/player/scoreboard"
	"sync"
)

// PlayerData holds all the session data and behaviors we care about, such as specific behavior callbacks and objects
type PlayerData struct {
	ScoreboardBehavior func(*player.Player, *scoreboard.Scoreboard)
	ScoreBoard         *scoreboard.Scoreboard
	ScoreTagBehavior   func(*player.Player)
	ClicksArray        *[]int64
	PearlCoolDown      float64
	Handler            *PlayerHandler
	Mutex              sync.RWMutex
}

func (data *PlayerData) PlayerDataLock() {
	data.Mutex.Lock()
}

func (data *PlayerData) PlayerDataUnlock() {
	data.Mutex.Unlock()
}

func (data *PlayerData) PlayerDataRLock() {
	data.Mutex.RLock()
}

func (data *PlayerData) PlayerDataRUnlock() {
	data.Mutex.RUnlock()
}

// GetScoreBoard gets the player's scoreboard attached to their session.
// PlayerSessionManager's update function should only need to access this function
func (session *PlayerSessionDataManager) GetScoreBoard(p *player.Player) *scoreboard.Scoreboard {
	session.Players[p].PlayerDataRLock()
	defer session.Players[p].PlayerDataRUnlock()
	return session.Players[p].ScoreBoard
}

// SetPlayerScoreboardBehavior sets the scoreboard behavior function for the given player
// the Scoreboard behavior function must have pointers to player and scoreboard as arguments
func (session *PlayerSessionDataManager) SetPlayerScoreboardBehavior(p *player.Player, behavior func(*player.Player, *scoreboard.Scoreboard)) {
	session.Players[p].PlayerDataLock()
	defer session.Players[p].PlayerDataUnlock()
	// only works if the player is in the session manager
	if data, ok := session.Players[p]; ok {
		// get rid of the old scoreboard just in case if there is one
		p.RemoveScoreboard()
		data.ScoreBoard = scoreboard.New("§bSwim Practice") // assign a new one the behavior will edit
		data.ScoreboardBehavior = behavior
		session.Players[p] = data
	}
}

// SetPlayerScoreTagBehavior sets the score tag behavior function for the given player
func (session *PlayerSessionDataManager) SetPlayerScoreTagBehavior(p *player.Player, behavior func(*player.Player)) {
	session.Players[p].PlayerDataLock()
	defer session.Players[p].PlayerDataUnlock()
	// only works if the player is in the session manager
	if data, ok := session.Players[p]; ok {
		data.ScoreTagBehavior = behavior
		session.Players[p] = data
	}
}

// SetPlayerHandler sets the player Handler for the given player in the PlayerData struct
// this exists because we will need to get the player's Handler, so we can change there behaviors
func (session *PlayerSessionDataManager) SetPlayerHandler(p *player.Player, handler *PlayerHandler) {
	session.Players[p].PlayerDataLock()
	defer session.Players[p].PlayerDataUnlock()
	// only works if the player is in the session manager
	if data, ok := session.Players[p]; ok {
		data.Handler = handler
		session.Players[p] = data
	}
}

// GetPlayerHandler gets the player Handler for the given player in the PlayerData struct
// this exists because we will need to get the player's Handler, so we can change there behaviors
func (session *PlayerSessionDataManager) GetPlayerHandler(p *player.Player) *PlayerHandler {
	session.Players[p].PlayerDataRLock()
	defer session.Players[p].PlayerDataRUnlock()
	return session.Players[p].Handler
}
