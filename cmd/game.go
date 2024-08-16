package main

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/fouched/go-adventure/internal/models"
	"os"
	"strings"
)

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var cyan = color.New(color.FgCyan)

func PlayGame() {

	adventurer := models.NewPlayer()
	currentGame := models.NewGame(*adventurer)
	room := models.NewRoom()
	currentGame.Room = *room

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

func exploreLabyrinth(currentGame *models.Game) {

	for {
		currentGame.Room.PrintDescription()

		var input string
		yellow.Print("-> ")
		fmt.Scanln(&input)

		if input == "quit" {
			green.Println("Overcome with terror, you flee the dungeon.")
			playAgain()
		} else if input == "help" {
			showHelp()
		} else {
			green.Println("I'm not sure what you mean... type help for available commands.")
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
