package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fouched/go-adventure/internal/models"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var cyan = color.New(color.FgCyan)

var directions = []string{"n", "s", "e", "w"}

func PlayGame() {

	adventurer := models.NewPlayer()
	currentGame := models.NewGame(adventurer)
	room := generateRoom()
	currentGame.Room = room
	welcome()

	// get the player input
	cyan.Println("Press enter to begin...")
	var input string
	fmt.Scanln(&input)
	currentGame.Room.PrintDescription()
	exploreLabyrinth(currentGame)
}

func welcome() {

	red.Println("                                                  D U N G E O N")
	green.Println(`
    The village of Honeywood has been terrorized by strange, deadly creatures for months now. Unable to endure any 
    longer, the villagers pooled their wealth and hired the most skilled adventurer they could find: you. After
    listening to their tale of woe, you agree to enter the labyrinth where most of the creatures seem to originate,
    and destroy the foul beasts. Armed with nothing but a bundle of torches, you descend into the labyrinth, 
    ready to do battle....
	`)
}

func generateRoom() models.Room {

	room := models.NewRoom()

	// there is a 25% chance that this room has an item
	if rand.IntN(100) < 26 {
		a := models.GetAllArmory()
		item := a[rand.IntN(len(a))]
		room.Items[item.Name] = item
	}

	// there is a 25% chance that this room has a monster
	if rand.IntN(100) < 26 {
		m := models.GetAllMonsters()
		monster := m[rand.IntN(len(m))]
		room.Monster = &monster
	}

	return room
}

func exploreLabyrinth(currentGame *models.Game) {

	for {

		for _, item := range currentGame.Room.Items {
			yellow.Printf("You see a %s\n", item.Name)
		}

		if currentGame.Room.Monster != nil {
			red.Printf("There is a %s here!\n", currentGame.Room.Monster.Name)
			f := getInput("Do you want to fight or flee?", []string{"fight", "flee"})
			for {
				if f == "flee" {
					cyan.Println("You turn and run, coward that you are...")
					break
				} else {
					winner := fight(currentGame)
					if winner == "player" {
						break
					} else if winner == "monster" {
						break
					} else { // flee
						break
					}

				}
			}
		}

		yellow.Print("-> ")
		input := readInput()

		// process input
		if input == "help" {
			showHelp()
			continue
		} else if input == "look" {
			currentGame.Room.PrintDescription()
			continue
		} else if strings.HasPrefix(input, "get") {
			if currentGame.Room.Items == nil {
				cyan.Println("There is nothing to pick up")
				continue
			} else {
				getAnItem(currentGame, input)
				continue
			}
		} else if strings.HasPrefix(input, "drop") {
			dropAnItem(currentGame, strings.TrimPrefix(input, "drop "))
			continue
		} else if strings.HasPrefix(input, "equip") {
			item := strings.TrimPrefix(input, "equip ")
			useItem(&currentGame.Player, item)
			continue
		} else if strings.HasPrefix(input, "use") {
			item := strings.TrimPrefix(input, "use ")
			useItem(&currentGame.Player, item)
			continue
		} else if strings.HasPrefix(input, "unequip") {
			item := strings.TrimPrefix(input, "unequip ")
			unequipItem(&currentGame.Player, item)
			continue
		} else if input == "inventory" || input == "inv" {
			showInventory(currentGame)
			continue
		} else if slices.Contains(directions, input) {
			cyan.Println("You move deeper into the dungeon.")
		} else if input == "status" {
			printStatus(&currentGame.Player)
			continue
		} else if input == "q" || input == "quit" {
			yellow.Println("Overcome with terror, you flee the dungeon.")
			playAgain()
		} else {
			cyan.Println("I'm not sure what you mean... type help for available commands.")
			continue
		}

		currentGame.Room = generateRoom()
		currentGame.Room.PrintDescription()
		currentGame.Player.Turns += 1

	}
}

func printStatus(player *models.Player) {
	cyan.Printf("You have played the game for %d turns, defeated %d monsters, and found %d gold pieces.\n",
		player.Turns, player.MonstersDefeated, player.Treasure)
	cyan.Printf("You have earned %d xp.\n", player.XP)
	cyan.Printf("You have %d hit points remaining, out of 100.\n", player.HP)
	cyan.Printf("Currently equipped weapon: %s.\n", player.CurrentWeapon.Name)
	cyan.Printf("Currently equipped armor: %s.\n", player.CurrentArmor.Name)
	cyan.Printf("Currently equipped shield: %s.\n", player.CurrentShield.Name)
}

func unequipItem(player *models.Player, item string) {
	_, hasItem := player.Inventory[item]
	if hasItem {
		if player.CurrentWeapon.Name == item {
			player.CurrentWeapon = models.GetDefaultArmory()["hands"]
			cyan.Printf("You stop using the %s.", item)
		} else if player.CurrentArmor.Name == item {
			player.CurrentArmor = models.GetDefaultArmory()["clothes"]
			cyan.Printf("You stop using the %s.", item)
		} else if player.CurrentShield.Name == item {
			player.CurrentShield = models.GetDefaultArmory()["no shield"]
			cyan.Printf("You stop using the %s.", item)
		} else {
			red.Printf("You don't have a %s equipped!", item)
		}
	} else {
		red.Printf("You don't have a %s", item)
	}
}

func useItem(player *models.Player, item string) {

	_, hasItem := player.Inventory[item]
	if hasItem {
		oldWeapon := player.CurrentWeapon

		if player.Inventory[item].Type == "weapon" {
			player.CurrentWeapon = player.Inventory[item]
			cyan.Printf("You arm yourself with a %s instead of your %s.\n", player.CurrentWeapon.Name, oldWeapon.Name)

			// you can't use a shield with a bow
			if item == "longbow" && player.CurrentShield.Name != "no shield" {
				player.CurrentShield = models.GetDefaultArmory()["no shield"]
				cyan.Printf("Since you can't use a shield with a %s, you sling it over your back.\n", player.CurrentWeapon.Name)
			}
		} else if player.Inventory[item].Type == "armor" {
			player.CurrentArmor = player.Inventory[item]
			cyan.Printf("You put on the %s.\n", player.CurrentArmor.Name)
		} else if player.Inventory[item].Type == "shield" {
			// you can't use a shield with a bow
			if player.CurrentShield.Name == "longbow" {
				red.Printf("You can't use a shield while using a bow\n")
			} else {
				player.CurrentShield = player.Inventory[item]
				cyan.Printf("You equip your %s.\n", player.CurrentShield.Name)
			}
		} else {
			red.Printf("You can't equip a %s.\n", item)
		}

	} else {
		red.Printf("You don't have an %s.\n", item)
	}

}

func dropAnItem(game *models.Game, input string) {

	item, hasItem := game.Player.Inventory[input]
	if hasItem {

		if input == game.Player.CurrentWeapon.Name {
			red.Println("You cannot drop your currently equipped weapon!")
		} else if input == game.Player.CurrentArmor.Name {
			red.Println("You cannot drop your currently equipped armor!")
		} else if input == game.Player.CurrentShield.Name {
			red.Println("You cannot drop your currently equipped shield!")
		} else {
			delete(game.Player.Inventory, input)
			cyan.Printf("You drop the %s\n", item.Name)
			game.Room.Items[item.Name] = item
		}
	} else {
		red.Printf("You are not carrying a %s\n", input)
	}
}

func showInventory(currentGame *models.Game) {
	cyan.Println("Your inventory:")
	cyan.Printf("    - %d pieces of gold.\n", currentGame.Player.Treasure)

	for _, item := range currentGame.Player.Inventory {

		if item.Name == currentGame.Player.CurrentWeapon.Name {
			cyan.Printf("    - %s (equipped)\n", item.Name)
		} else if item.Name == currentGame.Player.CurrentArmor.Name {
			cyan.Printf("    - %s (equipped)\n", item.Name)
		} else if item.Name == currentGame.Player.CurrentShield.Name {
			cyan.Printf("    - %s (equipped)\n", item.Name)
		} else {
			cyan.Printf("    - %s\n", item.Name)
		}
	}
}

func readInput() string {

	// we could also do below manually
	//reader := bufio.NewReader(os.Stdin)
	//input, _ := reader.ReadString('\n')
	//input = strings.TrimSuffix(input, "\n")
	//input = strings.TrimSuffix(input, "\r")
	var input string
	fmt.Scanln(&input)
	return strings.TrimSpace(strings.ToLower(input))
}

func getAnItem(game *models.Game, input string) {

	itemToGet := strings.TrimPrefix(input, "get")

	// cater for getting any item in a room - the user just typed get
	if len(game.Room.Items) > 0 && itemToGet == "" {
		for _, item := range game.Room.Items {
			itemToGet = item.Name
			continue
		}
	} else {
		itemToGet = strings.TrimPrefix(itemToGet, " ")
	}

	playerItem, ok := game.Player.Inventory[itemToGet]
	if ok {
		cyan.Printf("You already have a %s, and decide you don't need another.\n", playerItem.Name)
		return
	}

	roomItem, ok := game.Room.Items[itemToGet]
	if ok {
		delete(game.Room.Items, itemToGet)
		game.Player.Inventory[roomItem.Name] = roomItem
		cyan.Printf("You pick up the %s.\n", itemToGet)
	} else {
		if itemToGet == "" {
			red.Println("There is no item to pick up.")
		} else {
			red.Printf("There is no %s here\n", itemToGet)
		}
	}

}

func playAgain() {

	yn := getYN("Play again")
	if yn == "yes" {
		PlayGame()
	} else {
		yellow.Println("Until next time, adventurer.")
		os.Exit(0)
	}
}

func getYN(q string) string {

	valid := false
	answer := ""
	options := []string{"yes", "no", "y", "n"}

	for {
		cyan.Print(q + " (yes/no) -> ")
		input := readInput()

		for i := range options {
			if input == options[i] {
				valid = true
				answer = options[i]
				break
			}
		}

		if !valid {
			cyan.Println("Please enter yes (y) or no (n).")
		} else {
			if answer == "y" {
				answer = "yes"
			} else if answer == "n" {
				answer = "no"
			}
			break
		}
	}

	return answer
}

func getInput(q string, answers []string) string {

	for {
		cyan.Printf("%s", q)
		yellow.Print(" -> ")
		input := readInput()

		valid := false
		for i := range answers {
			if input == answers[i] {
				valid = true
				break
			}
		}

		if valid {
			return input
		} else {
			yellow.Println("Please enter a valid response.")
		}
	}
}

func showHelp() {

	green.Println(`Available commands:
    - n / s / e / w : move in a direction
    - map : show a map of the labyrinth
    - look : look around and describe you environment
    - use / equip <item> : use an item from your inventory
    - unequip <item> : stop using an item from your inventory
    - fight : attack a foe
    - examine <object> : examine an object more closely
    - get <item> : pick up an item
    - drop <item> : drop an item
    - rest : restore health by resting
    - inv / inventory : show current inventory
    - status : show current player status
    - q / quit : end the game`)
}
