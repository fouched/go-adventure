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
	currentGame := models.NewGame(*adventurer)

	welcome()
	green.Print("Press enter to begin...")
	var input string
	fmt.Scanln(&input)

	exploreLabyrinth(currentGame)
}

func welcome() {
	red.Println("                                                  D U N G E O N")
	green.Println(`
    The village of Honeywood has been terrorized by strange, deadly creatures for months now. Unable to endure any 
    longer, the villagers pooled their wealth and hired the most skilled adventurer they could find: you. After
    listening to their tale of woe, you agree to enter the labyrinth where most of the creatures seem to originate,
    and destroy the foul beasts. Armed with a longsword and a bundle of torches, you descend into the labyrinth, 
    ready to do battle....
	`)
}

func generateRoom() *models.Room {

	room := models.NewRoom()

	// there is a 25% chance that this room has an item
	if rand.IntN(100) < 26 {
		a := models.GetAllArmory()
		item := a[rand.IntN(len(a))]
		room.Items = append(room.Items, item)
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
		room := generateRoom()
		currentGame.Room = *room
		currentGame.Room.PrintDescription()

		for _, item := range room.Items {
			yellow.Printf("You see a %s\n", item.Name)
		}

		if currentGame.Room.Monster != nil {
			red.Printf("There is a %s here!\n", currentGame.Room.Monster.Name)
		}

		var input string
		yellow.Print("-> ")
		fmt.Scanln(&input)

		if input == "quit" {
			green.Println("Overcome with terror, you flee the dungeon.")
			playAgain()
		} else if slices.Contains(directions, input) {
			cyan.Println("You move deeper into the dungeon.")
			continue
		} else if input == "help" {
			showHelp()
		} else {
			red.Println("I'm not sure what you mean... type help for available commands.")
		}
	}
}

func playAgain() {
	yn := getYN("Play again")
	if yn == "yes" {
		PlayGame()
	} else {
		green.Println("Until next time, adventurer.")
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
			cyan.Println("Please enter yes or no.")
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
    - n/s/e/w : move in a direction
    - map : show a map of the labyrinth
    - look : look around and describe you environment
    - equip <item> : use an item from your inventory
    - unequip <item> : stop using an item from your inventory
    - fight : attack a foe
    - examine <object> : examine an object more closely
    - get <item> : pick up an item
    - drop <item> : drop an item
    - rest : restore health by resting
    - inventory : show current inventory
    - status : show current player status
    - quit : end the game`)
}
