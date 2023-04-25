package Managers

import (
	"github.com/df-mc/dragonfly/server/player"
	"sync"
	"time"
)

// PlayerSessionDataManager contains a map of all the Players and their session data
type PlayerSessionDataManager struct {
	playersMu sync.RWMutex
	Players   map[*player.Player]*PlayerData
}

// NewPlayerSessionDataManager creates a new session manager
func NewPlayerSessionDataManager() *PlayerSessionDataManager {
	return &PlayerSessionDataManager{
		Players: make(map[*player.Player]*PlayerData),
	}
}

// AddPlayer add a player to the session manager
func (session *PlayerSessionDataManager) AddPlayer(p *player.Player) {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	session.Players[p] = &PlayerData{
		ClicksArray:   new([]int64), // Initialize the ClicksArray field with an empty slice.
		PearlCoolDown: 0,
	}
}

// RemovePlayer remove a player from the session manager
func (session *PlayerSessionDataManager) RemovePlayer(p *player.Player) {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	delete(session.Players, p)
}

// SessionManagerRoutineStart starts the go routine for updating the player session every tick and every second
func SessionManagerRoutineStart(session *PlayerSessionDataManager) {
	go func() {
		tickUpdater := time.NewTicker(50 * time.Millisecond) // runs the functions that need to update every tick
		secondUpdater := time.NewTicker(1 * time.Second)     // runs the functions that need to update every second
		for {
			select {
			case <-tickUpdater.C:
				session.playerSessionTickUpdate()
			case <-secondUpdater.C:
				session.playerSessionSecondUpdate()
			}
		}
	}()
}

// playerSessionTickUpdate runs the tick update functions on the players
func (session *PlayerSessionDataManager) playerSessionTickUpdate() {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	// the functions ran are what need to be precise and rely on player data every server cycle (tick)
	for p := range session.Players {
		UpdateCPS(p)   // update the players cps
		UpdatePearl(p) // updates the players pearl cool down
	}
}

// playerSessionSecondUpdate runs the second update functions on the players
func (session *PlayerSessionDataManager) playerSessionSecondUpdate() {
	session.playersMu.Lock()
	defer session.playersMu.Unlock()
	// the functions ran here are fine to be activated every second
	for p, data := range session.Players {
		// run any active scoreboard behavior
		if data.ScoreboardBehavior != nil {
			data.ScoreboardBehavior(p, session.GetScoreBoard(p))
		}
		// run any active score tag behavior
		if data.ScoreTagBehavior != nil {
			data.ScoreTagBehavior(p)
		}
	}
}
