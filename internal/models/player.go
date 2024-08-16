package models

type Player struct {
	HP               int
	Treasure         int
	MonstersDefeated int
	XP               int
	Turns            int
}

func NewPlayer() *Player {
	return &Player{
		HP:               100,
		Treasure:         0,
		MonstersDefeated: 0,
		XP:               0,
		Turns:            0,
	}
}
