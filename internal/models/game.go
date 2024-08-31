package models

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
