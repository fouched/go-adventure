package main

import (
	"bufio"
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
	currentGame := models.NewGame(*adventurer)
	room := generateRoom()
	currentGame.Room = *room
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

func generateRoom() *models.Room {

	room := models.NewRoom()

	// there is a 25% chance that this room has an item
	if rand.IntN(100) < 99 {
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
			cyan.Printf("You see a %s\n", item.Name)
		}

		if currentGame.Room.Monster != nil {
			red.Printf("There is a %s here!\n", currentGame.Room.Monster.Name)
		}

		var input string
		yellow.Print("-> ")
		input = readInput(input)

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
		} else if input == "inventory" || input == "inv" {
			showInventory(currentGame)
			continue
		} else if slices.Contains(directions, input) {
			cyan.Println("You move deeper into the dungeon.")
		} else if input == "q" || input == "quit" {
			yellow.Println("Overcome with terror, you flee the dungeon.")
			playAgain()
		} else {
			cyan.Println("I'm not sure what you mean... type help for available commands.")
			continue
		}

		currentGame.Room = *generateRoom()
		currentGame.Room.PrintDescription()

	}
}

func dropAnItem(game *models.Game, input string) {

	item, ok := game.Player.Inventory[input]
	if ok {
		delete(game.Player.Inventory, input)
		cyan.Printf("You drop the %s\n", item.Name)
		game.Room.Items[item.Name] = item
	} else {
		red.Printf("You are not carrying a %s\n", input)
	}
}

func showInventory(currentGame *models.Game) {
	cyan.Println("Your inventory:")
	for _, item := range currentGame.Player.Inventory {
		cyan.Printf("    - %s\n", item.Name)
	}
}

func readInput(input string) string {

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")
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
		var input string
		cyan.Print(q + " (yes/no) -> ")
		fmt.Scanln(&input)
		input = strings.TrimSpace(strings.ToLower(input))

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

func showHelp() {

	green.Println(`Available commands:
    - n / s / e / w : move in a direction
    - map : show a map of the labyrinth
    - look : look around and describe you environment
    - equip <item> : use an item from your inventory
    - unequip <item> : stop using an item from your inventory
    - fight : attack a foe
    - examine <object> : examine an object more closely
    - get <item> : pick up an item
    - drop <item> : drop an item
    - rest : restore health by resting
    - inv/ inventory : show current inventory
    - status : show current player status
    - q / quit : end the game`)
}
