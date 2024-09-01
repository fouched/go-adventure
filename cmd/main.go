package main

import (
	"github.com/fatih/color"
	"github.com/fouched/go-adventure/internal/config"
	"github.com/fouched/go-adventure/internal/models"
	"os"
	"os/exec"
	"runtime"
)

// create a map for storing clear funcs
var clear map[string]func()

func main() {
	// go run github.com/fouched/go-adventure/cmd
	// go build -o ./tmp/main.exe github.com/fouched/go-adventure/cmd
	PlayGame()
}

// init runs before main and initializes clear functions
func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func ClearScreen() {
	//runtime.GOOS -> linux, windows, darwin etc.
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

func drawUI(currentGame *models.Game) {
	ClearScreen()
	black := color.New(color.FgBlack)
	greenBg := black.Add(color.BgGreen)
	// The village of Honeywood has been terrorized by strange, deadly creatures for months now. Unable to endure an
	greenBg.Printf("     Health: %d/%d                          Monsters defeated: %d/%d                                      XP: %d     \n\n",
		currentGame.Player.HP, config.PLAYER_HP, currentGame.Player.MonstersDefeated, currentGame.NumMonsters, currentGame.Player.XP)
}
