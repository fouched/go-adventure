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
}

func NewGame(p Player) *Game {
	return &Game{
		Player:      p,
		NumMonsters: 0,
		Rooms:       make(map[string]Room),
		X:           0,
		Y:           0,
	}
}
