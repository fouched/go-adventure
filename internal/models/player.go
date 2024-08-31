package models

import "github.com/fouched/go-adventure/internal/config"

type Player struct {
	HP               int
	Treasure         int
	MonstersDefeated int
	XP               int
	Turns            int
	Inventory        map[string]ArmoryItem
	CurrentWeapon    ArmoryItem
	CurrentArmor     ArmoryItem
	CurrentShield    ArmoryItem
	CoordX           int
	CoordY           int
	Visited          []string
}

func NewPlayer() Player {
	return Player{
		HP:               config.PLAYER_HP,
		Treasure:         0,
		MonstersDefeated: 0,
		XP:               0,
		Turns:            0,
		Inventory:        make(map[string]ArmoryItem),
		CurrentWeapon:    GetDefaultArmory()["hands"],
		CurrentArmor:     GetDefaultArmory()["clothes"],
		CurrentShield:    GetDefaultArmory()["no shield"],
	}
}
