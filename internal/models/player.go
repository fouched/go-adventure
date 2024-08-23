package models

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
}

func NewPlayer() Player {
	return Player{
		HP:               100,
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
