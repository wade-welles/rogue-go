package rogue

import (
	"os"
	"time"
)

type Engine struct {
	Display    string
	Process    *GameProcess
	Delay      time.Duration
	ActiveGame *Game
	GlowHack   *GlowHack
	RadarHack  *RadarHack
}

func NewEngine() *Engine {
	// TODO: Pull settings from YAML config file
	engine := &Engine{
		Process: FindGameProcess(),
		Display: os.Getenv("DISPLAY"),
		Delay:   time.Duration(950 * time.Millisecond),
	}
	if engine.Display == "" {
		Error(nil, "Active display must be assigned to the environmental variable: "+Yellow("env")+" "+Yellow("DISPLAY"))
		os.Exit(1)
	} else {
		Info("Using display: " + Yellow(engine.Display))
	}
	return engine
}

func (self *Engine) IsConnected() bool {
	// TODO: Use SignOnState memory data pulled from the process data in `/proc/*` subsystem
	return false
}

func (self *Engine) Render() {
	Info("Starting Engine.Render() function which calls Game.Render(), RadarHack.Render(), and GlowHack.Render() in its main render loop.")
	// TODO: Should check also if game is active, and if connected probably
	if self.Process == nil {
		return
	}
	//ticker := time.NewTicker(20 * time.Millisecond)
	tick := time.Tick(self.Delay)
	for {
		select {
		case <-tick:
			self.ActiveGame.Render()
			self.RadarHack.Render()
			self.GlowHack.Render()
		}
	}

}
