package rogue

//"fmt"

type ObjectType int

const (
	Avatar ObjectType = iota
	Allies
	Enemies
	Weapons
	Items
)

type Team int

const (
	CounterTerrorists = iota
	Terrorists
)

type Game struct {
	Avatar  *GameObject
	Allies  Team
	Enemies Team
	Round   int
	Alive   bool
	Health  int
	Objects []*GameObject
	// TODO: Just implement these by selecting them from the object group
	// by using an attribute in the GameObject struct
	//Weapons           []*GameObject
	//Items             []*GameObject
	//Players           []*GameObject
	//Terrorists        []*GameObject
	//CounterTerrorists []*GameObject
}

// TODO: All vector.cpp stuff will be moved in here, mainly
// regarding distance calculations so we can make glow hack
// work at distance
type Vector struct {
	X float64
	Y float64
	Z float64
}

type GameObject struct {
	ID       int
	Type     ObjectType
	Position Vector
	Health   int
	Alive    bool
}

//typedef struct {
//    char __buff_0x00[0xC];//0x00
//    float x;//0xC
//    char __buff_0x10[0xC];//0x10
//    float y;//0x1c
//    char __buff_0x20[0xC];//0x20
//    float z;//0x2c
//} BoneMatrix;

type BoneMatrix struct {
	Vector Vector
	Offset []byte // ? Is this the right name, I dont really understand why its called buff
}

func (self *Game) NewRound() {
	self.Round++
	self.Health = 100
	self.Alive = true
	// TODO Set active team
}

func (self *Game) ParseObjects() {
	// TODO: Use a mutex to update the entity list in a thread safe manner
	// might as well also return the list after parsing

	// TODO: Previously thre was GetEntityList, UpdateEntityList, but we can combine them
	// into one. And just make the list a public object within the engine.

	// TODO: This is where we will assign our avatar data and our team data as well from
	// memory
}

func (self *Game) ClearObjects() {
	self.Objects = []*GameObject{}
}

func (self *Game) AddGameObject(o *GameObject) {
	self.Objects = append(self.Objects, o)
}

func (self *Game) ObjectWithID(id int) *GameObject {
	// TODO: Use a read mutex to ensure the data being pulled is not changed
	// and ensure the entity does not equal 0 or nil
	for _, object := range self.Objects {
		if object.ID == id {
			return object
		}
	}
	return nil
}

func (self *Game) ObjectByIndex(index int) *GameObject {
	if index < len(self.Objects) {
		return self.Objects[index]
	}
	return nil
}

func (self *Game) Render() {
	Info("Entering Game.Render() function")
	self.ParseObjects()
}

type Weapon int

const (
	NONE              Weapon = 0
	DEAGLE            Weapon = 1
	ELITE             Weapon = 2
	FIVESEVEN         Weapon = 3
	GLOCK             Weapon = 4
	AK47              Weapon = 7
	AUG               Weapon = 8
	AWP               Weapon = 9
	FAMAS             Weapon = 10
	G3SG1             Weapon = 11
	GALILAR           Weapon = 13
	M249              Weapon = 14
	M4A1              Weapon = 16
	MAC10             Weapon = 17
	P90               Weapon = 19
	UMP45             Weapon = 24
	XM1014            Weapon = 25
	BIZON             Weapon = 26
	MAG7              Weapon = 27
	NEGEV             Weapon = 28
	SAWEDOFF          Weapon = 29
	TEC9              Weapon = 30
	TASER             Weapon = 31
	HKP2000           Weapon = 32
	MP7               Weapon = 33
	MP9               Weapon = 34
	NOVA              Weapon = 35
	P250              Weapon = 36
	SCAR20            Weapon = 38
	SG556             Weapon = 39
	SSG08             Weapon = 40
	KNIFE             Weapon = 42
	FLASHBANG         Weapon = 43
	HEGRENADE         Weapon = 44
	SMOKEGRENADE      Weapon = 45
	MOLOTOV           Weapon = 46
	DECOY             Weapon = 47
	INCENDIARYGRENADE Weapon = 48
	C4                Weapon = 49
	TERRORISTKNIFE    Weapon = 59
	M4A1SILENCER      Weapon = 60
	USPSILENCER       Weapon = 61
	CZ75A             Weapon = 63
	REVOLVER          Weapon = 64
	KNIFEBAYONET      Weapon = 500
	KNIFEFLIP         Weapon = 505
	KNIFEGUT          Weapon = 506
	KNIFEKARMABIT     Weapon = 507
	KNIFEM9BAYONET    Weapon = 508
	KNIFETACTICAL     Weapon = 509
	KNIFEFALCHION     Weapon = 512
	KNIFEBOWIE        Weapon = 514
	KNIFEBUTTERFLY    Weapon = 515
	KNIFEPUSH         Weapon = 516
)
