package rogue

//"fmt"

// TODO: Create a 'Hack' interface so glow, radar and whatever else can be
// held in the Engine object together

type GlowGroup struct {
	Enabled  bool
	Active   bool
	Distance int
	Color    GlowColor
}

type GlowColor struct {
	Red     float64
	Blue    float64
	Green   float64
	Opacity float64
}

type GlowHack struct {
	Enabled bool
	Active  bool
	Allies  *GlowGroup
	Enemies *GlowGroup
	Weapons *GlowGroup
	Items   *GlowGroup
}

func InitGlowHack() *GlowHack {
	// TODO: Load this data via YAML config
	return &GlowHack{
		Enabled: true,
		Active:  true,
		Allies: &GlowGroup{
			Enabled:  false,
			Active:   false,
			Distance: 9999,
			Color: GlowColor{
				Red:     0.8,
				Blue:    0.3,
				Green:   0.3,
				Opacity: 0.6,
			},
		},
		Enemies: &GlowGroup{
			Enabled:  true,
			Active:   true,
			Distance: 500,
			Color: GlowColor{
				Red:     0.8,
				Blue:    0.3,
				Green:   0.3,
				Opacity: 0.6,
			},
		},
		Weapons: &GlowGroup{
			Enabled:  false,
			Active:   false,
			Distance: 9999,
			Color: GlowColor{
				Red:     0.8,
				Blue:    0.8,
				Green:   0.8,
				Opacity: 0.8,
			},
		},
		Items: &GlowGroup{
			Enabled:  false,
			Active:   false,
			Distance: 9999,
			Color: GlowColor{
				Red:     0.8,
				Blue:    0.8,
				Green:   0.8,
				Opacity: 0.8,
			},
		},
	}
}

func (self *GlowHack) Start() bool {
	if self.Enabled {
		self.Active = true
		//TODO: Render objects
	}
	return self.Active
}

func (self *GlowHack) Stop() bool {
	if self.Active {
		self.Active = false
		//TODO: Stop rendering glow objects
	}
	return self.Active
}

func (self *GlowHack) GlowObject(object *GameObject) {
	// TODO: Render the object over the game display
}

func (self *GlowHack) GlowObjects(objects []*GameObject) {
	for _, object := range objects {
		self.GlowObject(object)
	}
}

func (self *GlowHack) RemoveObjectGlow(object *GameObject) {

}

func (self *GlowHack) RemoveObjectsGlow(objects []*GameObject) {
	for _, object := range objects {
		self.RemoveObjectGlow(object)
	}
}

func (self *GlowHack) Render() {
	Info("Entering GlowHack.Render() function")
	// TODO: This is where we would do actual rendering over display and itll
	// be called from the primary one in Engine
}
