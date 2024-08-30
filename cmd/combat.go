package main

import (
	"fmt"
	"github.com/fouched/go-adventure/internal/models"
	"math/rand/v2"
	"time"
)

func fight(currentGame *models.Game) string {
	winner := ""
	rm := currentGame.Room
	pl := currentGame.Player

	playerTurn := true
	if rand.IntN(2) == 1 {
		playerTurn = false
	}

	if playerTurn {
		cyan.Printf("You brace yourself and attack the %s.\n", rm.Monster.Name)
	} else {
		cyan.Printf("The %s moves quickly and attacks first!\n", rm.Monster.Name)
	}

	monsterHp := rand.IntN(rm.Monster.MaxHP-rm.Monster.MinHP) + rm.Monster.MinDamage
	monsterOriginalHp := monsterHp

	for {
		if playerTurn {
			// 50% chance to hit
			r := rand.IntN(100)
			mr := r + pl.CurrentWeapon.ToHit - rm.Monster.ArmorModifier

			if mr >= 50 {
				green.Printf("You hit the %s with your %s!\n", rm.Monster.Name, pl.CurrentWeapon.Name)
				damage := rand.IntN(pl.CurrentWeapon.MaxDamage-pl.CurrentWeapon.MinDamage) + pl.CurrentWeapon.MinDamage
				monsterHp = monsterHp - damage
			} else {
				green.Printf("You attack the %s with you %s and miss!\n", rm.Monster.Name, pl.CurrentWeapon.Name)
			}

			if monsterHp <= 0 {
				green.Printf("The %s falls on the floor, dead.\n", rm.Monster.Name)
				winner = "player"
			}
		} else {
			// 50% chance to hit
			r := rand.IntN(100)
			mr := r - (pl.CurrentShield.Defense + pl.CurrentArmor.Defense)

			if mr >= 50 {
				red.Printf("The %s attacks and hits!\n", rm.Monster.Name)
				damage := rand.IntN(rm.Monster.MaxDamage-rm.Monster.MinDamage) + rm.Monster.MinDamage
				pl.HP = pl.HP - damage
			} else {
				green.Printf("The %s attacks and misses!\n", rm.Monster.Name)
			}

			if pl.HP <= 0 {
				green.Printf("The %s kills you, and you fall on the floor, dead.\n", rm.Monster.Name)
				winner = "monster"
			}
		}

		if pl.HP <= 0 || monsterHp <= 0 {
			break
		}

		// feedback on monster state
		if monsterHp < monsterOriginalHp/2 {
			yellow.Printf("The %s is bleeding.\n", rm.Monster.Name)
		} else if monsterHp < monsterOriginalHp/3 {
			yellow.Printf("The %s is bleeding profusely, and looks to be nearly dead.\n", rm.Monster.Name)
		}

		// feedback on player state
		if pl.HP <= 0.2*PLAYER_HP {
			answer := getYN("You are near death. Do you want to continue")
			if answer == "no" {
				return "flee"
			}
		} else if pl.HP <= 0.4*PLAYER_HP {
			answer := getYN("You are badly wounded. Do you want to continue")
			if answer == "no" {
				return "flee"
			}
		} else if pl.HP <= 0.6*PLAYER_HP {
			answer := getYN("You are wounded. Do you want to continue")
			if answer == "no" {
				return "flee"
			}
		} else {
			cyan.Println("You are only lightly wounded. 'Tis but a scratch.")
		}

		time.Sleep(1 * time.Second)
		fmt.Println("")
		playerTurn = !playerTurn
	}

	return winner
}
