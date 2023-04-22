package Managers

// 90% of this is code written by DidntPot

import (
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/entity"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/df-mc/dragonfly/server/world/mcdb"
	"github.com/df-mc/goleveldb/leveldb/opt"
	"github.com/sirupsen/logrus"
	"os"
	"sync"
)

// WorldManager The world struct for where it exists on disk and its place in the world map
type WorldManager struct {
	s *server.Server

	folderPath string

	log *logrus.Logger

	worldsMu sync.RWMutex
	worlds   map[string]*world.World
}

// NewWorldManager Create a new world manager
func NewWorldManager(server *server.Server, folderPath string, log *logrus.Logger) *WorldManager {
	_ = os.Mkdir(folderPath, 0644)
	defaultWorld := server.World()
	return &WorldManager{
		s:          server,
		folderPath: folderPath,
		log:        log,
		worlds: map[string]*world.World{
			defaultWorld.Name(): defaultWorld,
		},
	}
}

// DefaultWorld Get our default world
func (m *WorldManager) DefaultWorld() *world.World {
	return m.s.World()
}

// Worlds Create the world map data structure
func (m *WorldManager) Worlds() []*world.World {
	m.worldsMu.RLock()
	worlds := make([]*world.World, 0, len(m.worlds))
	for _, w := range m.worlds {
		worlds = append(worlds, w)
	}
	m.worldsMu.RUnlock()
	return worlds
}

// AssertWorld check if world exists and is loaded
func (m *WorldManager) AssertWorld(name string) *world.World {
	m.worldsMu.RLock()
	defer m.worldsMu.RUnlock()
	if w, ok := m.worlds[name]; ok {
		return w
	}
	panic(fmt.Errorf("expected world %v, but is not loaded", name))
}

// World Get a world and its status
func (m *WorldManager) World(name string) (*world.World, bool) {
	m.worldsMu.RLock()
	w, ok := m.worlds[name]
	m.worldsMu.RUnlock()
	return w, ok
}

// DefaultWorldSettings sets a world settings to default
func DefaultWorldSettings(w *world.World) {
	w.SetTickRange(0)
	w.SetTime(6000)
	w.StopTime()
	w.StopWeatherCycle()
	w.StopRaining()
	w.StopThundering()
	w.SetDefaultGameMode(world.GameModeSurvival)
	w.SetDifficulty(world.DifficultyNormal)
}

// LoadWorld loads a world if it hasn't been loaded yet and sets default values
func (m *WorldManager) LoadWorld(worldPath, worldName string) error {
	// check if the world has been loaded yet
	if _, ok := m.World(worldName); ok {
		return fmt.Errorf("world is already loaded")
	}
	// attempt to load the world from disk
	log := m.log.WithField("dimension", "overworld")
	log.Debugf("Loading world...")
	prov, err := mcdb.New(log, worldPath, opt.DefaultCompression)
	if err != nil {
		return fmt.Errorf("error loading world: %v", err)
	}
	// set up the disk configs and other basics
	w := world.Config{
		Dim:      world.Overworld,
		Log:      m.log,
		ReadOnly: true,
		Provider: prov,
		Entities: entity.DefaultRegistry,
	}.New()
	// set up default values
	DefaultWorldSettings(w)
	// w.Handle(&Handler{})
	w.Handle(nil) // &Handler{} is not defined so doing this for now
	// change world name
	m.worldsMu.Lock() // we have to lock it to do this
	m.worlds[worldName] = w
	m.worldsMu.Unlock() // then we can unlock it right after doing so
	// log that we loaded the world and finish
	log.Infof(`Opened world "%v".`, w.Name())
	return nil
}

// UnloadWorld Unloads the world
func (m *WorldManager) UnloadWorld(w *world.World) error {
	// can't unload the default world
	if w == m.DefaultWorld() {
		return fmt.Errorf("the default world cannot be unloaded")
	}
	// no point to unload an already unloaded world
	if _, ok := m.World(w.Name()); !ok {
		return fmt.Errorf("world isn't loaded")
	}
	// right before unloading the world teleport all players back to the default world
	m.log.Debugf("Unloading world '%v'\n", w.Name())
	for _, p := range m.s.Players() {
		if p.World() == w {
			m.DefaultWorld().AddEntity(p)
			p.Teleport(m.DefaultWorld().Spawn().Vec3Middle())
		}
	}
	// remove the world from the world map data structure
	m.worldsMu.Lock()
	delete(m.worlds, w.Name())
	m.worldsMu.Unlock()
	// error handle/log stuff if needed and return
	if err := w.Close(); err != nil {
		return fmt.Errorf("error closing world: %v", err)
	}
	m.log.Infof("Unloaded world '%v'\n", w.Name())
	return nil
}

// Close closes all the loaded worlds except the default world
func (m *WorldManager) Close() error {
	m.worldsMu.Lock()
	for _, w := range m.worlds {
		// Let dragonfly close this.
		if w == m.DefaultWorld() {
			continue
		}
		m.log.Infof("Closing world '%v'\n", w.Name())
		if err := w.Close(); err != nil {
			return err
		}
	}
	m.worlds = nil
	m.worldsMu.Unlock()
	return nil
}
