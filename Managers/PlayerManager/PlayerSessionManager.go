package PlayerManager

import (
	"github.com/df-mc/dragonfly/server/player"
	"sync"
	"time"
)

// PlayerSessionDataManager contains a map of all the players and their session data
type PlayerSessionDataManager struct {
	playersMu sync.RWMutex
	players   map[*player.Player]PlayerData
}

// NewPlayerSessionDataManager creates a new session manager
func NewPlayerSessionDataManager() *PlayerSessionDataManager {
	return &PlayerSessionDataManager{
		players: make(map[*player.Player]PlayerData),
	}
}

// AddPlayer add a player to the session manager
func (session *PlayerSessionDataManager) AddPlayer(p *player.Player) {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	session.players[p] = PlayerData{}
}

// RemovePlayer remove a player from the session manager
func (session *PlayerSessionDataManager) RemovePlayer(p *player.Player) {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	delete(session.players, p)
}

// SessionManagerRoutineStart starts the go routine for updating the player session every second
func SessionManagerRoutineStart(m *PlayerSessionDataManager) {
	go func() {
		ticker := time.NewTicker(1 * time.Second)
		for range ticker.C {
			m.PlayerSessionUpdate()
		}
	}()
}

// PlayerSessionUpdate each update on the session is every second
func (session *PlayerSessionDataManager) PlayerSessionUpdate() {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	// loop through all the players
	for p, data := range session.players {
		// do the scoreboard behavior if it is set
		if data.ScoreboardBehavior != nil {
			data.ScoreboardBehavior(p)
		}
	}
}
