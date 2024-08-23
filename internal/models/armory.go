package models

type ArmoryItem struct {
	Name      string
	MinDamage int
	MaxDamage int
	ToHit     int
	Defense   int
	Type      string
}

func GetAllArmory() []ArmoryItem {

	var a = []ArmoryItem{
		{
			Name:      "longsword",
			MinDamage: 5,
			MaxDamage: 25,
			ToHit:     25,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "dagger",
			MinDamage: 1,
			MaxDamage: 10,
			ToHit:     10,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "flail",
			MinDamage: 10,
			MaxDamage: 25,
			ToHit:     15,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "mace",
			MinDamage: 10,
			MaxDamage: 30,
			ToHit:     15,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "broadsword",
			MinDamage: 10,
			MaxDamage: 25,
			ToHit:     20,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "broken sword",
			MinDamage: 2,
			MaxDamage: 10,
			ToHit:     0,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "longbow",
			MinDamage: 5,
			MaxDamage: 20,
			ToHit:     20,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "glowing sword",
			MinDamage: 20,
			MaxDamage: 50,
			ToHit:     25,
			Defense:   0,
			Type:      "weapon",
		},
		{
			Name:      "battered shield",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   5,
			Type:      "shield",
		},
		{
			Name:      "tower shield",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   25,
			Type:      "shield",
		},
		{
			Name:      "kite shield",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   20,
			Type:      "shield",
		},
		{
			Name:      "round shield",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   15,
			Type:      "shield",
		},
		{
			Name:      "small shield",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   10,
			Type:      "shield",
		},
		{
			Name:      "leather armor",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   10,
			Type:      "armor",
		},
		{
			Name:      "chain mail",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   20,
			Type:      "armor",
		},
		{
			Name:      "ring mail",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   15,
			Type:      "armor",
		},
		{
			Name:      "water skin",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   0,
			Type:      "item",
		},
		{
			Name:      "penny",
			MinDamage: 0,
			MaxDamage: 0,
			ToHit:     0,
			Defense:   0,
			Type:      "item",
		},
	}

	return a
}

func GetDefaultArmory() map[string]ArmoryItem {

	a := make(map[string]ArmoryItem)

	a["hands"] = ArmoryItem{
		Name:      "hands",
		MinDamage: 1,
		MaxDamage: 5,
		ToHit:     0,
		Defense:   0,
		Type:      "weapon",
	}

	a["clothes"] = ArmoryItem{
		Name:      "clothes",
		MinDamage: 0,
		MaxDamage: 0,
		ToHit:     0,
		Defense:   0,
		Type:      "armor",
	}
	a["no shield"] = ArmoryItem{
		Name:      "no shield",
		MinDamage: 0,
		MaxDamage: 0,
		ToHit:     0,
		Defense:   0,
		Type:      "shield",
	}

	return a
}
