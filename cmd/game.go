package main

import (
	"bufio"
	"fmt"
	"github.com/fatih/color"
	"github.com/fouched/go-adventure/internal/clr"
	"github.com/fouched/go-adventure/internal/config"
	"github.com/fouched/go-adventure/internal/models"
	"math/rand/v2"
	"os"
	"slices"
	"strings"
)

var directions = []string{"n", "s", "e", "w"}

func PlayGame() {

	adventurer := models.NewPlayer()
	currentGame := models.NewGame(adventurer, config.MAX_X_AXIS, config.MAX_Y_AXIS)

	allRooms, numMonsters := createWorld(currentGame)
	currentGame.NumMonsters = numMonsters
	currentGame.Rooms = allRooms

	entrance := "0,0"
	currentGame.Room = currentGame.Rooms[entrance]
	currentGame.Entrance = entrance
	currentGame.Room.Location = entrance

	welcome(currentGame)

	// get the player input
	clr.Cyan.Println("Press enter to begin...")
	var input string
	fmt.Scanln(&input)
	currentGame.Room.PrintDescription()
	exploreLabyrinth(currentGame)
}

func welcome(currentGame *models.Game) {

	clr.Red.Println("                                                  D U N G E O N")
	clr.Green.Println(`
    The village of Honeywood has been terrorized by strange, deadly creatures for months now. Unable to endure any 
    longer, the villagers pooled their wealth and hired the most skilled adventurer they could find: you. After
    listening to their tale of woe, you agree to enter the labyrinth where most of the creatures seem to originate,
    and destroy the foul beasts. Armed with nothing but a bundle of torches, you descend into the labyrinth, 
    ready to do battle....
	`)
	fmt.Println("")
	clr.Yellow.Printf("According to the people of Honeywood there are %d creatures in this labyrinth.\n", currentGame.NumMonsters)
	fmt.Println("")
}

func generateRoom(location string) models.Room {

	room := models.NewRoom(location)

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
			clr.Yellow.Printf("You see a %s\n", item.Name)
		}

		if currentGame.Room.Monster != nil {
			clr.Red.Printf("There is a %s here!\n", currentGame.Room.Monster.Name)
			f := getInput("Do you want to fight or flee?", []string{"fight", "flee"})
			for {
				if f == "flee" {
					clr.Cyan.Println("You turn and run, coward that you are...")
					break
				} else {
					winner := fight(currentGame)
					if winner == "player" {
						gold := rand.IntN(100) + 1
						clr.Cyan.Printf("You search the monster's body and find %d pieces of gold.\n", gold)
						currentGame.Player.Treasure = currentGame.Player.Treasure + gold
						currentGame.Player.XP = currentGame.Player.XP + 100
						currentGame.Player.MonstersDefeated = currentGame.Player.MonstersDefeated + 1
						currentGame.Room.Monster = nil
						break
					} else if winner == "monster" {
						clr.Red.Printf("You have failed in your mission, and your body lies in the labyrinth forever.\n")
						playAgain()
						break
					} else {
						clr.Cyan.Println("You flee in terror from the monster.\n")
						break
					}
				}
			}
		}

		clr.Yellow.Print("-> ")
		input := readInput()

		// process input
		if input == "help" {
			showHelp()
			continue
		} else if input == "look" {
			currentGame.Room.PrintDescription()
			continue
		} else if input == "map" {
			showMap(currentGame)
			continue
		} else if strings.HasPrefix(input, "get") {
			if currentGame.Room.Items == nil {
				clr.Cyan.Println("There is nothing to pick up")
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
		} else if input == "r" || input == "rest" {
			rest(&currentGame.Player)
			continue
		} else if input == "inv" || input == "inventory" {
			showInventory(currentGame)
			continue
		} else if slices.Contains(directions, input) {
			direction := input
			if currentGame.Room.Location == currentGame.Entrance && direction == "s" {
				yn := getYN("You are about to leave the dungeon, are you sure")
				if yn == "yes" {
					playAgain()
				} else {
					continue
				}
			}

			if direction == "n" {
				if currentGame.Player.CoordY < currentGame.Y {
					currentGame.Player.CoordY = currentGame.Player.CoordY + 1
				} else {
					clr.Red.Println("You bump into a stone wall.")
					continue
				}
			} else if direction == "s" {
				if currentGame.Player.CoordY > currentGame.Y*-1 {
					currentGame.Player.CoordY = currentGame.Player.CoordY - 1
				} else {
					clr.Red.Println("You bump into a stone wall.")
					continue
				}
			} else if direction == "e" {
				if currentGame.Player.CoordX < currentGame.X {
					currentGame.Player.CoordX = currentGame.Player.CoordX + 1
				} else {
					clr.Red.Println("You bump into a stone wall.")
					continue
				}
			} else if direction == "w" {
				if currentGame.Player.CoordX > currentGame.X*-1 {
					currentGame.Player.CoordX = currentGame.Player.CoordX - 1
				} else {
					clr.Red.Println("You bump into a stone wall.")
					continue
				}
			}
			clr.Cyan.Println("You move deeper into the dungeon.")

		} else if input == "status" {
			printStatus(&currentGame.Player)
			continue
		} else if input == "q" || input == "quit" {
			clr.Yellow.Println("Overcome with terror, you flee the dungeon.")
			playAgain()
		} else {
			clr.Cyan.Println("I'm not sure what you mean... type help for available commands.")
			continue
		}

		newLocation := fmt.Sprintf("%d,%d", currentGame.Player.CoordX, currentGame.Player.CoordY)
		currentGame.Room = currentGame.Rooms[newLocation]
		currentGame.Room.Location = newLocation

		if slices.Contains(currentGame.Player.Visited, newLocation) {
			clr.Yellow.Println("This place seems familiar...")
		} else {
			currentGame.Player.Visited = append(currentGame.Player.Visited, newLocation)
		}

		currentGame.Room.PrintDescription()
		currentGame.Player.Turns += 1

	}
}

func rest(player *models.Player) {

	if player.HP == config.PLAYER_HP {
		clr.Cyan.Println("You are fully rested, and feel great. There is no point in setting around...\n")
	} else {
		player.HP = player.HP + rand.IntN(10) + 1
		if player.HP > config.PLAYER_HP {
			player.HP = config.PLAYER_HP
		}
		clr.Cyan.Printf("You feel better (%d/%d) hit points.\n", player.HP, config.PLAYER_HP)
	}
}

func showMap(currentGame *models.Game) {

	// print top line
	for i := 0; i < config.MAX_X_AXIS*6+3; i++ {
		clr.Yellow.Print("-")
	}
	fmt.Println("")

	for y := config.MAX_Y_AXIS; y >= config.MAX_Y_AXIS*-1; y-- {
		for x := config.MAX_X_AXIS * -1; x <= config.MAX_X_AXIS; x++ {

			if fmt.Sprintf("%d,%d", x, y) == currentGame.Room.Location {
				// current location
				red := color.New(color.FgRed)
				whiteBg := red.Add(color.BgWhite)
				whiteBg.Print(" X ")
			} else if fmt.Sprintf("%d,%d", x, y) == currentGame.Entrance {
				green := color.New(color.FgGreen)
				whiteBg := green.Add(color.BgWhite)
				whiteBg.Print(" E ")
			} else if slices.Contains(currentGame.Player.Visited, fmt.Sprintf("%d,%d", x, y)) {
				// a place we've visited
				testRoom := currentGame.Rooms[fmt.Sprintf("%d,%d", x, y)]
				if testRoom.Monster == nil {
					fmt.Print("   ")
				} else {
					clr.Red.Print(" M ")
				}
			} else {
				clr.Yellow.Print(" ? ")
			}
		}
		fmt.Println("")
	}

	// print bottom line
	for i := 0; i < config.MAX_X_AXIS*6+3; i++ {
		clr.Yellow.Print("-")
	}

	// print the legend
	fmt.Println("")
	red := color.New(color.FgRed)
	whiteBg := red.Add(color.BgWhite)
	whiteBg.Print(" X ")
	clr.Red.Print(": You  ")
	clr.Red.Print(" M : Monster  ")
	green := color.New(color.FgGreen)
	whiteBg = green.Add(color.BgWhite)
	whiteBg.Print(" E ")
	clr.Green.Print(": Exit")

	fmt.Println("")

}

func printStatus(player *models.Player) {

	clr.Cyan.Printf("You have played the game for %d turns, defeated %d monsters, and found %d gold pieces.\n",
		player.Turns, player.MonstersDefeated, player.Treasure)
	clr.Cyan.Printf("You have earned %d xp.\n", player.XP)
	clr.Cyan.Printf("You have %d hit points remaining, out of 100.\n", player.HP)
	clr.Cyan.Printf("Currently equipped weapon: %s.\n", player.CurrentWeapon.Name)
	clr.Cyan.Printf("Currently equipped armor: %s.\n", player.CurrentArmor.Name)
	clr.Cyan.Printf("Currently equipped shield: %s.\n", player.CurrentShield.Name)
}

func unequipItem(player *models.Player, item string) {
	_, hasItem := player.Inventory[item]
	if hasItem {
		if player.CurrentWeapon.Name == item {
			player.CurrentWeapon = models.GetDefaultArmory()["hands"]
			clr.Cyan.Printf("You stop using the %s.", item)
		} else if player.CurrentArmor.Name == item {
			player.CurrentArmor = models.GetDefaultArmory()["clothes"]
			clr.Cyan.Printf("You stop using the %s.", item)
		} else if player.CurrentShield.Name == item {
			player.CurrentShield = models.GetDefaultArmory()["no shield"]
			clr.Cyan.Printf("You stop using the %s.", item)
		} else {
			clr.Red.Printf("You don't have a %s equipped!", item)
		}
	} else {
		clr.Red.Printf("You don't have a %s", item)
	}
}

func useItem(player *models.Player, item string) {

	_, hasItem := player.Inventory[item]
	if hasItem {
		oldWeapon := player.CurrentWeapon

		if player.Inventory[item].Type == "weapon" {
			player.CurrentWeapon = player.Inventory[item]
			clr.Cyan.Printf("You arm yourself with a %s instead of your %s.\n", player.CurrentWeapon.Name, oldWeapon.Name)

			// you can't use a shield with a bow
			if item == "longbow" && player.CurrentShield.Name != "no shield" {
				player.CurrentShield = models.GetDefaultArmory()["no shield"]
				clr.Cyan.Printf("Since you can't use a shield with a %s, you sling it over your back.\n", player.CurrentWeapon.Name)
			}
		} else if player.Inventory[item].Type == "armor" {
			player.CurrentArmor = player.Inventory[item]
			clr.Cyan.Printf("You put on the %s.\n", player.CurrentArmor.Name)
		} else if player.Inventory[item].Type == "shield" {
			// you can't use a shield with a bow
			if player.CurrentShield.Name == "longbow" {
				clr.Red.Printf("You can't use a shield while using a bow\n")
			} else {
				player.CurrentShield = player.Inventory[item]
				clr.Cyan.Printf("You equip your %s.\n", player.CurrentShield.Name)
			}
		} else {
			clr.Red.Printf("You can't equip a %s.\n", item)
		}

	} else {
		clr.Red.Printf("You don't have an %s.\n", item)
	}

}

func dropAnItem(game *models.Game, input string) {

	item, hasItem := game.Player.Inventory[input]
	if hasItem {

		if input == game.Player.CurrentWeapon.Name {
			clr.Red.Println("You cannot drop your currently equipped weapon!")
		} else if input == game.Player.CurrentArmor.Name {
			clr.Red.Println("You cannot drop your currently equipped armor!")
		} else if input == game.Player.CurrentShield.Name {
			clr.Red.Println("You cannot drop your currently equipped shield!")
		} else {
			delete(game.Player.Inventory, input)
			clr.Cyan.Printf("You drop the %s\n", item.Name)
			game.Room.Items[item.Name] = item
		}
	} else {
		clr.Red.Printf("You are not carrying a %s\n", input)
	}
}

func showInventory(currentGame *models.Game) {
	clr.Cyan.Println("Your inventory:")
	clr.Cyan.Printf("    - %d pieces of gold.\n", currentGame.Player.Treasure)

	for _, item := range currentGame.Player.Inventory {

		if item.Name == currentGame.Player.CurrentWeapon.Name {
			clr.Cyan.Printf("    - %s (equipped)\n", item.Name)
		} else if item.Name == currentGame.Player.CurrentArmor.Name {
			clr.Cyan.Printf("    - %s (equipped)\n", item.Name)
		} else if item.Name == currentGame.Player.CurrentShield.Name {
			clr.Cyan.Printf("    - %s (equipped)\n", item.Name)
		} else {
			clr.Cyan.Printf("    - %s\n", item.Name)
		}
	}
}

func readInput() string {

	// we could also do below manually
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
	// cannot use below it splits spaces...
	//var input string
	//fmt.Scanln(&input)
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
		clr.Cyan.Printf("You already have a %s, and decide you don't need another.\n", playerItem.Name)
		return
	}

	roomItem, ok := game.Room.Items[itemToGet]
	if ok {
		delete(game.Room.Items, itemToGet)
		game.Player.Inventory[roomItem.Name] = roomItem
		clr.Cyan.Printf("You pick up the %s.\n", itemToGet)
	} else {
		if itemToGet == "" {
			clr.Red.Println("There is no item to pick up.")
		} else {
			clr.Red.Printf("There is no %s here\n", itemToGet)
		}
	}

}

func playAgain() {

	yn := getYN("Play again")
	if yn == "yes" {
		PlayGame()
	} else {
		clr.Yellow.Println("Until next time, adventurer.")
		os.Exit(0)
	}
}

func getYN(q string) string {

	valid := false
	answer := ""
	options := []string{"yes", "no", "y", "n"}

	for {
		clr.Cyan.Print(q + " (yes/no) -> ")
		input := readInput()

		for i := range options {
			if input == options[i] {
				valid = true
				answer = options[i]
				break
			}
		}

		if !valid {
			clr.Cyan.Println("Please enter yes (y) or no (n).")
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
		clr.Cyan.Printf("%s", q)
		clr.Yellow.Print(" -> ")
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
			clr.Yellow.Println("Please enter a valid response.")
		}
	}
}

func showHelp() {

	clr.Green.Println(`Available commands:
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
