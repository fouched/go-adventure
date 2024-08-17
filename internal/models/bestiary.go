package models

type Monster struct {
	Name          string
	MinHP         int
	MaxHP         int
	MinDamage     int
	MaxDamage     int
	ArmorModifier int
}

func GetAllMonsters() []Monster {
	var m = []Monster{
		{
			Name:          "Orc",
			MinHP:         10,
			MaxHP:         20,
			MinDamage:     1,
			MaxDamage:     6,
			ArmorModifier: -20,
		},
		{
			Name:          "Goblin",
			MinHP:         10,
			MaxHP:         30,
			MinDamage:     1,
			MaxDamage:     10,
			ArmorModifier: 5,
		},
		{
			Name:          "Troll",
			MinHP:         25,
			MaxHP:         60,
			MinDamage:     5,
			MaxDamage:     25,
			ArmorModifier: 10,
		},
		{
			Name:          "Zombie",
			MinHP:         5,
			MaxHP:         20,
			MinDamage:     1,
			MaxDamage:     20,
			ArmorModifier: -10,
		},
		{
			Name:          "Vampire",
			MinHP:         25,
			MaxHP:         75,
			MinDamage:     10,
			MaxDamage:     75,
			ArmorModifier: 15,
		},
		{
			Name:          "Dragon",
			MinHP:         50,
			MaxHP:         100,
			MinDamage:     10,
			MaxDamage:     50,
			ArmorModifier: 25,
		},
		{
			Name:          "Ghoul",
			MinHP:         30,
			MaxHP:         60,
			MinDamage:     5,
			MaxDamage:     30,
			ArmorModifier: 20,
		},
		{
			Name:          "Kobold",
			MinHP:         20,
			MaxHP:         40,
			MinDamage:     5,
			MaxDamage:     20,
			ArmorModifier: 15,
		},
		{
			Name:          "Werewolf",
			MinHP:         60,
			MaxHP:         90,
			MinDamage:     10,
			MaxDamage:     50,
			ArmorModifier: 10,
		},
		{
			Name:          "Giant",
			MinHP:         50,
			MaxHP:         100,
			MinDamage:     20,
			MaxDamage:     60,
			ArmorModifier: 25,
		},
		{
			Name:          "Wraith",
			MinHP:         20,
			MaxHP:         70,
			MinDamage:     15,
			MaxDamage:     30,
			ArmorModifier: 40,
		},
	}

	return m
}
