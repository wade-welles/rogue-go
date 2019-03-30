package rogue

//"fmt"

type RadarHack struct {
	Radar []*GameObject
}

// TODO: The radar could be more sophisticated, for example it could show weapons
// or other items on the ground potentially

func (self *RadarHack) AddToRadar(o *GameObject) {
	// TODO: Write to memory that the object is spotted so it shows on radar
	// in the old version lTeam and rTeam were passeed, but in this version
	// we will just hold that data within the gameObject or store the team
	// data in the engine for each match
}

func (self *RadarHack) Render() {
	Info("Entering RadarHack.Render() function")
	// TODO: All rendering actions go here, it is called from Engine
}
