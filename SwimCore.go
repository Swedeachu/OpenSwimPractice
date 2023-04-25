package main

import (
	"SwimPractice/Commands"
	"SwimPractice/Managers"
	"fmt"
	"github.com/df-mc/dragonfly/server"
	"github.com/df-mc/dragonfly/server/player/chat"
	"github.com/pelletier/go-toml"
	"github.com/sirupsen/logrus"
	"os"
)

// main entry point of swim.gg, any essentials and start-ups are called here first
func main() {
	Commands.RegisterAllCommands() // need to register commands before starting server
	startLoggerAndServer()
}

// startLoggerAndServer starts the logger via logrus, and then starts the server
func startLoggerAndServer() {
	log := logrus.New()
	log.Formatter = &logrus.TextFormatter{ForceColors: true}
	log.Level = logrus.DebugLevel
	chat.Global.Subscribe(chat.StdoutSubscriber{})
	// read our config for some basic settings
	conf, err := readConfig(log)
	if err != nil {
		log.Fatalln(err)
	}
	// get the server instance set up
	Managers.Swim = &Managers.SwimPractice{}
	Managers.Swim.ServerInstance = conf.New()
	Managers.Swim.ServerInstance.CloseOnProgramEnd()
	// load the worlds before listening
	Managers.LoadWorlds()
	// set up the session manager
	Managers.Swim.SessionDataManager = Managers.NewPlayerSessionDataManager()
	Managers.SessionManagerRoutineStart(Managers.Swim.SessionDataManager)
	// start listening (enable the server)
	Managers.Swim.ServerInstance.Listen()
	Managers.JoinManager(Managers.Swim.ServerInstance) // call the function which constantly listens for player connections
}

// readConfig reads the configuration from the config.toml file, or creates the file if it does not yet exist.
// this function is called from startLoggerAndServer
func readConfig(log server.Logger) (server.Config, error) {
	c := server.DefaultConfig()
	var zero server.Config
	if _, err := os.Stat("config.toml"); os.IsNotExist(err) {
		data, err := toml.Marshal(c)
		if err != nil {
			return zero, fmt.Errorf("encode default config: %v", err)
		}
		if err := os.WriteFile("config.toml", data, 0644); err != nil {
			return zero, fmt.Errorf("create default config: %v", err)
		}
		return c.Config(log)
	}
	data, err := os.ReadFile("config.toml")
	if err != nil {
		return zero, fmt.Errorf("read config: %v", err)
	}
	if err := toml.Unmarshal(data, &c); err != nil {
		return zero, fmt.Errorf("decode config: %v", err)
	}
	return c.Config(log)
}
