package models

type Player struct {
	HP               int
	Treasure         int
	MonstersDefeated int
	XP               int
	Turns            int
	Inventory        map[string]ArmoryItem
}

func NewPlayer() *Player {
	return &Player{
		HP:               100,
		Treasure:         0,
		MonstersDefeated: 0,
		XP:               0,
		Turns:            0,
		Inventory:        make(map[string]ArmoryItem), // []ArmoryItem{}, // maybe this should rather be a map (for editing it)
	}
}
