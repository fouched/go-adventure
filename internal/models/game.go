package models

import "github.com/fatih/color"

var red = color.New(color.FgRed)
var green = color.New(color.FgGreen)
var yellow = color.New(color.FgYellow)
var cyan = color.New(color.FgCyan)

type Game struct {
	Player      Player
	Room        Room
	NumMonsters int
	Rooms       map[string]Room
	X           int
	Y           int
	Entrance    string
}

func NewGame(p Player, x int, y int) *Game {
	return &Game{
		Player:      p,
		NumMonsters: 0,
		Rooms:       make(map[string]Room),
		X:           x,
		Y:           y,
	}
}
